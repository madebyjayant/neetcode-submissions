func isAnagram(s string, t string) bool {
	n := len(s)
	m := len(t)

	if m!=n{return false}

	var arr [26]int

	for i, v := range s {
		arr[v - 'a']++
		arr[t[i] - 'a']--;
	}

	for _, v := range arr {
		if v!=0 {
			return false;
		}
	}

	return true
}
