package dna_stand

func dnaStrand(dna string) string {
	var newDna string
	for _, str := range dna {
		switch {
		case str == 'T':
			newDna += "A"
		case str == 'A':
			newDna += "T"
		case str == 'C':
			newDna += "G"
		case str == 'G':
			newDna += "C"
		}
	}
	return newDna
}

//"ATTGC" --> "TAACG"
//"GTAT" --> "CATA"

// a -> t t --> a
// c -> g g -> c
