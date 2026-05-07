func groupAnagrams(strs []string) [][]string {
	var finalResult [][]string
	mymap := make(map[[26]int][]string);

	for _, str := range strs {
		key := generateKey(str);
		mymap[key] = append(mymap[key], str)
	}

	for _, v := range mymap{
		finalResult = append(finalResult, v)
	}
	return finalResult

}

func generateKey(s string) [26]int{
	var key [26]int
	for _, c := range s{
		key[c-'a']++
	}
	return key
}