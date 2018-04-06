import "math"

func reverse(x int) int {
    var MaxInt int32 = 2147483647
    
    if x > int(MaxInt) || x < - int(MaxInt) { return 0 }
    
    tmp := 0
    run := x
    if (x < 0) { run = -run }
    for run > 0 {
        tmp *= 10
        digit := run % 10
        tmp += digit
        run /= 10
    }
    
    if tmp > int(MaxInt) { return 0 }
    
    if (x < 0) {
        return -tmp
    }
    return tmp
}