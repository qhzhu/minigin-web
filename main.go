package main

// $ curl http://localhost:9999/
// URL.Path = "/"
// $ curl http://localhost:9999/hello
// Header["Accept"] = ["*/*"]
// Header["User-Agent"] = ["curl/7.54.0"]
// curl http://localhost:9999/world
// 404 NOT FOUND: /world

import (
	"fmt"
	"minigin"
	"net/http"
)

func main() {
	e := minigin.New()
	e.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	e.Run(":8080")
}
