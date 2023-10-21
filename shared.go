package gox

const (
	sharedCite  = "cite"
	sharedData  = "data"
	sharedForm  = "form"
	sharedLabel = "label"
	sharedSlot  = "slot"
	sharedSpan  = "span"
	sharedStyle = "style"
	sharedTitle = "title"
)

func Cite(nodes ...Node) Node {
	return createShared(sharedCite, nodeElement, nodes...)
}

func Data(nodes ...Node) Node {
	return createShared(sharedData, nodeElement, nodes...)
}

func Form(nodes ...Node) Node {
	return createShared(sharedForm, nodeElement, nodes...)
}

func Label(nodes ...Node) Node {
	return createShared(sharedLabel, nodeElement, nodes...)
}

func Slot(nodes ...Node) Node {
	return createShared(sharedSlot, nodeElement, nodes...)
}

func Span(nodes ...Node) Node {
	return createShared(sharedSpan, nodeElement, nodes...)
}

func Style(nodes ...Node) Node {
	return createShared(sharedStyle, nodeAttribute, nodes...)
}

func Title(nodes ...Node) Node {
	return createShared(sharedTitle, nodeElement, nodes...)
}
