package order

import (
	"context"
	"testing"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/services/v202402/order"
	"github.com/knakazawa99/gam-sdk-go/pkg/typeutil"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/requestbody"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
)

func TestService_PerformOrderAction(t *testing.T) {
	type fields struct {
		client          soap.Client
		wsdl            string
		networkCode     int
		applicationName string
	}
	type args struct {
		ctx    context.Context
		action requestbody.PerformAction
		input  order.PerformOrderAction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(client *soap.MockClient)
		want    *order.PerformOrderActionResponse
		wantErr bool
	}{
		{
			name: "success/perform approve-order-action",
			fields: fields{
				wsdl:            "https://ads.google.com/apis/ads/publisher/v202402/CompanyService?wsdl",
				networkCode:     12345678,
				applicationName: "App_1",
			},
			args: args{
				ctx: context.Background(),
				action: order.ApproveOrders{
					SkipInventoryCheck: typeutil.ConvertToPointer(true),
				},
				input: order.PerformOrderAction{
					FilterStatement: &order.Statement{Query: typeutil.ConvertToPointer("WHERE id = 1")},
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
						<performOrderActionResponse xmlns="https://www.google.com/apis/ads/publisher/v202402">
							<rval>
								<numChanges>1</numChanges>
							</rval>
						</performOrderActionResponse>
					</soap:Body>
				</soap:Envelope>`), nil)
			},
			want: &order.PerformOrderActionResponse{
				Rval: &order.UpdateResult{
					NumChanges: typeutil.ConvertToPointer(1),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockSoapClient := soap.NewMockClient(ctrl)
			tt.prepare(mockSoapClient)

			s := order.NewService(
				mockSoapClient,
				tt.fields.networkCode,
				tt.fields.applicationName,
			)
			got, err := s.PerformOrderAction(tt.args.ctx, tt.args.action, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("PerformOrderAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("PerformOrderAction() (-want +got):\\n%s", diff)
			}
		})
	}
}
