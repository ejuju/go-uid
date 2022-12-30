package uid

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// ID represents a unique identifier.
// Generated using the machine's PRNG.
type UID []byte

// Generates a new UID.
func NewUID(length int) (UID, error) {
	buf := make(UID, length)
	_, err := rand.Read(buf)
	return buf, err
}

// String returns the UID as a base-64 string
func (uid UID) String() string { return base64.StdEncoding.EncodeToString(uid) }

// String returns the UID as bytes
func (uid UID) Bytes() []byte { return uid }

// Must is a utility function to check if the UID was generated without an error.
// Used in combination with NewUID function.
// Panics if error is not nil, else returns the UID.
func Must(id UID, err error) UID {
	if err != nil {
		panic(fmt.Errorf("uid: %w", err))
	}
	return id
}
