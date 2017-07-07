package item

import (
	"errors"
	tamber "github.com/tamber/tamber-go"
	"net/url"
	"sync"
	"time"
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

func Update(params *tamber.ItemParams) (*tamber.Item, *tamber.ResponseInfo, error) {
	return getClient().Update(params)
}

func (c Client) Update(params *tamber.ItemParams) (*tamber.Item, *tamber.ResponseInfo, error) {
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

func OpenChanStream(in chan *tamber.ItemParams, out *chan *tamber.Item, numThreads, bufSize int) (*tamber.ResponseInfo, error) {
	return getClient().OpenChanStream(in, out, numThreads, bufSize)
}

func (c Client) OpenChanStream(in chan *tamber.ItemParams, out *chan *tamber.Item, numThreads, bufSize int) (*tamber.ResponseInfo, error) {
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

func Stream(items []*tamber.ItemParams, out *chan *tamber.Item, numThreads, bufSize int) (*tamber.ResponseInfo, error) {
	return getClient().Stream(items, out, numThreads, bufSize)
}

func (c Client) Stream(items []*tamber.ItemParams, out *chan *tamber.Item, numThreads, bufSize int) (*tamber.ResponseInfo, error) {
	in := make(chan *tamber.ItemParams, bufSize)
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

func Remove(params *tamber.ItemParams) (*tamber.Item, *tamber.ResponseInfo, error) {
	return getClient().Remove(params)
}

func (c Client) Remove(params *tamber.ItemParams) (*tamber.Item, *tamber.ResponseInfo, error) {
	body := &url.Values{}
	params.AppendToBody(body)
	item := &tamber.ItemResponse{}
	var err error

	if len(params.Id) > 0 {
		err = c.S.Call("POST", "", c.ProjectKey, c.EngineKey, object, "remove", body, item)
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
