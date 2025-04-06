package pkg

import (
	"github.com/chrisbward/html2pug-go/pkg/html2pug-go/entities"
)

type Output struct {
	Options *entities.Html2JadeConvertorOptions
	Indents string
}

// NewOutput
func NewOutput(options *entities.Html2JadeConvertorOptions) (output entities.IOutput) {

	output = &Output{
		Options: options,
		Indents: "",
	}
	return
}

// Final implements entities.IOutput.
func (o *Output) Final() (output string) {
	panic("unimplemented 1")
}

// Write implements entities.IOutput.
func (o *Output) Write(data string, indent bool) {
	panic("unimplemented 2")

	if indent {
		print(o.Indents)
	}
	print(data)
}

// WriteLine implements entities.IOutput.
func (o *Output) WriteLine(data string, indent bool) {
	panic("unimplemented 3")

	if indent {
		print(o.Indents)
	}
	println(data)
}

func (o *Output) GetIndents() (indents string) {
	panic("unimplemented 4")
	return o.Indents
}

// Enter
func (o *Output) Enter() {

	if o.Options.UseTabs {
		o.Indents += "\t"
	} else {
		for i := 1; ; {
			if o.Options.NSpaces >= 1 {
				if i > o.Options.NSpaces {
					break
				}
				o.Indents += " "
				i++
			} else {
				if i < o.Options.NSpaces {
					break
				}
				o.Indents += " "
				i--
			}
		}
	}
}

// Leave implements entities.IOutput.
func (o *Output) Leave() {
	if o.Options.UseTabs {
		if len(o.Indents) > 0 {
			o.Indents = o.Indents[1:]
		}
	} else {
		if len(o.Indents) >= o.Options.NSpaces {
			o.Indents = o.Indents[o.Options.NSpaces:]
		}
	}
}
