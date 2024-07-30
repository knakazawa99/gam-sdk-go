package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/services/v202402/report"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
	"github.com/knakazawa99/gam-sdk-go/pkg/typeutil"
)

func main() {
	privateKeyPemString := "-----BEGIN PRIVATE KEY----END PRIVATE KEY-----"
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
	// Get Access Token
	token, err := gam.GetAccessToken(ctx, cfg)
	if err != nil {
		fmt.Println("Failed to get token:", err)
		os.Exit(1)
	}

	soapClient := soap.NewClient(token, cfg)
	reportService := report.NewService(soapClient, networkCode, applicationName)

	// Run Report Job
	dimensionDate := report.DimensionDate
	columnAdServerImpression := report.ColumnAdServerImpressions
	columnAdServerClicks := report.ColumnAdServerClicks
	columnAdServerCTR := report.ColumnAdServerCtr
	reportJob := report.ReportJob{
		ReportQuery: &report.ReportQuery{
			Dimensions: []*report.Dimension{&dimensionDate},
			Columns:    []*report.Column{&columnAdServerImpression, &columnAdServerClicks, &columnAdServerCTR},
			StartDate: &report.Date{
				Year:  typeutil.IntToPtr(2024),
				Month: typeutil.IntToPtr(1),
				Day:   typeutil.IntToPtr(1),
			},
			EndDate: &report.Date{
				Year:  typeutil.IntToPtr(2024),
				Month: typeutil.IntToPtr(1),
				Day:   typeutil.IntToPtr(31),
			},
		},
	}
	reportJobResult, err := reportService.RunReportJob(ctx, report.RunReportJob{ReportJob: &reportJob})
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}

	reportJobID := reportJobResult.Rval.Id
	fmt.Println("reportJobResult.Id", reportJobID)
	// Polling Report Job Status
	for {
		res, err := reportService.GetReportJobStatus(ctx, report.GetReportJobStatus{ReportJobId: reportJobID})
		reportJobStatus := res.Rval
		if err != nil {
			fmt.Println("GetReportJobStatus", err)
			os.Exit(1)
		}
		if *reportJobStatus == report.ReportJobStatusCompleted {
			fmt.Println("reportJobStatus: ", report.ReportJobStatusCompleted)
			break
		} else if *reportJobStatus == report.ReportJobStatusFailed {
			fmt.Println("reportJobStatus: ", report.ReportJobStatusFailed)
			os.Exit(1)
		} else if *reportJobStatus == report.ReportJobStatusInProgress {
			fmt.Println("reportJobStatus: ", report.ReportJobStatusInProgress)
		}
	}

	// Get Report Download URL
	resReportDownloadURL, err := reportService.GetReportDownloadURL(ctx, report.GetReportDownloadURL{
		ReportJobId:  reportJobID,
		ExportFormat: typeutil.ConvertToPointer(report.ExportFormatCsvDump),
	})
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}

	resp, err := http.Get(*resReportDownloadURL.Rval)
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	out, err := os.Create("./report.csv.gz")
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}
}
