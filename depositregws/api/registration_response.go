package api

//
// RegistrationResponse -- the basic response
//
type RegistrationResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Details []*Registration `json:"details,omitempty"`
}

//
// end of file
//
