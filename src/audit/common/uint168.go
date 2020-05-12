package common

import (
	"bytes"
	"errors"
	"io"
)

const (
	UINT168SIZE = 21
	// Address types
	STANDARD   = 0xAC
	DID        = 0xAD
	MULTISIG   = 0xAE
	CROSSCHAIN = 0xAF
)

type Uint168 [UINT168SIZE]uint8

func (u Uint168) String() string {
	return BytesToHexString(u.Bytes())
}

func (u Uint168) Compare(o Uint168) int {
	for i := UINT168SIZE - 1; i >= 0; i-- {
		if u[i] > o[i] {
			return 1
		}
		if u[i] < o[i] {
			return -1
		}
	}
	return 0
}

func (u Uint168) IsEqual(o Uint168) bool {
	return bytes.Equal(u[:], o[:])
}

func (u Uint168) Bytes() []byte {
	return u[:]
}

func (u *Uint168) Serialize(w io.Writer) error {
	_, err := w.Write(u[:])
	return err
}

func (u *Uint168) Deserialize(r io.Reader) error {
	_, err := io.ReadFull(r, u[:])
	return err
}

func Uint168FromBytes(bytes []byte) (*Uint168, error) {
	if len(bytes) != UINT168SIZE {
		return nil, errors.New("[Uint168FromBytes] error, len != 21")
	}

	var hash Uint168
	copy(hash[:], bytes)

	return &hash, nil
}
