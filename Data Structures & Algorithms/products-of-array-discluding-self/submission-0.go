func productExceptSelf(nums []int) []int {
    product := 1
    zeros := map[int]struct{}{}
    for i, num := range nums {
        if num == 0 {
            zeros[i] = struct{}{}
            continue
        }

        product *= num
    }

    switch len(zeros) {
        case 0:
            products := make([]int, 0, len(nums))
            for i, num :=  range nums {
                if _, ok := zeros[i]; ok {
                    products = append(products, product)
                    continue
                }
                products = append(products, product/num)
            }
            return products
        case 1:
            products := make([]int, len(nums))
            for k, _ :=  range zeros {
                products[k] = product
            }
            return products
        default:
            return make([]int, len(nums))
    }
}
