func minWindow(s string, t string) string {
	var need, have [128]int

	for idx, _ := range t{
		need['z'-t[idx]]++
	}
	required := 0
	for _, val := range need{
		if val!=0{required++}
	}

	formed := 0
	l,r := 0,0
	leftBest := 0
	bestLen := -1
	for r<len(s){
		have['z'-s[r]]++

		if need['z'-s[r]]>0 && have['z'-s[r]]==need['z'-s[r]]{
			formed++
		}

		for formed==required{
			if bestLen ==-1 || bestLen>r-l+1{
				bestLen = r-l+1
				leftBest = l
			}
			have['z'-s[l]]--
			if need['z'-s[l]]>0 && have['z'-s[l]]<need['z'-s[l]]{
				formed--
			}
			l++
		}
		r++
	}
	if bestLen==-1{
		return ""
	}
	return s[leftBest: leftBest+bestLen]
}
