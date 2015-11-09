package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type ItemParams struct {
	Id         string
	Properties map[string]interface{}
	Tags       []string
	Created    int64
	// GetSimilar DiscoverParams //Coming soon, not yet supported
}

type Item struct {
	Id         string
	Properties map[string]interface{}
	Tags       []string
	Created    int64
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

// Custom unmarshaling is needed because the result
// may be an id or the full struct.
func (i *Item) UnmarshalJSON(data []byte) error {
	var item Item
	err := json.Unmarshal(data, &item)
	if err == nil {
		*i = item
	} else {
		// the id is surrounded by "\" characters, so strip them
		i.Id = string(data[1 : len(data)-1])
	}

	return nil
}
