package tamber

import (
	"encoding/json"
	"net/url"
)

type PropertyParams struct {
	Name, Type string
}

type Property struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Created int64  `json:"created"`
}

func (params *PropertyParams) AppendToBody(v *url.Values) {
	v.Add("name", params.Name)
	if len(params.Type) > 0 {
		v.Add("type", params.Type)
	}
}

// Custom unmarshaling is needed because the result
// may be a name or the full struct.
func (p *Property) UnmarshalJSON(data []byte) error {
	var property Property
	err := json.Unmarshal(data, &property)
	if err == nil {
		*p = property
	} else {
		// the name is surrounded by "\" characters, so strip them
		p.Name = string(data[1 : len(data)-1])
	}

	return nil
}
