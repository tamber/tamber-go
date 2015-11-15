package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type DiscoverParams struct {
	//When setting ActorParams.GetRecs, Actor is not needed and wll be ignored if set.
	Actor  string                  `json:"actor"`
	Item   string                  `json:"item"`
	Number int                     `json:"number"`
	Page   int                     `json:"page"`
	Filter *map[string]interface{} `json:"filter,omitempty"`
}

type Discovery struct {
	Id         string  `json:"id"`
	Score      float64 `json:"score"`
	Popularity float64 `json:"popularity"`
	Hotness    float64 `json:"hotness"`
	Created    int64   `json:"created"`
}

type Discoveries []Discovery

type DiscoverResponse struct {
	Succ   bool        `json:"success"`
	Result Discoveries `json:"result"`
	Error  string      `json:"error"`
	Time   float64     `json:"time"`
}

func (params *DiscoverParams) AppendToBody(v *url.Values) {
	if len(params.Actor) > 0 {
		v.Add("actor", params.Actor)
	}
	if len(params.Item) > 0 {
		v.Add("item", params.Item)
	}
	if params.Number > 0 {
		v.Add("number", strconv.Itoa(params.Number))
	}
	if params.Page > 0 {
		v.Add("page", strconv.Itoa(params.Number))
	}
	filter, _ := json.Marshal(params.Filter)
	v.Add("filter", string(filter))
}
