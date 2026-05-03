package tests

import (
	"app/config"
	"app/database"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDB_Healhcheck(t *testing.T) {
	t.Run("database is reachable", func(t *testing.T) {
		sqlDB, err := testEnv.DB.Connection.DB()
		require.NoError(t, err)

		require.NoError(t, sqlDB.Ping())
	})
}

func TestDB_Connect(t *testing.T) {
	t.Run("invalid dsn", func(t *testing.T) {
		cfg := config.Config{
			PostgreSQL: config.PostgreSQL{
				User:     "test",
				Password: "strong",
				DBname:   "none",
				Host:     "my_own",
				Port:     "0000",
				SslMode:  "ignore",
			},
		}

		_, err := database.Connect(cfg)

		require.Error(t, err)
	})
}
