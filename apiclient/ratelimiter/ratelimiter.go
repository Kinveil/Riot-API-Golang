package ratelimiter

import (
	"net/http"
	"sync"
)

type ConserveUsage struct {
	RegionPercent int
	MethodPercent int
	IgnoreLimits  []MethodID
}

type RateLimiter struct {
	Requests       chan *APIRequest
	httpClient     *http.Client
	apiKey         string
	maxRetries     int
	conserveUsage  ConserveUsage
	regionMutex    sync.Mutex
	methodMutex    sync.Mutex
	regionLimiters map[string]*RateLimit
	methodLimiters map[string]*RateLimit
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewRateLimiter(requests chan *APIRequest, apiKey string) *RateLimiter {
	if requests == nil {
		panic("requests channel cannot be nil")
	}

	return &RateLimiter{
		Requests:   requests,
		httpClient: &http.Client{},
		apiKey:     apiKey,
		maxRetries: -1,
		conserveUsage: ConserveUsage{
			RegionPercent: 0,
			MethodPercent: 0,
			IgnoreLimits:  []MethodID{},
		},
		regionMutex:    sync.Mutex{},
		methodMutex:    sync.Mutex{},
		regionLimiters: make(map[string]*RateLimit),
		methodLimiters: make(map[string]*RateLimit),
	}
}

// SetUsageConservation sets the usage conservation percentages for regions and methods.
// Both region and method percentages must be between 0 and 100.
func (rl *RateLimiter) SetUsageConservation(conserveUsage ConserveUsage) {
	if conserveUsage.RegionPercent < 0 || conserveUsage.RegionPercent > 100 {
		panic("regionPercent must be between 0 and 100")
	}

	if conserveUsage.MethodPercent < 0 || conserveUsage.MethodPercent > 100 {
		panic("methodPercent must be between 0 and 100")
	}

	rl.conserveUsage = conserveUsage
}

func (rl *RateLimiter) SetAPIKey(apiKey string) {
	rl.apiKey = apiKey
}

// SetMaxRetries sets the maximum number of retries for a request.
// If maxRetries is less than 0, then the request will be retried indefinitely.
func (rl *RateLimiter) SetMaxRetries(maxRetries int) {
	if maxRetries < -1 {
		rl.maxRetries = -1
		return
	}

	rl.maxRetries = maxRetries
}

func (rl *RateLimiter) Start() {
	for req := range rl.Requests {
		go rl.handleRequest(req)
	}
}
