package soap

import (
	"bytes"
	"context"
	"encoding/xml"
	"io"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
)

type Client interface {
	Call(ctx context.Context, url string, soapBody string) ([]byte, error)
}
type client struct {
	Token  *oauth2.Token
	Config *jwt.Config
}

func NewClient(token *oauth2.Token, cfg *jwt.Config) Client {
	return &client{
		Token:  token,
		Config: cfg,
	}
}

func (c client) Call(ctx context.Context, url string, soapBody string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(soapBody)))

	if err != nil {
		return nil, err
	}

	// Refresh token if it's not valid
	if !c.Token.Valid() {
		token, err := c.Config.TokenSource(ctx).Token()
		if err != nil {
			return nil, err
		}
		c.Token = token
	}

	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("Authorization", "Bearer "+c.Token.AccessToken)

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Fault type
type Fault struct {
	XMLName xml.Name `xml:"Fault"`

	Code        string      `xml:"faultcode,omitempty"`
	FaultString string      `xml:"faultstring,omitempty"`
	Detail      FaultDetail `xml:"detail,omitempty"`
}

type FaultDetail struct {
	XMLName           xml.Name           `xml:"detail"`
	APIExecutionFault *APIExecutionFault `xml:"ApiExceptionFault"`
}

type APIExecutionFault struct {
	XMLName xml.Name `xml:"ApiExceptionFault"`
	Errors  []Errors `xml:"errors,omitempty"`
	Message string   `xml:"message,omitempty"`
}

type Errors struct {
	XMLName     xml.Name `xml:"errors"`
	ErrorType   string   `xml:"type,attr"`
	FieldPath   string   `xml:"fieldPath,omitempty"`
	ErrorString string   `xml:"errorString,omitempty"`
	Reason      string   `xml:"reason,omitempty"`
	Trigger     string   `xml:"trigger,omitempty"`
}

type UpdateResult struct {
	NumChanges int `xml:"numChanges"`
}
