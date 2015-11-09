package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type DiscoverParams struct {
	Actor, Item  string //When setting ActorParams.GetRecs, Actor is not needed and wll be ignored if set.
	Number, Page int
	Filter       map[string]string
}

type Discovery struct {
	Id         string  `json:"id"`
	Score      float64 `json:"score"`
	Popularity float64 `json:"popularity"`
	Hotness    float64 `json:"hotness"`
	Created    int64   `json:"created"`
}

type Discoveries []Discovery

// type DiscoverReturn struct {
// 	Succ   bool
// 	Result []Discovery
// 	Time   float64
// }

func (params *DiscoverParams) AppendToBody(v *url.Values) {
	if len(params.Actor) > 0 {
		v.Add("actor", params.Actor)
	}
	if len(params.Item) > 0 {
		v.Add("item", params.Item)
	}
	if params.Number > 0 {
		v.Add("number", strconv.Itoa(params.Number))
	}
	if params.Page > 0 {
		v.Add("page", strconv.Itoa(params.Number))
	}
	filter, _ := json.Marshal(params.Filter)
	v.Add("tags", string(filter))
}

// Custom unmarshaling is needed because the result
// may be an id or the full struct.
func (d *Discovery) UnmarshalJSON(data []byte) error {
	var discovery Discovery
	err := json.Unmarshal(data, &discovery)
	if err == nil {
		*d = discovery
	} else {
		// the id is surrounded by "\" characters, so strip them
		d.Id = string(data[1 : len(data)-1])
	}

	return nil
}
