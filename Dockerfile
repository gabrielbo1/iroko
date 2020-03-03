FROM  golang:1.13.4 as builder

LABEL maintainer="Gabriel Oliveira <barbosa.olivera1@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Fim Testes Inicio BUILD.
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#BUILD WEB
FROM node:10-alpine as builderweb

# Saving libraries to different layers avoids unnecessary downloads.
COPY /website/iroko-app/package.json ./
COPY /website/iroko-app/package-lock.json ./
RUN npm ci && mkdir /iroko-app && mv ./node_modules ./iroko-app
WORKDIR /iroko-app
COPY /website/iroko-app/ .

#Build Web
RUN npm run ng build -- --prod --output-path=dist

######## Start a new stage from scratch #######
FROM alpine:3.10.3
RUN apk --no-cache add ca-certificates
WORKDIR /root/public/
COPY  --from=builderweb /iroko-app/dist .
WORKDIR /root/
COPY --from=builder /app/main .
ENV TZ=America/Sao_Paulo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
CMD ["./main"]