package openapi

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func checkEnvs(keys ...string) bool {
	var notFound []string
	for _, key := range keys {
		if _, ok := os.LookupEnv(key); !ok {
			notFound = append(notFound, key)
		}
	}
	if len(notFound) != 0 {
		fmt.Printf("Missing env variables: %v\n", strings.Join(notFound, ", "))
		return false
	} else {
		return true
	}
}

var cfg *Configuration
var apiClient *APIClient

func TestMain(m *testing.M) {
	cfg = NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: os.Getenv("SWX_HOST")}}
	apiClient = NewAPIClient(cfg)

	if checkEnvs("SWX_HOST", "CLIENT_ID", "CLIENT_SECRET", "SCOPE", "SPACE", "SOURCE") {
		os.Exit(m.Run())
	} else {
		os.Exit(1)
	}
}

func TestNegativeAuthenticate(t *testing.T) {
	token, err := apiClient.Authenticate(
		os.Getenv("CLIENT_ID"),
		"ZjyKbKC47Oj3XEMx0BWvtgL0kStPMy",
		os.Getenv("SCOPE"),
	)
	if err == nil {
		t.Errorf("`APIClient.Authenticate` = %v; want `Client authentication failed` error\n", token)
	}
}

func TestPositiveAuthenticate(t *testing.T) {
	apiClient = NewAPIClient(cfg)
	_, err := apiClient.Authenticate(
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		os.Getenv("SCOPE"),
	)
	if err != nil {
		t.Errorf("`APIClient.Authenticate` = %v; want `access_token`\n", err)
	}
}
