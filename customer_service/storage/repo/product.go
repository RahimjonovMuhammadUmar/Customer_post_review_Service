package repo

import (
	pbc "exam/customer_service/genproto/customer"
)

// CustomerStorage
type CustomerStorage interface {
	CreateCustomer(*pbc.CustomerRequest) (*pbc.Customer, error)
	UpdateCustomer(*pbc.Customer) (*pbc.Customer, error)
	CheckIfCustomerExists(id int32) (*pbc.Exists, error)
	GetCustomer(id int32) (*pbc.Customer, error)
	DeleteCustomer(id int32) (*pbc.CustomerDeleted, error)
}
