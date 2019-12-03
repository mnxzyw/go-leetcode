package main

// 二维表， 横向为 背包容量增加，纵向为 物品种类增加
// n 件物品, 质量分别是 w1，w2,...,wn，价值分别为 v1, v2,...,vn
// 背包容量为 C
func two_dimensional(n int, c int) int {
	dp := make([][]int, n)
	// 横向填表，
	// 即纵向为 先只考虑一种物品，再考虑两种物品
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, c)
		for j := 0; i < c; i++ {

		}
	}
}

func main() {

}
