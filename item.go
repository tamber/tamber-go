package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type ItemParams struct {
	Id         string                  `json:"id"`
	Properties *map[string]interface{} `json:"properties,omitempty"`
	Tags       *[]string               `json:"tags,omitempty"`
	Created    int64                   `json:"created,omitempty"`
	// GetSimilar DiscoverParams //Coming soon, not yet supported
}

type Item struct {
	Id         string                 `json:"id"`
	Properties map[string]interface{} `json:"properties"`
	Tags       []string               `json:"tags"`
	Created    int64                  `json:"created"`
}

type ItemResponse struct {
	Succ   bool    `json:"success"`
	Result Item    `json:"result"`
	Error  string  `json:"error"`
	Time   float64 `json:"time"`
}

func (params *ItemParams) AppendToBody(v *url.Values) {

	v.Add("id", params.Id)

	props, _ := json.Marshal(params.Properties)
	v.Add("properties", string(props))

	tags, _ := json.Marshal(params.Tags)
	v.Add("tags", string(tags))

	if params.Created > 0 {
		v.Add("created", strconv.FormatInt(params.Created, 10))
	}
}
