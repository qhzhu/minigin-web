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
	e.GET("/dd", func(c *minigin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gin</h1>")
	})
	e.GET("/aa", func(c *minigin.Context) {
		c.String(http.StatusOK, "URL.Path = %q\n", c.Path)
	})
	e.GET("/aa/*name/dwd", func(c *minigin.Context) {
		c.String(http.StatusOK, "Your name = %q\n", c.Paras["*name"])
	})
	grp1 := e.NewRouterGroup("/admin")
	grp1.GET("/*type", func(c *minigin.Context) {
		c.String(http.StatusOK, "Your name = %q\n", c.Paras["*type"])
	})
	e.Run(":8080")
}
