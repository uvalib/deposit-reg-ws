package api

type Registration struct {
    Id             string   `json:"id,omitempty"`
    For            string   `json:"for,omitempty"`
    School         string   `json:"school,omitempty"`
    Degree         string   `json:"degree,omitempty"`
    RequestDate    string   `json:"request_date,omitempty"`
    DepositDate    string   `json:"deposit_date,omitempty"`
    Status         string   `json:"status,omitempty"`
}