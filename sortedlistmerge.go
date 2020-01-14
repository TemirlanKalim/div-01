package student

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil && n2 != nil {
		return n2
	} else if n1 != nil && n2 == nil {
		return n1
	} else if n1 == nil && n2 == nil {
		return nil
	}
	cur1 := n1
	cur2 := n2
	var temp *NodeI
	var main *NodeI
	for cur1 != nil && cur2 != nil {
		if cur1.Data > cur2.Data {
			if temp == nil {
				temp = cur2
				main = temp
			} else {
				temp.Next = cur2
				temp = temp.Next
			}
			cur2 = cur2.Next
		} else {
			if temp == nil {
				temp = cur1
				main = temp
			} else {
				temp.Next = cur1
				temp = temp.Next
			}
			cur1 = cur1.Next
		}
	}
	if cur1 != nil {
		//for cur1 != nil {
		temp.Next = cur1
		//temp = temp.Next
		//	cur1 = cur1.Next
		//}
	} else if cur2 != nil {
		//for cur2 != nil {
		temp.Next = cur2
		//temp = temp.Next
		//	cur2 = cur2.Next
		//}
	}
	return main
}
