// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/simplechecks/simplechecks-sdk-go/internal/apijson"
	"github.com/simplechecks/simplechecks-sdk-go/internal/param"
	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
)

// Manage who has access to an account and at what role (PR-Members/2). Five roles:
// owner / admin / member / billing / viewer. Owner is the strict superset of all
// other roles' scopes; every account always has at least one owner.
//
// MemberInvitationService contains methods and other services that help with
// interacting with the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMemberInvitationService] method instead.
type MemberInvitationService struct {
	Options []option.RequestOption
}

// NewMemberInvitationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMemberInvitationService(opts ...option.RequestOption) (r *MemberInvitationService) {
	r = &MemberInvitationService{}
	r.Options = opts
	return
}

// Stores a pending invitation and returns it (including the random token, surfaced
// once for the inviter to copy). The webapp emails the accept link on a separate
// step (PR-Members/3); for solo development the inviter can paste the
// `accept_url_path` directly. Requires the `members:write` scope and a user-bound
// key (account-wide keys can't attribute the invitation to a human).
func (r *MemberInvitationService) New(ctx context.Context, body MemberInvitationNewParams, opts ...option.RequestOption) (res *Invitation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/invitations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns pending (not-yet-accepted, not-yet-revoked) invitations. Newest first.
// Tokens are deliberately omitted — they're only returned at creation time so the
// inviter can copy/share the accept link. Requires the `members:read` scope.
func (r *MemberInvitationService) List(ctx context.Context, opts ...option.RequestOption) (res *MemberInvitationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/invitations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Marks the invitation revoked. The token becomes unusable. A fresh invite can be
// issued for the same email afterwards. Requires the `members:write` scope.
func (r *MemberInvitationService) Revoke(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/invitations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type MemberInvitationListResponse struct {
	Invitations []Invitation                     `json:"invitations" api:"required"`
	JSON        memberInvitationListResponseJSON `json:"-"`
}

// memberInvitationListResponseJSON contains the JSON metadata for the struct
// [MemberInvitationListResponse]
type memberInvitationListResponseJSON struct {
	Invitations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberInvitationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberInvitationListResponseJSON) RawJSON() string {
	return r.raw
}

type MemberInvitationNewParams struct {
	Email param.Field[string]                        `json:"email" api:"required" format:"email"`
	Role  param.Field[MemberInvitationNewParamsRole] `json:"role" api:"required"`
}

func (r MemberInvitationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberInvitationNewParamsRole string

const (
	MemberInvitationNewParamsRoleOwner   MemberInvitationNewParamsRole = "owner"
	MemberInvitationNewParamsRoleAdmin   MemberInvitationNewParamsRole = "admin"
	MemberInvitationNewParamsRoleMember  MemberInvitationNewParamsRole = "member"
	MemberInvitationNewParamsRoleBilling MemberInvitationNewParamsRole = "billing"
	MemberInvitationNewParamsRoleViewer  MemberInvitationNewParamsRole = "viewer"
)

func (r MemberInvitationNewParamsRole) IsKnown() bool {
	switch r {
	case MemberInvitationNewParamsRoleOwner, MemberInvitationNewParamsRoleAdmin, MemberInvitationNewParamsRoleMember, MemberInvitationNewParamsRoleBilling, MemberInvitationNewParamsRoleViewer:
		return true
	}
	return false
}
