package ukpolice

import (
	"net/url"
)

// Option specifies parameters to various methods that support multiple variable
// choices.
type Option func(*url.URL)

// WithDate adds date information to methods which accept a date
// parameter.
func WithDate(date string) Option {
	return func(u *url.URL) {
		q := u.Query()
		q.Set("date", string(date))
		u.RawQuery = q.Encode()
	}
}

// WithLatLong will add provided latitude and longitude information to methods
// which accept a latitude and longitude.
// Lat Long is exclusive and may not be combied with any other location variable
// in a single request, doing so will result in a panic.
func WithLatLong(latitude, longitude string) Option {
	return func(u *url.URL) {
		q := u.Query()
		q.Set("lat", string(latitude))
		q.Set("lng", longitude)
		u.RawQuery = q.Encode()
	}
}

// WithPolygon will add provided latitude and longitude information to methods
// which accept a latitude and longitude.
// WithPolygon is exclusive and may not be combied with any other location variable
// in a single request, doing so will result in a panic.
func WithPolygon(poly string) Option {
	return func(u *url.URL) {
		q := u.Query()
		q.Set("poly", poly)
		u.RawQuery = q.Encode()
	}
}

// WithLocationID will add provided latitude and longitude information to methods
// which accept a latitude and longitude.
// WithLocationID is exclusive and may not be combied with any other location variable
// in a single request, doing so will result in a panic.
func WithLocationID(id string) Option {
	return func(u *url.URL) {
		q := u.Query()
		q.Set("location_id", string(id))
		u.RawQuery = q.Encode()
		// u.Path = u.Path + fmt.Sprintf("location_id=%d&", id)
	}
}

func addOptions(baseURL string, opts ...Option) string {
	u, _ := url.Parse(baseURL)

	for _, opt := range opts {
		opt(u)
	}

	// check only one location method has been passed
	q := u.Query()
	d := q.Get("date")
	if len(opts) > 2 {
		panic("Usage: Must contain only one location method + optional date")
	}
	if len(opts) == 2 && d == "" {
		panic("Usage: Must contain only one location method + optional date")
	}

	return u.String()
}
