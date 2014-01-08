package hello

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	tt "github.com/rcrowley/go-tigertonic"
)

var (
	mux, v1Mux *tt.TrieServeMux
)

type HelloResponse struct {
	Message string `json:"message"`
}

func init() {
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)

	cors := tt.NewCORSBuilder().AddAllowedOrigins("*")

	v1Mux = tt.NewTrieServeMux()
	v1Mux.Handle("GET", "/hello", cors.Build(tt.Timed(tt.Marshaled(hello), "hello", nil)))
	v1Mux.Handle("GET", "/hello/{id}", cors.Build(tt.Timed(tt.Marshaled(helloParam), "helloParam", nil)))
	v1Mux.Handle("GET", "/if-true", tt.If(testTrue, tt.Marshaled(ifHandler)))
	v1Mux.Handle("GET", "/if-false", tt.If(testFalse, tt.Marshaled(ifHandler)))

	mux = tt.NewTrieServeMux()
	mux.HandleNamespace("/1.0", v1Mux)

	http.Handle("/", mux)
}

func hello(u *url.URL, h http.Header, _ interface{}) (int, http.Header, *HelloResponse, error) {
	return http.StatusOK, nil, &HelloResponse{"Hello Saki"}, nil
}

func helloParam(u *url.URL, h http.Header, _ interface{}) (int, http.Header, *HelloResponse, error) {
	return http.StatusOK, nil, &HelloResponse{fmt.Sprintf("Hello %s", u.Query().Get("id"))}, nil
}

func testTrue(r *http.Request) (http.Header, error) {
	return nil, nil
}

func testFalse(r *http.Request) (http.Header, error) {
	return nil, tt.InternalServerError{errors.New("internal server error")}
}

func ifHandler(u *url.URL, h http.Header, _ interface{}) (int, http.Header, *HelloResponse, error) {
	return http.StatusOK, nil, &HelloResponse{"Saki"}, nil
}
