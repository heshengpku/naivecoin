package utils

import (
	"crypto/sha256"
	"fmt"
)

func CalculateHash(any string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(any)))
}
