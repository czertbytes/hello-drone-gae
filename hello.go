package hello

import (
    "log"
    "net/http"
    "net/url"

    "github.com/rcrowley/go-tigertonic"
)

var (
    mux, nsMux *tigertonic.TrieServeMux
)

type HelloResponse struct {
    Message string `json:"message"`
}

func init() {
    log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)

    cors := tigertonic.NewCORSBuilder().AddAllowedOrigins("*")

    mux = tigertonic.NewTrieServeMux()
    mux.Handle("GET", "/hello", cors.Build(tigertonic.Timed(tigertonic.Marshaled(hello), "hello", nil)))

    nsMux = tigertonic.NewTrieServeMux()
    nsMux.HandleNamespace("", mux)
    nsMux.HandleNamespace("/1.0", mux)

    http.Handle("/", nsMux)
}

func hello(u *url.URL, h http.Header, _ interface{}) (int, http.Header, *HelloResponse, error) {
    return http.StatusOK, nil, &HelloResponse{"Saki Saki Saki Saki"}, nil
}
