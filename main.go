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
	e.GET("/aa/*dwdw/dwd", func(c *minigin.Context) {
		c.String(http.StatusOK, "URL.Path = %q\n", c.Path)
	})
	// e.GET("/test/hello/world", func(c *minigin.Context) {
	// 	c.String(http.StatusOK, "URL.Path = %q\n", c.Path)
	// })
	// e.GET("/test/hello/dw", func(c *minigin.Context) {
	// 	c.String(http.StatusOK, "URL.Path = %q\n", c.Path)
	// })
	// e.GET("/test/hello/world/this", func(c *minigin.Context) {
	// 	c.String(http.StatusOK, "URL.Path = %q\n", c.Path)
	// })
	// e.GET("/ee", func(c *minigin.Context) {
	// 	c.String(http.StatusOK, "URL.Path = %q\n", c.Path)
	// })
	// e.GET("/hello", func(c *minigin.Context) {
	// 	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	// })
	// e.POST("/login", func(c *minigin.Context) {
	// 	c.JSON(http.StatusOK, minigin.H{
	// 		"username": c.PostForm("username"),
	// 		"password": c.PostForm("password"),
	// 	})
	// })

	e.Run(":8080")
}
