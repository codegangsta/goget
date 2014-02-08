package main

import (
	"github.com/codegangsta/goget/packages"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Extensions: []string{".html"},
	}))

	m.Get("/:package", packages.ServeMaster)
	m.Get("/:package/:version", packages.ServePackage)

	m.Run()
}
