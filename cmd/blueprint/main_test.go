package main

import (
	"fmt"
	"testing"

	"github.com/serhijko/go-project-blueprint/cmd/blueprint/config"
	"github.com/stretchr/testify/assert"
)

// Example test to show usage of `make test`
func TestDummy(t *testing.T) {
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	assert.Equal(t, "Dummy Value", config.Config.ConfigVar)
}
