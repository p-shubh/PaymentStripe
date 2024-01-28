package main

import (
	router "backend/Router"
	"fmt"
)

func main() {

	var (
		id    string
		err   error
		email = "shubhamprajapati032@gmil.com"
		name  = "shubham prajapati"
		price = "23"
	)

	// router.Router()
	id, err = router.GetCustomerIDByEmail(email)
	if id == "" {
		fmt.Println("Error :", err.Error())
		if err.Error() == "customer not found with this email" {
			fmt.Println("Printing CreateCustomerIdByEmail")
			customerId, customerIdErr := router.CreateCustomerIdByEmail(email, name)
			if customerId != "" && customerIdErr == nil {
				router.CreateSubscription(customerId, price, 0)
			}
		}
		//

	} else {
		fmt.Println("id : ", id)
		fmt.Println("Printing response : ")
		response := router.CreateSubscription(id, price, 0)
		fmt.Printf("%+v\n", response)
	}
}
