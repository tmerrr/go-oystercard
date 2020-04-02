package main

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
