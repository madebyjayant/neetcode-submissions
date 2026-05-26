func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	l,r := 0,0
	q := []int{}
	for r<len(nums) {
		for len(q)>0 && nums[q[len(q)-1]] < nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)

		// window limit
		if l > q[0]{
			q = q[1:]
		}

		if r>=k-1{
			ans = append(ans, nums[q[0]])
			l++
		}
		r++
	}
	
	return ans
}