
func longestConsecutive(nums []int) int {
    set := map[int]struct{}{}
    for _, num := range nums {
        if _, ok := set[num]; !ok{
            set[num] = struct{}{}
        }
    }

    starts := map[int]int{}
    for n, _ :=  range set {
        if _, ok :=  set[n-1]; !ok {
            starts[n] = 1
        }
    }

    for start, _ :=  range starts {
        for i := start+1; i <= int(math.Pow(10, 9)); i++ {
            if _, exists := set[i]; !exists {
               break
            }
            starts[start]++
        }
    }
    
    max := 0
    for _, length := range starts {
        if length > max {
            max = length
        }
    }

    return max
}
