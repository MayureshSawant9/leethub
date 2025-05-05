func findClosestNumber(nums []int) int {
    closest_num := nums[0]
    lowest_dist := abs(nums[0])
    for _, num := range(nums){
        if dist := abs(num); dist < lowest_dist{
            closest_num = num
            lowest_dist = dist
        } else if abs(num) == lowest_dist && num > closest_num{
            closest_num = num
        }
    }
    return closest_num
}

func abs(x int) int{
    if x < 0{
        return -x
    }
    return x
}