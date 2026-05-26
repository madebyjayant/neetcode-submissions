type Stack interface {
	Push(rune) rune
	Pop () rune
	Peek() rune
}

type StackImpl struct {
	val []rune
}

func newStack() *StackImpl{
	return &StackImpl{}
}

func (s *StackImpl) Push(a rune) rune {
	s.val = append(s.val, a)
	return a
}

func (s *StackImpl) Pop() (rune,bool) {
	if len(s.val)==0{
		return 'a', false
	}
	li := len(s.val)-1
	x:=s.val[li]
	s.val = s.val[:li]
	return x, true
}

func (s *StackImpl) Peek() rune {
	if len(s.val)!=0{
		return s.val[len(s.val)-1]
	}
	return -1
}

func (s *StackImpl) IsEmpty() bool {
	return len(s.val)==0
}
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	
	st := newStack()
	for _, c := range s{
		switch c {
			case ')':
				if t, ok := st.Pop(); ok && t=='('{
					continue
				}
				return false
			case '}':
				if t, ok := st.Pop(); ok && t=='{'{
						continue
					}
					return false
			case ']':
				if t, ok := st.Pop(); ok && t=='['{
					continue
				}
				return false

			default :
				st.Push(c)
		}
	}
	return st.IsEmpty()
}
