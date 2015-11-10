package tamber

import (
	"net/url"
)

type PropertyParams struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Property struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Created int64  `json:"created"`
}

type PropertyResponse struct {
	Succ   bool     `json:"success"`
	Result Property `json:"result"`
	Error  string   `json:"error"`
	Time   float64  `json:"time"`
}

func (params *PropertyParams) AppendToBody(v *url.Values) {
	v.Add("name", params.Name)
	if len(params.Type) > 0 {
		v.Add("type", params.Type)
	}
}
