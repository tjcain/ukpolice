package ukpolice

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAvailabilityService_GetAvailabilityInfo(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crimes-street-dates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
			[
				{
					"date": "2015-06",
					"stop-and-search": ["bedfordshire","cleveland","durham"]
				},
				{
					"date": "2015-05",
					"stop-and-search": ["bedfordshire","city-of-london","cleveland"]
				}
			]
		`)
	})

	availabilityInfo, _, err := client.Availability.GetAvailabilityInfo(context.Background())
	if err != nil {
		t.Errorf("Availability.GetAvailabilityInfo returned error: '%+v'", err)
	}

	want := []AvailabilityInfo{
		{"2015-06", []string{"bedfordshire", "cleveland", "durham"}},
		{"2015-05", []string{"bedfordshire", "city-of-london", "cleveland"}},
	}

	if !reflect.DeepEqual(availabilityInfo, want) {
		t.Errorf("Availability.GetAvailabilityInfo returned %v, want %v", availabilityInfo, want)
	}

}
