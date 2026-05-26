func maxSlidingWindow(nums []int, k int) []int {
	q:=[]int{}
	ans := []int{}
	for i:=0;i<len(nums);i++{
		// pop low elements
		for len(q)>0 && nums[i]>nums[q[len(q)-1]]{
			q = q[:len(q)-1]
		}
		// append the current element
		q = append(q, i)

		if q[0]<=i-k{
			q = q[1:]
		}
		
		if i>=k-1{
			ans = append(ans, nums[q[0]])
		}
	}
	return ans
}