package pkg

import "github.com/chrisbward/html2pug-go/pkg/entities"

type StreamOutput struct {
	Output
}

func NewStreamOutput() (convertor entities.IStreamOutput) {

	convertor = &StreamOutput{}

	return

}
func (so *StreamOutput) GetIndents() (indents string) {
	return so.Indents
}
func (so *StreamOutput) WriteLine(data string, indent bool) {

}
func (so *StreamOutput) Write(data string, indent bool) {

}
func (so *StreamOutput) Final() (output string) {

	return
}
