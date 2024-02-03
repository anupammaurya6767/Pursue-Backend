package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

// CreateSession: Middleware function to create a payment session
func CreateSession(orderId string, amount int, customerCredentials types.JuspayCustomer) (string, error) {
	// Create a new Http client
	client := &http.Client{}

	// Defining our Request parameters
	requestUrl := os.Getenv("JUSPAY_SERVER_URL") + "/session"
	requestData := map[string]interface{}{
		"order_id":               orderId,
		"amount":                 strconv.Itoa(amount),
		"customer_id":            *customerCredentials.CustomerId,
		"customer_email":         *customerCredentials.EmailAddress,
		"customer_phone":         strconv.Itoa(customerCredentials.MobileNumber),
		"payment_page_client_id": "pursue",
		"action":                 "paymentPage",
	}

	// Serialising the above request body to json
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("version", "2024-01-07")
	req.Header.Set("x-merchantid", os.Getenv("MERCHANT_ID"))

	req.SetBasicAuth(os.Getenv("JUSPAY_API_KEY"), "")
	// Maing a post request programmatically
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// Reading the request body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(body))
	return string(body), nil
}
