package requestbody

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam"
)

var soapDefaultRequestBodyTemplate = fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope
        xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
        xmlns:xsd="http://www.w3.org/2001/XMLSchema"
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <soapenv:Header>
    <ns1:RequestHeader
         soapenv:actor="http://schemas.xmlsoap.org/soap/actor/next"
         soapenv:mustUnderstand="0"
         xmlns:ns1="%s">
      <ns1:networkCode>{{.NetworkCode}}</ns1:networkCode>
      <ns1:applicationName>{{.ApplicationName}}</ns1:applicationName>
    </ns1:RequestHeader>
  </soapenv:Header>
  <soapenv:Body>
    <{{.OperationName}} xmlns="%s">
		{{ .Parameters }}
    </{{.OperationName}}>
  </soapenv:Body>
</soapenv:Envelope>
`, gam.Namespace, gam.Namespace)

func getSoapDefaultBodyTemplate() *template.Template {
	return template.Must(template.New("soapBody").Parse(soapDefaultRequestBodyTemplate))
}

type DefaultRequestBody struct {
	NetworkCode     int
	ApplicationName string
	OperationName   string
	Parameters      string
}

func GenerateDefaultRequestSoapBody(data DefaultRequestBody) (string, error) {
	var buf bytes.Buffer
	if err := getSoapDefaultBodyTemplate().Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
