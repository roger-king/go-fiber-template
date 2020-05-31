// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package pkg

import (
	"github.com/gofiber/fiber"
	"github.com/google/wire"
)

var ServerSet = wire.NewSet(AppSet)

func New() (*fiber.App, error) {
	wire.Build(ServerSet)
	return nil, nil
}
