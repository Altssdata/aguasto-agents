// Package aguasto provides anonymous clients for Aguasto's public research and
// educational calculations. https://www.aguasto.com/developers
// All rights reserved. No GPU provisioning or inference service is provided.
package aguasto

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

//go:embed fixtures.json
var fixtures []byte

// OfflineFixtures returns independent static examples without a network request.
func OfflineFixtures() json.RawMessage { return append(json.RawMessage(nil), fixtures...) }

type APIError struct {
	Status        int
	Code, Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s (HTTP %d): %s", e.Code, e.Status, e.Message)
}

type Client struct {
	base string
	http *http.Client
}

// NewClient accepts an HTTPS origin, or HTTP loopback for local tests. Empty
// selects https://www.aguasto.com. Credentials, paths and redirects are rejected.
func NewClient(base string) (*Client, error) {
	if base == "" {
		base = "https://www.aguasto.com"
	}
	u, err := url.Parse(base)
	if err != nil {
		return nil, fmt.Errorf("invalid origin: %w", err)
	}
	loopback := u.Hostname() == "localhost"
	if ip := net.ParseIP(u.Hostname()); ip != nil {
		loopback = ip.IsLoopback()
	}
	if u.Host == "" || u.User != nil || u.RawQuery != "" || u.ForceQuery || u.Fragment != "" || (u.Path != "" && u.Path != "/") || (u.Scheme != "https" && !(u.Scheme == "http" && loopback)) {
		return nil, fmt.Errorf("use an HTTPS origin without credentials, path, query or fragment; HTTP is limited to loopback")
	}
	u.Path = ""
	return &Client{base: u.String(), http: &http.Client{Timeout: 15 * time.Second, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}}, nil
}

func (c *Client) call(ctx context.Context, method, path string, input any) (map[string]any, error) {
	var reader io.Reader
	if input != nil {
		data, err := json.Marshal(input)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(data)
	}
	request, err := http.NewRequestWithContext(ctx, method, c.base+path, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Accept", "application/json")
	if input != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	response, err := c.http.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(io.LimitReader(response.Body, (1<<20)+1))
	if err != nil {
		return nil, err
	}
	if len(data) > 1<<20 {
		return nil, fmt.Errorf("response exceeds 1 MiB")
	}
	var output map[string]any
	if err := json.Unmarshal(data, &output); err != nil {
		return nil, fmt.Errorf("HTTP %d returned invalid JSON", response.StatusCode)
	}
	if output == nil {
		return nil, fmt.Errorf("HTTP %d returned no JSON object", response.StatusCode)
	}
	failure := response.StatusCode < 200 || response.StatusCode >= 300
	if meta, ok := output["_meta"].(map[string]any); ok && meta["response_type"] == "failure" {
		failure = true
	}
	if failure {
		result := &APIError{Status: response.StatusCode, Code: "HTTP_ERROR", Message: "Request was rejected"}
		if detail, ok := output["error"].(map[string]any); ok {
			if v, ok := detail["code"].(string); ok {
				result.Code = v
			}
			if v, ok := detail["message"].(string); ok {
				result.Message = v
			}
		}
		return nil, result
	}
	return output, nil
}

func (c *Client) Catalog(ctx context.Context) (map[string]any, error) {
	return c.call(ctx, http.MethodGet, "/api/v1/catalog", nil)
}
func (c *Client) Search(ctx context.Context, query string) (map[string]any, error) {
	return c.call(ctx, http.MethodPost, "/ask", map[string]any{"query": map[string]string{"text": query}, "prefer": map[string]string{"mode": "list"}})
}
func (c *Client) Calculate(ctx context.Context, name string, parameters map[string]float64) (map[string]any, error) {
	return c.call(ctx, http.MethodPost, "/api/v1/calculate", map[string]any{"calculation": name, "parameters": parameters})
}
