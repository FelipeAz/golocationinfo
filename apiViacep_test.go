package locationhelper

import (
	"testing"
)

const defaultErrorMsg = "Expecting: %v. Found: %v."

func TestGetAPIData(t *testing.T) {
	tests := []struct {
		cep      string
		expected string
	}{
		{"17123082", "17123-082"},
		{"17123-082", "17123-082"},
		{"17015300", "17015-300"},
		{"", ""},
		{"11111111", ""},
		{"11111-111", ""},
	}

	for _, test := range tests {
		t.Logf("Mass: %v", test)
		result, _ := getAPIData(test.cep)

		if result.Cep != test.expected {
			t.Errorf(defaultErrorMsg, test.expected, result.Cep)
		}
	}
}
