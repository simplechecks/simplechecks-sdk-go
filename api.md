# Healthz

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go">simplechecks</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#HealthzCheckResponse">HealthzCheckResponse</a>

Methods:

- <code title="get /healthz">client.Healthz.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#HealthzService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go">simplechecks</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#HealthzCheckResponse">HealthzCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Account

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go">simplechecks</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#AccountGetResponse">AccountGetResponse</a>

Methods:

- <code title="get /v1/account">client.Account.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go">simplechecks</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#AccountGetResponse">AccountGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Checks

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go">simplechecks</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#Check">Check</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go">simplechecks</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#CheckListResponse">CheckListResponse</a>

Methods:

- <code title="post /v1/checks">client.Checks.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#CheckService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go">simplechecks</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#CheckNewParams">CheckNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go">simplechecks</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#Check">Check</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/checks/{id}">client.Checks.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#CheckService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go">simplechecks</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#Check">Check</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/checks/{id}">client.Checks.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#CheckService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go">simplechecks</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#CheckUpdateParams">CheckUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go">simplechecks</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#Check">Check</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/checks">client.Checks.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#CheckService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go">simplechecks</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#CheckListParams">CheckListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go">simplechecks</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#CheckListResponse">CheckListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/checks/{id}">client.Checks.<a href="https://pkg.go.dev/github.com/stainless-sdks/simplechecks-go#CheckService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
