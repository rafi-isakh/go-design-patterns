// Copyright (c) 2018 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package factory

import (
	"testing"
	"strings"
)

func TestGetPaymentMethodCash(t *testing.T)  {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist ")
	}

	msg := payment.Pay(10.45)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message is wrong")
	}
	t.Log("Log: ", msg)
}


func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("A payment method of type 'DebitCard' must exist ")
	}

	msg := payment.Pay(10.45)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message is wrong")
	}
	t.Log("Log: ", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := GetPaymentMethod(20)

	if err == nil {
		t.Error("Payment method not found")
	}
	t.Log("LOG: ", err)
}