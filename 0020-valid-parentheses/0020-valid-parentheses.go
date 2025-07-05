func isValid(s string) bool {
    
    stack := []rune{}
    
    closing := map[rune]rune{
        ')':'(',
        '}':'{',
        ']':'[',
    }

    for _, bracket := range s {

        if _, ok := closing[bracket]; !ok {
            stack = append(stack, bracket)
        } else {
            if len(stack) == 0 || stack[len(stack) - 1] != closing[bracket]{
                return false
            }
            stack = stack[:len(stack) - 1]
        }
    }
    return len(stack) == 0
}