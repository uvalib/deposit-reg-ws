package handlers

import (
   "depositregws/api"
   "depositregws/logger"
   "encoding/json"
   "fmt"
   "log"
   "net/http"
   "strings"
)

func encodeStandardResponse(w http.ResponseWriter, status int, message string, details []*api.Registration) {

   logger.Log(fmt.Sprintf("Status: %d (%s)\n", status, message))
   jsonAttributes(w)
   coorsAttributes(w)
   w.WriteHeader(status)
   if err := json.NewEncoder(w).Encode(api.StandardResponse{Status: status, Message: message, Details: details}); err != nil {
      log.Fatal(err)
   }
}

func encodeOptionsResponse(w http.ResponseWriter, status int, message string, options *api.Options) {

   jsonAttributes(w)
   coorsAttributes(w)
   w.WriteHeader(status)
   if err := json.NewEncoder(w).Encode(api.OptionsResponse{Status: status, Message: message, Options: options}); err != nil {
      log.Fatal(err)
   }
}

func encodeHealthCheckResponse(w http.ResponseWriter, status int, message string) {
   healthy := status == http.StatusOK
   jsonAttributes(w)
   w.WriteHeader(status)
   if err := json.NewEncoder(w).Encode(api.HealthCheckResponse{CheckType: api.HealthCheckResult{Healthy: healthy, Message: message}}); err != nil {
      log.Fatal(err)
   }
}

func encodeVersionResponse(w http.ResponseWriter, status int, version string) {
   jsonAttributes(w)
   w.WriteHeader(status)
   if err := json.NewEncoder(w).Encode(api.VersionResponse{Version: version}); err != nil {
      log.Fatal(err)
   }
}

func encodeRuntimeResponse(w http.ResponseWriter, status int, version string, cpus int, goroutines int, heapcount uint64, alloc uint64) {
   jsonAttributes(w)
   w.WriteHeader(status)
   if err := json.NewEncoder(w).Encode(api.RuntimeResponse{Version: version, CPUCount: cpus, GoRoutineCount: goroutines, ObjectCount: heapcount, AllocatedMemory: alloc}); err != nil {
      log.Fatal(err)
   }
}

func jsonAttributes(w http.ResponseWriter) {
   w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func coorsAttributes(w http.ResponseWriter) {
   w.Header().Set("Access-Control-Allow-Origin", "*")
   w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func jsonResponse(w http.ResponseWriter) {
   w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func notEmpty(param string) bool {
   return len(strings.TrimSpace(param)) != 0
}

//
// end of file
//
