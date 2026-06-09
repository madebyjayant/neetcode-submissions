type TimeMap struct {
	m map[string][]pair
}

type pair struct {
	timestamp int
	value string
}

func Constructor() TimeMap {
	return TimeMap{
		m : make(map[string][]pair, 0),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key]= append(this.m[key], pair{timestamp,value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	fmt.Println(this.m)
	// fetch the list of available values
	arr, ok := this.m[key]
	if !ok || len(arr)==0 || arr[0].timestamp>timestamp{
		return ""
	}

	l, r := 0, len(arr)-1
	ans := ""
	for l<=r {
		m := (l+r)/2
		if arr[m].timestamp==timestamp{
			return arr[m].value
		}

		if arr[m].timestamp>timestamp{
			r = m-1
		}else{
			ans = arr[m].value
			l = m+1
		}
	}

	return ans

}
