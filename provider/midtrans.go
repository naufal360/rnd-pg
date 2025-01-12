package provider

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"payment-gateway/config"
	"time"
)

func NewMidtrans() *Midtrans {
	return &Midtrans{
		// client: client,
	}
}

type Midtrans struct {
	// client *http.Client
}

type MidtransRequest struct {
	TransactionDetails TransactionDetails `json:"transaction_details"`
	CustomerDetails    *CustomerDetails   `json:"customer_details,omitempty"`
}

type MidtransErrorResponse struct {
	ErrorMessages []string `json:"error_messages"`
	Locale        string   `json:"locale,omitempty"`
}

type MidtransResponse struct {
	Token      string                 `json:"token,omitempty"`
	PaymentURL string                 `json:"redirect_url,omitempty"`
	Error      *MidtransErrorResponse `json:"error,omitempty"`
}

type TransactionDetails struct {
	OrderID     string `json:"order_id"`
	GrossAmount int    `json:"gross_amount"`
}

type CustomerDetails struct {
	FirstName      string          `json:"first_name"`
	LastName       string          `json:"last_name"`
	Email          string          `json:"email"`
	Phone          string          `json:"phone,omitempty"`
	BillingAddress *BillingAddress `json:"billing_address,omitempty"`
}

type BillingAddress struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone,omitempty"`
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

type MidtransInterface interface {
	SendPayment(payload MidtransRequest) (MidtransResponse, time.Time, error)
}

func (m *Midtrans) SendPayment(payload MidtransRequest) (MidtransResponse, time.Time, error) {
	env := config.GetEnv()

	url := fmt.Sprintf("%s/transactions", env.PAYMENT_HOST)
	paymentServerKeyStr := fmt.Sprintf("%s:", env.PAYMENT_SERVER_KEY)
	paymentServerKey := base64.StdEncoding.EncodeToString([]byte(paymentServerKeyStr))
	var result MidtransResponse

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return result, time.Now(), err
	}

	payloadReader := io.NopCloser(bytes.NewReader(payloadJSON))
	req, err := http.NewRequest(http.MethodPost, url, payloadReader)
	if err != nil {
		return result, time.Now(), err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Basic %s", paymentServerKey))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, time.Now(), err
	}
	createdTransactionTime := time.Now()

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return result, time.Now(), err
	}

	if res.StatusCode != http.StatusCreated {
		var errorResponse MidtransErrorResponse
		err := json.Unmarshal(body, &errorResponse)
		if err != nil {
			return result, time.Now(), err
		}
		result = MidtransResponse{
			Error: &errorResponse,
		}
		return result, time.Now(), fmt.Errorf("error response: %s", errorResponse.ErrorMessages)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, time.Now(), err
	}

	return result, createdTransactionTime, nil
}
