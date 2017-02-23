package behavior

import (
	"errors"
	tamber "github.com/tamber/tamber-go"
	"net/url"
)

type Client struct {
	S          *tamber.SessionConfig
	ProjectKey string
	EngineKey  string
}

var object = "behavior"

func Create(params *tamber.BehaviorParams) (*tamber.Behavior, error) {
	return getClient().Create(params)
}

func (c Client) Create(params *tamber.BehaviorParams) (*tamber.Behavior, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	behavior := &tamber.BehaviorResponse{}
	var err error

	if len(params.Name) > 0 && params.Desirability > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "create", body, behavior)
	} else {
		err = errors.New("Invalid behavior params: name, type, and desirability need to be set")
	}

	if !behavior.Succ {
		err = errors.New(behavior.Error)
	}
	return &behavior.Result, err
}

func Retrieve(params *tamber.BehaviorParams) (*tamber.Behavior, error) {
	return getClient().Retrieve(params)
}

func (c Client) Retrieve(params *tamber.BehaviorParams) (*tamber.Behavior, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	behavior := &tamber.BehaviorResponse{}
	var err error

	if len(params.Name) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "retrieve", body, behavior)
	} else {
		err = errors.New("Invalid behavior params: name needs to be set")
	}

	if !behavior.Succ {
		err = errors.New(behavior.Error)
	}
	return &behavior.Result, err
}

func getClient() Client {
	return Client{tamber.GetDefaultSessionConfig(), tamber.DefaultProjectKey, tamber.DefaultEngineKey}
}
