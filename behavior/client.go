package behavior

import (
	"errors"
	tamber "github.com/tamber/tamber-go"
	"net/url"
)

type Engine struct {
	S   *tamber.SessionConfig
	Key string
}

var object = "behavior"

func Create(params *tamber.BehaviorParams) (*tamber.Behavior, error) {
	return getEngine().Create(params)
}

func (e Engine) Create(params *tamber.BehaviorParams) (*tamber.Behavior, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	behavior := &tamber.BehaviorResponse{}
	var err error

	if len(params.Name) > 0 && len(params.Type) > 0 && params.Desirability > 0 {
		err = e.S.Call("POST", "", e.Key, object, "create", body, behavior)
	} else {
		err = errors.New("Invalid behavior params: name, type, and desirability need to be set")
	}

	if !behavior.Succ {
		err = errors.New(behavior.Error)
	}
	return &behavior.Result, err
}

func Retrieve(params *tamber.BehaviorParams) (*tamber.Behavior, error) {
	return getEngine().Create(params)
}

func (e Engine) Retrieve(params *tamber.BehaviorParams) (*tamber.Behavior, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	behavior := &tamber.BehaviorResponse{}
	var err error

	if len(params.Name) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "retrieve", body, behavior)
	} else {
		err = errors.New("Invalid behavior params: name needs to be set")
	}

	if !behavior.Succ {
		err = errors.New(behavior.Error)
	}
	return &behavior.Result, err
}

func Remove(params *tamber.BehaviorParams) (*tamber.Behavior, error) {
	return getEngine().Create(params)
}

func (e Engine) Remove(params *tamber.BehaviorParams) (*tamber.Behavior, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	behavior := &tamber.BehaviorResponse{}
	var err error

	if len(params.Name) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "remove", body, behavior)
	} else {
		err = errors.New("Invalid behavior params: name needs to be set")
	}

	if !behavior.Succ {
		err = errors.New(behavior.Error)
	}
	return &behavior.Result, err
}

func getEngine() Engine {
	return Engine{tamber.GetDefaultSessionConfig(), tamber.DefaultKey}
}
