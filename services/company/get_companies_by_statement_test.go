package company

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/mock/gomock"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/services/v202402/company"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
	"github.com/knakazawa99/gam-sdk-go/pkg/typeutil"
)

func TestService_GetCompaniesByStatement(t *testing.T) {
	type fields struct {
		wsdl            string
		networkCode     int
		applicationName string
	}
	type args struct {
		ctx   context.Context
		input company.GetCompaniesByStatement
	}

	companyTypeAdvertiser := company.CompanyTypeTypeAdvertiser
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(client *soap.MockClient)
		want    *company.GetCompaniesByStatementResponse
		wantErr error
	}{
		{
			name: "success/get companies by statement",
			fields: fields{
				wsdl:            "https://ads.google.com/apis/ads/publisher/v202402/CompanyService?wsdl",
				networkCode:     12345678,
				applicationName: "App_1",
			},
			args: args{
				ctx: context.Background(),
				input: company.GetCompaniesByStatement{
					FilterStatement: &company.Statement{
						Query: typeutil.ConvertToPointer("WHERE type = 'AGENCY'"),
					},
				},
			},
			prepare: func(client *soap.MockClient) {
				client.EXPECT().Call(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
					<soap:Header>
						<ResponseHeader xmlns="https://www.google.com/apis/ads/publisher/v202402">
							<requestId>6aba7760c2744b8ff3b7f567f04da5d1</requestId>
							<responseTime>378</responseTime>
						</ResponseHeader>
					</soap:Header>
					<soap:Body>
						<getCompaniesByStatementResponse xmlns="https://www.google.com/apis/ads/publisher/v202402">
							<rval>
								<totalResultSetSize>1</totalResultSetSize>
								<startIndex>0</startIndex>
								<results>
									<id>123456789</id>
									<name>company-2024042609</name>
									<type>ADVERTISER</type>
								    <lastModifiedDateTime>
								        <date>
											<year>2022</year>
											<month>4</month>	
											<day>26</day>
										</date>
								        <hour>12</hour>	
										<minute>0</minute>
										<second>0</second>	
                                    </lastModifiedDateTime>
								</results>
							</rval>
						</getCompaniesByStatementResponse>
					</soap:Body>
				</soap:Envelope>`), nil)
			},
			want: &company.GetCompaniesByStatementResponse{
				Rval: &company.CompanyPage{
					TotalResultSetSize: typeutil.ConvertToPointer(1),
					StartIndex:         typeutil.ConvertToPointer(0),
					Results: []*company.Company{
						{
							Id:   typeutil.Int64ToPtr(123456789),
							Name: typeutil.StringToPtr("company-2024042609"),
							Type: &companyTypeAdvertiser,
							LastModifiedDateTime: &company.DateTime{
								Date: &company.Date{
									Year:  typeutil.IntToPtr(2022),
									Month: typeutil.IntToPtr(4),
									Day:   typeutil.IntToPtr(26),
								},
								Hour:       typeutil.IntToPtr(12),
								Minute:     typeutil.IntToPtr(0),
								Second:     typeutil.IntToPtr(0),
								TimeZoneId: nil,
							},
						},
					},
				},
			},
		},
		{
			name: "fail/query statement ",
			fields: fields{
				wsdl:            "https://ads.google.com/apis/ads/publisher/v202402/CompanyService?wsdl",
				networkCode:     12345678,
				applicationName: "App_1",
			},
			args: args{
				ctx: context.Background(),
				input: company.GetCompaniesByStatement{
					FilterStatement: &company.Statement{
						Query: typeutil.ConvertToPointer("WHERE services = 'AGENCY'"),
					},
				},
			},
			prepare: func(client *soap.MockClient) {
				client.EXPECT().Call(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
					<soap:Header>
						<ResponseHeader xmlns="https://www.google.com/apis/ads/publisher/v202402">
							<requestId>0e552db90bfffafd40d6f2ecd1a8d81b</requestId>
							<responseTime>264</responseTime>
						</ResponseHeader>
					</soap:Header>
					<soap:Body>
						<soap:Fault>
							<faultcode>soap:Server</faultcode>
							<faultstring>[PublisherQueryLanguageContextError.UNEXECUTABLE @ ]</faultstring>
							<detail>
								<ApiExceptionFault xmlns="https://www.google.com/apis/ads/publisher/v202402">
									<message>[PublisherQueryLanguageContextError.UNEXECUTABLE @ ]</message>
									<errors xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
											xsi:type="PublisherQueryLanguageContextError">
										<fieldPath></fieldPath>
										<trigger></trigger>
										<errorString>PublisherQueryLanguageContextError.UNEXECUTABLE</errorString>
										<reason>UNEXECUTABLE</reason>
									</errors>
								</ApiExceptionFault>
							</detail>
						</soap:Fault>
					</soap:Body>
				</soap:Envelope>
				`), nil)
			},
			want:    nil,
			wantErr: company.NewPublisherQueryLanguageContextError("UNEXECUTABLE"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockSoapClient := soap.NewMockClient(ctrl)

			tt.prepare(mockSoapClient)
			s := company.NewService(
				mockSoapClient,
				tt.fields.networkCode,
				tt.fields.applicationName,
			)
			got, err := s.GetCompaniesByStatement(tt.args.ctx, tt.args.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetCompaniesByStatement() error = %v, wantErr %v", err, tt.wantErr)
			}
			opts := cmpopts.IgnoreUnexported(company.Company{})
			if diff := cmp.Diff(got, tt.want, opts); diff != "" {
				t.Errorf("GetCompaniesByStatement() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
