class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:

        m, n = len(matrix), len(matrix[0])
        i, j = 0, 0
        RIGHT, LEFT, UP, DOWN = 0, 1, 2, 3
        direction = RIGHT
        top_wall, bottom_wall, right_wall, left_wall = 0, m, n, -1
        total = m * n

        result = []

        while len(result) != total:
            if direction == RIGHT:
                while j < right_wall:
                    result.append(matrix[i][j])
                    j += 1
                j -= 1
                i += 1
                right_wall -= 1
                direction = DOWN

            elif direction == DOWN:
                while i < bottom_wall:
                    result.append(matrix[i][j])
                    i += 1
                i -= 1
                j -= 1
                bottom_wall -= 1
                direction = LEFT

            elif direction == LEFT:
                while j > left_wall:
                    result.append(matrix[i][j])
                    j -= 1
                i -= 1
                j += 1
                left_wall += 1
                direction = UP

            elif direction == UP:
                while i > top_wall:
                    result.append(matrix[i][j])
                    i -= 1
                i += 1
                j += 1
                top_wall += 1
                direction = RIGHT

        return result