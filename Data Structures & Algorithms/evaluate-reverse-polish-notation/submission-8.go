type stack struct {
    c []int
}

func New(n int) *stack{
    return &stack{
        c: []int{},
    }
}

func (s *stack) push(n int) {
    s.c = append(s.c, n)
}

func (s *stack) pop() int{
    if len(s.c)>=1{
        t := s.c[len(s.c)-1]
        s.c = s.c[:len(s.c)-1]
        return t
    }
    return -1
}

func evalRPN(tokens []string) int {
    s := New(len(tokens))
    for _, token := range tokens {
        if isNumber(token) {
            num,_:= strconv.Atoi(token)
            s.push(num)
        }else{
            switch{
                case token=="+":
                    f := s.pop()
                    l := s.pop()
                    t := l+f
                    s.push(t)
                case token=="-":
                    f := s.pop()
                    l := s.pop()
                    t := l-f
                    s.push(t)
                case token=="*":
                    f := s.pop()
                    l := s.pop()
                    t := l*f
                    s.push(t)
                case token=="/":
                    f := s.pop()
                    l := s.pop()
                    t := l/f
                    s.push(t)
                default: 
                    return -1
            }
        }
    }
    return s.pop()
}

func isNumber(s string) bool {
    _, err := strconv.Atoi(s)
    if err==nil{
        return true
    }
    return false
}