package simulator

import "testing"

func TestGetConfiguredResponse(t *testing.T) {
	// todo: improve testing here
	tcs := []struct {
		name             string
		matchEndpoint    string
		expectedEndpoint string
		expectMatch      bool
	}{
		{
			name:             "exact match",
			matchEndpoint:    "/foo",
			expectedEndpoint: "/foo",
			expectMatch:      true,
		},
		{
			name:             "fuzzy match hit",
			matchEndpoint:    "/fo",
			expectedEndpoint: "/foo",
			expectMatch:      true,
		},
		{
			name:             "fuzzy match miss",
			matchEndpoint:    "/doesnt-exist",
			expectedEndpoint: "",
			expectMatch:      false,
		},
		{
			name:             "no match",
			matchEndpoint:    "doesnt/exist",
			expectedEndpoint: "",
			expectMatch:      false,
		},
	}
	for _, tc := range tcs {
		cr := GetConfiguredResponse(tc.matchEndpoint)
		if cr.Endpoint == nil {
			if tc.expectMatch {
				t.Errorf("Test Name [%s] expected match for endpoint %s on %s but no match was found", tc.name, tc.matchEndpoint, tc.expectedEndpoint)
			}
			continue
		}
		if !tc.expectMatch && cr.Endpoint != nil {
			t.Errorf("Test Name [%s] expected no match for endpoint %s but matched on %s", tc.name, tc.matchEndpoint, cr.Endpoint.Path)
		}
		if tc.expectMatch && tc.expectedEndpoint != cr.Endpoint.Path {
			t.Errorf("Test Name [%s] expected match for endpoint %s on %s but matched on %s", tc.name, tc.matchEndpoint, tc.expectedEndpoint, cr.Endpoint.Path)
		}
	}
}
