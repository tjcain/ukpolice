package ukpolice

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

// Neighbourhoods
func TestNeighbourhoodService_GetNeighbourhoods(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/leicestershire/neighbourhoods", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
		[
			{
				"id": "NC04",
				"name": "City Centre"
			},
			{
				"id": "NC66",
				"name": "Cultural Quarter"
			},
			{
				"id": "NC67",
				"name": "Riverside"
			}
		]
		`)
	})

	neighbourhoods, _, err := client.Neighborhood.GetNeighbourhoods(context.Background(), "leicestershire")
	if err != nil {
		t.Errorf("Neighbourhood.GetNeighbourhoods returned error: '%+v'", err)
	}

	want := []Neighbourhood{
		Neighbourhood{ID: "NC04", Name: "City Centre"},
		Neighbourhood{ID: "NC66", Name: "Cultural Quarter"},
		Neighbourhood{ID: "NC67", Name: "Riverside"},
	}

	if !reflect.DeepEqual(neighbourhoods, want) {
		t.Errorf("Neighbourhood.GetNeighbourhoods returned %v, want %v", neighbourhoods, want)
	}
}

// Specific neighbourhood
func TestNeighbourhoodService_GetSpecificNeighbourhood(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/leicestershire/NC04", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
			{
				"url_force": "http://www.leics.police.uk/local-policing/city-centre",
				"contact_details": {
					"twitter": "http://www.twitter.com/centralleicsNPA",
					"facebook": "http://www.facebook.com/leicspolice",
					"telephone": "101",
					"email": "centralleicester.npa@leicestershire.pnn.police.uk"
				},
				"name": "City Centre",
				"links": [
					{
						"url": "http://www.leicester.gov.uk/",
						"description": null,
						"title": "Leicester City Council"
					}
				],
				"centre": {
					"latitude": "52.6389",
					"longitude": "-1.13619"
				},
				"locations": [
					{
						"name": "Mansfield House",
						"longitude": null,
						"postcode": "LE1 3GG",
						"address": "74 Belgrave Gate\n, Leicester",
						"latitude": null,
						"type": "station",
						"description": null
					}
				],
				"description": "<p>The Castle neighbourhood is a diverse covering all of the City Centre. In addition it covers De Montfort University, the University of Leicester, Leicester Royal Infirmary, the Leicester Tigers rugby ground and the Clarendon Park and Riverside communities.</p>\n<p>The Highcross and Haymarket shopping centres and Leicester's famous Market are all covered by this neighbourhood.</p>",
				"id": "NC04",
				"population": "0"
			}
		`)
	})

	neighbourhood, _, err := client.Neighborhood.GetSpecificNeighbourhood(context.Background(), "leicestershire", "NC04")
	if err != nil {
		t.Errorf("Neighbourhood.GetSpecificNeighbourhood returned error: '%+v'", err)
	}

	want := &Neighbourhood{
		ForceURL: "http://www.leics.police.uk/local-policing/city-centre",
		ContactDetails: map[string]string{
			"twitter":   "http://www.twitter.com/centralleicsNPA",
			"facebook":  "http://www.facebook.com/leicspolice",
			"telephone": "101",
			"email":     "centralleicester.npa@leicestershire.pnn.police.uk",
		},
		Name: "City Centre",
		Links: []map[string]string{
			{
				"url":         "http://www.leicester.gov.uk/",
				"description": "",
				"title":       "Leicester City Council",
			},
		},
		Centre: Location{
			Latitude:  "52.6389",
			Longitude: "-1.13619",
		},
		Locations: []Location{
			{
				Name:        "Mansfield House",
				Longitude:   "",
				Postcode:    "LE1 3GG",
				Address:     "74 Belgrave Gate\n, Leicester",
				Latitude:    "",
				Type:        "station",
				Description: "",
			},
		},
		Description: "<p>The Castle neighbourhood is a diverse covering all of the City Centre. In addition it covers De Montfort University, the University of Leicester, Leicester Royal Infirmary, the Leicester Tigers rugby ground and the Clarendon Park and Riverside communities.</p>\n<p>The Highcross and Haymarket shopping centres and Leicester's famous Market are all covered by this neighbourhood.</p>",
		ID:          "NC04",
		Population:  "0",
	}

	if !reflect.DeepEqual(neighbourhood, want) {
		t.Errorf("Neighbourhood.GetSpecificNeighbourhood returned %v, want %v", neighbourhood, want)
	}
}

