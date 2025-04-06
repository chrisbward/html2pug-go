package pkg

import (
	"regexp"
	"strings"

	"github.com/chrisbward/html2pug-go/pkg/html2pug-go/entities"
	"github.com/chrisbward/html2pug-go/pkg/html2pug-go/util"
	"golang.org/x/net/html"
)

type Writer struct {
	Options          *entities.Html2JadeConvertorOptions
	WrapLength       int
	Scalate          bool
	AttrSep          string
	AttrQuote        string
	NonAttrQuote     string
	AttrQuoteEscaped string
	NoEmptyPipe      bool
}

// NewWriter
func NewWriter(options *entities.Html2JadeConvertorOptions) (convertor entities.IWriter) {

	if options.WriterOptions == nil {
		options.WriterOptions = &entities.WriterOptions{}
	}
	wrapLength := 80
	if options.WriterOptions.WrapLength != nil {
		wrapLength = *options.WriterOptions.WrapLength
	}
	scalate := false
	if options.WriterOptions.Scalate != nil {
		scalate = *options.WriterOptions.Scalate
	}
	noAttrComma := false
	if options.WriterOptions.NoAttrComma != nil {
		noAttrComma = *options.WriterOptions.NoAttrComma
	}
	attrSep := ", "
	if scalate || noAttrComma {
		attrSep = " "
	}

	double := false
	if options.WriterOptions.Double != nil {
		double = *options.WriterOptions.Double
	}

	attrQuote := "'"
	nonAttrQuote := `"`
	if double {
		attrQuote = `"`
		nonAttrQuote = "'"
	}

	attrQuoteEscaped := "\\" + attrQuote

	noEmptyPipe := false
	if options.WriterOptions.NoEmptyPipe != nil {
		noEmptyPipe = *options.WriterOptions.NoEmptyPipe
	}

	convertor = &Writer{
		Options:          options,
		WrapLength:       wrapLength,
		Scalate:          scalate,
		AttrSep:          attrSep,
		AttrQuote:        attrQuote,
		NonAttrQuote:     nonAttrQuote,
		AttrQuoteEscaped: attrQuoteEscaped,
		NoEmptyPipe:      noEmptyPipe,
	}
	return
}

// BreakLine implements entities.IWriter.
func (w *Writer) BreakLine(line string) []string {
	// If line is empty, return an empty slice
	if len(line) == 0 {
		return []string{}
	}

	// If line has no spaces (a single word), return the line itself
	if !strings.Contains(line, " ") {
		return []string{line}
	}

	var lines []string
	var currentLine string
	words := strings.Fields(line) // splits the line into words based on whitespace

	for len(words) > 0 {
		word := words[0]
		words = words[1:]

		// If adding the word exceeds the wrapLength, push the current line and start a new one
		if len(currentLine)+len(word) > w.WrapLength {
			lines = append(lines, currentLine)
			currentLine = word
		} else if len(currentLine) > 0 {
			// Otherwise, append the word with a space
			currentLine += " " + word
		} else {
			// If currentLine is empty, just add the word
			currentLine = word
		}
	}

	// Push the remaining line if it's not empty
	if len(currentLine) > 0 {
		lines = append(lines, currentLine)
	}

	return lines
}

// BuildTagAttribute implements entities.IWriter.
func (w *Writer) BuildTagAttribute(attrName string, attrValue string) string {
	if !strings.Contains(attrValue, w.AttrQuote) {
		return attrName + "=" + w.AttrQuote + attrValue + w.AttrQuote
	} else if !strings.Contains(attrValue, w.NonAttrQuote) {
		return attrName + "=" + w.NonAttrQuote + attrValue + w.NonAttrQuote
	} else {
		escaped := strings.ReplaceAll(attrValue, w.AttrQuote, w.AttrQuoteEscaped)
		return attrName + "=" + w.AttrQuote + escaped + w.AttrQuote
	}
}

// ForEachChild implements entities.IWriter.
func (w *Writer) ForEachChild(parent *html.Node, cb func(child *html.Node)) {
	if parent == nil {
		return
	}

	for child := parent.FirstChild; child != nil; child = child.NextSibling {
		cb(child)
	}
}

