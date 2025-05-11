func spiralOrder(matrix [][]int) []int {
    
    m, n := len(matrix), len(matrix[0])
    i, j := 0, 0
    RIGHT, DOWN, LEFT, UP := 0, 1, 2, 3
    topWall, rightWall, bottomWall, leftWall := 0, n, m, -1
    direction := RIGHT
    total := m*n
    result := []int{}

    for len(result) < total {
        if direction == RIGHT {
            for j < rightWall {
                result = append(result, matrix[i][j])
                j++
            }
            i++
            j--
            rightWall--
            direction = DOWN
        } else if direction == DOWN {
            for i < bottomWall {
                result = append(result, matrix[i][j])
                i++
            }
            i--
            j--
            bottomWall--
            direction = LEFT
        } else if direction == LEFT {
            for j > leftWall {
                result = append(result, matrix[i][j])
                j--
            }
            i--
            j++
            leftWall++
            direction = UP
        } else {
            for i > topWall {
                result = append(result, matrix[i][j])
                i--
            }
            i++
            j++
            topWall++
            direction = RIGHT
        }
    }
    return result
}