// Neighbourhood boundary
func TestNeighbourhoodService_GetNeighbourhoodBoundary(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/leicestershire/NC04/boundary", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
		[
			{
				"latitude": "52.6394052587",
				"longitude": "-1.1458618876"
			},
			{
				"latitude": "52.6389452755",
				"longitude": "-1.1457057759"
			},
			{
				"latitude": "52.6383706746",
				"longitude": "-1.1455755443"
			}
		]
	`)
	})

	neighbourhood, _, err := client.Neighborhood.GetNeighbourhoodBoundary(context.Background(), "leicestershire", "NC04")
	if err != nil {
		t.Errorf("Neighbourhood.GetNeighbourhoodBoundary returned error: '%+v'", err)
	}

	want := []Location{
		{
			Latitude:  "52.6394052587",
			Longitude: "-1.1458618876",
		},
		{
			Latitude:  "52.6389452755",
			Longitude: "-1.1457057759",
		},
		{
			Latitude:  "52.6383706746",
			Longitude: "-1.1455755443",
		},
	}

	if !reflect.DeepEqual(neighbourhood, want) {
		t.Errorf("Neighbourhood.GetNeighbourhoodBoundary returned %v, want %v", neighbourhood, want)
	}
}

// Neighbourhood team
func TestNeighbourhoodService_GetNeighbourhoodTeam(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/leicestershire/NC04/people", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
			[
				{
					"bio": "<p>I joined Leicestershire Police in 1997 and have enjoyed a variety of roles across the County.<br />\n</p>\n<p>I have worked as a Sergeant in the City Centre for the last 6 years in Response and Neighbourhood Policing roles.</p>\n<p>I am now Deputy NPA Commander and am privileged to work with a great team of committed and enthusiastic people.</p>\n<p>We always try and listen to concerns then support and protect the City communities with a problem solving approach.  Please don?t hesitate to contact us with any issues or problems you may have.</p>",
					"contact_details": {},
					"name": "Andy Cooper",
					"rank": "Sgt"
				},
				{
					"bio":"<p>I am the Safer Neighbourhood Sergeant for the Cultural Quarter and recently joined the Neighbourhood Team at Mansfield House.</p>\n<p>I have been a police officer for 27 years and was promoted to Sergeant 14 years ago. Since then I have worked in a variety of roles in both the County and City my last role being at Euston Street managing resources and demand for the Response units based there. </p>\n<p>I am looking forward to policing the City Centre and will be responsible, along with my team for the Cultural Quarter so please contact myself with any issues that I can help with.<br />\n</p>",
					"contact_details": {},
					"name": "Andy Price",
					"rank": "Sgt"
				}
			]
	`)
	})

	team, _, err := client.Neighborhood.GetNeighbourhoodTeam(context.Background(), "leicestershire", "NC04")
	if err != nil {
		t.Errorf("Neighbourhood.GetNeighbourhoodTeam returned error: '%+v'", err)
	}

	want := []NeighbourhoodTeam{
		{
			Bio:            "<p>I joined Leicestershire Police in 1997 and have enjoyed a variety of roles across the County.<br />\n</p>\n<p>I have worked as a Sergeant in the City Centre for the last 6 years in Response and Neighbourhood Policing roles.</p>\n<p>I am now Deputy NPA Commander and am privileged to work with a great team of committed and enthusiastic people.</p>\n<p>We always try and listen to concerns then support and protect the City communities with a problem solving approach.  Please don?t hesitate to contact us with any issues or problems you may have.</p>",
			ContactDetails: map[string]string{},
			Name:           "Andy Cooper",
			Rank:           "Sgt",
		},
		{
			Bio:            "<p>I am the Safer Neighbourhood Sergeant for the Cultural Quarter and recently joined the Neighbourhood Team at Mansfield House.</p>\n<p>I have been a police officer for 27 years and was promoted to Sergeant 14 years ago. Since then I have worked in a variety of roles in both the County and City my last role being at Euston Street managing resources and demand for the Response units based there. </p>\n<p>I am looking forward to policing the City Centre and will be responsible, along with my team for the Cultural Quarter so please contact myself with any issues that I can help with.<br />\n</p>",
			ContactDetails: map[string]string{},
			Name:           "Andy Price",
			Rank:           "Sgt",
		},
	}

	if !reflect.DeepEqual(team, want) {
		t.Errorf("Neighbourhood.GetNeighbourhoodTeam returned %v, want %v", team, want)
	}
}

