package ukpolice

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

const (
	// DefaultBaseURL is the URL of the data.police.uk API
	DefaultBaseURL = "https://data.police.uk/api/"
	// DefaultUserAgent is the value to use in the User-Agent header if none
	// has been explicitly configured.
	DefaultUserAgent = "go-ukpolice"

	// I am not sure if the API returns these headers - something to test.
	headerErrorRateRemaining = "X-ESI-Error-Limit-Remain"
	headerErrorRateReset     = "X-ESI-Error-Limit-Reset"
)

// for testing?
var now = time.Now

type service struct {
	api *Client
}

// Client manages communication with the data.police.uk API
type Client struct {
	client *http.Client // HTTP client used to communicate with the API

	// BaseURL for API requests
	BaseURL *url.URL

	// UserAgent for communicating with the data.police.uk API
	UserAgent string

	mu struct {
		sync.Mutex
		Rate
	}

	// Logging holds optional loggers. If any are nil, logging is done via the
	// log package's standard logger.
	Logging struct {
		Info, Error, Debug *log.Logger
	}

	common service // Reuse a single struct instead of allocating one for each service.

	// Services used for talking to different parts of the data.police.uk API
	Avaliability *AvaliabilityService
	// @TODO: Force Related
	Force *ForceService
	// @TODO: Crime Related
	// @TODO: Neighborhood Related
	// @TODO: Stop and Search Related

}

// NewClient returns a new data.police.uk API client. If a nil httpClient is
// provided, http.DefaultClient will be used.
// There is no authentication required by the API.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(DefaultBaseURL)

	api := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: DefaultUserAgent,
	}

	api.common.api = api

	// endpoints
	api.Avaliability = (*AvaliabilityService)(&api.common)
	api.Force = (*ForceService)(&api.common)

	return api
}

// NewRequest creates an API request. An url relative to the BaseURL of the
// client is provided. No request body is required for interaction with this
// API
func (api *Client) NewRequest(method, url string, body interface{}) (*http.Request, error) {
	u, err := api.BaseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	// set header first.
	req.Header.Set("Accept", "application/json")

	if api.UserAgent != "" {
		req.Header.Set("User-Agent", api.UserAgent)
	}

	return req, nil
}

// Response is a data.police.uk API response. This wraps the standard
// http.Response returned from data.police.uk.
type Response struct {
	*http.Response
}

func makeResponse(r *http.Response) *Response {
	return &Response{Response: r}
}

// Error represents a data.police.uk API error.
// I'm not sure if the api ever returns a json error?
type Error struct {
	Response       *http.Response
	HTTPStatusCode int
	Err            string `json:"error"`

	Rate
}

// Implement Error interface
func (e Error) Error() string {
	return e.Err
}

func makeError(r *http.Response) *Error {
	e := &Error{Response: r}

	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, e)
	}

	e.HTTPStatusCode = r.StatusCode
	e.Rate = parseRate(r)

	return e
}

// Rate represents rate limit information.
type Rate struct {
	Remaining int
	Reset     time.Time
}

// Implement Stringer interface
func (r Rate) String() string {
	return fmt.Sprintf("error rate limit: %d remaining calls; reset in %.fs",
		r.Remaining, r.Reset.Sub(now()).Seconds())
}

func parseRate(r *http.Response) Rate {
	var rate Rate
	if remaining := r.Header.Get(headerErrorRateRemaining); remaining != "" {
		rate.Remaining, _ = strconv.Atoi(remaining)
	}

	if reset := r.Header.Get(headerErrorRateReset); reset != "" {
		if v, _ := strconv.Atoi(reset); v != 0 {
			rate.Reset = now().Add(time.Duration(v) * time.Second)
		}
	}

	return rate
}

// Do carries out a request and stores the result in v.
func (api *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)
	// send request
	resp, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}
	// deferred closing of response body
	defer resp.Body.Close()

	response := makeResponse(resp)

	if err := api.check(resp); err != nil {
		api.mu.Lock()
		api.mu.Rate = err.(*Error).Rate
		api.mu.Unlock()
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
				if err == io.EOF {
					err = nil
				}
				return response, err
			}
		}
	}

	return response, nil
}

func (api *Client) check(resp *http.Response) error {
	if rc := resp.StatusCode; 200 <= rc && rc <= 299 {
		// check for any waning headers and log them
		if v := resp.Header.Get("warning"); v != "" {
			logf(api.Logging.Error, "warning header received (%s %v): %s",
				resp.Request.Method, resp.Request.URL.Path, v,
			)
		}

		return nil
	}

	return makeError(resp)
}

func logf(logger *log.Logger, format string, args ...interface{}) {
	if logger != nil {
		logger.Printf(format, args...)
		return
	}

	log.Printf(format, args...)
}

// Bool is a helper function that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper function that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper function that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// Float64 is a helper function that allocates a new float64 value
// to store v and returns a pointer to it.
func Float64(v float64) *float64 { return &v }

// String is a helper function that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
