package event

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

var object = "event"

func Track(params *tamber.EventParams) (*tamber.EventResult, *tamber.ResponseInfo, error) {
	return getClient().Track(params)
}

func (c Client) Track(params *tamber.EventParams) (*tamber.EventResult, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	event := &tamber.EventResponse{}
	var err error

	if len(params.User) > 0 && len(params.Item) > 0 && len(params.Behavior) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "track", body, event)
	} else {
		err = errors.New("Invalid event params: user, item, and behavior need to be set")
	}

	if err == nil && !event.Succ {
		err = errors.New(event.Error)
	}
	return &event.Result, &event.ResponseInfo, err
}

func Retrieve(params *tamber.EventParams) (*tamber.EventResult, *tamber.ResponseInfo, error) {
	return getClient().Retrieve(params)
}

func (c Client) Retrieve(params *tamber.EventParams) (*tamber.EventResult, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	event := &tamber.EventResponse{}
	var err error

	err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "retrieve", body, event)

	if err == nil && !event.Succ {
		err = errors.New(event.Error)
	}
	return &event.Result, &event.ResponseInfo, err
}

func Batch(params *tamber.EventBatchParams) (*tamber.BatchResult, *tamber.ResponseInfo, error) {
	return getClient().Batch(params)
}

func (c Client) Batch(params *tamber.EventBatchParams) (*tamber.BatchResult, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	event := &tamber.BatchResponse{}
	var err error

	if len(params.Events) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "batch", body, event)
	} else {
		err = errors.New("Invalid batch params: events need to be set")
	}

	if err == nil && !event.Succ {
		err = errors.New(event.Error)
	}
	return &event.Result, &event.ResponseInfo, err
}

func getClient() Client {
	return Client{tamber.GetDefaultSessionConfig(), tamber.DefaultProjectKey, tamber.DefaultEngineKey}
}
