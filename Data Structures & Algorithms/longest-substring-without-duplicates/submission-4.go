func lengthOfLongestSubstring(s string) int {
    lastSeen := make(map[byte]int, 0)
    left,best := 0,0
    for right := 0; right < len(s); right++ {
        if idx, ok := lastSeen[s[right]]; ok && idx >= left {
            left = idx + 1
        }
        lastSeen[s[right]] = right
        if right-left+1 > best {
            best = right - left + 1
        }
    }
    return best
}
