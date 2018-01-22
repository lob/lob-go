package tracking

// Event represents a tracking event object as returned from the Lob API.
type Event struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Time         string `json:"time"`
	DateCreated  string `json:"date_created"`
	DateModified string `json:"date_modified"`
}

// TODO(damian): Add area mail compatibility
