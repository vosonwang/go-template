package util

import (
	"log"
	"os"
	"strconv"
)

// GetEnv looks up the given key from the environment, returning its value if
// it exists, and otherwise returning the given fallback value.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetInt looks up the given key from the environment and expects an integer,
// returning the integer value if it exists, and otherwise returning the given
// fallback value.
// If the environment variable has a value but it can't be parsed as an integer,
// GetInt terminates the program.
func GetInt(key string, fallback int) int {
	if s, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("bad value %q for %s: %v", s, key, err)
		}
		return v
	}
	return fallback
}

func GetBool(key string) bool {
	if s, ok := os.LookupEnv(key); ok {
		return s == "true"
	}
	return false
}
