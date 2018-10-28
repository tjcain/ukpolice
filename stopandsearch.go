package ukpolice

import "context"

// StopAndSearchService handles communication with the stop and search related
// method of the data.police.uk API.
type StopAndSearchService service

// Search holds information relating to individual stop and searches.
type Search struct {
	AgeRange                       string      `json:"age_range,omitempty"`
	Type                           string      `json:"type,omitempty"`
	Gender                         string      `json:"gender,omitempty"`
	Outcome                        interface{} `json:"outcome,omitempty"`
	InvolvedPerson                 bool        `json:"involved_person,omitempty"`
	SelfDefinedEthnicity           string      `json:"self_defined_ethnicity,omitempty"`
	OfficerDefinedEthnicity        string      `json:"officer_defined_ethnicity,omitempty"`
	DateTime                       string      `json:"datetime,omitempty"`
	RemovalOfMoreThanOuterClothing bool        `json:"removal_of_more_than_outer_clothing,omitempty"`
	OutcomeObject                  struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"outcome_object,omitempty"`
	Location              Location `json:"location,omitempty"`
	Operation             bool     `json:"operation,omitempty"`
	OperationName         string   `json:"operation_name,omitempty"`
	OutcomeLinkedToObject bool     `json:"outcome_linked_to_object_of_search,omitempty"`
	ObjectOfSearch        string   `json:"object_of_search,omitempty"`
	Legislation           string   `json:"legislation,omitempty"`
}

func (s Search) String() string {
	return Stringify(s)
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
