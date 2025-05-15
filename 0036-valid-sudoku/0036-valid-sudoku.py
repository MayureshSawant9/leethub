class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        row_dict = collections.defaultdict(set)
        column_dict = collections.defaultdict(set)
        box_dict = collections.defaultdict(set)

        for i in range(9):
            for j in range(9):
                current = board[i][j]
                if current==".": continue
                if current in row_dict[i] or current in column_dict[j] or current in box_dict[(i//3, j//3)]:
                    return False
                else:
                    row_dict[i].add(current)
                    column_dict[j].add(current)
                    box_dict[(i//3, j//3)].add(current)

        return True

        