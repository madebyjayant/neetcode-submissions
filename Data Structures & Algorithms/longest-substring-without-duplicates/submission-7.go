func lengthOfLongestSubstring(s string) int {
	if len(s)==1{
		return 1
	}

	seen := make(map[byte]int, 0)
	l,r:=0,0
	mx := 0
	for r<len(s) && l<=r {
		if _, ok := seen[s[r]]; ok && seen[s[r]]>0{
			mx = max(mx, r-l)
			seen[s[l]]--
			l++
			continue
		}
		seen[s[r]]++
		r++
	}
	return max(mx, r-l)
}
