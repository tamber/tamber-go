package property

import (
	. "github.com/tamber/tamber-go"
	"net/url"
	"strconv"
)

var object = "property"

func Create(params *PropertyParams) (*Property, error) {
	return getEngine().Create(params)
}

func (e Engine) Create(params *PropertyParams) (*Property, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	property := &Property{}
	var err error

	if len(params.Name) > 0 && len(params.Type) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "create", body, property)
	} else {
		err = errors.New("Invalid property params: name and type need to be set")
	}

	return property, err
}

func Retrieve(params *PropertyParams) (*Property, error) {
	return getEngine().Create(params)
}

func (e Engine) Retrieve(params *PropertyParams) (*Property, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	property := &Property{}
	var err error

	if len(params.Name) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "retrieve", body, property)
	} else {
		err = errors.New("Invalid property params: name needs to be set")
	}

	return property, err
}

func Remove(params *PropertyParams) (*Property, error) {
	return getEngine().Create(params)
}

func (e Engine) Remove(params *PropertyParams) (*Property, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	property := &Property{}
	var err error

	if len(params.Name) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "remove", body, property)
	} else {
		err = errors.New("Invalid property params: name needs to be set")
	}

	return property, err
}

func getEngine() Engine {
	return Engine{GetDefaultSessionConfig(), DefaultKey}
}
