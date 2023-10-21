package gox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {
	t.Run(
		"find node with modifier type", func(t *testing.T) {
			nodes := []Node{
				Text("test"),
				Div(Text("div children")),
				Attribute(),
			}
			assert.Equal(t, nodeModifier, findNodeWithType(nodeModifier, nodes...).nodeType)
		},
	)
	t.Run(
		"remove nodes with modifier type", func(t *testing.T) {
			nodes := []Node{
				Text("test"),
				Div(Text("div children")),
				Attribute(),
			}
			expectedNodes := []Node{
				Text("test"),
				Div(Text("div children")),
			}
			nodes = removeNodesWithType(nodeModifier, nodes...)
			assert.Equal(t, expectedNodes, nodes)
		},
	)
	t.Run(
		"minify html", func(t *testing.T) {
			html := `
				<html lang="en">
					<head>
						<title>Example page</title>
						<meta name="viewport" content="width=device-width,initial-scale=1.0" />
					</head>
					<body>
						<h1>Example page</h1>
						<p>Example paragraph</p>
					</body>
				</html>
				`
			expectedHtml := `<html lang="en"><head><title>Example page</title><meta name="viewport" content="width=device-width,initial-scale=1.0" /></head><body><h1>Example page</h1><p>Example paragraph</p></body></html>`
			assert.Equal(t, expectedHtml, Minify(html))
		},
	)
}
