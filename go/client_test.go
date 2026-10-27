package aguasto

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestRequestsAndFailures(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" || r.Header.Get("Cookie") != "" {
			t.Error("credentials sent")
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Error("missing JSON negotiation")
		}
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/v1/catalog":
			if r.Method != "GET" {
				t.Error("catalog method")
			}
			_, _ = w.Write([]byte(`{"pages":[{"name":"GPU reference"}]}`))
		case "/api/v1/calculate":
			if r.Method != "POST" || r.Header.Get("Content-Type") != "application/json" {
				t.Error("calculation request")
			}
			_, _ = w.Write([]byte(`{"result":{"gb":35}}`))
		default:
			_, _ = w.Write([]byte(`{"_meta":{"response_type":"failure"},"error":{"code":"NO_RESULTS","message":"No reference"}}`))
		}
	}))
	defer server.Close()
	client, err := NewClient(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	catalog, err := client.Catalog(context.Background())
	if err != nil || len(catalog["pages"].([]any)) != 1 {
		t.Fatalf("catalog: %v", err)
	}
	result, err := client.Calculate(context.Background(), "weight_memory", map[string]float64{"bits": 4, "parameters_billions": 70})
	if err != nil || result["result"].(map[string]any)["gb"] != float64(35) {
		t.Fatalf("calculate: %v", err)
	}
	_, err = client.Search(context.Background(), "no-match")
	var apiError *APIError
	if !errors.As(err, &apiError) || apiError.Code != "NO_RESULTS" {
		t.Fatalf("application failure: %v", err)
	}
}
func TestOriginAndResponseBoundaries(t *testing.T) {
	for _, base := range []string{"http://example.com", "https://user:secret@example.com", "https://example.com/private", "https://example.com?token=x", "file:///tmp/data"} {
		if _, err := NewClient(base); err == nil {
			t.Errorf("accepted %s", base)
		}
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "/api/v1/catalog", http.StatusFound) }))
	defer server.Close()
	client, _ := NewClient(server.URL)
	if _, err := client.Catalog(context.Background()); err == nil {
		t.Fatal("redirect followed")
	}
	large := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { _, _ = w.Write([]byte(strings.Repeat("x", (1<<20)+2))) }))
	defer large.Close()
	client, _ = NewClient(large.URL)
	if _, err := client.Catalog(context.Background()); err == nil || !strings.Contains(err.Error(), "1 MiB") {
		t.Fatalf("unbounded response: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := client.Catalog(ctx); err == nil {
		t.Fatal("context cancellation ignored")
	}
}
func TestLivePublicAPI(t *testing.T) {
	if os.Getenv("AGUASTO_LIVE_TEST") != "1" {
		t.Skip("set AGUASTO_LIVE_TEST=1 for anonymous production integration tests")
	}
	client, err := NewClient("")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := client.Catalog(context.Background()); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Search(context.Background(), "GPU rack cooling"); err != nil {
		t.Fatal(err)
	}
	for _, item := range []struct {
		name, field string
		input       map[string]float64
		expected    float64
	}{
		{"weight_memory", "gb", map[string]float64{"parameters_billions": 70, "bits": 4}, 35},
		{"facility_energy", "electricityUSD", map[string]float64{"it_kw": 100, "pue": 1.2, "hours": 24, "usd_per_kwh": 0.1}, 288},
		{"token_cost", "usdPerMillionCompletedTokens", map[string]float64{"cost_usd": 10, "completed_tokens": 2000000}, 5},
	} {
		result, err := client.Calculate(context.Background(), item.name, item.input)
		if err != nil {
			t.Fatal(err)
		}
		if result["result"].(map[string]any)[item.field] != item.expected {
			t.Fatalf("unexpected %s result", item.name)
		}
	}
}
