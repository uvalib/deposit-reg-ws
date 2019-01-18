package tests

import (
	"github.com/uvalib/deposit-reg-ws/depositregws/api"
	"github.com/uvalib/deposit-reg-ws/depositregws/client"
	"net/http"
	"testing"
)

var badDepartment = "xxxyyyzzz"
var badDegree = "aaabbbccc"

//
// options map update tests
//

func TestOptionsMapUpdateHappyDay(t *testing.T) {
	expected := http.StatusOK
	optionMap := makeNewOptionsMap(t)
	status := client.AddOptionMap(cfg.Endpoint, optionMap, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	status, options := client.GetMappedOptions(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	ensureCorrectMap(t, optionMap, options)
}

func TestOptionsMapUpdateBadDepartment(t *testing.T) {
	expected := http.StatusNotFound
	optionMap := makeGoodOptionsMap(t)
	optionMap.Department = badDepartment
	status := client.AddOptionMap(cfg.Endpoint, optionMap, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestOptionsMapUpdateBadDegree(t *testing.T) {
	expected := http.StatusNotFound
	optionMap := makeGoodOptionsMap(t)
	optionMap.Degrees[0] = badDegree
	status := client.AddOptionMap(cfg.Endpoint, optionMap, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestOptionsMapUpdateEmptyOptionMap(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.AddOptionMap(cfg.Endpoint, api.DepartmentMap{}, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestOptionsMapUpdateBadToken(t *testing.T) {
	expected := http.StatusForbidden
	optionMap := makeGoodOptionsMap(t)
	status := client.AddOptionMap(cfg.Endpoint, optionMap, badToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func makeNewOptionsMap(t *testing.T) api.DepartmentMap {

	expected := http.StatusOK
	departmentOption := makeNewOption(departmentType)
	status := client.AddOption(cfg.Endpoint, departmentOption, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	degreeOneOption := makeNewOption(degreeType)
	status = client.AddOption(cfg.Endpoint, degreeOneOption, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	degreeTwoOption := makeNewOption(degreeType)
	status = client.AddOption(cfg.Endpoint, degreeTwoOption, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	degrees := make([]string, 2)
	degrees[0] = degreeOneOption.Value
	degrees[1] = degreeTwoOption.Value

	return api.DepartmentMap{
		Department: departmentOption.Value,
		Degrees:    degrees,
	}
}

func makeGoodOptionsMap(t *testing.T) api.DepartmentMap {

	expected := http.StatusOK
	status, options := client.GetOptions(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if len(options.Departments) == 0 || len(options.Degrees) == 0 {
		t.Fatalf("Got empty options response\n")
	}

	degrees := make([]string, 2)
	degrees[0] = options.Degrees[0]
	degrees[1] = options.Degrees[1]

	return api.DepartmentMap{
		Department: options.Departments[0],
		Degrees:    degrees,
	}
}

func ensureCorrectMap(t *testing.T, mapSet api.DepartmentMap, allMaps []api.DepartmentMap) {

	for _, m := range allMaps {
		if m.Department == mapSet.Department {
			if len(mapSet.Degrees) != len(m.Degrees) {
				t.Fatalf("Degree count is different\n")
			}

			for _, d := range mapSet.Degrees {
				if arrayContains(d, m.Degrees) == false {
					t.Fatalf("Degree %s not located in options map\n", d)
				}
			}

			// done checking, all looks good
			return
		}
	}
	t.Fatalf("Department %s not located in options map\n", mapSet.Department)
}

func arrayContains(value string, array []string) bool {

	for _, s := range array {
		if s == value {
			return true
		}
	}
	return false
}

//
// end of file
//
