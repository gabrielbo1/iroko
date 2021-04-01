package config

import (
	"flag"
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

type confiVar struct {
	name  EnvironmentVariable
	value string
	usage string
}

var configVars []confiVar = []confiVar{
	{name: AppName, value: "iroko", usage: "App define continaier image name and PathPrefix to API URLs."},
	{name: HealthCheckPath, value: "/_health", usage: "Health Chech Path."},
	{name: Port, value: "7070", usage: "Server port application."},
	{name: ConsulActive, value: "false", usage: "Case true subscribe consul. Ex.: false"},
	{name: ConsulAddress, value: "127.0.0.1", usage: "Consul IP address or DNS. Ex.: 127.0.0.1"},
	{name: ConsulPort, value: "8500", usage: "Consul default port. Ex.: 8500"},
	{name: Base, value: "POSTGRESQL", usage: "Define data base app. Ex.: POSTGRESQL"},
	{name: BaseName, value: "iroko", usage: "Define base name. Ex.: iroko"},
	{name: BaseAddress, value: "127.0.0.1", usage: "Data base IP address or DNS. Ex.: 127.0.0.1"},
	{name: BasePort, value: "5432", usage: "Data base port. Ex.: 5432"},
	{name: BaseUser, value: "postgres", usage: "User data base user. Ex.: postgres"},
	{name: BasePassword, value: "123456", usage: "Data base password. Ex.: 123456"},
	{name: BaseSSL, value: "disable", usage: "Data base SSL mode(View PostgreSQL ssl configuration). Ex.:  require, verify-full, verify-ca and disable"},
	{name: AddressInstance, value: "", usage: "Ip application, case no set defined automatic. Ex.: 127.0.0.1"},
	{name: RandomFreePort, value: "", usage: " RandomFreePort - free PORT, case not set defined automatic, system use free port with consul configuration. E.: 5656"},
	{name: Domain, value: "localhost", usage: "DNS application or IP. Ex.: localhost"},
}

func setVar(envVar EnvironmentVariable, value string) {
	for i := range configVars {
		if configVars[i].name == envVar {
			configVars[i].value = value
		}
	}
}

func getVar(envVar EnvironmentVariable) *confiVar {
	for i := range configVars {
		if configVars[i].name == envVar {
			return &configVars[i]
		}
	}
	return nil
}

func setAddressInstance(envVar EnvironmentVariable) {
	if envVar == AddressInstance && getVar(AddressInstance).value == "" {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			panic(err)
		}
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					getVar(AddressInstance).value = ipnet.IP.String()
					return
				}
			}
		}
		panic(fmt.Errorf("Impossible to determinate IP Address"))
	}
}

func setRandomPort(envVar EnvironmentVariable) {
	if envVar == RandomFreePort && getVar(RandomFreePort).value == "" {
		rand.Seed(time.Now().UnixNano())
		getVar(RandomFreePort).value = strconv.Itoa(rand.Intn(20000-10000) + 10000)
	}
}

//FlagParse - Flags parsing and set values.
func FlagParse() {
	var values []*string
	if !flag.Parsed() {
		for i := range configVars {
			values = append(values, flag.String(string(configVars[i].name), configVars[i].value, configVars[i].usage))
		}
		flag.Parse()
	}
	for i := range configVars {
		configVars[i].value = *values[i]
	}
}

// EnvironmentVariableValue - Find to environment variable value
// or return default value of variable.
func EnvironmentVariableValue(variable EnvironmentVariable) string {
	//Set statics variables.
	setAddressInstance(variable)
	setRandomPort(variable)
	if value := os.Getenv(string(variable)); value != "" {
		setVar(variable, value)
	}
	if conf := getVar(variable); conf != nil {
		return conf.value
	}
	return ""
}
