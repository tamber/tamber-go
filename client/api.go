package client

import (
	"github.com/tamber/tamber-go"
	"github.com/tamber/tamber-go/behavior"
	"github.com/tamber/tamber-go/discover"
	"github.com/tamber/tamber-go/event"
	"github.com/tamber/tamber-go/item"
	"github.com/tamber/tamber-go/user"
)

type API struct {
	Event      *event.Client
	Discover   *discover.Client
	User       *user.Client
	Item       *item.Client
	Behavior   *behavior.Client
	ProjectKey string
	EngineKey  string
}

func (a *API) Init(projectKey string, engineKey string, config *tamber.SessionConfig) {
	if config == nil {
		config = tamber.GetDefaultSessionConfig()
	}
	a.Event = &event.Client{S: config, ProjectKey: projectKey, EngineKey: engineKey}
	a.Discover = &discover.Client{S: config, ProjectKey: projectKey, EngineKey: engineKey}
	a.User = &user.Client{S: config, ProjectKey: projectKey, EngineKey: engineKey}
	a.Item = &item.Client{S: config, ProjectKey: projectKey, EngineKey: engineKey}
	a.Behavior = &behavior.Client{S: config, ProjectKey: projectKey, EngineKey: engineKey}
	a.ProjectKey = projectKey
	a.EngineKey = engineKey
}

// New creates a new Tamber Engine object the appropriate key
// as well as providing the ability to override the backends as needed.
func New(projectKey string, engineKey string, config *tamber.SessionConfig) *API {
	api := API{}
	api.Init(projectKey, engineKey, config)
	return &api
}
