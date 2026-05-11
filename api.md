# Account

Response Types:

- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Account">Account</a>

Methods:

- <code title="get /v1/account">client.Account.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Checks

Params Types:

- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#AlertChannelParam">AlertChannelParam</a>
- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#AlertConfigParam">AlertConfigParam</a>
- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#MaintenanceWindowParam">MaintenanceWindowParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#AlertChannel">AlertChannel</a>
- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#AlertConfig">AlertConfig</a>
- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Check">Check</a>
- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#MaintenanceWindow">MaintenanceWindow</a>

Methods:

- <code title="post /v1/checks">client.Checks.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckNewParams">CheckNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Check">Check</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/checks/{id}">client.Checks.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Check">Check</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/checks/{id}">client.Checks.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckUpdateParams">CheckUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Check">Check</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/checks">client.Checks.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckListParams">CheckListParams</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go/packages/pagination#Offset">Offset</a>[<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Check">Check</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/checks/{id}">client.Checks.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Alerts

Response Types:

- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckAlertTestFireResponse">CheckAlertTestFireResponse</a>

Methods:

- <code title="get /v1/checks/{id}/alerts">client.Checks.Alerts.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckAlertService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#AlertConfig">AlertConfig</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/checks/{id}/alerts">client.Checks.Alerts.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckAlertService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /v1/checks/{id}/alerts">client.Checks.Alerts.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckAlertService.Replace">Replace</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckAlertReplaceParams">CheckAlertReplaceParams</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#AlertConfig">AlertConfig</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/checks/{id}/alerts:test">client.Checks.Alerts.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckAlertService.TestFire">TestFire</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckAlertTestFireResponse">CheckAlertTestFireResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Runs

Response Types:

- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Aggregate">Aggregate</a>
- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Run">Run</a>
- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#RunListResponse">RunListResponse</a>
- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#RunAggregatesResponse">RunAggregatesResponse</a>

Methods:

- <code title="get /v1/runs/{id}">client.Runs.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#RunService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Run">Run</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/runs">client.Runs.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#RunService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#RunListParams">RunListParams</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#RunListResponse">RunListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/runs/aggregates">client.Runs.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#RunService.Aggregates">Aggregates</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#RunAggregatesParams">RunAggregatesParams</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#RunAggregatesResponse">RunAggregatesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/runs/{id}/logs">client.Runs.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#RunService.Logs">Logs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/builtin#io.Reader">io.Reader</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Incidents

Response Types:

- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Incident">Incident</a>
- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#IncidentListResponse">IncidentListResponse</a>

Methods:

- <code title="get /v1/incidents">client.Incidents.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#IncidentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#IncidentListParams">IncidentListParams</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#IncidentListResponse">IncidentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Keys

Response Types:

- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#APIKey">APIKey</a>
- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#KeyNewResponse">KeyNewResponse</a>
- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#KeyListResponse">KeyListResponse</a>

Methods:

- <code title="post /v1/keys">client.Keys.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#KeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#KeyNewParams">KeyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#KeyNewResponse">KeyNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/keys">client.Keys.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#KeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#KeyListResponse">KeyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/keys/{id}">client.Keys.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#KeyService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Balance

Response Types:

- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Balance">Balance</a>

Methods:

- <code title="get /v1/balance">client.Balance.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#BalanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Balance">Balance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CheckoutSessions

Response Types:

- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckoutSession">CheckoutSession</a>

Methods:

- <code title="post /v1/checkout-session">client.CheckoutSessions.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckoutSessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckoutSessionNewParams">CheckoutSessionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckoutSession">CheckoutSession</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Members

Response Types:

- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Invitation">Invitation</a>
- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Member">Member</a>
- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#MemberListResponse">MemberListResponse</a>

Methods:

- <code title="patch /v1/members/{user_id}">client.Members.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#MemberService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#MemberUpdateParams">MemberUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Member">Member</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/members">client.Members.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#MemberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#MemberListResponse">MemberListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/members/{user_id}">client.Members.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#MemberService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Invitations

Response Types:

- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#MemberInvitationListResponse">MemberInvitationListResponse</a>

Methods:

- <code title="post /v1/invitations">client.Members.Invitations.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#MemberInvitationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#MemberInvitationNewParams">MemberInvitationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Invitation">Invitation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/invitations">client.Members.Invitations.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#MemberInvitationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#MemberInvitationListResponse">MemberInvitationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/invitations/{id}">client.Members.Invitations.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#MemberInvitationService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Locations

Response Types:

- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Location">Location</a>
- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#LocationListResponse">LocationListResponse</a>

Methods:

- <code title="get /v1/locations">client.Locations.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#LocationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#LocationListResponse">LocationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
