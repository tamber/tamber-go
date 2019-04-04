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

func (params *UserParams) GetUserParams() *UserParams {
	return params
}

type UserListParams struct {
	Number *int                   `json:"number,omitempty"`
	Page   *int                   `json:"page,omitempty"`
	Filter map[string]interface{} `json:"filter,omitempty"`
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

func (user *User) GetUserParams() *UserParams {
	if user == nil {
		return nil
	}
	p := &UserParams{
		Id: user.Id,
		Metadata: user.Metadata,
	}
	if user.Created != 0 {
		p.Created = &user.Created
	}
	return p
}

type UserResponse struct {
	Succ   bool   `json:"success"`
	Result User   `json:"result"`
	Error  string `json:"error"`
	ResponseInfo
}

type Users []User

type UsersResponse struct {
	Succ   bool   `json:"success"`
	Result Users  `json:"result"`
	Error  string `json:"error"`
	ResponseInfo
}

func (r *UserResponse) SetInfo(info ResponseInfo) {
	info.Time = r.Time
	r.ResponseInfo = info
}

func (r *UsersResponse) SetInfo(info ResponseInfo) {
	info.Time = r.Time
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

func (params *UserListParams) AppendToBody(v *url.Values) {

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
}

func (params *UserSearchParams) AppendToBody(v *url.Values) {
	filter, _ := json.Marshal(params.Filter)
	v.Add("filter", string(filter))
}
