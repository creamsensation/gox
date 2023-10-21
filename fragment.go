package gox

func Fragment(nodes ...Node) Node {
	_, children := processNodes(nodes)
	return node{
		nodeType: nodeFragment,
		children: children,
	}
}
