func isAnagram(s string, t string) bool {
    count := make(map[rune]int)

    for _, char := range s {
        count[char]++
    }

    for _, char := range t {
        if _, ok := count[char]; !ok {
            return false
        } else if count[char] == 1{
            delete(count, char)
        } else {
            count[char]--
        }
    }

    return len(count) == 0
}