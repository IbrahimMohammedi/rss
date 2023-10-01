package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts API key from the request's header
// Example : ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, "")

	if len(val) != 2 {
		return "", errors.New("error auth header")
	}

	if vals[0] != "Apikey" {
		return "", errors.New("error first part of auth header")
	}
	return vals[1], nil
}
