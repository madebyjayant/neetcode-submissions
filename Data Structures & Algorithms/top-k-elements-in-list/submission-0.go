func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++;
	}
	
	var res []int

	for k, _ := range freq {
		res = append(res, k)
	}
 
	sort.Slice(res, func(i, j int) bool {
		if freq[res[i]] != freq[res[j]] {
			return freq[res[i]] > freq[res[j]]
		}
		return res[i] < res[j]
	})

	if k>n{k = len(res)}
	return res[:k];
}
