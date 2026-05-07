func isAnagram(s string, t string) bool {
	n := len(s)
	m := len(t)

	if m!=n{return false}

	mymap := make(map[rune]int)

	for _, v := range s {
		mymap[v]++;
	}

	for _, v := range t {
		if mymap[v]==0 {
			return false;
		}
		mymap[v]--;
	}

	return true
}
