func mergeAlternately(word1 string, word2 string) string {
    result := []byte{}
    ptr1, ptr2 := 0, 0
    switchTurn := true

    for ptr1 < len(word1) && ptr2 < len(word2){
        if switchTurn{
            result = append(result, word1[ptr1])
            ptr1++
        } else {
            result = append(result, word2[ptr2])
            ptr2++
        }
        switchTurn = !switchTurn
    }

    for ptr1 < len(word1){
        result = append(result, word1[ptr1])
        ptr1++
    }

    for ptr2 < len(word2){
        result = append(result, word2[ptr2])
        ptr2++
    }
    return string(result)
}