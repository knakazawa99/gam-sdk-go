package order

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/mock/gomock"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/services/v202402/order"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
	"github.com/knakazawa99/gam-sdk-go/pkg/typeutil"
)

func TestService_CreateOrders(t *testing.T) {
	type fields struct {
		wsdl            string
		networkCode     int
		applicationName string
	}
	type args struct {
		ctx    context.Context
		orders order.CreateOrders
	}

	orderStatusDraft := order.OrderStatusDraft
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(client *soap.MockClient)
		want    *order.CreateOrdersResponse
		wantErr error
	}{
		{
			name: "success/create orders",
			fields: fields{
				wsdl:            "https://ads.google.com/apis/ads/publisher/v202402/OrderService?wsdl",
				networkCode:     12345678,
				applicationName: "App_1",
			},
			args: args{
				ctx: context.Background(),
				orders: order.CreateOrders{
					Orders: []*order.Order{
						{
							Name:          typeutil.StringToPtr("order-2024042609"),
							TraffickerId:  typeutil.Int64ToPtr(123456789),
							SalespersonId: typeutil.Int64ToPtr(123456789),
							StartDateTime: &order.DateTime{
								Date: &order.Date{
									Year:  typeutil.IntToPtr(2022),
									Month: typeutil.IntToPtr(4),
									Day:   typeutil.IntToPtr(28),
								},
								Hour:       typeutil.IntToPtr(12),
								Minute:     typeutil.IntToPtr(0),
								Second:     typeutil.IntToPtr(0),
								TimeZoneId: typeutil.StringToPtr("UTC"),
							},
							EndDateTime: &order.DateTime{
								Date: &order.Date{
									Year:  typeutil.IntToPtr(2022),
									Month: typeutil.IntToPtr(4),
									Day:   typeutil.IntToPtr(30),
								},
								Hour:       typeutil.IntToPtr(12),
								Minute:     typeutil.IntToPtr(0),
								Second:     typeutil.IntToPtr(0),
								TimeZoneId: typeutil.StringToPtr("UTC"),
							},
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
						<createOrdersResponse xmlns="https://www.google.com/apis/ads/publisher/v202402">
							<rval>
								<id>123456789</id>
								<advertiserId>123456789</advertiserId>
								<name>order-2024042609</name>
								<status>DRAFT</status>
								<traffickerId>123456789</traffickerId>
								<salespersonId>123456789</salespersonId>
								<startDateTime>
									<date>
										<year>2022</year>
										<month>4</month>	
										<day>28</day>
									</date>
									<hour>12</hour>	
									<minute>0</minute>
									<second>0</second>	
									<timeZoneId>UTC</timeZoneId>
								</startDateTime>
								<endDateTime>
									<date>
										<year>2022</year>
										<month>4</month>	
										<day>30</day>
									</date>
									<hour>12</hour>	
									<minute>0</minute>
									<second>0</second>	
									<timeZoneId>UTC</timeZoneId>
								</endDateTime>
								<lastModifiedDateTime>
									<date>
										<year>2024</year>
										<month>4</month>
										<day>27</day>
									</date>
									<hour>12</hour>
									<minute>00</minute>
									<second>00</second>
									<timeZoneId>UTC</timeZoneId>
								</lastModifiedDateTime>
							</rval>
						</createOrdersResponse>
					</soap:Body>
				</soap:Envelope>
				`), nil)
			},
			want: &order.CreateOrdersResponse{
				Rval: []*order.Order{
					{
						Id:            typeutil.Int64ToPtr(123456789),
						Name:          typeutil.StringToPtr("order-2024042609"),
						Status:        &orderStatusDraft,
						AdvertiserId:  typeutil.Int64ToPtr(123456789),
						TraffickerId:  typeutil.Int64ToPtr(123456789),
						SalespersonId: typeutil.Int64ToPtr(123456789),
						StartDateTime: &order.DateTime{
							Date: &order.Date{
								Year:  typeutil.IntToPtr(2022),
								Month: typeutil.IntToPtr(4),
								Day:   typeutil.IntToPtr(28),
							},
							Hour:       typeutil.IntToPtr(12),
							Minute:     typeutil.IntToPtr(0),
							Second:     typeutil.IntToPtr(0),
							TimeZoneId: typeutil.StringToPtr("UTC"),
						},
						EndDateTime: &order.DateTime{
							Date: &order.Date{
								Year:  typeutil.IntToPtr(2022),
								Month: typeutil.IntToPtr(4),
								Day:   typeutil.IntToPtr(30),
							},
							Hour:       typeutil.IntToPtr(12),
							Minute:     typeutil.IntToPtr(0),
							Second:     typeutil.IntToPtr(0),
							TimeZoneId: typeutil.StringToPtr("UTC"),
						},
						LastModifiedDateTime: &order.DateTime{
							Date: &order.Date{
								Year:  typeutil.IntToPtr(2024),
								Month: typeutil.IntToPtr(4),
								Day:   typeutil.IntToPtr(27),
							},
							Hour:       typeutil.IntToPtr(12),
							Minute:     typeutil.IntToPtr(0),
							Second:     typeutil.IntToPtr(0),
							TimeZoneId: typeutil.StringToPtr("UTC"),
						},
					},
				},
			},
		},
		{
			name: "fail/not unique order name",
			fields: fields{
				wsdl:            "https://ads.google.com/apis/ads/publisher/v202402/OrderService?wsdl",
				networkCode:     12345678,
				applicationName: "App_1",
			},
			args: args{
				ctx: context.Background(),
				orders: order.CreateOrders{
					Orders: []*order.Order{
						{
							Name: typeutil.StringToPtr("order-2024042609"),
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
								<faultstring>[UniqueError.NOT_UNIQUE @ [0].name; trigger:'order-2024042609']</faultstring>
								<detail>
									<ApiExceptionFault xmlns="https://www.google.com/apis/ads/publisher/v202402">
										<message>[UniqueError.NOT_UNIQUE @ [0].name; trigger:'order-2024042609']</message>
										<errors xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="UniqueError">
											<fieldPath>[0].name</fieldPath>
											<fieldPathElements>
												<index>0</index>
											</fieldPathElements>
											<fieldPathElements>
												<field>name</field>
											</fieldPathElements>
											<trigger>order-2024042609</trigger>
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
			wantErr: order.NewUniqueError("order-2024042609"),
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
			got, err := s.CreateOrders(tt.args.ctx, tt.args.orders)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("CreateOrders() error = %v, wantErr %v", err, tt.wantErr)
			}
			opts := cmpopts.IgnoreUnexported(order.Order{})
			if diff := cmp.Diff(tt.want, got, opts); diff != "" {
				t.Errorf("CreateOrders() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
