// Code generated by gamwsdl/main.go. DO NOT EDIT.
package inventory

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/requestbody"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
	gamxml "github.com/knakazawa99/gam-sdk-go/pkg/xml"
)

const (
	wsdl = "https://ads.google.com/apis/ads/publisher/v202405/InventoryService"

	OperationCreateAdUnits = "createAdUnits"

	OperationGetAdUnitSizesByStatement = "getAdUnitSizesByStatement"

	OperationGetAdUnitsByStatement = "getAdUnitsByStatement"

	OperationPerformAdUnitAction = "performAdUnitAction"

	OperationUpdateAdUnits = "updateAdUnits"
)

// InventoryServiceInterface
//
// Provides operations for creating, updating and retrieving {@link AdUnit} objects.
//
// <p>Line items connect a creative with its associated ad unit through targeting.
//
// <p>An ad unit represents a piece of inventory within a publisher. It contains all the settings
// that need to be associated with the inventory in order to serve ads. For example, the ad unit
// contains creative size restrictions and AdSense settings.
type InventoryServiceInterface interface {

	// CreateAdUnits
	//
	// Creates new {@link AdUnit} objects.
	//
	CreateAdUnits(ctx context.Context, input CreateAdUnits) (*CreateAdUnitsResponse, error)

	// GetAdUnitSizesByStatement
	//
	// Returns a set of all relevant {@link AdUnitSize} objects.
	//
	// <p>The given {@link Statement} is currently ignored but may be honored in future versions.
	//
	GetAdUnitSizesByStatement(ctx context.Context, input GetAdUnitSizesByStatement) (*GetAdUnitSizesByStatementResponse, error)

	// GetAdUnitsByStatement
	//
	// Gets a {@link AdUnitPage} of {@link AdUnit} objects that satisfy the given {@link
	// Statement#query}. The following fields are supported for filtering:
	//
	// <table>
	// <tr>
	// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
	// </tr>
	// <tr>
	// <td>{@code adUnitCode}</td>
	// <td>{@link AdUnit#adUnitCode}</td>
	// </tr>
	// <tr>
	// <td>{@code id}</td>
	// <td>{@link AdUnit#id}</td>
	// </tr>
	// <tr>
	// <td>{@code name}</td>
	// <td>{@link AdUnit#name}</td>
	// </tr>
	// <tr>
	// <td>{@code parentId}</td>
	// <td>{@link AdUnit#parentId}</td>
	// </tr>
	// <tr>
	// <td>{@code status}</td>
	// <td>{@link AdUnit#status}</td>
	// </tr>
	// <tr>
	// <td>{@code lastModifiedDateTime}</td>
	// <td>{@link AdUnit#lastModifiedDateTime}</td>
	// </tr>
	// </table>
	//
	GetAdUnitsByStatement(ctx context.Context, input GetAdUnitsByStatement) (*GetAdUnitsByStatementResponse, error)

	// PerformAdUnitAction
	//
	// Performs actions on {@link AdUnit} objects that match the given {@link Statement#query}.
	//
	PerformAdUnitAction(ctx context.Context, action requestbody.PerformAction, input PerformAdUnitAction) (*PerformAdUnitActionResponse, error)

	// UpdateAdUnits
	//
	// Updates the specified {@link AdUnit} objects.
	//
	UpdateAdUnits(ctx context.Context, input UpdateAdUnits) (*UpdateAdUnitsResponse, error)
}

type Envelope struct {
	xmlName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	xmlName xml.Name   `xml:"Body"`
	Fault   soap.Fault `xml:"Fault"`

	// CreateAdUnitsResponse
	//
	// Creates new {@link AdUnit} objects.
	//
	CreateAdUnitsResponse CreateAdUnitsResponse `xml:"createAdUnitsResponse"`

	// GetAdUnitSizesByStatementResponse
	//
	// Returns a set of all relevant {@link AdUnitSize} objects.
	//
	// <p>The given {@link Statement} is currently ignored but may be honored in future versions.
	//
	GetAdUnitSizesByStatementResponse GetAdUnitSizesByStatementResponse `xml:"getAdUnitSizesByStatementResponse"`

	// GetAdUnitsByStatementResponse
	//
	// Gets a {@link AdUnitPage} of {@link AdUnit} objects that satisfy the given {@link
	// Statement#query}. The following fields are supported for filtering:
	//
	// <table>
	// <tr>
	// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
	// </tr>
	// <tr>
	// <td>{@code adUnitCode}</td>
	// <td>{@link AdUnit#adUnitCode}</td>
	// </tr>
	// <tr>
	// <td>{@code id}</td>
	// <td>{@link AdUnit#id}</td>
	// </tr>
	// <tr>
	// <td>{@code name}</td>
	// <td>{@link AdUnit#name}</td>
	// </tr>
	// <tr>
	// <td>{@code parentId}</td>
	// <td>{@link AdUnit#parentId}</td>
	// </tr>
	// <tr>
	// <td>{@code status}</td>
	// <td>{@link AdUnit#status}</td>
	// </tr>
	// <tr>
	// <td>{@code lastModifiedDateTime}</td>
	// <td>{@link AdUnit#lastModifiedDateTime}</td>
	// </tr>
	// </table>
	//
	GetAdUnitsByStatementResponse GetAdUnitsByStatementResponse `xml:"getAdUnitsByStatementResponse"`

	// PerformAdUnitActionResponse
	//
	// Performs actions on {@link AdUnit} objects that match the given {@link Statement#query}.
	//
	PerformAdUnitActionResponse PerformAdUnitActionResponse `xml:"performAdUnitActionResponse"`

	// UpdateAdUnitsResponse
	//
	// Updates the specified {@link AdUnit} objects.
	//
	UpdateAdUnitsResponse UpdateAdUnitsResponse `xml:"updateAdUnitsResponse"`
}

