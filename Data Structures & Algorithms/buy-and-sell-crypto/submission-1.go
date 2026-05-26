func maxProfit(prices []int) int {
	sp := prices[0]
	mp := 0
	for i:=0;i<len(prices);i++{
		if prices[i]< sp{
			sp = prices[i]
		}
		if prices[i]-sp > mp {
			mp = prices[i]-sp
		}
	}

	return mp
}
