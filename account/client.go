package account

import (
	"errors"
	"net/url"
	"strconv"
	"time"

	tamber "github.com/tamber/tamber-go"
)

var (
	accountObject       = "account"
	projectObject       = "project"
	projectParentObject = "project_parent"
	engineObject        = "engine"
	blockObject         = "block"
)

func UploadEventsDataset(projectId uint32, filepath string) (*tamber.Dataset, error) {
	return getAccount().UploadEventsDataset(projectId, filepath)
}

func (a *Account) UploadEventsDataset(projectId uint32, filepath string) (*tamber.Dataset, error) {
	params := &tamber.UploadParams{ProjectId: projectId, Filepath: filepath, Type: tamber.EventsDatasetName}
	return a.UploadDataset(params)
}

// func UploadItemsDataset(projectId uint32, filepath string) (*tamber.Dataset, error) {
// 	return getAccount().UploadItemsDataset(projectId, filepath)
// }

// func (a *Account) UploadItemsDataset(projectId uint32, filepath string) (*tamber.Dataset, error) {
// 	params := &tamber.UploadParams{ProjectId: projectId, Filepath: filepath, Type: tamber.ItemsDatasetName}
// 	return a.UploadDataset(params)
// }

func (a *Account) UploadDataset(params *tamber.UploadParams) (*tamber.Dataset, error) {
	dataset := &tamber.UploadResponse{}
	var err error

	err = a.updateToken()
	if err != nil {
		return nil, err
	}

	if len(params.Filepath) > 0 {
		err = a.S.CallUpload("POST", "", a.AuthToken.AccountId, a.AuthToken.Token, projectObject, "upload_dataset", params, dataset)
	} else {
		err = errors.New("Invalid upload dataset params: filepath needs to be set")
	}

	if !dataset.Succ {
		err = errors.New(dataset.Error)
	}
	return &dataset.Result, err
}

func CreateProject(params *tamber.CreateProjectParams) (*tamber.Project, error) {
	return getAccount().CreateProject(params)
}

func (a *Account) CreateProject(params *tamber.CreateProjectParams) (*tamber.Project, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	project := &tamber.CreateProjectResponse{}
	var err error

	err = a.updateToken()
	if err != nil {
		return nil, err
	}

	if len(params.Name) > 0 {
		err = a.S.Call("POST", "", a.AuthToken.AccountId, a.AuthToken.Token, projectObject, "create", body, project)
	} else {
		err = errors.New("Invalid create project params: Name needs to be set")
	}

	if !project.Succ {
		err = errors.New(project.Error)
	}
	return &project.Result, err
}

func DeleteProject(projectId uint32) error {
	return getAccount().DeleteProject(projectId)
}

func (a *Account) DeleteProject(projectId uint32) error {
	body := &url.Values{}
	body.Add("id", strconv.FormatUint(uint64(projectId), 10))
	resp := &tamber.DeleteProjectResponse{}
	var err error

	err = a.updateToken()
	if err != nil {
		return err
	}

	err = a.S.Call("POST", "", a.AuthToken.AccountId, a.AuthToken.Token, projectObject, "delete", body, resp)

	if !resp.Succ {
		err = errors.New(resp.Error)
	}
	return err
}

func UpdateProjectDefaultEngine(projectId, engineId uint32) error {
	return getAccount().UpdateProjectDefaultEngine(projectId, engineId)
}

func (a *Account) UpdateProjectDefaultEngine(projectId, engineId uint32) error {
	body := &url.Values{}
	body.Add("id", strconv.FormatUint(uint64(projectId), 10))
	body.Add("eid", strconv.FormatUint(uint64(engineId), 10))
	resp := &tamber.CreateProjectResponse{}
	var err error

	err = a.updateToken()
	if err != nil {
		return err
	}

	err = a.S.Call("POST", "", a.AuthToken.AccountId, a.AuthToken.Token, projectObject, "update_default_engine", body, resp)
	if !resp.Succ {
		err = errors.New(resp.Error)
	}
	return err
}

func CreateEngine(params *tamber.CreateEngineParams) (*tamber.Engine, error) {
	return getAccount().CreateEngine(params)
}

func (a *Account) CreateEngine(params *tamber.CreateEngineParams) (*tamber.Engine, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	engine := &tamber.CreateEngineResponse{}
	var err error

	err = a.updateToken()
	if err != nil {
		return nil, err
	}

	if len(params.Name) > 0 {
		err = a.S.Call("POST", "", a.AuthToken.AccountId, a.AuthToken.Token, engineObject, "create", body, engine)
	} else {
		err = errors.New("Invalid create engine params: name needs to be set")
	}

	if !engine.Succ {
		err = errors.New(engine.Error)
	}
	return &engine.Result, err
}

