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
	// DECLARE CONSTANTS FOR TESTS
	s1 := station{"St Pauls", 1}
	s2 := station{"Finsbury Park", 2}
	s3 := station{"Wimbledon", 3}
	s4 := station{"Wembley Stadium", 4}
	s5 := station{"Twickenham", 5}
	s6 := station{"Kingston", 6}

	// VALID JOURNEY - SAME ZONE
	j := journey{&s1, &s1}
	n := calculateFare(j)
	assert.Equal(t, 1, n, "expected result to equal 1")

	// VALID JOURNEY - 1 ZONE DIFF (BOTH WAYS)
	j = journey{&s1, &s2}
	n = calculateFare(j)
	assert.Equal(t, 2, n, "expected result to equal 2")

	j = journey{&s2, &s1}
	n = calculateFare(j)
	assert.Equal(t, 2, n, "expected result to equal 2")

	// VALID JOURNEY - 2 ZONE DIFF (BOTH WAYS)
	j = journey{&s1, &s3}
	n = calculateFare(j)
	assert.Equal(t, 3, n, "expected result to equal 3")

	j = journey{&s3, &s1}
	n = calculateFare(j)
	assert.Equal(t, 3, n, "expected result to equal 3")

	// VALID JOURNEY - 3 ZONE DIFF (BOTH WAYS)
	j = journey{&s1, &s4}
	n = calculateFare(j)
	assert.Equal(t, 4, n, "expected result to equal 4")

	j = journey{&s4, &s1}
	n = calculateFare(j)
	assert.Equal(t, 4, n, "expected result to equal 4")

	// VALID JOURNEY - 4 ZONE DIFF (BOTH WAYS)
	j = journey{&s1, &s5}
	n = calculateFare(j)
	assert.Equal(t, 5, n, "expected result to equal 5")

	j = journey{&s5, &s1}
	n = calculateFare(j)
	assert.Equal(t, 5, n, "expected result to equal 5")

	// VALID JOURNEY - 5 ZONE DIFF (BOTH WAYS)
	j = journey{&s1, &s6}
	n = calculateFare(j)
	assert.Equal(t, 6, n, "expected result to equal 6")

	j = journey{&s6, &s1}
	n = calculateFare(j)
	assert.Equal(t, 6, n, "expected result to equal 6")

	// PENALTY FARE - NO END STATION
	j = journey{start: &s1}
	n = calculateFare(j)
	assert.Equal(t, 6, n, "expected result to equal 6")

	// PENALTY FARE - NO START STATION
	j = journey{end: &s2}
	n = calculateFare(j)
	assert.Equal(t, 6, n, "expected result to equal 6")

}
