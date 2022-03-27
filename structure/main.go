package structure

import (
	"math/rand"
)

type Order struct {
	Id       int
	Status   string
	Receiver Customer
	Sender   Customer
}

type Customer struct {
	Id_number    int
	Name         string
	Surname      string
	Address      string
	Phone_number int
}

func NewCustomer(id int, name, surname, address string, phone int) Customer {
	return Customer{
		Id_number:    id,
		Name:         name,
		Surname:      surname,
		Address:      address,
		Phone_number: phone,
	}
}

func NewOrder(status string, receiver Customer, sender Customer) Order {
	return Order{
		Id:       rand.Int(),
		Status:   status,
		Receiver: receiver,
		Sender:   sender,
	}
}

func (o *Order) SetOrderStatus(status string) {
	o.Status = status
}

func (c *Customer) SetCustomerAddress(address string) {
	c.Address = address
}
