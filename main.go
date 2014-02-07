package main

import (
	"github.com/codegangsta/martini"
	"github.com/yvasiyarov/gorelic"
	"os"
)

func main() {
	// setup newrelic
	if martini.Env == martini.Prod {
		agent := gorelic.NewAgent()
		agent.Verbose = true
		agent.NewrelicLicense = os.Getenv("NEW_RELIC_LICENSE")
		agent.Run()
	}

	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello World!"
	})

	m.Run()
}
