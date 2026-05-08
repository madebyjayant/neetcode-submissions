func twoSum(numbers []int, target int) []int {
	i:=0
	j:=len(numbers)-1

	for i<len(numbers) && j>=0 && i<j{
		curr := numbers[i]+numbers[j]
		if curr==target{
			return []int{i+1,j+1}
		}
		if curr<target {
			i++
			continue;
		}
		j--
	}
	return []int{-1,-1}
}
