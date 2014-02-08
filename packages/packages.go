package packages

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
)

// Packages is a map of packages to github repos. For now this is
// a variable but eventually it should be in a configuration file
// or database.
var packages []Package = []Package{
	{"martini", "codegangsta/martini"},
}

type Package struct {
	Name string
	Repo string
}

func FindPackage(name string) Package {
	for _, p := range packages {
		if p.Name == name {
			return p
		}
	}

	return Package{}
}

func ServeMaster(params martini.Params, r render.Render, req *http.Request) {
	p := FindPackage(params["package"])
	if len(p.Name) < 1 {
		r.Error(404)
		return
	}

	r.HTML(200, "master", struct {
		Host    string
		Package Package
	}{req.URL.Host, p})
}
