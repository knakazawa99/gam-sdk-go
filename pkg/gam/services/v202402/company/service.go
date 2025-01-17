// Code generated by gamwsdl/main.go. DO NOT EDIT.
package company

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/requestbody"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
	gamxml "github.com/knakazawa99/gam-sdk-go/pkg/xml"
)

const (
	wsdl = "https://ads.google.com/apis/ads/publisher/v202402/CompanyService"

	OperationCreateCompanies = "createCompanies"

	OperationGetCompaniesByStatement = "getCompaniesByStatement"

	OperationPerformCompanyAction = "performCompanyAction"

	OperationUpdateCompanies = "updateCompanies"
)

// CompanyServiceInterface
//
// Provides operations for creating, updating and retrieving {@link Company} objects.
type CompanyServiceInterface interface {

	// CreateCompanies
	//
	// Creates new {@link Company} objects.
	//
	CreateCompanies(ctx context.Context, input CreateCompanies) (*CreateCompaniesResponse, error)

	// GetCompaniesByStatement
	//
	// Gets a {@link CompanyPage} of {@link Company} objects that satisfy the given {@link
	// Statement#query}. The following fields are supported for filtering:
	//
	// <table>
	// <tr>
	// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
	// </tr>
	// <tr>
	// <td>{@code id}</td>
	// <td>{@link Company#id}</td>
	// </tr>
	// <tr>
	// <td>{@code name}</td>
	// <td>{@link Company#name}</td>
	// </tr>
	// <tr>
	// <td>{@code type}</td>
	// <td>{@link Company#type}</td>
	// </tr>
	// <tr>
	// <td>{@code lastModifiedDateTime}</td>
	// <td>{@link Company#lastModifiedDateTime}</td>
	// </tr>
	// </table>
	//
	GetCompaniesByStatement(ctx context.Context, input GetCompaniesByStatement) (*GetCompaniesByStatementResponse, error)

	// PerformCompanyAction
	//
	// Performs actions on {@link Company} objects that match the given {@code Statement}.
	//
	PerformCompanyAction(ctx context.Context, action requestbody.PerformAction, input PerformCompanyAction) (*PerformCompanyActionResponse, error)

	// UpdateCompanies
	//
	// Updates the specified {@link Company} objects.
	//
	UpdateCompanies(ctx context.Context, input UpdateCompanies) (*UpdateCompaniesResponse, error)
}

type Envelope struct {
	xmlName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	xmlName xml.Name   `xml:"Body"`
	Fault   soap.Fault `xml:"Fault"`

	// CreateCompaniesResponse
	//
	// Creates new {@link Company} objects.
	//
	CreateCompaniesResponse CreateCompaniesResponse `xml:"createCompaniesResponse"`

	// GetCompaniesByStatementResponse
	//
	// Gets a {@link CompanyPage} of {@link Company} objects that satisfy the given {@link
	// Statement#query}. The following fields are supported for filtering:
	//
	// <table>
	// <tr>
	// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
	// </tr>
	// <tr>
	// <td>{@code id}</td>
	// <td>{@link Company#id}</td>
	// </tr>
	// <tr>
	// <td>{@code name}</td>
	// <td>{@link Company#name}</td>
	// </tr>
	// <tr>
	// <td>{@code type}</td>
	// <td>{@link Company#type}</td>
	// </tr>
	// <tr>
	// <td>{@code lastModifiedDateTime}</td>
	// <td>{@link Company#lastModifiedDateTime}</td>
	// </tr>
	// </table>
	//
	GetCompaniesByStatementResponse GetCompaniesByStatementResponse `xml:"getCompaniesByStatementResponse"`

	// PerformCompanyActionResponse
	//
	// Performs actions on {@link Company} objects that match the given {@code Statement}.
	//
	PerformCompanyActionResponse PerformCompanyActionResponse `xml:"performCompanyActionResponse"`

	// UpdateCompaniesResponse
	//
	// Updates the specified {@link Company} objects.
	//
	UpdateCompaniesResponse UpdateCompaniesResponse `xml:"updateCompaniesResponse"`
}

type Service struct {
	client          soap.Client
	wsdl            string
	networkCode     int
	applicationName string
}

func NewService(client soap.Client, networkCode int, applicationName string) CompanyServiceInterface {
	return &Service{
		client:          client,
		wsdl:            wsdl,
		networkCode:     networkCode,
		applicationName: applicationName,
	}
}

// CompanyServiceInterface
//
// Provides operations for creating, updating and retrieving {@link Company} objects.
//

// CreateCompanies
//
// Creates new {@link Company} objects.
func (s *Service) CreateCompanies(ctx context.Context, input CreateCompanies) (*CreateCompaniesResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationCreateCompanies,
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
	return &envelope.Body.CreateCompaniesResponse, nil
}

// GetCompaniesByStatement
//
// Gets a {@link CompanyPage} of {@link Company} objects that satisfy the given {@link
// Statement#query}. The following fields are supported for filtering:
//
// <table>
// <tr>
// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
// </tr>
// <tr>
// <td>{@code id}</td>
// <td>{@link Company#id}</td>
// </tr>
// <tr>
// <td>{@code name}</td>
// <td>{@link Company#name}</td>
// </tr>
// <tr>
// <td>{@code type}</td>
// <td>{@link Company#type}</td>
// </tr>
// <tr>
// <td>{@code lastModifiedDateTime}</td>
// <td>{@link Company#lastModifiedDateTime}</td>
// </tr>
// </table>
func (s *Service) GetCompaniesByStatement(ctx context.Context, input GetCompaniesByStatement) (*GetCompaniesByStatementResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationGetCompaniesByStatement,
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
	return &envelope.Body.GetCompaniesByStatementResponse, nil
}

func (s *Service) PerformCompanyAction(ctx context.Context, action requestbody.PerformAction, input PerformCompanyAction) (*PerformCompanyActionResponse, error) {
	performActionBody, err := action.GetPerformActionBody()
	if err != nil {
		return nil, fmt.Errorf("failed to perform action: %w", err)
	}

	parameterXML, err := gamxml.DeepMarshal(input.Statement, true)
	if err != nil {
		return nil, fmt.Errorf("failed to DeepMarshal: %w", err)
	}

	data := requestbody.PerformActionBodyData{
		NetworkCode:        s.networkCode,
		ApplicationName:    s.applicationName,
		OperationName:      OperationPerformCompanyAction,
		ActionName:         "companyAction",
		ActionType:         action.GetActionType(),
		ActionField:        performActionBody,
		Statement:          "statement",
		StatementParameter: string(parameterXML),
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
	return &envelope.Body.PerformCompanyActionResponse, nil
}

// UpdateCompanies
//
// Updates the specified {@link Company} objects.
func (s *Service) UpdateCompanies(ctx context.Context, input UpdateCompanies) (*UpdateCompaniesResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationUpdateCompanies,
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
	return &envelope.Body.UpdateCompaniesResponse, nil
}
