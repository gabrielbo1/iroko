package permission

import (
	"github.com/gabrielbo1/iroko/domain"
	"testing"
)

func assertValidProfileCompany(profileCompany *ProfileCompany, t *testing.T, errCode string) {
	if _, err := ValidProfileCompany(profileCompany); err == nil || err.GetCode() != errCode {
		t.Error("Error code entity company, code " + errCode)
	}
}

func TestValidProfileCompany(t *testing.T) {
	profile := &ProfileCompany{}
	assertValidProfileCompany(profile, t, "profile10")
	profile.Name = "Admin"
	assertValidProfileCompany(profile, t, "profile20")
	profile.Permissions = []Permission{}
	assertValidProfileCompany(profile, t, "profile20")
	profile.Permissions = append(profile.Permissions, Permission{Create: true})
	assertValidProfileCompany(profile, t, "profilecompany10")
	profile.Company = domain.Company{ID: 1}
	if perm, err := ValidProfileCompany(profile); err != nil || perm == nil {
		t.Error("Error valid profile.")
	}
}
