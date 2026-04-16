package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		key     string
		value   string
		want    string
		wantErr string
	}{
		{
			wantErr: "no authorization header included",
		},
		{
			key:     "Authorization",
			value:   "not a valid api key",
			wantErr: "malformed authorization header",
		},
		{
			key:   "Authorization",
			value: "ApiKey 1234567890",
			want:  "1234567890",
		},
		{
			key:     "Authorization",
			value:   "ApiKey",
			wantErr: "malformed authorization header",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetApiKey Case #%v:", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			got, err := GetAPIKey(header)
			if err != nil {
				if err.Error() != test.wantErr {
					t.Errorf("Unexpected: TestGetApiKey: %v\n", err)
					return
				}
				return
			}

			if got != test.want {
				t.Errorf("Unexpected: TestGetApiKey:%s", got)
				return
			}
		})
	}
}
