package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig"
)

type supportedTypesT struct {
	Package string

	Types []templateDataT
}

type templateDataT struct {
	// Type is the name of the type.
	Type string

	// KonnectStatusType is the name of the konnect status type (.status.konnect).
	KonnectStatusType string

	// HasKonnectStatus is true if the type has a Konnect status field.
	HasKonnectStatus bool
}

const (
	apiPackageName           = "api"
	configurationPackageName = "configuration"
	konnectPackageName       = "konnect"
)

func main() {
	if err := renderTemplate(konnectFuncTemplate, konnectFuncOutputFileName, supportedTypesControlPlaneConfig, configurationPackageName); err != nil {
		panic(err)
	}
	if err := renderTemplate(konnectFuncStandaloneTemplate, konnectFuncOutputFileName, supportedTypesStandalone, konnectPackageName); err != nil {
		panic(err)
	}
}

func renderTemplate(
	templateContent string,
	outputFile string,
	supportedTypes []supportedTypesT,
	packagename string,
) error {
	tpl, err := template.New("tpl").Funcs(sprig.TxtFuncMap()).Parse(templateContent)
	if err != nil {
		return fmt.Errorf("failed to parse template for %s: %w", outputFile, err)
	}
	for _, st := range supportedTypes {
		contents := &bytes.Buffer{}
		path := filepath.Join(apiPackageName, packagename, st.Package, outputFile)
		if err := tpl.Execute(contents, st); err != nil {
			return fmt.Errorf("%s: failed to execute template for %s: %w", path, outputFile, err)
		}
		if err := os.WriteFile(path, contents.Bytes(), 0o600); err != nil {
			return fmt.Errorf("%s: failed to write file %s: %w", path, outputFile, err)
		}
	}
	return nil
}
