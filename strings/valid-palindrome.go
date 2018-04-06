func cleanup(s string) []byte {
    tmp := make([]byte, len(s))
    ctr := 0
    for _, v := range([]byte(s)) {
        if ('0' <= v && v <= '9') || ('a' <= v && v <= 'z') || ('A' <= v && v <= 'Z') {
            if 'A' <= v && v <= 'Z' {
                v += 32
            }
            tmp[ctr] = v
            ctr ++
        }
    }
    return tmp[0:ctr]
}

func isPalindrome(s string) bool {
    x := cleanup(s)
    for i := 0; i < len(x) / 2; i ++ {
        if x[i] != x[len(x) - 1 - i] {
            return false
        }
    }
    return true
}