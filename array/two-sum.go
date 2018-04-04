func twoSum(nums []int, target int) []int {
    needs := make(map[int]bool)
    needs_where := make(map[int]int)
    for k, v := range (nums) {
        if needs[v] == true {
            return []int{needs_where[v], k}
        } else {
            needs[target - v] = true
            needs_where[target - v] = k
        }
    }
    return []int{-1,-1}
}