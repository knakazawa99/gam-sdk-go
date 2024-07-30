// Code generated by gamwsdl/main.go. DO NOT EDIT.
package lineitem

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/requestbody"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
	gamxml "github.com/knakazawa99/gam-sdk-go/pkg/xml"
)

const (
	wsdl = "https://ads.google.com/apis/ads/publisher/v202405/LineItemService"

	OperationCreateLineItems = "createLineItems"

	OperationGetLineItemsByStatement = "getLineItemsByStatement"

	OperationPerformLineItemAction = "performLineItemAction"

	OperationUpdateLineItems = "updateLineItems"
)

// LineItemServiceInterface
//
// Provides methods for creating, updating and retrieving {@link LineItem} objects.
//
// <p>Line items define the campaign. For example, line items define:
//
// <ul>
// <li>a budget
// <li>a span of time to run
// <li>ad unit targeting
// </ul>
//
// <p>In short, line items connect all of the elements of an ad campaign.
//
// <p>Line items and creatives can be associated with each other through {@link
// LineItemCreativeAssociation} objects. An ad unit will host a creative through both this
// association and the {@link LineItem#targeting} to it. The delivery of a line item depends on its
// priority. More information on line item priorities can be found on the <a
// href="https://support.google.com/admanager/answer/177279">Ad Manager Help Center</a>.
type LineItemServiceInterface interface {

	// CreateLineItems
	//
	// Creates new {@link LineItem} objects.
	//
	CreateLineItems(ctx context.Context, input CreateLineItems) (*CreateLineItemsResponse, error)

	// GetLineItemsByStatement
	//
	// Gets a {@link LineItemPage} of {@link LineItem} objects that satisfy the given {@link
	// Statement#query}. The following fields are supported for filtering:
	//
	// <table>
	// <tbody>
	// <tr>
	// <th>PQL property</th>
	// <th>Entity property</th>
	// </tr>
	// <tr>
	// <td>
	// {@code CostType}
	// </td>
	// <td>
	// {@link LineItem#costType}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code CreationDateTime}
	// </td>
	// <td>
	// {@link LineItem#creationDateTime}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code DeliveryRateType}
	// </td>
	// <td>
	// {@link LineItem#deliveryRateType}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code EndDateTime}
	// </td>
	// <td>
	// {@link LineItem#endDateTime}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code ExternalId}
	// </td>
	// <td>
	// {@link LineItem#externalId}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code Id}
	// </td>
	// <td>
	// {@link LineItem#id}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code IsMissingCreatives}
	// </td>
	// <td>
	// {@link LineItem#isMissingCreatives}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code IsSetTopBoxEnabled}
	// </td>
	// <td>
	// {@link LineItem#isSetTopBoxEnabled}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code LastModifiedDateTime}
	// </td>
	// <td>
	// {@link LineItem#lastModifiedDateTime}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code LineItemType}
	// </td>
	// <td>
	// {@link LineItem#lineItemType}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code Name}
	// </td>
	// <td>
	// {@link LineItem#name}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code OrderId}
	// </td>
	// <td>
	// {@link LineItem#orderId}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code StartDateTime}
	// </td>
	// <td>
	// {@link LineItem#startDateTime}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code Status}
	// </td>
	// <td>
	// {@link LineItem#status}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code UnitsBought}
	// </td>
	// <td>
	// {@link LineItem#unitsBought}
	// </td>
	// </tr>
	// </tbody>
	// </table>
	//
	GetLineItemsByStatement(ctx context.Context, input GetLineItemsByStatement) (*GetLineItemsByStatementResponse, error)

	// PerformLineItemAction
	//
	// Performs actions on {@link LineItem} objects that match the given {@link Statement#query}.
	//
	PerformLineItemAction(ctx context.Context, action requestbody.PerformAction, input PerformLineItemAction) (*PerformLineItemActionResponse, error)

	// UpdateLineItems
	//
	// Updates the specified {@link LineItem} objects.
	//
	UpdateLineItems(ctx context.Context, input UpdateLineItems) (*UpdateLineItemsResponse, error)
}

