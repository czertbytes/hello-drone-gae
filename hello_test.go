package hello

import (
	"net/http"
	"testing"

	"github.com/rcrowley/go-tigertonic/mocking"
)

func TestHello(t *testing.T) {
	code, _, response, err := hello(
		mocking.URL(mux, "GET", "/1.0/hello"),
		mocking.Header(nil),
		nil,
	)

	if nil != err {
		t.Fatal(err)
	}

	if http.StatusOK != code {
		t.Fatal(code)
	}

	if "Hello Saki" != response.Message {
		t.Fatal(response)
	}
}
