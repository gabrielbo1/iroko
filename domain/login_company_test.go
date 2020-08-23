package domain

import "testing"

func assertValidLoginCompany(loginCmp *LoginCompany, t *testing.T, errCode string) {
	if _, err := ValidLoginCompany(loginCmp); err == nil || err.GetCode() != errCode {
		t.Error("Error code entity login company, code " + errCode)
	}
}

func TestValidLoginCompany(t *testing.T) {
	loginCmp := &LoginCompany{}

	assertValidLoginCompany(loginCmp, t, "logincompany10")
	loginCmp.Login = ""
	assertValidLoginCompany(loginCmp, t, "logincompany10")
	loginCmp.Login = "user"
	assertValidLoginCompany(loginCmp, t, "logincompany20")
	loginCmp.Password = ""
	assertValidLoginCompany(loginCmp, t, "logincompany20")
	loginCmp.Password = "pass"
	assertValidLoginCompany(loginCmp, t, "logincompany30")
	loginCmp.Company = Company{}
	assertValidLoginCompany(loginCmp, t, "logincompany30")
	loginCmp.Company.ID = 1
	assertValidLoginCompany(loginCmp, t, "logincompany40")
	loginCmp.Email = "Simple Text"
	assertValidLoginCompany(loginCmp, t, "logincompany40")
	loginCmp.Email = "barbosa.olivera1@gmail.com"

	if log, err := ValidLoginCompany(loginCmp); log == nil || err != nil {
		t.Error("Error valid entity login company: " + err.message)
	}
}
