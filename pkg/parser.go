package pkg

import (
	"io"

	"github.com/chrisbward/html2pug-go/pkg/entities"
	html "golang.org/x/net/html"
)

type Parser struct {
	Options *entities.Html2JadeConvertorOptions
}

// NewParser
func NewParser(options *entities.Html2JadeConvertorOptions) (convertor entities.IParser) {
	convertor = &Parser{
		Options: options,
	}
	return
}

func (p *Parser) Parse(htmlContentReader io.Reader, callback entities.ParserCallback) {
	window := entities.Window{}

	var errors []error
	doc, err := html.Parse(htmlContentReader)
	if err != nil {
		errors = append(errors, err)
	}

	window.Document = &entities.Document{
		Root: doc,
	}

	callback(errors, window)
}
