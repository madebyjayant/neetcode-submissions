func findDuplicate(nums []int) int {
	seen := make(map[int]int, 0)
	for _, v := range nums {
		if _, ok := seen[v];ok{
			return v;
		}
		seen[v]++
	}
	return 0
}
