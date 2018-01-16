package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type ItemParams struct {
	Id         string                 `json:"id"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Tags       []string               `json:"tags,omitempty"`
	Created    *int64                 `json:"created,omitempty"`
}

type ItemFeatures struct {
	Properties map[string]interface{} `json:"properties"`
	Tags       []string               `json:"tags"`
}

type ItemUpdates struct {
	Add    ItemFeatures `json:"add"`
	Remove ItemFeatures `json:"remove"`
}

type ItemUpdateParams struct {
	Id       string      `json:"id"`
	Updates  ItemUpdates `json:"updates"`
	NoCreate bool        `json:"no_create,omitempty"`
	Created  *int64      `json:"created,omitempty"`
	Hidden   *bool       `json:"hidden,omitempty"`
}

type Item struct {
	Id         string                 `json:"id"`
	Properties map[string]interface{} `json:"properties"`
	Tags       []string               `json:"tags"`
	Created    int64                  `json:"created"`
	Hidden     bool                   `json:"hidden"`
	// Trends data only returned if the engine key is set, or the project has a default engine
	Popularity float64 `json:"popularity"`
	Hotness    float64 `json:"hotness"`
}

type ItemResponse struct {
	Succ   bool    `json:"success"`
	Result Item    `json:"result"`
	Error  string  `json:"error"`
	Time   float64 `json:"time"`
	ResponseInfo
}

func (r *ItemResponse) SetInfo(info ResponseInfo) {
	r.ResponseInfo = info
}

func (params *ItemParams) AppendToBody(v *url.Values) {

	v.Add("id", params.Id)

	props, _ := json.Marshal(params.Properties)
	if props != nil {
		v.Add("properties", string(props))
	}
	tags, _ := json.Marshal(params.Tags)
	if tags != nil {
		v.Add("tags", string(tags))
	}

	if params.Created != nil {
		v.Add("created", strconv.FormatInt(*params.Created, 10))
	}
}

func (params *ItemUpdateParams) AppendToBody(v *url.Values) {

	v.Add("id", params.Id)

	updates, _ := json.Marshal(params.Updates)
	if updates != nil {
		v.Add("updates", string(updates))
	}

	if params.Created != nil {
		v.Add("created", strconv.FormatInt(*params.Created, 10))
	}
	if params.Hidden != nil {
		v.Add("hidden", strconv.FormatBool(*params.Hidden))
	}

	v.Add("no_create", strconv.FormatBool(params.NoCreate))
}
