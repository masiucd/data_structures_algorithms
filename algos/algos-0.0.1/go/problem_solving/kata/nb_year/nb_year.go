package nb_year

func nbYear(peopleAtStart int, percent float64, aug int, targetPopulation int) (years int) {
	for peopleAtStart < targetPopulation {
		peopleAtStart = peopleAtStart + int(float64(peopleAtStart)*percent/100) + aug
		years++
	}
	return years
}
