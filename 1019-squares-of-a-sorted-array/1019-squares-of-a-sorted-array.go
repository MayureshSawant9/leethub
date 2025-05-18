func sortedSquares(nums []int) []int {
    l, r, n := 0, len(nums) - 1, len(nums)
    result := []int{}

    for l <= r {
        if abs(nums[l]) < abs(nums[r]) {
            result = append(result, nums[r] * nums[r])
            r--
        } else {
            result = append(result, nums[l] * nums[l])
            l++
        }
    }
    
    for i := 0; i < n / 2; i++ {
        result[i], result[n - i - 1] = result[n - i - 1], result[i]
    }

    return result

}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}