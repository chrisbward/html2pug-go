package pkg

import (
	"strings"

	"github.com/chrisbward/html2pug-go/pkg/html2pug-go/entities"
)

type Html2PugConvertor struct {
	Options *entities.Html2JadeConvertorOptions
}

func NewHtml2PugConvertor(options *entities.Html2JadeConvertorOptions) (html2jadeConvertor entities.IHtml2JadeConvertor) {

	if options == nil {
		options = &entities.Html2JadeConvertorOptions{}
	}
	applyOptions(options)

	if options.Parser == nil {
		parser := NewParser(options)
		options.Parser = &parser
	}

	if options.Parser == nil {
		output := NewStringOutput(options)
		options.Output = (*entities.IStringWriter)(&output)
	}

	if options.Writer == nil {
		writer := NewWriter(options)
		options.Writer = &writer
	}
	if options.Converter == nil {
		convertor := NewConvertor(options)
		options.Converter = &convertor
	}

	html2jadeConvertor = &Html2PugConvertor{
		Options: options,
	}

	return
}

// ConvertHTML
func (h2jc *Html2PugConvertor) ConvertHTML(html string, callback entities.Html2JadeConvertorConvertDocumentCallback) {

	htmlReader := strings.NewReader(html)
	(*h2jc.Options.Parser).Parse(htmlReader, func(err []error, window entities.Window) {
		if len(err) > 0 {
			return
		}
		stringOutput := NewStringOutput(h2jc.Options).(entities.IStringWriter)

		(*h2jc.Options.Converter).Document(window.Document, &stringOutput)

		callback(nil, stringOutput.Final())

	})

}

func applyOptions(options *entities.Html2JadeConvertorOptions) {
}
