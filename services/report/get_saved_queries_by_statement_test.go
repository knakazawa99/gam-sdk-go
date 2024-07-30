package report

import (
	"context"
	"errors"
	"testing"

	"github.com/knakazawa99/gam-sdk-go/pkg/typeutil"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/mock/gomock"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/services/v202402/report"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
)

func TestService_GetSavedQueriesByStatement(t *testing.T) {
	type fields struct {
		wsdl            string
		networkCode     int
		applicationName string
	}
	type args struct {
		ctx   context.Context
		input report.GetSavedQueriesByStatement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(client *soap.MockClient)
		want    *report.GetSavedQueriesByStatementResponse
		wantErr error
	}{
		{
			name: "success/get saved queries by statement",
			fields: fields{
				wsdl:            "https://ads.google.com/apis/ads/publisher/v202402/OrderService?wsdl",
				networkCode:     12345678,
				applicationName: "App_1",
			},
			args: args{
				ctx: context.Background(),
				input: report.GetSavedQueriesByStatement{
					FilterStatement: &report.Statement{
						Query: typeutil.ConvertToPointer("WHERE id = 123456789"),
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
						<getSavedQueriesByStatementResponse xmlns="https://www.google.com/apis/ads/publisher/v202402">
							<rval>
								<totalResultSetSize>1</totalResultSetSize>
								<startIndex>0</startIndex>
								<results>
									<id>123456789</id>
									<name>report-123456</name>
								</results>
							</rval>
						</getSavedQueriesByStatementResponse>
					</soap:Body>
				</soap:Envelope>`), nil)
			},
			want: &report.GetSavedQueriesByStatementResponse{
				Rval: &report.SavedQueryPage{
					TotalResultSetSize: typeutil.ConvertToPointer(1),
					StartIndex:         typeutil.ConvertToPointer(0),
					Results: []*report.SavedQuery{
						{
							Id:   typeutil.ConvertToPointer(int64(123456789)),
							Name: typeutil.ConvertToPointer("report-123456"),
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockSoapClient := soap.NewMockClient(ctrl)

			tt.prepare(mockSoapClient)
			s := report.NewService(
				mockSoapClient,
				tt.fields.networkCode,
				tt.fields.applicationName,
			)
			got, err := s.GetSavedQueriesByStatement(tt.args.ctx, tt.args.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetSavedQueriesByStatement() error = %v, wantErr %v", err, tt.wantErr)
			}
			opts := cmpopts.IgnoreUnexported(report.ReportQuery{})
			if diff := cmp.Diff(tt.want, got, opts); diff != "" {
				t.Errorf("GetSavedQueriesByStatement() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
