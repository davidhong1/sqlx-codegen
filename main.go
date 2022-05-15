package main

import (
	"flag"
	"fmt"
	"go/build"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

const (
	DEFAULT_DIR          = "."
	DEFAULT_GEN_FILENAME = "sqlx_codegen.go"
	COMMAND              = "sqlx-codegen"
)

func main() {
	var typeNames, templateFile, packageName string

	flag.StringVar(&typeNames,
		"t", "", "comma-separated list of type names; must be set")
	flag.StringVar(&templateFile,
		"template", "", "template file in . directory or in sqlx-codegen/template directory; default is sqlx-codegen/template/base.tmpl")
	flag.StringVar(&packageName,
		"package", "", "packageName; default use package in ./xxx.go file")
	flag.Parse()
	if typeNames == "" {
		flag.Usage()
		os.Exit(1)
	}

	var err error
	if packageName == "" {
		packageName, err = parsePackageName(DEFAULT_DIR)
		if err != nil {
			log.Fatal(err)
		}
	}

	tmpl := defaultTemplates[baseTemplateName]
	if templateFile != "" {
		tmpl, err = NewTemplate(templateFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	sqlxCodeGenTemplate := SqlxCodeGenTemplate{
		Command:  COMMAND,
		Package:  packageName,
		TypeList: strings.Split(typeNames, ","),
		Tmpl:     tmpl,
		OutFile:  DEFAULT_GEN_FILENAME,
	}
	err = sqlxCodeGenTemplate.Render2File()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("generate code finished.")
}

func parsePackageName(dir string) (string, error) {
	pkg, err := build.Default.ImportDir(dir, 0)
	if err != nil {
		return "", fmt.Errorf("importing dir %s: %v", dir, err)
	}
	var names []string
	names = append(names, pkg.GoFiles...)

	fset := token.NewFileSet()
	for _, name := range names {
		if !strings.HasSuffix(name, ".go") {
			continue
		}
		f, err := parser.ParseFile(fset, name, nil, 0)
		if err != nil {
			return "", fmt.Errorf("parsing file %s: %v", name, err)
		}
		return f.Name.Name, nil
	}

	return "", fmt.Errorf("not found go file in %s directory", dir)
}
