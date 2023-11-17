package ratelimiter

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ConserveUsage struct {
	RegionPercent int
	MethodPercent int
	IgnoreLimits  []MethodId
}

type RateLimiter struct {
	Requests      chan *APIRequest
	httpClient    *http.Client
	apiKey        string
	maxRetries    int
	conserveUsage ConserveUsage
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
			IgnoreLimits:  []MethodId{},
		},
	}
}

// Both region and method percentages must be between 0 and 100
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

// Set the maximum number of retries for a request. If maxRetries is less than 0, then the request will be retried indefinitely
func (rl *RateLimiter) SetMaxRetries(maxRetries int) {
	if maxRetries < -1 {
		rl.maxRetries = -1
		return
	}

	rl.maxRetries = maxRetries
}

type APIRequest struct {
	Region   string
	MethodId MethodId
	URL      string
	Response chan<- *http.Response
	Retries  int
}

type RateLimit struct {
	shortLimiter *limiterMu
	longLimiter  *limiterMu
	blockedUntil time.Time
}

const (
	initialRegionLimit = 20
	initialMethodLimit = 5
)

func (rl *RateLimiter) Start() {
	regionLimiters := make(map[string]*RateLimit)
	methodLimiters := make(map[string]*RateLimit)

	regionMutex := sync.RWMutex{}
	methodMutex := sync.RWMutex{}

	for req := range rl.Requests {
		var regionLimiter *RateLimit
		var methodLimiter *RateLimit

		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()

			var ok bool
			regionMutex.RLock()
			regionLimiter, ok = regionLimiters[req.Region]
			regionMutex.RUnlock()

			if !ok {
				regionMutex.Lock()
				regionLimiters[req.Region] = &RateLimit{
					shortLimiter: newLimiterMu(initialRegionLimit),
					longLimiter:  newLimiterMu(initialRegionLimit),
					blockedUntil: time.Time{},
				}
				regionLimiter = regionLimiters[req.Region]
				regionMutex.Unlock()
			}
		}()

		go func() {
			defer wg.Done()

			var ok bool
			methodMutex.RLock()
			methodLimiter, ok = methodLimiters[req.Region+req.MethodId.String()]
			methodMutex.RUnlock()

			if !ok {
				methodMutex.Lock()
				methodLimiters[req.Region+req.MethodId.String()] = &RateLimit{
					shortLimiter: newLimiterMu(initialMethodLimit),
					longLimiter:  newLimiterMu(initialMethodLimit),
					blockedUntil: time.Time{},
				}
				methodLimiter = methodLimiters[req.Region+req.MethodId.String()]
				methodMutex.Unlock()
			}
		}()

		wg.Wait()

		// Check if the region is blocked
		if time.Now().Before(regionLimiter.blockedUntil) {
			time.Sleep(time.Until(regionLimiter.blockedUntil))
		}

		// Check if the method is blocked
		if time.Now().Before(methodLimiter.blockedUntil) {
			time.Sleep(time.Until(methodLimiter.blockedUntil))
		}

		// Obtain a lock on the region and method limiters
		// Add the request to the limiter channels
		regionLimiter.shortLimiter.Obtain()
		regionLimiter.longLimiter.Obtain()
		methodLimiter.shortLimiter.Obtain()

		go func(req *APIRequest) {
			// Create a new HTTP request and set the API key as a header
			httpRequest, err := http.NewRequest("GET", req.URL, nil)
			if err != nil {
				return
			}
			httpRequest.Header.Set("X-Riot-Token", rl.apiKey)

			// Send the HTTP request
			resp, err := (*rl.httpClient).Do(httpRequest)
			if err == nil && resp.StatusCode == http.StatusOK {
				rl.updateRateLimits(resp, req.MethodId, regionLimiter, methodLimiter)
				req.Response <- resp
			} else if err == nil && resp.StatusCode == http.StatusForbidden {
				req.Response <- resp
			} else if err == nil && resp.StatusCode == http.StatusTooManyRequests {
				// Retry the request if Retries is less than maxRetries, or if maxRetries is -1. Otherwise, send the response to the channel
				if req.Retries < rl.maxRetries || rl.maxRetries == -1 {
					req.Retries++
					rl.Requests <- req
				} else {
					req.Response <- resp
				}

				handleRateLimitedResponse(resp, regionLimiter, methodLimiter)

				// Remove the request from the limiter channels
				regionLimiter.shortLimiter.Release()
				regionLimiter.longLimiter.Release()
				methodLimiter.shortLimiter.Release()
			} else {
				if !isBadRequest(resp) && (req.Retries < rl.maxRetries || rl.maxRetries == -1) {
					req.Retries++
					rl.Requests <- req
				} else {
					req.Response <- resp
				}

				time.Sleep(15 * time.Second)

				// Remove the request from the limiter channels
				regionLimiter.shortLimiter.Release()
				regionLimiter.longLimiter.Release()
				methodLimiter.shortLimiter.Release()
			}
		}(req)
	}
}

