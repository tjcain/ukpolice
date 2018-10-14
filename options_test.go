package ukpolice

import (
	"testing"
)

func TestWrongOptionsPanic(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("addOptions should have panicked!")
			}
		}()
		// This function should cause a panic
		addOptions("crimes-at-location", WithLatLong("52.629729", "-1.131592"),
			WithPolygon("52.268,0.543:52.794,0.238:52.130,0.478"),
			WithDate("2017-02"))
	}()
}

func TestOptions(t *testing.T) {
	tt := []struct {
		name   string
		input  []Option
		output string
	}{
		{"LatLong", []Option{WithLatLong("52.629729", "-1.131592")}, "crimes-at-location?lat=52.629729&lng=-1.131592"},
		{"Polygon", []Option{WithPolygon("52.268,0.543:52.794,0.238:52.130,0.478")}, "crimes-at-location?poly=52.268%2C0.543%3A52.794%2C0.238%3A52.130%2C0.478"},
		{"ID", []Option{WithLocationID("884227")}, "crimes-at-location?location_id=884227"},
		{"With Date", []Option{WithDate("2017-02"), WithLocationID("884227")}, "crimes-at-location?date=2017-02&location_id=884227"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := addOptions("crimes-at-location", tc.input...)
			if s != tc.output {
				t.Errorf("output for %v should be %v; got %v",
					tc.name,
					tc.output,
					s)
			}
		})
	}
}
