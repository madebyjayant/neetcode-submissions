func findMin(nums []int) int {
	temp := nums[0]
	startingPoint := 0
	for idx, v := range nums {
		if v < temp{
			startingPoint = idx
			break
		}else {
			temp = v
		}
	}
	return nums[startingPoint]
}
