package mapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCountryCodeForTimezone(t *testing.T) {
	testCases := []struct {
		name                string
		inputTimeone        string
		expectedCountryCode string
		expectedError       string
	}{
		{
			name:                "Should return country code",
			inputTimeone:        "Asia/Kolkata",
			expectedCountryCode: "IN",
		},
		{
			name:          "Should return error",
			inputTimeone:  "random",
			expectedError: "Country-Code not found for random",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			code, err := FindCountryCodeForTimezone(tc.inputTimeone)
			if tc.expectedError != "" {
				assert.EqualError(t, err, tc.expectedError)
			} else {
				assert.Equal(t, tc.expectedCountryCode, code)
			}
		})
	}
}
