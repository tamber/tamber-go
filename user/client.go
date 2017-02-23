package user

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

var object = "user"

func Create(params *tamber.UserParams) (*tamber.User, error) {
	return getClient().Create(params)
}

func (c Client) Create(params *tamber.UserParams) (*tamber.User, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	user := &tamber.UserResponse{}
	var err error

	if len(params.Id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "create", body, user)
	} else {
		err = errors.New("Invalid user params: id needs to be set")
	}

	if !user.Succ {
		err = errors.New(user.Error)
	}
	return &user.Result, err
}

func Update(params *tamber.UserParams) (*tamber.User, error) {
	return getClient().Update(params)
}

func (c Client) Update(params *tamber.UserParams) (*tamber.User, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	user := &tamber.UserResponse{}
	var err error

	if len(params.Id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "update", body, user)
	} else {
		err = errors.New("Invalid user params: id needs to be set")
	}

	if !user.Succ {
		err = errors.New(user.Error)
	}
	return &user.Result, err
}

func Retrieve(params *tamber.UserParams) (*tamber.User, error) {
	return getClient().Retrieve(params)
}

func (c Client) Retrieve(params *tamber.UserParams) (*tamber.User, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	user := &tamber.UserResponse{}
	var err error

	if len(params.Id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "retrieve", body, user)
	} else {
		err = errors.New("Invalid user params: id needs to be set")
	}

	if !user.Succ {
		err = errors.New(user.Error)
	}
	return &user.Result, err
}

func getClient() Client {
	return Client{tamber.GetDefaultSessionConfig(), tamber.DefaultProjectKey, tamber.DefaultEngineKey}
}
