// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
)

// Per-check alert configuration + test-fire endpoint (PR-Alerts/1).
//
// CheckAlertService contains methods and other services that help with interacting
// with the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckAlertService] method instead.
type CheckAlertService struct {
	Options []option.RequestOption
}

// NewCheckAlertService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCheckAlertService(opts ...option.RequestOption) (r *CheckAlertService) {
	r = &CheckAlertService{}
	r.Options = opts
	return
}

// Subsequent runs will not be evaluated for alerts. State rows in `alert_state`
// and `alert_location_state` cascade with the underlying check; deleting just the
// config leaves them behind harmlessly. Requires the `checks:write` scope.
func (r *CheckAlertService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/checks/%s/alerts", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}
