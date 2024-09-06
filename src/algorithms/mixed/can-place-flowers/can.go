package mixed

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			leftSpotIsEmpty, rightSpotIsEmpty := i == 0 || flowerbed[i-1] == 0, i == len(flowerbed)-1 || flowerbed[i+1] == 0
			if leftSpotIsEmpty && rightSpotIsEmpty {
				// we can plant a flower
				flowerbed[i] = 1
				count++
			}
		}
	}
	return count >= n
}
