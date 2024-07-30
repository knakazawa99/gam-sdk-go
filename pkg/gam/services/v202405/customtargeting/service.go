// Code generated by gamwsdl/main.go. DO NOT EDIT.
package customtargeting

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/requestbody"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
	gamxml "github.com/knakazawa99/gam-sdk-go/pkg/xml"
)

const (
	wsdl = "https://ads.google.com/apis/ads/publisher/v202405/CustomTargetingService"

	OperationCreateCustomTargetingKeys = "createCustomTargetingKeys"

	OperationCreateCustomTargetingValues = "createCustomTargetingValues"

	OperationGetCustomTargetingKeysByStatement = "getCustomTargetingKeysByStatement"

	OperationGetCustomTargetingValuesByStatement = "getCustomTargetingValuesByStatement"

	OperationPerformCustomTargetingKeyAction = "performCustomTargetingKeyAction"

	OperationPerformCustomTargetingValueAction = "performCustomTargetingValueAction"

	OperationUpdateCustomTargetingKeys = "updateCustomTargetingKeys"

	OperationUpdateCustomTargetingValues = "updateCustomTargetingValues"
)

// CustomTargetingServiceInterface
//
// Provides operations for creating, updating and retrieving {@link CustomTargetingKey} and {@link
// CustomTargetingValue} objects.
type CustomTargetingServiceInterface interface {

	// CreateCustomTargetingKeys
	//
	// Creates new {@link CustomTargetingKey} objects.
	//
	// <p>The following fields are required:
	//
	// <ul>
	// <li>{@link CustomTargetingKey#name}
	// <li>{@link CustomTargetingKey#type}
	// </ul>
	//
	CreateCustomTargetingKeys(ctx context.Context, input CreateCustomTargetingKeys) (*CreateCustomTargetingKeysResponse, error)

	// CreateCustomTargetingValues
	//
	// Creates new {@link CustomTargetingValue} objects.
	//
	// <p>The following fields are required:
	//
	// <ul>
	// <li>{@link CustomTargetingValue#customTargetingKeyId}
	// <li>{@link CustomTargetingValue#name}
	// </ul>
	//
	CreateCustomTargetingValues(ctx context.Context, input CreateCustomTargetingValues) (*CreateCustomTargetingValuesResponse, error)

	// GetCustomTargetingKeysByStatement
	//
	// Gets a {@link CustomTargetingKeyPage} of {@link CustomTargetingKey} objects that satisfy the
	// given {@link Statement#query}. The following fields are supported for filtering:
	//
	// <table>
	// <tr>
	// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
	// </tr>
	// <tr>
	// <tr>
	// <td>{@code id}</td>
	// <td>{@link CustomTargetingKey#id}</td>
	// </tr>
	// <tr>
	// <td>{@code name}</td>
	// <td>{@link CustomTargetingKey#name}</td>
	// </tr>
	// <tr>
	// <td>{@code displayName}</td>
	// <td>{@link CustomTargetingKey#displayName}</td>
	// </tr>
	// <tr>
	// <td>{@code type}</td>
	// <td>{@link CustomTargetingKey#type}</td>
	// </tr>
	// </table>
	//
	GetCustomTargetingKeysByStatement(ctx context.Context, input GetCustomTargetingKeysByStatement) (*GetCustomTargetingKeysByStatementResponse, error)

	// GetCustomTargetingValuesByStatement
	//
	// Gets a {@link CustomTargetingValuePage} of {@link CustomTargetingValue} objects that satisfy
	// the given {@link Statement#query}.
	//
	// <p>The {@code WHERE} clause in the {@link Statement#query} must always contain {@link
	// CustomTargetingValue#customTargetingKeyId} as one of its columns in a way that it is AND'ed
	// with the rest of the query. So, if you want to retrieve values for a known set of key ids,
	// valid {@link Statement#query} would look like:
	//
	// <ol>
	// <li>"WHERE customTargetingKeyId IN ('17','18','19')" retrieves all values that are associated
	// with keys having ids 17, 18, 19.
	// <li>"WHERE customTargetingKeyId = '17' AND name = 'red'" retrieves values that are associated
	// with keys having id 17 and value name is 'red'.
	// </ol>
	//
	// <p>The following fields are supported for filtering:
	//
	// <table>
	// <tr>
	// <th scope="col">PQL Property</th>
	// <th scope="col">Object Property</th>
	// </tr>
	// <tr>
	// <td>{@code id}</td>
	// <td>{@link CustomTargetingValue#id}</td>
	// </tr>
	// <tr>
	// <td>{@code customTargetingKeyId}</td>
	// <td>{@link CustomTargetingValue#customTargetingKeyId}</td>
	// </tr>
	// <tr>
	// <td>{@code name}</td>
	// <td>{@link CustomTargetingValue#name}</td>
	// </tr>
	// <tr>
	// <td>{@code displayName}</td>
	// <td>{@link CustomTargetingValue#displayName}</td>
	// </tr>
	// <tr>
	// <td>{@code matchType}</td>
	// <td>{@link CustomTargetingValue#matchType}</td>
	// </tr>
	// </table>
	//
	GetCustomTargetingValuesByStatement(ctx context.Context, input GetCustomTargetingValuesByStatement) (*GetCustomTargetingValuesByStatementResponse, error)

	// PerformCustomTargetingKeyAction
	//
	// Performs actions on {@link CustomTargetingKey} objects that match the given {@link
	// Statement#query}.
	//
	PerformCustomTargetingKeyAction(ctx context.Context, action requestbody.PerformAction, input PerformCustomTargetingKeyAction) (*PerformCustomTargetingKeyActionResponse, error)

	// PerformCustomTargetingValueAction
	//
	// Performs actions on {@link CustomTargetingValue} objects that match the given {@link
	// Statement#query}.
	//
	PerformCustomTargetingValueAction(ctx context.Context, action requestbody.PerformAction, input PerformCustomTargetingValueAction) (*PerformCustomTargetingValueActionResponse, error)

	// UpdateCustomTargetingKeys
	//
	// Updates the specified {@link CustomTargetingKey} objects.
	//
	UpdateCustomTargetingKeys(ctx context.Context, input UpdateCustomTargetingKeys) (*UpdateCustomTargetingKeysResponse, error)

	// UpdateCustomTargetingValues
	//
	// Updates the specified {@link CustomTargetingValue} objects.
	//
	UpdateCustomTargetingValues(ctx context.Context, input UpdateCustomTargetingValues) (*UpdateCustomTargetingValuesResponse, error)
}

