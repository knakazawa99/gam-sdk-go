package lineitem

import (
	"context"
	"testing"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/services/v202402/lineitem"
	"github.com/knakazawa99/gam-sdk-go/pkg/typeutil"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/requestbody"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
)

func TestService_PerformLineItemAction(t *testing.T) {
	type fields struct {
		client          soap.Client
		wsdl            string
		networkCode     int
		applicationName string
	}
	type args struct {
		ctx    context.Context
		action requestbody.PerformAction
		input  lineitem.PerformLineItemAction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(client *soap.MockClient)
		want    *lineitem.PerformLineItemActionResponse
		wantErr bool
	}{
		{
			name: "success/perform activate line item action",
			fields: fields{
				wsdl:            "https://ads.google.com/apis/ads/publisher/v202402/LineItemService?wsdl",
				networkCode:     12345678,
				applicationName: "App_1",
			},
			args: args{
				ctx:    context.Background(),
				action: lineitem.ActivateLineItems{},
				input: lineitem.PerformLineItemAction{
					FilterStatement: &lineitem.Statement{
						Query: typeutil.ConvertToPointer("WHERE id = 1"),
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
						<performLineItemActionResponse xmlns="https://www.google.com/apis/ads/publisher/v202402">
							<rval>
								<numChanges>1</numChanges>
							</rval>
						</performLineItemActionResponse>
					</soap:Body>
				</soap:Envelope>`), nil)
			},
			want: &lineitem.PerformLineItemActionResponse{
				Rval: &lineitem.UpdateResult{
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

			s := lineitem.NewService(
				mockSoapClient,
				tt.fields.networkCode,
				tt.fields.applicationName,
			)
			got, err := s.PerformLineItemAction(tt.args.ctx, tt.args.action, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("PerformLineItemAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("PerformLineItemAction() (-want +got):\\n%s", diff)
			}
		})
	}
}
