package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabrielbo1/iroko/config"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/connect"
)

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, word!")
}

func main() {
	http.HandleFunc("/", helloServer)

	if config.EnvironmentVariableValue(config.ConsulActive) == "true" {
		// build client
		configConsul := api.DefaultConfig()
		configConsul.Address = config.EnvironmentVariableValue(config.ConsulAddress)
		configConsul.Address += ":" + config.EnvironmentVariableValue(config.ConsulPort)

		clientConsul, err := api.NewClient(configConsul)
		if err != nil {
			log.Fatalf("client err: %v", err)
		}

		svc, _ := connect.NewService("iroko-service", clientConsul)
		defer svc.Close()

		// Creating an HTTP server that serves via Connect
		server := &http.Server{
			Addr:      ":" + config.EnvironmentVariableValue(config.Port),
			TLSConfig: svc.ServerTLSConfig(),
		}
		server.ListenAndServeTLS("", "")
	}

	http.ListenAndServe(":"+config.EnvironmentVariableValue(config.Port), nil)
}
