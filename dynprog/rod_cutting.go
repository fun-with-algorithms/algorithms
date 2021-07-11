package dynprog

func RodCutting(length int, cuts map[int]int) ([]int, int) {
	return rodCuttingMemoized(length, cuts, make(map[int]int), make(map[int][]int))
}

func rodCuttingMemoized(length int, cuts map[int]int, priceCalcs map[int]int, cutCals map[int][]int) ([]int, int) {
	if length <= 0 {
		return []int{}, 0
	}

	if price, ok := priceCalcs[length]; ok {
		return cutCals[length], price
	}

	maxPrice := 0
	bestCuts := []int{}

	for cutLen, cutPrice := range cuts {
		if cutLen > length {
			continue
		}

		cuts, subPrice := rodCuttingMemoized(length - cutLen, cuts, priceCalcs, cutCals)

		price := cutPrice + subPrice

		if price > maxPrice {
			maxPrice = price
			bestCuts = append(cuts, cutLen)
		}
	}

	priceCalcs[length] = maxPrice
	cutCals[length] = bestCuts

	return bestCuts, maxPrice
}