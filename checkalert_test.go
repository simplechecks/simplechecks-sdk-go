// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/simplechecks/simplechecks-sdk-go"
	"github.com/simplechecks/simplechecks-sdk-go/internal/testutil"
	"github.com/simplechecks/simplechecks-sdk-go/option"
)

func TestCheckAlertGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := simplechecksgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Checks.Alerts.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *simplechecksgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckAlertDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := simplechecksgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Checks.Alerts.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *simplechecksgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckAlertReplaceWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := simplechecksgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Checks.Alerts.Replace(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		simplechecksgo.CheckAlertReplaceParams{
			AlertConfig: simplechecksgo.AlertConfigParam{
				ConsecutiveFailuresThreshold: simplechecksgo.F(int64(1)),
				ConsensusM:                   simplechecksgo.F(int64(1)),
				ConsensusN:                   simplechecksgo.F(int64(1)),
				Enabled:                      simplechecksgo.F(true),
				AccountID:                    simplechecksgo.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				CheckID:                      simplechecksgo.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
		},
	)
	if err != nil {
		var apierr *simplechecksgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
