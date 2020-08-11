package domain

import "testing"

func assertValidLogin(login *Login, t *testing.T, errCode string) {
	if _, err := ValidLogin(login); err == nil || err.GetCode() != errCode {
		t.Error("Error code entity login, code " + errCode)
	}
}

func TestValidLogin(t *testing.T) {
	login := &Login{}

	assertValidLogin(login, t, "login10")
	login.Login = ""
	assertValidLogin(login, t, "login10")
	login.Login = "user"
	assertValidLogin(login, t, "login20")
	login.Password = ""
	assertValidLogin(login, t, "login20")
	login.Password = "pass"

	if log, err := ValidLogin(login); log == nil || err != nil {
		t.Error("Error valid entity login")
	}
}
