package company

import (
	"context"
	"testing"

	"github.com/knakazawa99/gam-sdk-go/pkg/typeutil"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/requestbody"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/services/v202402/company"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
)

func TestService_PerformCompanyAction(t *testing.T) {
	type fields struct {
		client          soap.Client
		wsdl            string
		networkCode     int
		applicationName string
	}
	type args struct {
		ctx       context.Context
		action    requestbody.PerformAction
		statement company.PerformCompanyAction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(client *soap.MockClient)
		want    *company.PerformCompanyActionResponse
		wantErr bool
	}{
		{
			name: "success/perform re-invite-action",
			fields: fields{
				wsdl:            "https://ads.google.com/apis/ads/publisher/v202402/CompanyService?wsdl",
				networkCode:     12345678,
				applicationName: "App_1",
			},
			args: args{
				ctx: context.Background(),
				action: company.ReInviteAction{
					ProposedDelegationType:           typeutil.ConvertToPointer(company.DelegationTypeManageAccount),
					ProposedRevenueShareMillipercent: typeutil.ConvertToPointer(int64(1000)),
					ProposedEmail:                    typeutil.ConvertToPointer("example@example.com"),
				},
				statement: company.PerformCompanyAction{
					Statement: &company.Statement{
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
						<performCompanyActionResponse xmlns="https://www.google.com/apis/ads/publisher/v202402">
							<rval>
								<numChanges>1</numChanges>
							</rval>
						</performCompanyActionResponse>
					</soap:Body>
				</soap:Envelope>`), nil)
			},
			want: &company.PerformCompanyActionResponse{
				Rval: &company.UpdateResult{
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

			s := company.NewService(
				mockSoapClient,
				tt.fields.networkCode,
				tt.fields.applicationName,
			)
			got, err := s.PerformCompanyAction(tt.args.ctx, tt.args.action, tt.args.statement)
			if (err != nil) != tt.wantErr {
				t.Errorf("PerformCompanyAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("PerformCompanyAction() (-want +got):\\n%s", diff)
			}
		})
	}
}

func TestAction_GetPerformActionBody(t *testing.T) {
	tests := []struct {
		name    string
		action  requestbody.PerformAction
		want    string
		wantErr bool
	}{

		{
			name: "success/perform re-invite-action",
			action: company.ReInviteAction{
				ProposedDelegationType:           typeutil.ConvertToPointer(company.DelegationTypeManageAccount),
				ProposedRevenueShareMillipercent: typeutil.ConvertToPointer(int64(1000)),
				ProposedEmail:                    typeutil.ConvertToPointer("example@example.com"),
			},
			want: "<proposedDelegationType>MANAGE_ACCOUNT</proposedDelegationType><proposedRevenueShareMillipercent>1000</proposedRevenueShareMillipercent><proposedEmail>example@example.com</proposedEmail>",
		},
		{
			name:   "success/perform end-agreement-action",
			action: company.EndAgreementAction{},
			want:   "",
		},
		{
			name:   "success/perform resend-invitation-action",
			action: company.ResendInvitationAction{},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.action.GetPerformActionBody()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPerformActionBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("GetPerformActionBody() (-want +got):\\n%s", diff)
			}
		})
	}
}
