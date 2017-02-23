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

##Single Engine

If you only have one Tamber engine, and therefore one API Key, simply import the packages you would like to use, and use the following pattern:

```go
import (
    tamber "github.com/tamber/tamber-go"
    "github.com/tamber/tamber-go/event"
    "fmt"
)

tamber.DefaultProjectKey = "Mu6DUPXdDYe98cv5JIfX"

e, err := event.Track(&tamber.EventParams{
    User: "user_rlox8k927z7p",
    Behavior: "click",
    Item: "item_wmt4fn6o4zlk",
})

if err != nil {
   //Handle
}

tamber.DefaultEngineKey = "SbWYPBNdARfIDa0IIO9L"

recommendations, err := discover.Recommended(&tamber.DiscoverParams{
    User: "user_rlox8k927z7p",
})

if err != nil {
   //Handle
}

for _, rec := range recommendations{
    fmt.Printf("Item Id:%s :: Score:%f", rec.Item, rec.Score)
}
```

##Multiple Engines

If you have multiple Tamber engines, use the engine package.

```go
import (
    "github.com/tamber/tamber-go"
    "github.com/tamber/tamber-go/client"
    "github.com/tamber/tamber-go/event"
)

c := client.New("Mu6DUPXdDYe98cv5JIfX", "SbWYPBNdARfIDa0IIO9L", nil)

e, err := c.Event.Track(&tamber.EventParams{
    User: "user_rlox8k927z7p",
    Behavior: "click",
    Item: "item_wmt4fn6o4zlk",
})

if err != nil {
   //Handle
}
```

See [test.go](https://github.com/tamber/tamber-go/blob/master/test/test.go) for more examples.

