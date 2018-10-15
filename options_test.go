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
		addOptions("base-query", WithLatLong("52.629729", "-1.131592"),
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
		{"LatLong", []Option{WithLatLong("52.629729", "-1.131592")}, "base-query?lat=52.629729&lng=-1.131592"},
		{"Polygon", []Option{WithPolygon("52.268,0.543:52.794,0.238:52.130,0.478")}, "base-query?poly=52.268%2C0.543%3A52.794%2C0.238%3A52.130%2C0.478"},
		{"ID", []Option{WithLocationID("884227")}, "base-query?location_id=884227"},
		{"Date", []Option{WithDate("2017-02"), WithLocationID("884227")}, "base-query?date=2017-02&location_id=884227"},
		{"Crime Category", []Option{WithCrimeCategory("all-crime")}, "base-query?category=all-crime"},
		{"Force", []Option{WithForce("west-midlands")}, "base-query?force=west-midlands"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := addOptions("base-query", tc.input...)
			if s != tc.output {
				t.Errorf("output for %v should be %v; got %v",
					tc.name,
					tc.output,
					s)
			}
		})
	}
}
