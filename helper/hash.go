package helper

import (
	"encoding/hex"
	"hash"
)

const (
	Sha256     = "SHA256"
	HmacSha512 = "HMAC-SHA512"
	MD5        = "MD5"
)

func ComputeSecureHash(data string, hashAlgo string, hashSecret string) string {
	var h hash.Hash
	h.Reset()

	if hashAlgo == HmacSha512 {
		h.Write([]byte(data))
	}

	if hashAlgo == Sha256 || hashAlgo == MD5 {
		h.Write([]byte(hashSecret + data))
	}

	return hex.EncodeToString(h.Sum(nil))
}
