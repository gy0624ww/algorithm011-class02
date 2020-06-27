package main

func plusOne(digits []int) []int {
	digitsLen := len(digits)
	for i:= digitsLen-1; i>=0; i-- {
		// 从第二次循环开始，如果循环了则说明有进位，进位加1
		// 个位数加1
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	// 如果是999这种，则结果肯定是1000，数组延长1位，首位为1 其余0
	newDigits := make([]int, digitsLen + 1)
	newDigits[0] = 1
	return newDigits
}
