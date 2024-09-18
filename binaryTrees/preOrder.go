package binary_trees

func WalkInPreOder(curr *Node, path []int) []int {
	if curr == nil {
		return path
	}
	//recurse
	//pre
	path = append(path, curr.Data)
	path = WalkInPreOder(curr.Left, path)
	path = WalkInPreOder(curr.Right, path)

	return path
}

func PreOrderSearch(root *Node) []int {
	return WalkInPreOder(root, []int{})
}
