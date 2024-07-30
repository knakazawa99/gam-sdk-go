package main

import (
	"context"
	"fmt"
	"os"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/services/v202402/company"
	"github.com/knakazawa99/gam-sdk-go/pkg/typeutil"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
)

func main() {
	privateKeyPemString := ""
	privateKeyID := "private-key-id"
	serviceAccountID := "service-account-id"
	networkCode := 12345678
	applicationName := "app_1"

	ctx := context.Background()

	cfg, _ := gam.GetConfig(ctx, gam.AccessTokenInput{
		PrivateKeyPemString: &privateKeyPemString,
		PrivateKeyID:        privateKeyID,
		ServiceAccountID:    serviceAccountID,
	})
	token, err := gam.GetAccessToken(ctx, cfg)
	if err != nil {
		fmt.Println("Failed to get token:", err)
		os.Exit(1)
	}

	soapClient := soap.NewClient(token, cfg)
	companyService := company.NewService(soapClient, networkCode, applicationName)

	response, err := companyService.GetCompaniesByStatement(ctx,
		company.GetCompaniesByStatement{
			FilterStatement: &company.Statement{Query: typeutil.ConvertToPointer("WHERE type = 'AGENCY'")},
		},
	)

	if err != nil {
		fmt.Println("Failed to get companies:", err)
		os.Exit(1)
	}

	for _, result := range response.Rval.Results {
		fmt.Println("result: ", result.Name)
		fmt.Println("result: ", result.Type)
	}

}
