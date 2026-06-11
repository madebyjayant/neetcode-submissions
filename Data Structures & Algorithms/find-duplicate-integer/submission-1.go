import "slices"
func findDuplicate(nums []int) int {
	slices.Sort(nums)
	if nums==nil{
		return -1
	}
	prev := -1
	for _,v := range nums{
		if prev == v{
			return v
		}
		prev = v
	}
	return -1
}
