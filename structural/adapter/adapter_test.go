// Copyright (c) 2018 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package adapter

import "testing"

func TestAdapter(t *testing.T)  {
	msg := "Hello World"
	adapter := PrinterAdapter{
		OldPrinter: &MyLegacyPrinter{},
		Msg: msg,
	}

	returnedMsg := adapter.PrintStored()
	if returnedMsg != "Legacy printer: Adapter: Hello World\n" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}

	adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg = adapter.PrintStored()

	if returnedMsg != "Hello World" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}
}
