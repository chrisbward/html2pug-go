package util

import (
	"regexp"
	"strings"

	html "golang.org/x/net/html"
)

var (
	validJadeIdRegExp    = regexp.MustCompile(`^[\w\-]+$`)
	validJadeClassRegExp = regexp.MustCompile(`^[\w\-]+$`)
)

func IsValidJadeId(id string) (isValid bool) {
	id = strings.TrimSpace(id)
	return id != "" && validJadeIdRegExp.MatchString(id)
}
func IsValidJadeClassName(className string) (isValid bool) {
	className = strings.TrimSpace(className)
	return className != "" && validJadeClassRegExp.MatchString(className)
}

func NormalizeTextNode(parent *html.Node) {
	var prev *html.Node

	// Traverse through all children of the parent node
	for curr := parent.FirstChild; curr != nil; {
		next := curr.NextSibling // Save the next sibling to avoid losing track during iteration

		// Check if both the current node and the previous node are text nodes
		if curr.Type == html.TextNode && prev != nil && prev.Type == html.TextNode {

			// Merge the current text node's data with the previous text node's data
			prev.Data += curr.Data

			// Remove the current text node as it has been merged into the previous one
			parent.RemoveChild(curr)
		} else {

			// Update the 'prev' node to be the current node if no merge occurs
			prev = curr
		}

		// Move to the next sibling for the next iteration
		curr = next
	}
}

func HasAttr(node *html.Node, name string) bool {
	for _, attr := range node.Attr {
		if attr.Key == name {
			return true
		}
	}
	return false
}

func GetAttr(node *html.Node, name string) string {
	for _, attr := range node.Attr {
		if attr.Key == name {
			return attr.Val
		}
	}
	return ""
}
