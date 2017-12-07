package api

//
// OptionMapResponse -- response to the option map query
//
type OptionMapResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Options []DepartmentMap `json:"options,omitempty"`
}

//
// end of file
//
