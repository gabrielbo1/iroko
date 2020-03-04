package config

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

// EnvironmentVariable - Especific type to environment variables.
type EnvironmentVariable string

const (
	//AppName - App define continaier image name and PathPrefix to API URLs.
	AppName EnvironmentVariable = "APP_NAME"
	//HealthCheckPath - Health Chech Path.
	HealthCheckPath EnvironmentVariable = "HEALTH_CHECH_PATH"
	// Port -  Server port application.
	Port EnvironmentVariable = "PORT"
	// ConsulActive - Case true subscribe consul.
	ConsulActive = "CONSUL_ACTIVE"
	// ConsulAddress - Consul IP address or DNS.
	ConsulAddress = "CONSUL_ADDRESS"
	// ConsulPort - Consul subscribe port.
	ConsulPort = "CONSUL_PORT"
	// PostgresAddress - PostgreSQL IP  address or DNS.
	PostgresAddress = "POSTGRES_ADDRESS"
	// PostgresPort - PostgreSQL port.
	PostgresPort = "POSTGRES_PORT"
	// PostgresUser - PostgreSQL user.
	PostgresUser = "POSTGRES_USER"
	// PostgresPassword - PostgreSQL user.
	PostgresPassword = "POSTGRES_PASSWORD"
	// AddressInstance - Ip application.
	AddressInstance = "ADDRESS_INSTANCE"
	// RandomFreePort - free PORT.
	RandomFreePort = "RandomFreePort"
	// Domain - DNS application or IP.
	Domain = "DOMAIN"
)

func getVariable(name EnvironmentVariable, defaultValue string) string {
	if variable := os.Getenv(string(name)); variable != "" {
		return variable
	}
	return defaultValue
}

// EnvironmentVariableValue - Find to environment variable value
// or return default value of variable.
func EnvironmentVariableValue(variable EnvironmentVariable) string {
	switch variable {
	case AppName:
		return getVariable(AppName, "iroko")
	case HealthCheckPath:
		return getVariable(HealthCheckPath, "/_health")
	case Port:
		return getVariable(Port, "7070")
	case ConsulActive:
		return getVariable(ConsulActive, "false")
	case ConsulAddress:
		return getVariable(ConsulAddress, "127.0.0.1")
	case ConsulPort:
		return getVariable(ConsulPort, "8500")
	case PostgresAddress:
		return getVariable(PostgresAddress, "127.0.0.1")
	case PostgresPort:
		return getVariable(PostgresPort, "5432")
	case PostgresUser:
		return getVariable(PostgresUser, "postgres")
	case PostgresPassword:
		return getVariable(PostgresPassword, "123456")
	case AddressInstance:
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			panic(err)
		}
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.String()
				}
			}
		}
		panic(fmt.Errorf("Impossible to determinate IP Address"))
	case RandomFreePort:
		rand.Seed(time.Now().UnixNano())
		return strconv.Itoa(rand.Intn(20000-10000) + 10000)
	case Domain:
		return getVariable(Domain, "localhost")
	}
	return ""
}
