// Code generated by gamwsdl/main.go. DO NOT EDIT.
package activitygroup

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/requestbody"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
	gamxml "github.com/knakazawa99/gam-sdk-go/pkg/xml"
)

const (
	wsdl = "https://ads.google.com/apis/ads/publisher/v202402/ActivityGroupService"

	OperationCreateActivityGroups = "createActivityGroups"

	OperationGetActivityGroupsByStatement = "getActivityGroupsByStatement"

	OperationUpdateActivityGroups = "updateActivityGroups"
)

// ActivityGroupServiceInterface
//
// As of February 22, 2024 this service will become read only as part of Spotlight deprecation, <a
// href="https://support.google.com/admanager/answer/7519021#spotlight">learn more</a>..
//
// <p>Provides methods for creating, updating and retrieving {@link ActivityGroup} objects.
//
// <p>An activity group contains {@link Activity} objects. Activities have a many-to-one
// relationship with activity groups, meaning each activity can belong to only one activity group,
// but activity groups can have multiple activities. An activity group can be used to manage the
// activities it contains.
type ActivityGroupServiceInterface interface {

	// CreateActivityGroups
	//
	// Creates a new {@link ActivityGroup} objects.
	//
	CreateActivityGroups(ctx context.Context, input CreateActivityGroups) (*CreateActivityGroupsResponse, error)

	// GetActivityGroupsByStatement
	//
	// Gets an {@link ActivityGroupPage} of {@link ActivityGroup} objects that satisfy the given
	// {@link Statement#query}. The following fields are supported for filtering:
	//
	// <table>
	// <tr>
	// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
	// </tr>
	// <tr>
	// <td>{@code id}</td>
	// <td>{@link ActivityGroup#id}</td>
	// </tr>
	// <tr>
	// <td>{@code name}</td>
	// <td>{@link ActivityGroup#name}</td>
	// </tr>
	// <tr>
	// <td>{@code impressionsLookback}</td>
	// <td>{@link ActivityGroup#impressionsLookback}</td>
	// </tr>
	// <tr>
	// <td>{@code clicksLookback}</td>
	// <td>{@link ActivityGroup#clicksLookback}</td>
	// </tr>
	// <tr>
	// <td>{@code status}</td>
	// <td>{@link ActivityGroup#status}</td>
	// </tr>
	// </table>
	//
	GetActivityGroupsByStatement(ctx context.Context, input GetActivityGroupsByStatement) (*GetActivityGroupsByStatementResponse, error)

	// UpdateActivityGroups
	//
	// Updates the specified {@link ActivityGroup} objects.
	//
	UpdateActivityGroups(ctx context.Context, input UpdateActivityGroups) (*UpdateActivityGroupsResponse, error)
}

type Envelope struct {
	xmlName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	xmlName xml.Name   `xml:"Body"`
	Fault   soap.Fault `xml:"Fault"`

	// CreateActivityGroupsResponse
	//
	// Creates a new {@link ActivityGroup} objects.
	//
	CreateActivityGroupsResponse CreateActivityGroupsResponse `xml:"createActivityGroupsResponse"`

	// GetActivityGroupsByStatementResponse
	//
	// Gets an {@link ActivityGroupPage} of {@link ActivityGroup} objects that satisfy the given
	// {@link Statement#query}. The following fields are supported for filtering:
	//
	// <table>
	// <tr>
	// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
	// </tr>
	// <tr>
	// <td>{@code id}</td>
	// <td>{@link ActivityGroup#id}</td>
	// </tr>
	// <tr>
	// <td>{@code name}</td>
	// <td>{@link ActivityGroup#name}</td>
	// </tr>
	// <tr>
	// <td>{@code impressionsLookback}</td>
	// <td>{@link ActivityGroup#impressionsLookback}</td>
	// </tr>
	// <tr>
	// <td>{@code clicksLookback}</td>
	// <td>{@link ActivityGroup#clicksLookback}</td>
	// </tr>
	// <tr>
	// <td>{@code status}</td>
	// <td>{@link ActivityGroup#status}</td>
	// </tr>
	// </table>
	//
	GetActivityGroupsByStatementResponse GetActivityGroupsByStatementResponse `xml:"getActivityGroupsByStatementResponse"`

	// UpdateActivityGroupsResponse
	//
	// Updates the specified {@link ActivityGroup} objects.
	//
	UpdateActivityGroupsResponse UpdateActivityGroupsResponse `xml:"updateActivityGroupsResponse"`
}

type Service struct {
	client          soap.Client
	wsdl            string
	networkCode     int
	applicationName string
}

func NewService(client soap.Client, networkCode int, applicationName string) ActivityGroupServiceInterface {
	return &Service{
		client:          client,
		wsdl:            wsdl,
		networkCode:     networkCode,
		applicationName: applicationName,
	}
}

// ActivityGroupServiceInterface
//
// As of February 22, 2024 this service will become read only as part of Spotlight deprecation, <a
// href="https://support.google.com/admanager/answer/7519021#spotlight">learn more</a>..
//
// <p>Provides methods for creating, updating and retrieving {@link ActivityGroup} objects.
//
// <p>An activity group contains {@link Activity} objects. Activities have a many-to-one
// relationship with activity groups, meaning each activity can belong to only one activity group,
// but activity groups can have multiple activities. An activity group can be used to manage the
// activities it contains.
//

// CreateActivityGroups
//
// Creates a new {@link ActivityGroup} objects.
func (s *Service) CreateActivityGroups(ctx context.Context, input CreateActivityGroups) (*CreateActivityGroupsResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationCreateActivityGroups,
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
	return &envelope.Body.CreateActivityGroupsResponse, nil
}

// GetActivityGroupsByStatement
//
// Gets an {@link ActivityGroupPage} of {@link ActivityGroup} objects that satisfy the given
// {@link Statement#query}. The following fields are supported for filtering:
//
// <table>
// <tr>
// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
// </tr>
// <tr>
// <td>{@code id}</td>
// <td>{@link ActivityGroup#id}</td>
// </tr>
// <tr>
// <td>{@code name}</td>
// <td>{@link ActivityGroup#name}</td>
// </tr>
// <tr>
// <td>{@code impressionsLookback}</td>
// <td>{@link ActivityGroup#impressionsLookback}</td>
// </tr>
// <tr>
// <td>{@code clicksLookback}</td>
// <td>{@link ActivityGroup#clicksLookback}</td>
// </tr>
// <tr>
// <td>{@code status}</td>
// <td>{@link ActivityGroup#status}</td>
// </tr>
// </table>
func (s *Service) GetActivityGroupsByStatement(ctx context.Context, input GetActivityGroupsByStatement) (*GetActivityGroupsByStatementResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationGetActivityGroupsByStatement,
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
	return &envelope.Body.GetActivityGroupsByStatementResponse, nil
}

// UpdateActivityGroups
//
// Updates the specified {@link ActivityGroup} objects.
func (s *Service) UpdateActivityGroups(ctx context.Context, input UpdateActivityGroups) (*UpdateActivityGroupsResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationUpdateActivityGroups,
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
	return &envelope.Body.UpdateActivityGroupsResponse, nil
}