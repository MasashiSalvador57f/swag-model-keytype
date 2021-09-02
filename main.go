package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/ghodss/yaml"
)

var (
	swaggerFileOpt = flag.String("f", "./swagger.yaml", "path to swagger file")
	outputFileOpt  = flag.String("o", "./KeyTypes.ts", "output file name of key types")
)

const keyTypeTemplate = `
{{- $typeName := .TypeName }}
const {{ $typeName }} = {
	{{ range $i := .Properties -}}
	{{ .Ref }}:"{{ .KeyName -}}",
	{{ end -}}
} as const;
export type {{ $typeName }} = typeof {{ $typeName }}[keyof typeof {{ $typeName -}}]
`

// Property is definition model property.
type Property struct {
	KeyName string
	Ref     string
}

// Model represens model with  Title & Properties for swager definition.
type Model struct {
	TypeName   string
	Properties []Property
}

func main() {
	flag.Parse()
	swaggerFilePath := *swaggerFileOpt
	outputFilePath := *outputFileOpt
	input, err := ioutil.ReadFile(swaggerFilePath)
	if err != nil {
		panic(err)
	}
	outFile, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}

	var docYAML openapi2.T
	if err = yaml.Unmarshal(input, &docYAML); err != nil {
		panic(err)
	}

	models := []Model{}
	for i := range docYAML.Definitions {
		definition := docYAML.Definitions[i]
		title := definition.Value.Title
		if len(title) == 0 {
			continue
		}
		m := Model{
			TypeName:   fmt.Sprintf("%sKey", title),
			Properties: []Property{},
		}

		for k := range definition.Value.Properties {
			keyStr := fmt.Sprint(k)
			p := Property{
				Ref:     strings.Title(keyStr),
				KeyName: keyStr,
			}
			m.Properties = append(m.Properties, p)
		}
		models = append(models, m)
	}

	fmt.Fprint(outFile, "// auto generated file DO NOT EDIT.")
	tpl, err := template.New("keyTypeTempalte").Parse(keyTypeTemplate)
	for i := range models {
		if tpl.Execute(outFile, models[i]); err != nil {
			panic(err)
		}
	}

	fmt.Println("finished")
}
