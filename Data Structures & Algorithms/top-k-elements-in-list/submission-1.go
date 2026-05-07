func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++;
	}

	buckets := make([][]int, n+1)
	for num, f := range freq {
		buckets[f] = append(buckets[f],num)
	}

	result := make([]int, 0, k)
	for i:=len(buckets)-1;i>=0 && len(result)<k; i-- {
		result = append(result, buckets[i]...)
	}
	return result
}
