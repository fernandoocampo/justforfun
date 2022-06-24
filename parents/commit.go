package parents

func GetCommonParent(cone, ctwo *TreeNode) *TreeNode {
	if cone == nil || ctwo == nil {
		return nil
	}
	done := getDepth(cone)
	dtwo := getDepth(ctwo)

	for done > dtwo {
		cone = cone.Parent
		done--
	}

	for dtwo > done {
		ctwo = ctwo.Parent
		dtwo--
	}

	for cone != ctwo && done > 0 {
		cone = cone.Parent
		ctwo = ctwo.Parent
		done--
	}

	return cone
}

func getDepth(node *TreeNode) int {
	var result int
	for {
		node = node.Parent
		result++
		if node.Parent == nil {
			break
		}
	}
	return result
}
