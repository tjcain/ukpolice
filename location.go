package ukpolice

// Location holds location information shared by multiple methods.
type Location struct {
	// Universal
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`

	// Used by Crime Methods
	Street Street `json:"street,omitempty"`

	// Used by Neighbourhood methods
	Name        string `json:"name,omitempty"`
	Postcode    string `json:"postcode,omitempty"`
	Address     string `json:"address,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

// Street holds street-level location information.
type Street struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
