package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Event struct {
	User     string  `json:"user"`
	Item     string  `json:"item"`
	Behavior string  `json:"behavior"`
	Value    float64 `json:"value"`
	Created  int64   `json:"created"`
	Object   string  `json:"object"`
}

type EventParams struct {
	User, Item, Behavior        string // required
	Value                       *float64
	Created                     *int64
	GetRecs                     *DiscoverParams
	CreatedSince, CreatedBefore *int64 //Only used by Retrieve method
	Number                      int    //Only used by Retrieve method | default:200 | max:500
}

type EventResult struct {
	Events []Event      `json:"events"`
	Recs   *[]Discovery `json:"recommended,omitempty"`
}

type EventResponse struct {
	Succ   bool        `json:"success"`
	Result EventResult `json:"result"`
	Error  string      `json:"error"`
	Time   float64     `json:"time"`
}

func (params *EventParams) AppendToBody(v *url.Values) {
	if len(params.User) > 0 {
		v.Add("user", params.User)
	}
	if len(params.Item) > 0 {
		v.Add("item", params.Item)
	}
	if len(params.Behavior) > 0 {
		v.Add("behavior", params.Behavior)
	}

	if params.Value != nil {
		v.Add("value", strconv.FormatFloat(*(params.Value), 'f', -1, 64))
	}
	if params.Created != nil {
		v.Add("created", strconv.FormatInt(*(params.Created), 10))
	}
	getRecs, _ := json.Marshal(params.GetRecs)
	v.Add("get_recs", string(getRecs))

	//For Retrieve Method Only
	if params.CreatedSince != nil {
		v.Add("created_since", strconv.FormatInt(*(params.CreatedSince), 10))
	}
	if params.CreatedBefore != nil {
		v.Add("created_before", strconv.FormatInt(*(params.CreatedBefore), 10))
	}
	if params.Number != 0 {
		v.Add("number", strconv.Itoa(params.Number))
	}
}

//Batch
type EventBatchParams struct {
	Events []Event `json:"events"`
}

type BatchResult struct {
	Object         string `json:"object"`
	NumBatchEvents int    `json:"num_batch_events"` //events in the batch that have been tracked
	NumBatchUsers  int    `json:"num_batch_users"`  //total user count in the batch
	NumBatchItems  int    `json:"num_batch_items"`  //total item count in the batch
	NumUsersAdded  int    `json:"num_users_added"`  //users added from the batch
	NumItemsAdded  int    `json:"num_items_added"`  //items added from the batch
}

type BatchResponse struct {
	Succ   bool        `json:"success"`
	Result BatchResult `json:"result"`
	Error  string      `json:"error"`
	Time   float64     `json:"time"`
}

func (params *EventBatchParams) AppendToBody(v *url.Values) {
	events, _ := json.Marshal(params.Events)
	v.Add("events", string(events))
}
