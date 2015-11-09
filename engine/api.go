package engine

import (
	. "github.com/tamber/tamber-go"
	"github.com/tamber/tamber-go/actor"
	"github.com/tamber/tamber-go/behavior"
	"github.com/tamber/tamber-go/discover"
	"github.com/tamber/tamber-go/item"
	"github.com/tamber/tamber-go/property"
)

type API struct {
	Actors      *actor.Engine
	Items       *item.Engine
	Behaviors   *behavior.Engine
	Properties  *property.Engine
	Discoveries *discover.Engine
}

func (a *API) Init(key string, config *SessionConfig) {
	if config == nil {
		config = GetDefaultSessionConfig()
	}
	a.Actors = &actor.Engine{S: config, Key: key}
	a.Items = &item.Engine{S: config, Key: key}
	a.Behaviors = &behavior.Engine{S: config, Key: key}
	a.Properties = &property.Engine{S: config, Key: key}
	a.Discoveries = &discover.Engine{S: config, Key: key}
}

// New creates a new Tamber Engine object the appropriate key
// as well as providing the ability to override the backends as needed.
func New(key string, config *SessionConfig) *API {
	api := API{}
	api.Init(key, config)
	return &api
}
