package main

import (
	"fmt"
	"net/http"

	"github.com/gabrielbo1/iroko/config"
)

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, word!")
}

func main() {
	http.HandleFunc("/", helloServer)
	http.ListenAndServe(":"+config.EnvironmentVariableValue(config.Port), nil)
}
