import "regexp"
func isPalindrome(s string) bool {
    pattern := regexp.MustCompile("[^A-Za-z0-9]")
    phrase := strings.ToLower(pattern.ReplaceAllString(s, ""))

    l, r := 0, len(phrase) - 1
    fmt.Print(phrase)
    for l < r {
        if phrase[l] != phrase[r]{
            return false
        }
        l++
        r--
    }

    return true
}