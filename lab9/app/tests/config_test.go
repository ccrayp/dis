package tests

import (
	"app/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_LoadConfig(t *testing.T) {
	t.Run("invalid path", func(t *testing.T) {
		_, err := config.LoadConfig("invalid/path")

		require.Error(t, err)
	})
}
