package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainSucess(t *testing.T) {
	c, err := kafkaConfig()
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, c.Address, "IDFE")
}
func TestMainFaliure(t *testing.T) {
	c, err := kafkaConfig()
	assert.Nil(t, err)
	assert.NotNil(t, c)
}

