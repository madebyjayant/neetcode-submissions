import (
	"slices"
)
func groupAnagrams(strs []string) [][]string {
	var finalResult [][]string
	mymap := make(map[string][]string);

	for _, str := range strs {
		b := []byte(str)
		slices.Sort(b)
		key := string(b)
		mymap[key] = append(mymap[key], str)
	}

	// compile
	for _, v := range mymap{
		finalResult = append(finalResult, v)
	}
	return finalResult

}