package tamber

import (
	"strconv"
)

const (
	RL_KEY  = "X-Rate-Limit-Limit"
	RLR_KEY = "X-Rate-Limit-Remaining"
	RLT_KEY = "X-Rate-Limit-Reset"
)

type Response interface {
	SetInfo(ResponseInfo)
}

type ResponseInfo struct {
	HTTPCode           int // HTTP status code
	RateLimit          int // Limit-per-period for request method
	RateLimitRemaining int // Requests remaining in current window for request method
	RateLimitReset     int // Time in seconds until rate limits are reset
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
