func intersect(nums1 []int, nums2 []int) []int {
    found2 := make([]bool, len(nums2))
    found1 := make([]bool, len(nums1))
    count := 0
    for j, val := range(nums1) {
        for k, val2 := range(nums2) {
            if found1[j] == false && found2[k] == false && val==val2 {
                found2[k]= true
                found1[j]= true
                count ++
            }
        }
    }
    ret := make([]int, count)
    indx := 0
    for k, f := range found2 {
        if f { 
            ret[indx] = nums2[k]
            indx++
        }
    }
    return ret
}