import (
	"slices"
)


func carFleet(target int, position []int, speed []int) int {
	cars := make([][2]int, 0)

	for idx, val := range position {
		cars = append(cars, [2]int{val, speed[idx]})
	}

	// sort based on first index
	slices.SortFunc(cars, func (a, b [2]int) int {
		return a[0]-b[0]
	})

	// create a stack
	stack := []float64{}
	for idx:=len(cars)-1;idx>=0;idx-- {
		tcc := float64((target-cars[idx][0]))/float64(cars[idx][1])
		if len(stack)==0 {
			stack = append(stack, tcc)
			continue
		}

		topSt := stack[len(stack)-1]
		if topSt < tcc {
			stack = append(stack, tcc)
		}
	}

	return len(stack)

}
