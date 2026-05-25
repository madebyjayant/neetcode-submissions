func maxProfit(prices []int) int {
    maxProfit,minPrice := 0,101
    for _, v := range prices {
        minPrice = min(v,minPrice)
        maxProfit = max(v-minPrice,maxProfit)
    }
    return maxProfit
}