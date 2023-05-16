package service

import (
	"labora-api/API/model"
	"testing"
)

type totalPriceTest struct {
	arg1     int
	arg2     string
	arg3     string
	arg4     string
	arg5     int
	arg6     int
	expected int
}

var priceTests = []totalPriceTest{
	{0, "", "", "", 3, 1000, 3000},
	{0, "", "", "", 4, 2000, 8000},
	{0, "", "", "", 1, 2000, 2000},
	{0, "", "", "", 9, 2000, 18000},
	{0, "", "", "", 4, 6000, 24000},
}

func TestNewItem(t *testing.T) {
	for _, test := range priceTests {
		if output := model.NewItem(test.arg1, test.arg2, test.arg3, test.arg4, test.arg5, test.arg6); output.TotalPrice != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
