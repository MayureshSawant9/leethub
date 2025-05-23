func reverseString(s []byte)  {
    // n := len(s)

    // for i:=0; i<n/2; i++ {
    //     s[i], s[n-i-1] = s[n-i-1], s[i]
    // }

    // two pointer
    l, r := 0, len(s) - 1

    for l < r {
        s[l], s[r] = s[r], s[l]
        l++
        r--
    }

}