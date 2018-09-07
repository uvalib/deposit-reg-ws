package api

//
// Options -- a particular option; we have many of these
//
type Options struct {
	Departments []string `json:"departments,omitempty"`
	Degrees     []string `json:"degrees,omitempty"`
}

//
// end of file
//
