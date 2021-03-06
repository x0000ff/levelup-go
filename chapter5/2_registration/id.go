package main

import (
	"crypto/rand"
	"fmt"
)

// Source String used when generationg a random identifier
const idSource = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Save the length in a constant so we don't look it up each time
const idSourceLength = byte(len(idSource))

// GenerateID creates a prefixed random identifier
func GenerateID(prefix string, length int) string {
	// Create an array with the correct capacity
	id := make([]byte, length)

	// Fill our array with random numbers
	rand.Read(id)

	// Replace each random number with an alphanumeric value
	for i, b := range id {
		id[i] = idSource[b%idSourceLength]
	}

	return fmt.Sprintf("%s_%s", prefix, string(id))
}
