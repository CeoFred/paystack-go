package paystack

import (
	"fmt"
	"testing"
)

func TestRefundCRUD(t *testing.T) {
	fullrefund := &RefundRequest{
		Transaction: "018d3146-d22c-7e5b-93aa-2c61344950f8--018d3147-57fa-7590-9ce0-94a1ecfc0255",
	}

	txn, err := c.Refund.Create(fullrefund)
	if err != nil {
		t.Errorf("UnExpected CREATE Refund, got %+v", err)
	}
	fmt.Println("refund ID: ", txn.ID)

}
