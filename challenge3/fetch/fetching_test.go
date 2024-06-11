package fetch

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFetchMeatText(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test meat text"))
	}))
	defer ts.Close()

	//Setenv
	os.Setenv("API_URL", ts.URL)

	// call function
	text, err := FetchMeatText()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// verify got กับ expected
	expected := "Test meat text"
	if text != expected {
		t.Fatalf("Expected %s, got %s", expected, text)
	}
}
