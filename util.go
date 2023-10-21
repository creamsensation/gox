package gox

import "strings"

func Minify(html string) string {
	replacer := strings.NewReplacer("\t", "", "\n", "")
	return replacer.Replace(html)
}

func findNodeWithType(nodeType int, nodes ...Node) node {
	for _, item := range nodes {
		n, ok := item.(node)
		if !ok {
			continue
		}
		if n.nodeType == nodeType {
			return n
		}
	}
	return node{}
}

func removeNodesWithType(nodeType int, nodes ...Node) []Node {
	result := make([]Node, 0)
	for _, item := range nodes {
		n, ok := item.(node)
		if !ok {
			result = append(result, item)
			continue
		}
		if n.nodeType == nodeType {
			continue
		}
		result = append(result, item)
	}
	return result
}
