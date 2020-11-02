package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
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
func MakeAuthHeader() (string, bool) {
	token := os.Getenv("GITLAB_TOKEN")
	if token == "" {
		return "", false
	}

	return "Bearer " + token, true
}

// SendRequest ...
func SendRequest(
	endpoint string,
	parameters url.Values,
	responseData interface{},
) error {
	url := MakeURL(endpoint, parameters)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("unable to create the request: %v", err)
	}

	authHeader, ok := MakeAuthHeader()
	if ok {
		request.Header.Add("Authorization", authHeader)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("unable to send the request: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"unable to request was failed with the status: %d",
			response.StatusCode,
		)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("unable to read the request body: %v", err)
	}

	if err = json.Unmarshal(responseBytes, responseData); err != nil {
		return fmt.Errorf("unable to unmarshal the request body: %v", err)
	}

	return nil
}
