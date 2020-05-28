package common

import (
	"encoding/hex"
)

func BytesToHexString(data []byte) string {
	return hex.EncodeToString(data)
}

func Uint32ToBytes(n uint32) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
	}
}
