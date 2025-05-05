func romanToInt(s string) int {
    value := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    i := 0
    result := 0
    for i < len(s){
        if i < len(s) - 1 && value[s[i]] < value[s[i+1]]{
            result += value[s[i+1]] - value[s[i]]
            i += 2
        } else {
            result += value[s[i]]
            i++
        }
    }
    return result
}