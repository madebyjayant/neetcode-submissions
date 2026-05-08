func maxArea(heights []int) int {
	if len(heights)==1{return 0}
	if len(heights)==2{
		return min(heights[0],heights[1]);
	}

	left:=0
	maxArea:=0
	for left<len(heights)-1 {
		right := left+1

		for right<len(heights){
			currArea := (right-left)*min(heights[left], heights[right]);
			if currArea>maxArea{
				maxArea=currArea
			}
			right++
		}
		left++;
	}
	return maxArea
}