package main

import (
	"bytes"
	"embed"
	"io/ioutil"
	"log"
	"text/template"
)

var (
	baseTemplateName = "sqlx-codegen/template/base.tmpl"
	//go:embed templates/base.tmpl
	baseTmplFile embed.FS

	initTemplates = map[string]embed.FS{
		baseTemplateName: baseTmplFile,
	}

	defaultTemplates = make(map[string]*template.Template)
)

func init() {
	for k, v := range initTemplates {
		tbs, err := v.ReadFile("templates/base.tmpl")
		if err != nil {
			log.Fatal(err)
		}
		t, err := template.New(k).ParseGlob(string(tbs))
		if err != nil {
			log.Fatal(err)
		}
		defaultTemplates[k] = t
	}
}

type SqlxCodeGenTemplate struct {
	Command  string
	Package  string
	TypeList []string
	Tmpl     *template.Template
	OutFile  string
}

func (s *SqlxCodeGenTemplate) Render2File() error {
	var result bytes.Buffer
	err := s.Tmpl.Execute(&result, s)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.OutFile, result.Bytes(), 0644)
}

func NewTemplate(fileName string) (*template.Template, error) {
	t, err := template.New(fileName).ParseFiles(fileName)
	if err != nil {
		return nil, err
	}
	return t, nil
}
