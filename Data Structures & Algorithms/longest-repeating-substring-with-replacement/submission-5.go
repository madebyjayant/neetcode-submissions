func characterReplacement(s string, k int) int {
	mf := 0
	l,r := 0,0
	ans := 0
	seen := make(map[byte]int, 26)
	for r<len(s) {
		seen[s[r]]++
		if mf < seen[s[r]] {mf = seen[s[r]]}
		if k >= r-l+1-mf {
			ans = max(ans, r-l+1)
		}else{
			seen[s[l]]--
			l++
		}
		r++
	}

	return ans
}
