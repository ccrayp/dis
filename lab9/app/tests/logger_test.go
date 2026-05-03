package tests

import (
	"app/internal/logger"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	db := testEnv.DB
	auditRepository := logger.NewAuditRepository(db)
	log := logger.NewLogger(auditRepository)

	require.NoError(t, log.MakeLog("tester", "test", "none", "integration logger test", logger.Info))
}
