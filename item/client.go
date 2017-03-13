package item

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

var object = "item"

func Create(params *tamber.ItemParams) (*tamber.Item, error) {
	return getClient().Create(params)
}

func (c Client) Create(params *tamber.ItemParams) (*tamber.Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.ItemResponse{}
	var err error

	if len(params.Id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "create", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}

	if err == nil && !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, err
}

func Update(params *tamber.ItemParams) (*tamber.Item, error) {
	return getClient().Update(params)
}

func (c Client) Update(params *tamber.ItemParams) (*tamber.Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.ItemResponse{}
	var err error

	if len(params.Id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "update", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}

	if err == nil && !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, err
}

func Retrieve(params *tamber.ItemParams) (*tamber.Item, error) {
	return getClient().Retrieve(params)
}

func (c Client) Retrieve(params *tamber.ItemParams) (*tamber.Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.ItemResponse{}
	var err error

	if len(params.Id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "retrieve", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}

	if err == nil && !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, err
}

func Remove(params *tamber.ItemParams) (*tamber.Item, error) {
	return getClient().Remove(params)
}

func (c Client) Remove(params *tamber.ItemParams) (*tamber.Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.ItemResponse{}
	var err error

	if len(params.Id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "remove", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}
	if err == nil && !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, err
}

func getClient() Client {
	return Client{tamber.GetDefaultSessionConfig(), tamber.DefaultProjectKey, tamber.DefaultEngineKey}
}
