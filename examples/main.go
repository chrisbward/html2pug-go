package main

import (
	"fmt"

	html2puggo "github.com/chrisbward/html2pug-go/pkg/html2pug-go"
	html2puggo_entities "github.com/chrisbward/html2pug-go/pkg/html2pug-go/entities"
)

func main() {

	targetHtml := ``

	pugConvertor := html2puggo.NewHtml2PugConvertor(&html2puggo_entities.Html2JadeConvertorOptions{})
	pugConvertor.ConvertHTML(targetHtml, func(err error, jadeOutput string) {
		fmt.Println(jadeOutput)
	})
}
