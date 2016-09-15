package tamber

import (
	"encoding/json"
	"net/url"
)

var (
	EventsDatasetName = "behaviors_dataset"
	ItemsDatasetName  = "items_dataset"
)

// Inputs
type UploadParams struct {
	Filepath string
	Type     string
}

type CreateEngineParams struct {
	Name            string
	EventsDatasetId string
	ItemsDatasetId  string
	Behaviors       map[string]Behavior
}

// Types
type Dataset struct {
	Id        string                 `json:"id"`
	Name      string                 `json:"name"`
	Type      string                 `json:"type"`     // behaviors_dataset, items_dataset
	FileType  string                 `json:"filetype"` // json or csv
	FileSize  int64                  `json:"filesize"`
	Behaviors []string               `json:"behaviors"`
	Info      map[string]interface{} `json:"info"`
	Settings  map[string]interface{} `json:"settings"`
	Object    string                 `json:"object"`
}

type AccountInfo struct {
	Id       string             `json:"id"`
	Username string             `json:"username"`
	Engines  map[string]Engine  `json:"clusters"` // key = engine.Id
	Datasets map[string]Dataset `json:"datasets"` // key = dataset.Id
}

// Responses
type AccountResponse struct {
	Succ   bool        `json:"success"`
	Result AccountInfo `json:"result"`
	Error  string      `json:"error"`
	Time   float64     `json:"time"`
}

type UploadResponse struct {
	Succ   bool    `json:"success"`
	Result Dataset `json:"result"`
	Error  string  `json:"error"`
	Time   float64 `json:"time"`
}

type CreateEngineResponse struct {
	Succ   bool    `json:"success"`
	Result Engine  `json:"result"`
	Error  string  `json:"error"`
	Time   float64 `json:"time"`
}

func (params *CreateEngineParams) AppendToBody(v *url.Values) {
	if len(params.Name) > 0 {
		v.Add("name", params.Name)
	}
	if len(params.EventsDatasetId) > 0 {
		v.Add("events_dataset_id", params.EventsDatasetId)
	}
	if len(params.ItemsDatasetId) > 0 {
		v.Add("items_dataset_id", params.ItemsDatasetId)
	}
	behaviors, _ := json.Marshal(params.Behaviors)
	v.Add("behaviors", string(behaviors))
}
