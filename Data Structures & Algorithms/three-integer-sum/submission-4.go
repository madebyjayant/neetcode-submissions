import "slices"
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	ans := make([][]int, 0)

	for i:=0;i<len(nums);i++{
		newTarget := 0-nums[i];
		curr := nums[i]
		left, right := i+1,len(nums)-1
		for left<right{
			if nums[left]+nums[right]==newTarget{
				ans = addUnique(ans, []int{curr,nums[left],nums[right]})
			}
			if nums[left]+nums[right]>newTarget{
				right--
				continue
			}
			left++
		}
	}
	return ans
}

func addUnique(all [][]int, triplet []int) [][]int{
	for _, sl := range all {
		if slices.Equal(sl, triplet){
			return all
		}
	}
	return append(all, triplet)
}
