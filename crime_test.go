package ukpolice

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var (
	rawCrime = `
		[
			{
				"category": "anti-social-behaviour",
				"location_type": "Force",
				"location": {
					"latitude": "52.640961",
					"longitude": "-1.126371",
					"street": {
						"id": 884343,
						"name": "On or near Wharf Street North"
					}
				},
				"context": "",
				"outcome_status": null,
				"persistent_id": "",
				"id": 54164419,
				"location_subtype": "",
				"month": "2017-01"
			}
		]`

	expectedCrime = Crime{
		Category:     "anti-social-behaviour",
		LocationType: "Force",
		Location: Location{
			Latitude: "52.640961",
			Street: Street{
				ID:   884343,
				Name: "On or near Wharf Street North",
			},
			Longitude: "-1.126371",
		},
		Context:         "",
		OutcomeStatus:   nil,
		PersistentID:    "",
		ID:              54164419,
		LocationSubtype: "",
		Month:           "2017-01",
	}
)

// Street level crimes
func TestCrimeService_GetStreetLevelCrimes(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crimes-street/all-crime", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, rawCrime)
	})

	crimes, _, err := client.Crime.GetStreetLevelCrimes(context.Background(),
		WithLatLong("52.629729", "-1.131592"), WithDate("2017-01"))
	if err != nil {
		t.Errorf("Crime.GetStreetLevelCrimes returned error: '%s'", err)
	}

	want := []Crime{expectedCrime}

	if !reflect.DeepEqual(crimes, want) {
		t.Errorf("Crime.GetForceDetails returned %v, want %v", crimes, want)
	}
}

// Street level outcomes
func TestCrimeService_GetStreetLevelOutcomes(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/outcomes-at-location", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
			[
				{
					"category": {
						"code": "unable-to-prosecute",
						"name": "Unable to prosecute suspect"
					},
					"date": "2017-01",
					"person_id": null,
					"crime": {
						"category": "theft-from-the-person",
						"location_type": "Force",
						"location": {
							"latitude": "52.634474",
							"street": {
								"id": 883498,
								"name": "On or near Kate Street"
							},
							"longitude": "-1.149197"
						},
						"context": "",
						"persistent_id": "a5a98275facee535b959b236130f5ec05205869fb3d0804c9b14296fcd0bce46",
						"id": 53566126,
						"location_subtype": "ROAD",
						"month": "2016-12"
					}
				}
			]
		`)
	})

	outcomes, _, err := client.Crime.GetStreetLevelOutcomes(context.Background(),
		WithLatLong("52.629729", "-1.131592"), WithDate("2017-01"))
	if err != nil {
		t.Errorf("Crime.GetStreetLevelCrimes returned error: '%s'", err)
	}

	want := []Outcome{
		{
			Category: map[string]string{
				"code": "unable-to-prosecute",
				"name": "Unable to prosecute suspect",
			},
			Date:     "2017-01",
			PersonID: 0,
			Crime: Crime{
				Category:     "theft-from-the-person",
				LocationType: "Force",
				Location: Location{
					Latitude:  "52.634474",
					Longitude: "-1.149197",
					Street: Street{
						ID:   883498,
						Name: "On or near Kate Street",
					},
				},
				Context:         "",
				PersistentID:    "a5a98275facee535b959b236130f5ec05205869fb3d0804c9b14296fcd0bce46",
				ID:              53566126,
				LocationSubtype: "ROAD",
				Month:           "2016-12",
			},
		},
	}

	if !reflect.DeepEqual(outcomes, want) {
		t.Errorf("Crime.GetStreetLevelOutcomes returned %v, want %v", outcomes, want)
	}
}

// Crimes at location
func TestCrimeService_GetCrimesAtLocation(t *testing.T) {

}

// Crimes with no location
func TestCrimeService_GetCrimesWithNoLocation(t *testing.T) {

}

// Crime categories
func TestCrimeService_GetCrimeCatagories(t *testing.T) {

}

// Last updated
func TestCrimeService_GetLastUpdated(t *testing.T) {

}

// Outcomes for a specific crime
func TestCrimeService_GetOutcomesForSpecificCrime(t *testing.T) {

}
