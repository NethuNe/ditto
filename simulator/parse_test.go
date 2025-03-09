package simulator

import "testing"

func TestParseConfig(t *testing.T) {
	tcs := []struct {
		testName             string
		numExpectedEndpoints int
		wantErr              bool
		filePath             string
	}{
		{
			testName:             "Valid Schema Single Endpoint",
			numExpectedEndpoints: 1,
			wantErr:              false,
			filePath:             "simulator/testSchema.json",
		},
		{
			testName:             "Invalid Filepath",
			numExpectedEndpoints: 0,
			wantErr:              true,
			filePath:             "./does-not-exist.naur",
		},
	}
	for _, tc := range tcs {
		out, outErr := ParseConfig(tc.filePath)
		if outErr != nil && !tc.wantErr {
			t.Errorf("Test Case [%s]: Failed with unexpected error: %s", tc.testName, outErr.Error())
		} else if outErr == nil && tc.wantErr {
			t.Errorf("Test Case [%s]: Failed to throw expected error with output %v", tc.testName, *out)
		}
		// Don't try to use the returned config if there's an error.
		if outErr == nil {
			if len(out.Endpoints) != tc.numExpectedEndpoints {
				t.Errorf("Test Case [%s]: Returned config with wrong number of endpoints; expected %d, got %d",
					tc.testName, tc.numExpectedEndpoints, len(out.Endpoints))
			}
		}
	}
}
