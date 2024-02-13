package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	res := addTwoNumbers(l1, l2)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1 := l1
	cur2 := l2

	var overhead int

	curNode := &ListNode{}
	res := curNode

	for {
		curNode.Val = (cur1.Val + cur2.Val + overhead) % 10
		overhead = (cur1.Val + cur2.Val + overhead) / 10

		if cur1.Next == nil && cur2.Next == nil && overhead == 0 {
			break
		}

		if cur1.Next == nil {
			cur1.Next = &ListNode{}
		}
		cur1 = cur1.Next

		if cur2.Next == nil {
			cur2.Next = &ListNode{}
		}
		cur2 = cur2.Next

		nextNode := &ListNode{}
		curNode.Next = nextNode
		curNode = nextNode
	}

	return res
}
