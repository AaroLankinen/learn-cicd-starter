package auth

import (
	"errors"
	"net/http"
	"testing"
)

// TestGetAPIKey tests the GetAPIKey function with various Authorization header scenarios.
func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectedErr error
		errContains string
	}{
		{
			name:        "Valid ApiKey header",
			headers:     http.Header{"Authorization": []string{"ApiKey secret123"}},
			expectedKey: "secret123",
			expectedErr: nil,
		},
		{
			name:        "No Authorization header",
			headers:     http.Header{},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:        "Malformed Authorization header - wrong prefix",
			headers:     http.Header{"Authorization": []string{"Bearer secret123"}},
			expectedKey: "",
			errContains: "malformed authorization header",
		},
		{
			name:        "Malformed Authorization header - no key",
			headers:     http.Header{"Authorization": []string{"ApiKey"}},
			expectedKey: "",
			errContains: "malformed authorization header",
		},
		{
			name:        "Malformed Authorization header - empty header",
			headers:     http.Header{"Authorization": []string{""}},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.headers)
			if tt.expectedErr != nil {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("GetAPIKey() error = %v, want expectedErr %v", err, tt.expectedErr)
				}
			} else if tt.errContains != "" {
				if err == nil || err.Error() != tt.errContains {
					t.Errorf("GetAPIKey() error = %v, want errContains %q", err, tt.errContains)
				}
			} else if err != nil {
				t.Errorf("GetAPIKey() unexpected error = %v", err)
			}

			if gotKey != tt.expectedKey {
				t.Errorf("GetAPIKey() gotKey = %v, want expectedKey %v", gotKey, tt.expectedKey)
			}
		})
	}
}
