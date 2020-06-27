package main

// 比较巧妙的解法
func moveZeros(nums []int) {
	if nums == nil {
		return
	}
	j := 0
	numsLen := len(nums)
	for i:= 0; i<numsLen; i++ {
		if nums[i] != 0 {
			if i > j {
				nums[j] = nums[i]
				nums[i] = 0
			}
			j++
		}
	}
}

// 类似快速排序的解法，以0作为中间点，把不是0的数交换到左边
func moveZeros2(nums []int) {
	if nums == nil {
		return
	}
	j := 0
	numsLen := len(nums)
	for i:=0; i<numsLen; i ++ {
		if nums[i] != 0 {
			if i >j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
}

// 先把非零的元素统计起来依此覆盖数组前面元素，再之后直到数组结尾补零
func moveZeors3(nums []int) {
	if nums == nil {
		return
	}
	j := 0
	numsLen := len(nums)
	for i:=0; i< numsLen; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for i:=j; i< numsLen; i++ {
		nums[i] = 0
	}
}
