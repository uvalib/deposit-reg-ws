package api

//
// StandardResponse -- the basic response
//
type StandardResponse struct {
   Status  int             `json:"status"`
   Message string          `json:"message"`
   Details []*Registration `json:"details,omitempty"`
}

//
// end of file
//
