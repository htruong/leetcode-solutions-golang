func findNonConformZeros(nums []int) ([]bool, int) {
    out := make([]bool, len(nums))
    
    expectingForZero := true
    count := 0
    for i:=len(nums) - 1; i >= 0; i -- {
        if expectingForZero {
            if nums[i] != 0 {
                expectingForZero = false
            }
        } else {
            if nums[i] == 0 {
                out[i] = true
                count++
            }
        }
    }
    return out, count
}

func moveZeroes(nums []int)  {
    nczs, nczsc := findNonConformZeros(nums)
    shifts := 0
    for k, _ := range(nczs) {
        if (k < len(nums) - nczsc) {
            if nczs[k] == true {
                shifts ++
                idx := k + shifts
                val := nczs[idx]
                for val != true && idx < len(nums) {
                    nums[idx] = nums[idx+shifts]
                    idx++
                    val = nczs[idx]
                }
            }
        } else {
            nums[k] = 0
        }
    }
}