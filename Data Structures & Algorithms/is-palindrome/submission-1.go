func isPalindrome(s string) bool {
	if len(s)==0{return true}
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	ns := strings.ToLower(re.ReplaceAllString(s, ""))
	j:=len(ns)-1
	i:=0
	for i<len(s) && j>=0 && i<=j{
		if ns[i]!=ns[j]{
			return false
		}
		i++;j--
	}
	return true
}
