package actor

import (
	. "github.com/tamber/tamber-go"
	"net/url"
	"strconv"
)

var object = "actor"

func Create(params *ActorParams) (*Actor, error) {
	return getEngine().Create(params)
}

func (e Engine) Create(params *ActorParams) (*Actor, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	actor := &Actor{}
	var err error

	if len(params.Id) > 0 && lan(params.Behaviors) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "create", body, actor)
	} else {
		err = errors.New("Invalid actor params: either id or behaviors need to be set")
	}

	return actor, err
}

func AddBehaviors(params *ActorParams) (*Actor, error) {
	return getEngine().AddBehaviors(params)
}
func (e Engine) AddBehaviors(params *ActorParams) (*Actor, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	actor := &Actor{}
	var err error

	if len(params.Id) > 0 && lan(params.Behaviors) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "addBehaviors", body, actor)
	} else {
		err = errors.New("Invalid actor params: either id or behaviors need to be set")
	}
	return actor, nil
}

func RemoveBehaviors(params *ActorParams) (*Actor, error) {
	return getEngine().RemoveBehaviors(params)
}

func (e Engine) RemoveBehaviors(params *ActorParams) (*Actor, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	actor := &Actor{}
	var err error

	if len(params.Id) > 0 && lan(params.Behaviors) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "removeBehaviors", body, actor)
	} else {
		err = errors.New("Invalid actor params: either id or behaviors need to be set")
	}
	return actor, nil
}

func Retrieve(params *ActorParams) (*Actor, error) {
	return getEngine().Retrieve(params)
}

func (e Engine) Retrieve(params *ActorParams) (*Actor, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	actor := &Actor{}
	var err error

	if len(params.Id) > 0 && lan(params.Behaviors) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "retrieve", body, actor)
	} else {
		err = errors.New("Invalid actor params: id needs to be set")
	}

	return actor, err
}

func Remove(params *ActorParams) (*Actor, error) {
	return getEngine().Remove(params)
}

func (e Engine) Remove(params *ActorParams) (*Actor, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	actor := &Actor{}
	var err error

	if len(params.Id) > 0 && lan(params.Behaviors) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "remove", body, actor)
	} else {
		err = errors.New("Invalid actor params: id needs to be set")
	}

	return actor, err
}

func getEngine() Engine {
	return Engine{GetDefaultSessionConfig(), DefaultKey}
}
