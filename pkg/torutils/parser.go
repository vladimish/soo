package torutils

import (
	"bytes"
	"encoding/base32"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	"strings"
)

// ED25519ToAddress getting 64 chars hex string,
// representing ed25519 public key
// and returns .onion url address.
func ED25519ToAddress(publicKey string) string {
	pub, err := hex.DecodeString(publicKey)
	if err != nil {
		panic(err)
	}
	version := "\x03"
	checksumData := append([]byte(".onion checksum"), pub...)
	checksumData = append(checksumData, []byte(version)...)

	var checksum []byte
	h := sha3.New256()
	h.Write(checksumData)
	checksum = h.Sum(checksum)[:2]

	var finalBytes []byte
	finalWriter := bytes.NewBuffer(finalBytes)

	e := base32.NewEncoder(base32.StdEncoding, finalWriter)
	e.Write(append(append(pub, checksum...), version...))

	return strings.ToLower(finalWriter.String() + ".onion")
}
