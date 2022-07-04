package pkg

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gorilla/mux"
)

// NewRouter - Returns all APIs implemented and mapped in rotes.go
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		if route.Name == "Login" ||
			route.Name == "Refresh" {
			router.
				Methods(route.Method).
				Path("/api" + route.Pattern).
				Handler(route.HandlerFunc).
				Name(route.Name)
			continue
		}

		var authHandler http.Handler
		authHandler = route.HandlerFunc
		authHandler = authenticateHandler(authHandler)

		router.
			Methods(route.Method).
			Path("/api" + route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc).
			Handler(authHandler)
	}
	return router
}

// JWT-based authentication model,
// following tutorial available at: https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/
func authenticateHandler(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, _, ok := getTokenAndValid(w, r); !ok {
			return
		}
		inner.ServeHTTP(w, r)
	})
}

// Claims Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string   `json:"username"`
	Modules  []string `json:"modules"`
	jwt.StandardClaims
}

// Login - Loging and get token.
func Login(w http.ResponseWriter, r *http.Request) {
	username, password, authOK := r.BasicAuth()
	if authOK == false {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	valid, modules := validUser(username, password)

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	moment := time.Now()
	expirationTime := moment.Add(5 * time.Minute)

	claims := &Claims{
		Username: username,
		Modules:  modules,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  moment.Unix(),
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string with SingKey.
	tokenString, err := token.SignedString(GetSignJwtKey())
	if !valid {
		tokenString, err = token.SignedString([]byte("INVALID_SINGNATURE"))
	}

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set token and returning client.
	w.Header().Add("Authorization", tokenString)
}

// Refresh - Refresh user token.
func Refresh(w http.ResponseWriter, r *http.Request) {
	if _, claims, ok := getTokenAndValid(w, r); ok {
		// We ensure that a new token is not issued until enough time has elapsed
		// In this case, a new token will only be issued if the old token is within
		// 30 seconds of expiry. Otherwise, return a bad request status
		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Now, create a new token for the current use, with a renewed expiration time
		claims.ExpiresAt = time.Now().Add(5 * time.Minute).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(GetSignJwtKey())
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Authorization", tokenString)
	}
}

// Test - Test valid token.
func Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func getTokenAndValid(w http.ResponseWriter, r *http.Request) (*jwt.Token, *Claims, bool) {
	signingKey := GetSignJwtKey()
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(r.Header.Get("Authorization"), claims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil || !token.Valid {
		if (err != nil && err.Error() == jwt.ErrSignatureInvalid.Error()) || !token.Valid {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		} else {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		return nil, nil, false
	}

	return token, claims, true
}

// Returning user valid and modules user acess.
func validUser(username, passowrd string) (bool, []string) {
	return username == "gabriel" && passowrd == "12345678",
		[]string{ConfigVars.EnvironmentVariableValue(AppName)}
}
