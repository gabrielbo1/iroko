package domain

import (
	"strings"
	"testing"

	"github.com/gabrielbo1/iroko/pkg"
)

func newUserTestBuild(user *User, expectedErrCode string, t *testing.T) {
	userValid, err := NewUser(*user)
	if userValid != nil || err.GetCode() != expectedErrCode {
		t.Errorf("Validate user entity, error %s", expectedErrCode)
		t.Errorf("Expected error code %s. Error code returned %s", expectedErrCode, err.GetCode())
		t.Fail()
	}
}

func TestNewUser(t *testing.T) {
	user := &User{}
	newUserTestBuild(user, "PERMISSION_USER_10", t)

	user.Nick = strings.Repeat("a", pkg.MaxNameSize+1)
	newUserTestBuild(user, "PERMISSION_USER_10", t)

	user.Nick = "gabrielbo1"
	newUserTestBuild(user, "PERMISSION_USER_20", t)

	user.Nick = "gabrielbo1"
	user.Email = strings.Repeat("a", pkg.MaxNameSize+1)
	newUserTestBuild(user, "PERMISSION_USER_20", t)

	user.Nick = "gabrielbo1"
	user.Email = "invalid_email_user"
	newUserTestBuild(user, "PERMISSION_USER_20", t)

	user.Nick = "gabrielbo1"
	user.Email = "gabrielbo1@gmail.com"
	newUserTestBuild(user, "PERMISSION_USER_30", t)

	user.Nick = "gabrielbo1"
	user.Email = "gabrielbo1@gmail.com"
	user.Password = "password"

	userValid, err := NewUser(*user)
	if userValid == nil || err != nil {
		t.Error("NewUserTest no create valid entity.")
		t.Fail()
	}
}