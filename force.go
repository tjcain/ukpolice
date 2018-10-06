package ukpolice

import (
	"context"
	"fmt"
)

// ForceService handles communication with the force related
// method of the data.police.uk API
type ForceService service

// Force holds information about a police force
type Force struct {
	ID          string              `json:"id,omitempty"`
	Name        string              `json:"name,omitempty"`
	Telephone   string              `json:"telephone,omitempty"`
	URL         string              `json:"url,omitempty"`
	Description string              `json:"description,omitempty"`
	Engagement  []EngagementMethods `json:"engagement_methods,omitempty"`
}

// EngagementMethods holds information on a specific police force's social media.
type EngagementMethods struct {
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	Title       string `json:"title,omitempty"`
}

func (f Force) String() string {
	return Stringify(f)
}

// SeniorOfficer holds information on Senior Officers within a police force.
type SeniorOfficer struct {
	Bio            string            `json:"bio,omitempty"`
	ContactDetails map[string]string `json:"contact_details,omitempty"`
	Name           string            `json:"name,omitempty"`
	Rank           string            `json:"rank,omitempty"`
}

func (so SeniorOfficer) String() string {
	return Stringify(so)
}

// GetForces returns a slice containing all avaliable police forces.
func (f *ForceService) GetForces(ctx context.Context) ([]Force, *Response, error) {
	u := "forces"

	req, err := f.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var forces []Force
	resp, err := f.api.Do(ctx, req, &forces)
	if err != nil {
		return nil, nil, err
	}

	return forces, resp, nil
}

// GetForceDetails returns more information about the provided force
func (f *ForceService) GetForceDetails(ctx context.Context, force string) (Force, *Response, error) {
	u := fmt.Sprintf("forces/%s", force)
	forceDetails := Force{}
	req, err := f.api.NewRequest("GET", u, nil)
	if err != nil {
		return forceDetails, nil, err
	}

	resp, err := f.api.Do(ctx, req, &forceDetails)
	if err != nil {
		return forceDetails, nil, err
	}

	return forceDetails, resp, nil
}

// GetPeople returns a slice containing details of the senior police officers
// of the requested police force.
func (f *ForceService) GetPeople(ctx context.Context, force string) ([]SeniorOfficer, *Response, error) {
	u := fmt.Sprintf("forces/%s/people", force)

	req, err := f.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var officers []SeniorOfficer
	resp, err := f.api.Do(ctx, req, &officers)
	if err != nil {
		return nil, nil, err
	}

	return officers, resp, nil
}
