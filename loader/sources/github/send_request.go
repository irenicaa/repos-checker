package github

import (
	"fmt"
	"net/url"

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

// SendRequest ...
func SendRequest(
	endpoint string,
	parameters url.Values,
	responseData interface{},
) error {
	url := MakeURL(endpoint, parameters)
	authHeader :=
		systemutils.MakeBasicAuthHeader("GITHUB_USERNAME", "GITHUB_TOKEN")
	return systemutils.SendRequest(url, authHeader, responseData)
}
