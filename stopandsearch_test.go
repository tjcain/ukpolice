package ukpolice

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var rawSearch = `
	[
		{
			"age_range": "10-17",
			"self_defined_ethnicity": "White - White British (W1)",
			"outcome_linked_to_object_of_search": true,
			"datetime": "2017-01-14T20:50:00+00:00",
			"removal_of_more_than_outer_clothing": null,
			"operation": null,
			"officer_defined_ethnicity": "White",
			"object_of_search": "Controlled drugs",
			"involved_person": true,
			"gender": "Female",
			"legislation": "Misuse of Drugs Act 1971 (section 23)",
			"location": {
				"latitude": "52.634407",
				"street": {
					"id": 883407,
					"name": "On or near Shopping Area"
				},
				"longitude": "-1.133653"
			},
			"outcome": "Local resolution",
			"type": "Person search",
			"operation_name": null
		}
	]
`

var expectedSearch = []Search{
	// AgeRange:                       "10-17",
	// SelfDefinedEthnicity:           "White - White British (W1)",
	// OutcomeLinkedToObject:          true,
	// DateTime:                       "2017-01-14T20:50:00+00:00",
	// RemovalOfMoreThanOuterClothing: false,
	// Operation:                      false,
	// OfficerDefinedEthnicity:        "White",
	// ObjectOfSearch:                 "Controlled drugs",
	// InvolvedPerson:                 true,
	// Gender:                         "Female",
	// Legislation:                    "Misuse of Drugs Act 1971 (section 23)",
}

// Stop and searches by area
func TestCrimeService_GetStopAndSearchesByArea(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/stops-street", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, rawSearch)
	})

	searches, _, err := client.StopAndSearch.GetStopAndSearchesByArea(context.Background(),
		WithDate("2017-01"), WithLatLong("52.629729", "-1.131592"))
	if err != nil {
		t.Errorf("StopAndSearch.GetStopAndSearchesByArea returned error: '%s'", err)
	}

	want := []Search{}
	err = json.Unmarshal([]byte(rawSearch), &want)
	if err != nil {
		t.Errorf("could not unmarshal json: '%s'", err)
	}

	if !reflect.DeepEqual(searches, want) {
		t.Errorf("StopAndSearch.GetStopAndSearchesByArea returned %v, want %v", searches, want)
	}
}

// Stop and searches by location
func TestCrimeService_GetStopAndSearchesByLocation(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/stops-at-location", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, rawSearch)
	})

	searches, _, err := client.StopAndSearch.GetStopAndSearchesByLocation(context.Background(),
		WithDate("2017-01"), WithLocationID("883407"))
	if err != nil {
		t.Errorf("StopAndSearch.GetStopAndSearchesByLocation returned error: '%s'", err)
	}

	want := []Search{}
	err = json.Unmarshal([]byte(rawSearch), &want)
	if err != nil {
		t.Errorf("could not unmarshal json: '%s'", err)
	}

	if !reflect.DeepEqual(searches, want) {
		t.Errorf("StopAndSearch.GetStopAndSearchesByLocation returned %v, want %v", searches, want)
	}
}

// Stop and searches with no location
// func TestCrimeService_GetStopAndSearchesWithNoLocation(t *testing.T) {
// 	client, mux, _, teardown := setup()
// 	defer teardown()

// 	mux.HandleFunc("/stops-no-location", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		fmt.Fprint(w, `
// 			[
// 				{
// 					"id": 0,
// 					"age_range": "over 34",
// 					"self_defined_ethnicity": "White - White British (W1)",
// 					"outcome_linked_to_object_of_search": null,
// 					"datetime": "2017-01-24T01:50:00+00:00",
// 					"removal_of_more_than_outer_clothing": null,
// 					"operation": null,
// 					"officer_defined_ethnicity": "White",
// 					"object_of_search": "Controlled drugs",
// 					"involved_person": true,
// 					"gender": "Male",
// 					"legislation": "Misuse of Drugs Act 1971 (section 23)",
// 					"location": null,
// 					"outcome": false,
// 					"type": "Person search",
// 					"operation_name": null
// 				}
// 			]
// 		`)
// 	})

// 	searches, _, err := client.StopAndSearch.GetStopAndSearchesWithNoLocation(context.Background(),
// 		WithDate("2017-01"), WithForce("cleveland"))
// 	if err != nil {
// 		t.Errorf("StopAndSearch.GetStopAndSearchesWithNoLocation returned error: '%s'", err)
// 	}

// 	want := []Search{}
// 	err = json.Unmarshal([]byte(rawSearch), &want)
// 	if err != nil {
// 		t.Errorf("could not unmarshal json: '%s'", err)
// 	}

// 	if !reflect.DeepEqual(searches, want) {
// 		t.Errorf("StopAndSearch.GetStopAndSearchesWithNoLocation returned %v, want %v", searches, want)
// 	}
// }

// Stop and searches by force
func TestCrimeService_GetStopAndSearchesByForce(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/stops-force", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, rawSearch)
	})

	searches, _, err := client.StopAndSearch.GetStopAndSearchesByForce(context.Background(),
		WithDate("2017-01"), WithForce("avon-and-summerset"))
	if err != nil {
		t.Errorf("StopAndSearch.GetStopAndSearchesByForce returned error: '%s'", err)
	}

	want := []Search{}
	err = json.Unmarshal([]byte(rawSearch), &want)
	if err != nil {
		t.Errorf("could not unmarshal json: '%s'", err)
	}

	if !reflect.DeepEqual(searches, want) {
		t.Errorf("StopAndSearch.GetStopAndSearchesByForce returned %v, want %v", searches, want)
	}
}
