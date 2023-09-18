package service


import (
"github.com/kamlesh1807/Rest_based_microservices_in_Golang/domain")

type CutomerService interface{
	getAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {

	repo domain.CustomerRepository
}

func (s DefaultCustomerService) getAllCustomer() ([]domain.Customer, error){

	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{repository}
}

