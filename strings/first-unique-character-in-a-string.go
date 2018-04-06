func firstUniqChar(s string) int {
    seen := make(map[byte]bool)
    //marked := make(map[byte]bool)
    fseen := make(map[byte]int)
    dup := make([]bool, len(s))
    
    
    
    for k, v := range([]byte(s)) {
        //v := byte(v1)
        
        if seen[v] {
            dup[k] = true
            //if ! marked[v] {
                dup[fseen[v]] = true
                //marked[v] = true
            //}
        } else {
            seen[v] = true
            fseen[v] = k
        }
    }
    
    for k, v := range(dup) {
        if v == false {
            return k
        }
    }
    
    return -1
}