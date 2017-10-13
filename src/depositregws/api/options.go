package api

//
// Options -- a perticular option; we have many of these
//
type Options struct {
   Department []string `json:"department,omitempty"`
   Degree     []string `json:"degree,omitempty"`
}

//
// end of file
//