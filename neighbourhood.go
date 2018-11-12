package ukpolice

import (
	"context"
	"fmt"
)

// NeighbourhoodService handles communication with the neighbourhood related
// method of the data.police.uk API.
type NeighbourhoodService service

// Neighbourhood holds details of neighbourhoods.
type Neighbourhood struct {
	ForceURL       string              `json:"url_force,omitempty"`
	ContactDetails map[string]string   `json:"contact_details,omitempty"`
	Name           string              `json:"name,omitempty"`
	Links          []map[string]string `json:"links,omitempty"`
	Centre         Location            `json:"centre,omitempty"`
	Locations      []Location          `json:"locations,omitempty"`
	Description    string              `json:"description,omitempty"`
	ID             string              `json:"id,omitempty"`
	Population     string              `json:"population,omitempty"`

	// Used by LocateNeighbourhood
	Force         string `json:"force,omitempty"`
	Neighbourhood string `json:"neighbourhood,omitempty"`
}

func (n Neighbourhood) String() string {
	return Stringify(n)
}

// NeighbourhoodTeam holds details of neighbourhood teams.
type NeighbourhoodTeam struct {
	Bio            string            `json:"bio,omitempty"`
	ContactDetails map[string]string `json:"contact_details,omitempty"`
	Name           string            `json:"name,omitempty"`
	Rank           string            `json:"rank,omitempty"`
}

func (n NeighbourhoodTeam) String() string {
	return Stringify(n)
}

// NeighbourhoodEvent holds details of neighbourhood events.
type NeighbourhoodEvent struct {
	ContactDetails map[string]string `json:"contact_details,omitempty"`
	Description    string            `json:"description,omitempty"`
	Title          string            `json:"title,omitempty"`
	Address        string            `json:"address,omitempty"`
	Type           string            `json:"type,omitempty"`

	//@TODO: make time objects.
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

func (n NeighbourhoodEvent) String() string {
	return Stringify(n)
}

// NeighbourhoodPriorities holds details of neighbourhood priorities.
type NeighbourhoodPriorities struct {
	Action string `json:"action,omitempty"`
	Issue  string `json:"issue,omitempty"`

	//@TODO: make time objects.
	IssueDate  string `json:"issue-date,omitempty"`
	ActionDate string `json:"action-date,omitempty"`
}

func (n NeighbourhoodPriorities) String() string {
	return Stringify(n)
}

// GetNeighbourhoods returns a the neighbourhood details for a given police force.
func (n *NeighbourhoodService) GetNeighbourhoods(ctx context.Context, force string) ([]Neighbourhood, *Response, error) {
	u := fmt.Sprintf("%s/neighbourhoods", force)

	req, err := n.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var neighbourhoods []Neighbourhood
	resp, err := n.api.Do(ctx, req, &neighbourhoods)
	if err != nil {
		return nil, nil, err
	}
	return neighbourhoods, resp, nil

}

// GetSpecificNeighbourhood returns the details of a specific neighbourhood given
// a police force and neighbourhood ID
func (n *NeighbourhoodService) GetSpecificNeighbourhood(ctx context.Context, force, ID string) (*Neighbourhood, *Response, error) {
	u := fmt.Sprintf("%s/%s", force, ID)

	req, err := n.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var neighbourhood *Neighbourhood
	resp, err := n.api.Do(ctx, req, &neighbourhood)
	if err != nil {
		return nil, nil, err
	}
	return neighbourhood, resp, nil

}

// GetNeighbourhoodBoundary returns a list of latitude/longitude pairs that make
// up the boundary of a neighbourhood.
func (n *NeighbourhoodService) GetNeighbourhoodBoundary(ctx context.Context, force, NeighbourhoodID string) ([]Location, *Response, error) {
	u := fmt.Sprintf("%s/%s/boundary", force, NeighbourhoodID)

	req, err := n.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var boundary []Location
	resp, err := n.api.Do(ctx, req, &boundary)
	if err != nil {
		return nil, nil, err
	}
	return boundary, resp, nil
}

// GetNeighbourhoodTeam returns a list of team information for a given force and neighbourhood.
func (n *NeighbourhoodService) GetNeighbourhoodTeam(ctx context.Context, force, NeighbourhoodID string) ([]NeighbourhoodTeam, *Response, error) {
	u := fmt.Sprintf("%s/%s/people", force, NeighbourhoodID)

	req, err := n.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var team []NeighbourhoodTeam
	resp, err := n.api.Do(ctx, req, &team)
	if err != nil {
		return nil, nil, err
	}
	return team, resp, nil
}

// GetNeighbourhoodEvents returns a list of events information for a given force and neighbourhood.
func (n *NeighbourhoodService) GetNeighbourhoodEvents(ctx context.Context, force, NeighbourhoodID string) ([]NeighbourhoodEvent, *Response, error) {
	u := fmt.Sprintf("%s/%s/events", force, NeighbourhoodID)

	req, err := n.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var events []NeighbourhoodEvent
	resp, err := n.api.Do(ctx, req, &events)
	if err != nil {
		return nil, nil, err
	}
	return events, resp, nil
}

// GetNeighbourhoodPriorities returns a list of priorities for a given force and neighbourhood
func (n *NeighbourhoodService) GetNeighbourhoodPriorities(ctx context.Context, force, NeighbourhoodID string) ([]NeighbourhoodPriorities, *Response, error) {
	u := fmt.Sprintf("%s/%s/priorities", force, NeighbourhoodID)

	req, err := n.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var priorities []NeighbourhoodPriorities
	resp, err := n.api.Do(ctx, req, &priorities)
	if err != nil {
		return nil, nil, err
	}
	return priorities, resp, nil
}

// LocateNeighbourhood returns the neighbourhood policing team responsible for a given
// latitude and longitude
func (n *NeighbourhoodService) LocateNeighbourhood(ctx context.Context, lat, long string) (*Neighbourhood, *Response, error) {
	u := fmt.Sprintf("locate-neighbourhood?q=%s,%s", lat, long)

	req, err := n.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var neighbourhood *Neighbourhood
	resp, err := n.api.Do(ctx, req, &neighbourhood)
	if err != nil {
		return nil, nil, err
	}
	return neighbourhood, resp, nil

}
