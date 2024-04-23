package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner
func RandomOwner() string {
	return RandomString(6)
}

func RandomRole() string {

	role := []string{
		"root",
		"admin",
	}

	return role[rand.Intn(1)]
}

func RandomeEnable() int {
	return rand.Intn(1)
}

// RandomEmail
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
