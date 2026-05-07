func groupAnagrams(strs []string) [][]string {
	var finalResult [][]string
	mymap := make(map[[26]int][]string);

	for _, str := range strs {
		var key [26]int
		for _, c := range str{
			key[c-'a']++
		}
		mymap[key] = append(mymap[key], str)
	}

	for _, v := range mymap{
		finalResult = append(finalResult, v)
	}
	return finalResult

}