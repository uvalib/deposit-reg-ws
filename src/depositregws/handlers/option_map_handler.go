package handlers

import (
	"depositregws/api"
	"depositregws/dao"
	"depositregws/logger"
	"fmt"
	"net/http"
)

//
// OptionMapGet -- get option map request handler
//
func OptionMapGet(w http.ResponseWriter, r *http.Request) {

	optionsSet, err := dao.DB.GetMappedOptions()
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		status := http.StatusInternalServerError
		encodeOptionMapResponse(w, status,
			fmt.Sprintf("%s (%s)", http.StatusText(status), err),
			nil)
		return
	}

	options := createOptionsMap(optionsSet)

	status := http.StatusOK
	encodeOptionMapResponse(w, status, http.StatusText(status), options)
}

func createOptionsMap(pairs []dao.StringPair) []api.DepartmentMap {

	results := make([]api.DepartmentMap, 0)
	for _, v := range pairs {

		// if we have a blank value, it means that there are no mapped fields for this value
		// just add it to the list with an empty list
		if len( v.B ) == 0 {
			results = append( results, api.DepartmentMap{ Department: v.A, Degrees: []string{}})
			continue
		}

		// otherwise, see if we have already got a mapped value and if so append it to the list or create a new
		// entry with a single mapped value
		ix := indexOf(results, v.A)
		if ix >= 0 {
			results[ix].Degrees = append(results[ix].Degrees, v.B)
		} else {
			results = append(results, api.DepartmentMap{Department: v.A, Degrees: []string{v.B}})
		}
	}
	return results
}

func indexOf(options []api.DepartmentMap, option string) int {
	for ix, v := range options {

		if v.Department == option {
			return ix
		}
	}
	// not found
	return -1
}

//
// end of file
//
