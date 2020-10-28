package github

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// MakeAuthHeader ...
func MakeAuthHeader() (string, bool) {
	username, token := os.Getenv("GITHUB_USERNAME"), os.Getenv("GITHUB_TOKEN")
	if username == "" || token == "" {
		return "", false
	}

	credentials := username + ":" + token
	credentials = base64.StdEncoding.EncodeToString([]byte(credentials))

	return "Basic " + credentials, true
}

// SendRequest ...
func SendRequest(
	endpoint string,
	parameters url.Values,
	responseData interface{},
) error {
	url := fmt.Sprintf(
		"https://api.github.com%s?%s",
		endpoint,
		parameters.Encode(),
	)

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
