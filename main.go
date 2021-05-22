package main

import (
	"time"

	"github.com/bifrost-inc/Heimdall/core"
)

func main() {
	mainAPi := core.Api{}
	mainAPi.Title = "my first api"
	mainAPi.Backend = core.Backend{Path: "http://backend:8080/",
		HealthcheckSettings: core.HealthcheckDefinition{
			Interval:   time.Now(),
			MethodUsed: healthcheck.Ping},
	}
}
