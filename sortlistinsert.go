package student

type NodeI struct {
	Data int
	Next *NodeI
}

//SortListInsert a function  that inserts data_ref in the linked list l while keeping the list sorted in ascending order.
func SortListInsert(l *NodeI, data_ref int) *NodeI {
	cur := l
	newNode := &NodeI{Data: data_ref}
	if l == nil {
		cur.Data = data_ref
		return cur
	}
	for cur.Next != nil {
		if cur.Data >= data_ref {
			newNode.Next = cur
			return newNode
		} else if cur.Next.Data >= data_ref {
			newNode.Next = cur.Next
			cur.Next = newNode
			return l
		}
		cur = cur.Next
	}
	cur.Next = newNode
	return l
}
