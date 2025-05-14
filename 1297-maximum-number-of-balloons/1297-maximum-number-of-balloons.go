func maxNumberOfBalloons(text string) int {
    count := make(map[rune]int)
    for _, char := range text {
        count[char]++
    }

    return min(
        count['b'],
        count['a'],
        count['l']/2,
        count['o']/2,
        count['n'],
    )
}