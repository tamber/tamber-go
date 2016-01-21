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
    "github.com/tamber/tamber-go/actor"
)

tamber.DefaultKey = "sBW1WHQ4bP4Ryfz3AQOo"

a, err := actor.AddBehaviors(&tamber.ActorParams{
    Id: "2197054086",
    Behaviors: &[]tamber.ActorBehavior{
        tamber.ActorBehavior{
            Behavior: "like",
            Item:     "HZNP",
            Value:    1.0,
            Created:  1446417346,
        },
    },
})
if err != nil {
   //Handle
}
```

##Multiple Engines

If you have multiple Tamber engines, use the engine package.

```go
import (
    "github.com/tamber/tamber-go"
    "github.com/tamber/tamber-go/engine"
    "github.com/tamber/tamber-go/actor"
)

e := &engine.API{}
e.Init("80r2oX10Uw4XfZSxfh4O", nil)
a, err := e.Actors.AddBehaviors(&tamber.ActorParams{
    Id: "2197054086",
    Behaviors: &[]tamber.ActorBehavior{
        tamber.ActorBehavior{
            Behavior: "like",
            Item:     "HZNP",
            Value:    1.0,
            Created:  1446417346,
        },
    },
})
if err != nil {
   //Handle
}
```

See [test.go](https://github.com/tamber/tamber-go/blob/master/test/test.go) for more examples.

