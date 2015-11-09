package item

import (
	. "github.com/tamber/tamber-go"
	"net/url"
	"strconv"
)

var object = "item"

func Create(params *ItemParams) (*Item, error) {
	return getEngine().Create(params)
}

func (e Engine) Create(params *ItemParams) (*Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &Item{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "create", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}

	return item, err
}

func AddProperties(params *ItemParams) (*Item, error) {
	return getEngine().AddBehaviors(params)
}
func (e Engine) AddProperties(params *ItemParams) (*Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &Item{}
	var err error

	if len(params.Id) > 0 && lan(params.Properties) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "addProperties", body, item)
	} else {
		err = errors.New("Invalid item params: id and properties need to be set")
	}
	return item, nil
}

func RemoveProperties(params *ItemParams) (*Item, error) {
	return getEngine().RemoveBehaviors(params)
}

func (e Engine) RemoveProperties(params *ItemParams) (*Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &Item{}
	var err error

	if len(params.Id) > 0 && lan(params.Properties) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "removeProperties", body, item)
	} else {
		err = errors.New("Invalid item params: id and properties need to be set")
	}
	return item, nil
}

func AddTags(params *ItemParams) (*Item, error) {
	return getEngine().AddBehaviors(params)
}
func (e Engine) AddTags(params *ItemParams) (*Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &Item{}
	var err error

	if len(params.Id) > 0 && lan(params.Tags) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "addTags", body, item)
	} else {
		err = errors.New("Invalid item params: id and tags need to be set")
	}
	return item, nil
}

func RemoveTags(params *ItemParams) (*Item, error) {
	return getEngine().RemoveBehaviors(params)
}

func (e Engine) RemoveTags(params *ItemParams) (*Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &Item{}
	var err error

	if len(params.Id) > 0 && lan(params.Tags) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "removeTags", body, item)
	} else {
		err = errors.New("Invalid item params: id and tags need to be set")
	}
	return item, nil
}

func Retrieve(params *ItemParams) (*Item, error) {
	return getEngine().Retrieve(params)
}

func (e Engine) Retrieve(params *ItemParams) (*Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &Item{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "retrieve", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}

	return item, err
}

func Remove(params *ItemParams) (*Item, error) {
	return getEngine().Remove(params)
}

func (e Engine) Remove(params *ItemParams) (*Item, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &Item{}
	var err error

	if len(params.Id) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "remove", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}
	return item, err
}

func getEngine() Engine {
	return Engine{GetDefaultSessionConfig(), DefaultKey}
}
