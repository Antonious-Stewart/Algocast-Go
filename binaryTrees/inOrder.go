package binary_trees

func WalkInOrder(curr *Node, path []int) []int {
	if curr == nil {
		return path
	}
	//recurse
	//pre
	path = WalkInOrder(curr.Left, path)
	path = append(path, curr.Data)
	path = WalkInOrder(curr.Right, path)

	return path
}

func InOrderSearch(root *Node) []int {
	return WalkInOrder(root, []int{})
}
