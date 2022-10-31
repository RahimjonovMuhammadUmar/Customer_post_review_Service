package postgres

import (
	pbc "exam/customer_service/genproto/customer"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type customerRepo struct {
	db *sqlx.DB
}

// NewCustomerRepo ...
func NewCustomerRepo(db *sqlx.DB) *customerRepo {
	return &customerRepo{db: db}
}

func (c *customerRepo) CreateCustomer(req *pbc.CustomerRequest) (*pbc.CustomerWithoutPost, error) {
	inserted := &pbc.CustomerWithoutPost{}
	tx, err := c.db.Begin()
	if err != nil {
		fmt.Println("Error while declaring tx", err)
		return &pbc.CustomerWithoutPost{}, err
	}
	err = tx.QueryRow(`INSERT INTO customers(
	 first_name,
	 last_name,
	 bio,
	 email,
	 phone_number,
	 refresh_token) VALUES($1, $2, $3, $4, $5, $6) RETURNING 
	 id, 
	 first_name, 
	 last_name, 
	 bio, 
	 email, 
	 phone_number,
	 refresh_token`, req.FirstName, req.LastName, req.Bio, req.Email, req.PhoneNumber, req.Token).Scan(
		&inserted.Id,
		&inserted.FirstName,
		&inserted.LastName,
		&inserted.Bio,
		&inserted.Email,
		&inserted.PhoneNumber,
		&inserted.RefreshToken,
	)
	if err != nil {
		fmt.Println("error while inserting into customers", err)
		tx.Rollback()
		return &pbc.CustomerWithoutPost{}, err
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("Error while commiting", err)
		tx.Rollback()
	}
	addresses := []*pbc.Address{}
	for _, address := range req.Addresses {
		address_ := &pbc.Address{}
		err := c.db.QueryRow(`INSERT INTO addresses(
		street, 
		house_number, 
		customer_id) VALUES($1, $2, $3) RETURNING 
		id, 
		street, 
		house_number`, address.Street, address.HouseNumber, inserted.Id).Scan(
			&address_.Id,
			&address_.Street,
			&address_.HouseNumber,
		)
		if err != nil {
			fmt.Println("error while inserting into addresses", err)
			tx.Rollback()
			return &pbc.CustomerWithoutPost{}, err
		}
		addresses = append(addresses, address_)
	}
	inserted.Addresses = addresses

	err = tx.Commit()
	if err != nil {
		fmt.Println("Error while commiting 2", err)
		tx.Rollback()
	}
	return inserted, nil
}

func (c *customerRepo) UpdateCustomer(req *pbc.CustomerWithoutPost) (*pbc.CustomerWithoutPost, error) {
	_, err := c.db.Exec(`UPDATE customers SET 
	first_name = $1,
	last_name = $2,
	bio = $3,
	email = $4,
	phone_number = $5, updated_at = NOW() WHERE id = $6`, req.FirstName, req.LastName, req.Bio, req.Email, req.PhoneNumber, req.Id)
	if err != nil {
		fmt.Println("error while updating customer", err)
		return &pbc.CustomerWithoutPost{}, err
	}
	for _, address := range req.Addresses {
		_, err = c.db.Exec(`UPDATE addresses SET street = $1, house_number = $2 WHERE id = $3`, address.Street, address.HouseNumber, address.Id)
		if err != nil {
			fmt.Println("Error while updating customers", err)
			return &pbc.CustomerWithoutPost{}, err
		}
	}

	return req, nil
}

func (c *customerRepo) CheckIfCustomerExists(id int32) (*pbc.Exists, error) {
	var check int
	err := c.db.QueryRow(`SELECT 1 FROM customers WHERE id = $1`, id).Scan(&check)
	if check == 0 {
		return &pbc.Exists{
			Exists: false}, nil
	} else if err != nil {
		fmt.Println("error while checking id from customers", err)
		return &pbc.Exists{}, err
	}

	return &pbc.Exists{
		Exists: true}, nil
}

func (c *customerRepo) GetCustomer(id int32) (*pbc.Customer, error) {
	customerData := &pbc.Customer{}
	err := c.db.QueryRow(`SELECT 
	id, 
	first_name, 
	last_name, 
	bio, 
	email, 
	phone_number,
	refresh_token FROM customers WHERE id = $1 and deleted_at IS NULL`, id).Scan(
		&customerData.Id,
		&customerData.FirstName,
		&customerData.LastName,
		&customerData.Bio,
		&customerData.Email,
		&customerData.PhoneNumber,
	)
	// if err.Error() == "sql: no rows in result set" {
	// 	fmt.Println("133")
	// 	return &pbc.Customer{}, err
	// }
	if err != nil {
		fmt.Println("error while selecting customer by id", err)
		return &pbc.Customer{}, err
	}
	addresses := []*pbc.Address{}
	addresses_rows, err := c.db.Query(`SELECT  id, street, house_number from addresses WHERE customer_id = $1`, customerData.Id)
	if err != nil {
		fmt.Println("error while selecting addresses of customer", err)
		return &pbc.Customer{}, err
	}

	for addresses_rows.Next() {
		address := &pbc.Address{}
		err = addresses_rows.Scan(
			&address.Id,
			&address.Street,
			&address.HouseNumber,
		)
		if err != nil {
			fmt.Println("Error while scanning to address from addresses_rows", err)
			return &pbc.Customer{}, err
		}
		addresses = append(addresses, address)
	}
	customerData.Addresses = addresses
	return customerData, nil
}

func (c *customerRepo) DeleteCustomer(id int32) (*pbc.CustomerDeleted, error) {
	result, err := c.db.Exec(`UPDATE customers SET deleted_at = NOW() WHERE id = $1`, id)
	if err != nil {
		fmt.Println("Error while deleting from customer", err)
		return &pbc.CustomerDeleted{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("error while rowsAffected, err := result.RowsAffected();", err)
		return &pbc.CustomerDeleted{}, err
	}
	if rowsAffected == 1 {
		return &pbc.CustomerDeleted{
			CustomerDeleted: true,
		}, nil
	}

	return &pbc.CustomerDeleted{
		CustomerDeleted: false,
	}, nil
}

func (c *customerRepo) CheckField(field, value string) (*pbc.Exists, error) {
	var query = fmt.Sprintf("SELECT 1 FROM customers WHERE %s = $1", field)
	var check int
	err := c.db.QueryRow(query, value).Scan(&check)
	if check == 0 {
		return &pbc.Exists{
			Exists: false}, nil
	} else if err != nil {
		fmt.Println("error while checking field from customers", err)
		return &pbc.Exists{}, err
	}

	return &pbc.Exists{
		Exists: true}, nil
}
func (c *customerRepo) SearchCustomer(field, value, orderBy, ascOrDesc string, limit, page int32) (*pbc.PossibleCustomers, error) {
	query := fmt.Sprintf("SELECT id, first_name, last_name, bio, email, phone_number FROM customers WHERE %s ~ '%s'", field, value)
	if orderBy != "" {
		query += " ORDER BY " + orderBy
	}
	if ascOrDesc != ""{
		query = query + " " + ascOrDesc
	}
	rows, err := c.db.Query(query+" LIMIT $1 OFFSET $2", limit, ((page - 1) * 10))
	if err != nil {
		fmt.Println("error while searching by customer", err)
		return &pbc.PossibleCustomers{}, err
	}

	defer rows.Close()
	alike_customers := &pbc.PossibleCustomers{}
	for rows.Next() {
		customerData := &pbc.CustomerWithoutPost{}
		err = rows.Scan(
			&customerData.Id,
			&customerData.FirstName,
			&customerData.LastName,
			&customerData.Bio,
			&customerData.Email,
			&customerData.PhoneNumber,
		)
		if err != nil {
			fmt.Println("error while scanning alike data", err)
			return &pbc.PossibleCustomers{}, err
		}
		alike_customers.PossibleCustomers = append(alike_customers.PossibleCustomers, customerData)

	}

	return alike_customers, nil
}