type Envelope struct {
	xmlName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	xmlName xml.Name   `xml:"Body"`
	Fault   soap.Fault `xml:"Fault"`

	// CreateLineItemsResponse
	//
	// Creates new {@link LineItem} objects.
	//
	CreateLineItemsResponse CreateLineItemsResponse `xml:"createLineItemsResponse"`

	// GetLineItemsByStatementResponse
	//
	// Gets a {@link LineItemPage} of {@link LineItem} objects that satisfy the given {@link
	// Statement#query}. The following fields are supported for filtering:
	//
	// <table>
	// <tbody>
	// <tr>
	// <th>PQL property</th>
	// <th>Entity property</th>
	// </tr>
	// <tr>
	// <td>
	// {@code CostType}
	// </td>
	// <td>
	// {@link LineItem#costType}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code CreationDateTime}
	// </td>
	// <td>
	// {@link LineItem#creationDateTime}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code DeliveryRateType}
	// </td>
	// <td>
	// {@link LineItem#deliveryRateType}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code EndDateTime}
	// </td>
	// <td>
	// {@link LineItem#endDateTime}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code ExternalId}
	// </td>
	// <td>
	// {@link LineItem#externalId}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code Id}
	// </td>
	// <td>
	// {@link LineItem#id}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code IsMissingCreatives}
	// </td>
	// <td>
	// {@link LineItem#isMissingCreatives}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code IsSetTopBoxEnabled}
	// </td>
	// <td>
	// {@link LineItem#isSetTopBoxEnabled}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code LastModifiedDateTime}
	// </td>
	// <td>
	// {@link LineItem#lastModifiedDateTime}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code LineItemType}
	// </td>
	// <td>
	// {@link LineItem#lineItemType}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code Name}
	// </td>
	// <td>
	// {@link LineItem#name}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code OrderId}
	// </td>
	// <td>
	// {@link LineItem#orderId}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code StartDateTime}
	// </td>
	// <td>
	// {@link LineItem#startDateTime}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code Status}
	// </td>
	// <td>
	// {@link LineItem#status}
	// </td>
	// </tr>
	// <tr>
	// <td>
	// {@code UnitsBought}
	// </td>
	// <td>
	// {@link LineItem#unitsBought}
	// </td>
	// </tr>
	// </tbody>
	// </table>
	//
	GetLineItemsByStatementResponse GetLineItemsByStatementResponse `xml:"getLineItemsByStatementResponse"`

	// PerformLineItemActionResponse
	//
	// Performs actions on {@link LineItem} objects that match the given {@link Statement#query}.
	//
	PerformLineItemActionResponse PerformLineItemActionResponse `xml:"performLineItemActionResponse"`

	// UpdateLineItemsResponse
	//
	// Updates the specified {@link LineItem} objects.
	//
	UpdateLineItemsResponse UpdateLineItemsResponse `xml:"updateLineItemsResponse"`
}

type Service struct {
	client          soap.Client
	wsdl            string
	networkCode     int
	applicationName string
}

func NewService(client soap.Client, networkCode int, applicationName string) LineItemServiceInterface {
	return &Service{
		client:          client,
		wsdl:            wsdl,
		networkCode:     networkCode,
		applicationName: applicationName,
	}
}

// LineItemServiceInterface
//
// Provides methods for creating, updating and retrieving {@link LineItem} objects.
//
// <p>Line items define the campaign. For example, line items define:
//
// <ul>
// <li>a budget
// <li>a span of time to run
// <li>ad unit targeting
// </ul>
//
// <p>In short, line items connect all of the elements of an ad campaign.
//
// <p>Line items and creatives can be associated with each other through {@link
// LineItemCreativeAssociation} objects. An ad unit will host a creative through both this
// association and the {@link LineItem#targeting} to it. The delivery of a line item depends on its
// priority. More information on line item priorities can be found on the <a
// href="https://support.google.com/admanager/answer/177279">Ad Manager Help Center</a>.
//

