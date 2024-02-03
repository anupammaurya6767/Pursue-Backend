package middlewares

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// CheckOrderStatus: Middleware Method to Check Order Status
func CheckOrderStatus(orderId string) (string, error) {
	// Creating an HTTP client
	client := http.Client{}

	orderUrl := os.Getenv("JUSPAY_SERVER_URL") + "/orders/" + orderId
	merchantId := os.Getenv("MERCHANT_ID")
	currentDate := "2024-01-07"
	juspayAPIKey := os.Getenv("JUSPAY_API_KEY")
	fmt.Println(orderUrl, merchantId, currentDate, juspayAPIKey)
	// Setting headers to the request
	req, err := http.NewRequest("GET", orderUrl, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Setting the values of Headers
	req.Header.Set("version", currentDate)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-merchantid", merchantId)

	// Setting up Basic Authentication
	req.SetBasicAuth(juspayAPIKey, "")

	// Making the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
