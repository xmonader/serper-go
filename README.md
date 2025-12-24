# serper-go

A production-ready Go client for the [serper.dev](https://serper.dev) Google Search API.

## Features

- **Robust Design**: Uses the Functional Options pattern for flexible configuration.
- **Mockable**: Defined via `ClientInterface` for easy testing in your own projects.
- **Typed Constants**: Predefined constants for common regions (`gl`) and languages (`hl`).
- **Comprehensive Coverage**: Support for Search, Images, News, Videos, and Places.
- **Error Handling**: Custom `APIError` type for programmatic access to status codes and messages.
- **Context Support**: Full support for `context.Context` for cancellation and timeouts.
- **Zero Dependencies**: Uses only the Go standard library.

## Installation

```bash
go get github.com/ahmedthabet/serper-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"

    "github.com/ahmedthabet/serper-go/serper"
)

func main() {
    // Initialize with optional timeout
    client := serper.NewClient(
        os.Getenv("SERPER_API_KEY"),
        serper.WithTimeout(10*time.Second),
    )

    // Use predefined constants for regions and languages
    resp, err := client.Search(context.Background(), &serper.Request{
        Q:  "Apple Inc",
        Gl: serper.GLUnitedStates,
        Hl: serper.HLEnglish,
    })
    if err != nil {
        panic(err)
    }

    for _, r := range resp.Organic {
        fmt.Printf("%s - %s\n", r.Title, r.Link)
    }
}
```

## Advanced Configuration

The client supports several functional options:

```go
client := serper.NewClient(apiKey,
    serper.WithHTTPClient(customClient),
    serper.WithBaseURL("https://your-proxy.com"),
    serper.WithTimeout(5*time.Second),
)
```

## Error Handling

You can check for specific API errors:

```go
resp, err := client.Search(ctx, req)
if err != nil {
    var apiErr *serper.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("API Error: %s (Status: %d)\n", apiErr.Message, apiErr.StatusCode)
    }
}
```

## Mocking for Tests

Because the client implements `ClientInterface`, you can easily mock it in your application's unit tests.

## License

MIT
