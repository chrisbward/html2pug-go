package pkg

import (
	"fmt"
	"regexp"
	"strings"

	util "github.com/chrisbward/html2pug-go/internal"
	"github.com/chrisbward/html2pug-go/pkg/html2pug-go/entities"
	html "golang.org/x/net/html"
)

type Convertor struct {
	Options              *entities.Html2JadeConvertorOptions
	PublicIdDocTypeNames map[string]string
	SystemIdDocTypeNames map[string]string
	Writer               *entities.IWriter
}

func NewConvertor(options *entities.Html2JadeConvertorOptions) (convertor entities.IConvertor) {

	convertor = &Convertor{
		Options: options,
		Writer:  options.Writer,
	}

	return

}

// Comment implements entities.IConvertor.
func (c *Convertor) Comment(node *html.Node, output *entities.IStringWriter) {
	// Match the condition pattern
	re := regexp.MustCompile(`\s*\[(if\s+[^\]]+)\]`)
	condition := re.FindStringSubmatch(node.Data)

	if condition == nil {
		// If no condition, handle as normal comment
		data := node.Data
		if len(data) == 0 || !strings.ContainsAny(data, "\r\n") {
			// Output single-line comment
			(*output).WriteLine("// "+strings.TrimSpace(data), true)
		} else {
			// Multi-line comment handling
			(*output).WriteLine("//", true)
			(*output).Enter()
			lines := strings.Split(data, "\r\n")
			for _, line := range lines {
				(*c.Writer).WriteTextLine(node, line, output, entities.TextOptions{
					Pipe: false,
					Trim: true,
					Wrap: false,
				})
			}
			(*output).Leave()
		}
	} else {
		// If condition exists, process it
		c.Conditional(node, condition[1], output)
	}

}

// Conditional implements entities.IConvertor.
func (c *Convertor) Conditional(node *html.Node, condition string, output *entities.IStringWriter) {
	// Get the inner HTML (textContent in JS) and clean it up
	innerHTML := strings.TrimSpace(node.Data)
	innerHTML = strings.Replace(innerHTML, " [if "+condition+"]> ", "", -1)
	innerHTML = strings.Replace(innerHTML, "<![endif]", "", -1)

	// Check if the inner HTML starts with a comment
	if strings.HasPrefix(innerHTML, "<!") {
		condition = " [" + condition + "] <!"
		innerHTML = "" // Set innerHTML to null if it starts with "<!"
	}

	// Create the new conditional element
	conditionalElem := &html.Node{
		Type:       html.ElementNode,
		Data:       "conditional", // Tag name 'conditional'
		Attr:       []html.Attribute{{Key: "condition", Val: condition}},
		FirstChild: nil, // Will be populated below if needed
	}

	// If there is inner HTML, we need to set it as children (text nodes)
	if innerHTML != "" {
		conditionalElem.FirstChild = &html.Node{
			Type: html.TextNode,
			Data: innerHTML,
		}
	}

	// Insert the new conditional element after the current node
	node.Parent.AppendChild(conditionalElem)
}

// Script implements entities.IConvertor.
func (c *Convertor) Script(node *html.Node, output *entities.IStringWriter, tagHead, tagAttr string) {
	// Check if scalate flag is set (equivalent to this.scalate in JavaScript)
	if c.Options.Scalate {
		// If scalate is true, output ':javascript'
		(*output).WriteLine(":javascript", true)

		// Write the text content of the script node
		(*c.Writer).WriteTextContent(node, output, entities.TextOptions{
			Pipe: false,
			Wrap: false,
		})
	} else {
		// If scalate is false, output the tag header and attributes followed by a period
		(*output).WriteLine(tagHead+tagAttr+".", true)

		// Write the text content of the script node with additional options
		(*c.Writer).WriteTextContent(node, output, entities.TextOptions{
			Pipe:            false,
			Trim:            true,
			Wrap:            false,
			EscapeBackslash: true,
		})
	}
}

// Style implements entities.IConvertor.
func (c *Convertor) Style(node *html.Node, output *entities.IStringWriter, tagHead string, tagAttr string) {
	if c.Options.Scalate {
		// In scalate mode, emit shorthand for embedded CSS
		(*output).WriteLine(":css", true)
		(*c.Writer).WriteTextContent(node, output, entities.TextOptions{
			Pipe: false,
			Wrap: false,
		})
	} else {
		// Otherwise, output full tag and its content
		(*output).WriteLine(tagHead+tagAttr+".", true)
		(*c.Writer).WriteTextContent(node, output, entities.TextOptions{
			Pipe: false,
			Trim: true,
			Wrap: false,
		})
	}
}

// Text implements entities.IConvertor.
func (c *Convertor) Text(node *html.Node, output *entities.IStringWriter, textOptions entities.TextOptions) {
	util.NormalizeTextNode(node)
	(*c.Writer).WriteText(node, output, textOptions)
}

