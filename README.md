# shopify-go-api

The new home of Conversio's Shopify Go library.

forked from [bold-commerce/go-shopify](https://github.com/bold-commerce/go-shopify)

## Supported Go Versions

This library has been tested against the following versions of Go
* 1.10
* 1.11
* 1.12
* 1.13
* 1.14
* 1.15

## Install

```console
$ go get github.com/youbudong/shopify-go-api
```

## Use

```go
import "github.com/youbudong/shopify-go-api"
```

This gives you access to the `shopify` package.

#### Oauth

If you don't have an access token yet, you can obtain one with the oauth flow.
Something like this will work:

```go
// Create an app somewhere.
app := shopify.App{
    ApiKey: "abcd",
    ApiSecret: "efgh",
    RedirectUrl: "https://example.com/shopify/callback",
    Scope: "read_products,read_orders",
}

// Create an oauth-authorize url for the app and redirect to it.
// In some request handler, you probably want something like this:
func MyHandler(w http.ResponseWriter, r *http.Request) {
    shopName := r.URL.Query().Get("shop")
    state := "nonce"
    authUrl := app.AuthorizeUrl(shopName, state)
    http.Redirect(w, r, authUrl, http.StatusFound)
}

// Fetch a permanent access token in the callback
func MyCallbackHandler(w http.ResponseWriter, r *http.Request) {
    // Check that the callback signature is valid
    if ok, _ := app.VerifyAuthorizationURL(r.URL); !ok {
        http.Error(w, "Invalid Signature", http.StatusUnauthorized)
        return
    }

    query := r.URL.Query()
    shopName := query.Get("shop")
    code := query.Get("code")
    token, err := app.GetAccessToken(shopName, code)

    // Do something with the token, like store it in a DB.
}
```

#### Api calls with a token

With a permanent access token, you can make API calls like this:

```go
// Create an app somewhere.
app := shopify.App{
    ApiKey: "abcd",
    ApiSecret: "efgh",
    RedirectUrl: "https://example.com/shopify/callback",
    Scope: "read_products",
}

// Create a new API client
client := shopify.NewClient(app, "shopname", "token")

// Fetch the number of products.
numProducts, err := client.Product.Count(nil)
```

#### Private App Auth

Private Shopify apps use basic authentication and do not require going through the OAuth flow. Here is an example:

```go
// Create an app somewhere.
app := shopify.App{
	ApiKey: "apikey",
	Password: "apipassword",
}

// Create a new API client (notice the token parameter is the empty string)
client := shopify.NewClient(app, "shopname", "")

// Fetch the number of products.
numProducts, err := client.Product.Count(nil)
```
### Client Options
When creating a client there are configuration options you can pass to NewClient. Simply use the last variadic param and 
pass in the built in options or create your own and manipulate the client. See [options.go](https://github.com/bold-commerce/go-shopify/blob/master/options.go)
for more details.

#### WithVersion
Read more details on the [Shopify API Versioning](https://shopify.dev/concepts/about-apis/versioning)
to understand the format and release schedules. You can use `WithVersion` to specify a specific version 
of the API. If you do not use this option you will be defaulted to the oldest stable API.

```go
client := shopify.NewClient(app, "shopname", "", shopify.WithVersion("2019-04"))
```

#### WithRetry
Shopify [Rate Limits](https://shopify.dev/concepts/about-apis/rate-limits) their API and if this happens to you they 
will send a back off (usually 2s) to tell you to retry your request. To support this functionality seamlessly within 
the client a `WithRetry` option exists where you can pass an `int` of how many times you wish to retry per-request 
before returning an error. `WithRetry` additionally supports retrying HTTP503 errors.

```go
client := shopify.NewClient(app, "shopname", "", shopify.WithRetry(3))
```

#### Query options

Most API functions take an options `interface{}` as parameter. You can use one
from the library or create your own. For example, to fetch the number of
products created after January 1, 2016, you can do:

```go
// Create standard CountOptions
date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
options := shopify.CountOptions{createdAtMin: date}

// Use the options when calling the API.
numProducts, err := client.Product.Count(options)
```

The options are parsed with Google's
[go-querystring](https://github.com/google/go-querystring) library so you can
use custom options like this:

```go
// Create custom options for the orders.
// Notice the `url:"status"` tag
options := struct {
    Status string `url:"status"`
}{"any"}

// Fetch the order count for orders with status="any"
orderCount, err := client.Order.Count(options)
```

#### Using your own models

Not all endpoints are implemented right now. In those case, feel free to
implement them and make a PR, or you can create your own struct for the data
and use `NewRequest` with the API client. This is how the existing endpoints
are implemented.

For example, let's say you want to fetch webhooks. There's a helper function
`Get` specifically for fetching stuff so this will work:

```go
// Declare a model for the webhook
type Webhook struct {
    ID int         `json:"id"`
    Address string `json:"address"`
}

// Declare a model for the resource root.
type WebhooksResource struct {
    Webhooks []Webhook `json:"webhooks"`
}

func FetchWebhooks() ([]Webhook, error) {
    path := "admin/webhooks.json"
    resource := new(WebhooksResource)
    client := shopify.NewClient(app, "shopname", "token")

    // resource gets modified when calling Get
    err := client.Get(path, resource, nil)

    return resource.Webhooks, err
}
```

#### Webhooks verification

In order to be sure that a webhook is sent from ShopifyApi you could easily verify
it with the `VerifyWebhookRequest` method.

For example:
```go
func ValidateWebhook(httpRequest *http.Request) (bool) {
    shopifyApp := shopify.App{ApiSecret: "ratz"}
    return shopifyApp.VerifyWebhookRequest(httpRequest)
}
```

## Develop and test
`docker` and `docker-compose` must be installed

