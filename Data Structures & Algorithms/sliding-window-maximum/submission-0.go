func maxSlidingWindow(nums []int, k int) []int {
	// result is stored in ans
	var ans []int

	if len(nums)<k {
		return []int{}
	}

	for i:=0;i<len(nums)-k+1;i++{
		// for every ith till i+kth position
		// brute force
		max := findMax(nums[i:i+k])
		ans = append(ans,max)
	}

	return ans
}

func findMax(nums []int) int {
	max:=-1
	for _, v := range nums {
		if v>max{
			max = v
		}
	}
	return max
}
