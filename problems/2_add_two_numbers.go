package main

import (
	"fmt"
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 数字倒序的列表
func calculateList(l *ListNode) *big.Int {
	n := new(big.Int)
	bit := big.NewInt(1)
	for tmp := l; tmp != nil; tmp = tmp.Next {
		value := big.NewInt(int64(tmp.Val))
		//个位乘以1
		if tmp != l {
			bit.Mul(bit, big.NewInt(10))
		}
		//以后每次乘以10
		n.Add(n, value.Mul(value, bit))
	}
	return n
}

func bigLen(s string) *big.Int {
	n := new(big.Int)
	for range s {
		n.Add(n, big.NewInt(1))
	}
	return n
}

//根据数字生成列表
func genList(n *big.Int) *ListNode {
	var head *ListNode
	var tail = new(ListNode)
	s := n.String()
	var i int
	for _, tmp := range s {
		i, _ = strconv.Atoi(string(tmp))
		node := &ListNode{i, nil}
		if head == nil {
			head = node
		} else {
			tail.Next = node
		}
		tail = node
	}
	return head
}

func reverseList(l *ListNode) *ListNode {
	var cur = l
	var prev *ListNode
	// 这里无需临时变量是因为
	// 执行顺序从左到右，但左边被赋值的变量在表达式的下一行才生效。
	// 即 cur.Next 在本行还是原本未赋值之前的cur.Next
	// 如 a, b = b, a+1 可理解成
	//   temp_a = a
	//   a = b
	//   b = temp_a + 1
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	// 下面这种循环方式需要用到临时变量
	//for cur != nil {
	//	tmp := cur.Next
	//	cur.Next = prev
	//	prev = cur
	//	cur = tmp
	//}
	return prev
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := calculateList(l1)
	n2 := calculateList(l2)

	n3 := new(big.Int)
	n3 = n3.Add(n1, n2)
	head := genList(n3)

	return reverseList(head)
}

func printList(l *ListNode) {
	for tmp := l; tmp != nil; tmp = tmp.Next {
		fmt.Println("value", tmp.Val)
	}
}

func main() {
	a := genList(big.NewInt(123))
	b := genList(big.NewInt(456))
	c := addTwoNumbers(a, b)

	printList(c)
}
