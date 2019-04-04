package discover

import (
	"errors"
	tamber "github.com/tamber/tamber-go"
	"net/url"
)

type UserTrendClient struct {
	S          *tamber.SessionConfig
	ProjectKey string
	EngineKey  string
}

var UserTrend UserTrendClient

func (c UserTrendClient) Popular(params *tamber.DiscoverUserTrendParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	if c.S == nil {
		return getClient().UserTrend.Popular(params)
	}
	body := &url.Values{}
	params.AppendToBody(body)

	discoveries := &tamber.DiscoverResponse{}
	err := c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "user_trend/popular", body, discoveries)

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}

func (c UserTrendClient) Hot(params *tamber.DiscoverUserTrendParams) (*tamber.Discoveries, *tamber.ResponseInfo, error) {
	if c.S == nil {
		return getClient().UserTrend.Hot(params)
	}
	body := &url.Values{}
	params.AppendToBody(body)
	discoveries := &tamber.DiscoverResponse{}
	var err error

	err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "user_trend/hot", body, discoveries)

	if err == nil && !discoveries.Succ {
		err = errors.New(discoveries.Error)
	}
	return &discoveries.Result, &discoveries.ResponseInfo, err
}

