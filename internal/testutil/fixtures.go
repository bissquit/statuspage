package testutil

import (
	"fmt"
	"math/rand"
)

// RandomSlug generates a random slug for testing.
func RandomSlug(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, rand.Intn(100000))
}

// RandomEmail generates a random email for testing.
func RandomEmail() string {
	return fmt.Sprintf("test-%d@example.com", rand.Intn(100000))
}
