package auth

import (
	"fmt"
	"net/http"
	"strings"
)

// Extract API Key from headers of HTTP request
func GetAPIKey(headers http.Header) (string, error) {
	apiKey := headers.Get("Authorization")
	if apiKey == "" {
		return "", fmt.Errorf("API Key not found")
	}

	keys := strings.Split(apiKey, " ")
	if len(keys) != 2 {
		return "", fmt.Errorf("invalid api key format")
	}
	if keys[0] != "ApiKey" {
		return "", fmt.Errorf("invalid first part of api key")
	}
	return keys[1], nil
}
