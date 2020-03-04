package main

import (
	"fmt"
	"github.com/gabrielbo1/iroko/api"
	"github.com/gabrielbo1/iroko/config"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
	"time"
)

func init() {
	// Only log the warning severity or above.
	log.SetLevel(log.FatalLevel)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func main() {

	doneChan := make(chan struct{})
	defer close(doneChan)

	config.ConsulStart(doneChan)

	router := api.NewRouter()
	router.HandleFunc(config.EnvironmentVariableValue(config.HealthCheckPath), health)
	http.Handle("/", router)

	n := negroni.Classic()
	n.UseHandler(cors.AllowAll().Handler(router))
	n.Run(":" + config.EnvironmentVariableValue(config.Port))

	time.Sleep(time.Second * 90)
}
