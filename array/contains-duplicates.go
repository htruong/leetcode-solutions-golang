import "sort"

// Could be done with a map too

func containsDuplicate(nums []int) bool {
    sort.Sort(sort.IntSlice(nums))
    last := 0
    for k, v := range(nums) {
        if k == 0 {
            last = v
        } else {
            if v == last {
                return true
            }
            last = v
        }
    }
    return false
}