func minWindow(s string, t string) string {
	var need, have [128]int

	for idx, _ := range t {
		need[t[idx]]++
	}

	required := 0
	for _, v := range need {
		if v!=0{required++}
	}

	formed :=0 
	l,bestLeft:=0,0
	bestLength:=-1
	for r:=0;r<len(s);r++{
		c := s[r]
		have[c]++
		if need[c]>0 && have[c]==need[c]{
			formed++
		}
		for required==formed {
			if bestLength==-1 || bestLength>len(s[l:r+1]){
				bestLeft = l
				bestLength=r-l+1
			}

			have[s[l]]--
			if need[s[l]]>0 && have[s[l]]<need[s[l]]{
				formed--
			}
			l++
		}
	}
	if bestLength==-1{
		return ""
	}
	return s[bestLeft:bestLeft+bestLength]
}