type Envelope struct {
	xmlName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	xmlName xml.Name   `xml:"Body"`
	Fault   soap.Fault `xml:"Fault"`

	// CreateCustomTargetingKeysResponse
	//
	// Creates new {@link CustomTargetingKey} objects.
	//
	// <p>The following fields are required:
	//
	// <ul>
	// <li>{@link CustomTargetingKey#name}
	// <li>{@link CustomTargetingKey#type}
	// </ul>
	//
	CreateCustomTargetingKeysResponse CreateCustomTargetingKeysResponse `xml:"createCustomTargetingKeysResponse"`

	// CreateCustomTargetingValuesResponse
	//
	// Creates new {@link CustomTargetingValue} objects.
	//
	// <p>The following fields are required:
	//
	// <ul>
	// <li>{@link CustomTargetingValue#customTargetingKeyId}
	// <li>{@link CustomTargetingValue#name}
	// </ul>
	//
	CreateCustomTargetingValuesResponse CreateCustomTargetingValuesResponse `xml:"createCustomTargetingValuesResponse"`

	// GetCustomTargetingKeysByStatementResponse
	//
	// Gets a {@link CustomTargetingKeyPage} of {@link CustomTargetingKey} objects that satisfy the
	// given {@link Statement#query}. The following fields are supported for filtering:
	//
	// <table>
	// <tr>
	// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
	// </tr>
	// <tr>
	// <tr>
	// <td>{@code id}</td>
	// <td>{@link CustomTargetingKey#id}</td>
	// </tr>
	// <tr>
	// <td>{@code name}</td>
	// <td>{@link CustomTargetingKey#name}</td>
	// </tr>
	// <tr>
	// <td>{@code displayName}</td>
	// <td>{@link CustomTargetingKey#displayName}</td>
	// </tr>
	// <tr>
	// <td>{@code type}</td>
	// <td>{@link CustomTargetingKey#type}</td>
	// </tr>
	// </table>
	//
	GetCustomTargetingKeysByStatementResponse GetCustomTargetingKeysByStatementResponse `xml:"getCustomTargetingKeysByStatementResponse"`

	// GetCustomTargetingValuesByStatementResponse
	//
	// Gets a {@link CustomTargetingValuePage} of {@link CustomTargetingValue} objects that satisfy
	// the given {@link Statement#query}.
	//
	// <p>The {@code WHERE} clause in the {@link Statement#query} must always contain {@link
	// CustomTargetingValue#customTargetingKeyId} as one of its columns in a way that it is AND'ed
	// with the rest of the query. So, if you want to retrieve values for a known set of key ids,
	// valid {@link Statement#query} would look like:
	//
	// <ol>
	// <li>"WHERE customTargetingKeyId IN ('17','18','19')" retrieves all values that are associated
	// with keys having ids 17, 18, 19.
	// <li>"WHERE customTargetingKeyId = '17' AND name = 'red'" retrieves values that are associated
	// with keys having id 17 and value name is 'red'.
	// </ol>
	//
	// <p>The following fields are supported for filtering:
	//
	// <table>
	// <tr>
	// <th scope="col">PQL Property</th>
	// <th scope="col">Object Property</th>
	// </tr>
	// <tr>
	// <td>{@code id}</td>
	// <td>{@link CustomTargetingValue#id}</td>
	// </tr>
	// <tr>
	// <td>{@code customTargetingKeyId}</td>
	// <td>{@link CustomTargetingValue#customTargetingKeyId}</td>
	// </tr>
	// <tr>
	// <td>{@code name}</td>
	// <td>{@link CustomTargetingValue#name}</td>
	// </tr>
	// <tr>
	// <td>{@code displayName}</td>
	// <td>{@link CustomTargetingValue#displayName}</td>
	// </tr>
	// <tr>
	// <td>{@code matchType}</td>
	// <td>{@link CustomTargetingValue#matchType}</td>
	// </tr>
	// </table>
	//
	GetCustomTargetingValuesByStatementResponse GetCustomTargetingValuesByStatementResponse `xml:"getCustomTargetingValuesByStatementResponse"`

	// PerformCustomTargetingKeyActionResponse
	//
	// Performs actions on {@link CustomTargetingKey} objects that match the given {@link
	// Statement#query}.
	//
	PerformCustomTargetingKeyActionResponse PerformCustomTargetingKeyActionResponse `xml:"performCustomTargetingKeyActionResponse"`

	// PerformCustomTargetingValueActionResponse
	//
	// Performs actions on {@link CustomTargetingValue} objects that match the given {@link
	// Statement#query}.
	//
	PerformCustomTargetingValueActionResponse PerformCustomTargetingValueActionResponse `xml:"performCustomTargetingValueActionResponse"`

	// UpdateCustomTargetingKeysResponse
	//
	// Updates the specified {@link CustomTargetingKey} objects.
	//
	UpdateCustomTargetingKeysResponse UpdateCustomTargetingKeysResponse `xml:"updateCustomTargetingKeysResponse"`

	// UpdateCustomTargetingValuesResponse
	//
	// Updates the specified {@link CustomTargetingValue} objects.
	//
	UpdateCustomTargetingValuesResponse UpdateCustomTargetingValuesResponse `xml:"updateCustomTargetingValuesResponse"`
}

