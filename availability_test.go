package ukpolice

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAvaliabilityService_GetAvaliabilityInfo(t *testing.T) {
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

	avaliabilityInfo, _, err := client.Avaliability.GetAvaliabilityInfo(context.Background())
	if err != nil {
		t.Errorf("Avaliability.GetAvaliabilityInfo returned error: '%+v'", err)
	}

	want := []AvaliabilityInfo{
		{"2015-06", []string{"bedfordshire", "cleveland", "durham"}},
		{"2015-05", []string{"bedfordshire", "city-of-london", "cleveland"}},
	}

	if !reflect.DeepEqual(avaliabilityInfo, want) {
		t.Errorf("Avaliability.GetAvaliabilityInfo returned %v, want %v", avaliabilityInfo, want)
	}

}
