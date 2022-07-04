package pkg

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
	//ConsulJWTUpdatePath - JWT update notification.
	ConsulJWTPath EnvironmentVariable = "CONSUL_JWT_NOTIFY_PATH"
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
	// Modules - Define active modules EX.: MODULE_A, MODULE_B
	Modules EnvironmentVariable = "MODULES"
)

//ConfiVar - Template app environment variable.
type ConfiVar struct {
	name  EnvironmentVariable
	value string
	usage string
}

//Vars - Array with variables defined.
type Vars struct {
	Vars []ConfiVar
}

var ConfigVars *Vars

//NewVars - Create a new environment variables with default values, set custom default values.
func NewVars() {
	if ConfigVars == nil {
		ConfigVars = &Vars{
			[]ConfiVar{
				{name: AppName, value: "iroko", usage: "App define continaier image name and PathPrefix to API URLs."},
				{name: HealthCheckPath, value: "/_health", usage: "Health Check Path."},
				{name: ConsulJWTPath, value: "/_consulJwt", usage: "JWT notify get JWT key."},
				{name: Port, value: "9090", usage: "Server port application."},
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
				{name: Modules, value: "MODULE_A", usage: "Define active modules EX.: MODULE_A, MODULE_B"},
			},
		}
	}
}

//SetVar - Set default environment variables value.
func (vars *Vars) SetVar(envVar EnvironmentVariable, value string) {
	for i := range vars.Vars {
		if vars.Vars[i].name == envVar {
			vars.Vars[i].value = value
		}
	}
}

//GetVar - Get environment variable object.
func (vars *Vars) GetVar(envVar EnvironmentVariable) *ConfiVar {
	for i := range vars.Vars {
		if vars.Vars[i].name == envVar {
			return &vars.Vars[i]
		}
	}
	return nil
}

//setAddressInstance - Configuration to IP address app.
func (vars *Vars) setAddressInstance(envVar EnvironmentVariable) {
	if envVar == AddressInstance && vars.GetVar(AddressInstance).value == "" {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			panic(err)
		}
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					vars.GetVar(AddressInstance).value = ipnet.IP.String()
					return
				}
			}
		}
		panic(fmt.Errorf("%s", "Impossible to determinate IP Address"))
	}
}

//setRandomPort - Define random port app case Port variable is empty.
func (vars *Vars) setRandomPort(envVar EnvironmentVariable) {
	if envVar == RandomFreePort && vars.GetVar(Port).value != "" {
		vars.GetVar(RandomFreePort).value = vars.GetVar(Port).value
	}
	if envVar == RandomFreePort && vars.GetVar(RandomFreePort).value == "" {
		rand.Seed(time.Now().UnixNano())
		vars.GetVar(RandomFreePort).value = strconv.Itoa(rand.Intn(20000-10000) + 10000)
		vars.GetVar(Port).value = vars.GetVar(RandomFreePort).value
	}
}

//FlagParse - Flags variables parsing and operational system variables parsing, set values.
func (vars *Vars) FlagParse() {
	//Flags parse
	if !flag.Parsed() {
		var flagValues []*string
		for i := range vars.Vars {
			flagValues = append(flagValues, flag.String(string(vars.Vars[i].name), vars.Vars[i].value, vars.Vars[i].usage))
		}
		flag.Parse()
		for i := range vars.Vars {
			vars.Vars[i].value = *flagValues[i]
		}
	}

	//Operational System parse
	for i := range vars.Vars {
		if value := os.Getenv(string(vars.Vars[i].name)); value != "" {
			vars.Vars[i].value = value
		}
	}
}

// EnvironmentVariableValue - Find to environment variable value
// or return default value of variable.
func (vars *Vars) EnvironmentVariableValue(variable EnvironmentVariable) string {
	//Set statics variables.
	vars.setAddressInstance(variable)
	vars.setRandomPort(variable)
	if value := vars.GetVar(variable); value != nil {
		return value.value
	}
	return ""
}
