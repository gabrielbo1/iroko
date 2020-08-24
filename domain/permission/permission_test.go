package permission

import "testing"

func assertValidPermission(permission *Permission, t *testing.T, errCode string) {
	if _, err := ValidPermission(permission); err == nil || err.GetCode() != errCode {
		t.Error("Error code entity company, code " + errCode)
	}
}

func TestValidPermission(t *testing.T) {
	permission := &Permission{}
	assertValidPermission(permission, t, "permission10")
	permission.Create = true
	if perm, err := ValidPermission(permission); err != nil || perm == nil {
		t.Error("Error valid permission.")
	}
}
