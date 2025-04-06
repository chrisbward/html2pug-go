package entities

type Window struct {
	Document *Document
}

type Doctype struct {
	Name     string
	PublicId string
	SystemId string
}

type Element struct {
	TagName string
}

type ProgramInputType string

const (
	HTMLProgramInputType ProgramInputType = "html"
	URLProgramInputType  ProgramInputType = "url"
)

type WriterOptions struct {
	WrapLength  *int
	Scalate     *bool
	NoAttrComma *bool
	Double      *bool
	NoEmptyPipe *bool
}

type Html2JadeConvertorOptions struct {
	UseTabs          bool
	NSpaces          int
	KeepHead         bool
	Bodyless         bool
	Scalate          bool
	WriterOptions    *WriterOptions
	InputType        ProgramInputType
	OutDirectoryPath string

	Parser    *IParser
	Converter *IConvertor
	Output    *IStringWriter
	Writer    *IWriter
}
type Html2JadeConvertorConvertDocumentCallback func(err error, jadeOutput string)

type ParserCallback func(err []error, window Window)

var DoIndent = true

type TextOptions struct {
	EncodeEntityRef bool
	Pipe            bool
	Trim            bool
	Wrap            bool
	EscapeBackslash bool
}
