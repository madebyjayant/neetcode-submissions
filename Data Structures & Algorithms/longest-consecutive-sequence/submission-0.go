func longestConsecutive(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _,v:=range nums {
		m[v]=true;
	}

	
	startPoints := make([]int, 0)
	for _,v := range nums {
		if !m[v-1]{
			startPoints = append(startPoints, v);
		}
	}

	longest := 0
	for _, pt := range startPoints {
		length := 1
		for{
			if m[pt+1]{
				length++
				pt++
				continue
			}
			break
		}

		if length > longest {longest=length}
	}

	return longest
}