func RetrainEngine(engineId uint32) (*tamber.Engine, error) {
	return getAccount().RetrainEngine(engineId)
}

func (a *Account) RetrainEngine(engineId uint32) (*tamber.Engine, error) {
	body := &url.Values{}
	body.Add("id", strconv.FormatUint(uint64(engineId), 10))
	engine := &tamber.CreateEngineResponse{}

	var err error

	err = a.updateToken()
	if err != nil {
		return nil, err
	}

	a.S.Call("GET", "", a.AuthToken.AccountId, a.AuthToken.Token, engineObject, "retrain", body, engine)

	if !engine.Succ {
		err = errors.New(engine.Error)
	}
	return &engine.Result, err
}

func DeleteEngine(engineId uint32) error {
	return getAccount().DeleteEngine(engineId)
}

func (a *Account) DeleteEngine(engineId uint32) error {
	body := &url.Values{}
	body.Add("id", strconv.FormatUint(uint64(engineId), 10))
	resp := &tamber.DeleteProjectResponse{}
	var err error

	err = a.updateToken()
	if err != nil {
		return err
	}

	err = a.S.Call("POST", "", a.AuthToken.AccountId, a.AuthToken.Token, engineObject, "delete", body, resp)

	if !resp.Succ {
		err = errors.New(resp.Error)
	}
	return err
}

func CreateBlock(params *tamber.CreateBlockParams) (*tamber.Block, error) {
	return getAccount().CreateBlock(params)
}

func (a *Account) CreateBlock(params *tamber.CreateBlockParams) (*tamber.Block, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	response := &tamber.BlockResponse{}
	var err error

	err = a.updateToken()
	if err != nil {
		return nil, err
	}

	if len(params.Name) > 0 {
		err = a.S.Call("POST", "", a.AuthToken.AccountId, a.AuthToken.Token, blockObject, "create", body, response)
	} else {
		err = errors.New("Invalid create engine params: name needs to be set")
	}

	if !response.Succ {
		err = errors.New(response.Error)
	}
	return &response.Result, err
}

func RemoveBlock(params *tamber.RemoveBlockParams) (*tamber.Block, error) {
	return getAccount().RemoveBlock(params)
}

func (a *Account) RemoveBlock(params *tamber.RemoveBlockParams) (*tamber.Block, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	response := &tamber.BlockResponse{}
	var err error

	err = a.updateToken()
	if err != nil {
		return nil, err
	}

	if len(params.Name) > 0 {
		err = a.S.Call("POST", "", a.AuthToken.AccountId, a.AuthToken.Token, blockObject, "create", body, response)
	} else {
		err = errors.New("Invalid create engine params: name needs to be set")
	}

	if !response.Succ {
		err = errors.New(response.Error)
	}
	return &response.Result, err
}

func Retrieve() (*tamber.AccountInfo, error) {
	return getAccount().Retrieve()
}

func (a *Account) Retrieve() (*tamber.AccountInfo, error) {
	body := &url.Values{}
	resp := &tamber.AccountResponse{}
	var err error

	err = a.updateToken()
	if err != nil {
		return nil, err
	}

	err = a.S.Call("POST", "", a.AuthToken.AccountId, a.AuthToken.Token, accountObject, "retrieve", body, resp)

	if !resp.Succ {
		err = errors.New(resp.Error)
	}
	return &resp.Result, err
}

func (a *Account) updateToken() error {
	if a.AuthToken == nil || a.AuthToken.ExpireTime < time.Now().UnixNano()/int64(time.Millisecond) {
		authToken, err := a.Login()
		if err != nil {
			return err
		}
		if a.AuthToken == tamber.DefaultAuthToken {
			tamber.DefaultAuthToken = authToken
		}
		a.AuthToken = authToken
	}
	return nil
}

func (a *Account) Login() (*tamber.AuthToken, error) {
	body := &url.Values{}
	resp := &tamber.LoginResponse{}
	var err error

	err = a.S.Call("POST", "", a.Email, a.Password, accountObject, "login", body, resp)

	if !resp.Succ {
		err = errors.New(resp.Error)
	}
	return &resp.Result, err
}

func getAccount() *Account {
	return &Account{GetDefaultAccountSessionConfig(), tamber.DefaultAccountEmail, tamber.DefaultAccountPassword, tamber.DefaultAuthToken}
}
