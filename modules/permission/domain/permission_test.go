package domain

import (
	"strings"
	"testing"

	"github.com/gabrielbo1/iroko/pkg"
	"github.com/google/uuid"
)

func newRotineTestBuild(rotine *Rotine, expectedErrCode string, t *testing.T) {
	rotineValid, err := NewRotine(*rotine)
	if rotineValid != nil || err.GetCode() != expectedErrCode {
		t.Errorf("Validate user entity, error %s", expectedErrCode)
		t.Errorf("Expected error code %s. Error code returned %s", expectedErrCode, err.GetCode())
		t.Fail()
	}
}

func NewRotineTest(t *testing.T) {
	rotine := &Rotine{}
	newRotineTestBuild(rotine, "PERMISSION_ROTINE_10", t)

	rotine.SystemId = uuid.NewString()
	newRotineTestBuild(rotine, "PERMISSION_ROTINE_20", t)

	rotine.SystemId = uuid.NewString()
	rotine.Rotine = "rotine_name"
	rotine.Path = strings.Repeat("a", pkg.MaxNameSize+1)
	newRotineTestBuild(rotine, "PERMISSION_ROTINE_30", t)

	rotine.Path = ""
	rotineValid, err := NewRotine(*rotine)
	if rotineValid == nil || err != nil {
		t.Error("NewUserTest no create valid entity.")
		t.Fail()
	}

	rotine.Path = strings.Repeat("a", pkg.MaxNameSize)
	if rotineValid == nil || err != nil {
		t.Error("NewUserTest no create valid entity.")
		t.Fail()
	}
}
