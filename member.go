// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/simplechecks/simplechecks-sdk-go/internal/apijson"
	"github.com/simplechecks/simplechecks-sdk-go/internal/param"
	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
)

// Manage who has access to an account and at what role (PR-Members/2). Five roles:
// owner / admin / member / billing / viewer. Owner is the strict superset of all
// other roles' scopes; every account always has at least one owner.
//
// MemberService contains methods and other services that help with interacting
// with the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMemberService] method instead.
type MemberService struct {
	Options []option.RequestOption
	// Manage who has access to an account and at what role (PR-Members/2). Five roles:
	// owner / admin / member / billing / viewer. Owner is the strict superset of all
	// other roles' scopes; every account always has at least one owner.
	Invitations *MemberInvitationService
}

// NewMemberService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMemberService(opts ...option.RequestOption) (r *MemberService) {
	r = &MemberService{}
	r.Options = opts
	r.Invitations = NewMemberInvitationService(opts...)
	return
}

// Sets the member's role. Refuses to demote the last owner; the webapp surfaces
// this as "promote another owner first." Cannot modify your own role — ask another
// owner to do it. Requires the `members:write` scope.
func (r *MemberService) Update(ctx context.Context, userID string, body MemberUpdateParams, opts ...option.RequestOption) (res *Member, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/members/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns every (user, role, joined_at) tuple for the caller's account. Ordered
// owner-first (oldest membership). Backs the Settings → Members tab in the webapp.
// Requires the `members:read` scope.
func (r *MemberService) List(ctx context.Context, opts ...option.RequestOption) (res *MemberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/members"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deletes the (account, user) membership. Refuses to remove the last owner. Cannot
// remove yourself — use the "leave account" flow instead. Note that this does NOT
// revoke the user's API keys; the webapp orchestrates a follow-up keys:write call
// if the caller wants a hard cut-off. Requires the `members:write` scope.
func (r *MemberService) Remove(ctx context.Context, userID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/members/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A pending invitation to join the account.
type Invitation struct {
	ID              string         `json:"id" api:"required"`
	CreatedAt       time.Time      `json:"created_at" api:"required" format:"date-time"`
	Email           string         `json:"email" api:"required" format:"email"`
	ExpiresAt       time.Time      `json:"expires_at" api:"required" format:"date-time"`
	InvitedByUserID string         `json:"invited_by_user_id" api:"required"`
	Role            InvitationRole `json:"role" api:"required"`
	// Random URL-safe token. Only returned at creation time (POST /v1/invitations);
	// GET responses omit this field.
	Token string `json:"token"`
	// Convenience: the relative path the webapp routes to for redemption. Only present
	// on creation responses.
	AcceptURLPath string         `json:"accept_url_path"`
	JSON          invitationJSON `json:"-"`
}

// invitationJSON contains the JSON metadata for the struct [Invitation]
type invitationJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Email           apijson.Field
	ExpiresAt       apijson.Field
	InvitedByUserID apijson.Field
	Role            apijson.Field
	Token           apijson.Field
	AcceptURLPath   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Invitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invitationJSON) RawJSON() string {
	return r.raw
}

type InvitationRole string

const (
	InvitationRoleOwner   InvitationRole = "owner"
	InvitationRoleAdmin   InvitationRole = "admin"
	InvitationRoleMember  InvitationRole = "member"
	InvitationRoleBilling InvitationRole = "billing"
	InvitationRoleViewer  InvitationRole = "viewer"
)

func (r InvitationRole) IsKnown() bool {
	switch r {
	case InvitationRoleOwner, InvitationRoleAdmin, InvitationRoleMember, InvitationRoleBilling, InvitationRoleViewer:
		return true
	}
	return false
}

// A user's membership in the caller's account.
type Member struct {
	// When the member joined this account.
	CreatedAt time.Time  `json:"created_at" api:"required" format:"date-time"`
	Email     string     `json:"email" api:"required" format:"email"`
	Role      MemberRole `json:"role" api:"required"`
	// UUID of the member.
	UserID string     `json:"user_id" api:"required"`
	JSON   memberJSON `json:"-"`
}

// memberJSON contains the JSON metadata for the struct [Member]
type memberJSON struct {
	CreatedAt   apijson.Field
	Email       apijson.Field
	Role        apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Member) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberJSON) RawJSON() string {
	return r.raw
}

type MemberRole string

const (
	MemberRoleOwner   MemberRole = "owner"
	MemberRoleAdmin   MemberRole = "admin"
	MemberRoleMember  MemberRole = "member"
	MemberRoleBilling MemberRole = "billing"
	MemberRoleViewer  MemberRole = "viewer"
)

func (r MemberRole) IsKnown() bool {
	switch r {
	case MemberRoleOwner, MemberRoleAdmin, MemberRoleMember, MemberRoleBilling, MemberRoleViewer:
		return true
	}
	return false
}

type MemberListResponse struct {
	Members []Member               `json:"members" api:"required"`
	JSON    memberListResponseJSON `json:"-"`
}

// memberListResponseJSON contains the JSON metadata for the struct
// [MemberListResponse]
type memberListResponseJSON struct {
	Members     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberListResponseJSON) RawJSON() string {
	return r.raw
}

type MemberUpdateParams struct {
	Role param.Field[MemberUpdateParamsRole] `json:"role" api:"required"`
}

func (r MemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRole string

const (
	MemberUpdateParamsRoleOwner   MemberUpdateParamsRole = "owner"
	MemberUpdateParamsRoleAdmin   MemberUpdateParamsRole = "admin"
	MemberUpdateParamsRoleMember  MemberUpdateParamsRole = "member"
	MemberUpdateParamsRoleBilling MemberUpdateParamsRole = "billing"
	MemberUpdateParamsRoleViewer  MemberUpdateParamsRole = "viewer"
)

func (r MemberUpdateParamsRole) IsKnown() bool {
	switch r {
	case MemberUpdateParamsRoleOwner, MemberUpdateParamsRoleAdmin, MemberUpdateParamsRoleMember, MemberUpdateParamsRoleBilling, MemberUpdateParamsRoleViewer:
		return true
	}
	return false
}
