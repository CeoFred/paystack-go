package paystack

import "fmt"

// RefundsService handles operations related to user refunds
// For more details see https://paystack.com/docs/payments/refunds/#create-a-refund
type RefundsService service

// Refund represents a Paystack refund
// For more details see https://paystack.com/docs/payments/refunds/#create-a-refund
type Refund struct {
	Integration    int         `json:"integration"`
	DeductedAmount int         `json:"deducted_amount"`
	Channel        string      `json:"channel"`
	MerchantNote   string      `json:"merchant_note"`
	CustomerNote   string      `json:"customer_note"`
	Status         string      `json:"status"`
	RefundedBy     string      `json:"refunded_by"`
	ExpectedAt     string      `json:"expected_at"`
	Currency       string      `json:"currency"`
	Domain         string      `json:"domain"`
	Amount         int         `json:"amount"`
	FullyDeducted  bool        `json:"fully_deducted"`
	ID             int         `json:"id"`
	CreatedAt      string      `json:"createdAt"`
	UpdatedAt      string      `json:"updatedAt"`
}

// RefundRequest represents a Paystack refund request
type RefundRequest struct {
	Transaction string `json:"transaction,omitempty"`
	Amount      string `json:"amount,omitempty"`
}

// RefundRequestList is a list object for refund.
type RefundRequestList struct {
	Meta   ListMeta
	Values []Refund `json:"data"`
}

// Create creates a refund 
// For more details see https://paystack.com/docs/payments/refunds/#create-a-refund
func (s *RefundsService) Create(refund *RefundRequest) (*Refund, error) {
	u := "/refund"
	r := &Refund{}
	fmt.Println(refund,r,u)
	err := s.client.Call("POST", u, refund, r)
	return r, err
}


// List returns a list of refunds.
// For more details see https://paystack.com/docs/payments/refunds/#list-refunds
func (s *RefundsService) List() (*RefundRequestList, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of refunds
// For more details see https://paystack.com/docs/payments/refunds/#list-refunds
func (s *RefundsService) ListN(count, offset int) (*RefundRequestList, error) {
	u := paginateURL("/refund", count, offset)
	sub := &RefundRequestList{}
	err := s.client.Call("GET", u, nil, sub)
	return sub, err
}

