func checkInclusion(s1 string, s2 string) bool {
	
	if len(s1)>len(s2){return false}

	var count1, count2 [26]int
	for idx, _ := range s1{
		count1[s1[idx]-'a']++
		count2[s2[idx]-'a']++
	}

	l,r := 0,len(s1)-1
	for r<len(s2){
		if count1==count2{
			return true
		}
		count2[s2[l]-'a']--
		l++
		r++
		if r>len(s2)-1{return false}
		count2[s2[r]-'a']++
	}
	return false
}
