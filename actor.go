package tamber

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type ActorParams struct {
	Id        string           `json:"id"`
	Behaviors *[]ActorBehavior `json:"behaviors,omitempty"`
	GetRecs   *DiscoverParams  `json:"getRecs,omitempty"` //DiscoverParams.Actor field is not needed and wll be ignored if set.
	Created   int64            `json:"created,omitempty"` //cannot be set to 0
}

type ActorBehavior struct {
	Behavior string  `json:"behavior"`
	Item     string  `json:"item"`
	Value    float64 `json:"value"`
	Created  int64   `json:"created"`
}

type Actor struct {
	Id        string           `json:"id"`
	Behaviors *[]ActorBehavior `json:"behaviors"`
	Recs      Discoveries      `json:"recs"`
	Created   int64            `json:"created"`
}

type ActorResponse struct {
	Succ   bool    `json:"success"`
	Result Actor   `json:"result"`
	Error  string  `json:"error"`
	Time   float64 `json:"time"`
}

func (params *ActorParams) AppendToBody(v *url.Values) {

	v.Add("id", params.Id)

	behaviors, _ := json.Marshal(params.Behaviors)
	v.Add("behaviors", string(behaviors))

	getRecs, _ := json.Marshal(params.GetRecs)
	v.Add("getRecs", string(getRecs))
	fmt.Printf("%v", v)

	if params.Created > 0 {
		v.Add("created", strconv.FormatInt(params.Created, 10))
	}
}