// TagAttribute implements entities.IWriter.
func (w *Writer) TagAttribute(node *html.Node, indents string) string {
	if indents == "" {
		indents = ""
	}

	if node == nil || node.Type != html.ElementNode || len(node.Attr) == 0 {
		return ""
	}

	var result []string

	for _, attr := range node.Attr {
		attrName := attr.Key
		attrValue := attr.Val

		if attrName == "id" && util.IsValidJadeId(attrValue) {
			// Skip valid IDs (they're already used in TagHead)
			continue
		} else if attrName == "class" {
			classNames := strings.Fields(attrValue)
			var invalidClassNames []string
			for _, name := range classNames {
				if name != "" && !util.IsValidJadeClassName(name) {
					invalidClassNames = append(invalidClassNames, name)
				}
			}
			if len(invalidClassNames) > 0 {
				joined := strings.Join(invalidClassNames, " ")
				result = append(result, w.BuildTagAttribute(attrName, joined))
			}
		} else {
			// Replace newlines + optional whitespace with \n and the indent
			re := regexp.MustCompile(`(\r|\n)\s*`)
			escaped := re.ReplaceAllString(attrValue, `\$1`+indents)
			result = append(result, w.BuildTagAttribute(attrName, escaped))
		}
	}

	if len(result) > 0 {
		return "(" + strings.Join(result, w.AttrSep) + ")"
	}
	return ""
}

// TagHead implements entities.IWriter.
func (w *Writer) TagHead(node *html.Node) string {
	if node == nil || node.Type != html.ElementNode {
		return "div"
	}

	tagName := node.Data
	result := ""
	if strings.ToLower(tagName) != "div" {
		result = strings.ToLower(tagName)
	}

	var id string
	var classAttr string

	for _, attr := range node.Attr {
		switch attr.Key {
		case "id":
			id = attr.Val
		case "class":
			classAttr = attr.Val
		}
	}

	if id != "" && util.IsValidJadeId(id) {
		result += "#" + id
	}

	if classAttr != "" {
		classNames := strings.Fields(classAttr)
		validClassNames := []string{}
		for _, name := range classNames {
			if name != "" && util.IsValidJadeClassName(name) {
				validClassNames = append(validClassNames, name)
			}
		}
		if len(validClassNames) > 0 {
			result += "." + strings.Join(validClassNames, ".")
		}
	}

	if result == "" {
		result = "div"
	}

	return result
}

// TagText implements entities.IWriter.
func (w *Writer) TagText(node *html.Node) *string {
	if node == nil || node.FirstChild == nil {
		return nil
	}

	first := node.FirstChild
	if first.Type != html.TextNode {
		return nil
	}

	// Make sure it's the only child
	if first != node.LastChild {
		return nil
	}

	data := first.Data
	if len(data) > w.WrapLength || regexp.MustCompile(`\r|\n`).MatchString(data) {
		return nil
	}

	return &data
}

// WriteText implements entities.IWriter.
func (w *Writer) WriteText(node *html.Node, output *entities.IStringWriter, textOptions entities.TextOptions) {
	if node.Type == html.TextNode {
		data := node.Data
		if len(data) > 0 {
			re := regexp.MustCompile(`\r|\n`)
			lines := re.Split(data, -1)
			for _, line := range lines {
				w.WriteTextLine(node, line, output, textOptions)
			}
		}
	}
}

// WriteTextContent implements entities.IWriter.
func (w *Writer) WriteTextContent(node *html.Node, output *entities.IStringWriter, textOptions entities.TextOptions) {
	(*output).Enter()
	w.ForEachChild(node, func(child *html.Node) {
		w.WriteText(child, output, textOptions)
	})
	(*output).Leave()
}

// WriteTextLine implements entities.IWriter.
func (w *Writer) WriteTextLine(node *html.Node, line string, output *entities.IStringWriter, textOptions entities.TextOptions) {
	// Handle pipe and noEmptyPipe
	if textOptions.Pipe && w.NoEmptyPipe && len(strings.TrimSpace(line)) == 0 {
		return
	}

	// Prefix for the line
	prefix := ""
	if textOptions.Pipe {
		prefix = "| "
	}

	// Trim the line if needed
	// if the node is not nil, and previous sibling is not nil and the previous sibling type is an element
	if node != nil && node.PrevSibling != nil && node.PrevSibling.Type != html.ElementNode {
		line = strings.TrimLeft(line, " ")
	}

	// if the node is not nil, and next sibling is not nil and the previous sibling type is an element
	if node != nil && node.NextSibling != nil && node.NextSibling.Type != html.ElementNode {
		line = strings.TrimRight(line, " ")
	}

	// Handle non-empty line
	if len(line) > 0 {
		if textOptions.EncodeEntityRef {
			line = html.EscapeString(line)
		}

		if textOptions.EscapeBackslash {
			line = strings.ReplaceAll(line, "\\", "\\\\")
		}

		if !textOptions.Wrap || len(line) <= w.WrapLength {
			(*output).WriteLine(prefix+line, true)
		} else {
			// Split the line if it's too long
			lines := w.BreakLine(line)
			if len(lines) == 1 {
				(*output).WriteLine(prefix+line, true)
			} else {
				for _, l := range lines {
					w.WriteTextLine(node, l, output, textOptions)
				}
			}
		}
	}
}
