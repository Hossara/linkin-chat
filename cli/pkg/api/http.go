package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Hossara/linkin-chat/cli/types"
	"github.com/Hossara/linkin-chat/pkg/utils"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"strings"
)

type RequestBody map[string]interface{}
type RequestQuery map[string]string
type Headers map[string]string

type Api struct {
	BaseUrl string
	Port    uint
}

func GetAuthHeaders() *Headers {
	return &Headers{
		"Authorization": fmt.Sprintf("Bearer %s", viper.GetString("login.token")),
	}
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

func Post[T any](a *Api, endpoint string, body *RequestBody, query *RequestQuery, headers *Headers) (*T, error) {
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

	// Create a new POST request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if headers != nil {
		for key, value := range *headers {
			req.Header.Set(key, value)
		}
	}

	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}
	defer res.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Handle the response status code
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized {
			return nil, fmt.Errorf("your login has expired")
		}

		var errRes *types.ErrorResponse
		if err := json.Unmarshal(respBody, &errRes); err != nil {
			return nil, fmt.Errorf("received non-OK status code: %d\nResponse Body: %s", res.StatusCode, respBody)
		}

		return nil, fmt.Errorf(
			errRes.Error + utils.IfThenElse(errRes.Message == "", "",
				fmt.Sprintf(" %s", errRes.Message),
			).(string),
		)
	}

	var result *T
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return result, nil
}

func Get[T any](a *Api, endpoint string, query *RequestQuery, headers *Headers) (*T, error) {
	// Build URL with query parameters
	url := a.buildURL(endpoint, query)

	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if headers != nil {
		for key, value := range *headers {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}
	defer res.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Handle the response status code
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized {
			return nil, fmt.Errorf("your login has expired")
		}

		var errRes *types.ErrorResponse
		if err := json.Unmarshal(respBody, &errRes); err != nil {
			return nil, fmt.Errorf("received non-OK status code: %d\nResponse Body: %s", res.StatusCode, respBody)
		}

		return nil, fmt.Errorf(
			errRes.Error + utils.IfThenElse(errRes.Message == "", "",
				fmt.Sprintf(" %s", errRes.Message),
			).(string),
		)
	}

	var result *T
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return result, nil
}
