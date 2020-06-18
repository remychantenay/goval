package goval

import (
	"testing"
)

// TestStruct represents the struct under test
type TestStruct struct {
	Value string `goval:"country_code,required=true,exclude=US,excludeEU=true"`
}

func TestValidateStruct_Err(t *testing.T) {
	tests := []struct {
		description      string
		expectedErrorMsg string
		with             TestStruct
	}{
		{
			description:      "Too short with constraint",
			expectedErrorMsg: "Value should be 2 characters long",
			with: TestStruct{
				Value: "E",
			},
		},
		{
			description:      "Too long with constraint",
			expectedErrorMsg: "Value should be 2 characters long",
			with: TestStruct{
				Value: "EEEE",
			},
		},
		{
			description:      "Invalid",
			expectedErrorMsg: "Value is an invalid country code",
			with: TestStruct{
				Value: "RR",
			},
		},
		{
			description:      "Empty",
			expectedErrorMsg: "Value cannot be blank",
			with: TestStruct{
				Value: "",
			},
		},
		{
			description:      "Country excluded",
			expectedErrorMsg: "Value is excluded",
			with: TestStruct{
				Value: "US",
			},
		},
	}

	for _, test := range tests {
		err := ValidateStruct(test.with)[0]
		if err == nil {
			t.Fatalf("%s -> was expecting %s", test.description, test.expectedErrorMsg)
		}

		if err.Error() != test.expectedErrorMsg {
			t.Fatalf("%s -> Got %s but expected %s ", test.description, err.Error(), test.expectedErrorMsg)
		}
	}
}

func TestValidateStruct_Success(t *testing.T) {
	tests := []struct {
		description string
		with        TestStruct
	}{
		{
			description: "Success with constraint",
			with: TestStruct{
				Value: "CH",
			},
		},
	}

	for _, test := range tests {
		errs := ValidateStruct(test.with)
		if errs != nil && len(errs) > 0 {
			t.Fatalf("%s -> was not expecting error but got %s", test.description, errs[0].Error())
		}
	}
}

func BenchmarkValidateStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarks := []struct {
			description string
			with        TestStruct
		}{
			{
				description: "Success with constraint",
				with: TestStruct{
					Value: "CH",
				},
			},
		}

		for _, benchmark := range benchmarks {
			ValidateStruct(benchmark.with)
		}
	}
}
