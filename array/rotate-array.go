// This is really bad, I will need to fix it

func shr(nums []int) {
    shout := nums[len(nums)-1]
    for i := len(nums) - 1; i > 0; i-- {
        nums[i] = nums[i-1]
    }
    nums[0] = shout
}

func rotate(nums []int, k int)  {
    
    k = k % len(nums)
    if k == 0 {
        return
    }
    
    for i := 0; i< k; i++ {
        shr(nums)
    }
}