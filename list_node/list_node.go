package list_node

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	ln := new(ListNode)
	ln.Val = val

	return ln
}

func NewListNodeRec(val int, next *ListNode) *ListNode {
	ln := new(ListNode)
	ln.Val = val
	ln.Next = next

	return ln
}

func (ln *ListNode) ToString() string {
	if ln == nil {
		return "nil"
	}

	return fmt.Sprintf("{val: %d; next: %s}", ln.Val, ln.Next.ToString())
}

func ArrayToList(arr []int) *ListNode {
	l := len(arr)
	if l < 1 {
		return nil
	}

	i := l - 1
	head := NewListNode(arr[i])
	i -= 1
	for i >= 0 {
		ln := NewListNode(arr[i])
		i -= 1
		ln.Next = head
		head = ln
	}

	return head
}

func ArrayToListCyc(arr []int, cycle int) *ListNode {
	l := len(arr)
	if l < 1 {
		return nil
	}

	i := l - 1
	head := NewListNode(arr[i])
	i -= 1
	tail := head
	for i >= 0 {
		ln := NewListNode(arr[i])
		i -= 1
		if (cycle >= 0) && (i+1 == cycle) {
			tail.Next = ln
		}
		ln.Next = head
		head = ln
	}

	return head
}

func ArraysToIntersectedLists(listA, listB []int, skipA, skipB int) (headA, headB *ListNode) {
	la := len(listA)
	if (la < 1) || (la <= skipA) {
		return nil, nil
	}

	lb := len(listB)
	if (lb < 1) || (lb <= skipB) {
		return nil, nil
	}

	i := la - 1
	tail := NewListNode(listA[i])
	i -= 1
	for i >= skipA {
		ln := NewListNode(listA[i])
		i -= 1
		ln.Next = tail
		tail = ln
	}

	i = skipA - 1
	headA = tail
	for i >= 0 {
		ln := NewListNode(listA[i])
		i -= 1
		ln.Next = headA
		headA = ln
	}

	i = skipB - 1
	headB = tail
	for i >= 0 {
		ln := NewListNode(listB[i])
		i -= 1
		ln.Next = headB
		headB = ln
	}

	return headA, headB
}

// 141. Linked List Cycle
// Floyd's Cycle Finding Algorithm
func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
