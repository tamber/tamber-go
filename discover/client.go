package discover

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

var object = "discover"

func Next(params *tamber.DiscoverNextParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	return getClient().Next(params)
}

func (c Client) Next(params *tamber.DiscoverNextParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)

	discoveries := &tamber.DiscoverResponse{}
	err := c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "next", body, discoveries)

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}

func Recommended(params *tamber.DiscoverParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	return getClient().Recommended(params)
}

func (c Client) Recommended(params *tamber.DiscoverParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	discoveries := &tamber.DiscoverResponse{}
	var err error

	if len(params.User) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "recommended", body, discoveries)
	} else {
		err = errors.New("Invalid discover params: user needs to be set")
	}

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}

func Similar(params *tamber.DiscoverParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	return getClient().Similar(params)
}

func (c Client) Similar(params *tamber.DiscoverParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	discoveries := &tamber.DiscoverResponse{}
	var err error

	if len(params.Item) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "similar", body, discoveries)
	} else {
		err = errors.New("Invalid discover params: item needs to be set")
	}

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}

func RecommendedSimilar(params *tamber.DiscoverParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	return getClient().RecommendedSimilar(params)
}

func (c Client) RecommendedSimilar(params *tamber.DiscoverParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	discoveries := &tamber.DiscoverResponse{}
	var err error

	if len(params.User) > 0 && len(params.Item) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "recommended_similar", body, discoveries)
	} else {
		err = errors.New("Invalid discover params: user and item need to be set")
	}

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}

func Popular(params *tamber.DiscoverParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	return getClient().Popular(params)
}

func (c Client) Popular(params *tamber.DiscoverParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)

	discoveries := &tamber.DiscoverResponse{}
	err := c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "popular", body, discoveries)

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}

func Hot(params *tamber.DiscoverParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	return getClient().Hot(params)
}

func (c Client) Hot(params *tamber.DiscoverParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)

	discoveries := &tamber.DiscoverResponse{}
	err := c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "hot", body, discoveries)

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}

func getClient() Client {
	return Client{tamber.GetDefaultSessionConfig(), tamber.DefaultProjectKey, tamber.DefaultEngineKey}
}
