package pkg

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/google/wire"
	"github.com/roger-king/gh-template/pkg/handlers"
	"github.com/roger-king/gh-template/pkg/repositories"
	log "github.com/sirupsen/logrus"
)

// AppSet -
var AppSet = wire.NewSet(NewApp, repositories.Set, handlers.Set)

// NewApp -
func NewApp(status *handlers.StatusHandler) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	api := app.Group("/api")
	// Server status endpoint - sanity check that the server is running
	statusGroup := api.Group("/status")
	statusGroup.Get("/", status.Get)

	log.Info("Application is running...")
	return app
}
