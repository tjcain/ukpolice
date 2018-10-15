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
			Category: OutcomeCategory{
				Code: "unable-to-prosecute",
				Name: "Unable to prosecute suspect",
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
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crimes-at-location", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, rawCrime)
	})

	crimes, _, err := client.Crime.GetCrimesAtLocation(context.Background(),
		WithLatLong("52.629729", "-1.131592"), WithDate("2017-01"))
	if err != nil {
		t.Errorf("Crime.GetCrimesAtLocation returned error: '%s'", err)
	}

	want := []Crime{expectedCrime}

	if !reflect.DeepEqual(crimes, want) {
		t.Errorf("Crime.GetCrimesAtLocation returned %v, want %v", crimes, want)
	}
}

// Crimes with no location
func TestCrimeService_GetCrimesWithNoLocation(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crimes-no-location", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, rawCrime)
	})

	crimes, _, err := client.Crime.GetCrimesWithNoLocation(context.Background(),
		WithCrimeCategory("all-crime"), WithForce("staffordshire"))
	if err != nil {
		t.Errorf("Crime.GetCrimesWithNoLocation returned error: '%s'", err)
	}

	want := []Crime{expectedCrime}

	if !reflect.DeepEqual(crimes, want) {
		t.Errorf("Crime.GetCrimesWithNoLocation returned %v, want %v", crimes, want)
	}
}

// Crime categories
func TestCrimeService_GetCrimeCategories(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crime-categories", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, "[{\"url\":\"all-crime\", \"name\":\"All crime and ASB\"}]")
	})

	categories, _, err := client.Crime.GetCrimeCategories(context.Background(),
		WithDate("2011-08"))
	if err != nil {
		t.Errorf("Crime.GetCrimeCategories returned error: '%s'", err)
	}

	want := []CrimeCategory{{URL: "all-crime", Name: "All crime and ASB"}}
	if !reflect.DeepEqual(categories, want) {
		t.Errorf("Crime.GetCrimeCategories returned %v, want %v", categories, want)
	}

}

// Last updated
func TestCrimeService_GetLastUpdated(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crime-last-updated", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, "{\"date\": \"2018-08-01\"}")
	})

	date, _, err := client.Crime.GetLastUpdated(context.Background())
	if err != nil {
		t.Errorf("Crime.GetLastUpdated returned error: '%s'", err)
	}

	want := &Date{Date: "2018-08-01"}
	if !reflect.DeepEqual(date, want) {
		t.Errorf("Crime.GetLastUpdated returned %v, want %v", date, want)
	}
}

// Outcomes for a specific crime
func TestCrimeService_GetOutcomesForSpecificCrime(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/outcomes-for-crime/590d68b69228a9ff95b675bb4af591b38de561aa03129dc09a03ef34f537588c",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `
			{
				"crime": {
					"category": "violent-crime", 
					"persistent_id": "590d68b69228a9ff95b675bb4af591b38de561aa03129dc09a03ef34f537588c", 
					"location_subtype": "", 
					"location_type": "Force", 
					"location": {
						"latitude": "52.639814", 
						"street": {
							"id": 883235, 
							"name": "On or near Sanvey Gate"
						}, 
						"longitude": "-1.139118"
					}, 
					"context": "", 
					"month": "2017-05", 
					"id": 56880258
				}, 
				"outcomes": [
					{
						"category": {
							"code": "under-investigation", 
							"name": "Under investigation"
						}, 
						"date": "2017-05", 
						"person_id": null
					}, 
					{
						"category": {
							"code": "formal-action-not-in-public-interest", 
							"name": "Formal action is not in the public interest"
						}, 
						"date": "2017-06", 
						"person_id": null
					}
				]
			}
		`)
		})

	outcomes, _, err := client.Crime.GetSpecificOutcomes(context.Background(),
		"590d68b69228a9ff95b675bb4af591b38de561aa03129dc09a03ef34f537588c")
	if err != nil {
		t.Errorf("Crime.GetSpecificOutcomes returned error: '%s'", err)
	}

	want := &OutcomesForSpecificCrime{
		Crime: Crime{
			Category:        "violent-crime",
			PersistentID:    "590d68b69228a9ff95b675bb4af591b38de561aa03129dc09a03ef34f537588c",
			LocationSubtype: "",
			LocationType:    "Force",
			Location: Location{
				Latitude:  "52.639814",
				Longitude: "-1.139118",
				Street: Street{
					ID:   883235,
					Name: "On or near Sanvey Gate",
				},
			},
			Context: "",
			Month:   "2017-05",
			ID:      56880258,
		},
		Outcomes: []Outcome{
			{
				Category: OutcomeCategory{
					Code: "under-investigation",
					Name: "Under investigation",
				},
				Date:     "2017-05",
				PersonID: 0,
			},
			{
				Category: OutcomeCategory{
					Code: "formal-action-not-in-public-interest",
					Name: "Formal action is not in the public interest",
				},
				Date:     "2017-06",
				PersonID: 0,
			},
		},
	}
	if !reflect.DeepEqual(outcomes, want) {
		t.Errorf("Crime.GetSpecificOutcomes returned %v, want %v", outcomes, want)
	}

}
