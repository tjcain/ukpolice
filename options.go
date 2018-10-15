package ukpolice

import (
	"net/url"
)

// Option specifies parameters to various methods that support multiple variable
// choices.
type Option func(*url.Values)

// WithDate sets provided date URL parameters.
func WithDate(date string) Option {
	return func(v *url.Values) {
		v.Set("date", string(date))
	}
}

// WithLatLong sets provided latitude and longitude URL parameters.
func WithLatLong(latitude, longitude string) Option {
	return func(v *url.Values) {
		if v.Get("poly") != "" || v.Get("location_id") != "" {
			panic("oops")
		}
		v.Set("lat", string(latitude))
		v.Set("lng", longitude)
	}
}

// WithPolygon sets provided polygon URL parameters.
func WithPolygon(poly string) Option {
	return func(v *url.Values) {
		if v.Get("lat") != "" || v.Get("lng") != "" || v.Get("location_id") != "" {
			panic("oops")
		}
		v.Set("poly", poly)
	}
}

// WithLocationID sets provided locationID URL parameters.
func WithLocationID(id string) Option {
	return func(v *url.Values) {
		if v.Get("lat") != "" || v.Get("lng") != "" || v.Get("poly") != "" {
			panic("oops")
		}
		v.Set("location_id", id)
	}
}

// WithCrimeCategory sets provided crime category URL parameters.
func WithCrimeCategory(category string) Option {
	return func(v *url.Values) {
		v.Set("category", category)
	}
}

// WithForce sets provided force URL parameters.
func WithForce(force string) Option {
	return func(v *url.Values) {
		v.Set("force", force)
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
