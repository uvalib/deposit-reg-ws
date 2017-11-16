package api

//
// OptionsResponse -- response to the options query
//
type OptionsResponse struct {
   Status     int      `json:"status"`
   Message    string   `json:"message"`
   Options [] Options  `json:"options,omitempty"`
}

//
// end of file
//