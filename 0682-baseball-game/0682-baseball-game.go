func calPoints(operations []string) int {
    
    stack := []int{}

    for _, operation := range operations {
        switch operation {
            case "+":
                stack = append(stack, stack[len(stack) - 1] + stack[len(stack) - 2])

            case "D":
                stack = append(stack, stack[len(stack) - 1] * 2)

            case "C":
                stack = stack[:len(stack) - 1]

            default:
                value, _ := strconv.Atoi(operation)
                stack = append(stack, value)
        }
    }

    total := 0

    for _, value := range stack {
        total += value
    }

    return total
}