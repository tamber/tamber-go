package discover

import (
	"errors"
	tamber "github.com/tamber/tamber-go"
	"net/url"
)

type Engine struct {
	S   *tamber.SessionConfig
	Key string
}

var object = "discover"

func Recommended(params *tamber.DiscoverParams) (*tamber.Discoveries, error) {
	return getEngine().Recommended(params)
}

func (e Engine) Recommended(params *tamber.DiscoverParams) (*tamber.Discoveries, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	discoveries := &tamber.DiscoverResponse{}
	var err error

	if len(params.User) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "recommended", body, discoveries)
	} else {
		err = errors.New("Invalid discover params: user needs to be set")
	}

	if !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, err
}

func Similar(params *tamber.DiscoverParams) (*tamber.Discoveries, error) {
	return getEngine().Similar(params)
}

func (e Engine) Similar(params *tamber.DiscoverParams) (*tamber.Discoveries, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	discoveries := &tamber.DiscoverResponse{}
	var err error

	if len(params.Item) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "similar", body, discoveries)
	} else {
		err = errors.New("Invalid discover params: item needs to be set")
	}

	if !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, err
}

func RecommendedSimilar(params *tamber.DiscoverParams) (*tamber.Discoveries, error) {
	return getEngine().RecommendedSimilar(params)
}

func (e Engine) RecommendedSimilar(params *tamber.DiscoverParams) (*tamber.Discoveries, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	discoveries := &tamber.DiscoverResponse{}
	var err error

	if len(params.User) > 0 && len(params.Item) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "recommendedSimilar", body, discoveries)
	} else {
		err = errors.New("Invalid discover params: user and item need to be set")
	}

	if !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, err
}

func Popular(params *tamber.DiscoverParams) (*tamber.Discoveries, error) {
	return getEngine().Popular(params)
}

func (e Engine) Popular(params *tamber.DiscoverParams) (*tamber.Discoveries, error) {
	body := &url.Values{}
	params.AppendToBody(body)

	discoveries := &tamber.DiscoverResponse{}
	err := e.S.Call("POST", "", e.Key, object, "popular", body, discoveries)

	if !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, err
}

func Hot(params *tamber.DiscoverParams) (*tamber.Discoveries, error) {
	return getEngine().Hot(params)
}

func (e Engine) Hot(params *tamber.DiscoverParams) (*tamber.Discoveries, error) {
	body := &url.Values{}
	params.AppendToBody(body)

	discoveries := &tamber.DiscoverResponse{}
	err := e.S.Call("POST", "", e.Key, object, "hot", body, discoveries)

	if !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, err
}

func getEngine() Engine {
	return Engine{tamber.GetDefaultSessionConfig(), tamber.DefaultKey}
}
