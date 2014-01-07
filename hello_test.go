package hello

import (
    "net/http"
    "testing"

    "github.com/rcrowley/go-tigertonic"
    "github.com/rcrowley/go-tigertonic/mocking"
)

var (
    hMux tigertonic.HostServeMux
)

func init() {
    hMux = tigertonic.NewHostServeMux()
    hMux.Handle("example.com", nsMux)
}

func TestHello(t *testing.T) {
    code, _, response, err := hello(
        mocking.URL(hMux, "GET", "http://example.com/1.0/hello"),
        mocking.Header(nil),
        nil,
    )

    if nil != err {
        t.Fatal(err)
    }

    if http.StatusOK != code {
        t.Fatal(code)
    }

    if "Saki Saki Saki Saki" != response.Message {
        t.Fatal(response)
    }
}
