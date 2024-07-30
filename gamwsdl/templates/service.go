package gamtemplates

var GamServiceTemplate = `// Code generated by gamwsdl/main.go. DO NOT EDIT.
package {{ generatePackageName .Definition.Service.Name }}

import (
	"context"
	"errors"
	"encoding/xml"
	"fmt"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/requestbody"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
	gamxml "github.com/knakazawa99/gam-sdk-go/pkg/xml"
)


const (
	wsdl = "{{ .Definition.Service.Port.Address.Location }}"
	{{ range .Definition.Binding.Operations	 }}
		{{ $operationName := .Name }}
			Operation{{ $operationName | toUpperCamelFromCamelCase }} = "{{ $operationName }}"
	{{ end }}
)

{{ $performFilterName := "" }}
{{/* Generate Interface */}}
{{ if ne .Definition.PortType nil }}
	// {{ .Definition.PortType.Name }} {{ if .Definition.PortType.Documentations }} {{ .Definition.PortType.Documentations | makeComment }}{{ end }}
	type  {{ .Definition.PortType.Name }} interface {
		{{ range .Definition.PortType.Operations }}
			{{ $outputName := .Output.Name }}
			{{ $operationType := .Name | getOperationTypeFromOperationName }}
			{{ if eq $operationType "Perform" }}
				// {{ .Name | toUpperCamelFromCamelCase }} {{ if .Documentations }} {{ .Documentations | makeComment }}{{ end }}
				{{ .Name | toUpperCamelFromCamelCase }}(ctx context.Context, action requestbody.PerformAction, input {{ .Name | toUpperCamelFromCamelCase}}) (*{{ $outputName | toUpperCamelFromCamelCase }}, error)
			{{ else }}
				// {{ .Name | toUpperCamelFromCamelCase }} {{ if .Documentations }} {{ .Documentations | makeComment }}{{ end }}
				{{ .Name | toUpperCamelFromCamelCase }}(ctx context.Context, input {{ .Name | toUpperCamelFromCamelCase}}) (*{{ $outputName | toUpperCamelFromCamelCase }}, error)
			{{ end }}
		{{ end }}
	}
{{ end }}


{{/* Generate Response */}}
type Envelope struct {
	xmlName xml.Name ` + "`" + `xml:"Envelope"` + "`" + `
	Body    Body     ` + "`" + `xml:"Body"` + "`" + `
}

{{ if ne .Definition.PortType nil }}
	type Body struct {
		xmlName xml.Name ` + "`" + `xml:"Body"` + "`" + `
		Fault soap.Fault ` + "`" + `xml:"Fault"` + "`" + `
		{{ range .Definition.PortType.Operations }}
			{{ $outputName := .Output.Name }}
			// {{ $outputName | toUpperCamelFromCamelCase }} {{ if .Documentations }} {{ .Documentations | makeComment }}{{ end }}
			{{ $outputName | toUpperCamelFromCamelCase }} {{ $outputName | toUpperCamelFromCamelCase }} ` + "`" + `xml:"{{ $outputName }}"` + "`" + `
		{{ end }}
	}
{{ end }}

{{/* Generate Service Struct */}}
{{ if ne .Definition.PortType nil }}
	type Service struct {
		client          soap.Client
		wsdl            string
		networkCode     int
		applicationName string
	}
	
	func NewService(client soap.Client, networkCode int, applicationName string) {{ .Definition.PortType.Name }} {
		return &Service{
			client:          client,
			wsdl:            wsdl,
			networkCode:     networkCode,
			applicationName: applicationName,
		}
	}

	// {{ .Definition.PortType.Name }} {{ if .Definition.PortType.Documentations }} {{ .Definition.PortType.Documentations | makeComment }}{{ end }}
	{{ range .Definition.PortType.Operations }}
		{{ $outputName := .Output.Name }}
		{{ $operationName := .Name }}
		{{ $operationType := .Name | getOperationTypeFromOperationName }}
		
		{{ if eq $operationType "Default" }}
			// {{ .Name | toUpperCamelFromCamelCase }} {{ if .Documentations }} {{ .Documentations | makeComment }}{{ end }}
			func (s *Service) {{ .Name | toUpperCamelFromCamelCase }}(ctx context.Context, input {{ .Name | toUpperCamelFromCamelCase}}) (*{{ $outputName | toUpperCamelFromCamelCase}}, error) {
				parameterXML, err := gamxml.DeepMarshal(input, true)
				data := requestbody.DefaultRequestBody{
					NetworkCode:     s.networkCode,
					ApplicationName: s.applicationName,
					OperationName:   Operation{{  .Name | toUpperCamelFromCamelCase }},
					Parameters:      string(parameterXML),
				}

				soapBody, err := requestbody.GenerateDefaultRequestSoapBody(data)
				if err != nil {
					return nil, fmt.Errorf("failed to generate soap body: %w", err)
				}
			
				body, err := s.client.Call(ctx, s.wsdl, soapBody)
				if err != nil {
					return nil, err
				}
			
				var envelope Envelope
				if err := xml.Unmarshal(body, &envelope); err != nil {
					return nil, err
				}
			
				if envelope.Body.Fault.Detail.APIExecutionFault != nil {
					apiExceptionFaultErrors := envelope.Body.Fault.Detail.APIExecutionFault.Errors[0]
					return nil, RaiseError(apiExceptionFaultErrors.ErrorType, apiExceptionFaultErrors.Reason)
				}
				return &envelope.Body.{{ $outputName | toUpperCamelFromCamelCase }}, nil	
			}
		{{ else if eq $operationType "Perform" }}
			func (s *Service) {{ .Name | toUpperCamelFromCamelCase }}(ctx context.Context, action requestbody.PerformAction, input {{ .Name | toUpperCamelFromCamelCase}}) (*{{ $outputName | toUpperCamelFromCamelCase }}, error) {
				performActionBody, err := action.GetPerformActionBody()
				if err != nil {
					return nil, fmt.Errorf("failed to perform action: %w", err)
				}
				
				{{ $statementName := getStatementName }}
				parameterXML, err := gamxml.DeepMarshal(input.{{ $statementName | toUpperCamelFromCamelCase }}, true)
				if err != nil {
					return nil, fmt.Errorf("failed to DeepMarshal: %w", err)
				}

				data := requestbody.PerformActionBodyData{
					NetworkCode:     s.networkCode,
					ApplicationName: s.applicationName,
					OperationName:   Operation{{  .Name | toUpperCamelFromCamelCase }},
					ActionName:      "{{ .Name | removePerform }}",
					ActionType:      action.GetActionType(),
					ActionField:     performActionBody,
					Statement: 	 "{{ $statementName }}",
					StatementParameter:   string(parameterXML),
				}
			
				soapBody, err := requestbody.GeneratePerformActionSoapBody(data)
				if err != nil {
					return nil, fmt.Errorf("failed to generate soap body: %w", err)
				}
			
				body, err := s.client.Call(ctx, s.wsdl, soapBody)
				if err != nil {
					return nil, fmt.Errorf("failed to call: %w", err)
				}
			
				var envelope Envelope
				if err := xml.Unmarshal(body, &envelope); err != nil {
					return nil, fmt.Errorf("failed to unmarshal: %w", err)
				}
			
				if envelope.Body.Fault.Detail.APIExecutionFault != nil {
					// TODO handle multiple errors
					apiExceptionFaultErrors := envelope.Body.Fault.Detail.APIExecutionFault.Errors[0]
					return nil, RaiseError(apiExceptionFaultErrors.ErrorType, apiExceptionFaultErrors.Reason)
				}
				return &envelope.Body.{{ $outputName | toUpperCamelFromCamelCase }}, nil
			}
		{{ end }}
	{{ end }}
{{ end }}
`