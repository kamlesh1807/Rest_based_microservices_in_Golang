package domain


type CustomerRepositoryStub struct {
	customers []Customer
  }


func (s CustomerRepositoryStub) FindAll() ([]Customer, error){

	return s.customers, nil
  }


func NewCustomerRepositoryStub() CustomerRepositoryStub{
    customers:= []Customer {
	{"10001","ashish","new delhi","110011","2000-08-09", "1"},
	{"2001","Ankit"," Noida","110022","2003-6-30", "1"},
    
    }
  return CustomerRepositoryStub{customers}
  }