package main

import (
	"bytes"
	"crypto/sha256"
	"encoding"
	"fmt"
	"hash"
	"log"
)

// Hash functions are used in cryptography and other areas. These data structures are
// presented with code examples related to cryptography. There are two ways to implement a
// hash function in Go: with crc32 or sha256. Marshaling (changing the string to an encoded
// form) saves the internal state.

func main() {
	const (
		example1 = "this is a example "
		example2 = "second example"
	)

	var firstHash hash.Hash
	firstHash = sha256.New()
	firstHash.Write([]byte(example1))

	var marshaler encoding.BinaryMarshaler
	var ok bool
	marshaler, ok = firstHash.(encoding.BinaryMarshaler)
	if !ok {
		log.Fatal("first Hash does not implement encoding.BinaryMarshaler")
	}
	var data []byte
	var err error
	data, err = marshaler.MarshalBinary()
	if err != nil {
		log.Fatal("unable to marshal first Hash:", err)
	}

	var secondHash hash.Hash
	secondHash = sha256.New()

	var unmarshaler encoding.BinaryUnmarshaler

	unmarshaler, ok = secondHash.(encoding.BinaryUnmarshaler)
	if !ok {
		log.Fatal("second Hash does not implement encoding.BinaryUnmarshaler")
	}
	if err := unmarshaler.UnmarshalBinary(data); err != nil {
		log.Fatal("unable to unmarshal hash:", err)
	}

	firstHash.Write([]byte(example2))
	secondHash.Write([]byte(example2))

	fmt.Printf("%x\n", firstHash.Sum(nil))
	fmt.Println(bytes.Equal(firstHash.Sum(nil), secondHash.Sum(nil)))
}
