package dfs
import "fmt"

// 二叉树的侧面视角
// @Author: hyc
// @Description: 来自云账户笔试题
// @Date: 2021/9/16 22:08
func main() {

	var (
		root = new(TreeNode)
	)
	root.Val = 1
	root.Left = &TreeNode{
		Val:   2,
	}

	root.Right = &TreeNode{
		Val:   3,
	}

	root.Left.Right = &TreeNode{
		Val:   5,
	}

	root.Right.Right = &TreeNode{
		Val:   4,
	}
	fmt.Printf("%+v\n",root)

	ans := RightSideView(root)

	fmt.Printf("%+v\n",ans)
}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func RightSideView( root *TreeNode ) []int {
	// write code here
	if root == nil {
		return []int{}
	}

	hight := 0
	max := 0

	tmp := make([]*TreeNode,0)
	ans := make([]int,0)

	var solver func(root *TreeNode)
	solver = func(root *TreeNode) {

		if root != nil {
			hight++
			if hight > max {
				max = hight
			}

			solver(root.Left)
			solver(root.Right)
			tmp = append(tmp,root)
			hight--
		}
	}
	solver(root)

	for i := 0;i <= len(tmp);i++{
		if i != max {
			ans = append(ans,tmp[len(tmp)-1-i].Val)
		} else {
			break
		}
	}
	return ans
}