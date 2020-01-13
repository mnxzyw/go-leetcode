package main

import "fmt"

type BaseInfo struct {
	// 物品的重量 舍弃数组的0位置元素
	Weight []int
	// 物品的价值 舍弃数组的0位置元素
	Value []int
	// 动态规划网格
	DP [][]int
	// 用于回溯选择的商品 selected[i]=1 表示放进背包，0表示不放进背包
	Selected []int
}

// 二维表， 横向为 背包容量增加，纵向为 物品种类增加
// n 件物品, 质量分别是 w1，w2,...,wn，价值分别为 v1, v2,...,vn
// 背包容量为 C
/*
* Dynamic
*
* PARAMS:
*   - c: 背包容量
*
* RETURNS:
*   最大总价值
 */
func (bs *BaseInfo) Dynamic(c int) (max int) {
	n := len(bs.Value)
	bs.DP = make([][]int, n)

	// 纵向 可选物品种类 i [0,n)
	// 即横向为 背包容量动态变化 j [0,c]
	// j 为背包容量
	for i := 0; i < n; i++ {
		bs.DP[i] = make([]int, c+1)
		for j := 0; j <= c; j++ {
			if i == 0 || j == 0 {
				bs.DP[i][j] = 0
			} else {
				bs.DP[i][j] = bs.maxValue(i, j)
			}
		}
	}
	return bs.DP[n-1][c]
}

// 判断当前单元格最大价值方法
func (bs *BaseInfo) maxValue(i, j int) int {
	// 当前商品无法放入背包，返回当前背包所能容纳的最大价值
	if j < bs.Weight[i] {
		return bs.DP[i-1][j]
	}
	// 可放进背包时候，计算放入当前商品后的最大价值
	curr := bs.Value[i] + bs.DP[i-1][j-bs.Weight[i]]
	// 与当前商品不放进背包的最大价值比较，看是否应该放进背包
	if curr >= bs.DP[i-1][j] {
		return curr
	}
	return bs.DP[i-1][j]
}

// 回溯已选择商品
func (bs *BaseInfo) findSelected() {
	bs.Selected = make([]int, len(bs.Value))
	j := len(bs.DP[0]) - 1
	for i := len(bs.Value) - 1; i > 0; i-- {
		if bs.DP[i][j] > bs.DP[i-1][j] {
			bs.Selected[i] = 1
			j = j - bs.Weight[i]
		}
	}
}

func main() {
	bs := BaseInfo{
		Weight:   []int{0, 1, 4, 3},
		Value:    []int{0, 1500, 3000, 2000},
		DP:       nil,
		Selected: nil,
	}
	max := bs.Dynamic(7)
	fmt.Println(max)
	bs.findSelected()
	fmt.Println(bs.Selected)
}
