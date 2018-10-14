package ukpolice

// Location holds location information shared by multiple methods.
type Location struct {
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Street    Street `json:"street,omitempty"`
}

// Street holds street-level location information.
type Street struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
