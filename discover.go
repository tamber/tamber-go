package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type DiscoverUser interface {
	GetUserParams() *UserParams
}

type DiscoverItem interface {
	GetItemParams() *ItemParams
}

type DiscoverCommonParams struct {
	User          DiscoverUser `json:"user"`
	Number        *int         `json:"number,omitempty"`
	ExcludeItems  []string     `json:"exclude_items,omitempty"`
	GetProperties bool         `json:"get_properties"`
	NoCreate      bool         `json:"no_create"`
}

func (params *DiscoverCommonParams) AppendToBody(v *url.Values) {
	if user, _ := json.Marshal(params.User.GetUserParams()); user != nil {
		v.Add("user", string(user))
	}
	if params.Number != nil {
		v.Add("number", strconv.Itoa(*params.Number))
	}
	if len(params.ExcludeItems) > 0 {
		if exclude_items, _ := json.Marshal(params.ExcludeItems); exclude_items != nil {
			v.Add("exclude_items", string(exclude_items))
		}
	}
	v.Add("get_properties", strconv.FormatBool(params.GetProperties))
	v.Add("no_create", strconv.FormatBool(params.NoCreate))
}

type DiscoverRecommendedParams struct {
	User          DiscoverUser `json:"user"`
	Number        *int         `json:"number,omitempty"`
	ExcludeItems  []string     `json:"exclude_items,omitempty"`
	GetProperties bool         `json:"get_properties"`
	NoCreate      bool         `json:"no_create"`

	Filter       map[string]interface{} `json:"filter,omitempty"`
	Variability  *float64               `json:"variability,omitempty"`
	Continuation bool                   `json:"continuation"`
}

func (p *DiscoverRecommendedParams) GetDiscoverCommonParams() *DiscoverCommonParams {
	return &DiscoverCommonParams{
		User:          p.User,
		Number:        p.Number,
		ExcludeItems:  p.ExcludeItems,
		GetProperties: p.GetProperties,
		NoCreate:      p.NoCreate,
	}
}

func (params *DiscoverRecommendedParams) AppendToBody(v *url.Values) {
	params.GetDiscoverCommonParams().AppendToBody(v)
	if filter, _ := json.Marshal(params.Filter); filter != nil {
		v.Add("filter", string(filter))
	}
	if params.Variability != nil {
		v.Add("variability", strconv.FormatFloat(*params.Variability, 'f', -1, 64))
	}
	v.Add("continuation", strconv.FormatBool(params.Continuation))
}

type DiscoverNextParams struct {
	User          DiscoverUser `json:"user"`
	Number        *int         `json:"number,omitempty"`
	ExcludeItems  []string     `json:"exclude_items,omitempty"`
	GetProperties bool         `json:"get_properties"`

	Item         DiscoverItem           `json:"item"` // ignores empty string
	Filter       map[string]interface{} `json:"filter,omitempty"`
	Variability  *float64               `json:"variability,omitempty"`
	Continuation bool                   `json:"continuation"`
	NoCreate     string                 `json:"no_create"`
}

func (p *DiscoverNextParams) GetDiscoverCommonParams() *DiscoverCommonParams {
	return &DiscoverCommonParams{
		User:          p.User,
		Number:        p.Number,
		ExcludeItems:  p.ExcludeItems,
		GetProperties: p.GetProperties,
	}
}

func (params *DiscoverNextParams) AppendToBody(v *url.Values) {
	params.GetDiscoverCommonParams().AppendToBody(v)
	if item, _ := json.Marshal(params.Item.GetItemParams()); item != nil {
		v.Add("item", string(item))
	}
	if filter, _ := json.Marshal(params.Filter); filter != nil {
		v.Add("filter", string(filter))
	}
	if params.Variability != nil {
		v.Add("variability", strconv.FormatFloat(*params.Variability, 'f', -1, 64))
	}
	v.Add("continuation", strconv.FormatBool(params.Continuation))
	if len(params.NoCreate) > 0 {
		v.Add("no_create", params.NoCreate)
	}
}

type DiscoverPeriodicParams = DiscoverCommonParams

type DiscoverUserTrendParams struct {
	User          DiscoverUser           `json:"user"`
	Number        *int                   `json:"number,omitempty"`
	Filter        map[string]interface{} `json:"filter,omitempty"`
	GetProperties bool                   `json:"get_properties"`
	NoCreate      bool                   `json:"no_create"`
}

func (params *DiscoverUserTrendParams) AppendToBody(v *url.Values) {
	if user, _ := json.Marshal(params.User.GetUserParams()); user != nil {
		v.Add("user", string(user))
	}
	if params.Number != nil {
		v.Add("number", strconv.Itoa(*params.Number))
	}
	if len(params.Filter) > 0 {
		if filter, _ := json.Marshal(params.Filter); filter != nil {
			v.Add("filter", string(filter))
		}
	}
	v.Add("get_properties", strconv.FormatBool(params.GetProperties))
	v.Add("no_create", strconv.FormatBool(params.NoCreate))
}

type DiscoverBasicParams struct {
	User          string                 `json:"user"` // ignores empty string
	Item          string                 `json:"item"` // ignores empty string
	Number        *int                   `json:"number,omitempty"`
	Page          *int                   `json:"page,omitempty"`
	Filter        map[string]interface{} `json:"filter,omitempty"`
	TestEvents    []EventParams          `json:"test_events,omitempty"`
	GetProperties bool                   `json:"get_properties"`
}

func (params *DiscoverBasicParams) AppendToBody(v *url.Values) {
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
	if filter, _ := json.Marshal(params.Filter); filter != nil {
		v.Add("filter", string(filter))
	}
	if test_events, _ := json.Marshal(params.TestEvents); test_events != nil {
		v.Add("test_events", string(test_events))
	}
	v.Add("get_properties", strconv.FormatBool(params.GetProperties))
}

type DiscoverParams struct {
	Number        *int                   `json:"number,omitempty"`
	Filter        map[string]interface{} `json:"filter,omitempty"`
	ExcludeItems  []string               `json:"exclude_items,omitempty"`
	GetProperties bool                   `json:"get_properties"`
	Variability   *float64               `json:"variability,omitempty"`
}

func (params *DiscoverParams) AppendToBody(v *url.Values) {
	if params.Number != nil {
		v.Add("number", strconv.Itoa(*params.Number))
	}
	if params.Variability != nil {
		v.Add("variability", strconv.FormatFloat(*params.Variability, 'f', -1, 64))
	}
	if filter, _ := json.Marshal(params.Filter); filter != nil {
		v.Add("filter", string(filter))
	}
	if len(params.ExcludeItems) > 0 {
		if exclude_items, _ := json.Marshal(params.ExcludeItems); exclude_items != nil {
			v.Add("exclude_items", string(exclude_items))
		}
	}
	v.Add("get_properties", strconv.FormatBool(params.GetProperties))
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
