package report

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/mock/gomock"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/services/v202402/report"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
	"github.com/knakazawa99/gam-sdk-go/pkg/typeutil"
)

func TestService_GetReportJobStatus(t *testing.T) {
	type fields struct {
		wsdl            string
		networkCode     int
		applicationName string
	}
	type args struct {
		ctx   context.Context
		input report.GetReportJobStatus
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(client *soap.MockClient)
		want    *report.GetReportJobStatusResponse
		wantErr bool
	}{
		{
			name: "success/get report download url",
			fields: fields{
				wsdl:            "https://ads.google.com/apis/ads/publisher/v202402/ReportService?wsdl",
				networkCode:     12345678,
				applicationName: "App_1",
			},
			args: args{
				ctx: context.Background(),
				input: report.GetReportJobStatus{
					ReportJobId: typeutil.ConvertToPointer(int64(123456789)),
				},
			},
			prepare: func(client *soap.MockClient) {
				client.EXPECT().Call(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
					<soap:Header>
						<ResponseHeader xmlns="https://www.google.com/apis/ads/publisher/v202402">
							<requestId>19dd5e334acc53708bb1d7b5855f2a6b</requestId><responseTime>278</responseTime>
						</ResponseHeader>
					</soap:Header>
					<soap:Body>
						<getReportJobStatusResponse xmlns="https://www.google.com/apis/ads/publisher/v202402">
						<rval>COMPLETED</rval></getReportJobStatusResponse>
					</soap:Body></soap:Envelope>
				`), nil)
			},
			want: &report.GetReportJobStatusResponse{
				Rval: typeutil.ConvertToPointer(report.ReportJobStatusCompleted),
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
			got, err := s.GetReportJobStatus(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReportJobStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported()); diff != "" {
				t.Errorf("GetReportJobStatus() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
