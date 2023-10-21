package gox

func Range[T any](s []T, fn func(item T) Node) Node {
	children := make([]node, len(s))
	for i, val := range s {
		item := fn(val)
		n, ok := item.(node)
		if !ok {
			continue
		}
		children[i] = n
	}
	return node{
		nodeType: nodeFragment,
		children: children,
	}
}

func If(condition bool, nodes ...Node) Node {
	if !condition {
		return node{
			nodeType: nodeFragment,
		}
	}
	children := make([]node, len(nodes))
	for i, item := range nodes {
		n, ok := item.(node)
		if !ok {
			continue
		}
		children[i] = n
	}
	return node{
		nodeType: nodeFragment,
		children: children,
	}
}
