class Solution:
    def isValid(self, s: str) -> bool:

        stack = deque()

        closing = {')':'(', '}':'{',']':'['}

        for bracket in s:
            
            if bracket not in closing:
                stack.append(bracket)

            elif stack and stack[-1] == closing[bracket]:
                stack.pop()

            else:
                return False

        return len(stack) == 0
        