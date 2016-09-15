package item

import (
	"errors"
	tamber "github.com/tamber/tamber-go"
	"net/url"
)

type Engine struct {
	S   *tamber.SessionConfig
	Key string
}

var object = "item"

func Create(params *tamber.ItemParams) (*tamber.Item, error) {
	return getEngine().Create(params)
}

func (e Engine) Create(params *tamber.ItemParams) (*tamber.Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.ItemResponse{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, "", object, "create", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}
	if !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, err
}

func Update(params *tamber.ItemParams) (*tamber.Item, error) {
	return getEngine().Update(params)
}

func (e Engine) Update(params *tamber.ItemParams) (*tamber.Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.ItemResponse{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, "", object, "update", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}

	if !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, err
}

func Retrieve(params *tamber.ItemParams) (*tamber.Item, error) {
	return getEngine().Retrieve(params)
}

func (e Engine) Retrieve(params *tamber.ItemParams) (*tamber.Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.ItemResponse{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, "", object, "retrieve", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}

	if !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, err
}

func Remove(params *tamber.ItemParams) (*tamber.Item, error) {
	return getEngine().Remove(params)
}

func (e Engine) Remove(params *tamber.ItemParams) (*tamber.Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.ItemResponse{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, "", object, "remove", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}
	if !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, err
}

func getEngine() Engine {
	return Engine{tamber.GetDefaultSessionConfig(), tamber.DefaultKey}
}
