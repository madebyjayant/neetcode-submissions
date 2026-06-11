func findDuplicate(nums []int) int {
	marks := make([]int, len(nums))
	for _,v := range nums {
		if marks[v-1]!=0{
			return v
		}
		marks[v-1]++
	}
	return -1
}
