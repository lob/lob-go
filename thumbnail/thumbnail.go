package thumbnail

// Thumbnail represents a thumbnail object as returned from the Lob API.
type Thumbnail struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}
