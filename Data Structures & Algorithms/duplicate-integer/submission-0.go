func hasDuplicate(nums []int) bool {
    hits := map[int]struct{}{}
    for _, num := range nums {
        if _, ok := hits[num]; ok{
            return true
        }
        hits[num] = struct{}{}
    }
    return false
}
