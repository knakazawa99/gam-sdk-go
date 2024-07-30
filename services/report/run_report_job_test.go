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

func TestService_RunReportJob(t *testing.T) {
	type fields struct {
		wsdl            string
		networkCode     int
		applicationName string
	}
	type args struct {
		ctx   context.Context
		input report.RunReportJob
	}

	dimensionDate := report.DimensionDate

	columnAdServerImpression := report.ColumnAdServerImpressions
	columnAdServerClicks := report.ColumnAdServerClicks
	adUnitViewTopLevel := report.ReportQueryAdUnitViewTypeTopLevel
	timezomeTypePublisher := report.TimeZoneTypePublisher
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(client *soap.MockClient)
		want    *report.RunReportJobResponse
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
				input: report.RunReportJob{
					ReportJob: &report.ReportJob{
						ReportQuery: &report.ReportQuery{
							Dimensions: []*report.Dimension{&dimensionDate},
							Columns:    []*report.Column{&columnAdServerImpression, &columnAdServerClicks},
							StartDate: &report.Date{
								Year:  typeutil.IntToPtr(2024),
								Month: typeutil.IntToPtr(1),
								Day:   typeutil.IntToPtr(1),
							},
							EndDate: &report.Date{
								Year:  typeutil.IntToPtr(2024),
								Month: typeutil.IntToPtr(1),
								Day:   typeutil.IntToPtr(31),
							},
						},
					},
				},
			},
			prepare: func(client *soap.MockClient) {
				client.EXPECT().Call(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
					<soap:Header>
						<ResponseHeader xmlns="https://www.google.com/apis/ads/publisher/v202402">
							<requestId>5da2b6330833db1ab91b7ed3bb9c9f55</requestId><responseTime>593</responseTime>
						</ResponseHeader>
					</soap:Header><soap:Body>
					<runReportJobResponse xmlns="https://www.google.com/apis/ads/publisher/v202402">
						<rval><id>14595523419</id>
						<reportQuery>
							<dimensions>DATE</dimensions><adUnitView>TOP_LEVEL</adUnitView>
							<columns>AD_SERVER_IMPRESSIONS</columns><columns>AD_SERVER_CLICKS</columns>
							<startDate><year>2024</year><month>1</month><day>1</day></startDate>
							<endDate><year>2024</year><month>1</month><day>31</day></endDate>
							<timeZoneType>PUBLISHER</timeZoneType>
						</reportQuery>
						</rval>
					</runReportJobResponse></soap:Body></soap:Envelope>
				`), nil)
			},
			want: &report.RunReportJobResponse{
				Rval: &report.ReportJob{
					Id: typeutil.Int64ToPtr(14595523419),
					ReportQuery: &report.ReportQuery{
						Dimensions: []*report.Dimension{&dimensionDate},
						Columns:    []*report.Column{&columnAdServerImpression, &columnAdServerClicks},
						StartDate: &report.Date{
							Year:  typeutil.IntToPtr(2024),
							Month: typeutil.IntToPtr(1),
							Day:   typeutil.IntToPtr(1),
						},
						EndDate: &report.Date{
							Year:  typeutil.IntToPtr(2024),
							Month: typeutil.IntToPtr(1),
							Day:   typeutil.IntToPtr(31),
						},
						AdUnitView:   &adUnitViewTopLevel,
						TimeZoneType: &timezomeTypePublisher,
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
			got, err := s.RunReportJob(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunReportJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmpopts.IgnoreUnexported(report.ReportQuery{})
			if diff := cmp.Diff(got, tt.want, opts); diff != "" {
				t.Errorf("RunReportJob() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
