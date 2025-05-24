import "sort"
func threeSum(nums []int) [][]int {
    
    sort.Ints(nums)
    n := len(nums)
    result := [][]int{}

    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            break
        } else if i > 0 && nums[i] == nums[i - 1] {
            continue
        }

        low, high := i + 1, n - 1
        
        for low < high {
            sum := nums[i] + nums[low] + nums[high]

            if sum == 0 {
                result = append(result, []int{nums[i], nums[low], nums[high]})
                low++
                high--

                for low < high && nums[low] == nums[low - 1] {
                    low++
                }

                for low < high && nums[high] == nums[high + 1] {
                    high--
                }
            } else if sum < 0 {
                low++
            } else {
                high--
            }
        }
    }
    return result
}