func search(nums []int, target int) int {
	l,r := 0, len(nums)-1
	if nums[l]==target {
		return l
	}
	if nums[r]==target {
		return r
	}

	for l <= r {
		m := (l+r)/2
		if nums[m] == target {
			return m
		}

		if nums[m] >= nums[l] {
			if nums[m]<target || nums[l]>target {
				l = m+1
			}else{
				r = m-1
			}
		}else {
			if target < nums[m] || target > nums[r] {
                r = m - 1
            } else {
                l = m + 1
            }
		}
	}
	return -1
}
