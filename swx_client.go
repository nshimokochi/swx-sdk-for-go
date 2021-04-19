package openapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"golang.org/x/oauth2"
)

// SWXToken stores the Smartworks Token
type SWXToken struct {
	oauth2.Token
	Scope string `json:"scope,omitempty"`
}

// Authentication Error response from backend
type AuthNError struct {
	ErrorMessage     string `json:"error,omitempty"`
	ErrorDebug       string `json:"error_debug,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	ErrorVerbose     string `json:"error_verbose,omitempty"`
	StatusCode       int    `json:"status_code,omitempty"`
}

func (e *AuthNError) Error() string { return e.ErrorDescription }

func (c *APIClient) Authenticate(clientID string, clientSecret string, scope string) (string, error) {
	// Check required parameters
	required := []string{clientID, clientSecret, scope}
	for _, v := range required {
		if v == "" {
			return "", fmt.Errorf("Authenticate requires client_id, client_secret, and scope")
		}
	}

	// Create request URL
	serverURL, err := c.cfg.ServerURL(0, nil)
	if serverURL == "" || err != nil {
		return "", fmt.Errorf("Web of Things Server URL not found")
	}
	requestURL := serverURL + "/oauth2/token"

	// Make the request
	resp, err := c.cfg.HTTPClient.PostForm(requestURL, url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"scope":         {scope},
	})
	if err != nil {
		return "", fmt.Errorf("Error when calling `HTTPClient.PostForm`: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error when reading `Response.Body`: %v", err)
	}

	// Decode Error
	authNErr := AuthNError{}
	json.Unmarshal(body, &authNErr)
	if authNErr.StatusCode != 0 {
		return "", &authNErr
	}

	// Decode Token
	latestToken := SWXToken{}
	json.Unmarshal(body, &latestToken)
	if latestToken.AccessToken == "" {
		return "", fmt.Errorf("AccessToken Not Found in `Response.Body`: %v", body)
	}

	// Set Authorization header
	c.cfg.AddDefaultHeader("Authorization", "Bearer "+latestToken.AccessToken)

	return latestToken.AccessToken, nil
}
