package main

import (
	"fmt"
	"time"
)

// embedding
type customer struct {
	name  string
	phone string
}

// order representation struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
	customer            //embedding other struct
}

// func newOrder(id string, amount float32, status string) *order {
// 	// initial setup goes here
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
// }

// // receiver type
// func (o *order) changeStatus(status string) {
// 	o.status = status //no need of using derefrence again because struct does the work automatically
// }

// func (o order) getAmount() float32 { //use pointer when only changing the value otherwise it already uses the address
// 	return o.amount
// }

func main() {

	// newCustomer := customer{
	// 	name:  "jphn",
	// 	phone: "167899",
	// }
	newOrder := order{
		id:     "1",
		amount: 30,
		status: "pending",
		customer: customer{name: "john",
			phone: "12356"},
	}
	newOrder.customer.name = "robin"
	fmt.Println(newOrder.customer)

	// in-line struct
	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{"Golang", true}
	// fmt.Println(language)

	// myOrder := newOrder("1", 20.5, "received")
	// var  order order =
	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50,
	// 	status: "recieved",
	// 	//nil value in unfilled field, deafult is zero values
	// }

	// myOrder.changeStatus("delivered")
	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.status)
	// fmt.Println(myOrder.getAmount())
	// fmt.Println(myOrder)

	// myOrder2 := order{
	// 	id:        "2",
	// 	amount:    20,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }

	// myOrder.status = "paid"
	// fmt.Println("orders struct", myOrder, "\n", myOrder2)

}
