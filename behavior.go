package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type BehaviorParams struct {
	Name, Type   string
	Desirability float64
	Params       map[string]interface{}
}

type Behavior struct {
	Name         string                 `json:"name"`
	Type         string                 `json:"type"`
	Desirability float64                `json:"desirability"`
	Params       map[string]interface{} `json:"params"`
	Created      int64                  `json:"created"`
}

func (params *BehaviorParams) AppendToBody(v *url.Values) {
	v.Add("name", params.Name)
	if len(params.Type) > 0 {
		v.Add("type", params.Type)
	}
	if params.Desirability > 0 {
		v.Add("desirability", strconv.FormatFloat(params.Desirability, 'E', -1, 64))
	}
	bParams, _ := json.Marshal(params.Params)
	v.Add("params", string(bParams))
}

// Custom unmarshaling is needed because the result
// may be a name or the full struct.
func (b *Behavior) UnmarshalJSON(data []byte) error {
	var behavior Behavior
	err := json.Unmarshal(data, &behavior)
	if err == nil {
		*b = behavior
	} else {
		// the name is surrounded by "\" characters, so strip them
		b.Name = string(data[1 : len(data)-1])
	}
	return nil
}
