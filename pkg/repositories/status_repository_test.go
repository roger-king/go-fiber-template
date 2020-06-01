package repositories

import (
	"testing"

	"github.com/roger-king/gh-template/pkg/config"
	"github.com/stretchr/testify/suite"
)

var (
	expectedVersion = "v1.0.0"
	expectedEnv     = "test"
)

type StatusRepositorySuite struct {
	suite.Suite
	svc *StatusRepository
}

func (s *StatusRepositorySuite) SetupTest() {
	s.svc = NewStatusRepository(&config.Config{
		Version: expectedVersion,
		GoEnv:   expectedEnv,
	})
}

func (s *StatusRepositorySuite) TestGetApplicationStatus_Success() {
	status := s.svc.GetApplicationStatus()
	s.Equal("OK", status.Status)
	s.Equal(expectedVersion, status.Version)
	s.Equal(expectedEnv, status.Environment)
}

func TestStatusRepositorySuite(t *testing.T) {
	suite.Run(t, new(StatusRepositorySuite))
}
