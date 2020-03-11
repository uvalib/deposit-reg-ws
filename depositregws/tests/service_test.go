package tests

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/uvalib/deposit-reg-ws/depositregws/api"
	"github.com/uvalib/deposit-reg-ws/depositregws/client"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"testing"
	"time"
)

type testConfig struct {
	Endpoint string
	Secret   string
}

var cfg = loadConfig()

var goodID = "1"
var notFoundID = "x"
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
	status, results := client.CreateDepositRequest(cfg.Endpoint, reg, goodToken(cfg.Secret))
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

func makeNewOption(optionType string) api.Option {
	return api.Option{
		Option: optionType,
		Value:  fmt.Sprintf("%s-%s", optionType, randomValue()),
	}
}

func makeOptionMap(department string, degrees []string) api.DepartmentMap {
	return api.DepartmentMap{
		Department: department,
		Degrees:    degrees,
	}
}

func randomValue() string {

	// see the RNG
	rand.Seed(time.Now().UnixNano())

	// list of possible characters
	possible := []rune(numericChars)

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

	log.Printf("endpoint   [%s]\n", c.Endpoint)
	log.Printf("secret     [%s]\n", c.Secret)

	return c
}

func badToken(secret string) string {

	// Declare the expiration time of the token
	expirationTime := time.Now().Add(-5 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &jwt.StandardClaims{
		// In JWT, the expiry time is expressed as unix milliseconds
		ExpiresAt: expirationTime.Unix(),
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

func goodToken(secret string) string {

	// Declare the expiration time of the token
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &jwt.StandardClaims{
		// In JWT, the expiry time is expressed as unix milliseconds
		ExpiresAt: expirationTime.Unix(),
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

//
// end of file
//
