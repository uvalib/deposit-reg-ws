package tests

import (
	"depositregws/api"
	"depositregws/client"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"fmt"
	"math/rand"
	"time"
)

type testConfig struct {
	Endpoint string
	Token    string
}

var cfg = loadConfig()

var goodID = "1"
var notFoundID = "x"
var goodToken = cfg.Token
var badToken = "badness"
var empty = " "
var numericChars = "0123456789"
var departmentType = "department"
var degreeType = "degree"

func ensureValidRegistrations(t *testing.T, details []*api.Registration) {

	for _, e := range details {
		if emptyField(e.ID) ||
			emptyField(e.Requester) ||
			emptyField(e.For) ||
			emptyField(e.Department) ||
			emptyField(e.Degree) {
			//emptyField( e.RequestDate ) ||
			//emptyField( e.Status ) {
			t.Fatalf("Expected non-empty field but one is empty\n")
		}
	}
}

func ensureValidMappedOptions(t *testing.T, options []api.DepartmentMap) {

	for _, o := range options {
		if emptyField(o.Department) {
			t.Fatalf("Expected non-empty department field but one is empty\n")
		}
		for _, d := range o.Degrees {
			if emptyField(d) {
				t.Fatalf("Expected non-empty degree field but one is empty\n")
			}
		}
	}
}

func ensureValidOptions(t *testing.T, options api.Options) {

	for _, dep := range options.Departments {
		if emptyField(dep) {
			t.Fatalf("Expected non-empty department field but one is empty\n")
		}
	}
	for _, deg := range options.Degrees {
		if emptyField(deg) {
			t.Fatalf("Expected non-empty degree field but one is empty\n")
		}
	}
}

func createNewReg(t *testing.T) string {
	reg := makeSingleRegistration()
	expected := http.StatusOK
	status, results := client.CreateDepositRequest(cfg.Endpoint, reg, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
	if results == nil || len(results) != 1 {
		t.Fatalf("Incomplete registration details returned")
	}

	return results[0].ID
}

func emptyField(field string) bool {
	return len(strings.TrimSpace(field)) == 0
}

func makeSingleRegistration() api.Registration {
	return api.Registration{
		For:        "dpg3k",
		Requester:  "dpg3k",
		Department: "Engineering",
		Degree:     "Ph.D"}
}

func makeMultiRegistration() api.Registration {
	return api.Registration{
		For:        "dpg3k, tss6n",
		Requester:  "dpg3k",
		Department: "Engineering",
		Degree:     "Ph.D"}
}

func makeNewOption( optionType string ) api.Option {
	return api.Option{
		Option: optionType,
		Value: fmt.Sprintf( "%s-%s", optionType, randomValue( ) ),
	}
}

func makeOptionMap( department string, degrees [] string ) api.DepartmentMap {
	return api.DepartmentMap{
		Department: department,
		Degrees: degrees,
	}
}

func randomValue( ) string {

	// see the RNG
	rand.Seed(time.Now().UnixNano())

	// list of possible characters
	possible := []rune( numericChars )

	return randomString(possible, 10)
}

func randomString(possible []rune, sz int) string {

	b := make([]rune, sz)
	for i := range b {
		b[i] = possible[rand.Intn(len(possible))]
	}
	return string(b)
}

func loadConfig() testConfig {

	data, err := ioutil.ReadFile("service_test.yml")
	if err != nil {
		log.Fatal(err)
	}

	var c testConfig
	if err := yaml.Unmarshal(data, &c); err != nil {
		log.Fatal(err)
	}

	log.Printf("Test config; endpoint   [%s]\n", c.Endpoint)
	log.Printf("Test config; auth token [%s]\n", c.Token)

	return c
}

//
// end of file
//
