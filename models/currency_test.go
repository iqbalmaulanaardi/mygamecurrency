package models

import "testing"

var (
	currency Currency
)

func TestValidateCurrency(t *testing.T) {
	t.Logf("Validate : %s", currency.Validate())
	if currency.Validate().Error() != "Invalid Param(s)" {
		t.Errorf("should be  %s", "Invalid Param(s)")
	}
}
