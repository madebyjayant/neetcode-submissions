func trap(height []int) int {
    left,right := 0,len(height)-1
    max := 0
    maxLeft, maxRight := 0,0
    for left<right {
        if height[left]<height[right] {
            if maxLeft > height[left] {
                max+=maxLeft-height[left]
            }else {
                maxLeft = height[left]
            }
            left++
        }else{
            if maxRight > height[right] {
                max+=maxRight-height[right]
            }else {
                maxRight = height[right]
            }
            right--
        }
    }

    return max
}
