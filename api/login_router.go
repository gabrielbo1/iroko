package api

import (
	"github.com/gabrielbo1/iroko/domain"
	"github.com/gabrielbo1/iroko/service"
	"net/http"
)

//PostLogin - Create new login.
func PostLogin(w http.ResponseWriter, r *http.Request) {
	var login domain.Login
	onParserLogin(w, readBody(r), login, parseJSON, service.CreateLogin)
}

//PutLogin - Update login.
func PutLogin(w http.ResponseWriter, r *http.Request) {
	var login domain.Login
	onParserLogin(w, readBody(r), login, parseJSON, service.UpdateLogin)
}

//DeleteLogin - Delete login.
func DeleteLogin(w http.ResponseWriter, r *http.Request) {
	var errDomain *domain.Err
	var paramInt map[string]int

	if paramInt, errDomain = findURLParamInt(r, []string{"id", "Id login entity invalid or empty."}); errDomain != nil {
		if errDomain = service.DeleteLogin(paramInt["id"]); errDomain != nil {
			responseJson(w, http.StatusOK, simpleResponsePattern{"Login record deleted success."})
			return
		}
	}
	responseJson(w, http.StatusBadRequest, errDomain)
}

func onParserLogin(w http.ResponseWriter, body []byte, entity domain.Login,
	parseFunc func(w http.ResponseWriter, body []byte, entity interface{}) *domain.Err,
	service func(entity *domain.Login) *domain.Err) {
	var err *domain.Err
	if err = parseFunc(w, body, &entity); err == nil {
		if err = service(&entity); err == nil {
			responseJson(w, http.StatusOK, entity)
		} else {
			responseJson(w, http.StatusBadRequest, err)
		}
	}
}
