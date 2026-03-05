package p1650

import "localhost/leetcode/util"

func lowestCommonAncestor(p *util.Node, q *util.Node) *util.Node {
	ppath := make(map[*util.Node]struct{})

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
