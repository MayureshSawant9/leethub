func countSort(word string) string {
        counts := make([]int, 26)
        for _, char := range word {
            counts[int(char) - int('a')]++
        }
        var builder strings.Builder
        for index, count := range counts {
            for i := 0; i < count; i++ {
                builder.WriteByte(byte(index + int('a')))
            }
        }
        return builder.String()
}

func groupAnagrams(strs []string) [][]string {

    hashMap := make(map[string][]string)

    for _, word := range strs {
        sortedWord := countSort(word)
        if _, ok := hashMap[sortedWord]; !ok {
            hashMap[sortedWord] = make([]string, 0)
        }
        hashMap[sortedWord] = append(hashMap[sortedWord], word)
    }

    result := [][]string{}
    for _, value := range hashMap {
        result = append(result, value)
    }
    
    return result
}