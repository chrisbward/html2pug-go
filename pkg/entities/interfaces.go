package entities

import (
	"io"

	"golang.org/x/net/html"
)

type IHtml2JadeConvertor interface {
	ConvertHTML(html string, options *Html2JadeConvertorOptions, callback *Html2JadeConvertorConvertDocumentCallback)
}

type IStringWriter interface {
	Write(data string, indent bool)
	WriteLine(data string, indent bool)
	GetIndents() (indents string)
	Final() (output string)
	Enter()
	Leave()
}
type IOutput interface {
	IStringWriter
}
type IStreamOutput interface {
	IStringWriter
}
type IStringOutput interface {
	IStringWriter
}
type IConvertor interface {
	Document(document *Document, output *IStringWriter)
	Element(node *html.Node, output *IStringWriter, doNotEncode bool)
	Children(doc *html.Node, output *IStringWriter, flag bool)
	Text(*html.Node, *IStringWriter, TextOptions)
	Comment(*html.Node, *IStringWriter)
	Conditional(node *html.Node, condition string, output *IStringWriter)
	Script(*html.Node, *IStringWriter, string, string)
	Style(*html.Node, *IStringWriter, string, string)
}

type IParser interface {
	Parse(html io.Reader, callback ParserCallback)
}

type IWriter interface {
	TagHead(*html.Node) string
	TagAttribute(*html.Node, string) string
	BuildTagAttribute(string, string) string
	TagText(node *html.Node) *string
	ForEachChild(parent *html.Node, cb func(child *html.Node))
	WriteTextContent(*html.Node, *IStringWriter, TextOptions)
	WriteText(*html.Node, *IStringWriter, TextOptions)
	WriteTextLine(*html.Node, string, *IStringWriter, TextOptions)
	BreakLine(string) []string
}
