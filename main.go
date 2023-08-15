package main

// $ curl http://localhost:9999/
// URL.Path = "/"
// $ curl http://localhost:9999/hello
// Header["Accept"] = ["*/*"]
// Header["User-Agent"] = ["curl/7.54.0"]
// curl http://localhost:9999/world
// 404 NOT FOUND: /world

import (
	"minigin"
	"net/http"
)

func main() {
	e := minigin.New()
	e.GET("/", func(c *minigin.Context) {
		c.String(http.StatusOK, "URL.Path = %q\n", c.Path)
	})
	e.GET("/test", func(c *minigin.Context) {
		c.String(http.StatusOK, "URL.Path = %q\n", c.Path)
	})
	e.Run(":8080")
}
