package behavior

import (
	. "github.com/tamber/tamber-go"
	"net/url"
	"strconv"
)

var object = "behavior"

func Create(params *BehaviorParams) (*Behavior, error) {
	return getEngine().Create(params)
}

func (e Engine) Create(params *BehaviorParams) (*Behavior, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	behavior := &Behavior{}
	var err error

	if len(params.Name) > 0 && len(params.Type) > 0 && params.Desirability > 0 {
		err = e.S.Call("POST", "", e.Key, object, "create", body, behavior)
	} else {
		err = errors.New("Invalid behavior params: name, type, and desirability need to be set")
	}

	return behavior, err
}

func Retrieve(params *BehaviorParams) (*Behavior, error) {
	return getEngine().Create(params)
}

func (e Engine) Retrieve(params *BehaviorParams) (*Behavior, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	behavior := &Behavior{}
	var err error

	if len(params.Name) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "retrieve", body, behavior)
	} else {
		err = errors.New("Invalid behavior params: name needs to be set")
	}

	return behavior, err
}

func Remove(params *BehaviorParams) (*Behavior, error) {
	return getEngine().Create(params)
}

func (e Engine) Remove(params *BehaviorParams) (*Behavior, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	behavior := &Behavior{}
	var err error

	if len(params.Name) > 0 {
		err = e.S.Call("POST", "", e.Key, object, "remove", body, behavior)
	} else {
		err = errors.New("Invalid behavior params: name needs to be set")
	}

	return behavior, err
}

func getEngine() Engine {
	return Engine{GetDefaultSessionConfig(), DefaultKey}
}
