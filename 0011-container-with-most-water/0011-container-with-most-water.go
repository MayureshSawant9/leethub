func maxArea(height []int) int {
    
    n := len(height)
    maxArea := 0
    l, r := 0, n-1

    for l < r {
        maxArea = max((r-l) *  min(height[l], height[r]), maxArea)
        fmt.Printf("Area : %v", maxArea)
        if height[l] < height[r] {
            l ++
        } else {
            r --
        }
    }

    return maxArea
}

func max(num1, num2 int) int {
    if num1 < num2 {
        return num2
    } else {
        return num1
    }
}

func min(num1, num2 int) int {
    if num1 < num2 {
        return num1
    } else {
        return num2
    }
}