package auth // first, declare name of package you want to test

import (
	"net/http"
	"testing" // import testing package
)

func TestGetAPIKey(t *testing.T) { // by convention, Test + function name
	dummyHeader := http.Header{}
	dummyHeader.Add("Authorization", "ApiKey 123456")
	gotKey, gotErr := GetAPIKey(dummyHeader)
	wantKey := "123456"
	if gotKey != wantKey {
		t.Fatalf("got API key %s, want API key %s", gotKey, wantKey)
	}
	if gotErr != nil {
		t.Fatalf("got error %v, want nil", gotErr)
	}
}
