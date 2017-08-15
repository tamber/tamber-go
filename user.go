package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type UserParams struct {
	Id       string                 `json:"id"`
	Events   []EventParams          `json:"events,omitempty"`
	GetRecs  *DiscoverParams        `json:"get_recs,omitempty"` //DiscoverParams.User field is not needed and wll be ignored if set.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Created  *int64                 `json:"created,omitempty"`
}

type UserSearchParams struct {
	Filter map[string]interface{} `json:"filter,omitempty"`
}

type UserMergeParams struct {
	From     string `json:"from"`
	To       string `json:"to"`
	NoCreate bool   `json:"no_create,omitempty"`
}

type User struct {
	Id       string                 `json:"id"`
	Object   string                 `json:"object"`
	Events   []Event                `json:"events"`
	Recs     Discoveries            `json:"recommended, omitempty"`
	Metadata map[string]interface{} `json:"metadata"`
	Created  int64                  `json:"created"`
}

type UserResponse struct {
	Succ   bool    `json:"success"`
	Result User    `json:"result"`
	Error  string  `json:"error"`
	Time   float64 `json:"time"`
	ResponseInfo
}

type Users []User

type UserSearchResponse struct {
	Succ   bool    `json:"success"`
	Result Users   `json:"result"`
	Error  string  `json:"error"`
	Time   float64 `json:"time"`
	ResponseInfo
}

func (r *UserResponse) SetInfo(info ResponseInfo) {
	r.ResponseInfo = info
}

func (r *UserSearchResponse) SetInfo(info ResponseInfo) {
	r.ResponseInfo = info
}

func (params *UserParams) AppendToBody(v *url.Values) {

	v.Add("id", params.Id)

	events, _ := json.Marshal(params.Events)
	if events != nil {
		v.Add("events", string(events))
	}

	getRecs, _ := json.Marshal(params.GetRecs)
	if getRecs != nil {
		v.Add("get_recs", string(getRecs))
	}

	metadata, _ := json.Marshal(params.Metadata)
	if metadata != nil {
		v.Add("metadata", string(metadata))
	}

	if params.Created != nil {
		v.Add("created", strconv.FormatInt(*params.Created, 10))
	}
}

func (params *UserMergeParams) AppendToBody(v *url.Values) {
	v.Add("from", params.From)
	v.Add("to", params.To)
	v.Add("no_create", strconv.FormatBool(params.NoCreate))
}

func (params *UserSearchParams) AppendToBody(v *url.Values) {
	filter, _ := json.Marshal(params.Filter)
	v.Add("filter", string(filter))
}
