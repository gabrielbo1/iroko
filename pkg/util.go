package pkg

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomBytes Author: Matt Silverlock
// Url: https://blog.questionable.services/article/generating-secure-random-numbers-crypto-rand/
// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString Author: Matt Silverlock
// Url: https://blog.questionable.services/article/generating-secure-random-numbers-crypto-rand/
// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

// Page - Paginated queries pattern.
type Page struct {

	//First - Is first page.
	First bool `json:"first"`

	//Last - Is last page.
	Last bool `json:"last"`

	//PageNumber - The page number.
	PageNumber int `json:"pageNumber"`

	//PageSize - The page size.
	PageSize int `json:"pageSize"`

	//Content - The content page.
	Content interface{} `json:"content"`
}
