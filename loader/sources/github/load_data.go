package github

import (
	"fmt"
	"net/http"
	"net/url"

	httputils "github.com/irenicaa/go-http-utils"
	"github.com/irenicaa/go-http-utils/auth"
)

// MakeURL ...
func MakeURL(endpoint string, parameters url.Values) string {
	return fmt.Sprintf(
		"https://api.github.com%s?%s",
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
	authHeader := auth.MakeBasicAuthHeader("GITHUB_USERNAME", "GITHUB_TOKEN")
	return httputils.LoadJSONData(
		http.DefaultClient,
		url,
		authHeader,
		responseData,
	)
}
