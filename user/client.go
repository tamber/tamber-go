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

func Create(params *tamber.UserParams) (*tamber.User, *tamber.ResponseInfo, error) {
	return getClient().Create(params)
}

func (c Client) Create(params *tamber.UserParams) (*tamber.User, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	user := &tamber.UserResponse{}
	var err error

	if len(params.Id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "create", body, user)
	} else {
		err = errors.New("Invalid user params: id needs to be set")
	}

	if err == nil && !user.Succ {
		err = errors.New(user.Error)
	}
	return &user.Result, &user.ResponseInfo, err
}

func Update(params *tamber.UserParams) (*tamber.User, *tamber.ResponseInfo, error) {
	return getClient().Update(params)
}

func (c Client) Update(params *tamber.UserParams) (*tamber.User, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	user := &tamber.UserResponse{}
	var err error

	if len(params.Id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "update", body, user)
	} else {
		err = errors.New("Invalid user params: id needs to be set")
	}

	if err == nil && !user.Succ {
		err = errors.New(user.Error)
	}
	return &user.Result, &user.ResponseInfo, err
}

func Retrieve(params *tamber.UserParams) (*tamber.User, *tamber.ResponseInfo, error) {
	return getClient().Retrieve(params)
}

func (c Client) Retrieve(params *tamber.UserParams) (*tamber.User, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	user := &tamber.UserResponse{}
	var err error

	if len(params.Id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "retrieve", body, user)
	} else {
		err = errors.New("Invalid user params: id needs to be set")
	}

	if err == nil && !user.Succ {
		err = errors.New(user.Error)
	}
	return &user.Result, &user.ResponseInfo, err
}

func Search(params *tamber.UserSearchParams) (*tamber.Users, *tamber.ResponseInfo, error) {
	return getClient().Search(params)
}

func (c Client) Search(params *tamber.UserSearchParams) (*tamber.Users, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	result := &tamber.UserSearchResponse{}
	var err error

	if params.Filter != nil {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "retrieve", body, result)
	} else {
		err = errors.New("Invalid user params: id needs to be set")
	}

	if err == nil && !result.Succ {
		err = errors.New(result.Error)
	}
	return &result.Result, &result.ResponseInfo, err
}

func Merge(params *tamber.UserMergeParams) (*tamber.User, *tamber.ResponseInfo, error) {
	return getClient().Merge(params)
}

func (c Client) Merge(params *tamber.UserMergeParams) (*tamber.User, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	user := &tamber.UserResponse{}
	var err error

	if len(params.From) > 0 && len(params.To) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "retrieve", body, user)
	} else {
		err = errors.New("Invalid merge params: `from` and `to` fields need to be set")
	}

	if err == nil && !user.Succ {
		err = errors.New(user.Error)
	}
	return &user.Result, &user.ResponseInfo, err
}

func getClient() Client {
	return Client{tamber.GetDefaultSessionConfig(), tamber.DefaultProjectKey, tamber.DefaultEngineKey}
}
