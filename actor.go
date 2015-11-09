package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type ActorParams struct {
	Id        string
	Behaviors *[]ActorBehavior
	GetRecs   DiscoverParams //DiscoverParams.Actor field is not needed and wll be ignored if set.
	Created   int64          //cannot be set to 0
}

type ActorBehavior struct {
	Behavior, Item string
	Value          float64
	Created        int64
}

type Actor struct {
	Id        string           `json:"id"`
	Behaviors *[]ActorBehavior `json:"behaviors"`
	Recs      []Discovery      `json:"recs"`
	Created   int64            `json:"created"`
}

type ActorReturn struct {
	Succ   bool
	Result Actor
	Error  string
	Time   float64
}

func (params *ActorParams) AppendToBody(v *url.Values) {

	v.Add("id", params.Id)

	behaviors, _ := json.Marshal(params.Behaviors)
	v.Add("behaviors", string(behaviors))

	getRecs, _ := json.Marshal(params.GetRecs)
	v.Add("getRecs", string(getRecs))

	if params.Created > 0 {
		v.Add("created", strconv.FormatInt(params.Created, 10))
	}
}

// Custom unmarshaling is needed because the result
// may be an id or the full struct.
func (a *Actor) UnmarshalJSON(data []byte) error {
	var actor Actor
	err := json.Unmarshal(data, &actor)
	if err == nil {
		*a = actor
	} else {
		// the id is surrounded by "\" characters, so strip them
		a.Id = string(data[1 : len(data)-1])
	}

	return nil
}
