package almostincreasingsequence

func almostIncreasingSequence(sequence []int) bool {
	invalidCount := 0
	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] >= sequence[i+1] {
			invalidCount++
			if invalidCount > 1 {
				return false
			}
			if i > 0 && sequence[i-1] >= sequence[i+1] && i+2 < len(sequence) && sequence[i] >= sequence[i+2] {
				return false
			}
		}
	}
	return true
}
