package tamber

import (
	"strconv"
	"time"
)

const (
	RL_KEY  = "X-Rate-Limit-Limit"
	RLR_KEY = "X-Rate-Limit-Remaining"
	RLT_KEY = "X-Rate-Limit-Reset"
)

type Response interface {
	SetInfo(ResponseInfo)
}

type TamberResponse struct {
	Succ  bool    `json:"success"`
	Error string  `json:"error"`
	Time  float64 `json:"time"`
	ResponseInfo
}

func (r *TamberResponse) SetInfo(info ResponseInfo) {
	r.ResponseInfo = info
}

type ResponseInfo struct {
	HTTPCode           int     // HTTP status code
	RateLimit          int     // Limit-per-period for request method
	RateLimitRemaining int     // Requests remaining in current window for request method
	RateLimitReset     int     // Time in seconds until rate limits are reset
	Time               float64 `json:"time"`
}

func (info ResponseInfo) Duration() time.Duration {
	return time.Duration(int64(info.Time*1000)) * time.Millisecond
}

func (s *SessionConfig) NewResponse(HTTPCode int, Headers map[string][]string) ResponseInfo {
	var rl, rlr, rlt int
	var err error
	for k, vals := range Headers {
		switch k {
		case RL_KEY:
			if len(vals) > 0 {
				rl, err = strconv.Atoi(vals[0])
				if err != nil {
					s.errFunc("Cannot parse "+k+" header", err)
				}
			}
		case RLR_KEY:
			rlr, err = strconv.Atoi(vals[0])
			if err != nil {
				s.errFunc("Cannot parse "+k+" header", err)
			}
		case RLT_KEY:
			rlt, err = strconv.Atoi(vals[0])
			if err != nil {
				s.errFunc("Cannot parse "+k+" header", err)
			}
		}
	}
	return ResponseInfo{
		HTTPCode:           HTTPCode,
		RateLimit:          rl,
		RateLimitRemaining: rlr,
		RateLimitReset:     rlt,
	}
}
