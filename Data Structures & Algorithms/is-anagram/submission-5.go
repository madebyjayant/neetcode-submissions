func isAnagram(s string, t string) bool {
	n := len(s)

    m := len(t)

    if m!=n{return false}

    var seen [26]int

    for j,v := range s {
        seen[t[j]-'a']++
        seen[v-'a']--
    }

    for _, v := range seen {
        if v!=0{
            return false
        }
    }

    return true
}
