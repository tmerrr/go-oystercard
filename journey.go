package main

type journey struct {
	start *station
	end   *station
}

func isPenalty(j journey) bool {
	return j.start == nil || j.end == nil
}

func calculateFare(j journey) int {
	// CHECK FOR PENALTY
	if isPenalty(j) {
		return penaltyFare
	}

	return minFare
}
