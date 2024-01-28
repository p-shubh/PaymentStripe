package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/subscription"
)

func Router() {

	d := gin.Default()

	Routes(d)

	gin.SetMode(gin.ReleaseMode)

	d.Run(":8080")
}

func Routes(r *gin.Engine) {

	r.POST("/subscribe", Subscribe)

}

type CardDetails struct {
	Card_numbers string
	CVV          string
	Expiration   string
}

func Subscribe(c *gin.Context) {

	req := CardDetails{}

	if err := c.Bind(&req); err != nil {
		c.JSON(422, gin.H{"error": "Invalid request"})
	}

	c.JSON(http.StatusOK, gin.H{})

}

func SetStripeKey() {
	// Set your secret key. Remember to switch to your live secret key in production.
	// See your keys here: https://dashboard.stripe.com/apikeys
	stripe.Key = "sk_test_51Od72qBCiBx32pTcjP5YW8DowWimdqZopXyyNourpiNZn8cur5WWoCJZwhfzcUONRjORsgFIYuB9E6rUBkNcHqjB00t0FahmNS"
}

// // CreateSubscription creates a subscription for a customer with the given price ID
// func CreateMonthlySubscription(customerID *string, priceID string) (*stripe.Subscription, error) {
// 	SetStripeKey()
// 	params := &stripe.SubscriptionParams{
// 		Customer: customerID,
// 		Items: []*stripe.SubscriptionItemsParams{
// 			{
// 				Price: stripe.String(priceID),
// 			},
// 		},
// 		BillingCycleAnchorConfig: &stripe.SubscriptionBillingCycleAnchorConfigParams{
// 			DayOfMonth: stripe.Int64(31),
// 		},
// 	}

// 	result, err := sub.New(params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// GetCustomerIDByEmail retrieves customer ID by email
func GetCustomerIDByEmail(email string) (string, error) {
	SetStripeKey()

	params := &stripe.CustomerListParams{
		Email: stripe.String(email),
	}

	iter := customer.List(params)
	for iter.Next() {
		return iter.Customer().ID, nil

	}

	if err := iter.Err(); err != nil {
		return "", err
	}

	// If customer not found, you can choose to handle it here
	return "", fmt.Errorf("customer not found with this email")
}

func CreateCustomerIdByEmail(email, name string) (string, error) {

	SetStripeKey()

	params := &stripe.CustomerParams{
		Name:  &name,
		Email: &email,
	}

	result, err := customer.New(params)

	if err != nil {
		return "", err
	}

	fmt.Println(result.ID)

	return result.ID, err

}

func CreateSubscription(id, price string, subscriptionDays int) *stripe.Subscription {

	// Set your secret key. Remember to switch to your live secret key in production.
	// See your keys here: https://dashboard.stripe.com/apikeys
	SetStripeKey()

	params := &stripe.SubscriptionParams{
		Customer:         stripe.String(id),
		CollectionMethod: stripe.String(string(stripe.SubscriptionCollectionMethodChargeAutomatically)),
		Items: []*stripe.SubscriptionItemsParams{
			// &stripe.SubscriptionItemsParams{Price: stripe.String(price)},
			{Price: stripe.String("price_1OdC64BCiBx32pTcPNtd4AMY")}, //stripe price id not price amount
		},
		DefaultPaymentMethod: stripe.String(string("card_1OdCMbBCiBx32pTcQTnNU4sV")),
	}
	response, err := subscription.New(params)
	if err == nil {
		return response
	}
	return nil
}

func CreateDirectAmount() {

}
