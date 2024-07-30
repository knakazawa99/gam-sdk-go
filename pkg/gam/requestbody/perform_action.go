package requestbody

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam"
)

var soapPerformActionTemplate = fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope
     xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
     xmlns:xsd="http://www.w3.org/2001/XMLSchema"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xmlns:ns="https://www.google.com/apis/ads/publisher/v202402">
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
   <{{.ActionName}} xsi:type="ns:{{.ActionType}}">{{.ActionField}}</{{.ActionName}}>
   <{{.Statement}}>
     {{ .StatementParameter }}
   </{{.Statement}}>
 </{{.OperationName}}>
</soapenv:Body>
</soapenv:Envelope>
`, gam.Namespace, gam.Namespace)

func getSoapPerformActionTemplate() *template.Template {
	return template.Must(template.New("soapBody").Parse(soapPerformActionTemplate))
}

type PerformActionBodyData struct {
	NetworkCode        int
	ApplicationName    string
	OperationName      string
	ActionName         string
	ActionType         string
	ActionField        string
	Statement          string
	StatementParameter string
}

func GeneratePerformActionSoapBody(data PerformActionBodyData) (string, error) {
	var buf bytes.Buffer
	if err := getSoapPerformActionTemplate().Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

type PerformAction interface {
	GetPerformActionBody() (string, error)
	GetActionType() string
}
