package ukpolice

import (
	"context"
)

// AvaliabilityService handles communication with the avaliability related
// method of the data.police.uk API
type AvaliabilityService service

// AvaliabilityInfo holds information about data avaliability.
type AvaliabilityInfo struct {
	// @TODO: Convert Date to a time type
	Date          string   `json:"date,omitempty"`
	StopAndSearch []string `json:"stop-and-search,omitempty"`
}

// // AvaliabilityInfoResponse holds the slice of returned AvaliabilityInfo
// type AvaliabilityInfoResponse []*AvaliabilityInfo

func (a AvaliabilityInfo) String() string {
	return Stringify(a)
}

// GetAvaliabilityInfo returns information about data avaliability.
func (a *AvaliabilityService) GetAvaliabilityInfo(ctx context.Context) ([]AvaliabilityInfo, *Response, error) {
	u := "crimes-street-dates"

	req, err := a.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var avaliabilityInfo []AvaliabilityInfo
	resp, err := a.api.Do(ctx, req, &avaliabilityInfo)
	if err != nil {
		return nil, nil, err
	}

	return avaliabilityInfo, resp, nil
}
