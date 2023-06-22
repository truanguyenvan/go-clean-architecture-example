package router

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewCragRouter,
)