// CreateLineItems
//
// Creates new {@link LineItem} objects.
func (s *Service) CreateLineItems(ctx context.Context, input CreateLineItems) (*CreateLineItemsResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationCreateLineItems,
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
	return &envelope.Body.CreateLineItemsResponse, nil
}

// GetLineItemsByStatement
//
// Gets a {@link LineItemPage} of {@link LineItem} objects that satisfy the given {@link
// Statement#query}. The following fields are supported for filtering:
//
// <table>
// <tbody>
// <tr>
// <th>PQL property</th>
// <th>Entity property</th>
// </tr>
// <tr>
// <td>
// {@code CostType}
// </td>
// <td>
// {@link LineItem#costType}
// </td>
// </tr>
// <tr>
// <td>
// {@code CreationDateTime}
// </td>
// <td>
// {@link LineItem#creationDateTime}
// </td>
// </tr>
// <tr>
// <td>
// {@code DeliveryRateType}
// </td>
// <td>
// {@link LineItem#deliveryRateType}
// </td>
// </tr>
// <tr>
// <td>
// {@code EndDateTime}
// </td>
// <td>
// {@link LineItem#endDateTime}
// </td>
// </tr>
// <tr>
// <td>
// {@code ExternalId}
// </td>
// <td>
// {@link LineItem#externalId}
// </td>
// </tr>
// <tr>
// <td>
// {@code Id}
// </td>
// <td>
// {@link LineItem#id}
// </td>
// </tr>
// <tr>
// <td>
// {@code IsMissingCreatives}
// </td>
// <td>
// {@link LineItem#isMissingCreatives}
// </td>
// </tr>
// <tr>
// <td>
// {@code IsSetTopBoxEnabled}
// </td>
// <td>
// {@link LineItem#isSetTopBoxEnabled}
// </td>
// </tr>
// <tr>
// <td>
// {@code LastModifiedDateTime}
// </td>
// <td>
// {@link LineItem#lastModifiedDateTime}
// </td>
// </tr>
// <tr>
// <td>
// {@code LineItemType}
// </td>
// <td>
// {@link LineItem#lineItemType}
// </td>
// </tr>
// <tr>
// <td>
// {@code Name}
// </td>
// <td>
// {@link LineItem#name}
// </td>
// </tr>
// <tr>
// <td>
// {@code OrderId}
// </td>
// <td>
// {@link LineItem#orderId}
// </td>
// </tr>
// <tr>
// <td>
// {@code StartDateTime}
// </td>
// <td>
// {@link LineItem#startDateTime}
// </td>
// </tr>
// <tr>
// <td>
// {@code Status}
// </td>
// <td>
// {@link LineItem#status}
// </td>
// </tr>
// <tr>
// <td>
// {@code UnitsBought}
// </td>
// <td>
// {@link LineItem#unitsBought}
// </td>
// </tr>
// </tbody>
// </table>
func (s *Service) GetLineItemsByStatement(ctx context.Context, input GetLineItemsByStatement) (*GetLineItemsByStatementResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationGetLineItemsByStatement,
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
	return &envelope.Body.GetLineItemsByStatementResponse, nil
}

func (s *Service) PerformLineItemAction(ctx context.Context, action requestbody.PerformAction, input PerformLineItemAction) (*PerformLineItemActionResponse, error) {
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
		OperationName:      OperationPerformLineItemAction,
		ActionName:         "lineItemAction",
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
	return &envelope.Body.PerformLineItemActionResponse, nil
}

// UpdateLineItems
//
// Updates the specified {@link LineItem} objects.
func (s *Service) UpdateLineItems(ctx context.Context, input UpdateLineItems) (*UpdateLineItemsResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationUpdateLineItems,
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
	return &envelope.Body.UpdateLineItemsResponse, nil
}