type Service struct {
	client          soap.Client
	wsdl            string
	networkCode     int
	applicationName string
}

func NewService(client soap.Client, networkCode int, applicationName string) CustomTargetingServiceInterface {
	return &Service{
		client:          client,
		wsdl:            wsdl,
		networkCode:     networkCode,
		applicationName: applicationName,
	}
}

// CustomTargetingServiceInterface
//
// Provides operations for creating, updating and retrieving {@link CustomTargetingKey} and {@link
// CustomTargetingValue} objects.
//

// CreateCustomTargetingKeys
//
// Creates new {@link CustomTargetingKey} objects.
//
// <p>The following fields are required:
//
// <ul>
// <li>{@link CustomTargetingKey#name}
// <li>{@link CustomTargetingKey#type}
// </ul>
func (s *Service) CreateCustomTargetingKeys(ctx context.Context, input CreateCustomTargetingKeys) (*CreateCustomTargetingKeysResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationCreateCustomTargetingKeys,
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
	return &envelope.Body.CreateCustomTargetingKeysResponse, nil
}

// CreateCustomTargetingValues
//
// Creates new {@link CustomTargetingValue} objects.
//
// <p>The following fields are required:
//
// <ul>
// <li>{@link CustomTargetingValue#customTargetingKeyId}
// <li>{@link CustomTargetingValue#name}
// </ul>
func (s *Service) CreateCustomTargetingValues(ctx context.Context, input CreateCustomTargetingValues) (*CreateCustomTargetingValuesResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationCreateCustomTargetingValues,
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
	return &envelope.Body.CreateCustomTargetingValuesResponse, nil
}

