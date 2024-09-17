package binary_trees

func Walk(curr *Node, path []int) []int {
	if curr == nil {
		return path
	}
	//recurse
	//pre
	path = append(path, curr.Data)
	path = Walk(curr.Left, path)
	path = Walk(curr.Right, path)

	return path
}

func PreOrderSearch(root *Node) []int {
	return Walk(root, []int{})
}
