// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecks_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/simplechecks-go"
	"github.com/stainless-sdks/simplechecks-go/internal/testutil"
	"github.com/stainless-sdks/simplechecks-go/option"
)

func TestCheckNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := simplechecks.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Checks.New(context.TODO(), simplechecks.CheckNewParams{
		Enabled:     true,
		Location:    "location",
		Name:        "name",
		Provider:    "provider",
		Schedule:    "*/5 * * * *",
		TargetURL:   "https://example.com",
		Type:        "http",
		ArtifactURL: simplechecks.String("artifact_url"),
		Config: map[string]any{
			"foo": "bar",
		},
		TimeoutMs: simplechecks.Int(0),
	})
	if err != nil {
		var apierr *simplechecks.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := simplechecks.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Checks.Get(context.TODO(), "id")
	if err != nil {
		var apierr *simplechecks.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := simplechecks.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Checks.Update(
		context.TODO(),
		"id",
		simplechecks.CheckUpdateParams{
			ArtifactURL: simplechecks.String("artifact_url"),
			Config: map[string]any{
				"foo": "bar",
			},
			Enabled:   simplechecks.Bool(true),
			Name:      simplechecks.String("name"),
			Schedule:  simplechecks.String("schedule"),
			TargetURL: simplechecks.String("https://example.com"),
			TimeoutMs: simplechecks.Int(0),
			Type:      simplechecks.String("type"),
		},
	)
	if err != nil {
		var apierr *simplechecks.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := simplechecks.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Checks.List(context.TODO(), simplechecks.CheckListParams{
		Limit:  simplechecks.Int(1),
		Offset: simplechecks.Int(0),
	})
	if err != nil {
		var apierr *simplechecks.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := simplechecks.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Checks.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *simplechecks.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
