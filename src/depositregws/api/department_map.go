package api

//
// DepartmentMap -- a particular option; we have many of these
//
type DepartmentMap struct {
	Department string   `json:"department,omitempty"`
	Degrees    []string `json:"degrees,omitempty"`
}

//
// end of file
//
