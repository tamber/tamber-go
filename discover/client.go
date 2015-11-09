package discover

import (
	"net/url"
	"strconv"
	. "tamber-go-master"
)

var object = "discover"

func Recommended(params *DiscoverParams) (*Discoveries, error) {
	return getEngine().Recommended(params)
}

func (e Engine) Recommended(params *DiscoverParams) (*Discoveries, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	discoveries := &Discoveries{}
	var err error

	if len(params.Actor) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "getRecs", body, discoveries)
	} else {
		err = errors.New("Invalid discover params: actor needs to be set")
	}

	return discoveries, err
}

func Similar(params *DiscoverParams) (*Discoveries, error) {
	return getEngine().Similar(params)
}

func (e Engine) Similar(params *DiscoverParams) (*Discoveries, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	discoveries := &Discoveries{}
	var err error

	if len(params.Item) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "getSimilar", body, discoveries)
	} else {
		err = errors.New("Invalid discover params: item needs to be set")
	}

	return discoveries, err
}

func RecommendedSimilar(params *DiscoverParams) (*Discoveries, error) {
	return getEngine().RecommendedSimilar(params)
}

func (e Engine) RecommendedSimilar(params *DiscoverParams) (*Discoveries, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	discoveries := &Discoveries{}
	var err error

	if len(params.Actor) > 0 && len(params.Item) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "getRecommendedSimilar", body, discoveries)
	} else {
		err = errors.New("Invalid discover params: actor and item need to be set")
	}

	return discoveries, err
}

func Popular(params *DiscoverParams) (*Discoveries, error) {
	return getEngine().Popular(params)
}

func (e Engine) Popular(params *DiscoverParams) (*Discoveries, error) {
	body := &url.Values{}
	params.AppendToBody(body)

	discoveries := &Discoveries{}
	err := e.S.Call("POST", "", e.Key, object, "getPopular", body, discoveries)

	return discoveries, err
}

func Hot(params *DiscoverParams) (*Discoveries, error) {
	return getEngine().Hot(params)
}

func (e Engine) Hot(params *DiscoverParams) (*Discoveries, error) {
	body := &url.Values{}
	params.AppendToBody(body)

	discoveries := &Discoveries{}
	err := e.S.Call("POST", "", e.Key, object, "getHot", body, discoveries)

	return discoveries, err
}

func getEngine() Engine {
	return Engine{GetDefaultSessionConfig(), DefaultKey}
}
