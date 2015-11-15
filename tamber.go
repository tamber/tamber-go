package tamber

import (
	"bytes"
	"encoding/json"
	// "errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	apiUrl = "https://dev.tamber.com/v1"
)

// apiversion is the currently supported API version
const apiversion = "2015-11-7"

// clientversion is the binding version
const clientversion = "0.0.1"

// defaultHTTPTimeout is the default timeout on the http.Client used by the library.
const defaultHTTPTimeout = 80 * time.Second

type SessionErrFunction func(exp string, err interface{})

// Session is an interface for making calls against Tamber services.
// This interface exists to enable mocking for during testing if needed.
type Session interface {
	Call(method, key string, body *url.Values, resp interface{}) error
	// CallMultipart(method, path, key, boundary string, body io.Reader, v interface{}) error
}

type SessionConfig struct {
	URL        string
	HTTPClient *http.Client
	errFunc    SessionErrFunction
}

type Engine struct {
	Key string
	S   *SessionConfig
}

//Default global engine key
var DefaultKey string

var httpClient = &http.Client{Timeout: defaultHTTPTimeout}

func NewEngine(key string, config *SessionConfig) Engine {
	if config == nil {
		config = GetDefaultSessionConfig()
	}
	return Engine{key, config}
}

func GetDefaultSessionConfig() *SessionConfig {
	return &SessionConfig{apiUrl, httpClient, defaultErrFunc}
}

func (s *SessionConfig) Call(method, path, key, object, command string, form *url.Values, resp interface{}) error {
	var body io.Reader
	if form != nil && len(*form) > 0 {
		form.Add("command", command)
		data := form.Encode()
		if strings.ToUpper(method) == "GET" {
			path += "?" + data
		} else {
			body = bytes.NewBufferString(data)
		}
	}
	path += object
	req, err := s.NewRequest(method, path, key, "application/x-www-form-urlencoded", body)
	if err != nil {
		return err
	}

	if err := s.Do(req, resp); err != nil {
		return err
	}

	return nil
}

// NewRequest is used by Call to generate an http.Request. It handles encoding
// parameters and attaching the appropriate headers.
func (s *SessionConfig) NewRequest(method, path, key, contentType string, body io.Reader) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = s.URL + path

	req, err := http.NewRequest(method, path, body)

	if err != nil {
		s.errFunc("Cannot create Tamber request", err)
		return nil, err
	}

	req.SetBasicAuth(key, "")

	req.Header.Add("Tamber-Version", apiversion)
	req.Header.Add("User-Agent", "Tamber/v1 GoBindings/"+clientversion)
	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// Do is used by Call to execute an API request and parse the response. It uses
// the backend's HTTP client to execute the request and unmarshals the response
// into v. It also handles unmarshaling errors returned by the API.
func (s *SessionConfig) Do(req *http.Request, v interface{}) error {

	res, err := s.HTTPClient.Do(req)

	if err != nil {
		s.errFunc("Request to Tamber failed", err)
		return err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		s.errFunc("Cannot parse Tamber response", err)
		return err
	}
	err = json.Unmarshal(resBody, v)
	if err != nil {
		s.errFunc("Json error", err)
	}

	return nil
}

// Set a new error handling function, which handles errors encountered
// When executing API requests. By default this is a log.Printf
func (s *SessionConfig) SetErrFunc(errFunc SessionErrFunction) {
	s.errFunc = errFunc
}

func defaultErrFunc(exp string, err interface{}) {
	log.Printf("\n%s: %v\n", exp, err)
}
