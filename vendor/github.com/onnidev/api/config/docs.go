package config

import (
	"io/ioutil"

	"github.com/cescoferraro/docgen"
	"github.com/go-chi/chi"
)

// Docs TODO: NEEDS COMMENT INFO
func Docs(r chi.Router) {
	docs := (docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
		ProjectPath: "github.com/onnidev/api",
		Intro:       "A Golang API to support ONNi",
	}))
	ioutil.WriteFile("README.md", []byte(docs), 0644)

}
