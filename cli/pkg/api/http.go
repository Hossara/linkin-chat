package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Hossara/linkin-chat/cli/types"
	"io"
	"net/http"
	"strings"
)

type RequestBody map[string]interface{}
type RequestQuery map[string]string

type Api struct {
	BaseUrl string
	Port    uint
}

func NewApiHandler(baseUrl string, port uint) *Api {
	if !strings.HasPrefix(baseUrl, "http://") && !strings.HasPrefix(baseUrl, "https://") {
		baseUrl = "http://" + baseUrl
	}

	return &Api{BaseUrl: baseUrl, Port: port}
}

func (a *Api) buildURL(endpoint string, query *RequestQuery) string {
	url := fmt.Sprintf("%s:%d%s", a.BaseUrl, a.Port, endpoint)
	if query != nil && len(*query) > 0 {
		url += "?"

		for key, value := range *query {
			url += fmt.Sprintf("%s=%s&", key, value)
		}

		url = url[:len(url)-1] // Remove the trailing "&"
	}
	return url
}

func Post[T any](a *Api, endpoint string, body *RequestBody, query *RequestQuery) (*T, error) {
	// Build URL with query parameters
	url := a.buildURL(endpoint, query)

	// Convert body map to JSON
	var jsonBody []byte
	var err error

	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshalling JSON: %w", err)
		}
	} else {
		jsonBody = []byte("{}")
	}

	reqBody := bytes.NewBuffer(jsonBody)

	// Create a new POST request
	res, err := http.Post(url, "application/json", reqBody)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Read the response body
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Handle the response status code
	if res.StatusCode != http.StatusOK {
		var errRes *types.ErrorResponse
		if err := json.Unmarshal(respBody, &errRes); err != nil {
			return nil, fmt.Errorf("received non-OK status code: %d\nResponse Body: %s", res.StatusCode, respBody)
		}

		return nil, fmt.Errorf(errRes.Error)
	}

	var result *T
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return result, nil
}

func Get[T any](a *Api, endpoint string, query *RequestQuery) (*T, error) {
	// Build URL with query parameters
	url := a.buildURL(endpoint, query)

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Handle the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK status code: %d\nResponse Body: %s", resp.StatusCode, respBody)
	}

	var result *T
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return result, nil
}
