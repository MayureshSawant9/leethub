class Solution:
    def evalRPN(self, tokens: List[str]) -> int:

        def operate(operand1, operand2, operator):

            match operator:
                case '+':
                    return operand1 + operand2
                case '-':
                    return operand1 - operand2
                case '*':
                    return operand1 * operand2
                case '/':
                    return operand1 / operand2

        stack = []
        operators = ('+','-','*','/')

        for token in tokens:

            if token not in operators:
                stack.append(token)

            else:
                num2, num1 = int(stack.pop()), int(stack.pop())
                stack.append(operate(num1, num2, token))


        return int(stack[-1])


        