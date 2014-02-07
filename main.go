package main

import (
	"github.com/codegangsta/martini"
	"github.com/yvasiyarov/gorelic"
	"log"
	"os"
)

func main() {
	// setup newrelic
	agent := gorelic.NewAgent()
	agent.Verbose = true
	agent.NewrelicLicense = os.Getenv("NEW_RELIC_LICENSE")
	err := agent.Run()
	if err != nil {
		log.Fatal(err)
	}

	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello World!"
	})

	m.Run()
}
