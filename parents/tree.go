package parents

type TreeNode struct {
	Key    int
	Parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
}

func Int(value int) *int {
	return &value
}

func NewTreeWithLeafs(leafs []int) *TreeNode {
	var newTree *TreeNode
	var stack []*TreeNode
	currentNode := &TreeNode{}
	if len(leafs) > 0 {
		currentNode.Key = leafs[0]
		newTree = currentNode
		stack = append(stack, currentNode)
	}
	for i := 1; i < len(leafs); i++ {
		newNode := TreeNode{
			Key: leafs[i],
		}
		currentNode = stack[0]
		newNode.Parent = currentNode
		if currentNode.Left == nil {
			currentNode.Left = &newNode
			stack = append(stack, &newNode)
			continue
		}
		currentNode.Right = &newNode
		stack = append(stack, &newNode)
		stack = stack[1:]
	}
	return newTree
}

func GetNode(head *TreeNode, key int) *TreeNode {
	if head == nil || head.Key == key {
		return head
	}
	node := GetNode(head.Left, key)
	if node == nil {
		node = GetNode(head.Right, key)
	}
	return node
}

func GetNodes(head *TreeNode, keys ...int) []*TreeNode {
	var result []*TreeNode
	for _, v := range keys {
		result = append(result, GetNode(head, v))
	}
	return result
}
