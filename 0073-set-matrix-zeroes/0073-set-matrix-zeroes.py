class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        row_zeros = set()
        column_zeros = set()

        m = len(matrix)
        n = len(matrix[0])

        for i in range(m):
            for j in range(n):
                if not matrix[i][j]: row_zeros.add(i); column_zeros.add(j)

        for i in range(m):
            for j in range(n):
                if i in row_zeros or j in column_zeros:
                    matrix[i][j] = 0