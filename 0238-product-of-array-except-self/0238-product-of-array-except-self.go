func productExceptSelf(nums []int) []int {
    result := make([]int, len(nums))

    prefixProduct := 1
    for i := 0; i < len(nums); i++ {
        result[i] = prefixProduct
        prefixProduct *= nums[i]
    }

    suffixProduct := 1
    for i := len(nums) - 1; i>=0; i-- {
        result[i] *= suffixProduct
        suffixProduct *= nums[i]
    }

    return result
}