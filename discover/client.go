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
	Basic      BasicClient
	UserTrend  UserTrendClient
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

func Recommended(params *tamber.DiscoverRecommendedParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	return getClient().Recommended(params)
}

func (c Client) Recommended(params *tamber.DiscoverRecommendedParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)

	discoveries := &tamber.DiscoverResponse{}
	err := c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "recommended", body, discoveries)

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}

func Weekly(params *tamber.DiscoverPeriodicParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	return getClient().Weekly(params)
}

func (c Client) Weekly(params *tamber.DiscoverPeriodicParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)

	discoveries := &tamber.DiscoverResponse{}
	err := c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "weekly", body, discoveries)

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}

func Daily(params *tamber.DiscoverPeriodicParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	return getClient().Daily(params)
}

func (c Client) Daily(params *tamber.DiscoverPeriodicParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)

	discoveries := &tamber.DiscoverResponse{}
	err := c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "daily", body, discoveries)

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}

func Popular(params *tamber.DiscoverBasicParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	return getClient().Popular(params)
}

func (c Client) Popular(params *tamber.DiscoverBasicParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)

	discoveries := &tamber.DiscoverResponse{}
	err := c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "popular", body, discoveries)

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}

func Hot(params *tamber.DiscoverBasicParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	return getClient().Hot(params)
}

func (c Client) Hot(params *tamber.DiscoverBasicParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)

	discoveries := &tamber.DiscoverResponse{}
	err := c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "hot", body, discoveries)

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}

func NewClient(s *tamber.SessionConfig, project, engine string) *Client {
	basic := BasicClient{s, project, engine}
	userTrend := UserTrendClient{s, project, engine}
	return &Client{s, project, engine, basic, userTrend}
}

func getClient() Client {
	return *NewClient(tamber.GetDefaultSessionConfig(), tamber.DefaultProjectKey, tamber.DefaultEngineKey)
}
