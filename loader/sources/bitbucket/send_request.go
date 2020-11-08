package bitbucket

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
		"https://api.bitbucket.org/2.0%s?%s",
		endpoint,
		parameters.Encode(),
	)
}

// MakeAuthHeader ...
func MakeAuthHeader() string {
	username, password :=
		os.Getenv("BITBUCKET_USERNAME"), os.Getenv("BITBUCKET_PASSWORD")
	if username == "" || password == "" {
		return ""
	}

	credentials := username + ":" + password
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
