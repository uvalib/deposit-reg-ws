package api

//
// Registration -- details of an optional registration
//
type Registration struct {
   ID          string `json:"id,omitempty"`
   Requester   string `json:"requester,omitempty"`
   For         string `json:"for,omitempty"`
   Department  string `json:"department,omitempty"`
   Degree      string `json:"degree,omitempty"`
   RequestDate string `json:"request_date,omitempty"`
   DepositDate string `json:"deposit_date,omitempty"`
   Status      string `json:"status,omitempty"`
}

//
// end of file
//