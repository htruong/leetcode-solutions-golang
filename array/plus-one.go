func plusOne(digits []int) []int {
    carry := 1
    for i := len(digits) - 1; i >= 0; i -- {
        digits[i] = digits[i] + carry
        if digits[i] >= 10 {
            digits[i] = 0
            carry = 1
        } else {
            carry = 0
        }
    }
    if carry == 1 {
        out := make([]int, len(digits) + 1)
        copy(out[1:], digits)
        out[0] = 1
        return out
    } else {
        return digits
    }
}