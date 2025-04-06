package pkg

import (
	"strings"

	"github.com/chrisbward/html2pug-go/pkg/entities"
)

type Html2JadeConvertor struct {
	Parser    *entities.IParser
	Converter *entities.IConvertor
	Output    *entities.IStringWriter
	Writer    *entities.IWriter
}

func NewHtml2JadeConvertor(output *entities.IStringWriter, convertor *entities.IConvertor, writer *entities.IWriter, parser *entities.IParser) (html2jadeConvertor entities.IHtml2JadeConvertor) {

	html2jadeConvertor = &Html2JadeConvertor{
		Output:    output,
		Converter: convertor,
		Parser:    parser,
		Writer:    writer,
	}

	return
}

// ConvertHTML
func (h2jc *Html2JadeConvertor) ConvertHTML(html string, options *entities.Html2JadeConvertorOptions, callback *entities.Html2JadeConvertorConvertDocumentCallback) {

	if options == nil {
		options = &entities.Html2JadeConvertorOptions{}
	}
	applyOptions(options)
	if options.Parser == nil {
		options.Parser = NewParser(options)
	}
	htmlReader := strings.NewReader(html)
	options.Parser.Parse(htmlReader, func(err []error, window entities.Window) {
		if len(err) > 0 {
			return
		}
		stringOutput := NewStringOutput(options).(entities.IStringWriter)
		if options.Converter == nil {
			options.Converter = NewConvertor(options, h2jc.Writer)
		}
		options.Converter.Document(window.Document, &stringOutput)

		if callback != nil {
			(*callback)(nil, stringOutput.Final())
		}

	})

}

func applyOptions(options *entities.Html2JadeConvertorOptions) {
}
