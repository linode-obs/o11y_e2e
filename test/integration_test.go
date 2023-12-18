package integrationtest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"o11y_e2e/internal/server"
)

const expectedStatusCode = 200

func setupTestServer() *httptest.Server {
	handler := server.SetupServer()
	return httptest.NewServer(handler)
}

func validateResponse(t *testing.T, resp *http.Response, expectedBodyContents ...string) {
	if resp.StatusCode != expectedStatusCode {
		t.Fatalf("Expected status %d, got: %d", expectedStatusCode, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read body: %v", err)
	}

	for _, content := range expectedBodyContents {
		if !strings.Contains(string(body), content) {
			t.Fatalf("Expected to find %s in response, but not found. Full content: %v", content, string(body))
		}
	}
}

func TestO11yE2ERootEndpoint(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	validateResponse(t, resp, "<head><title>o11y_e2e</title></head>")
}

func TestO11YE2EMetricsEndpoint(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/metrics")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	validateResponse(t, resp, "promhttp_metric_handler_requests_in_flight 1")
}
