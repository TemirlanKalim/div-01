package student

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	cur := l.Head
	var prev *NodeL
	for cur != nil {
		if l.Head.Data == data_ref {
			l.Head = cur.Next
			cur = l.Head
		} else if cur.Data == data_ref {
			prev.Next = cur.Next
			if cur.Next == nil || cur.Next.Data == data_ref {
				cur = cur.Next
			} else {
				prev = cur.Next
				cur = cur.Next
			}
		} else {
			prev = cur
			cur = cur.Next
		}
	}
}

func ListPushBack(l *List, data interface{}) *List {
	node := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = node
		return l
	}
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
	return l
}
