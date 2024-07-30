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

func TestService_CreateCompanies(t *testing.T) {
	type fields struct {
		wsdl            string
		networkCode     int
		applicationName string
	}
	type args struct {
		ctx       context.Context
		companies company.CreateCompanies
	}

	companyCreditStatusActive := company.CompanyCreditStatusTypeActive
	companyTypeAdvertiser := company.CompanyTypeTypeAdvertiser
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(client *soap.MockClient)
		want    *company.CreateCompaniesResponse
		wantErr error
	}{
		{
			name: "success/create companies",
			fields: fields{
				wsdl:            "https://ads.google.com/apis/ads/publisher/v202402/CompanyService?wsdl",
				networkCode:     12345678,
				applicationName: "App_1",
			},
			args: args{
				ctx: context.Background(),
				companies: company.CreateCompanies{
					Companies: []*company.Company{
						{
							Name: typeutil.StringToPtr("company-2024042609"),
							Type: &companyTypeAdvertiser,
						},
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
						<createCompaniesResponse xmlns="https://www.google.com/apis/ads/publisher/v202402">
							<rval>
								<id>12345678</id>
								<name>company-12345678</name>
								<type>ADVERTISER</type>
								<creditStatus>ACTIVE</creditStatus>
								<lastModifiedDateTime>
									<date>
										<year>2024</year>
										<month>4</month>
										<day>27</day>
									</date>
									<hour>12</hour>
									<minute>00</minute>
									<second>00</second>
									<timeZoneId>America/Los_Angeles</timeZoneId>
								</lastModifiedDateTime>
							</rval>
						</createCompaniesResponse>
					</soap:Body>
				</soap:Envelope>
				`), nil)
			},
			want: &company.CreateCompaniesResponse{
				Rval: []*company.Company{
					{
						Id:           typeutil.Int64ToPtr(12345678),
						Name:         typeutil.StringToPtr("company-12345678"),
						Type:         &companyTypeAdvertiser,
						CreditStatus: &companyCreditStatusActive,
						LastModifiedDateTime: &company.DateTime{
							Date: &company.Date{
								Year:  typeutil.IntToPtr(2024),
								Month: typeutil.IntToPtr(4),
								Day:   typeutil.IntToPtr(27),
							},
							Hour:       typeutil.IntToPtr(12),
							Minute:     typeutil.IntToPtr(0),
							Second:     typeutil.IntToPtr(0),
							TimeZoneId: typeutil.StringToPtr("America/Los_Angeles"),
						},
					},
				},
			},
		},
		{
			name: "fail/not unique company name",
			fields: fields{
				wsdl:            "https://ads.google.com/apis/ads/publisher/v202402/CompanyService?wsdl",
				networkCode:     12345678,
				applicationName: "App_1",
			},
			args: args{
				ctx: context.Background(),
				companies: company.CreateCompanies{
					Companies: []*company.Company{
						{
							Name: typeutil.StringToPtr("company-2024042609"),
							Type: &companyTypeAdvertiser,
						},
					},
				},
			},
			prepare: func(client *soap.MockClient) {
				client.EXPECT().Call(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
						<soap:Header>
							<ResponseHeader xmlns="https://www.google.com/apis/ads/publisher/v202402">
								<requestId>7f5b4a5ed975cf2c0161d8468b446e97</requestId>
								<responseTime>467</responseTime>
							</ResponseHeader>
						</soap:Header>
						<soap:Body>
							<soap:Fault>
								<faultcode>soap:Server</faultcode>
								<faultstring>[UniqueError.NOT_UNIQUE @ [0].name; trigger:'company-2024042609']</faultstring>
								<detail>
									<ApiExceptionFault xmlns="https://www.google.com/apis/ads/publisher/v202402">
										<message>[UniqueError.NOT_UNIQUE @ [0].name; trigger:'company-2024042609']</message>
										<errors xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="UniqueError">
											<fieldPath>[0].name</fieldPath>
											<fieldPathElements>
												<index>0</index>
											</fieldPathElements>
											<fieldPathElements>
												<field>name</field>
											</fieldPathElements>
											<trigger>company-2024042609</trigger>
											<errorString>UniqueError.NOT_UNIQUE</errorString>
										</errors>
									</ApiExceptionFault>
								</detail>
							</soap:Fault>
						</soap:Body>
					</soap:Envelope>
				`), nil)
			},
			want:    nil,
			wantErr: company.NewUniqueError("company-2024042609"),
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
			got, err := s.CreateCompanies(tt.args.ctx, tt.args.companies)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("CreateCompanies() error = %v, wantErr %v", err, tt.wantErr)
			}
			opts := cmpopts.IgnoreUnexported(company.Company{})
			if diff := cmp.Diff(tt.want, got, opts); diff != "" {
				t.Errorf("CreateCompanies() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
