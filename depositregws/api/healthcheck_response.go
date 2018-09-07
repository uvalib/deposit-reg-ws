package api

//
// HealthCheckResponse -- response to the health check query
//
type HealthCheckResponse struct {
	CheckType HealthCheckResult `json:"mysql"`
}

//
// end of file
//
