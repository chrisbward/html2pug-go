package entities

import "golang.org/x/net/html"

type Document struct {
	Doctype         *Doctype
	DocumentElement *Element
	Root            *html.Node
}

func (d *Document) GetDocType() (docType *Doctype) {
	for n := d.Root; n != nil; n = n.NextSibling {
		if n.Type == html.DoctypeNode {
			d.Doctype = &Doctype{}
		}
	}
	return d.Doctype
}
func (d *Document) GetElementsByTagName(tagName string) (nodes []*html.Node) {
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tagName {
			nodes = append(nodes, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	if d.Root != nil {
		traverse(d.Root)
	}
	return
}
