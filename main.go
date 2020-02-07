package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gabrielbo1/iroko/config"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, word!")
}

func main() {
	http.HandleFunc("/", helloServer)
	http.HandleFunc("/_health", health)

	doneChan := make(chan struct{})
	defer close(doneChan)

	config.ConsulStart(doneChan)
	http.ListenAndServe(":"+config.EnvironmentVariableValue(config.Port), nil)
	time.Sleep(time.Second * 90)
}
