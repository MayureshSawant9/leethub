import (
    "fmt"
    "strconv"
)
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	result := []string{}
	start := nums[0]

	for i := 1; i <= len(nums); i++ { // loop until the end
		if i == len(nums) || nums[i] != nums[i-1]+1 {
			// Check if the current range is just a single number
			if start == nums[i-1] {
				result = append(result, strconv.Itoa(start)) // Use strconv for int to string conversion
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, nums[i-1])) // Use fmt.Sprintf for ranges
			}
			// If we haven't reached the end, update start for the next range
			if i < len(nums) {
				start = nums[i]
			}
		}
	}
	return result
}