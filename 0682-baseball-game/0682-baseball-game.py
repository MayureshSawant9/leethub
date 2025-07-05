class Solution:
    def calPoints(self, operations: list[str]) -> int:
        score_stack = deque()

        for operation in operations:
            match operation:
                case "+":
                    score_stack.append(score_stack[-1] + score_stack[-2])
                case "D":
                    score_stack.append(2 * score_stack[-1])
                case "C":
                    score_stack.pop()
                case _:
                    # Assume it's a numeric string
                    score_stack.append(int(operation))

        return sum(score_stack)