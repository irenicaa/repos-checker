package gitlab

import (
	"fmt"
	"net/url"
	"os"

	systemutils "github.com/irenicaa/repos-checker/system-utils"
)

// MakeURL ...
func MakeURL(endpoint string, parameters url.Values) string {
	return fmt.Sprintf(
		"https://gitlab.com/api/v4%s?%s",
		endpoint,
		parameters.Encode(),
	)
}

// MakeAuthHeader ...
func MakeAuthHeader() string {
	token := os.Getenv("GITLAB_TOKEN")
	if token == "" {
		return ""
	}

	return "Bearer " + token
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
