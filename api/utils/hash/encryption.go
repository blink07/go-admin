package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func Encryption(str []byte) string {
	h := sha256.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

