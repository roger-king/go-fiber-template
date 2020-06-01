package config

import "github.com/google/wire"

// Set - Creates configuration set
var Set = wire.NewSet(NewConfig)