// GetCustomTargetingKeysByStatement
//
// Gets a {@link CustomTargetingKeyPage} of {@link CustomTargetingKey} objects that satisfy the
// given {@link Statement#query}. The following fields are supported for filtering:
//
// <table>
// <tr>
// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
// </tr>
// <tr>
// <tr>
// <td>{@code id}</td>
// <td>{@link CustomTargetingKey#id}</td>
// </tr>
// <tr>
// <td>{@code name}</td>
// <td>{@link CustomTargetingKey#name}</td>
// </tr>
// <tr>
// <td>{@code displayName}</td>
// <td>{@link CustomTargetingKey#displayName}</td>
// </tr>
// <tr>
// <td>{@code type}</td>
// <td>{@link CustomTargetingKey#type}</td>
// </tr>
// </table>
func (s *Service) GetCustomTargetingKeysByStatement(ctx context.Context, input GetCustomTargetingKeysByStatement) (*GetCustomTargetingKeysByStatementResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationGetCustomTargetingKeysByStatement,
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
	return &envelope.Body.GetCustomTargetingKeysByStatementResponse, nil
}

// GetCustomTargetingValuesByStatement
//
// Gets a {@link CustomTargetingValuePage} of {@link CustomTargetingValue} objects that satisfy
// the given {@link Statement#query}.
//
// <p>The {@code WHERE} clause in the {@link Statement#query} must always contain {@link
// CustomTargetingValue#customTargetingKeyId} as one of its columns in a way that it is AND'ed
// with the rest of the query. So, if you want to retrieve values for a known set of key ids,
// valid {@link Statement#query} would look like:
//
// <ol>
// <li>"WHERE customTargetingKeyId IN ('17','18','19')" retrieves all values that are associated
// with keys having ids 17, 18, 19.
// <li>"WHERE customTargetingKeyId = '17' AND name = 'red'" retrieves values that are associated
// with keys having id 17 and value name is 'red'.
// </ol>
//
// <p>The following fields are supported for filtering:
//
// <table>
// <tr>
// <th scope="col">PQL Property</th>
// <th scope="col">Object Property</th>
// </tr>
// <tr>
// <td>{@code id}</td>
// <td>{@link CustomTargetingValue#id}</td>
// </tr>
// <tr>
// <td>{@code customTargetingKeyId}</td>
// <td>{@link CustomTargetingValue#customTargetingKeyId}</td>
// </tr>
// <tr>
// <td>{@code name}</td>
// <td>{@link CustomTargetingValue#name}</td>
// </tr>
// <tr>
// <td>{@code displayName}</td>
// <td>{@link CustomTargetingValue#displayName}</td>
// </tr>
// <tr>
// <td>{@code matchType}</td>
// <td>{@link CustomTargetingValue#matchType}</td>
// </tr>
// </table>
func (s *Service) GetCustomTargetingValuesByStatement(ctx context.Context, input GetCustomTargetingValuesByStatement) (*GetCustomTargetingValuesByStatementResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationGetCustomTargetingValuesByStatement,
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
	return &envelope.Body.GetCustomTargetingValuesByStatementResponse, nil
}

func (s *Service) PerformCustomTargetingKeyAction(ctx context.Context, action requestbody.PerformAction, input PerformCustomTargetingKeyAction) (*PerformCustomTargetingKeyActionResponse, error) {
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
		OperationName:      OperationPerformCustomTargetingKeyAction,
		ActionName:         "customTargetingKeyAction",
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
	return &envelope.Body.PerformCustomTargetingKeyActionResponse, nil
}

func (s *Service) PerformCustomTargetingValueAction(ctx context.Context, action requestbody.PerformAction, input PerformCustomTargetingValueAction) (*PerformCustomTargetingValueActionResponse, error) {
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
		OperationName:      OperationPerformCustomTargetingValueAction,
		ActionName:         "customTargetingValueAction",
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
	return &envelope.Body.PerformCustomTargetingValueActionResponse, nil
}

// UpdateCustomTargetingKeys
//
// Updates the specified {@link CustomTargetingKey} objects.
func (s *Service) UpdateCustomTargetingKeys(ctx context.Context, input UpdateCustomTargetingKeys) (*UpdateCustomTargetingKeysResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationUpdateCustomTargetingKeys,
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
	return &envelope.Body.UpdateCustomTargetingKeysResponse, nil
}

// UpdateCustomTargetingValues
//
// Updates the specified {@link CustomTargetingValue} objects.
func (s *Service) UpdateCustomTargetingValues(ctx context.Context, input UpdateCustomTargetingValues) (*UpdateCustomTargetingValuesResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationUpdateCustomTargetingValues,
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
	return &envelope.Body.UpdateCustomTargetingValuesResponse, nil
}