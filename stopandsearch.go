package ukpolice

import (
	"context"
	"time"
)

// StopAndSearchService handles communication with the stop and search related
// method of the data.police.uk API.
type StopAndSearchService service

// Search holds information relating to individual stop and searches.
type Search struct {
	ID                             int           `json:"id"`
	AgeRange                       string        `json:"age_range"`
	Type                           string        `json:"type"`
	Gender                         string        `json:"gender"`
	Outcome                        SearchOutcome `json:"outcome"`
	InvolvedPerson                 bool          `json:"involved_person"`
	SelfDefinedEthnicity           string        `json:"self_defined_ethnicity"`
	OfficerDefinedEthnicity        string        `json:"officer_defined_ethnicity"`
	DateTime                       time.Time     `json:"datetime"`
	RemovalOfMoreThanOuterClothing bool          `json:"removal_of_more_than_outer_clothing"`
	Location                       Location      `json:"location"`
	Operation                      bool          `json:"operation"`
	OperationName                  string        `json:"operation_name"`
	OutcomeLinkedToObject          bool          `json:"outcome_linked_to_object_of_search"`
	ObjectOfSearch                 string        `json:"object_of_search"`
	Legislation                    string        `json:"legislation"`
	// Force is not supplied natively by the API - if you want to record which
	// force a search belongs to update this field after fetching.
	Force string `json:"force"`
}

func (s Search) String() string {
	return Stringify(s)
}

// SearchOutcome holds details of search outcomes. The 'outcome' result provided
// by the data.police.uk api returns both string and bool types, this struct
// and the custom UnmarshalJSON satisfy type security.
type SearchOutcome struct {
	Desc           string `json:"outcome_desc"`
	SearchHappened bool   `json:"searched"`
}

func (o SearchOutcome) String() string {
	return Stringify(o)
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (o *SearchOutcome) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "false" {
		*o = SearchOutcome{SearchHappened: false, Desc: ""}
	} else {
		*o = SearchOutcome{SearchHappened: true, Desc: s}

	}
	return nil
}

// GetStopAndSearchesByArea returns stop and searches at street-level;
// either within a 1 mile radius of a single point, or within a custom area.
func (s *StopAndSearchService) GetStopAndSearchesByArea(ctx context.Context, opts ...Option) ([]Search, *Response, error) {
	u := "stops-street"

	u = addOptions(u, opts...)
	req, err := s.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var searches []Search
	resp, err := s.api.Do(ctx, req, &searches)
	if err != nil {
		return nil, nil, err
	}
	return searches, resp, nil
}

// GetStopAndSearchesByLocation returns stop and searches at a particular location.
func (s *StopAndSearchService) GetStopAndSearchesByLocation(ctx context.Context, opts ...Option) ([]Search, *Response, error) {
	u := "stops-at-location"

	u = addOptions(u, opts...)
	req, err := s.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var searches []Search
	resp, err := s.api.Do(ctx, req, &searches)
	if err != nil {
		return nil, nil, err
	}
	return searches, resp, nil
}

// GetStopAndSearchesWithNoLocation returns stop and searches with no location
// provided for a given police force.
func (s *StopAndSearchService) GetStopAndSearchesWithNoLocation(ctx context.Context, opts ...Option) ([]Search, *Response, error) {
	u := "stops-no-location"

	u = addOptions(u, opts...)
	req, err := s.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var searches []Search
	resp, err := s.api.Do(ctx, req, &searches)
	if err != nil {
		return nil, nil, err
	}
	return searches, resp, nil
}

// GetStopAndSearchesByForce returns stop and searches reported by a given police force.
func (s *StopAndSearchService) GetStopAndSearchesByForce(ctx context.Context, opts ...Option) ([]Search, *Response, error) {
	u := "stops-force"

	u = addOptions(u, opts...)
	req, err := s.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var searches []Search
	resp, err := s.api.Do(ctx, req, &searches)
	if err != nil {
		return nil, nil, err
	}
	return searches, resp, nil
}
