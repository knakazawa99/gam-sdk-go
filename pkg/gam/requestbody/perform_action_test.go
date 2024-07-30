package requestbody

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGeneratePerformActionSoapBody(t *testing.T) {
	type args struct {
		data PerformActionBodyData
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				data: PerformActionBodyData{
					NetworkCode:        123456,
					ApplicationName:    "app",
					OperationName:      "performCompanyAction",
					ActionName:         "companyAction",
					ActionType:         "EndAgreementAction",
					ActionField:        "",
					Statement:          "statement",
					StatementParameter: "<query>WHERE id = 1</query>",
				},
			},
			want: `<?xml version="1.0" encoding="UTF-8"?>
				<soapenv:Envelope
					 xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
					 xmlns:xsd="http://www.w3.org/2001/XMLSchema"
					 xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
					xmlns:ns="https://www.google.com/apis/ads/publisher/v202402">
				<soapenv:Header>
				 <ns1:RequestHeader
					  soapenv:actor="http://schemas.xmlsoap.org/soap/actor/next"
					  soapenv:mustUnderstand="0"
					  xmlns:ns1="https://www.google.com/apis/ads/publisher/v202402">
				   <ns1:networkCode>123456</ns1:networkCode>
				   <ns1:applicationName>app</ns1:applicationName>
				 </ns1:RequestHeader>
				</soapenv:Header>
				<soapenv:Body>
				 <performCompanyAction xmlns="https://www.google.com/apis/ads/publisher/v202402">
				   <companyAction xsi:type="ns:EndAgreementAction"></companyAction>
				   <statement>
					 <query>WHERE id = 1</query>
				   </statement>
				 </performCompanyAction>
				</soapenv:Body>
				</soapenv:Envelope>
			`,
		},
	}
	trimWordGaps := func(s string) string {
		return strings.Join(strings.Fields(s), " ")
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GeneratePerformActionSoapBody(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePerformActionSoapBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(trimWordGaps(tt.want), trimWordGaps(got)); diff != "" {
				t.Errorf("GeneratePerformActionSoapBody() (-want +got):\\n%s", diff)
			}
		})
	}
}
