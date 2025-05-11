func productExceptSelf(nums []int) []int {
    result := make([]int, len(nums))

    prefix_product := 1
    for i := 0; i < len(nums); i++ {
        result[i] = prefix_product
        prefix_product *= nums[i]
    }

    suffix_product := 1
    for i := len(nums) - 1; i>=0; i-- {
        result[i] *= suffix_product
        suffix_product *= nums[i]
    }

    return result
}