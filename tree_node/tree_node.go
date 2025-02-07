package tree_node

import (
	"fmt"
	"os"
	"valpere/leetcode/queue"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode() *TreeNode {
	return new(TreeNode)
}

func NewTreeNodeVal(val int) *TreeNode {
	tn := new(TreeNode)
	tn.Val = val
	return tn
}

func NewTreeNodeAll(val int, left, right *TreeNode) *TreeNode {
	tn := new(TreeNode)
	tn.Val = val
	tn.Left = left
	tn.Right = right
	return tn
}

func ArrayToTree(arr []any) (head *TreeNode) {
	if (arr == nil) || (len(arr) == 0) {
		return nil
	}

	q := queue.NewQueue[*TreeNode](len(arr))

	n := len(arr)
	r := 0
	head = NewTreeNodeVal(arr[r].(int))
	r += 1
	q.Enqueue(head)
	for r < n {
		root, _ := q.Dequeue()

		il := r
		r += 1
		intL := arr[il]
		if intL != nil {
			root.Left = NewTreeNodeVal(intL.(int))
			q.Enqueue(root.Left)
		}

		if r >= n {
			break
		}

		ir := r
		r += 1
		intR := arr[ir]
		if intR != nil {
			root.Right = NewTreeNodeVal(intR.(int))
			q.Enqueue(root.Right)
		}

	}

	return head
}

func TreeToArray(root *TreeNode) (arr []any) {
	return treeToArray(root, false)
}

func treeToArray(root *TreeNode, rl bool) (arr []any) {
	ans := make([]any, 0)
	if root == nil {
		return ans
	}

	q := queue.NewQueue[*TreeNode](0)
	q.Enqueue(root)

	for !q.IsEmpty() {
		node, _ := q.Dequeue()
		if node == nil {
			ans = append(ans, nil)
		} else {
			ans = append(ans, node.Val)
			if rl {
				q.Enqueue(node.Right)
				q.Enqueue(node.Left)
			} else {
				q.Enqueue(node.Left)
				q.Enqueue(node.Right)
			}
		}
	}

	i := len(ans) - 1
	for ; (i >= 0) && (ans[i] == nil); i-- {
	}

	ans = ans[:i+1]

	return ans
}

func (tn *TreeNode) ToString() string {
	if tn == nil {
		return "nil"
	}

	if (tn.Left == nil) && (tn.Right == nil) {
		return fmt.Sprintf("{val: %d}", tn.Val)
	}

	return fmt.Sprintf("{val: %d, left: %s, right: %s}", tn.Val, tn.Left.ToString(), tn.Right.ToString())
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

var nodeId int = 0

func (tn *TreeNode) ToMermaid(mmdFile string) {
	if tn == nil {
		return
	}
	f, err := os.Create(mmdFile)
	CheckError(err)
	defer f.Close()

	f.WriteString("graph TB\n")

	nodeToMermaid(tn, f)
	return
}

func nodeToMermaid(node *TreeNode, f *os.File) {
	if node == nil {
		return
	}

	numN := nodeId
	if node.Left != nil {
		nodeId++
		numL := nodeId
		f.WriteString(fmt.Sprintf("    node%03d((\"%v\"))-->node%03d((\"%v\"))\n", numN, node.Val, numL, node.Left.Val))
		nodeToMermaid(node.Left, f)
	}
	if node.Right != nil {
		nodeId++
		numR := nodeId
		f.WriteString(fmt.Sprintf("    node%03d((\"%v\"))-->node%03d((\"%v\"))\n", numN, node.Val, numR, node.Right.Val))
		nodeToMermaid(node.Right, f)
	}

}
