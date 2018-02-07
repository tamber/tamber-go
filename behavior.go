package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

const (
	ExponentialBehaviorType = "exponential"
	RatingBehaviorType      = "rating"
)

type BehaviorParams struct {
	Name         string
	Desirability float64
	Type         *string
	Params       map[string]interface{}
}

type Behavior struct {
	Name         string                 `json:"name"`
	Type         string                 `json:"type"`
	Desirability float64                `json:"desirability"`
	Params       map[string]interface{} `json:"params"`
	Created      int64                  `json:"created"`
}

type BehaviorResponse struct {
	Succ   bool     `json:"success"`
	Result Behavior `json:"result"`
	Error  string   `json:"error"`
	ResponseInfo
}

func (r *BehaviorResponse) SetInfo(info ResponseInfo) {
	info.Time = r.Time
	r.ResponseInfo = info
}

func (params *BehaviorParams) AppendToBody(v *url.Values) {
	v.Add("name", params.Name)
	if params.Type != nil {
		v.Add("type", *params.Type)
	}

	v.Add("desirability", strconv.FormatFloat(params.Desirability, 'f', -1, 64))

	if params.Params != nil {
		bParams, _ := json.Marshal(params.Params)
		v.Add("params", string(bParams))
	}
}
