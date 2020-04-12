package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfigJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		filepath := "../../config.json"

		err, config := readConfigJSON(filepath)

		assert.NoError(t, err)
		assert.NotEmpty(t, config)
	})

	t.Run("wrong filePath", func(t *testing.T) {
		filepath := ""

		err, config := readConfigJSON(filepath)

		assert.Error(t, err)
		assert.Empty(t, config)
	})
}

func TestLoad(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		filepath := "../../config.json"

		err, config := Load(filepath)

		assert.NoError(t, err)
		assert.NotEmpty(t, config)
	})

	t.Run("wrong filePath", func(t *testing.T) {
		filepath := ""

		err, config := Load(filepath)

		assert.Error(t, err)
		assert.Empty(t, config)
	})
}
