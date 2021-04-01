package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gabrielbo1/iroko/api"
	"github.com/gabrielbo1/iroko/config"
	"github.com/gabrielbo1/iroko/infrastructure/repository"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func init() {
	// Only log the warning severity or above.
	log.SetLevel(log.FatalLevel)

	//Flag parsing
	config.FlagParse()
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func main() {
	//Start migration.
	repository.MigrationInit()

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
