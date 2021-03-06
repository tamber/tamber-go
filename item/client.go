package item

import (
	"errors"
	"net/url"
	"sync"
	"time"

	tamber "github.com/tamber/tamber-go"
)

type Client struct {
	S          *tamber.SessionConfig
	ProjectKey string
	EngineKey  string
}

var object = "item"

func Create(params *tamber.ItemParams) (*tamber.Item, *tamber.ResponseInfo, error) {
	return getClient().Create(params)
}

func (c Client) Create(params *tamber.ItemParams) (*tamber.Item, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.ItemResponse{}
	var err error

	if len(params.Id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "create", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}

	if err == nil && !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, &item.ResponseInfo, err
}

func Save(params *tamber.ItemSaveParams) (*tamber.Item, *tamber.ResponseInfo, error) {
	return getClient().Save(params)
}

func (c Client) Save(params *tamber.ItemSaveParams) (*tamber.Item, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.ItemResponse{}
	var err error

	if len(params.Id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "save", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}

	if err == nil && !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, &item.ResponseInfo, err
}

func Update(params *tamber.ItemUpdateParams) (*tamber.Item, *tamber.ResponseInfo, error) {
	return getClient().Update(params)
}

func (c Client) Update(params *tamber.ItemUpdateParams) (*tamber.Item, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.ItemResponse{}
	var err error

	if len(params.Id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "update", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}

	if err == nil && !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, &item.ResponseInfo, err
}

func Batch(params *tamber.ItemBatchParams) (*tamber.Items, *tamber.ResponseInfo, error) {
	return getClient().Batch(params)
}

func (c Client) Batch(params *tamber.ItemBatchParams) (*tamber.Items, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.ItemsResponse{}
	var err error

	if len(params.Items) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "batch", body, item)
	} else {
		err = errors.New("Invalid batch params: `Items` needs to be set")
	}

	if err == nil && !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, &item.ResponseInfo, err
}

func Stream(items []*tamber.ItemUpdateParams, out *chan *tamber.Item, numThreads, bufSize int) (*tamber.ResponseInfo, error) {
	return getClient().Stream(items, out, numThreads, bufSize)
}

func (c Client) Stream(items []*tamber.ItemUpdateParams, out *chan *tamber.Item, numThreads, bufSize int) (*tamber.ResponseInfo, error) {
	in := make(chan *tamber.ItemUpdateParams, bufSize)
	var wg sync.WaitGroup

	stop := make(chan struct {
		info *tamber.ResponseInfo
		err  error
	}, 1)

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(thread int) {
			defer wg.Done()
			for itemParams := range in {
				select {
				case resp := <-stop:
					stop <- resp
					return
				default:
				}
				item, info, err := c.Update(itemParams)
				if err != nil {
					resp := struct {
						info *tamber.ResponseInfo
						err  error
					}{info, err}
					select {
					case stop <- resp:
						return
					default:
						return
					}
				}
				if out != nil {
					*out <- item
				}
				if info.RateLimitRemaining <= numThreads {
					waitv := time.Second * time.Duration(info.RateLimitReset)
					time.Sleep(waitv)
				}
			}
		}(i)
	}
	for i, itemParams := range items {
		// ensure rate limits are acceptable to begin multi-threaded streaming
		if i == 0 {
			item, info, err := c.Update(itemParams)
			if err != nil && info.HTTPCode != 429 {
				return info, err
			} else if info.RateLimitRemaining < numThreads {
				time.Sleep(time.Second * time.Duration(info.RateLimitReset))
				if err != nil { // update failed due to rate limits
					in <- itemParams
				} else if out != nil { // update successful
					*out <- item
				}
			} else if err != nil {
				return info, err
			} else if out != nil {
				*out <- item
			}
		} else {
			in <- itemParams
		}
	}
	close(in)
	wg.Wait()
	select {
	case resp := <-stop:
		return resp.info, resp.err
	default:
		return nil, nil
	}

}

func OpenChanStream(in chan *tamber.ItemUpdateParams, out *chan *tamber.Item, numThreads, bufSize int) (*tamber.ResponseInfo, error) {
	return getClient().OpenChanStream(in, out, numThreads, bufSize)
}

func (c Client) OpenChanStream(in chan *tamber.ItemUpdateParams, out *chan *tamber.Item, numThreads, bufSize int) (*tamber.ResponseInfo, error) {
	var wg sync.WaitGroup

	stop := make(chan struct {
		info *tamber.ResponseInfo
		err  error
	}, 1)
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for itemParams := range in {
				select {
				case resp := <-stop:
					stop <- resp
					return
				default:
				}
				item, info, err := c.Update(itemParams)
				if err != nil {
					resp := struct {
						info *tamber.ResponseInfo
						err  error
					}{info, err}
					select {
					case stop <- resp:
						return
					default:
						return
					}
				}
				if out != nil {
					*out <- item
				}
				if info.RateLimitRemaining < numThreads {
					time.Sleep(time.Second * time.Duration(info.RateLimitReset))
				}
			}
		}()
	}
	wg.Wait()
	select {
	case resp := <-stop:
		return resp.info, resp.err
	default:
		return nil, nil
	}
}

func Retrieve(params *tamber.ItemParams) (*tamber.Item, *tamber.ResponseInfo, error) {
	return getClient().Retrieve(params)
}

func (c Client) Retrieve(params *tamber.ItemParams) (*tamber.Item, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.ItemResponse{}
	var err error

	if len(params.Id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "retrieve", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}

	if err == nil && !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, &item.ResponseInfo, err
}

func List(params *tamber.ItemListParams) (tamber.Items, *tamber.ResponseInfo, error) {
	return getClient().List(params)
}

func (c Client) List(params *tamber.ItemListParams) (tamber.Items, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	result := &tamber.ItemsResponse{}

	err := c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "list", body, result)

	if err == nil && !result.Succ {
		err = errors.New(result.Error)
	}
	return result.Result, &result.ResponseInfo, err
}

func Hide(id string) (*tamber.Item, *tamber.ResponseInfo, error) {
	return getClient().Hide(id)
}

func (c Client) Hide(id string) (*tamber.Item, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	body.Add("id", id)
	item := &tamber.ItemResponse{}
	var err error

	if len(id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "hide", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}
	if err == nil && !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, &item.ResponseInfo, err
}

func Unhide(id string) (*tamber.Item, *tamber.ResponseInfo, error) {
	return getClient().Unhide(id)
}

func (c Client) Unhide(id string) (*tamber.Item, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	body.Add("id", id)
	item := &tamber.ItemResponse{}
	var err error

	if len(id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "unhide", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}
	if err == nil && !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, &item.ResponseInfo, err
}

func Delete(id string) (*tamber.Item, *tamber.ResponseInfo, error) {
	return getClient().Delete(id)
}

func (c Client) Delete(id string) (*tamber.Item, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	body.Add("id", id)
	item := &tamber.ItemResponse{}
	var err error

	if len(id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "delete", body, item)
	} else {
		err = errors.New("Invalid item params: id needs to be set")
	}
	if err == nil && !item.Succ {
		err = errors.New(item.Error)
	}
	return &item.Result, &item.ResponseInfo, err
}

func getClient() Client {
	return Client{tamber.GetDefaultSessionConfig(), tamber.DefaultProjectKey, tamber.DefaultEngineKey}
}
