package design_principles

import (
	"database/sql"
	"errors"
	"regexp"
)

var (
	InvalidEmailErr         = errors.New("invalid email")
	AuthorizationFailureErr = errors.New("authorization failure")
)

type Customer struct{}

type CustomerService struct {
	CustomerRepo
}

func NewCustomerService(customerRepo CustomerRepo) *CustomerService {
	return &CustomerService{CustomerRepo: customerRepo} // 依赖注入
}

func (s *CustomerService) Login(email, password string) (*Customer, error) {
	existed, err := s.CustomerRepo.CheckIfUserExisted(email, password)
	if err != nil {
		return nil, err
	}

	if !existed {
		return nil, AuthorizationFailureErr
	}

	return s.CustomerRepo.GetUserByEmail(email)
}

type CustomerRepo struct {
	db *sql.DB
}

func (r *CustomerRepo) CheckIfUserExisted(email, password string) (bool, error) {
	if !isValidEmail(email) {
		return false, InvalidEmailErr
	}

	if !isValidPassword(password) {
		return false, InvalidPasswordErr
	}

	// ...query db to check if email&password exists...
	_, err := r.db.Query("")
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *CustomerRepo) GetUserByEmail(email string) (*Customer, error) {
	if !isValidEmail(email) {
		return nil, InvalidEmailErr
	}

	//...query db to get user by email...
	_, err := r.db.Query("")
	if err != nil {
		return nil, err
	}

	return &Customer{}, nil
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(email) {
		return false
	}
	return true
}

func isValidPassword(password string) bool {
	return false
}

// ------------------------------
// 去除重复逻辑

type CustomerServiceV2 struct {
	CustomerRepoV2
}

func NewCustomerServiceV2(customerRepoV2 CustomerRepoV2) *CustomerServiceV2 {
	return &CustomerServiceV2{CustomerRepoV2: customerRepoV2}
}

func (s *CustomerServiceV2) Login(email, password string) (*Customer, error) {
	if !isValidEmail(email) {
		return nil, InvalidEmailErr
	}

	if !isValidPassword(password) {
		return nil, InvalidPasswordErr
	}

	return s.CustomerRepoV2.GetUserByEmail(email)
}

type CustomerRepoV2 struct {
	db *sql.DB
}

func (r *CustomerRepoV2) CheckIfUserExisted(email, password string) (bool, error) {
	return false, nil
}

func (r CustomerRepoV2) GetUserByEmail(email string) (*Customer, error) {
	return nil, nil
}
