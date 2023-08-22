package list_cardiopackage

func totalYearsLived(input []Inventor) int {
	var result int
	for _, inventor := range input {
		yearLived := inventor.passed - inventor.year
		result += yearLived
	}
	return result

}
