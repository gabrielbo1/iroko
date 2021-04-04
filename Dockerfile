FROM  golang:1.13.4 as builder

LABEL maintainer="Gabriel Oliveira <barbosa.olivera1@gmail.com>"
WORKDIR /app
RUN go get -u github.com/nicksnyder/go-i18n/v2/goi18n
COPY go.mod go.sum ./
RUN go mod download
COPY . .

#TEST AND COVERAGE
RUN go test ./... -cover

# Fim Testes Inicio BUILD.
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#BUILD WEB
FROM node:15.13.0-alpine3.13 as builderweb

#Phyton
#RUN apk add --update nodejs bash python make g++ git
#dockRUN python --version

# Saving libraries to different layers avoids unnecessary downloads.
COPY /website/iroko-app/package.json ./
COPY /website/iroko-app/package-lock.json ./
RUN npm install --cache /tmp/empty-Fcache
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
WORKDIR /root/infrastructure/repository/postgres/migration
ADD ./infrastructure/repository/postgres/migration .
ENV TZ=America/Sao_Paulo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
WORKDIR /root/
RUN chmod a+x main
CMD ["./main"]
