package pkg

import (
	"time"
)

type JwtValueKey struct {
	Key  string
	Date time.Time
}

func NewJwtValueKey() JwtValueKey {
	singKey, _ := GenerateRandomString(128)
	return JwtValueKey{singKey, time.Now().Add(time.Duration(24 * time.Hour))}
}
