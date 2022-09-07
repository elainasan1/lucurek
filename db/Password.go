package db

import (
	"crypto/sha256"
	"encoding/hex"
)

//HashPassword uses SHA256 to hash the password
func HashPassword(password string) string {

	return hex.EncodeToString(sha256.New().Sum([]byte(password)))
}