type Service struct {
	client          soap.Client
	wsdl            string
	networkCode     int
	applicationName string
}

func NewService(client soap.Client, networkCode int, applicationName string) InventoryServiceInterface {
	return &Service{
		client:          client,
		wsdl:            wsdl,
		networkCode:     networkCode,
		applicationName: applicationName,
	}
}

// InventoryServiceInterface
//
// Provides operations for creating, updating and retrieving {@link AdUnit} objects.
//
// <p>Line items connect a creative with its associated ad unit through targeting.
//
// <p>An ad unit represents a piece of inventory within a publisher. It contains all the settings
// that need to be associated with the inventory in order to serve ads. For example, the ad unit
// contains creative size restrictions and AdSense settings.
//

// CreateAdUnits
//
// Creates new {@link AdUnit} objects.
func (s *Service) CreateAdUnits(ctx context.Context, input CreateAdUnits) (*CreateAdUnitsResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationCreateAdUnits,
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
	return &envelope.Body.CreateAdUnitsResponse, nil
}

// GetAdUnitSizesByStatement
//
// Returns a set of all relevant {@link AdUnitSize} objects.
//
// <p>The given {@link Statement} is currently ignored but may be honored in future versions.
func (s *Service) GetAdUnitSizesByStatement(ctx context.Context, input GetAdUnitSizesByStatement) (*GetAdUnitSizesByStatementResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationGetAdUnitSizesByStatement,
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
	return &envelope.Body.GetAdUnitSizesByStatementResponse, nil
}

// GetAdUnitsByStatement
//
// Gets a {@link AdUnitPage} of {@link AdUnit} objects that satisfy the given {@link
// Statement#query}. The following fields are supported for filtering:
//
// <table>
// <tr>
// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
// </tr>
// <tr>
// <td>{@code adUnitCode}</td>
// <td>{@link AdUnit#adUnitCode}</td>
// </tr>
// <tr>
// <td>{@code id}</td>
// <td>{@link AdUnit#id}</td>
// </tr>
// <tr>
// <td>{@code name}</td>
// <td>{@link AdUnit#name}</td>
// </tr>
// <tr>
// <td>{@code parentId}</td>
// <td>{@link AdUnit#parentId}</td>
// </tr>
// <tr>
// <td>{@code status}</td>
// <td>{@link AdUnit#status}</td>
// </tr>
// <tr>
// <td>{@code lastModifiedDateTime}</td>
// <td>{@link AdUnit#lastModifiedDateTime}</td>
// </tr>
// </table>
func (s *Service) GetAdUnitsByStatement(ctx context.Context, input GetAdUnitsByStatement) (*GetAdUnitsByStatementResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationGetAdUnitsByStatement,
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
	return &envelope.Body.GetAdUnitsByStatementResponse, nil
}

func (s *Service) PerformAdUnitAction(ctx context.Context, action requestbody.PerformAction, input PerformAdUnitAction) (*PerformAdUnitActionResponse, error) {
	performActionBody, err := action.GetPerformActionBody()
	if err != nil {
		return nil, fmt.Errorf("failed to perform action: %w", err)
	}

	parameterXML, err := gamxml.DeepMarshal(input.FilterStatement, true)
	if err != nil {
		return nil, fmt.Errorf("failed to DeepMarshal: %w", err)
	}

	data := requestbody.PerformActionBodyData{
		NetworkCode:        s.networkCode,
		ApplicationName:    s.applicationName,
		OperationName:      OperationPerformAdUnitAction,
		ActionName:         "adUnitAction",
		ActionType:         action.GetActionType(),
		ActionField:        performActionBody,
		Statement:          "filterStatement",
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
	return &envelope.Body.PerformAdUnitActionResponse, nil
}

// UpdateAdUnits
//
// Updates the specified {@link AdUnit} objects.
func (s *Service) UpdateAdUnits(ctx context.Context, input UpdateAdUnits) (*UpdateAdUnitsResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationUpdateAdUnits,
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
	return &envelope.Body.UpdateAdUnitsResponse, nil
}
