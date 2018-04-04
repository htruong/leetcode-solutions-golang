func removeDuplicates(nums []int) int {
    var last int = 0 // last seen number
    var cur int = 1 // current insertion point
    
    if len(nums) == 0 {
        return 0
    }
    
    last = nums[0]
    
    for i := 1; i < len(nums); i++ {
        if nums[i] != last {
            last = nums[i]
            nums[cur] = last
            cur++
        }
    }
    
    return cur
}