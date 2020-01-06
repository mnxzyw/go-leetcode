package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 作者：wellszhane
// 链接：https://leetcode-cn.com/problems/merge-k-sorted-lists/solution/golang-fen-zhi-fa-by-wellszhane/
// *惊奇之处在于归并排序的合并步骤竟然也可用递归实现*
func mergeTwoList(L1 *ListNode, L2 *ListNode) *ListNode {
	if nil == L1 {
		return L2
	}
	if nil == L2 {
		return L1
	}
	if L1.Val > L2.Val {
		L2.Next = mergeTwoList(L1, L2.Next)
		return L2
	} else {
		L1.Next = mergeTwoList(L1.Next, L2)
		return L1
	}
}

func merge(lists []*ListNode, listFrom int, listTo int) *ListNode {
	if listFrom == listTo {
		return lists[listFrom]
	}
	mid := (listFrom + listTo) / 2
	return mergeTwoList(merge(lists, listFrom, mid), merge(lists, mid+1, listTo))
}

// 注意 list 为数组，从 0 开始
func mergeKLists(lists []*ListNode) *ListNode {
	if 0 == len(lists) {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func genListFromInt(n int) *ListNode {
	var head *ListNode
	var tail = new(ListNode)
	s := strconv.Itoa(n)
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

func genLists() []*ListNode {
	lists := make([]*ListNode, 5)
	l1 := genListFromInt(2378)
	l2 := genListFromInt(1247)
	l3 := genListFromInt(35789)
	l4 := genListFromInt(13579)
	l5 := genListFromInt(2468)
	lists = append(lists, l1, l2, l3, l4, l5)
	return lists
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func main() {
	LS := genLists()
	L := mergeKLists(LS)
	printList(L)
}
