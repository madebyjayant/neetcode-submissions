func checkInclusion(s1 string, s2 string) bool {
    if len(s1)>len(s2){
        return false
    }

    var count1, count2 [26]int

    for idx, _ := range s1 {
        count1[s1[idx]-'a']++
        count2[s2[idx]-'a']++
    }
    
    l, r := 0,len(s1)-1
    for l<=r {
        if count1==count2{
            return true
        }
        
        // remove the first of the window
        count2[s2[l]-'a']--
        l++
        // add the next
        r++
        if r>len(s2)-1 {break}
        count2[s2[r]-'a']++
    }
    return false
}
