package actor

import (
	"errors"
	tamber "github.com/tamber/tamber-go"
	"net/url"
)

type Engine struct {
	S   *tamber.SessionConfig
	Key string
}

var object = "actor"

func Create(params *tamber.ActorParams) (*tamber.Actor, error) {
	return getEngine().Create(params)
}

func (e Engine) Create(params *tamber.ActorParams) (*tamber.Actor, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	actor := &tamber.Actor{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "create", body, actor)
	} else {
		err = errors.New("Invalid actor params: id needs to be set")
	}

	return actor, err
}

func AddBehaviors(params *tamber.ActorParams) (*tamber.Actor, error) {
	return getEngine().AddBehaviors(params)
}
func (e Engine) AddBehaviors(params *tamber.ActorParams) (*tamber.Actor, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	actor := &tamber.Actor{}
	var err error

	if len(params.Id) > 0 && len(*params.Behaviors) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "addBehaviors", body, actor)
	} else {
		err = errors.New("Invalid actor params: either id or behaviors need to be set")
	}
	return actor, err
}

func RemoveBehaviors(params *tamber.ActorParams) (*tamber.Actor, error) {
	return getEngine().RemoveBehaviors(params)
}

func (e Engine) RemoveBehaviors(params *tamber.ActorParams) (*tamber.Actor, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	actor := &tamber.Actor{}
	var err error

	if len(params.Id) > 0 && len(*params.Behaviors) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "removeBehaviors", body, actor)
	} else {
		err = errors.New("Invalid actor params: either id or behaviors need to be set")
	}
	return actor, err
}

func Retrieve(params *tamber.ActorParams) (*tamber.Actor, error) {
	return getEngine().Retrieve(params)
}

func (e Engine) Retrieve(params *tamber.ActorParams) (*tamber.Actor, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	actor := &tamber.Actor{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "retrieve", body, actor)
	} else {
		err = errors.New("Invalid actor params: id needs to be set")
	}

	return actor, err
}

func Remove(params *tamber.ActorParams) (*tamber.Actor, error) {
	return getEngine().Remove(params)
}

func (e Engine) Remove(params *tamber.ActorParams) (*tamber.Actor, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	actor := &tamber.Actor{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "remove", body, actor)
	} else {
		err = errors.New("Invalid actor params: id needs to be set")
	}

	return actor, err
}

func getEngine() Engine {
	return Engine{tamber.GetDefaultSessionConfig(), tamber.DefaultKey}
}
