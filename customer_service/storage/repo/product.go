package repo

import (
	pbc "exam/customer_service/genproto/customer"
)

// CustomerStorage
type CustomerStorage interface {
	CreateCustomer(*pbc.CustomerRequest) (*pbc.CustomerWithoutPost, error)
	UpdateCustomer(*pbc.CustomerWithoutPost) (*pbc.CustomerWithoutPost, error)
	CheckIfCustomerExists(id int32) (*pbc.Exists, error)
	GetCustomer(id int32) (*pbc.Customer, error)
	GetCustomerForLogin(email string) (*pbc.CustomerWithoutPost, error)
	DeleteCustomer(id int32) (*pbc.CustomerDeleted, error)
	CheckField(field, value string) (*pbc.Exists, error)
	SearchCustomer(field, value, orderBy, ascOrDesc string, limit, page int32) (*pbc.PossibleCustomers, error)
	IsAdmin(username string) (*pbc.Admin, error)
	IsModerator(username string) (*pbc.Admin, error)
}
