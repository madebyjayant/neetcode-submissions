func characterReplacement(s string, k int) int {
    left := 0
    best := 0
    seen := make(map[byte]int, 26)
    freq := 0
    for right := 0;right < len(s); right++ {
        seen[s[right]]++
        if freq < seen[s[right]]{
            // this means the current char has become more frequent
            freq = seen[s[right]]
        }
        windowSize := right-left+1
        if k >= windowSize-freq{
            best = max(best, windowSize)
        }else{
            seen[s[left]]--
            left++
        }
    }
    return best
}