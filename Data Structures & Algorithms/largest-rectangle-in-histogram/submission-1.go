func largestRectangleArea(heights []int) int {
	stack := make([][]int, 0)
	m := 0
	for idx, height := range heights {
		if len(stack)==0 {
			stack = append(stack, []int{idx, height})
			m = height
			continue
		}

		t := idx
		for len(stack)!=0 && height < stack[len(stack)-1][1]{
			t = stack[len(stack)-1][0]
			// pop
			top := stack[len(stack)-1]
			m = max(m, top[1]*(idx-top[0]))
			stack = stack[:len(stack)-1]
		}
		if len(stack)>0 && height == stack[len(stack)-1][1]{
			continue
		}
		// push the current one with the prev idx
		stack = append(stack, []int{t, height})
		m = max(m, height*(idx-t+1))
	}

	for len(stack)>0{
		top:=stack[len(stack)-1]
		m = max(m,top[1]*(len(heights)-top[0]))
		stack = stack[:len(stack)-1]
	}

	return m
}
