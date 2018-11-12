package ukpolice

// Location holds location information shared by multiple methods.
type Location struct {
	// Universal
	Latitude  string `json:"latitude,omitempty" db:"latitude"`
	Longitude string `json:"longitude,omitempty" db:"longitude"`

	// Used by Crime Methods
	Street struct {
		ID   uint   `json:"id,omitempty" db:"streetid"`
		Name string `json:"name,omitempty" db:"streetname"`
	} `json:"street,omitempty"`

	// Used by Neighbourhood methods
	Name        string `json:"name,omitempty"`
	Postcode    string `json:"postcode,omitempty"`
	Address     string `json:"address,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}
