package account

import (
	"errors"
	tamber "github.com/tamber/tamber-go"
	"net/url"
	"time"
)

var (
	object = "account"
)

func UploadEventsDataset(filepath string) (*tamber.Dataset, error) {
	return getAccount().UploadEventsDataset(filepath)
}

func (a *Account) UploadEventsDataset(filepath string) (*tamber.Dataset, error) {
	params := &tamber.UploadParams{Filepath: filepath, Type: tamber.EventsDatasetName}
	return a.UploadDataset(params)
}

func UploadItemsDataset(filepath string) (*tamber.Dataset, error) {
	return getAccount().UploadItemsDataset(filepath)
}

func (a *Account) UploadItemsDataset(filepath string) (*tamber.Dataset, error) {
	params := &tamber.UploadParams{Filepath: filepath, Type: tamber.ItemsDatasetName}
	return a.UploadDataset(params)
}

func (a *Account) UploadDataset(params *tamber.UploadParams) (*tamber.Dataset, error) {
	dataset := &tamber.UploadResponse{}
	var err error

	err = a.updateToken()
	if err != nil {
		return nil, err
	}

	if len(params.Filepath) > 0 {
		err = a.S.CallUpload("POST", "", a.AuthToken.AccountId, a.AuthToken.Token, object, "uploadDataset", params.Filepath, params.Type, dataset)
	} else {
		err = errors.New("Invalid upload dataset params: filepath needs to be set")
	}

	if !dataset.Succ {
		err = errors.New(dataset.Error)
	}
	return &dataset.Result, err
}

func CreateProjectParent(params *tamber.CreateProjectParentParams) (*tamber.ProjectParent, error) {
	return getAccount().CreateProjectParent(params)
}

func (a *Account) CreateProjectParent(params *tamber.CreateProjectParentParams) (*tamber.ProjectParent, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	parent := &tamber.CreateProjectParentResponse{}
	var err error

	err = a.updateToken()
	if err != nil {
		return nil, err
	}

	if len(params.Name) > 0 {
		err = a.S.Call("POST", "", a.AuthToken.AccountId, a.AuthToken.Token, object, "createProjectParent", body, parent)
	} else {
		err = errors.New("Invalid create parent params: name needs to be set")
	}

	if !parent.Succ {
		err = errors.New(parent.Error)
	}
	return &parent.Result, err
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

	if len(params.ProjectParentId) > 0 {
		err = a.S.Call("POST", "", a.AuthToken.AccountId, a.AuthToken.Token, object, "createProject", body, project)
	} else {
		err = errors.New("Invalid create project params: ProjectParentId needs to be set")
	}

	if !project.Succ {
		err = errors.New(project.Error)
	}
	return &project.Result, err
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
		err = a.S.Call("POST", "", a.AuthToken.AccountId, a.AuthToken.Token, object, "createEngine", body, engine)
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
	a.S.Call("GET", "", a.Email, a.Password, object, "retrainEngine", body, engine)

	if !engine.Succ {
		err = errors.New(engine.Error)
	}
	return &engine.Result, err
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

	err = a.S.Call("POST", "", a.AuthToken.AccountId, a.AuthToken.Token, object, "retrieve", body, resp)

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

	err = a.S.Call("POST", "", a.Email, a.Password, object, "login", body, resp)

	if !resp.Succ {
		err = errors.New(resp.Error)
	}
	return &resp.Result, err
}

func getAccount() *Account {
	return &Account{GetDefaultAccountSessionConfig(), tamber.DefaultAccountEmail, tamber.DefaultAccountPassword, tamber.DefaultAuthToken}
}
