package core

import (
	"time"

	"github.com/bifrost-inc/Heimdall/src/enumerators/healthcheck"
)

type Api struct {
	Title     string
	BaseRoute string
	Endpoints []Endpoint
	Backend   Backend
}

type Endpoint struct {
	Path            string
	AcceptedMethods []string
}

type Backend struct {
	Path                string
	HealthcheckSettings HealthcheckDefinition
}

type HealthcheckDefinition struct {
	Interval   time.Time
	MethodUsed healthcheck.Method
}
