package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStationInit(t *testing.T) {
	s := station{
		name: "Old Street",
		zone: 1,
	}
	assert.Equal(t, s.name, "Old Street", "expected station name to be Old Street")
	assert.Equal(t, s.zone, 1, "expected station zone to be 1")
}
