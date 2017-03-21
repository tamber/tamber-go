# Tamber API Client for Go

You can sign up for a Tamber account at https://tamber.com.

For full API documentation, refer to https://tamber.com/docs/api.

Installation
============

```sh
go get github.com/tamber/tamber-go
```

Usage
=====

There are two ways to use Tamber:

##Single Project / Engine

If you only have one Tamber project and/or engine, simply import the packages you would like to use, and use the following pattern:

```go
import (
    tamber "github.com/tamber/tamber-go"
    "github.com/tamber/tamber-go/event"
    "fmt"
)

tamber.DefaultProjectKey = "Mu6DUPXdDYe98cv5JIfX"

e, info, err := event.Track(&tamber.EventParams{
    User: "user_rlox8k927z7p",
    Behavior: "click",
    Item: "item_wmt4fn6o4zlk",
})

if err != nil {
   //Handle
}

tamber.DefaultEngineKey = "SbWYPBNdARfIDa0IIO9L" // Discover endpoint requires engines

recommendations, info, err := discover.Recommended(&tamber.DiscoverParams{
    User: "user_rlox8k927z7p",
})

if err != nil {
   //Handle
}

for _, rec := range recommendations{
    fmt.Printf("Item Id:%s :: Score:%f", rec.Item, rec.Score)
}
```

##Stream Items

If you want to add properties or tags to your items, the Stream method allows you to efficiently stream item updates.

```go
import (
    tamber "github.com/tamber/tamber-go"
    "github.com/tamber/tamber-go/event"
    "fmt"
)

const (
    NUM_THREADS = 10
    BUF_SIZE    = 64 * 1024
)

tamber.DefaultProjectKey = "Mu6DUPXdDYe98cv5JIfX"

items := Database.LoadItems()
itemUpdates := make([]*tamber.ItemParams, len(items))
for i, item := range items {
    itemUpdates[i] = &tamber.ItemParams{Id: item.Id, Updates: &tamber.ItemUpdates{Add: tamber.ItemFeatures{Properties: item.Properties}}}
}

// You may optionally supply a channel to read updated items.
out := make(chan *tamber.Item, BUF_SIZE)
go func() {
    for {
        select {
        case item := <-out:
            fmt.Println("updated item:", *item)
        }
    }
}()
info, err := tamber_item.Stream(itemUpdates, &out, N_THREADS, BUF_SIZE)
if err != nil {
    //Handle
}
```

##Multiple Projects / Engines

If you have multiple Tamber projects or engines, use the client module to separate instances.

```go
import (
    "github.com/tamber/tamber-go"
    "github.com/tamber/tamber-go/client"
    "github.com/tamber/tamber-go/event"
)

c := client.New("Mu6DUPXdDYe98cv5JIfX", "SbWYPBNdARfIDa0IIO9L", nil)

e, info, err := c.Event.Track(&tamber.EventParams{
    User: "user_rlox8k927z7p",
    Behavior: "click",
    Item: "item_wmt4fn6o4zlk",
})

if err != nil {
   //Handle
}
```

###API Response Info

The Tamber API includes useful HTTP status codes and headers in its responses. The ResponseInfo type provides access to these values, and is returned by all methods (see the `info` value in the examples).

```
type ResponseInfo struct {
    HTTPCode           int // HTTP status code
    RateLimit          int // Limit-per-period for request method
    RateLimitRemaining int // Requests remaining in current window for request method
    RateLimitReset     int // Time in seconds until rate limits are reset
}
```

See [test.go](https://github.com/tamber/tamber-go/blob/master/test/test.go) for more examples.

