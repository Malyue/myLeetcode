package main

// 整数数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。
// 每一天你都可以选择是否 购买/出售 股票，你在任何时候最多只能持有 一股 股票，可以先购买，然后同一天出售
// 设计一个算法来计算你所能获取的最大利润。
// 例如 [7,1,5,3,6,4] 则结果为 5-1 + 6 - 3 =7
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// 动态规划
	// 每天其实有两种情况，即完成交易之后有股票和没股票
	// dp[i][0]表示第i天交易完之后手里没有股票的最大利润，dp[i][1]表示第i天交易完之后手里持有股票的最大利润
	// 当天交易完之后手里没有股票的可能就是当天没交易，或者卖掉了，如果是当天没交易，那当天没有利润，所以利润就是前一天手里没有股票的利润
	// 如果是当天卖掉了，那这个时候当天没有股票的利润要取前一天手里有股票的利润加上当天股票能卖的价格，这两种情况取最大
	// dp[i][0] = max(dp[i-1][0],dp[i-1][1]+prices[i]);

	// 当天交易完之后手里有股票的可能也有两种
	// 一种是当天没有进行任何交易，所以前天已经持有股票了
	// 另外一种是当天买入股票,说明前一天没有股票
	// dp[i][1] = max(dp[i-1][1],dp[i-1][0]-prices[i])

	// 边界条件，dp[0][1] = prices[0],dp[0][0] = 0
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
