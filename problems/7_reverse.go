package main

import "fmt"

// 可以简单理解为 go python "%" 为 取模，cpp java php "%" 为 取余
// 因为题干中说"假设我们的环境只能存储得下 32 位的有符号整数"，是否意味着不能申请int64的变量
// go 的 int 长度随操作系统变化, 而题干中x却声明为int，所以感觉题干也不太清晰
func reverse(x int) int {
	var y, tmp int
	var b bool
	if x >= 0 {
		b = true
	} else {
		b = false
	}
	for x != 0 {
		tmp = y*10 + x%10
		// 正数加法不会减小，负数加法不会增大
		if b == true && tmp < y ||
			b == false && tmp > y {
			return 0
		}
		x = x / 10
		y = tmp
	}
	// 理论上32位环境根本不用加下面的判断
	if y > ((1<<31)-1) || y < (-(1 << 31)) {
		return 0
	}
	return y
}

func main() {
	x := -12345
	y := reverse(x)
	fmt.Println(y)
}
