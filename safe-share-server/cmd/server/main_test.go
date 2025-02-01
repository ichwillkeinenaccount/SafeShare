package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadConfig_Success(t *testing.T) {
	assert.NotPanics(t, func() {
		readConfig("config.yaml")
	})
}

func TestReadConfig_FileNotFound(t *testing.T) {
	assert.Panics(t, func() {
		readConfig("nonexistent.yaml")
	})
}
