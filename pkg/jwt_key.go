package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type JwtValueKey struct {
	Key  string    `json:"key"`
	Date time.Time `json:"date"`
}

var key *JwtValueKey

func newJwtValueKey() *JwtValueKey {
	singKey, _ := GenerateRandomString(128)
	return &JwtValueKey{singKey, time.Now().Add(time.Duration(24 * time.Hour))}
}

func getJwtConsulVariable() bool {
	if consulJson, err := GetConsulVariable(JwtKey); err != nil && consulJson != nil {
		consulKey := JwtValueKey{}
		json.Unmarshal(consulJson, consulKey)
		*key = consulKey
		return true
	}
	return false
}

func notifyAllConsul() {
	if services, err := consulClient.Agent().Services(); err == nil {
		for _, service := range services {
			url := fmt.Sprintf("http://%s:%d"+ConfigVars.EnvironmentVariableValue(ConsulJWTPath), service.Address, service.Port)
			http.Get(url)
		}
	}
}

func GetSignJwtKey() []byte {
	if key != nil && key.Date.After(time.Now()) {
		return []byte(key.Key)
	}
	key = newJwtValueKey()
	if ConsulOk() {
		if !getJwtConsulVariable() && key.Date.Before(time.Now()) {
			keyJson, _ := json.Marshal(*key)
			PutConsulVariable(JwtKey, string(keyJson))
			notifyAllConsul()
		}
	}
	return []byte(key.Key)
}

func UpdateConsulJwt(w http.ResponseWriter, r *http.Request) {
	if ConsulOk() {
		getJwtConsulVariable()
	}
	log.Println("CONSUL JWT UPDATE TOKEN")
	fmt.Fprint(w, "OK")
}
