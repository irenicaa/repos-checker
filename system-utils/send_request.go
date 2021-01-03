package systemutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SendRequest ...
func SendRequest(
	url string,
	authHeader string,
	responseData interface{},
) error {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("unable to create the request: %v", err)
	}

	if authHeader != "" {
		request.Header.Add("Authorization", authHeader)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("unable to send the request: %v", err)
	}
	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("unable to read the request body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"request was failed: %d %s",
			response.StatusCode,
			responseBytes,
		)
	}

	if err = json.Unmarshal(responseBytes, responseData); err != nil {
		return fmt.Errorf("unable to unmarshal the request body: %v", err)
	}

	return nil
}
