package main

import (
	"context"
	"fmt"
	"os"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/services/v202402/company"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
	"github.com/knakazawa99/gam-sdk-go/pkg/typeutil"
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

	companyTypeAdvertiser := company.CompanyTypeTypeAdvertiser
	res, err := companyService.CreateCompanies(ctx, company.CreateCompanies{
		Companies: []*company.Company{
			{
				Name: typeutil.StringToPtr("company-2024042605"),
				Type: &companyTypeAdvertiser,
			},
			{
				Name: typeutil.StringToPtr("company-2024042606"),
				Type: &companyTypeAdvertiser,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(*res.Rval[0].Name)
}
