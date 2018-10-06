package ukpolice

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestForceService_GetForces(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/forces", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
			[
				{
					"id": "avon-and-somerset",
					"name": "Avon and Somerset Constabulary"
				},
				{
					"id": "bedfordshire",
					"name": "Bedfordshire Police"
				},
				{
					"id": "cambridgeshire",
					"name": "Cambridgeshire Constabulary"
				}
			]
		`)
	})

	forces, _, err := client.Force.GetForces(context.Background())
	if err != nil {
		t.Errorf("Force.GetForces returned error: '%+v'", err)
	}

	want := []Force{
		{ID: "avon-and-somerset", Name: "Avon and Somerset Constabulary"},
		{ID: "bedfordshire", Name: "Bedfordshire Police"},
		{ID: "cambridgeshire", Name: "Cambridgeshire Constabulary"},
	}

	if !reflect.DeepEqual(forces, want) {
		t.Errorf("Force.GetForces returned returned %v, want %v", forces, want)
	}

}

func TestForceService_GetForceDetails(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/forces/leicestershire", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
			{
				"description": "This is a lovely police force",
				"url": "http://www.leics.police.uk/",
				"engagement_methods": [
					{
						"url": "http://www.facebook.com/pages/Leicester/Leicestershire-Police/76807881169",
						"description": "Become friends with Leicestershire Constabulary",
						"title": "Facebook"
					},
					{
						"url": "http://www.twitter.com/leicspolice",
						"description": "Keep up to date with Leicestershire Constabulary on Twitter",
						"title": "Twitter"
					}
				],
				"telephone": "0116 222 2222",
				"id": "leicestershire",
				"name": "Leicestershire Constabulary"
			}
		`)
	})

	force, _, err := client.Force.GetForceDetails(context.Background(), "leicestershire")
	if err != nil {
		t.Errorf("Force.GetForceDetails returned error: '%s'", err)
	}

	want := Force{
		Description: "This is a lovely police force",
		URL:         "http://www.leics.police.uk/",
		Engagement: []EngagementMethods{
			{
				URL:         "http://www.facebook.com/pages/Leicester/Leicestershire-Police/76807881169",
				Description: "Become friends with Leicestershire Constabulary",
				Title:       "Facebook",
			},
			{
				URL:         "http://www.twitter.com/leicspolice",
				Description: "Keep up to date with Leicestershire Constabulary on Twitter",
				Title:       "Twitter",
			},
		},
		Telephone: "0116 222 2222",
		ID:        "leicestershire",
		Name:      "Leicestershire Constabulary",
	}

	if !reflect.DeepEqual(force, want) {
		t.Errorf("Force.GetForceDetails returned returned %v, want %v", force, want)
	}
}

func TestForceService_GetPeople(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/forces/leicestershire/people", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
			[
				{
					"bio": "Roger joined Lincolnshire Police in 1988...",
					"contact_details": {
						"twitter": "http://www.twitter.com/ACCCLeicsPolice"
					},
					"name": "Roger Bannister",
					"rank": "Assistant Chief Officer (Crime)"
				}
			]
		`)
	})

	forces, _, err := client.Force.GetPeople(context.Background(), "leicestershire")
	if err != nil {
		t.Errorf("Force.GetForces returned error: '%+v'", err)
	}

	want := []SeniorOfficer{
		{
			Bio:            "Roger joined Lincolnshire Police in 1988...",
			ContactDetails: map[string]string{"twitter": "http://www.twitter.com/ACCCLeicsPolice"},
			Name:           "Roger Bannister",
			Rank:           "Assistant Chief Officer (Crime)",
		},
	}

	if !reflect.DeepEqual(forces, want) {
		t.Errorf("Force.GetForces returned returned %v, want %v", forces, want)
	}

}
