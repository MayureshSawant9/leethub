func canConstruct(ransomNote string, magazine string) bool {
    
    count := make(map[rune]int)
    
    for _, letter := range magazine {
        count[letter]++
    }

    for _, letter := range ransomNote {
        if _, ok := count[letter]; !ok {
            return false
        } else if count[letter] == 1{
            delete(count, letter)
        } else {
            count[letter]--
        }
    }
    return true
}