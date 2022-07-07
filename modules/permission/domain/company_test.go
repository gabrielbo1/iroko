package domain

import (
	"strings"
	"testing"

	"github.com/gabrielbo1/iroko/pkg"
	"github.com/google/uuid"
)

func newSubsidiaryTestBuild(sub *Subsidiary, expectedErrCode string, t *testing.T) {
	subValid, err := NewSubsidiary(*sub)
	if subValid != nil || err.GetCode() != expectedErrCode {
		t.Errorf("Validade subsidiary entity, error %s", expectedErrCode)
		t.Errorf("Expected error code %s. Error code returned %s", expectedErrCode, err.GetCode())
		t.Fail()
	}
}

func TestNewSubsidiary(t *testing.T) {
	subsidiary := &Subsidiary{}
	newSubsidiaryTestBuild(subsidiary, "PERMISSION_SUBSIDIARY_10", t)

	subsidiary.CompanyId = uuid.NewString()
	newSubsidiaryTestBuild(subsidiary, "PERMISSION_SUBSIDIARY_20", t)

	subsidiary.Name = strings.Repeat("a", pkg.MaxNameSize+1)
	newSubsidiaryTestBuild(subsidiary, "PERMISSION_SUBSIDIARY_20", t)

	subsidiary.Name = "Sub1"
	subValid, err := NewSubsidiary(*subsidiary)
	if subValid == nil || err != nil {
		t.Error("NewSubsidiary no create valid entity.")
		t.Fail()
	}
}

func newCompanyTestBuild(company *Company, expectedErrCode string, t *testing.T) {
	companyValid, err := NewCompany(*company)
	if companyValid != nil || err.GetCode() != expectedErrCode {
		t.Errorf("Validade companyValid entity, error %s", expectedErrCode)
		t.Errorf("Expected error code %s. Error code returned %s", expectedErrCode, err.GetCode())
		t.Fail()
	}
}

func TestNewCompany(t *testing.T) {
	company := &Company{}
	newCompanyTestBuild(company, "PERMISSION_COMPANY_10", t)

	company.Name = strings.Repeat("a", pkg.MaxNameSize+1)
	newCompanyTestBuild(company, "PERMISSION_COMPANY_10", t)

	company.Name = "Comp1"
	compValid, err := NewCompany(*company)
	if compValid == nil || err != nil {
		t.Error("NewCompany no create valid entity.")
		t.Fail()
	}
}
