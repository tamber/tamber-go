package engine

import (
	. "github.com/tamber/tamber-go"
	"github.com/tamber/tamber-go/behavior"
	"github.com/tamber/tamber-go/discover"
	"github.com/tamber/tamber-go/event"
	"github.com/tamber/tamber-go/item"
	"github.com/tamber/tamber-go/user"
)

type API struct {
	Event    *event.Engine
	Discover *discover.Engine
	User     *user.Engine
	Item     *item.Engine
	Behavior *behavior.Engine
}

func (a *API) Init(key string, config *SessionConfig) {
	if config == nil {
		config = GetDefaultSessionConfig()
	}
	a.Event = &event.Engine{S: config, Key: key}
	a.Discover = &discover.Engine{S: config, Key: key}
	a.User = &user.Engine{S: config, Key: key}
	a.Item = &item.Engine{S: config, Key: key}
	a.Behavior = &behavior.Engine{S: config, Key: key}
}

// New creates a new Tamber Engine object the appropriate key
// as well as providing the ability to override the backends as needed.
func New(key string, config *SessionConfig) *API {
	api := API{}
	api.Init(key, config)
	return &api
}
