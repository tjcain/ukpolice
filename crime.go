package ukpolice

import (
	"context"
)

// @TODO: handle 503 error, which is returned when a request contains over 1000
// results: If a custom area contains more than 10,000 crimes, the API will
// return a 503 status code.

// CrimeService handles communication with the crime related
// method of the data.police.uk API
type CrimeService service

// Crime holds information about individual crimes recorded
type Crime struct {
	Category        string            `json:"category,omitempty"`
	LocationType    string            `json:"location_type,omitempty"`
	Location        Location          `json:"location,omitempty"`
	Context         string            `json:"context,omitempty"`
	OutcomeStatus   map[string]string `json:"outcome_status,omitempty"`
	PersistentID    string            `json:"persistent_id,omitempty"`
	ID              uint              `json:"id,omitempty"`
	LocationSubtype string            `json:"location_subtype,omitempty"`
	Month           string            `json:"month,omitempty"`
}

// Outcome holds information on the outcome of a crime at street-level.
type Outcome struct {
	Category map[string]string `json:"category,omitempty"`
	Date     string            `json:"date,omitempty"`
	PersonID uint              `json:"person_id,omitempty"`
	Crime    Crime             `json:"crime,omitempty"`
}

func (c Crime) String() string {
	return Stringify(c)
}

// GetStreetLevelCrimes returns a list of street level crimes that satisfy the
// criteria provied by a variable of type CrimeQueryOptions. An empty slice
// indicates no data matching the query exists.
func (c *CrimeService) GetStreetLevelCrimes(ctx context.Context, opts ...Option) ([]Crime, *Response, error) {
	u := "crimes-street/all-crime"

	u = addOptions(u, opts...)
	req, err := c.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var crimes []Crime
	resp, err := c.api.Do(ctx, req, &crimes)
	if err != nil {
		return nil, nil, err
	}
	return crimes, resp, nil
}

// GetStreetLevelOutcomes returns Outcomes at street-level; either at a specific
// latitude or longitude, a specific locationID, or within a custom polygonal area.
func (c *CrimeService) GetStreetLevelOutcomes(ctx context.Context, opts ...Option) ([]Outcome, *Response, error) {
	u := "outcomes-at-location"

	u = addOptions(u, opts...)
	req, err := c.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var outcomes []Outcome
	resp, err := c.api.Do(ctx, req, &outcomes)
	if err != nil {
		return nil, nil, err
	}
	return outcomes, resp, nil
}

// Crimes at location
// Crimes with no location
// Crime categories
// Last updated
// Outcomes for a specific crime
