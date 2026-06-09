func findMin(nums []int) int {
	l,r := 0, len(nums)-1
	if nums[l]<nums[r] {return nums[l]}
	if len(nums)==1{
		return nums[0]
	}

	if len(nums)==2{
		return min(nums[0],nums[1])
	}
	for l<r {
		m := (l+r)/2

		if nums[m]<nums[r]{
			// starting point is in between l, m
			// but right now it is m
			r = m
		}else {
			l = m+1
		}
	}

	return nums[l]
}
