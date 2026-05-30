type stack struct {
	s [][]int
}

func (s *stack) push(a,b int) {
	s.s = append(s.s, []int{a,b})
}

func (s *stack) pop() []int{
	if len(s.s)!=0 {
		c:= s.s[len(s.s)-1]
		s.s = s.s[:len(s.s)-1]
		return c
	}
	return []int{}
}

func (s *stack) top() []int{
	if len(s.s)!=0 {
		return s.s[len(s.s)-1]
	}
	return []int{}
}

func (s *stack) show() {
	fmt.Println("[Printing Stack]:")
	for _, val := range s.s{
		fmt.Println(val)
	}
}

func dailyTemperatures(temperatures []int) []int {
	s := &stack{
		s : [][]int{},
	}

	result := make([]int, len(temperatures))

	for idx, v := range temperatures {
		c:= s.top()
		for len(s.s)!=0 && c[0]<v{
			result[c[1]] = idx - c[1]
			s.pop()
			c = s.top()
			continue
		}
		s.push(v, idx)
	}
	return result
}
