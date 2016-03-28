package main

import (
	"github.com/zenazn/goji"
	"net/http"
	"github.com/zenazn/goji/web"
	"fmt"
	"flag"
)

func main() {
	flag.Set("bind", ":80")

	goji.Get("/", test)
	goji.Serve()
}

func test(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, it works!")
}

