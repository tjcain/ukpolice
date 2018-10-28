package ukpolice

import (
	"context"
	"fmt"
)

// CrimeService handles communication with the crime related
// method of the data.police.uk API.
type CrimeService service

// Crime holds information about individual crimes recorded.
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
	Category struct {
		Code string `json:"code,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"category,omitempty"`
	Date     string `json:"date,omitempty"`
	PersonID uint   `json:"person_id,omitempty"`
	Crime    Crime  `json:"crime,omitempty"`
}

// CrimeCategory holds of valid categories.
type CrimeCategory struct {
	URL  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
}

// OutcomesForSpecificCrime holds information returned about the outcomes of a
// specific crime.
type OutcomesForSpecificCrime struct {
	Crime    Crime     `json:"crime,omitempty"`
	Outcomes []Outcome `json:"outcomes,omitempty"`
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

// GetCrimesAtLocation Returns just the crimes which occurred at the specified
//location, rather than those within a radius. If given latitude and longitude,
//finds the nearest pre-defined location and returns the crimes which occurred there.
func (c *CrimeService) GetCrimesAtLocation(ctx context.Context, opts ...Option) ([]Crime, *Response, error) {
	u := "crimes-at-location"
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

// GetCrimesWithNoLocation returns a list of crimes associated to a specified
// police force that could not be mapped to a location. Force is mandatory.
// if no catergory is provided all-crime will be used as default
func (c *CrimeService) GetCrimesWithNoLocation(ctx context.Context, opts ...Option) ([]Crime, *Response, error) {
	u := "crimes-no-location"
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

// GetCrimeCategories returns a list of valid crime categories for a given date.
func (c *CrimeService) GetCrimeCategories(ctx context.Context, date Option) ([]CrimeCategory, *Response, error) {
	u := "crime-categories"
	u = addOptions(u, date)

	req, err := c.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var categories []CrimeCategory
	resp, err := c.api.Do(ctx, req, &categories)
	if err != nil {
		return nil, nil, err
	}

	return categories, resp, nil
}

// GetLastUpdated returns the date when the API was last updated.
func (c *CrimeService) GetLastUpdated(ctx context.Context) (*Date, *Response, error) {
	u := "crime-last-updated"

	req, err := c.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var date *Date
	resp, err := c.api.Do(ctx, req, &date)
	if err != nil {
		return nil, nil, err
	}

	return date, resp, nil
}

// GetSpecificOutcomes returns the crime details and outcome details for a specific crime
func (c *CrimeService) GetSpecificOutcomes(ctx context.Context, persistentID string) (*OutcomesForSpecificCrime, *Response, error) {
	u := fmt.Sprintf("outcomes-for-crime/%s", persistentID)

	req, err := c.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var crimeOutcomes *OutcomesForSpecificCrime
	resp, err := c.api.Do(ctx, req, &crimeOutcomes)
	if err != nil {
		return nil, nil, err
	}

	return crimeOutcomes, resp, nil

}
