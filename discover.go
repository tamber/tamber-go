package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type DiscoverParams struct {
	User          string                  `json:"user"`
	Item          string                  `json:"item"`
	Number        int                     `json:"number"`
	Page          int                     `json:"page"`
	Filter        *map[string]interface{} `json:"filter,omitempty"`
	TestEvents    *[]Event                `json:"test_events,omitempty"`
	GetProperties bool                    `json:"get_properties"`
}

type Discovery struct {
	Item        string                 `json:"item"`
	Object      string                 `json:"object"`
	Score       float64                `json:"score"`
	Popularity  float64                `json:"popularity"`
	Hotness     float64                `json:"hotness"`
	Properties  map[string]interface{} `json:"properties"`
	Tags        []string               `json:"tags"`
	ItemCreated int64                  `json:"item_created"`
}

type Discoveries []Discovery

type DiscoverResponse struct {
	Succ   bool        `json:"success"`
	Result Discoveries `json:"result"`
	Error  string      `json:"error"`
	Time   float64     `json:"time"`
}

func (params *DiscoverParams) AppendToBody(v *url.Values) {
	if len(params.User) > 0 {
		v.Add("user", params.User)
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

	test_events, _ := json.Marshal(params.TestEvents)
	v.Add("test_events", string(test_events))
}
