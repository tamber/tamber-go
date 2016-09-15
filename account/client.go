package account

import (
	"errors"
	tamber "github.com/tamber/tamber-go"
	"net/url"
)

var (
	object = "account"
)

func UploadEventsDataset(filepath string) (*tamber.Dataset, error) {
	return getAccount().UploadEventsDataset(filepath)
}

func (a Account) UploadEventsDataset(filepath string) (*tamber.Dataset, error) {
	params := &tamber.UploadParams{Filepath: filepath, Type: tamber.EventsDatasetName}
	return a.UploadDataset(params)
}

func UploadItemsDataset(filepath string) (*tamber.Dataset, error) {
	return getAccount().UploadItemsDataset(filepath)
}

func (a Account) UploadItemsDataset(filepath string) (*tamber.Dataset, error) {
	params := &tamber.UploadParams{Filepath: filepath, Type: tamber.ItemsDatasetName}
	return a.UploadDataset(params)
}

func (a Account) UploadDataset(params *tamber.UploadParams) (*tamber.Dataset, error) {
	dataset := &tamber.UploadResponse{}
	var err error

	if len(params.Filepath) > 0 {
		err = a.S.CallUpload("POST", "", a.Email, a.Password, object, "uploadDataset", params.Filepath, params.Type, dataset)
	} else {
		err = errors.New("Invalid upload dataset params: filepath needs to be set")
	}

	if !dataset.Succ {
		err = errors.New(dataset.Error)
	}
	return &dataset.Result, err
}

func CreateEngine(params *tamber.CreateEngineParams) (*tamber.Engine, error) {
	return getAccount().CreateEngine(params)
}

func (a Account) CreateEngine(params *tamber.CreateEngineParams) (*tamber.Engine, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	engine := &tamber.CreateEngineResponse{}
	var err error

	if len(params.Name) > 0 {
		err = a.S.Call("POST", "", a.Email, a.Password, object, "createEngine", body, engine)
	} else {
		err = errors.New("Invalid create engine params: name needs to be set")
	}

	if !engine.Succ {
		err = errors.New(engine.Error)
	}
	return &engine.Result, err
}

func Retrieve() (*tamber.AccountInfo, error) {
	return getAccount().Retrieve()
}

func (a Account) Retrieve() (*tamber.AccountInfo, error) {
	body := &url.Values{}
	resp := &tamber.AccountResponse{}
	var err error

	err = a.S.Call("POST", "", a.Email, a.Password, object, "retrieve", body, resp)

	if !resp.Succ {
		err = errors.New(resp.Error)
	}
	return &resp.Result, err
}

func getAccount() Account {
	return Account{GetDefaultAccountSessionConfig(), tamber.DefaultAccountEmail, tamber.DefaultAccountPassword}
}
