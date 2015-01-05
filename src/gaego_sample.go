package gaego_sample

import (
	"fmt"
	"net/http"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func init() {
	goji.Get("/hello/:name", hello)
	goji.Serve()
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}
