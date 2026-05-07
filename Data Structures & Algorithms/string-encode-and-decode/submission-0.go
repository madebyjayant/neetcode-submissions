type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var r strings.Builder
	for _, st := range strs {
		r.WriteString(fmt.Sprintf("%s#", strconv.Itoa(len(st))))
		r.WriteString(st)
	}
	fmt.Println(r.String())
	return r.String()
}

func (s *Solution) Decode(encoded string) []string {
	var res []string

	for i:=0; i<len(encoded); {
		j := i
		for encoded[j]!='#'{
			j++
		}

		m, _ := strconv.Atoi(encoded[i:j])
		res = append(res, encoded[j+1:j+1+m])
		i =  j + 1 + m
	}

	return res
}
