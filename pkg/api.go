package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Err - Define pattern error return.
type Err struct {
	code     string `json:"code"`
	message  string `json:"message"`
	err      error  `json:"err"`
	msgParam map[string]string
}

//NewErr - New Err set properties with builder function.
func NewErr() *Err {
	return &Err{}
}

//WithCode - Set Code.
func (e *Err) WithCode(code string) *Err {
	e.code = code
	return e
}

//GetCode - Code internationalization message.
func (e *Err) GetCode() string {
	return e.code
}

//WithMessage - Set Message.
func (e *Err) WithMessage(msg string) *Err {
	e.message = msg
	return e
}

//WithMessagef - Set Message.
func (e *Err) WithMessagef(msg string, a ...any) *Err {
	e.message = fmt.Sprintf(msg, a...)
	return e
}

//WithError - Set Err.
func (e *Err) WithError(error error) *Err {
	e.err = error
	return e
}

//WithMsgParam - Set params.
func (e *Err) WithMsgParam(param map[string]string) *Err {
	e.msgParam = param
	return e
}

//GetMsgParam - Return msg params.
func (e *Err) GetMsgParam() map[string]string {
	return e.msgParam
}

//WithErr - Build Err  with Err.
func (e *Err) WithErr(k *Err) *Err {
	e.code = k.code
	e.message = k.message
	e.msgParam = k.msgParam
	e.err = k.err
	return e
}

//OnError - Convert Err type to error type.
func (e *Err) OnError() error {
	if e == nil {
		return nil
	}
	return errors.New(fmt.Sprintf("%s - %s", e.code, e.message))
}

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
func parseJSON(w http.ResponseWriter, body []byte, entity interface{}) *Err {
	if err := json.Unmarshal(body, entity); err != nil {
		log.Println(err)
		err := NewErr().WithMessage("Error parsing json request.").WithError(err)
		responseJson(w, http.StatusUnprocessableEntity, err)
		return err
	}
	return nil
}

//findURLParam - Find by URL parameter.
func findURLParam(r *http.Request, keys []string) (map[string]string, *Err) {
	result := make(map[string]string)
	vars := mux.Vars(r)
	for i := 0; i < len(keys); i += 2 {
		if k := vars[keys[i]]; k == "" {
			return nil, NewErr().WithMessage("ERRO_PARAM_URL " + keys[i+1])
		} else {
			result[keys[i]] = k
		}
	}
	return result, nil
}

//findURLParam - Find by integer URL parameter.
func findURLParamInt(r *http.Request, keys []string) (keyResolv map[string]int, erro *Err) {
	result := make(map[string]int)
	vars := mux.Vars(r)
	for i := 0; i < len(keys); i += 2 {
		if k, err := strconv.Atoi(vars[keys[i]]); err != nil {
			return nil, NewErr().WithMessage("ERRO_PARAM_URL " + keys[i+1]).WithError(err)
		} else {
			result[keys[i]] = k
		}
	}
	return result, nil
}
