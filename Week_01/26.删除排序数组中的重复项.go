package main

func RemoveDuplicates(nums []int) int {
    i:= 0
    j:= i + 1
    numsLen := len(nums)
    for j< numsLen {
        if nums[i] != nums[j] {
            nums[i+1] = nums[j]
            i++
        }
        j++
    }
    // because begin from 0
    return i+1
}
