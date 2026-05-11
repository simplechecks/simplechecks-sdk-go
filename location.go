// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/simplechecks/simplechecks-sdk-go/internal/apijson"
	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
)

// Catalog of (provider, location) deployments Simple Checks runs checks from, with
// geographic metadata + live status. Used to drive the region picker and the
// dashboard's locations map.
//
// LocationService contains methods and other services that help with interacting
// with the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLocationService] method instead.
type LocationService struct {
	Options []option.RequestOption
}

// NewLocationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLocationService(opts ...option.RequestOption) (r *LocationService) {
	r = &LocationService{}
	r.Options = opts
	return
}

// Returns every (provider, location) Simple Checks deploys garrisons at, enriched
// with geographic metadata (city, country, continent, IATA-style metro code,
// lat/lon) and live garrison status. The catalog of locations is code-defined and
// identical across customers; only `status` is dynamic.
//
// Locations whose backing garrison hasn't been provisioned yet are returned with
// `status: "unprovisioned"` so the dashboard can grey them out while keeping the
// catalog visible.
//
// Requires the `account:read` scope — listing locations is incidental to account
// configuration, not a per-check write.
func (r *LocationService) List(ctx context.Context, opts ...option.RequestOption) (res *LocationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// One deployed (provider, data-center) tuple where Simple Checks runs garrisons,
// with geographic metadata + live status.
type Location struct {
	// Composite identifier; `<provider>:<location>` (e.g. `aws:us-east-1`).
	ID        string            `json:"id" api:"required"`
	City      string            `json:"city" api:"required"`
	Continent LocationContinent `json:"continent" api:"required"`
	// ISO 3166-1 alpha-2 country code.
	Country string `json:"country" api:"required"`
	// Provider-native data-center id (varies in format per provider).
	Location string `json:"location" api:"required"`
	// Cloud provider.
	Provider string `json:"provider" api:"required"`
	// Live garrison status. `unprovisioned` means the location is code-defined but no
	// garrison row exists yet (deploy pending); dashboard typically greys these out.
	Status LocationStatus `json:"status" api:"required"`
	// Metro-center latitude (degrees, WGS84).
	Lat float64 `json:"lat"`
	// Metro-center longitude (degrees, WGS84).
	Lon float64 `json:"lon"`
	// IATA-style 3-letter code for the nearest major metro. Empty for the mock
	// provider; "loose anchor" (not a precise claim) for non-airport-adjacent sites
	// like Hetzner Falkenstein.
	Metro string       `json:"metro"`
	JSON  locationJSON `json:"-"`
}

// locationJSON contains the JSON metadata for the struct [Location]
type locationJSON struct {
	ID          apijson.Field
	City        apijson.Field
	Continent   apijson.Field
	Country     apijson.Field
	Location    apijson.Field
	Provider    apijson.Field
	Status      apijson.Field
	Lat         apijson.Field
	Lon         apijson.Field
	Metro       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationJSON) RawJSON() string {
	return r.raw
}

type LocationContinent string

const (
	LocationContinentNa LocationContinent = "NA"
	LocationContinentSa LocationContinent = "SA"
	LocationContinentEu LocationContinent = "EU"
	LocationContinentAs LocationContinent = "AS"
	LocationContinentAf LocationContinent = "AF"
	LocationContinentOc LocationContinent = "OC"
	LocationContinentAn LocationContinent = "AN"
)

func (r LocationContinent) IsKnown() bool {
	switch r {
	case LocationContinentNa, LocationContinentSa, LocationContinentEu, LocationContinentAs, LocationContinentAf, LocationContinentOc, LocationContinentAn:
		return true
	}
	return false
}

// Live garrison status. `unprovisioned` means the location is code-defined but no
// garrison row exists yet (deploy pending); dashboard typically greys these out.
type LocationStatus string

const (
	LocationStatusReady         LocationStatus = "ready"
	LocationStatusDraining      LocationStatus = "draining"
	LocationStatusMaintenance   LocationStatus = "maintenance"
	LocationStatusUnprovisioned LocationStatus = "unprovisioned"
)

func (r LocationStatus) IsKnown() bool {
	switch r {
	case LocationStatusReady, LocationStatusDraining, LocationStatusMaintenance, LocationStatusUnprovisioned:
		return true
	}
	return false
}

type LocationListResponse struct {
	Locations []Location               `json:"locations" api:"required"`
	JSON      locationListResponseJSON `json:"-"`
}

// locationListResponseJSON contains the JSON metadata for the struct
// [LocationListResponse]
type locationListResponseJSON struct {
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationListResponseJSON) RawJSON() string {
	return r.raw
}
