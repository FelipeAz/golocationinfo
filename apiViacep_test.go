package locationhelper

import (
	"testing"
)

const defaultErrorMsg = "Expecting: %v. Found: %v."

func TestGetAPIData(t *testing.T) {
	tests := []struct {
		cep         string
		cepExpected string
	}{
		{cep: "17123082", cepExpected: "17123-082"},
		{cep: "17123-082", cepExpected: "17123-082"},
		{cep: "17015300", cepExpected: "17015-300"},
		{cep: "", cepExpected: ""},
		{cep: "11111111", cepExpected: ""},
		{cep: "11111-111", cepExpected: ""},
	}

	for _, test := range tests {
		t.Logf("Mass: %v", test)
		result, _ := getAPIData(test.cep)

		if result.Cep != test.cepExpected {
			t.Errorf(defaultErrorMsg, test.cepExpected, result.Cep)
		}
	}
}
