package github

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os"

	systemutils "github.com/irenicaa/repos-checker/system-utils"
)

// MakeURL ...
func MakeURL(endpoint string, parameters url.Values) string {
	return fmt.Sprintf(
		"https://api.github.com%s?%s",
		endpoint,
		parameters.Encode(),
	)
}

// MakeAuthHeader ...
func MakeAuthHeader() string {
	username, token := os.Getenv("GITHUB_USERNAME"), os.Getenv("GITHUB_TOKEN")
	if username == "" || token == "" {
		return ""
	}

	credentials := username + ":" + token
	credentials = base64.StdEncoding.EncodeToString([]byte(credentials))

	return "Basic " + credentials
}

// SendRequest ...
func SendRequest(
	endpoint string,
	parameters url.Values,
	responseData interface{},
) error {
	url := MakeURL(endpoint, parameters)
	authHeader := MakeAuthHeader()
	return systemutils.SendRequest(url, authHeader, responseData)
}
