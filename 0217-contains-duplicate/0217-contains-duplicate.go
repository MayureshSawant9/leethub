func containsDuplicate(nums []int) bool {
    numSet := make(map[int]struct{})

    for _, num := range nums {
        if _, ok := numSet[num]; ok {
            return true
        }

        numSet[num] = struct{}{}
    }

    return false
}