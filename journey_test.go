package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJourneyInit(t *testing.T) {
	s1 := station{"Old Street", 1}
	s2 := station{"London Bridge", 2}
	j := journey{s1, s2}
	assert.Equal(t, j.start, s1, "expected journey start to be station 1")
	assert.Equal(t, j.end, s2, "expected journey end to be station 2")
}
