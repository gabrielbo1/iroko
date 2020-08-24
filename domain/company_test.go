package domain

import "testing"

func assertValidCompany(company *Company, t *testing.T, errCode string) {
	if _, err := ValidCompany(company); err == nil || err.GetCode() != errCode {
		t.Error("Error code entity company, code " + errCode)
	}
}

func TestValidCompany(t *testing.T) {
	company := &Company{}

	assertValidCompany(company, t, "company10")
	company.Name = ""
	assertValidCompany(company, t, "company10")
	company.Name = "Company Name"
	if comp, err := ValidCompany(company); err != nil || comp == nil {
		t.Error("Error valid company.")
	}
}
