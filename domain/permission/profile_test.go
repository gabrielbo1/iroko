package permission

import "testing"

func assertValidProfile(profile *Profile, t *testing.T, errCode string) {
	if _, err := ValidProfile(profile); err == nil || err.GetCode() != errCode {
		t.Error("Error code entity company, code " + errCode)
	}
}

func TestValidProfile(t *testing.T) {
	profile := &Profile{}
	assertValidProfile(profile, t, "profile10")
	profile.Name = "Admin"
	assertValidProfile(profile, t, "profile20")
	profile.Permissions = []Permission{}
	assertValidProfile(profile, t, "profile20")
	profile.Permissions = append(profile.Permissions, Permission{Create: true})
	if perm, err := ValidProfile(profile); err != nil || perm == nil {
		t.Error("Error valid profile.")
	}
}
