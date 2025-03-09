package simulator

import "testing"

func TestGetConfiguredResponse(t *testing.T) {
	// todo: improve testing here
	expectedEndpoint := "/foo"
	cr := GetConfiguredResponse(expectedEndpoint)
	if cr.Endpoint == nil {
		t.Errorf("Failed to get expected endpoint")
	}
	expectedEndpoint = "/doesnt-exist"
	cr = GetConfiguredResponse(expectedEndpoint)
	if cr.Endpoint != nil {
		t.Errorf("Incorrectly matched request %s to input %s", cr.Endpoint.Path, expectedEndpoint)
	}
}
