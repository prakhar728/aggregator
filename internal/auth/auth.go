package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API KEY from
// the headers of an HTTP Request
// This is useful for authenticated endpoints
// Example
// Authorization : ApiKey {api_key_here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil
}
