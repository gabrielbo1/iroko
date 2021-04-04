package main

import (
	"fmt"
	"github.com/gabrielbo1/iroko/pkg"
	"net/http"
	"time"

	"github.com/gabrielbo1/iroko/api"
	"github.com/gabrielbo1/iroko/infrastructure/repository"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func init() {
	// Only log the warning severity or above.
	log.SetLevel(log.FatalLevel)

	//Flag parsing
	pkg.NewVars()
	pkg.ConfigVars.FlagParse()
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK %s:%s", pkg.ConfigVars.EnvironmentVariableValue(pkg.AppName), pkg.ConfigVars.EnvironmentVariableValue(pkg.Port))
}

func main() {
	//Start migration.
	repository.MigrationInit()

	doneChan := make(chan struct{})
	defer close(doneChan)

	pkg.ConsulStart(doneChan)

	router := api.NewRouter()
	router.HandleFunc(pkg.ConfigVars.EnvironmentVariableValue(pkg.HealthCheckPath), health)
	router.HandleFunc(pkg.ConfigVars.EnvironmentVariableValue(pkg.ConsulJWTPath), pkg.UpdateConsulJwt)
	http.Handle("/", router)

	n := negroni.Classic()
	n.UseHandler(cors.AllowAll().Handler(router))
	n.Run(":" + pkg.ConfigVars.EnvironmentVariableValue(pkg.Port))

	time.Sleep(time.Second * 90)
}
