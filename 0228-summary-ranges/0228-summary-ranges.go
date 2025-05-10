import "fmt"
func summaryRanges(nums []int) []string {

    if len(nums) == 0 {
        return []string{}
    }

    result := []string{}
    start := nums[0]

    for i := 1; i <= len(nums); i++ {
        if i == len(nums) || nums[i] != nums[i-1] + 1{
            
            if start == nums[i-1] {
                result = append(result, fmt.Sprint(start))
            } else {
                result = append(result, fmt.Sprintf("%v->%v", start, nums[i-1]))
            }

            if i < len(nums) {
                start = nums[i]
            }
        }
    } 
    return result

}