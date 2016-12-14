package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

var (
	EventsDatasetName = "behaviors_dataset"
	ItemsDatasetName  = "items_dataset"
)

// Inputs
type UploadParams struct {
	Filepath string `json:"filepath"`
	Type     string `json:"type"`
}

type CreateProjectParentParams struct {
	AccountId    string   `json:"accountid"`
	Name         string   `json:"name"`
	Environments []string `json:"environments"`
}

type CreateProjectParams struct {
	AccountId       string `json:"accountid"`
	Environment     string `json:"environment"`
	ProjectParentId string `json:"parentid"`
}

type CreateEngineParams struct {
	Name        string                 `json:"name"`
	AccountId   string                 `json:"accountid"`
	ProjectId   uint32                 `json:"projectid"`
	Behaviors   map[string]Behavior    `json:"behaviors"`
	ItemsFilter map[string]interface{} `json:"filter"`
}

// Types
type DashboardData struct {
	BehaviorCount int64 `json:"behavior_count"`
	ItemCount     int64 `json:"item_count"`
	UserCount     int64 `json:"user_count"`
}

type Project struct {
	Id              uint32                 `json:"id"`
	Key             string                 `json:"key"`
	Name            string                 `json:"name"`
	Environment     string                 `json:"environment"`
	AccountId       string                 `json:"accountid"`
	ProjectParentId string                 `json:"parentid"`
	ApiVersion      string                 `json:"apiversion"`
	Engines         []string               `json:"engines"`
	Metadata        map[string]interface{} `json:"metadata"`
	State           int                    `json:"state"`
	Behaviors       []string               `json:"behaviors"`
	Dashboard       DashboardData          `json:"dashboard"`
}

type ProjectKey struct {
	Id          uint32 `json:"id"`
	Environment string `json:"environment"`
}

type ProjectParent struct {
	Id         string       `json:"id"`
	Name       string       `json:"name"`
	ApiVersion string       `json:"apiversion"`
	Projects   []ProjectKey `json:"projects"`
}

type Engine struct {
	Key         string
	EngineId    uint32 `bson:"engine_id"`
	Id          string
	ProjectId   uint32
	Name        string
	Status      int
	ApiVersion  string
	Dashboard   DashboardData
	Behaviors   map[string]Behavior    `json:"behaviors"`
	ItemsFilter map[string]interface{} `json:"filter"`
}

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
	Id             string                   `json:"id"`
	Username       string                   `json:"username"`
	ProjectParents map[string]ProjectParent `json:"project_parents"` // key = projectParent.Id
	Projects       map[uint32]Project       `json:"projects"`        // key = project.Id
	Engines        map[string]Engine        `json:"engines"`         // key = engine.Id
	Datasets       map[string]Dataset       `json:"datasets"`        // key = dataset.Id
}

type AuthToken struct {
	Token      string `json:"token"`
	AccountId  string `json:"accountid`
	ExpireTime int64  `json:"expiration_timestamp"`
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

type CreateProjectParentResponse struct {
	Succ   bool          `json:"success"`
	Result ProjectParent `json:"result"`
	Error  string        `json:"error"`
	Time   float64       `json:"time"`
}

type CreateProjectResponse struct {
	Succ   bool    `json:"success"`
	Result Project `json:"result"`
	Error  string  `json:"error"`
	Time   float64 `json:"time"`
}

type CreateEngineResponse struct {
	Succ   bool    `json:"success"`
	Result Engine  `json:"result"`
	Error  string  `json:"error"`
	Time   float64 `json:"time"`
}

type LoginResponse struct {
	Succ   bool      `json:"success"`
	Result AuthToken `json:"result"`
	Error  string    `json:"error"`
	Time   float64   `json:"time"`
}

func (params *CreateProjectParentParams) AppendToBody(v *url.Values) {
	if len(params.AccountId) > 0 {
		v.Add("accountid", params.AccountId)
	}
	if len(params.Name) > 0 {
		v.Add("name", params.Name)
	}
	environments, _ := json.Marshal(params.Environments)
	v.Add("environments", string(environments))
}

func (params *CreateProjectParams) AppendToBody(v *url.Values) {
	if len(params.AccountId) > 0 {
		v.Add("accountid", params.AccountId)
	}
	if len(params.Environment) > 0 {
		v.Add("environment", params.Environment)
	}
	if len(params.ProjectParentId) > 0 {
		v.Add("parentid", params.ProjectParentId)
	}
}

func (params *CreateEngineParams) AppendToBody(v *url.Values) {
	if len(params.Name) > 0 {
		v.Add("name", params.Name)
	}
	if len(params.AccountId) > 0 {
		v.Add("accountid", params.AccountId)
	}
	v.Add("projectid", strconv.FormatUint(uint64(params.ProjectId), 10))
	behaviors, _ := json.Marshal(params.Behaviors)
	v.Add("behaviors", string(behaviors))
	filter, _ := json.Marshal(params.ItemsFilter)
	v.Add("filter", string(filter))
}
