# Account

Response Types:

- <a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Account">Account</a>

Methods:

- <code title="get /v1/account">client.Account.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go">simplechecksgo</a>.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Checks

Methods:

- <code title="delete /v1/checks/{id}">client.Checks.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Alerts

Methods:

- <code title="delete /v1/checks/{id}/alerts">client.Checks.Alerts.<a href="https://pkg.go.dev/github.com/simplechecks/simplechecks-sdk-go#CheckAlertService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

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

# Keys

Methods:

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
