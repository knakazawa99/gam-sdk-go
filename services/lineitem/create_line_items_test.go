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

func TestService_CreateCompanies(t *testing.T) {
	type fields struct {
		wsdl            string
		networkCode     int
		applicationName string
	}
	type args struct {
		ctx       context.Context
		lineItems lineitem.CreateLineItems
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(client *soap.MockClient)
		want    *lineitem.CreateLineItemsResponse
		wantErr error
	}{
		{
			name: "success/create line items",
			fields: fields{
				wsdl:            "https://ads.google.com/apis/ads/publisher/v202402/LineItemService?wsdl",
				networkCode:     12345678,
				applicationName: "App_1",
			},
			args: args{
				ctx: context.Background(),
				lineItems: lineitem.CreateLineItems{
					LineItems: []*lineitem.LineItem{
						{
							LineItemSummary: &lineitem.LineItemSummary{
								OrderId:           typeutil.Int64ToPtr(1000),
								Name:              typeutil.StringToPtr("lineItem2000"),
								StartDateTimeType: typeutil.ConvertToPointer(lineitem.StartDateTimeTypeImmediately),
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
								CreativeRotationType: typeutil.ConvertToPointer(lineitem.CreativeRotationTypeManual),
								DeliveryRateType:     typeutil.ConvertToPointer(lineitem.DeliveryRateTypeAsFastAsPossible),
								LineItemType:         typeutil.ConvertToPointer(lineitem.LineItemTypeStandard),
								Priority:             typeutil.IntToPtr(8),
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
								EnvironmentType:         typeutil.ConvertToPointer(lineitem.EnvironmentTypeBrowser),
								CompanionDeliveryOption: typeutil.ConvertToPointer(lineitem.CompanionDeliveryOptionUnknown),
								AllowOverbook:           typeutil.ConvertToPointer(true),
								SkipInventoryCheck:      typeutil.ConvertToPointer(true),
								ReserveAtCreation:       typeutil.ConvertToPointer(false),
								Budget: &lineitem.Money{
									CurrencyCode: typeutil.StringToPtr("USD"),
									MicroAmount:  typeutil.Int64ToPtr(100000000),
								},
								WebPropertyCode: typeutil.StringToPtr(""),
								DisableSameAdvertiserCompetitiveExclusion: typeutil.ConvertToPointer(false),
								CompetitiveConstraintScope:                typeutil.ConvertToPointer(lineitem.CompetitiveConstraintScopePod),
								PrimaryGoal: &lineitem.Goal{
									GoalType: typeutil.ConvertToPointer(lineitem.GoalTypeLifetime),
									UnitType: typeutil.ConvertToPointer(lineitem.UnitTypeImpressions),
									Units:    typeutil.Int64ToPtr(1),
								},
								ChildContentEligibility: typeutil.ConvertToPointer(lineitem.ChildContentEligibilityDisallowed),
							},
							Targeting: &lineitem.Targeting{
								InventoryTargeting: &lineitem.InventoryTargeting{
									TargetedPlacementIds: []*int64{typeutil.Int64ToPtr(3000)},
								},
							},
						},
					},
				},
			},
			prepare: func(client *soap.MockClient) {
				client.EXPECT().Call(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte(`<?xml version="1.0" encoding="UTF-8"?>
					<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
					<soap:Header><ResponseHeader xmlns="https://www.google.com/apis/ads/publisher/v202402"><requestId>05e6c8196001f4d1d5d0f901981e77ba</requestId><responseTime>706</responseTime></ResponseHeader></soap:Header>
					<soap:Body>
						<createLineItemsResponse xmlns="https://www.google.com/apis/ads/publisher/v202402">
							<rval>
								<orderId>1000</orderId><id>2000</id><name>lineItem2000</name><orderName>order1000</orderName>
								<startDateTime><date><year>2024</year><month>5</month><day>11</day></date><hour>12</hour><minute>0</minute><second>0</second><timeZoneId>America/Los_Angeles</timeZoneId></startDateTime>
								<startDateTimeType>USE_START_DATE_TIME</startDateTimeType>
								<endDateTime><date><year>2024</year><month>5</month><day>20</day></date><hour>12</hour><minute>0</minute><second>0</second><timeZoneId>America/Los_Angeles</timeZoneId></endDateTime>
								<autoExtensionDays>0</autoExtensionDays><unlimitedEndDateTime>false</unlimitedEndDateTime>
								<creativeRotationType>MANUAL</creativeRotationType><deliveryRateType>AS_FAST_AS_POSSIBLE</deliveryRateType>
								<deliveryForecastSource>HISTORICAL</deliveryForecastSource><roadblockingType>ONE_OR_MORE</roadblockingType>
								<skippableAdType>DISABLED</skippableAdType><lineItemType>STANDARD</lineItemType><priority>8</priority>
								<costPerUnit><currencyCode>USD</currencyCode><microAmount>0</microAmount></costPerUnit>
								<valueCostPerUnit><currencyCode>USD</currencyCode><microAmount>0</microAmount></valueCostPerUnit>
								<costType>CPM</costType><discountType>PERCENTAGE</discountType><discount>0.0</discount>
								<contractedUnitsBought>0</contractedUnitsBought>
								<creativePlaceholders>
									<size><width>1</width><height>1</height><isAspectRatio>false</isAspectRatio></size>
									<expectedCreativeCount>1</expectedCreativeCount><creativeSizeType>PIXEL</creativeSizeType>
									<isAmpOnly>false</isAmpOnly>
								</creativePlaceholders>
								<environmentType>BROWSER</environmentType>
								<companionDeliveryOption>UNKNOWN</companionDeliveryOption><allowOverbook>false</allowOverbook>
								<skipInventoryCheck>false</skipInventoryCheck><skipCrossSellingRuleWarningChecks>false</skipCrossSellingRuleWarningChecks>
								<reserveAtCreation>false</reserveAtCreation><budget><currencyCode>USD</currencyCode><microAmount>100000000</microAmount></budget>
								<status>DRAFT</status><reservationStatus>RESERVED</reservationStatus><isArchived>false</isArchived>
								<webPropertyCode></webPropertyCode><disableSameAdvertiserCompetitiveExclusion>false</disableSameAdvertiserCompetitiveExclusion>
								<lastModifiedByApp>app_1</lastModifiedByApp><competitiveConstraintScope>POD</competitiveConstraintScope>
								<lastModifiedDateTime><date><year>2024</year><month>5</month><day>11</day></date><hour>12</hour><minute>0</minute><second>0</second><timeZoneId>America/Los_Angeles</timeZoneId></lastModifiedDateTime>
								<creationDateTime><date><year>2024</year><month>5</month><day>11</day></date><hour>12</hour><minute>0</minute><second>0</second><timeZoneId>America/Los_Angeles</timeZoneId></creationDateTime>
								<isMissingCreatives>false</isMissingCreatives><youtubeKidsRestricted>false</youtubeKidsRestricted><videoMaxDuration>0</videoMaxDuration>
								<primaryGoal><goalType>LIFETIME</goalType><unitType>IMPRESSIONS</unitType><units>1</units></primaryGoal>
								<childContentEligibility>DISALLOWED</childContentEligibility>
								<targeting><inventoryTargeting><targetedPlacementIds>3000</targetedPlacementIds></inventoryTargeting><requestPlatformTargeting></requestPlatformTargeting></targeting>
							</rval></createLineItemsResponse></soap:Body></soap:Envelope>
				`), nil)
			},
			want: &lineitem.CreateLineItemsResponse{
				Rval: []*lineitem.LineItem{
					{
						LineItemSummary: &lineitem.LineItemSummary{
							OrderId:   typeutil.Int64ToPtr(1000),
							Id:        typeutil.Int64ToPtr(2000),
							OrderName: typeutil.StringToPtr("order1000"),
							Name:      typeutil.StringToPtr("lineItem2000"),
							StartDateTime: &lineitem.DateTime{
								Date: &lineitem.Date{
									Year:  typeutil.IntToPtr(2024),
									Month: typeutil.IntToPtr(5),
									Day:   typeutil.IntToPtr(11),
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
							DeliveryRateType:       typeutil.ConvertToPointer(lineitem.DeliveryRateTypeAsFastAsPossible),
							DeliveryForecastSource: typeutil.ConvertToPointer(lineitem.DeliveryForecastSourceHistorical),
							RoadblockingType:       typeutil.ConvertToPointer(lineitem.RoadblockingTypeOneOrMore),
							SkippableAdType:        typeutil.ConvertToPointer(lineitem.SkippableAdTypeDisabled),
							FrequencyCaps:          nil,
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
							Budget: &lineitem.Money{
								CurrencyCode: typeutil.StringToPtr("USD"),
								MicroAmount:  typeutil.Int64ToPtr(100000000),
							},
							Status:            typeutil.ConvertToPointer(lineitem.ComputedStatusDraft),
							ReservationStatus: typeutil.ConvertToPointer(lineitem.LineItemSummaryReservationStatusTypeReserved),
							IsArchived:        typeutil.ConvertToPointer(false),
							WebPropertyCode:   typeutil.StringToPtr(""),
							DisableSameAdvertiserCompetitiveExclusion: typeutil.ConvertToPointer(false),
							LastModifiedByApp:                         typeutil.StringToPtr("app_1"),
							CompetitiveConstraintScope:                typeutil.ConvertToPointer(lineitem.CompetitiveConstraintScopePod),
							LastModifiedDateTime: &lineitem.DateTime{
								Date: &lineitem.Date{
									Year:  typeutil.IntToPtr(2024),
									Month: typeutil.IntToPtr(5),
									Day:   typeutil.IntToPtr(11),
								},
								Hour:       typeutil.IntToPtr(12),
								Minute:     typeutil.IntToPtr(0),
								Second:     typeutil.IntToPtr(0),
								TimeZoneId: typeutil.StringToPtr("America/Los_Angeles"),
							},
							CreationDateTime: &lineitem.DateTime{
								Date: &lineitem.Date{
									Year:  typeutil.IntToPtr(2024),
									Month: typeutil.IntToPtr(5),
									Day:   typeutil.IntToPtr(11),
								},
								Hour:       typeutil.IntToPtr(12),
								Minute:     typeutil.IntToPtr(0),
								Second:     typeutil.IntToPtr(0),
								TimeZoneId: typeutil.StringToPtr("America/Los_Angeles"),
							},
							IsMissingCreatives:    typeutil.ConvertToPointer(false),
							YoutubeKidsRestricted: typeutil.ConvertToPointer(false),
							VideoMaxDuration:      typeutil.Int64ToPtr(0),
							PrimaryGoal: &lineitem.Goal{
								GoalType: typeutil.ConvertToPointer(lineitem.GoalTypeLifetime),
								UnitType: typeutil.ConvertToPointer(lineitem.UnitTypeImpressions),
								Units:    typeutil.Int64ToPtr(1),
							},
							ChildContentEligibility: typeutil.ConvertToPointer(lineitem.ChildContentEligibilityDisallowed),
						},
						Targeting: &lineitem.Targeting{
							InventoryTargeting: &lineitem.InventoryTargeting{
								TargetedPlacementIds: []*int64{typeutil.Int64ToPtr(3000)},
							},
							RequestPlatformTargeting: &lineitem.RequestPlatformTargeting{},
						},
					},
				},
			},
		},
		{
			name: "fail/not unique line item name",
			fields: fields{
				wsdl:            "https://ads.google.com/apis/ads/publisher/v202402/LineItemService?wsdl",
				networkCode:     12345678,
				applicationName: "App_1",
			},
			args: args{
				ctx: context.Background(),
				lineItems: lineitem.CreateLineItems{
					LineItems: []*lineitem.LineItem{
						{
							LineItemSummary: &lineitem.LineItemSummary{
								OrderId: typeutil.Int64ToPtr(1000),
								Name:    typeutil.StringToPtr("lineItem2000"),
							},
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
								<faultstring>[UniqueError.NOT_UNIQUE @ [0].name; trigger:'lineItem2000']</faultstring>
								<detail>
									<ApiExceptionFault xmlns="https://www.google.com/apis/ads/publisher/v202402">
										<message>[UniqueError.NOT_UNIQUE @ [0].name; trigger:'lineItem2000']</message>
										<errors xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="UniqueError">
											<fieldPath>[0].name</fieldPath>
											<fieldPathElements>
												<index>0</index>
											</fieldPathElements>
											<fieldPathElements>
												<field>name</field>
											</fieldPathElements>
											<trigger>lineItem2000</trigger>
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
			wantErr: lineitem.NewUniqueError("lineItem2000"),
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
			got, err := s.CreateLineItems(tt.args.ctx, tt.args.lineItems)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("CreateLineItems() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("CreateLineItems() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
