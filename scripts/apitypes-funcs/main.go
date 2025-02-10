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
	PackageVersion string

	AdditionalImports []string

	Types []templateDataT
}

type templateDataT struct {
	// Type is the name of the type.
	Type string

	// KonnectStatusType is the name of the konnect status type (.status.konnect).
	// If it's not provided Konnect status functions will not be generated.
	KonnectStatusType string

	// ControlPlaneRefType is the ControlPlaneRef type to be used in the template (with the package name if it's outside
	// the type's package).
	ControlPlaneRefType string

	// ControlPlaneRefRequired is true if the ControlPlaneRef is required for the type.
	ControlPlaneRefRequired bool
}

const (
	apiPackageName             = "api"
	configurationPackageName   = "configuration"
	konnectPackageName         = "konnect"
	gatewayOperatorPackageName = "gateway-operator"
)

func main() {
	if err := renderTemplate(konnectFuncTemplate, konnectFuncOutputFileName, supportedKonnectTypesControlPlaneConfig, configurationPackageName); err != nil {
		panic(err)
	}
	if err := renderTemplate(konnectFuncStandaloneTemplate, konnectFuncOutputFileName, supportedKonnectTypesStandalone, konnectPackageName); err != nil {
		panic(err)
	}
	if err := renderTemplate(listFuncTemplate, listFuncOutputFileNamme, supportedKonnectPackageTypesWithList, konnectPackageName); err != nil {
		panic(err)
	}
	if err := renderTemplate(listFuncTemplate, listFuncOutputFileNamme, supportedConfigurationPackageTypesWithList, configurationPackageName); err != nil {
		panic(err)
	}
	if err := renderTemplate(listFuncTemplate, listFuncOutputFileNamme, supportedGatewayOperatorPackageTypesWithList, gatewayOperatorPackageName); err != nil {
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
		path := filepath.Join(apiPackageName, packagename, st.PackageVersion, outputFile)
		if err := tpl.Execute(contents, st); err != nil {
			return fmt.Errorf("%s: failed to execute template for %s: %w", path, outputFile, err)
		}
		if err := os.WriteFile(path, contents.Bytes(), 0o600); err != nil {
			return fmt.Errorf("%s: failed to write file %s: %w", path, outputFile, err)
		}
	}
	return nil
}
