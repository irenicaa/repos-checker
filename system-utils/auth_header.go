package systemutils

import (
	"encoding/base64"
	"os"
)

// MakeBasicAuthHeader ...
func MakeBasicAuthHeader(usernameEnv string, passwordEnv string) string {
	username, password := os.Getenv(usernameEnv), os.Getenv(passwordEnv)
	if username == "" || password == "" {
		return ""
	}

	credentials := username + ":" + password
	credentials = base64.StdEncoding.EncodeToString([]byte(credentials))

	return "Basic " + credentials
}
