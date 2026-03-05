package solutions

func lowestCommonAncestor(p *Node, q *Node) *Node {
	ppath := make(map[*Node]struct{})

	for p != nil {
		ppath[p] = struct{}{}
		p = p.Parent
	}

	for q != nil {
		if _, ok := ppath[q]; ok {
			return q
		}
		q = q.Parent
	}
	return nil
}
