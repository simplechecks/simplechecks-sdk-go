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

func TestCheckDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := simplechecksgo.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.Checks.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *simplechecksgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
