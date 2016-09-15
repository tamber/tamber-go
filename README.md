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
    "github.com/tamber/tamber-go/discover"
    "fmt"
)

tamber.DefaultKey = "key_sBW1WHQ4bP4Ryfz3AQOo"

e, err := event.Track(&tamber.EventParams{
    User: "user_rlox8k927z7p",
    Behavior: "click",
    Item: "item_wmt4fn6o4zlk",
})

if err != nil {
   //Handle
}

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
    "github.com/tamber/tamber-go/engine"
    "github.com/tamber/tamber-go/event"
)

e := &engine.API{}
e.Init("80r2oX10Uw4XfZSxfh4O", nil)

e, err := e.Event.Track(&tamber.EventParams{
    User: "user_rlox8k927z7p",
    Behavior: "click",
    Item: "item_wmt4fn6o4zlk",
})

if err != nil {
   //Handle
}
```

See [test.go](https://github.com/tamber/tamber-go/blob/master/test/test.go) for more examples.

