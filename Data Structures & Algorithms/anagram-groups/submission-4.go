func groupAnagrams(strs []string) [][]string {
    var result [][]string
    mymap := make(map[[26]int][]string, len(strs))
    for _, str := range strs{
        var key [26]int
        for _, v := range str{
            key[v-'a']++
        }
        mymap[key] = append(mymap[key], str)
    }
    for _, elem := range mymap{
        result = append(result, elem)
    }
    return result
}