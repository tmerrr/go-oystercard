package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJourneyLogInit(t *testing.T) {
	jl := journeyLog{}
	assert.IsType(t, jl.journeys, []journey{}, "expected journeys to be slice of type journey")
	assert.Len(t, jl.journeys, 0, "expected journeys to have length of 0")
	assert.Nil(t, jl.currentJourney, "expected currentJourney to be nil")
}

func TestStartJourney(t *testing.T) {
	jl := journeyLog{}
	s := station{"London Bridge", 1}
	jl.startJourney(&s)
	assert.Len(t, jl.journeys, 0, "expected journeys to have length of 0")
	assert.Equal(t, jl.currentJourney.start, &s, "expected currentJourney start to be correct station")
	assert.Nil(t, jl.currentJourney.end, "expected currentJourney end to be nil")
}

func TestEndJourney(t *testing.T) {
	jl := journeyLog{}
	s1 := station{"London Bridge", 1}
	s2 := station{"Old Street", 1}
	jl.currentJourney = &journey{start: &s1}
	j := jl.endJourney(&s2)
	assert.Len(t, jl.journeys, 1, "expected journeys to have length of 1")
	assert.Equal(t, jl.journeys[0], journey{&s1, &s2}, "expected recorded journey to have correct values")
	assert.Nil(t, jl.currentJourney, "expected currentJourney to be nil")
	assert.Equal(t, j, jl.journeys[0], "expected return value to be completed journey")
}
