package app

import "github.com/google/wire"

var Set = wire.NewSet(
	NewApplication,
)
