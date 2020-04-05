package main

type journey struct {
	start *station
	end   *station
}

func isPenalty(j journey) bool {
	return j.start == nil || j.end == nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateFare(j journey) int {
	// CHECK FOR PENALTY
	if isPenalty(j) {
		return penaltyFare
	}
	return abs(j.start.zone-j.end.zone) + minFare
}
