package main

import "fmt"

//653. 两数之和 IV - 输入 BST
//给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
//
//案例 1:
//
//输入:
//5
/// \
//3   6
/// \   \
//2   4   7
//
//Target = 9
//
//输出: True
//
//
//案例 2:
//
//输入:
//5
/// \
//3   6
/// \   \
//2   4   7
//
//Target = 28
//
//输出: False
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ms = make([]int, 0)

func traverse(root *TreeNode) {
	if root != nil {
		traverse(root.Left)
		ms = append(ms, root.Val)
		traverse(root.Right)
	}
}

func findTarget(root *TreeNode, k int) bool {
	traverse(root)
	ln := len(ms)
	if k <= ms[ln-1]+ms[ln-2] && k >= ms[0]+ms[1] {
		for i := 0; i < ln-1; i++ {
			for j := i + 1; j < ln; j++ {
				if ms[i]+ms[j] == k {
					return true
				}
			}
		}
		return false
	} else {
		return false
	}
}

func genBST() *TreeNode {
	var root *TreeNode
	var tmp *TreeNode
	tmp = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root = tmp
	tmp = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Left = tmp
	tmp = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Right = tmp
	return root
}

func main() {
	bst := genBST()
	rsp := findTarget(bst, 4)
	fmt.Println("rsp = ", rsp)
}