func (rl *RateLimiter) updateRateLimits(resp *http.Response, methodId MethodId, regionLimiter *RateLimit, methodLimiter *RateLimit) {
	appRateLimitHeader := resp.Header.Get("X-App-Rate-Limit")
	appRateLimitCountHeader := resp.Header.Get("X-App-Rate-Limit-Count")
	methodRateLimitHeader := resp.Header.Get("X-Method-Rate-Limit")
	methodRateLimitCountHeader := resp.Header.Get("X-Method-Rate-Limit-Count")

	if appRateLimitHeader != "" && appRateLimitCountHeader != "" {
		shortLimitInfo, longLimitInfo := getShortAndLongLimits(appRateLimitHeader)
		shortCountInfo, longCountInfo := getShortAndLongLimits(appRateLimitCountHeader)

		rl.updateRateLimit(methodId, shortLimitInfo, shortCountInfo, regionLimiter, regionLimiter.shortLimiter, &regionLimiter.blockedUntil, rl.conserveUsage.RegionPercent, true)
		rl.updateRateLimit(methodId, longLimitInfo, longCountInfo, regionLimiter, regionLimiter.longLimiter, &regionLimiter.blockedUntil, rl.conserveUsage.RegionPercent, true)
	} else {
		// Remove the request from the limiter channels
		go func() {
			time.Sleep(time.Duration(15) * time.Second)
			regionLimiter.shortLimiter.Release()
			regionLimiter.longLimiter.Release()
		}()
	}

	if methodRateLimitHeader != "" && methodRateLimitCountHeader != "" {
		rl.updateRateLimit(methodId, methodRateLimitHeader, methodRateLimitCountHeader, methodLimiter, methodLimiter.shortLimiter, &methodLimiter.blockedUntil, rl.conserveUsage.MethodPercent, false)
	} else {
		// Remove the request from the limiter channels
		go func() {
			time.Sleep(time.Duration(15) * time.Second)
			methodLimiter.shortLimiter.Release()
		}()
	}
}

func (rl *RateLimiter) updateRateLimit(methodId MethodId, limitInfo, countInfo string, limiter *RateLimit, limiterChannel *limiterMu, blockedUntil *time.Time, conservePercent int, isRegionHeader bool) {
	limitSplit := strings.Split(limitInfo, ":")
	countSplit := strings.Split(countInfo, ":")

	limit, _ := strconv.Atoi(limitSplit[0])
	limitTimeout, _ := strconv.Atoi(limitSplit[1])
	count, _ := strconv.Atoi(countSplit[0])

	var limitWithConservation int = limit

	var useConservation bool = false
	if conservePercent > 0 {
		useConservation = true

		if !isRegionHeader {
			for i := 0; i < len(rl.conserveUsage.IgnoreLimits); i++ {
				if rl.conserveUsage.IgnoreLimits[i] == methodId {
					useConservation = false
					break
				}
			}
		}
	}

	if useConservation {
		limitWithConservation = limit - (limit * conservePercent / 100)
	} else {
		limitWithConservation = limit - 1
	}

	// If the limit has been reached, block the limiter channel until the limit resets
	if count >= limitWithConservation && time.Now().After(*blockedUntil) {
		*blockedUntil = time.Now().Add(time.Duration(limitTimeout) * time.Second)
	}

	// Resize the limiter channel if needed
	if limiterChannel.Capacity() != limitWithConservation {
		limiterChannel.SetCapacity(limitWithConservation)
	}

	// Add a goroutine to remove an element from the limiter channel after the limit timeout
	go func() {
		time.Sleep(time.Duration(limitTimeout) * time.Second)
		limiterChannel.Release()
	}()
}

func isBadRequest(resp *http.Response) bool {
	return resp == nil || resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden ||
		resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusMethodNotAllowed || resp.StatusCode == http.StatusUnsupportedMediaType
}

func getShortAndLongLimits(limitHeader string) (string, string) {
	limits := strings.Split(limitHeader, ",")
	return limits[0], limits[1]
}

func handleRateLimitedResponse(resp *http.Response, regionLimiter *RateLimit, methodLimiter *RateLimit) {
	retryAfterHeader := resp.Header.Get("Retry-After")
	rateLimitTypeHeader := resp.Header.Get("X-Rate-Limit-Type")
	retryAfter, _ := strconv.Atoi(retryAfterHeader)
	retryAfterDuration := time.Duration(retryAfter) * time.Second

	if rateLimitTypeHeader == "application" {
		regionLimiter.blockedUntil = time.Now().Add(retryAfterDuration)
	} else if rateLimitTypeHeader == "method" {
		methodLimiter.blockedUntil = time.Now().Add(retryAfterDuration)
	}

	time.Sleep(retryAfterDuration)
}
