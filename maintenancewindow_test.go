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

func TestMaintenanceWindowNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MaintenanceWindows.New(context.TODO(), simplechecksgo.MaintenanceWindowNewParams{
		DurationMs:       simplechecksgo.F(int64(0)),
		Name:             simplechecksgo.F("name"),
		ScheduleKind:     simplechecksgo.F(simplechecksgo.MaintenanceWindowNewParamsScheduleKindOneTime),
		StartUnixMs:      simplechecksgo.F(int64(0)),
		CheckIDs:         simplechecksgo.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		CheckTags:        simplechecksgo.F([]string{"string"}),
		RepeatEndsUnixMs: simplechecksgo.F(int64(0)),
		RepeatInterval:   simplechecksgo.F(int64(0)),
		RepeatUnit:       simplechecksgo.F(simplechecksgo.MaintenanceWindowNewParamsRepeatUnitDay),
		Timezone:         simplechecksgo.F("timezone"),
	})
	if err != nil {
		var apierr *simplechecksgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMaintenanceWindowGet(t *testing.T) {
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
	_, err := client.MaintenanceWindows.Get(context.TODO(), "id")
	if err != nil {
		var apierr *simplechecksgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMaintenanceWindowUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MaintenanceWindows.Update(
		context.TODO(),
		"id",
		simplechecksgo.MaintenanceWindowUpdateParams{
			CheckIDs:         simplechecksgo.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			CheckTags:        simplechecksgo.F([]string{"string"}),
			DurationMs:       simplechecksgo.F(int64(0)),
			Name:             simplechecksgo.F("name"),
			RepeatEndsUnixMs: simplechecksgo.F(int64(0)),
			RepeatInterval:   simplechecksgo.F(int64(0)),
			RepeatUnit:       simplechecksgo.F(simplechecksgo.MaintenanceWindowUpdateParamsRepeatUnitDay),
			ScheduleKind:     simplechecksgo.F(simplechecksgo.MaintenanceWindowUpdateParamsScheduleKindOneTime),
			StartUnixMs:      simplechecksgo.F(int64(0)),
			Timezone:         simplechecksgo.F("timezone"),
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

func TestMaintenanceWindowListWithOptionalParams(t *testing.T) {
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
	_, err := client.MaintenanceWindows.List(context.TODO(), simplechecksgo.MaintenanceWindowListParams{
		Cursor: simplechecksgo.F("cursor"),
		Limit:  simplechecksgo.F(int64(1)),
	})
	if err != nil {
		var apierr *simplechecksgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMaintenanceWindowDelete(t *testing.T) {
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
	err := client.MaintenanceWindows.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *simplechecksgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
