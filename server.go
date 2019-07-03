package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	dir  = "./docs"
	port = 8080
)

func main() {
	fs := http.FileServer(http.Dir(dir))
	log.Printf("Serving %s on http://localhost:%d", dir, port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		fs.ServeHTTP(resp, req)
	}))
}
