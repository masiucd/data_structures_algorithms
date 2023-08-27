package list_cardiopackage

//1. Filter the list of inventors for those who were born in the 1500's

func onlyBorn1500(input []Inventor) []Inventor {
	var result []Inventor
	for _, v := range input {
		if v.year >= 1500 && v.year <= 1599 {
			result = append(result, v)
		}
	}

	return result
}
