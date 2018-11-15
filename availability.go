package ukpolice

import (
	"context"
)

// AvailabilityService handles communication with the availability related
// method of the data.police.uk API
type AvailabilityService service

// AvailabilityInfo holds information about data availability. Date is returned
// as a string not time.Time.
type AvailabilityInfo struct {
	Date          string   `json:"date,omitempty"`
	StopAndSearch []string `json:"stop-and-search,omitempty"`
}

// // AvailabilityInfoResponse holds the slice of returned AvailabilityInfo
// type AvailabilityInfoResponse []*AvailabilityInfo

func (a AvailabilityInfo) String() string {
	return Stringify(a)
}

// GetAvailabilityInfo returns information about data availability.
func (a *AvailabilityService) GetAvailabilityInfo(ctx context.Context) ([]AvailabilityInfo, *Response, error) {
	u := "crimes-street-dates"

	req, err := a.api.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var availabilityInfo []AvailabilityInfo
	resp, err := a.api.Do(ctx, req, &availabilityInfo)
	if err != nil {
		return nil, nil, err
	}

	return availabilityInfo, resp, nil
}
