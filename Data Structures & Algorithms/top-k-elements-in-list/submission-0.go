import "slices"
func topKFrequent(nums []int, k int) []int {
	numFreq := map[int][]int{}

	for _, num := range nums {
		values := numFreq[num]
		if len(values) != 2 {
			numFreq[num] = []int{num, 1}
			continue
		}
		numFreq[num] = []int{num, values[1]+ 1}
	}

	results := make([][]int, 0, len(numFreq))
	for _, vals := range numFreq {
		results = append(results, vals)
	}

	slices.SortFunc(results, func(a, b []int) int {
		if a[1] < b[1] {
			return 1
		} else if a[1] > b[1] {
			return -1
		} else {
			return 0
		}
	})

	items := make([]int, 0, k)
	for i := 0; i < k; i++ {
		items = append(items, results[i][0])
	}
	return items
}
