package hello

import (
	"fmt"
	"net/http"

	_ "code.google.com/p/goauth2/oauth"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
