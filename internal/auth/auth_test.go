package auth

import "testing"

func TestGetApiKey(t *testing.T) {
	type test struct {
		name     string
		headers  map[string][]string
		expected string
		err      error
	}

	tests := []test{
		{
			name:     "no auth header",
			headers:  map[string][]string{},
			expected: "",
			err:      ErrNoAuthHeaderIncluded,
		},
		{
			name:     "malformed auth header",
			headers:  map[string][]string{"Authorization": {"Bearer"}},
			expected: "",
			err:      ErrMalformedAuthHeader,
		},
		{
			name:     "valid auth header",
			headers:  map[string][]string{"Authorization": {"ApiKey 123"}},
			expected: "123",
			err:      nil,
		},
		{
			name:     "invalid auth header with correct size",
			headers:  map[string][]string{"Authorization": {"Bearer 123"}},
			expected: "",
			err:      ErrMalformedAuthHeader,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tc.headers)
			if apiKey != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, apiKey)
			}
			if err != tc.err {
				t.Errorf("expected %v, got %v", tc.err, err)
			}
		})
	}
}
