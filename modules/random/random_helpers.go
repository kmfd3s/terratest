package random

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

// UniqueString returns a unique string with prefix
func UniqueString(prefix string) string {
	uniqueID := UniqueId()
	name := fmt.Sprintf("%s-%s", prefix, uniqueID)
	return name
}

// UniqueLowerString returns a unique string with prefix by lowercase
func UniqueLowerString(prefix string) string {
	uniqueID := UniqueId()
	name := fmt.Sprintf("%s-%s", strings.ToLower(prefix), strings.ToLower(uniqueID))
	return name
}

// UniqueAlphaNumericString returns a unique alpha numeric with prefix by lowercase
func UniqueAlphaNumericString(prefix string) string {
	uniqueID := UniqueId()
	name := fmt.Sprintf("%s%s", strings.ToLower(prefix), strings.ToLower(uniqueID))
	return name
}

// RndString returns a random alpha numeric string by lowercase
func RndString() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}
