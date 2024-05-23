package dao

import (
	"github.com/uvalib/deposit-reg-ws/depositregws/api"
)

// our storage interface
type Storage interface {
	Check() error
	GetDepositRequest(id string) ([]*api.Registration, error)
	SearchDepositRequest(id string) ([]*api.Registration, error)
	CreateDepositRequest(reg api.Registration) (*api.Registration, error)
	DeleteDepositRequest(id string) (int64, error)
	GetMappedOptions() ([]StringPair, error)
	GetAllOptions() ([]StringPair, error)
	CreateOption(option api.Option) error
	UpdateOptionMap(optionMap api.DepartmentMap) error
	//Destroy() error
}

// StringPair -- used for some results; not idiomatic
type StringPair struct {
	A string
	B string
}

// our singleton store
var Store Storage

// our factory
func NewDatastore() error {
	var err error
	// mock implementation here
	Store, err = newDBStore()
	return err
}

//
// end of file
//