func (c *Convertor) Document(document *entities.Document, output *entities.IStringWriter) {
	var docTypeName string
	docType := document.GetDocType()
	// Traverse to find the DoctypeNode

	if docType != nil {
		publicId := docType.PublicId
		systemId := docType.SystemId

		if publicId != "" {
			if val, ok := c.PublicIdDocTypeNames[publicId]; ok {
				docTypeName = val
			}
		} else if systemId != "" {
			if val, ok := c.SystemIdDocTypeNames[systemId]; ok {
				docTypeName = val
			}
		} else if docType.Name != "" && strings.ToLower(docType.Name) == "html" {
			docTypeName = "html"
		}

		if docTypeName != "" {
			(*output).WriteLine(fmt.Sprintf("doctype %s", docTypeName), entities.DoIndent)
		}
	}

	if document.DocumentElement != nil {
		c.Children(document.Root, output, false)
	} else {
		// Simulating getElementsByTagName('html') logic
		htmlEls := document.GetElementsByTagName("html")
		if len(htmlEls) > 0 {
			c.Element(htmlEls[0], output, false)
		}
	}

}

// Placeholder methods for Children and Element
func (c *Convertor) Children(parent *html.Node, output *entities.IStringWriter, indent bool) {

	if indent {
		(*output).Enter()
	}

	for child := parent.FirstChild; child != nil; child = child.NextSibling {
		switch child.Type {
		case html.ElementNode:
			c.Element(child, output, false)
		case html.TextNode:
			// Check if the parent is a <code> element
			if strings.ToLower(parent.Data) == "code" {
				c.Text(child, output, entities.TextOptions{
					EncodeEntityRef: true,
					Pipe:            true,
				})
			} else {
				c.Text(child, output, entities.TextOptions{
					EncodeEntityRef: true, // set to false if you want doNotEncode behavior
				})
			}
		case html.CommentNode:
			c.Comment(child, output)
		}
	}

	if indent {
		(*output).Leave()
	}
}

func (c *Convertor) Element(node *html.Node, output *entities.IStringWriter, doNotEncode bool) {
	if node == nil || node.Type != html.ElementNode {
		return
	}

	tagName := strings.ToLower(node.Data)
	tagHead := (*c.Writer).TagHead(node)
	tagAttr := (*c.Writer).TagAttribute(node, (*output).GetIndents())
	tagText := (*c.Writer).TagText(node)

	switch tagName {
	case "script", "style":
		if util.HasAttr(node, "src") {
			(*output).WriteLine(tagHead+tagAttr, true)
			(*c.Writer).WriteTextContent(node, output, entities.TextOptions{
				Pipe: false,
				Wrap: false,
			})
		} else if tagName == "script" {
			c.Script(node, output, tagHead, tagAttr)
		} else if tagName == "style" {
			c.Style(node, output, tagHead, tagAttr)
		}
	case "conditional":
		cond := util.GetAttr(node, "condition")
		(*output).WriteLine("//"+cond, true)
		c.Children(node, output, true)
	case "pre":
		(*output).WriteLine(tagHead+tagAttr+".", true)
		(*output).Enter()
		firstLine := true
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			if child.Type == html.TextNode {
				data := child.Data
				if strings.TrimSpace(data) != "" {
					if firstLine {
						data = strings.TrimLeft(data, "\r\n")
						data = "\\n" + data
						firstLine = false
					}
					data = strings.ReplaceAll(data, "\t", "\\t")
					data = strings.ReplaceAll(data, "\r\n", "\n"+(*output).GetIndents())
					data = strings.ReplaceAll(data, "\r", "\n"+(*output).GetIndents())
					data = strings.ReplaceAll(data, "\n", "\n"+(*output).GetIndents())
					(*output).Write(data, true)
				}
			}
		}
		(*output).WriteLine("", true)
		(*output).Leave()
	default:
		if c.Options.Bodyless && (tagName == "html" || tagName == "body") {
			// bodyless in options, skip the output and jump to next
			c.Children(node, output, false)
		} else if !c.Options.KeepHead && (tagName == "head") {
			// headless in options, skip the children of head
			// c.Children(node, output, false)
		} else if tagText != nil {
			if doNotEncode {
				(*output).WriteLine(tagHead+tagAttr+" "+*tagText, true)
			} else {
				(*output).WriteLine(tagHead+tagAttr+" "+html.EscapeString(*tagText), true)
			}
		} else {
			(*output).WriteLine(tagHead+tagAttr, true)
			c.Children(node, output, true)
		}
	}

	// (*output).WriteLine(fmt.Sprintf("Processing element: %v", el.Type), entities.DoIndent)
}

func GetElementsByTagName(doc *entities.Document, tag string) []*entities.Element {
	// Simulated lookup
	if doc.DocumentElement != nil && strings.ToLower(doc.DocumentElement.TagName) == tag {
		return []*entities.Element{doc.DocumentElement}
	}
	return []*entities.Element{}
}
