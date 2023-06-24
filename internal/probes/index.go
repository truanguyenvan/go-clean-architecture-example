package probes

import "github.com/google/wire"

var Set = wire.NewSet(
	NewHealthChecker,
)
