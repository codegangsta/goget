package packages

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"os"
)

var host = os.Getenv("GOGET_HOST")

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

func ServeMaster(params martini.Params, r render.Render) {
	p := FindPackage(params["package"])
	if len(p.Name) < 1 {
		r.HTML(404, "error", "Package not found")
		return
	}

	r.HTML(200, "master", struct {
		Host    string
		Package Package
	}{host, p})
}

func ServePackage(params martini.Params, r render.Render) {
	p := FindPackage(params["package"])
	t, err := GetTag(p.Repo, params["version"])
	if err != nil {
		r.HTML(500, "error", err)
		return
	}

	r.HTML(200, "package", struct {
		Host    string
		Package Package
		Tag     Tag
	}{host, p, t})
}
