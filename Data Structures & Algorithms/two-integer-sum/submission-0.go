func twoSum(nums []int, target int) []int {
    mymap := make(map[int]int, len(nums))
	var res []int
	for i, v := range nums {
		if _, ok := mymap[v]; !ok{
			mymap[target-v] = i;
			continue;
		}
		res = append(res, mymap[v],i)
		break
	}
	return res
}
