package gamwsdl

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	gamtemplates "github.com/knakazawa99/gam-sdk-go/gamwsdl/templates"

	"github.com/pkg/errors"
)

type TemplateOutput struct {
	Definition     WSDLDefinition
	SourceFilePath string
}

type GenerateManagement struct {
	Version  string       `json:"version"`
	Services []GAMService `json:"services"`
}

type GAMService struct {
	Name                string    `json:"name"`
	Version             string    `json:"version"`
	URL                 string    `json:"url"`
	GenerateDestination string    `json:"generateDestination"`
	LastUpdatedAt       time.Time `json:"lastUpdatedAt"`
}

// GenerateGoCodeFromGAMWSDL generates go code from GAM WSDL
func GenerateGoCodeFromGAMWSDL(url, outputFilePath string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "failed to get from url")
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			err = errors.Wrap(err, "failed to close")
		}
	}()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return errors.Wrap(err, "failed to copy")
	}

	var gamWSDLDefinition WSDLDefinition
	if err := xml.Unmarshal(buf.Bytes(), &gamWSDLDefinition); err != nil {
		return errors.Wrap(err, "failed to unmarshal xml")
	}

	fileTemplateMap := []map[string]string{
		{"fileName": "type", "template": gamtemplates.GamTypeTemplate},
		{"fileName": "service", "template": gamtemplates.GamServiceTemplate},
		{"fileName": "error", "template": gamtemplates.GamErrorTemplate},
	}
	for _, fileName := range fileTemplateMap {
		if err := generateCodeFromWSDLDefinition(gamWSDLDefinition, outputFilePath, fileName); err != nil {
			return errors.Wrap(err, "failed to generate code")
		}
	}

	return nil
}

func generateCodeFromWSDLDefinition(gamWSDLDefinition WSDLDefinition, outputFilePath string, fileNameTemplateMap map[string]string) error {
	statementName := "statement"
	for _, element := range gamWSDLDefinition.Types.Schema.Elements {
		if element.ComplexType == nil || element.ComplexType.Sequence == nil {
			continue
		}
		if !strings.HasPrefix(element.Name, "perform") {
			continue
		}
		for _, sequenceElement := range element.ComplexType.Sequence.Elements {
			if sequenceElement.Name == "filterStatement" {
				statementName = "filterStatement"
			}
		}

	}
	tmpl, err := template.New("").Funcs(
		template.FuncMap{
			"generatePackageName":               generatePackageName,
			"ptrStrToStr":                       func(s *string) string { return *s },
			"makeComment":                       makeComment,
			"toUpperCamelFromUpperSnakeCase":    toUpperCamelFromUpperSnakeCase,
			"toUpperCamelFromCamelCase":         toUpperCamelFromCamelCase,
			"removeDotFromString":               removeDotFromString,
			"getTypeKind":                       getTypeKind,
			"getTypeName":                       getTypeName,
			"getTypeFromElementType":            getTypeFromElementType,
			"isElementTypeError":                isElementTypeError,
			"getGoBuiltinType":                  getGoBuiltinType,
			"getOperationTypeFromOperationName": getOperationTypeFromOperationName,
			"getStatementName":                  func() string { return statementName },
			"isNameIncludeError": func(s string) bool {
				if strings.Contains(s, "Error") {
					return true
				}
				if strings.Contains(s, "Fault") {
					return true
				}
				if strings.Contains(s, "Exception") {
					return true
				}
				return false
			},
			"isPerformActionBase": func(s string) bool {
				if strings.HasSuffix(s, "Action") {
					return true
				}
				return false
			},
			"removePerform": func(s string) string {
				return toCamelCaseFromUpperCamelCase(strings.ReplaceAll(s, "perform", ""))
			},
		},
	).Parse(fileNameTemplateMap["template"])

	if err != nil {
		return errors.Wrap(err, "failed to parse template")
	}

	templateOutput := TemplateOutput{
		Definition: gamWSDLDefinition,
	}

	var newFileContent bytes.Buffer
	if err := tmpl.Execute(&newFileContent, templateOutput); err != nil {
		return errors.Wrap(err, "failed to execute template")
	}

	if _, err := os.Stat(outputFilePath); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(outputFilePath, 0755); err != nil {
				return errors.Wrap(err, "failed to make directory")
			}
		}
	}

	goFileName := fmt.Sprintf("%s/%s.go", outputFilePath, fileNameTemplateMap["fileName"])
	if _, err := os.Stat(goFileName); err == nil {
		return errors.Wrap(err, "failed to stat file")
	}
	if err := os.WriteFile(goFileName, newFileContent.Bytes(), 0644); err != nil {
		return errors.Wrap(err, "failed to write file")
	}
	return nil
}
