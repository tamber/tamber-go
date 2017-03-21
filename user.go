package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type UserParams struct {
	Id       string                  `json:"id"`
	Events   *[]Event                `json:"events,omitempty"`
	GetRecs  *DiscoverParams         `json:"get_recs,omitempty"` //DiscoverParams.User field is not needed and wll be ignored if set.
	Metadata *map[string]interface{} `json:"metadata"`
	Created  int64                   `json:"created,omitempty"` //0 values ignored
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

func (r *UserResponse) SetInfo(info ResponseInfo) {
	r.ResponseInfo = info
}

func (params *UserParams) AppendToBody(v *url.Values) {

	v.Add("id", params.Id)

	if params.Events != nil {
		events, _ := json.Marshal(params.Events)
		v.Add("events", string(events))
	}

	if params.GetRecs != nil {
		getRecs, _ := json.Marshal(params.GetRecs)
		v.Add("get_recs", string(getRecs))
	}

	metadata, _ := json.Marshal(params.Metadata)
	v.Add("metadata", string(metadata))

	if params.Created > 0 {
		v.Add("created", strconv.FormatInt(params.Created, 10))
	}

}
