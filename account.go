package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

var (
	EventsDatasetName = "events_dataset"
	ItemsDatasetName  = "items_dataset"
)

// Inputs
type UploadParams struct {
	ProjectId uint32 `json:"projectid"`
	Filepath  string `json:"filepath"`
	Type      string `json:"type"`
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
	ApiVersion      string                 `json:"api_version"`
	Engines         []string               `json:"engines"`
	Metadata        map[string]interface{} `json:"metadata"`
	State           int                    `json:"state"`
	Behaviors       []string               `json:"behaviors"`
	Dashboard       DashboardData          `json:"dashboard"`
	Object          string                 `json:"object"`
	Created         int64                  `json:"created"`
	Datasets        map[string]Dataset     `json:"datasets"`
}

type ProjectKey struct {
	Id          uint32 `json:"id"`
	Environment string `json:"environment"`
}

type ProjectParent struct {
	Id         string       `json:"id"`
	Name       string       `json:"name"`
	ApiVersion string       `json:"api_version"`
	Projects   []ProjectKey `json:"projects"`
}

type Engine struct {
	Key         string                 `json:"key"`
	Id          uint32                 `json:"id"`
	IdStr       string                 `json:"id_str"`
	ProjectId   uint32                 `json:"projectid"`
	ProjectKey  string                 `json:"project_key"`
	Name        string                 `json:"name"`
	Status      int                    `json:"status"`
	ApiVersion  string                 `json:"api_version"`
	Dashboard   DashboardData          `json:"dashboard"`
	Behaviors   map[string]Behavior    `json:"behaviors"`
	ItemsFilter map[string]interface{} `json:"filter"`
	Object      string                 `json:"object"`
	Created     int64                  `json:"created"`
}

type Dataset struct {
	Id        string `json:"id"`
	ProjectId uint32 `json:"projectid"`
	AccountId string `json:"accountid"`
	Name      string `json:"name"`
	Type      string `json:"type"`     // behaviors_dataset, items_dataset
	FileType  string `json:"filetype"` // json or csv
	FileSize  int64  `json:"filesize"`
	Object    string `json:"object"`
	Created   int64  `json:"created"`
}

type AccountInfo struct {
	Id             string                   `json:"id"`
	Username       string                   `json:"username"`
	ProjectParents map[string]ProjectParent `json:"project_parents"` // mapkey is projectParent.Id
	Projects       map[uint32]Project       `json:"projects"`        // mapkey is project.Id
	Engines        map[uint32]Engine        `json:"engines"`         // mapkey is engine.Id
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
