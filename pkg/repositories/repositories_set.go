package repositories

import "github.com/google/wire"

// Set - Creates wire dependencies for repositories and interacting with your datastores
var Set = wire.NewSet(
	NewStatusRepository,
	wire.Bind(new(StatusRepositoryIface), new(*StatusRepository)),
)
