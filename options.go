package ukpolice

import (
	"net/url"
)

// Option specifies parameters to various methods that support multiple variable
// choices.
type Option func(*url.Values)

// WithDate adds date information to methods which accept a date
// parameter.
func WithDate(date string) Option {
	return func(v *url.Values) {
		v.Set("date", string(date))
	}
}

// WithLatLong will add provided latitude and longitude information to methods
// which accept a latitude and longitude.
// Lat Long is exclusive and may not be combied with any other location variable
// in a single request, doing so will result in a panic.
func WithLatLong(latitude, longitude string) Option {
	return func(v *url.Values) {
		if v.Get("poly") != "" || v.Get("location_id") != "" {
			panic("oops")
		}
		v.Set("lat", string(latitude))
		v.Set("lng", longitude)
	}
}

// WithPolygon will add provided latitude and longitude information to methods
// which accept a latitude and longitude.
// WithPolygon is exclusive and may not be combied with any other location variable
// in a single request, doing so will result in a panic.
func WithPolygon(poly string) Option {
	return func(v *url.Values) {
		if v.Get("lat") != "" || v.Get("lng") != "" || v.Get("location_id") != "" {
			panic("oops")
		}
		v.Set("poly", poly)
	}
}

// WithLocationID will add provided latitude and longitude information to methods
// which accept a latitude and longitude.
// WithLocationID is exclusive and may not be combied with any other location variable
// in a single request, doing so will result in a panic.
func WithLocationID(id string) Option {
	return func(v *url.Values) {
		if v.Get("lat") != "" || v.Get("lng") != "" || v.Get("poly") != "" {
			panic("oops")
		}
		v.Set("location_id", id)
	}
}

func addOptions(baseURL string, opts ...Option) string {
	u, _ := url.Parse(baseURL)
	q := u.Query()
	for _, opt := range opts {
		opt(&q)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
