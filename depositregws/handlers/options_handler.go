package handlers

import (
	"fmt"
	"github.com/uvalib/deposit-reg-ws/depositregws/api"
	"github.com/uvalib/deposit-reg-ws/depositregws/dao"
	"github.com/uvalib/deposit-reg-ws/depositregws/logger"
	"net/http"
)

//
// OptionsGet -- get options request handler
//
func OptionsGet(w http.ResponseWriter, r *http.Request) {

	optionsSet, err := dao.Store.GetAllOptions()
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		status := http.StatusInternalServerError
		encodeOptionsResponse(w, status,
			fmt.Sprintf("%s (%s)", http.StatusText(status), err),
			api.Options{})
		return
	}

	options := createOptions(optionsSet)

	status := http.StatusOK
	encodeOptionsResponse(w, status, http.StatusText(status), options)
}

func createOptions(pairs []dao.StringPair) api.Options {

	results := api.Options{}
	for _, v := range pairs {
		if v.A == "department" {
			results.Departments = append(results.Departments, v.B)
		} else if v.A == "degree" {
			results.Degrees = append(results.Degrees, v.B)
		}
	}
	return results
}

//
// end of file
//
