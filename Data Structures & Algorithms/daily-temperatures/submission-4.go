type stack struct {
    c []int
}

func New(n int) *stack {
    return &stack{
        c : make([]int,n),
    }
}

func (s *stack) push(n int) {
    s.c = append(s.c, n)
}

func (s *stack) pop() int {
    t := s.c[len(s.c)-1]
    s.c = s.c[:len(s.c)-1]
    return t
}

func (s *stack) isEmpty() bool {
    return len(s.c)==0
}

func (s *stack) top() int{ 
    return s.c[len(s.c)-1]
}

func dailyTemperatures(temperatures []int) []int {
    st := New(len(temperatures))
    l := len(temperatures)
    result := make([]int, l)
    for i, temp := range temperatures {
        if st.isEmpty(){
            st.push(i)
            continue
        }

        for (!st.isEmpty() && temperatures[st.top()]<temp){
            t := st.pop()
            result[t] = i-t
        }
        st.push(i)
    }
    result[st.top()]=0
    return result
}
