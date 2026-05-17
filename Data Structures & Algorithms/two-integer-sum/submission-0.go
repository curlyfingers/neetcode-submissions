func twoSum(nums []int, target int) []int {
    solutions := map[int]int{}

	for i, num := range nums {
		if j, ok := solutions[num]; ok {
			return []int{min(i, j), max(i, j)}
		}
		solutions[target-num] = i
	}

	return []int{}
}
