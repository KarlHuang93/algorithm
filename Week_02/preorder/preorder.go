package preorder

type Node struct {
	Val      int
	Children []*Node
}

var res []int

func preorder(root *Node) []int {
	res = []int{}
	dfs(root)
	return res
}

func dfs(root *Node) {
	if root != nil {
		res = append(res, root.Val)
		for _, n := range root.Children {
			dfs(n)
		}
	}
}


func preorder1(root *Node) []int {
	var res []int
	var stack = []*Node{root}
	for 0 < len(stack) {
		for root != nil {
			res = append(res, root.Val) //前序输出
			if 0 == len(root.Children) {
				break
			}
			for i := len(root.Children) - 1; 0 < i; i-- {
				stack = append(stack, root.Children[i]) //入栈
			}
			root = root.Children[0]
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}

