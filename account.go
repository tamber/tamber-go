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

type CreateProjectParams struct {
	Name string `json:"name"`
}

type EngineConfig struct {
	DiscoverConfigs            map[string]interface{} `bson:"discover_configs,omitempty" json:"discover_configs,omitempty"`
	Notifications              map[string]interface{} `bson:"notifications,omitempty" json:"notifications,omitempty"`
	RecOnlyContexts            map[string]struct{}    `bson:"rec_only_ctx" json:"rec_only_ctx"`
	HideItemBehaviorThresholds map[string]float64     `bson:"hide_item_behavior_thresh,omitempty" json:"hide_item_behavior_thresh,omitempty"`
}

type CreateEngineParams struct {
	Name           string                 `json:"name"`
	ProjectId      uint32                 `json:"projectid"`
	Behaviors      map[string]Behavior    `json:"behaviors"`
	ItemsFilter    map[string]interface{} `json:"filter"`
	ProjectDefault *bool                  `json:"project_default"`
	Config         *EngineConfig          `json:"conf"`
}

// Types
type DashboardData struct {
	BehaviorCount int64 `json:"event_count"`
	ItemCount     int64 `json:"item_count"`
	UserCount     int64 `json:"user_count"`
}

type Project struct {
	Id              uint32                 `json:"id"`
	Key             string                 `json:"key"`
	Name            string                 `json:"name"`
	AccountId       string                 `json:"accountid"`
	ProjectParentId string                 `json:"parentid"`
	ApiVersion      string                 `json:"api_version"`
	Engines         []uint32               `json:"engines"`
	Metadata        map[string]interface{} `json:"metadata"`
	State           int                    `json:"state"`
	Behaviors       []string               `json:"behaviors"`
	Dashboard       DashboardData          `json:"dashboard"`
	Object          string                 `json:"object"`
	Created         int64                  `json:"created"`
	Datasets        map[string]Dataset     `json:"datasets"`
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
	Id       string             `json:"id"`
	Username string             `json:"username"`
	Projects map[uint32]Project `json:"projects"` // mapkey is project.Id
	Engines  map[uint32]Engine  `json:"engines"`  // mapkey is engine.Id
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
	ResponseInfo
}

type UploadResponse struct {
	Succ   bool    `json:"success"`
	Result Dataset `json:"result"`
	Error  string  `json:"error"`
	Time   float64 `json:"time"`
	ResponseInfo
}

type CreateProjectResponse struct {
	Succ   bool    `json:"success"`
	Result Project `json:"result"`
	Error  string  `json:"error"`
	Time   float64 `json:"time"`
	ResponseInfo
}

type DeleteProjectResponse struct {
	Succ   bool    `json:"success"`
	Result uint32  `json:"result"`
	Error  string  `json:"error"`
	Time   float64 `json:"time"`
	ResponseInfo
}

type CreateEngineResponse struct {
	Succ   bool    `json:"success"`
	Result Engine  `json:"result"`
	Error  string  `json:"error"`
	Time   float64 `json:"time"`
	ResponseInfo
}

type LoginResponse struct {
	Succ   bool      `json:"success"`
	Result AuthToken `json:"result"`
	Error  string    `json:"error"`
	Time   float64   `json:"time"`
	ResponseInfo
}

func (r *AccountResponse) SetInfo(info ResponseInfo) {
	r.ResponseInfo = info
}
func (r *UploadResponse) SetInfo(info ResponseInfo) {
	r.ResponseInfo = info
}
func (r *CreateProjectResponse) SetInfo(info ResponseInfo) {
	r.ResponseInfo = info
}
func (r *DeleteProjectResponse) SetInfo(info ResponseInfo) {
	r.ResponseInfo = info
}
func (r *CreateEngineResponse) SetInfo(info ResponseInfo) {
	r.ResponseInfo = info
}
func (r *LoginResponse) SetInfo(info ResponseInfo) {
	r.ResponseInfo = info
}

func (params *CreateProjectParams) AppendToBody(v *url.Values) {
	if len(params.Name) > 0 {
		v.Add("name", params.Name)
	}
}

func (params *CreateEngineParams) AppendToBody(v *url.Values) {
	if len(params.Name) > 0 {
		v.Add("name", params.Name)
	}
	if params.ProjectDefault != nil {
		v.Add("project_default", strconv.FormatBool(*params.ProjectDefault))
	}
	v.Add("projectid", strconv.FormatUint(uint64(params.ProjectId), 10))
	behaviors, _ := json.Marshal(params.Behaviors)
	v.Add("behaviors", string(behaviors))
	filter, _ := json.Marshal(params.ItemsFilter)
	v.Add("filter", string(filter))
	config, _ := json.Marshal(params.Config)
	v.Add("config", string(config))

}
