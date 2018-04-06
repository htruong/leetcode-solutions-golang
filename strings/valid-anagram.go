func isAnagram(s string, t string) bool {
    cnt_s := make(map[rune]int)
    cnt_t := make(map[rune]int)
    
    for _, v := range(s) {
        val, ok := cnt_s[v]
        if (ok) {
            val ++
            cnt_s[v] = val
        } else {
            cnt_s[v] = 1
        }
    }
    
    for _, v := range(t) {
        val, ok := cnt_t[v]
        if (ok) {
            val ++
            cnt_t[v] = val
        } else {
            cnt_t[v] = 1
        }
    }
    
    for k, v := range(cnt_s) {
        if cnt_t[k] != v { return false }
    }
    
    for k, v := range(cnt_t) {
        if cnt_s[k] != v { return false }
    }
    
    return true
}