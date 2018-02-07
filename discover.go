package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type DiscoverParams struct {
	User          string                 `json:"user"` // ignores empty string
	Item          string                 `json:"item"` // ignores empty string
	Number        *int                   `json:"number,omitempty"`
	Page          *int                   `json:"page,omitempty"`
	Filter        map[string]interface{} `json:"filter,omitempty"`
	TestEvents    []EventParams          `json:"test_events,omitempty"`
	GetProperties bool                   `json:"get_properties"`
}

type DiscoverNextParams struct {
	User          string                 `json:"user"` // ignores empty string
	Item          string                 `json:"item"` // ignores empty string
	Number        *int                   `json:"number,omitempty"`
	Filter        map[string]interface{} `json:"filter,omitempty"`
	ExcludeItems  []string               `json:"exclude_items,omitempty"`
	Variability   *float64               `json:"variability,omitempty"`
	GetProperties bool                   `json:"get_properties"`
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
	ResponseInfo
}

func (r *DiscoverResponse) SetInfo(info ResponseInfo) {
	info.Time = r.Time
	r.ResponseInfo = info
}

func (params *DiscoverParams) AppendToBody(v *url.Values) {
	if len(params.User) > 0 {
		v.Add("user", params.User)
	}
	if len(params.Item) > 0 {
		v.Add("item", params.Item)
	}
	if params.Number != nil {
		v.Add("number", strconv.Itoa(*params.Number))
	}
	if params.Page != nil {
		v.Add("page", strconv.Itoa(*params.Page))
	}

	filter, _ := json.Marshal(params.Filter)
	if filter != nil {
		v.Add("filter", string(filter))
	}

	test_events, _ := json.Marshal(params.TestEvents)
	if test_events != nil {
		v.Add("test_events", string(test_events))
	}

	v.Add("get_properties", strconv.FormatBool(params.GetProperties))
}

func (params *DiscoverNextParams) AppendToBody(v *url.Values) {
	if len(params.User) > 0 {
		v.Add("user", params.User)
	}
	if len(params.Item) > 0 {
		v.Add("item", params.Item)
	}
	if params.Number != nil {
		v.Add("number", strconv.Itoa(*params.Number))
	}
	if params.Variability != nil {
		v.Add("variability", strconv.FormatFloat(*params.Variability, 'f', -1, 64))
	}
	filter, _ := json.Marshal(params.Filter)
	if filter != nil {
		v.Add("filter", string(filter))
	}

	exclude_items, _ := json.Marshal(params.ExcludeItems)
	if exclude_items != nil {
		v.Add("exclude_items", string(exclude_items))
	}

	v.Add("get_properties", strconv.FormatBool(params.GetProperties))
}
