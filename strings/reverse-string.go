func reverseString(s string) string {
    
    r := []rune(s)
    
    for i := 0; i < len(r) / 2; i ++ {
        tmp := r[i]
        r[i] = r[len(s) - 1 - i]
        r[len(s) - 1 - i] = tmp
    }
    
    return string(r)
}