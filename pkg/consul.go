package pkg

import (
	"log"
	"os"
	"strconv"

	"github.com/hashicorp/consul/api"
)

var consulActive bool
var consulClient *api.Client

// ConsulStart - Checks if the Consul is enabled.
// If you are registering for a service at the
// consul following the past settings or with the
// standard environment settings.
func ConsulStart(doneChan chan struct{}) {
	if ConfigVars.EnvironmentVariableValue(ConsulActive) == "true" {
		// build consulClient
		c, err := api.NewClient(&api.Config{
			Address: ConfigVars.EnvironmentVariableValue(ConsulAddress) + ":" + ConfigVars.EnvironmentVariableValue(ConsulPort),
			Scheme:  "http",
		})

		if err != nil {
			panic(err)
		}
		consulClient = c

		address := ConfigVars.EnvironmentVariableValue(AddressInstance)

		//Random port with Consul
		port, _ := strconv.Atoi(ConfigVars.EnvironmentVariableValue(RandomFreePort))
		err = os.Setenv(string(Port), strconv.Itoa(port))
		if err != nil {
			panic(err)
		}
		log.Printf("Random port application = %v", port)

		// Unic ID application.
		id := ConfigVars.EnvironmentVariableValue(AppName) + ":" + ConfigVars.EnvironmentVariableValue(RandomFreePort)

		err = consulClient.Agent().ServiceRegister(&api.AgentServiceRegistration{
			Address: address,
			Port:    port,
			ID:      id,                                           // Unique for each node
			Name:    ConfigVars.EnvironmentVariableValue(AppName), // Can be service type
			Tags:    []string{"primary"},
			Check: &api.AgentServiceCheck{
				HTTP:     "http://" + address + ":" + strconv.Itoa(port) + "/_health",
				Interval: "60s",
			},
		})

		if err != nil {
			panic(err)
		}

		sessionID, _, err := consulClient.Session().Create(&api.SessionEntry{
			Name:     "service/monitoring/leader", // distributed lock
			Behavior: "delete",
			TTL:      "10s",
		}, nil)

		if err != nil {
			panic(err)
		}

		isLeader, _, err := consulClient.KV().Acquire(&api.KVPair{
			Key:     "service/monitoring/leader", // distributed lock
			Value:   []byte(sessionID),
			Session: sessionID,
		}, nil)

		if err != nil {
			panic(err)
		}

		go func() {
			// RenewPeriodic is used to periodically invoke Session.Renew on a
			// session until a doneChan is closed. This is meant to be used in a long running
			// goroutine to ensure a session stays valid.
			consulClient.Session().RenewPeriodic(
				"90s",
				sessionID,
				nil,
				doneChan,
			)
		}()

		consulActive = true
		log.Printf("Consul Active = %v.", isLeader)
	}
}

// ConsulOk - True case consul OK.
func ConsulOk() bool {
	return consulActive
}

// ConsulVariable - Especific type consul
// key variables.
type ConsulVariable string

const (
	//JwtKey - The signing key JWT shares with other instances of iroko (Gateway Pattern).
	JwtKey ConsulVariable = "IROKO_JWT_KEY"
)

// PutConsulVariable - Put simple variabe in Consul.
func PutConsulVariable(variable ConsulVariable, value string) error {
	var keyPair *api.KVPair = &api.KVPair{}
	keyPair.Key = string(variable)
	keyPair.Value = []byte(value)

	if _, e := consulClient.KV().Put(keyPair, nil); e != nil {
		return e
	}
	return nil
}

// GetConsulVariable - Retrieves simple consul variable.
func GetConsulVariable(variable ConsulVariable) ([]byte, error) {
	kp, _, err := consulClient.KV().Get(string(variable), nil)
	if err != nil {
		return nil, err
	}
	if kp == nil {
		return nil, nil
	}
	return kp.Value, nil
}
