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
	item := &tamber.Item{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "create", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}
	return item, err
}

func AddProperties(params *tamber.ItemParams) (*tamber.Item, error) {
	return getEngine().AddProperties(params)
}

func (e Engine) AddProperties(params *tamber.ItemParams) (*tamber.Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.Item{}
	var err error

	if len(params.Id) > 0 && len(params.Properties) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "addProperties", body, item)
	} else {
		err = errors.New("Invalid item params: id and properties need to be set")
	}
	return item, err
}

func RemoveProperties(params *tamber.ItemParams) (*tamber.Item, error) {
	return getEngine().RemoveProperties(params)
}

func (e Engine) RemoveProperties(params *tamber.ItemParams) (*tamber.Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.Item{}
	var err error

	if len(params.Id) > 0 && len(params.Properties) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "removeProperties", body, item)
	} else {
		err = errors.New("Invalid item params: id and properties need to be set")
	}
	return item, err
}

func AddTags(params *tamber.ItemParams) (*tamber.Item, error) {
	return getEngine().AddTags(params)
}
func (e Engine) AddTags(params *tamber.ItemParams) (*tamber.Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.Item{}
	var err error

	if len(params.Id) > 0 && len(params.Tags) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "addTags", body, item)
	} else {
		err = errors.New("Invalid item params: id and tags need to be set")
	}
	return item, err
}

func RemoveTags(params *tamber.ItemParams) (*tamber.Item, error) {
	return getEngine().RemoveTags(params)
}

func (e Engine) RemoveTags(params *tamber.ItemParams) (*tamber.Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.Item{}
	var err error

	if len(params.Id) > 0 && len(params.Tags) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "removeTags", body, item)
	} else {
		err = errors.New("Invalid item params: id and tags need to be set")
	}
	return item, err
}

func Retrieve(params *tamber.ItemParams) (*tamber.Item, error) {
	return getEngine().Retrieve(params)
}

func (e Engine) Retrieve(params *tamber.ItemParams) (*tamber.Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.Item{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "retrieve", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}

	return item, err
}

func Remove(params *tamber.ItemParams) (*tamber.Item, error) {
	return getEngine().Remove(params)
}

func (e Engine) Remove(params *tamber.ItemParams) (*tamber.Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.Item{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "remove", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}
	return item, err
}

func getEngine() Engine {
	return Engine{tamber.GetDefaultSessionConfig(), tamber.DefaultKey}
}
