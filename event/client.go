package event

import (
	"errors"
	tamber "github.com/tamber/tamber-go"
	"net/url"
)

type Engine struct {
	S   *tamber.SessionConfig
	Key string
}

var object = "event"

func Track(params *tamber.EventParams) (*tamber.EventResult, error) {
	return getEngine().Track(params)
}

func (e Engine) Track(params *tamber.EventParams) (*tamber.EventResult, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	event := &tamber.EventResponse{}
	var err error

	if len(params.User) > 0 && len(params.Item) > 0 && len(params.Behavior) > 0 {
		err = e.S.Call("POST", "", e.Key, "", object, "track", body, event)
	} else {
		err = errors.New("Invalid event params: user, item, and behavior need to be set")
	}

	if !event.Succ {
		err = errors.New(event.Error)
	}
	return &event.Result, err
}

func Retrieve(params *tamber.EventParams) (*tamber.EventResult, error) {
	return getEngine().Retrieve(params)
}

func (e Engine) Retrieve(params *tamber.EventParams) (*tamber.EventResult, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	event := &tamber.EventResponse{}
	var err error

	err = e.S.Call("POST", "", e.Key, "", object, "retrieve", body, event)

	if !event.Succ {
		err = errors.New(event.Error)
	}
	return &event.Result, err
}

func Batch(params *tamber.EventBatchParams) (*tamber.BatchResult, error) {
	return getEngine().Batch(params)
}

func (e Engine) Batch(params *tamber.EventBatchParams) (*tamber.BatchResult, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	event := &tamber.BatchResponse{}
	var err error

	if len(params.Events) > 0 {
		err = e.S.Call("POST", "", e.Key, "", object, "batch", body, event)
	} else {
		err = errors.New("Invalid batch params: events need to be set")
	}

	if !event.Succ {
		err = errors.New(event.Error)
	}
	return &event.Result, err
}

func getEngine() Engine {
	return Engine{tamber.GetDefaultSessionConfig(), tamber.DefaultKey}
}
