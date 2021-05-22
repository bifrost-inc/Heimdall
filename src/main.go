package main

import (
	"time"

	"github.com/bifrost-inc/Heimdall/src/core"
        "github.com/bifrost-inc/Heimdall/src/enumerators/heathcheck"
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
