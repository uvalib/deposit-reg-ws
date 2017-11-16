package api

//
// Options -- a particular option; we have many of these
//
type Options struct {
	Department string   `json:"department,omitempty"`
	Degrees    []string `json:"degrees,omitempty"`
}

//
// end of file
//
