package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_V1(t *testing.T) {
	// Тесты внутри модуля могут сравнивать указатели внутренних переменных
	cfg := V1()
	assert.Equal(t, cfg, v1)
}

func Test_ReloadV1(t *testing.T) {
	assert.NoError(t, ReloadV1("test_v1.yml"))
}

func Test_cfg_v1_Validate(t *testing.T) {
	assert.NoError(t, V1().Validate())
}
