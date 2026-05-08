func maxArea(heights []int) int {
	if len(heights)==1{return 0}
	if len(heights)==2{
		return min(heights[0],heights[1]);
	}

	left,right:=0, len(heights)-1
	maxArea:=0

	for left<right{
		currArea:=(right-left)*min(heights[left],heights[right])
		fmt.Println(currArea)
		if currArea>maxArea{maxArea=currArea}
		if heights[left]<heights[right]{
			left++
			continue
		}
		right--
	}
	return maxArea
}