package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ShestakovDA/warlog/config"
)

const (
	testV1_0Path = "test_data/test_1_0.yml"
)

func TestFromFile(t *testing.T) {
	assert.True(t, false)
}

func TestV1_0FromFile(t *testing.T) {
	expected := config.V1_0{}
	config.FromFile(testV1_0Path, expected)

	actual := config.V1_0{}
	actual.FromFile(testV1_0Path)

	assert.Equal(t, expected, actual)
}
