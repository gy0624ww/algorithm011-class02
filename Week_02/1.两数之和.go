package main

func twoSum(nums []int, target int) []int {
	numsLen := len(nums)
	nmap := make(map[int]int)
	for i:=0;i<numsLen;i++ {
		OtherNum := target-nums[i]
		if _,ok := nmap[OtherNum]; ok {
			return []int{nmap[OtherNum], i}
		}
		nmap[nums[i]] = i
	}
	return nil
}


