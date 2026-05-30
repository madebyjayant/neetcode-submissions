type Stack struct {
	st []int
}

func (s *Stack) push(a int) {
	s.st = append(s.st, a)
}

func (s *Stack) pop() {
	if s.st==nil {return}
	s.st = s.st[:len(s.st)-1]
}

func (s *Stack) top() int{
	if s.st==nil{return -1}
	return s.st[len(s.st)-1]
}

func (s *Stack) show() {
	fmt.Println("[Printing Stack]")
	for _, v := range s.st {
		fmt.Println(v)
	}
}
func evalRPN(tokens []string) int {
	stack := Stack{st: make([]int, len(tokens))}
	for _, token := range tokens {
		switch token {
			case "+":
				stack.push(stack.add())
			case "-":
				stack.push(stack.sub())
			case "*":
				stack.push(stack.mul())
			case "/":
				stack.push(stack.div())
			default:
				t, _ := strconv.Atoi(token)
				stack.push(t)
		}
		stack.show()
	}
	return stack.top()
}

func (s *Stack) add() int{
	if len(s.st)<2 {
		return 0;
	}
	l := len(s.st)
	x := s.st[l-2]+s.st[l-1]
	s.pop()
	s.pop()
	return x
}

func (s *Stack) sub() int{
	if len(s.st)<2 {
		return 0;
	}
	l := len(s.st)
	x := s.st[l-2]-s.st[l-1]
	s.pop()
	s.pop()
	return x
}

func (s *Stack) mul() int{
	if len(s.st)<2 {
		return 0;
	}
	l := len(s.st)
	x := s.st[l-2]*s.st[l-1]
	s.pop()
	s.pop()
	return x
}

func (s *Stack) div() int{
	if len(s.st)<2 {
		return 0;
	}
	l := len(s.st)
	x:= int(math.Floor(float64(s.st[l-2]/s.st[l-1])))
	s.pop()
	s.pop()
	return x
}


