package main

import "strconv"

type journeyLog struct {
	journeys       []journey
	currentJourney *journey
}

func (jl *journeyLog) startJourney(s *station) {
	jl.currentJourney = &journey{start: s}
}

func (jl *journeyLog) endJourney(s *station) journey {
	if jl.currentJourney == nil {
		jl.currentJourney = &journey{}
	}
	j := jl.currentJourney
	j.end = s
	jl.journeys = append(jl.journeys, *j)
	jl.currentJourney = nil
	return *j
}

func formatStationNames(s *station) string {
	if s == nil {
		return "N/A"
	}
	return s.name
}

func (jl *journeyLog) history() string {
	if len(jl.journeys) == 0 {
		return "No Journeys"
	}
	h := ""
	for i, j := range jl.journeys {
		h += strconv.Itoa(i+1) + ". Start: " + formatStationNames(j.start) + " End: " + formatStationNames(j.end) + "\n"
	}
	return h
}
