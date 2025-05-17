func longestConsecutive(nums []int) int {
    set := make(map[int]struct{})
    maximum := 0

    for _, num := range nums {
        set[num] = struct{}{}
    }

    for key, _ := range set {
        if _, contains := set[key - 1]; !contains {
            current := key + 1 
            longest := 1
            for {
                if _, contains := set[current]; contains {
                    current++
                    longest++
                } else {
                    break
                }
            }
            maximum = max(maximum, longest)
        }
    }

    return maximum
}