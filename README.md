# Minigin

Minigin: A lightweight HTTP web framework written in Golang, inspired by the popular [Gin framework](https://gin-gonic.com/).

Check [here](#reference) for the reference of this codebase.

## Table of Contents 

- [Description](#description)
- [Installation](#installation)
- [Usage](#usage)
- [Todo](#todo)
- [Reference](#reference)

## Description

Minigin is a lightweight HTTP web framework written in Golang. Its design and implementation follow the approach of the popular [Gin framework](https://gin-gonic.com/). Key features implemented in Minigin include:
- Gin-like easy-to-use features with context packaged.
```go
	e := minigin.New()
	e.GET("/", func(c *minigin.Context) {
		c.String(http.StatusOK, "Hi there")
	})
    e.Run(":8080")
```
- Dynamic routing implemented with a Trie data structure.
```go
    e := minigin.New()
	e.GET("/index", func(c *minigin.Context) {
		c.String(http.StatusOK, "Hi there")
	})
	e.GET("/index/*name", func(c *minigin.Context) {
		c.String(http.StatusOK, "Your name = %q\n", c.Paras["*name"])
	})
    e.Run(":8080")
```
- Support for responses in HTML/JSON formats and query parameters.
```go
	e := minigin.New()
	e.GET("/", func(c *minigin.Context) {
		c.HTML(http.StatusOK, "<h1>Hi there</h1>")
	})
    e.GET("/index", func(c *minigin.Context) {
		c.JSON(http.StatusOK, minigin.H{
			"foo": "bar",
		})
	})
    e.GET("/*type", func(c *minigin.Context) {
		c.String(http.StatusOK, "Your name = %q\n", c.Paras["*type"])
	})
    e.Run(":8080")
```
- Route grouping.
```go
    e := minigin.New()
	grp := e.NewRouterGroup("/admin")
	grp.GET("/*type", func(c *minigin.Context) {
		c.String(http.StatusOK, "Your name = %q\n", c.Paras["*type"])
	})
    e.Run(":8080") // try with `curl http://localhost:8080/admin/yourname`
```
- Middleware integraton.
```go
    e := minigin.New()
    e.RegisterMiddleware(minigin.Logger())
    e.Run(":8080")
```
- Proper panic handling strategy.
```go
    e := minigin.New()
    e.RegisterMiddleware(minigin.PanicRecovery(), minigin.Logger())
    e.Run(":8080")
```

## Installation

1. Clone the repository: `git clone git@github.com:qhzhu/minigin-web.git`
2. Copy `minigin` to your working directory
3. Start working with `mingin` by 
```go
    package yourproject

    import (
        "minigin"
    )
    
    func main() {
        engine := minigin.Default()
        engine.GET("/", func(c *minigin.Context) {
            c.HTML(http.StatusOK, "<h1>Hello Minigin!</h1>")
        })
        engine.Run(":8080")
    }
```

## Usage
<a id="usage"></a>
Here's how you can use Minigin...(to be finished)

## Todo
- [ ] Adding comments to the codebase
- [ ] Implementing static resources and rendering
- [ ] Adding [Usage](#usage) information 
- [ ] Changing the manner of installation to `go get github.com/qhzhu/minigin`
## Reference
- [7days-golang](https://geektutu.com/post/gee.html)
