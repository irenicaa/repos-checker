package bitbucket

import (
	"fmt"
	"net/http"
	"net/url"

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

// LoadData ...
func LoadData(
	endpoint string,
	parameters url.Values,
	responseData interface{},
) error {
	url := MakeURL(endpoint, parameters)
	authHeader :=
		systemutils.MakeBasicAuthHeader("BITBUCKET_USERNAME", "BITBUCKET_PASSWORD")
	return systemutils.LoadJSONData(
		http.DefaultClient,
		url,
		authHeader,
		responseData,
	)
}
