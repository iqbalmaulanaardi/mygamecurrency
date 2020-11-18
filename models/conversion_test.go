package models

import "testing"

var (
	conversion Conversion
)

func TestValidateConversion(t *testing.T) {
	t.Logf("Validate : %s", conversion.Validate())
	if conversion.Validate().Error() != "Invalid Param(s)" {
		t.Errorf("should be  %s", "Invalid Param(s)")
	}
}
