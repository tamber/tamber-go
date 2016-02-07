package user

import (
	"errors"
	tamber "github.com/tamber/tamber-go"
	"net/url"
)

type Engine struct {
	S   *tamber.SessionConfig
	Key string
}

var object = "user"

func Create(params *tamber.UserParams) (*tamber.User, error) {
	return getEngine().Create(params)
}

func (e Engine) Create(params *tamber.UserParams) (*tamber.User, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	user := &tamber.UserResponse{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "create", body, user)
	} else {
		err = errors.New("Invalid user params: id needs to be set")
	}

	if !user.Succ {
		err = errors.New(user.Error)
	}
	return &user.Result, err
}

func Update(params *tamber.UserParams) (*tamber.User, error) {
	return getEngine().Update(params)
}

func (e Engine) Update(params *tamber.UserParams) (*tamber.User, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	user := &tamber.UserResponse{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "update", body, user)
	} else {
		err = errors.New("Invalid user params: id needs to be set")
	}

	if !user.Succ {
		err = errors.New(user.Error)
	}
	return &user.Result, err
}

func Retrieve(params *tamber.UserParams) (*tamber.User, error) {
	return getEngine().Retrieve(params)
}

func (e Engine) Retrieve(params *tamber.UserParams) (*tamber.User, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	user := &tamber.UserResponse{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "retrieve", body, user)
	} else {
		err = errors.New("Invalid user params: id needs to be set")
	}

	if !user.Succ {
		err = errors.New(user.Error)
	}
	return &user.Result, err
}

func getEngine() Engine {
	return Engine{tamber.GetDefaultSessionConfig(), tamber.DefaultKey}
}
