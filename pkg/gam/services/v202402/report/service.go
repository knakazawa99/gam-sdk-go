// Code generated by gamwsdl/main.go. DO NOT EDIT.
package report

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/requestbody"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
	gamxml "github.com/knakazawa99/gam-sdk-go/pkg/xml"
)

const (
	wsdl = "https://ads.google.com/apis/ads/publisher/v202402/ReportService"

	OperationGetReportDownloadURL = "getReportDownloadURL"

	OperationGetReportDownloadUrlWithOptions = "getReportDownloadUrlWithOptions"

	OperationGetReportJobStatus = "getReportJobStatus"

	OperationGetSavedQueriesByStatement = "getSavedQueriesByStatement"

	OperationRunReportJob = "runReportJob"
)

// ReportServiceInterface
//
// Provides methods for executing a {@link ReportJob} and retrieving performance and statistics
// about ad campaigns, networks, inventory and sales.
//
// <p>Follow the steps outlined below:
//
// <p>
//
// <ul>
// <li>Create the {@code ReportJob} object by invoking {@link ReportService#runReportJob}.
// <li>Poll the report job status using {@link ReportService#getReportJobStatus}.
// <li>Continue to poll until the status is equal to {@link ReportJobStatus#COMPLETED} or {@link
// ReportJobStatus#FAILED}.
// <li>If successful, fetch the URL for downloading the report by invoking {@link
// ReportService#getReportDownloadURL}.
// </ul>
//
// <h4>Test network behavior</h4>
//
// <p>The networks created using {@link NetworkService#makeTestNetwork} are unable to provide
// reports that would be comparable to the production environment because reports require traffic
// history. In the test networks, reports will consistently return no data for all reports.
type ReportServiceInterface interface {

	// GetReportDownloadURL
	//
	// Returns the URL at which the report file can be downloaded.
	//
	// <p>The report will be generated as a gzip archive, containing the report file itself.
	//
	GetReportDownloadURL(ctx context.Context, input GetReportDownloadURL) (*GetReportDownloadURLResponse, error)

	// GetReportDownloadUrlWithOptions
	//
	// Returns the URL at which the report file can be downloaded, and allows for customization of the
	// downloaded report.
	//
	// <p>By default, the report will be generated as a gzip archive, containing the report file
	// itself. This can be changed by setting {@link ReportDownloadOptions#useGzipCompression} to
	// false.
	//
	GetReportDownloadUrlWithOptions(ctx context.Context, input GetReportDownloadUrlWithOptions) (*GetReportDownloadUrlWithOptionsResponse, error)

	// GetReportJobStatus
	//
	// Returns the {@link ReportJobStatus} of the report job with the specified ID.
	//
	GetReportJobStatus(ctx context.Context, input GetReportJobStatus) (*GetReportJobStatusResponse, error)

	// GetSavedQueriesByStatement
	//
	// Retrieves a page of the saved queries either created by or shared with the current user. Each
	// {@link SavedQuery} in the page, if it is compatible with the current API version, will contain
	// a {@link ReportQuery} object which can be optionally modified and used to create a {@link
	// ReportJob}. This can then be passed to {@link ReportService#runReportJob}. The following fields
	// are supported for filtering:
	//
	// <table>
	// <tr>
	// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
	// </tr>
	// <tr>
	// <td>{@code id}</td>
	// <td>{@link SavedQuery#id}</td>
	// </tr>
	// <tr>
	// <td>{@code name}</td>
	// <td>{@link SavedQuery#name}</td>
	// </tr>
	// </table>
	//
	GetSavedQueriesByStatement(ctx context.Context, input GetSavedQueriesByStatement) (*GetSavedQueriesByStatementResponse, error)

	// RunReportJob
	//
	// Initiates the execution of a {@link ReportQuery} on the server.
	//
	// <p>The following fields are required:
	//
	// <ul>
	// <li>{@link ReportJob#reportQuery}
	// </ul>
	//
	RunReportJob(ctx context.Context, input RunReportJob) (*RunReportJobResponse, error)
}

type Envelope struct {
	xmlName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	xmlName xml.Name   `xml:"Body"`
	Fault   soap.Fault `xml:"Fault"`

	// GetReportDownloadURLResponse
	//
	// Returns the URL at which the report file can be downloaded.
	//
	// <p>The report will be generated as a gzip archive, containing the report file itself.
	//
	GetReportDownloadURLResponse GetReportDownloadURLResponse `xml:"getReportDownloadURLResponse"`

	// GetReportDownloadUrlWithOptionsResponse
	//
	// Returns the URL at which the report file can be downloaded, and allows for customization of the
	// downloaded report.
	//
	// <p>By default, the report will be generated as a gzip archive, containing the report file
	// itself. This can be changed by setting {@link ReportDownloadOptions#useGzipCompression} to
	// false.
	//
	GetReportDownloadUrlWithOptionsResponse GetReportDownloadUrlWithOptionsResponse `xml:"getReportDownloadUrlWithOptionsResponse"`

	// GetReportJobStatusResponse
	//
	// Returns the {@link ReportJobStatus} of the report job with the specified ID.
	//
	GetReportJobStatusResponse GetReportJobStatusResponse `xml:"getReportJobStatusResponse"`

	// GetSavedQueriesByStatementResponse
	//
	// Retrieves a page of the saved queries either created by or shared with the current user. Each
	// {@link SavedQuery} in the page, if it is compatible with the current API version, will contain
	// a {@link ReportQuery} object which can be optionally modified and used to create a {@link
	// ReportJob}. This can then be passed to {@link ReportService#runReportJob}. The following fields
	// are supported for filtering:
	//
	// <table>
	// <tr>
	// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
	// </tr>
	// <tr>
	// <td>{@code id}</td>
	// <td>{@link SavedQuery#id}</td>
	// </tr>
	// <tr>
	// <td>{@code name}</td>
	// <td>{@link SavedQuery#name}</td>
	// </tr>
	// </table>
	//
	GetSavedQueriesByStatementResponse GetSavedQueriesByStatementResponse `xml:"getSavedQueriesByStatementResponse"`

	// RunReportJobResponse
	//
	// Initiates the execution of a {@link ReportQuery} on the server.
	//
	// <p>The following fields are required:
	//
	// <ul>
	// <li>{@link ReportJob#reportQuery}
	// </ul>
	//
	RunReportJobResponse RunReportJobResponse `xml:"runReportJobResponse"`
}

