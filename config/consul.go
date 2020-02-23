package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
	uuid "github.com/hashicorp/go-uuid"
)

const (
	// TTLInterval - .
	TTLInterval = time.Second * 15
	// TTLRefreshInterval - .
	TTLRefreshInterval = time.Second * 10
	// TTLDeregisterCriticalServiceAfter - .
	TTLDeregisterCriticalServiceAfter = time.Minute
)

// ConsulStart - Checks if the Consul is enabled.
// If you are registering for a service at the
// consul following the past settings or with the
// standard environment settings.
func ConsulStart(doneChan chan struct{}) {
	if EnvironmentVariableValue(ConsulActive) == "true" {
		time.Sleep(time.Second * 10)

		// build client
		client, err := api.NewClient(&api.Config{
			Address: EnvironmentVariableValue(ConsulAddress) + ":" + EnvironmentVariableValue(ConsulPort),
			Scheme:  "http",
		})

		if err != nil {
			panic(err)
		}

		address := EnvironmentVariableValue(AddressInstance)

		//Random port with Consul
		port, _ := strconv.Atoi(EnvironmentVariableValue(RandomFreePort))
		os.Setenv(string(Port), strconv.Itoa(port))
		log.Printf("Random port application = %v", port)

		// Unic ID application.
		id, _ := uuid.GenerateUUID()

		err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
			Address: address,
			Port:    port,
			ID:      id,      // Unique for each node
			Name:    "iroko", // Can be service type
			Tags:    []string{"iroko"},
			Check: &api.AgentServiceCheck{
				HTTP:     "http://" + address + ":" + strconv.Itoa(port) + "/_health",
				Interval: "60s",
			},
		})

		if err != nil {
			panic(err)
		}

		sessionID, _, err := client.Session().Create(&api.SessionEntry{
			Name:     "service/monitoring/leader", // distributed lock
			Behavior: "delete",
			TTL:      "10s",
		}, nil)

		if err != nil {
			panic(err)
		}

		isLeader, _, err := client.KV().Acquire(&api.KVPair{
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
			client.Session().RenewPeriodic(
				"90s",
				sessionID,
				nil,
				doneChan,
			)
		}()

		log.Printf("Consul Active = %v.", isLeader)
	}
}
