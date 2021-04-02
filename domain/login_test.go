package domain

import "testing"

func assertValidLogin(login *Login, t *testing.T, errCode string) {
	if _, err := login.ValidLogin(); err == nil || err.GetCode() != errCode {
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
	assertValidLogin(login, t, "login30")
	login.Email = "Simple Text"
	assertValidLogin(login, t, "login30")
	login.Email = "barbosa.olivera1@gmail.com"

	if log, err := login.ValidLogin(); log == nil || err != nil {
		t.Error("Error valid entity login")
	}
}
