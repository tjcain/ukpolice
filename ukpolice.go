package ukpolice

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/time/rate"
)

const (
	// DefaultBaseURL is the URL of the data.police.uk API
	DefaultBaseURL = "https://data.police.uk/api/"
	// DefaultUserAgent is the value to use in the User-Agent header if none
	// has been explicitly configured.
	DefaultUserAgent = "go-ukpolice"
	// RequestLimit is set to the rate limit of the data.police.uk api
	RequestLimit = 15
	// BurstLimit is set to the single second burst limit of the
	// data.police.uk api.
	BurstLimit = 30
)

var limiter = rate.NewLimiter(RequestLimit, BurstLimit)

// for testing
var now = time.Now()

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

	common service // Reuse a single struct instead of allocating one for each service.

	// Services used for talking to different parts of the data.police.uk API
	Avaliability  *AvaliabilityService
	Force         *ForceService
	Crime         *CrimeService
	Neighborhood  *NeighbourhoodService
	StopAndSearch *StopAndSearchService
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
	api.Crime = (*CrimeService)(&api.common)
	api.Neighborhood = (*NeighbourhoodService)(&api.common)
	api.StopAndSearch = (*StopAndSearchService)(&api.common)

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

// Do carries out a request and stores the result in v.
func (api *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)

	err := limiter.Wait(ctx)
	if err != nil {
		return nil, err
	}

	// send request
	resp, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}

	// deferred closing of response body
	defer resp.Body.Close()

	response := makeResponse(resp)

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

// Date represents a date in the format YYYY-MM
type Date struct {
	Date string `json:"date,omitempty" url:"date"`
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
