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
	//Base - Define data base app. Ex.: POSTGRESQL
	Base EnvironmentVariable = "BASE"
	//BaseName - Define base name.
	BaseName EnvironmentVariable = "BASE_NAME"
	// BaseAddress - Data base IP address or DNS.
	BaseAddress = "BASE_ADDRESS"
	// BasePort - Data base port.
	BasePort = "BASE_PORT"
	// BaseUser - User data base user.
	BaseUser = "BASE_USER"
	// BasePassword - Data base password.
	BasePassword = "BASE_PASSWORD"
	// BaseSSL - Data base SSL mode(View PostgreSQL ssl configuration).
	BaseSSL = "BASE_SSL"
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
	case Base:
		return getVariable(Base, "POSTGRESQL")
	case BaseName:
		return getVariable(BaseName, "iroko")
	case BaseAddress:
		return getVariable(BaseAddress, "127.0.0.1")
	case BasePort:
		return getVariable(BasePort, "5432")
	case BaseUser:
		return getVariable(BaseUser, "postgres")
	case BasePassword:
		return getVariable(BasePassword, "123456")
	case BaseSSL:
		return getVariable(BaseSSL, "disable")
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
