func trap(height []int) int {
	l,r := 0, len(height)-1
	ml, mr := -1,-1
	mw := 0
	for l<r {
		if height[l]<height[r]{
			if ml<height[l]{
				ml=height[l]
				continue
			}
			w := ml-height[l]
			if w>0 {
				mw+=w
			}
			l++
		}else{
			w := mr-height[r]
			if height[r]>mr{
				mr=height[r]
				continue
			}
			if w>0{
				mw+=w
			}
			r--
		}
	}

	return mw

}
