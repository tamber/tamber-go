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

Tamber learns from user behaviors, so to get started all you need to do is track Events (user-item interactions) just as you would for any analytics service. Then you can initialize learning by launching an engine in the [dashboard][dashboard] and start discovering recommendations!

## Track real time events

Track all events (user-item interactions in your app like 'clicked', 'shared', 'purchased', etc.) to your project in real time, just like you would for a data analytics service. Note that novel users and items will automatically be created.

We recommend performing your event tracking from the frontend to catch all user actions and their contexts as part of normal action handling. If you would like to handle this in the frontend, checkout our [other SDKs][sdks], including [Node][tamber-node], [iOS][tamber-ios], [Javascript][tamber.js], and [Android][tamber-android].

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
```

### Discover

Once you have tracked enough events and created your engine, you may begin using `discover` to put personalized recommendations in your app.

The primary methods of discovery in Tamber are the `discover.Next` and `discover.Recommended` methods. `discover.Next` is often the most impactful tool for driving lift, allowing you to turn your item pages into steps on personalized paths of discovery – it returns the optimal set of items that the user should be shown next on a given item page.

`discover.Recommended` works similarly, but is optimized for a recommended section, often located on a homepage.

#### Up Next

Keep users engaged by creating a path of discovery as they navigate from item to item, always showing the right mix of items they should check out next. Just set the user's id and the id of the item that they are navigating to / looking at.

```go
import (
    tamber "github.com/tamber/tamber-go"
    "github.com/tamber/tamber-go/discover"
    "fmt"
)

tamber.DefaultProjectKey = "Mu6DUPXdDYe98cv5JIfX" 

// Be sure to set the default engine for your project.
// Otherwise you can also set the engine manually:
// tamber.DefaultEngineKey = "SbWYPBNdARfIDa0IIO9L"

// Get items to display directly to the user on a given item page
recommendations, info, err := discover.Next(&tamber.DiscoverParams{
    User: "user_rlox8k927z7p",
    Item: "item_wmt4fn6o4zlk",
    Number: 8,
})

if err != nil {
   //Handle
}

for _, rec := range recommendations{
    fmt.Printf("Item Id:%s :: Score:%f", rec.Item, rec.Score)
}
```

#### Recommended

To put personalized recommendations on your homepage, or in any recommended section, just call `discover.Recommended` with the user's id and the number of recommendations you want to display.


```go
import (
    tamber "github.com/tamber/tamber-go"
    "github.com/tamber/tamber-go/discover"
)

tamber.DefaultProjectKey = "Mu6DUPXdDYe98cv5JIfX" 

// Get items to display directly to the user
recommendations, info, err := discover.Recommended(&tamber.DiscoverParams{
    User: "user_rlox8k927z7p",
})
```

#### Weekly and Daily Periodicals

Instantly deploy your own Spotify-style Discover Weekly feature, or a daily periodical with fresh recommendations updated every 24 hours.

```js
recommendations, info, err := discover.Weekly(&tamber.DiscoverParams{
    User: "user_rlox8k927z7p",
    Number: 35,
})

recommendations, info, err := discover.Daily(&tamber.DiscoverParams{
    User: "user_rlox8k927z7p",
    Number: 35,
})
```

#### Build Your Own Features

Tamber allows you to use lower-level methods to get lists of recommended items, similar item matches, and similar items for a given user with which you can build your own discovery experiences. Importantly, these methods return raw recommendation data and are not intended to be pushed directly to users.

```js
recommendations, info, err := discover.Basic.Recommended(&tamber.DiscoverBasicParams{
    User: "user_rlox8k927z7p",
})

recommendations, info, err := discover.Basic.Similar(&tamber.DiscoverBasicParams{
    Item: "item_wmt4fn6o4zlk",
})

recommendations, info, err := discover.Basic.RecommendedSimilar(&tamber.DiscoverBasicParams{
    User: "user_rlox8k927z7p",
    Item: "item_wmt4fn6o4zlk",
})
```

Features
========

The Tamber client library provides additional features that make it easy to build and run your engines.

## Create historical events dataset

If you have historical events data you would like to upload to your project, the `tamber-go` library makes it easy to stream events to a csv file, ready for upload ([head here][historic-events] for proper instructions).

```go
package main

import (
    tamber "github.com/tamber/tamber-go"
    "github.com/tamber/tamber-go/event"
    "fmt"
    DB "myapp/database"
)

const (
    BatchSize = 1000
    EventsFilepath = "./events.csv"
    MaxTimestamp = 1524101472 // Unix Timestamp of when we began streaming real time events
)

func main() {
    for offset := 0; ; offset += BatchSize {
        events, err := DB.LoadEvents(BatchSize, offset, MaxTimestamp) // load events created before MaxTimestamp
        if err != nil {
            panic(err)
        }
        if len(events) == 0 {
            break
        }
        err = event.BatchEventsToCSV(events, EventsFilepath)
        if err != nil {
            panic(err)
        }
    }
    tamber.Gzip(EventsFilepath) // saves to EventsFilepath + ".gz"
}
```

## Stream Items

If you want to add properties or tags to your items, the Stream method allows you to efficiently stream item updates. By default, item updates will automatically create novel items (you may deactivate this behavior by setting the `NoCreate` field to false).

```go
import (
    tamber "github.com/tamber/tamber-go"
    tamber_item "github.com/tamber/tamber-go/item"
    "fmt"
)

const (
    NUM_THREADS = 10
    BUF_SIZE    = 64 * 1024
)

tamber.DefaultProjectKey = "Mu6DUPXdDYe98cv5JIfX"

items := Database.LoadItems()
itemUpdates := make([]*tamber.ItemUpdateParams, len(items))
for i, item := range items {
    itemUpdates[i] = &tamber.ItemUpdateParams{Id: item.Id, Updates: tamber.ItemUpdates{Add: tamber.ItemFeatures{Properties: item.Properties}}}
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

## Multiple Projects / Engines

If you have multiple Tamber projects or engines, use the client module to separate instances.

```go
import (
    "github.com/tamber/tamber-go"
    "github.com/tamber/tamber-go/client"
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

### API Response Info

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

[homepage]: https://tamber.com/
[docs]: https://tamber.com/docs/
[dashboard]: https://dashboard.tamber.com/
[quickstart]: https://tamber.com/docs/start/
[historic-events]: https://tamber.com/docs/start/#upload-history
[sdks]: https://tamber.com/docs/libs/
[tamber-node]: https://github.com/tamber/tamber-node
[tamber-ios]: https://github.com/tamber/tamber-ios
[tamber.js]: https://github.com/tamber/tamber.js
[tamber-android]: https://github.com/tamber/tamber-android
[tamber-ruby]: https://github.com/tamber/tamber-ruby
[tamber-go]: https://github.com/tamber/tamber-go
[tamber-python]: https://github.com/tamber/tamber-python
[tamber-java]: https://github.com/tamber/tamber-java