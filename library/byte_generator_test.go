package library

import "testing"

func TestGenerateRawByteData(t *testing.T) {
	tcs := []struct {
		name            string
		wantErr         bool
		numBytes        int
		expectedDataLen int
	}{
		{
			name:            "base case",
			wantErr:         false,
			numBytes:        500,
			expectedDataLen: 500,
		},
		{
			name:            "large base case",
			wantErr:         false,
			numBytes:        1e5,
			expectedDataLen: 1e5,
		},
		{
			name:            "larger than cap case",
			wantErr:         false,
			numBytes:        1e9,
			expectedDataLen: MAX_BYTE_COUNT,
		},
		{
			name:            "invalid number of bytes",
			wantErr:         true,
			numBytes:        0,
			expectedDataLen: 0,
		},
		{
			name:            "negative number of bytes",
			wantErr:         true,
			numBytes:        -1,
			expectedDataLen: 0,
		},
	}

	for _, tc := range tcs {
		rawBytes, err := GenerateRawByteData(tc.numBytes)
		if err != nil {
			if !tc.wantErr {
				t.Errorf("Test Case [%s] failed with unwanted error %s", tc.name, err.Error())
			}
			continue
		}
		if tc.wantErr {
			t.Errorf("Test Case [%s] expected error, did not error. Returned %d bytes", tc.name, len(rawBytes))
		}
		if len(rawBytes) != tc.expectedDataLen {
			t.Errorf("Test Case [%s] failed with mismatched number of bytes. Expected %d, got %d", tc.name, tc.numBytes, len(rawBytes))
		}
	}
}
