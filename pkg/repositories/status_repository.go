package repositories

import (
	"github.com/roger-king/gh-template/pkg/config"
	"github.com/roger-king/gh-template/pkg/entities"
)

// StatusRepository -
type StatusRepository struct {
	Version     string `json:"version"`
	Environment string `json:"environment"`
}

// StatusRepositoryIface -
type StatusRepositoryIface interface {
	GetApplicationStatus() *entities.Status
}

// NewStatusRepository - initializes the status service
func NewStatusRepository(c *config.Config) *StatusRepository {
	return &StatusRepository{
		Version:     c.Version,
		Environment: c.GoEnv,
	}
}

// GetApplicationStatus - Fetches for the health status of our application
func (s *StatusRepository) GetApplicationStatus() *entities.Status {
	return &entities.Status{
		Status:      "OK",
		Version:     s.Version,
		Environment: s.Environment,
	}
}
