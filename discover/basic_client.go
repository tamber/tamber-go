package discover

import (
	"errors"
	tamber "github.com/tamber/tamber-go"
	"net/url"
)

type BasicClient struct {
	S          *tamber.SessionConfig
	ProjectKey string
	EngineKey  string
}

var Basic BasicClient

func (c BasicClient) Recommended(params *tamber.DiscoverBasicParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	if c.S == nil {
		return getClient().Basic.Recommended(params)
	}
	body := &url.Values{}
	params.AppendToBody(body)

	discoveries := &tamber.DiscoverResponse{}
	err := c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "basic/recommended", body, discoveries)

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}

func (c BasicClient) Similar(params *tamber.DiscoverBasicParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	if c.S == nil {
		return getClient().Basic.Similar(params)
	}
	body := &url.Values{}
	params.AppendToBody(body)
	discoveries := &tamber.DiscoverResponse{}
	var err error

	if len(params.Item) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "basic/similar", body, discoveries)
	} else {
		err = errors.New("Invalid discover params: item needs to be set")
	}

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}

func (c BasicClient) RecommendedSimilar(params *tamber.DiscoverBasicParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	if c.S == nil {
		return getClient().Basic.RecommendedSimilar(params)
	}
	body := &url.Values{}
	params.AppendToBody(body)
	discoveries := &tamber.DiscoverResponse{}
	var err error

	if len(params.User) > 0 && len(params.Item) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "basic/recommended_similar", body, discoveries)
	} else {
		err = errors.New("Invalid discover params: user and item need to be set")
	}

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}
