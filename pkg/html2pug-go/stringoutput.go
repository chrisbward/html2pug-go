package pkg

import (
	"strings"

	"github.com/chrisbward/html2pug-go/pkg/html2pug-go/entities"
)

type StringOutput struct {
	Output
	Fragments []string
}

func NewStringOutput(options *entities.Html2JadeConvertorOptions) (stringOutput entities.IStringOutput) {

	stringOutput = &StringOutput{
		Output: Output{
			Options: options,
		},
		Fragments: make([]string, 0),
	}

	return

}
func (so *StringOutput) GetIndents() (indents string) {
	return so.Indents
}
func (so *StringOutput) Write(data string, indent bool) {

	if indent {
		so.Fragments = append(so.Fragments, so.Indents+data)
	} else {
		so.Fragments = append(so.Fragments, data)
	}
}
func (so *StringOutput) WriteLine(data string, indent bool) {

	if strings.Trim(data, " ") == "" {
		return
	}

	if indent {
		so.Fragments = append(so.Fragments, so.Indents+data+"\n")
	} else {
		so.Fragments = append(so.Fragments, data+"\n")
	}
}

func (so *StringOutput) Final() (output string) {
	output = strings.Join(so.Fragments, "")
	so.Fragments = []string{}
	return
}
