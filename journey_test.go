package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJourneyInit(t *testing.T) {
	// full journey
	s1 := station{"Old Street", 1}
	s2 := station{"London Bridge", 2}
	j := journey{&s1, &s2}
	assert.Equal(t, *j.start, s1, "expected journey start to be station 1")
	assert.Equal(t, *j.end, s2, "expected journey end to be station 2")
	// no end station
	j = journey{start: &s1}
	assert.Equal(t, *j.start, s1, "expected journey start to be station 1")
	assert.Nil(t, j.end, "expected journey end to be nil")
	// no start station
	j = journey{end: &s2}
	assert.Equal(t, *j.end, s2, "expected journey start to be station 1")
	assert.Nil(t, j.start, "expected journey end to be nil")
}

func TestCalculateFare(t *testing.T) {
	// DECLARE CONSTANTS FOR TEST
	s1 := station{"Bank", 1}
	s2 := station{"Moorgate", 1}
	// VALID JOURNEY
	j := journey{&s1, &s2}
	n := calculateFare(j)
	assert.Equal(t, n, 1, "expected result to equal 1")
	// PENALTY FARE - NO END STATION
	j = journey{start: &s1}
	n = calculateFare(j)
	assert.Equal(t, n, 6, "expected result to equal 6")
	// PENALTY FARE - NO START STATION
	j = journey{end: &s2}
	n = calculateFare(j)
	assert.Equal(t, n, 6, "expected result to equal 6")

}