type Service struct {
	client          soap.Client
	wsdl            string
	networkCode     int
	applicationName string
}

func NewService(client soap.Client, networkCode int, applicationName string) ReportServiceInterface {
	return &Service{
		client:          client,
		wsdl:            wsdl,
		networkCode:     networkCode,
		applicationName: applicationName,
	}
}

// ReportServiceInterface
//
// Provides methods for executing a {@link ReportJob} and retrieving performance and statistics
// about ad campaigns, networks, inventory and sales.
//
// <p>Follow the steps outlined below:
//
// <p>
//
// <ul>
// <li>Create the {@code ReportJob} object by invoking {@link ReportService#runReportJob}.
// <li>Poll the report job status using {@link ReportService#getReportJobStatus}.
// <li>Continue to poll until the status is equal to {@link ReportJobStatus#COMPLETED} or {@link
// ReportJobStatus#FAILED}.
// <li>If successful, fetch the URL for downloading the report by invoking {@link
// ReportService#getReportDownloadURL}.
// </ul>
//
// <h4>Test network behavior</h4>
//
// <p>The networks created using {@link NetworkService#makeTestNetwork} are unable to provide
// reports that would be comparable to the production environment because reports require traffic
// history. In the test networks, reports will consistently return no data for all reports.
//

// GetReportDownloadURL
//
// Returns the URL at which the report file can be downloaded.
//
// <p>The report will be generated as a gzip archive, containing the report file itself.
func (s *Service) GetReportDownloadURL(ctx context.Context, input GetReportDownloadURL) (*GetReportDownloadURLResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationGetReportDownloadURL,
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
	return &envelope.Body.GetReportDownloadURLResponse, nil
}

// GetReportDownloadUrlWithOptions
//
// Returns the URL at which the report file can be downloaded, and allows for customization of the
// downloaded report.
//
// <p>By default, the report will be generated as a gzip archive, containing the report file
// itself. This can be changed by setting {@link ReportDownloadOptions#useGzipCompression} to
// false.
func (s *Service) GetReportDownloadUrlWithOptions(ctx context.Context, input GetReportDownloadUrlWithOptions) (*GetReportDownloadUrlWithOptionsResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationGetReportDownloadUrlWithOptions,
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
	return &envelope.Body.GetReportDownloadUrlWithOptionsResponse, nil
}

// GetReportJobStatus
//
// Returns the {@link ReportJobStatus} of the report job with the specified ID.
func (s *Service) GetReportJobStatus(ctx context.Context, input GetReportJobStatus) (*GetReportJobStatusResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationGetReportJobStatus,
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
	return &envelope.Body.GetReportJobStatusResponse, nil
}

// GetSavedQueriesByStatement
//
// Retrieves a page of the saved queries either created by or shared with the current user. Each
// {@link SavedQuery} in the page, if it is compatible with the current API version, will contain
// a {@link ReportQuery} object which can be optionally modified and used to create a {@link
// ReportJob}. This can then be passed to {@link ReportService#runReportJob}. The following fields
// are supported for filtering:
//
// <table>
// <tr>
// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
// </tr>
// <tr>
// <td>{@code id}</td>
// <td>{@link SavedQuery#id}</td>
// </tr>
// <tr>
// <td>{@code name}</td>
// <td>{@link SavedQuery#name}</td>
// </tr>
// </table>
func (s *Service) GetSavedQueriesByStatement(ctx context.Context, input GetSavedQueriesByStatement) (*GetSavedQueriesByStatementResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationGetSavedQueriesByStatement,
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
	return &envelope.Body.GetSavedQueriesByStatementResponse, nil
}

// RunReportJob
//
// Initiates the execution of a {@link ReportQuery} on the server.
//
// <p>The following fields are required:
//
// <ul>
// <li>{@link ReportJob#reportQuery}
// </ul>
func (s *Service) RunReportJob(ctx context.Context, input RunReportJob) (*RunReportJobResponse, error) {
	parameterXML, err := gamxml.DeepMarshal(input, true)
	data := requestbody.DefaultRequestBody{
		NetworkCode:     s.networkCode,
		ApplicationName: s.applicationName,
		OperationName:   OperationRunReportJob,
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
	return &envelope.Body.RunReportJobResponse, nil
}
