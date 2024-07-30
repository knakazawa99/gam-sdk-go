package lineitem

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"

	"github.com/knakazawa99/gam-sdk-go/pkg/gam/services/v202402/lineitem"
	"github.com/knakazawa99/gam-sdk-go/pkg/gam/soap"
	"github.com/knakazawa99/gam-sdk-go/pkg/typeutil"
)

func TestService_GetLineItemsByStatement(t *testing.T) {
	type fields struct {
		wsdl            string
		networkCode     int
		applicationName string
	}
	type args struct {
		ctx   context.Context
		input lineitem.GetLineItemsByStatement
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(client *soap.MockClient)
		want    *lineitem.GetLineItemsByStatementResponse
		wantErr error
	}{
		{
			name: "success/get companies by statement",
			fields: fields{
				wsdl:            "https://ads.google.com/apis/ads/publisher/v202402/LineItemService?wsdl",
				networkCode:     12345678,
				applicationName: "App_1",
			},
			args: args{
				ctx: context.Background(),
				input: lineitem.GetLineItemsByStatement{
					FilterStatement: &lineitem.Statement{
						Query: typeutil.ConvertToPointer("WHERE LIMIT 5"),
					},
				},
			},
			prepare: func(client *soap.MockClient) {
				client.EXPECT().Call(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
					<soap:Header><ResponseHeader xmlns="https://www.google.com/apis/ads/publisher/v202402"><requestId>995f552904060eecee1479f845d28a31</requestId><responseTime>587</responseTime></ResponseHeader></soap:Header>
					<soap:Body>
						<getLineItemsByStatementResponse xmlns="https://www.google.com/apis/ads/publisher/v202402">
							<rval>
								<totalResultSetSize>273</totalResultSetSize>
								<startIndex>0</startIndex>
								<results>
									<orderId>1000</orderId>
									<id>2000</id>
									<name>lineItem2000</name>
									<orderName>order1000</orderName>
									<startDateTime>
										<date>
											<year>2024</year>
											<month>5</month>
											<day>1</day>
										</date>
										<hour>12</hour>
										<minute>0</minute>
										<second>0</second>
										<timeZoneId>America/Los_Angeles</timeZoneId>
									</startDateTime>
									<startDateTimeType>USE_START_DATE_TIME</startDateTimeType>
									<endDateTime>
										<date>
											<year>2024</year>
											<month>5</month>
											<day>20</day>
										</date>
										<hour>12</hour>
										<minute>0</minute>
										<second>0</second>
										<timeZoneId>America/Los_Angeles</timeZoneId>
									</endDateTime>
									<autoExtensionDays>0</autoExtensionDays>
									<unlimitedEndDateTime>false</unlimitedEndDateTime>
									<creativeRotationType>MANUAL</creativeRotationType>
									<deliveryRateType>EVENLY</deliveryRateType>
									<deliveryForecastSource>HISTORICAL</deliveryForecastSource>
									<roadblockingType>ONE_OR_MORE</roadblockingType>
									<skippableAdType>DISABLED</skippableAdType>
									<lineItemType>STANDARD</lineItemType>
									<priority>8</priority>
									<costPerUnit>
										<currencyCode>USD</currencyCode>
										<microAmount>0</microAmount>
									</costPerUnit>
									<valueCostPerUnit>
										<currencyCode>USD</currencyCode>
										<microAmount>0</microAmount>
									</valueCostPerUnit>
									<costType>CPM</costType>
									<discountType>PERCENTAGE</discountType>
									<discount>0.0</discount>
									<contractedUnitsBought>0</contractedUnitsBought>
									<creativePlaceholders>
										<size>
											<width>1</width>
											<height>1</height>
											<isAspectRatio>false</isAspectRatio>
										</size>
										<expectedCreativeCount>1</expectedCreativeCount>
										<creativeSizeType>PIXEL</creativeSizeType>
										<isAmpOnly>false</isAmpOnly>
									</creativePlaceholders>
									<environmentType>BROWSER</environmentType>
									<companionDeliveryOption>UNKNOWN</companionDeliveryOption>
									<allowOverbook>false</allowOverbook>
									<skipInventoryCheck>false</skipInventoryCheck>
									<skipCrossSellingRuleWarningChecks>false</skipCrossSellingRuleWarningChecks>
									<reserveAtCreation>false</reserveAtCreation>
									<stats>
										<impressionsDelivered>0</impressionsDelivered>
										<clicksDelivered>0</clicksDelivered>
										<videoCompletionsDelivered>0</videoCompletionsDelivered>
										<videoStartsDelivered>0</videoStartsDelivered>
										<viewableImpressionsDelivered>0</viewableImpressionsDelivered>
									</stats>
									<deliveryIndicator>
										<expectedDeliveryPercentage>100.0</expectedDeliveryPercentage>
										<actualDeliveryPercentage>0.0</actualDeliveryPercentage>
									</deliveryIndicator>
									<budget>
										<currencyCode>USD</currencyCode>
										<microAmount>100000000</microAmount>
									</budget>
									<status>COMPLETED</status>
									<reservationStatus>UNRESERVED</reservationStatus>
									<isArchived>false</isArchived>
									<webPropertyCode></webPropertyCode>
									<disableSameAdvertiserCompetitiveExclusion>false</disableSameAdvertiserCompetitiveExclusion>
									<lastModifiedByApp>Goog_DFPUI</lastModifiedByApp>
									<competitiveConstraintScope>POD</competitiveConstraintScope>
									<lastModifiedDateTime>
										<date>
											<year>2024</year>
											<month>4</month>
											<day>30</day>
										</date>
										<hour>12</hour>
										<minute>0</minute>
										<second>0</second>
										<timeZoneId>America/Los_Angeles</timeZoneId>
									</lastModifiedDateTime>
									<creationDateTime>
										<date>
											<year>2024</year>
											<month>4</month>
											<day>29</day>
										</date>
										<hour>12</hour>
										<minute>0</minute>
										<second>0</second>
										<timeZoneId>America/Los_Angeles</timeZoneId>
									</creationDateTime>
									<isMissingCreatives>true</isMissingCreatives>
									<youtubeKidsRestricted>false</youtubeKidsRestricted>
									<videoMaxDuration>0</videoMaxDuration>
									<primaryGoal>
										<goalType>LIFETIME</goalType>
										<unitType>IMPRESSIONS</unitType>
										<units>1</units>
									</primaryGoal>
									<childContentEligibility>DISALLOWED</childContentEligibility>
									<targeting>
										<inventoryTargeting>
											<targetedPlacementIds>3000</targetedPlacementIds>
										</inventoryTargeting>
										<requestPlatformTargeting></requestPlatformTargeting>
									</targeting>
								</results>
							</rval>
						</getLineItemsByStatementResponse>
					</soap:Body></soap:Envelope>`), nil)
			},
			want: &lineitem.GetLineItemsByStatementResponse{
				Rval: &lineitem.LineItemPage{
					TotalResultSetSize: typeutil.ConvertToPointer(273),
					StartIndex:         typeutil.ConvertToPointer(0),
					Results: []*lineitem.LineItem{
						{
							Targeting: &lineitem.Targeting{
								InventoryTargeting: &lineitem.InventoryTargeting{
									TargetedPlacementIds: []*int64{typeutil.Int64ToPtr(3000)},
								},
								RequestPlatformTargeting: &lineitem.RequestPlatformTargeting{},
							},
							LineItemSummary: &lineitem.LineItemSummary{
								OrderId:   typeutil.Int64ToPtr(1000),
								Id:        typeutil.Int64ToPtr(2000),
								Name:      typeutil.StringToPtr("lineItem2000"),
								OrderName: typeutil.StringToPtr("order1000"),
								StartDateTime: &lineitem.DateTime{
									Date: &lineitem.Date{
										Year:  typeutil.IntToPtr(2024),
										Month: typeutil.IntToPtr(5),
										Day:   typeutil.IntToPtr(1),
									},
									Hour:       typeutil.IntToPtr(12),
									Minute:     typeutil.IntToPtr(0),
									Second:     typeutil.IntToPtr(0),
									TimeZoneId: typeutil.StringToPtr("America/Los_Angeles"),
								},
								StartDateTimeType: typeutil.ConvertToPointer(lineitem.StartDateTimeTypeUseStartDateTime),
								EndDateTime: &lineitem.DateTime{
									Date: &lineitem.Date{
										Year:  typeutil.IntToPtr(2024),
										Month: typeutil.IntToPtr(5),
										Day:   typeutil.IntToPtr(20),
									},
									Hour:       typeutil.IntToPtr(12),
									Minute:     typeutil.IntToPtr(0),
									Second:     typeutil.IntToPtr(0),
									TimeZoneId: typeutil.StringToPtr("America/Los_Angeles"),
								},
								AutoExtensionDays:      typeutil.IntToPtr(0),
								UnlimitedEndDateTime:   typeutil.ConvertToPointer(false),
								CreativeRotationType:   typeutil.ConvertToPointer(lineitem.CreativeRotationTypeManual),
								DeliveryRateType:       typeutil.ConvertToPointer(lineitem.DeliveryRateTypeEvenly),
								DeliveryForecastSource: typeutil.ConvertToPointer(lineitem.DeliveryForecastSourceHistorical),
								RoadblockingType:       typeutil.ConvertToPointer(lineitem.RoadblockingTypeOneOrMore),
								SkippableAdType:        typeutil.ConvertToPointer(lineitem.SkippableAdTypeDisabled),
								LineItemType:           typeutil.ConvertToPointer(lineitem.LineItemTypeStandard),
								Priority:               typeutil.IntToPtr(8),
								CostPerUnit: &lineitem.Money{
									CurrencyCode: typeutil.StringToPtr("USD"),
									MicroAmount:  typeutil.Int64ToPtr(0),
								},
								ValueCostPerUnit: &lineitem.Money{
									CurrencyCode: typeutil.StringToPtr("USD"),
									MicroAmount:  typeutil.Int64ToPtr(0),
								},
								CostType:              typeutil.ConvertToPointer(lineitem.CostTypeCpm),
								DiscountType:          typeutil.ConvertToPointer(lineitem.LineItemDiscountTypePercentage),
								Discount:              typeutil.ConvertToPointer(0.0),
								ContractedUnitsBought: typeutil.Int64ToPtr(0),
								CreativePlaceholders: []*lineitem.CreativePlaceholder{
									{
										Size: &lineitem.Size{
											Width:         typeutil.IntToPtr(1),
											Height:        typeutil.IntToPtr(1),
											IsAspectRatio: typeutil.ConvertToPointer(false),
										},
										ExpectedCreativeCount: typeutil.IntToPtr(1),
										CreativeSizeType:      typeutil.ConvertToPointer(lineitem.CreativeSizeTypePixel),
										IsAmpOnly:             typeutil.ConvertToPointer(false),
									},
								},
								EnvironmentType:                   typeutil.ConvertToPointer(lineitem.EnvironmentTypeBrowser),
								CompanionDeliveryOption:           typeutil.ConvertToPointer(lineitem.CompanionDeliveryOptionUnknown),
								AllowOverbook:                     typeutil.ConvertToPointer(false),
								SkipInventoryCheck:                typeutil.ConvertToPointer(false),
								SkipCrossSellingRuleWarningChecks: typeutil.ConvertToPointer(false),
								ReserveAtCreation:                 typeutil.ConvertToPointer(false),
								Stats: &lineitem.Stats{
									ImpressionsDelivered:         typeutil.Int64ToPtr(0),
									ClicksDelivered:              typeutil.Int64ToPtr(0),
									VideoCompletionsDelivered:    typeutil.Int64ToPtr(0),
									VideoStartsDelivered:         typeutil.Int64ToPtr(0),
									ViewableImpressionsDelivered: typeutil.Int64ToPtr(0),
								},
								DeliveryIndicator: &lineitem.DeliveryIndicator{
									ExpectedDeliveryPercentage: typeutil.ConvertToPointer(100.0),
									ActualDeliveryPercentage:   typeutil.ConvertToPointer(0.0),
								},
								Budget: &lineitem.Money{
									CurrencyCode: typeutil.StringToPtr("USD"),
									MicroAmount:  typeutil.Int64ToPtr(100000000),
								},
								Status:            typeutil.ConvertToPointer(lineitem.ComputedStatusCompleted),
								ReservationStatus: typeutil.ConvertToPointer(lineitem.LineItemSummaryReservationStatusTypeUnreserved),
								IsArchived:        typeutil.ConvertToPointer(false),
								WebPropertyCode:   typeutil.StringToPtr(""),
								DisableSameAdvertiserCompetitiveExclusion: typeutil.ConvertToPointer(false),
								LastModifiedByApp:                         typeutil.StringToPtr("Goog_DFPUI"),
								CompetitiveConstraintScope:                typeutil.ConvertToPointer(lineitem.CompetitiveConstraintScopePod),
								LastModifiedDateTime: &lineitem.DateTime{
									Date: &lineitem.Date{
										Year:  typeutil.IntToPtr(2024),
										Month: typeutil.IntToPtr(4),
										Day:   typeutil.IntToPtr(30),
									},
									Hour:       typeutil.IntToPtr(12),
									Minute:     typeutil.IntToPtr(0),
									Second:     typeutil.IntToPtr(0),
									TimeZoneId: typeutil.StringToPtr("America/Los_Angeles"),
								},
								CreationDateTime: &lineitem.DateTime{
									Date: &lineitem.Date{
										Year:  typeutil.IntToPtr(2024),
										Month: typeutil.IntToPtr(4),
										Day:   typeutil.IntToPtr(29),
									},
									Hour:       typeutil.IntToPtr(12),
									Minute:     typeutil.IntToPtr(0),
									Second:     typeutil.IntToPtr(0),
									TimeZoneId: typeutil.StringToPtr("America/Los_Angeles"),
								},
								IsMissingCreatives:    typeutil.ConvertToPointer(true),
								YoutubeKidsRestricted: typeutil.ConvertToPointer(false),
								VideoMaxDuration:      typeutil.Int64ToPtr(0),
								PrimaryGoal: &lineitem.Goal{
									GoalType: typeutil.ConvertToPointer(lineitem.GoalTypeLifetime),
									UnitType: typeutil.ConvertToPointer(lineitem.UnitTypeImpressions),
									Units:    typeutil.Int64ToPtr(1),
								},
								ChildContentEligibility: typeutil.ConvertToPointer(lineitem.ChildContentEligibilityDisallowed),
							},
						},
					},
				},
			},
		},
		{
			name: "fail/query statement ",
			fields: fields{
				wsdl:            "https://ads.google.com/apis/ads/publisher/v202402/OrderService?wsdl",
				networkCode:     12345678,
				applicationName: "App_1",
			},
			args: args{
				ctx: context.Background(),
				input: lineitem.GetLineItemsByStatement{
					FilterStatement: &lineitem.Statement{
						Query: typeutil.ConvertToPointer("WHERE statusss = 'APPROVED'"),
					},
				},
			},
			prepare: func(client *soap.MockClient) {
				client.EXPECT().Call(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
					<soap:Header>
						<ResponseHeader xmlns="https://www.google.com/apis/ads/publisher/v202402">
							<requestId>0e552db90bfffafd40d6f2ecd1a8d81b</requestId>
							<responseTime>264</responseTime>
						</ResponseHeader>
					</soap:Header>
					<soap:Body>
						<soap:Fault>
							<faultcode>soap:Server</faultcode>
							<faultstring>[PublisherQueryLanguageContextError.UNEXECUTABLE @ ]</faultstring>
							<detail>
								<ApiExceptionFault xmlns="https://www.google.com/apis/ads/publisher/v202402">
									<message>[PublisherQueryLanguageContextError.UNEXECUTABLE @ ]</message>
									<errors xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
											xsi:type="PublisherQueryLanguageContextError">
										<fieldPath></fieldPath>
										<trigger></trigger>
										<errorString>PublisherQueryLanguageContextError.UNEXECUTABLE</errorString>
										<reason>UNEXECUTABLE</reason>
									</errors>
								</ApiExceptionFault>
							</detail>
						</soap:Fault>
					</soap:Body>
				</soap:Envelope>
				`), nil)
			},
			want:    nil,
			wantErr: lineitem.NewPublisherQueryLanguageContextError("UNEXECUTABLE"),
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
			got, err := s.GetLineItemsByStatement(tt.args.ctx, tt.args.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetLineItemsByStatement() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetLineItemsByStatement() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
