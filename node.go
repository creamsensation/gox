package gox

import (
	"fmt"
)

type Node interface {
	Node() Node
}

type node struct {
	nodeType   int
	name       string
	value      any
	attributes []node
	children   []node
}

const (
	nodeElement = iota
	nodeAttribute
	nodeText
	nodeFragment
	nodeRaw
	nodeModifier
)

func (n node) Node() Node {
	return n
}

func createElement(name string, nodes ...Node) Node {
	attributes, children := processNodes(nodes)
	return node{
		nodeType:   nodeElement,
		name:       name,
		attributes: attributes,
		children:   children,
	}
}

func createAttribute[T any](name string, value ...T) Node {
	n := node{
		nodeType: nodeAttribute,
		name:     name,
	}
	if len(value) > 0 {
		n.value = value[0]
	}
	return n
}

func createShared(name string, nodeType int, nodes ...Node) Node {
	if len(nodes) > 0 {
		modifier := findNodeWithType(nodeModifier, nodes...)
		if modifier.nodeType == nodeModifier && modifier.value != nil {
			nodeType = modifier.value.(int)
			nodes = removeNodesWithType(nodeModifier, nodes...)
		}
	}
	if nodeType == nodeAttribute {
		if len(nodes) == 0 {
			return createAttribute[any](name)
		}
		n, ok := nodes[0].(node)
		if !ok {
			return createAttribute[any](name)
		}
		return createAttribute(name, fmt.Sprintf("%v", n.value))
	}
	return createElement(name, nodes...)
}
