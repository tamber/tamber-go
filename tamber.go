package tamber

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	ApiUrl = "https://api.tamber.com/v1"
)

// apiversion is the currently supported API version
const apiversion = "2017-3-8"

// clientversion is the binding version
const clientversion = "0.0.4"

// defaultHTTPTimeout is the default timeout on the http.Client used by the library.
const defaultHTTPTimeout = 80 * time.Second

type SessionErrFunction func(exp string, err interface{})

type SessionConfig struct {
	URL        string
	HTTPClient *http.Client
	errFunc    SessionErrFunction
}

// Default global keys
var (
	DefaultProjectKey string
	DefaultEngineKey  string

	DefaultAccountEmail    string
	DefaultAccountPassword string

	DefaultAuthToken *AuthToken
	DefaultErrFunc   SessionErrFunction = func(exp string, err interface{}) {
		log.Printf("\n%s: %v\n", exp, err)
	}
)

var httpClient = &http.Client{Timeout: defaultHTTPTimeout}

func GetDefaultSessionConfig() *SessionConfig {
	return &SessionConfig{ApiUrl, httpClient, DefaultErrFunc}
}

func (s *SessionConfig) Call(method, path, key, ext, object, command string, form *url.Values, resp Response) error {
	var body io.Reader
	if form != nil && len(*form) > 0 {
		data := form.Encode()
		if strings.ToUpper(method) == "GET" {
			command += "?" + data
		} else {
			body = bytes.NewBufferString(data)
		}
	}
	path += object + "/" + command
	req, err := s.NewRequest(method, path, key, ext, "application/x-www-form-urlencoded", body)
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
func (s *SessionConfig) NewRequest(method, path, key, ext, contentType string, body io.Reader) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = s.URL + path

	req, err := http.NewRequest(method, path, body)

	if err != nil {
		s.errFunc("Cannot create Tamber request", err)
		return nil, err
	}

	req.SetBasicAuth(key, ext)

	req.Header.Add("Tamber-Version", apiversion)
	req.Header.Add("User-Agent", "Tamber/v1 GoBindings/"+clientversion)
	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// Do is used by Call to execute an API request and parse the response. It uses
// the backend's HTTP client to execute the request and unmarshals the response
// into v. It also handles unmarshaling errors returned by the API.
func (s *SessionConfig) Do(req *http.Request, v Response) error {

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

	info := s.NewResponse(res.StatusCode, res.Header)
	v.SetInfo(info)

	return nil
}

func (s *SessionConfig) AccountDo(req *http.Request, v interface{}) error {

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
