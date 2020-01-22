package main

import (
	"github.com/NYTimes/gziphandler"
	"log"
	"net/http"
	"strings"
)

const dir = "./frontend"

func main() {
	//fs := http.FileServer(http.Dir(dir))
	fs := gziphandler.GzipHandler(http.FileServer(http.Dir(dir)))

	log.Print("Serving " + dir + " on http://localhost:8080")
	http.ListenAndServe(":8080", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		fs.ServeHTTP(resp, req)
	}))
}
