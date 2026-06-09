import (
	"slices"
)
func minEatingSpeed(piles []int, h int) int {
	slices.Sort(piles)
	l, r := 1, slices.Max(piles)

	ans := r
	for l<=r {
		m := (l+r)/2

		// for mid pile, check if the total time <=h
		t := 0
		for _, val := range piles {
			t+=int(math.Ceil(float64(val) / float64(m)))
		}
		if t>h{
			l = m+1
		}else{
			ans = m
			r = m-1
		}
	}
	return ans
}
