package api

import (
	"encoding/json"
	"github.com/gabrielbo1/iroko/domain"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

//simpleResponsePattern - Simple response pattern.
type simpleResponsePattern struct {
	Message string `json:"message"`
}

//responseJson - Parse entity json and write bytes json in response.
func responseJson(w http.ResponseWriter, code int, entity interface{}) {
	response, _ := json.Marshal(entity)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(response)
}

//readBody - Read request bytes.
func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	return body
}

//parseJSON - Parse JSON request.
func parseJSON(w http.ResponseWriter, body []byte, entity interface{}) *domain.Err {
	if err := json.Unmarshal(body, entity); err != nil {
		log.Println(err)
		err := domain.NewErr().WithMessage("Error parsing json request.").WithError(err)
		responseJson(w, http.StatusUnprocessableEntity, err)
		return err
	}
	return nil
}

//findURLParam - Find by URL parameter.
func findURLParam(r *http.Request, keys []string) (map[string]string, *domain.Err) {
	result := make(map[string]string)
	vars := mux.Vars(r)
	for i := 0; i < len(keys); i += 2 {
		if k := vars[keys[i]]; k == "" {
			return nil, domain.NewErr().WithMessage("ERRO_PARAM_URL " + keys[i+1])
		} else {
			result[keys[i]] = k
		}
	}
	return result, nil
}

//findURLParam - Find by integer URL parameter.
func findURLParamInt(r *http.Request, keys []string) (keyResolv map[string]int, erro *domain.Err) {
	result := make(map[string]int)
	vars := mux.Vars(r)
	for i := 0; i < len(keys); i += 2 {
		if k, err := strconv.Atoi(vars[keys[i]]); err != nil {
			return nil, domain.NewErr().WithMessage("ERRO_PARAM_URL " + keys[i+1]).WithError(err)
		} else {
			result[keys[i]] = k
		}
	}
	return result, nil
}
