package common

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
)

var EmptyBytes = make([]byte, 0)

func BytesReverse(u []byte) []byte {
	for i, j := 0, len(u)-1; i < j; i, j = i+1, j-1 {
		u[i], u[j] = u[j], u[i]
	}
	return u
}

func BytesToHexString(data []byte) string {
	return hex.EncodeToString(data)
}

func HexStringToBytes(value string) ([]byte, error) {
	return hex.DecodeString(value)
}

func IntToBytes(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, tmp)
	return bytesBuffer.Bytes()
}

func BytesToInt16(b []byte) int16 {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int16
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int16(tmp)
}

func ClearBytes(arr []byte) {
	for i := 0; i < len(arr); i++ {
		arr[i] = 0
	}
}

func Sha256D(data []byte) [32]byte {
	once := sha256.Sum256(data)
	return sha256.Sum256(once[:])
}

func Hash(data []byte) Uint256 {
	return Uint256(Sha256D(data))
}

func Uint32ToBytes(n uint32) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
	}
}