// Neighbourhood events
func TestNeighbourhoodService_GetNeighbourhoodEvents(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/leicestershire/NC04/events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
			[
				{
					"contact_details": {},
					"description": null,
					"title": "Drop In Beat Surgery",
					"address": "Nagarjuna Buddhist Centre, 17 Guildhall Lane",
					"type": "meeting",
					"start_date": "2016-09-17T12:00:00",
					"end_date": "2016-09-17T14:00:00"
				}
			]
	`)
	})

	events, _, err := client.Neighborhood.GetNeighbourhoodEvents(context.Background(), "leicestershire", "NC04")
	if err != nil {
		t.Errorf("Neighbourhood.GetNeighbourhoodEvents returned error: '%+v'", err)
	}

	want := []NeighbourhoodEvent{
		{
			ContactDetails: map[string]string{},
			Description:    "",
			Title:          "Drop In Beat Surgery",
			Address:        "Nagarjuna Buddhist Centre, 17 Guildhall Lane",
			Type:           "meeting",
			StartDate:      "2016-09-17T12:00:00",
			EndDate:        "2016-09-17T14:00:00",
		},
	}

	if !reflect.DeepEqual(events, want) {
		t.Errorf("Neighbourhood.GetNeighbourhoodEvents returned %v, want %v", events, want)
	}
}

// Neighbourhood priorities
func TestNeighbourhoodService_GetNeighbourhoodPriorities(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/leicestershire/NC04/priorities", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
			[
				{
					"action": null,
					"issue-date": "2016-04-14T00:00:00",
					"action-date": null,
					"issue": "<p>To reduce the amount of Anti-Social Behaviour Humberstone Gate, Leicester.</p>"
				}
			]
	`)
	})

	events, _, err := client.Neighborhood.GetNeighbourhoodPriorities(context.Background(), "leicestershire", "NC04")
	if err != nil {
		t.Errorf("Neighbourhood.GetNeighbourhoodPriorities returned error: '%+v'", err)
	}

	want := []NeighbourhoodPriorities{
		{
			Action:     "",
			Issue:      "<p>To reduce the amount of Anti-Social Behaviour Humberstone Gate, Leicester.</p>",
			IssueDate:  "2016-04-14T00:00:00",
			ActionDate: "",
		},
	}

	if !reflect.DeepEqual(events, want) {
		t.Errorf("Neighbourhood.GetNeighbourhoodPriorities returned %v, want %v", events, want)
	}
}

// Locate neighbourhood
func TestNeighbourhoodService_LocateNeighbourhood(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/locate-neighbourhood", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
			{
				"force": "metropolitan",
				"neighbourhood": "00BKX6"
			}
	`)
	})

	neighbourhood, _, err := client.Neighborhood.LocateNeighbourhood(context.Background(), "51.500617", "-0.124629")
	if err != nil {
		t.Errorf("Neighbourhood.LocateNeighbourhood returned error: '%+v'", err)
	}

	want := &Neighbourhood{Force: "metropolitan", Neighbourhood: "00BKX6"}

	if !reflect.DeepEqual(neighbourhood, want) {
		t.Errorf("Neighbourhood.LocateNeighbourhood returned %v, want %v", neighbourhood, want)
	}
}
