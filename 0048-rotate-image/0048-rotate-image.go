func rotate(matrix [][]int)  {
    n := len(matrix)
    half := n/2

    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }

    for i := 0; i < n; i++ {
        for j := half; j < n; j++ {
            matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
        }
    }
}