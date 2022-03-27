package main

import (
	"fmt"
	"kar_go/structure"
)

func main() {
	/* Customer Create */
	receiver1 := structure.NewCustomer(1, "Şevket", "Yılmaz", "Tepeören", 0524765432)
	//fmt.Println(receiver1)
	sender1 := structure.NewCustomer(2, "Adem", "Polat", "Erdek Sokak", 0234123456)
	//fmt.Println(sender1)

	/* Order create */
	order1 := structure.NewOrder("packaging", receiver1, sender1)
	order2 := structure.NewOrder("delivered", receiver1, sender1)
	order3 := structure.NewOrder("returned", receiver1, sender1)
	order4 := structure.NewOrder("packaging", receiver1, sender1)
	order5 := structure.NewOrder("delivered", receiver1, sender1)
	//fmt.Println(order1)

	orders := map[int]structure.Order{
		order1.Id: order1,
		order2.Id: order2,
		order3.Id: order3,
		order4.Id: order4,
		order5.Id: order5,
	}

	/* All Orders List */
	fmt.Println("All Orders")
	for k, v := range orders {
		fmt.Print("Id: ", k)
		fmt.Println("\tStatus: ", v.Status)
	}

	/* Cargo Filtering */
	fmt.Println("\nOrder Filter: Packaging")
	for k, v := range orders {
		if v.Status == "packaging" {
			fmt.Println("Order ID: ", k)
		}
	}

	/* Cargo Status Inquiry */
	var order_inquery int
	fmt.Printf("\nEnter order number: ")
	fmt.Scan(&order_inquery)

	for k, v := range orders {
		if k == order_inquery {
			fmt.Println("Your Order: ", v)
		}
	}
}
