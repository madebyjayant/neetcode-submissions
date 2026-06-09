func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2
	if len(A) > len(B) {
		A,B = B,A
	}

	n,m := len(A), len(B)
	half := (n+m+1)/2

	l,r := 0, n

	for l<=r {
		i := (l+r)/2
		j := half-i

		Aright := math.MaxInt64
		if i < len(A) {
			Aright = A[i]
		}

		Aleft := math.MinInt64
		if i > 0 {
			Aleft = A[i-1]
		}

		Bright := math.MaxInt64
		if j < len(B) {
			Bright = B[j]
		}
		Bleft := math.MinInt64
		if j > 0 {
			Bleft = B[j-1]
		}

		if Aleft <= Bright && Bleft <= Aright {
			if (n+m)%2!=0{
				return float64(max(Aleft, Bleft))
			}
			return (float64(max(Aleft, Bleft)) + float64(min(Aright, Bright))) / 2.0
		}else if Aleft > Bright{
			r = i-1
		}else{
			l = i+1
		}

	}
	return -1
}
