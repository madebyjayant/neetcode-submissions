func hasDuplicate(nums []int) bool {
    mymap := make(map[int]struct{})
    for _, num := range nums {
        if _, ok := mymap[num]; !ok {
            mymap[num]=struct{}{}
            continue
        }
        return true
    }
    return false;
}
