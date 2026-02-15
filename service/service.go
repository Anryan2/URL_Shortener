package service

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

func GenerateHash(full string) string {
	sum := sha256.Sum224([]byte(full))
	encoded := base64.URLEncoding.EncodeToString(sum[:])
	return strings.ReplaceAll(strings.TrimRight(encoded, "="), "-", "_")[:10]
}
