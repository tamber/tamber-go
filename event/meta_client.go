package event

import (
	"errors"
	tamber "github.com/tamber/tamber-go"
	"net/url"
)

type MetaClient struct {
	S          *tamber.SessionConfig
	ProjectKey string
	EngineKey  string
}

var Meta MetaClient

func (c MetaClient) Like(params *tamber.EventMetaParams) (*tamber.MetaPreference, *tamber.ResponseInfo, error) {
	if c.S == nil {
		return getClient().Meta.Like(params)
	}
	body := &url.Values{}
	params.AppendToBody(body)

	resp := &tamber.EventMetaResponse{}
	err := c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "meta/like", body, resp)

	if err == nil && !resp.Succ {
		err = errors.New(resp.Error)
	}
	return &resp.Result, &resp.ResponseInfo, err
}

func (c MetaClient) Unlike(params *tamber.EventMetaParams) (*tamber.MetaPreference, *tamber.ResponseInfo, error) {
	if c.S == nil {
		return getClient().Meta.Like(params)
	}
	body := &url.Values{}
	params.AppendToBody(body)

	resp := &tamber.EventMetaResponse{}
	err := c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "meta/unlike", body, resp)

	if err == nil && !resp.Succ {
		err = errors.New(resp.Error)
	}
	return &resp.Result, &resp.ResponseInfo, err
}
