// Code generated by gamwsdl/main.go. DO NOT EDIT.
package order

import (
	"errors"
)

func NewApiError(reason string) error {
	return &ApiError{
		Reason: reason,
	}
}

func (e *ApiError) Error() string {
	return "ApiError error: " + e.Reason
}

func (e *ApiError) Is(target error) bool {
	var errType *ApiError
	return errors.As(target, &errType)
}

func NewApiException(reason string) error {
	return &ApiException{
		Reason: reason,
	}
}

func (e *ApiException) Error() string {
	return "ApiException error: " + e.Reason
}

func (e *ApiException) Is(target error) bool {
	var errType *ApiException
	return errors.As(target, &errType)
}

func NewApiVersionError(reason string) error {
	return &ApiVersionError{
		Reason: reason,
	}
}

func (e *ApiVersionError) Error() string {
	return "ApiVersionError error: " + e.Reason
}

func (e *ApiVersionError) Is(target error) bool {
	var errType *ApiVersionError
	return errors.As(target, &errType)
}

func NewApplicationException(reason string) error {
	return &ApplicationException{
		Reason: reason,
	}
}

func (e *ApplicationException) Error() string {
	return "ApplicationException error: " + e.Reason
}

func (e *ApplicationException) Is(target error) bool {
	var errType *ApplicationException
	return errors.As(target, &errType)
}

func NewAssetError(reason string) error {
	return &AssetError{
		Reason: reason,
	}
}

func (e *AssetError) Error() string {
	return "AssetError error: " + e.Reason
}

func (e *AssetError) Is(target error) bool {
	var errType *AssetError
	return errors.As(target, &errType)
}

func NewAudienceExtensionError(reason string) error {
	return &AudienceExtensionError{
		Reason: reason,
	}
}

func (e *AudienceExtensionError) Error() string {
	return "AudienceExtensionError error: " + e.Reason
}

func (e *AudienceExtensionError) Is(target error) bool {
	var errType *AudienceExtensionError
	return errors.As(target, &errType)
}

func NewAuthenticationError(reason string) error {
	return &AuthenticationError{
		Reason: reason,
	}
}

func (e *AuthenticationError) Error() string {
	return "AuthenticationError error: " + e.Reason
}

func (e *AuthenticationError) Is(target error) bool {
	var errType *AuthenticationError
	return errors.As(target, &errType)
}

func NewClickTrackingLineItemError(reason string) error {
	return &ClickTrackingLineItemError{
		Reason: reason,
	}
}

func (e *ClickTrackingLineItemError) Error() string {
	return "ClickTrackingLineItemError error: " + e.Reason
}

func (e *ClickTrackingLineItemError) Is(target error) bool {
	var errType *ClickTrackingLineItemError
	return errors.As(target, &errType)
}

func NewCollectionSizeError(reason string) error {
	return &CollectionSizeError{
		Reason: reason,
	}
}

func (e *CollectionSizeError) Error() string {
	return "CollectionSizeError error: " + e.Reason
}

func (e *CollectionSizeError) Is(target error) bool {
	var errType *CollectionSizeError
	return errors.As(target, &errType)
}

func NewCommonError(reason string) error {
	return &CommonError{
		Reason: reason,
	}
}

func (e *CommonError) Error() string {
	return "CommonError error: " + e.Reason
}

func (e *CommonError) Is(target error) bool {
	var errType *CommonError
	return errors.As(target, &errType)
}

func NewCompanyCreditStatusError(reason string) error {
	return &CompanyCreditStatusError{
		Reason: reason,
	}
}

func (e *CompanyCreditStatusError) Error() string {
	return "CompanyCreditStatusError error: " + e.Reason
}

func (e *CompanyCreditStatusError) Is(target error) bool {
	var errType *CompanyCreditStatusError
	return errors.As(target, &errType)
}

func NewCreativeError(reason string) error {
	return &CreativeError{
		Reason: reason,
	}
}

func (e *CreativeError) Error() string {
	return "CreativeError error: " + e.Reason
}

func (e *CreativeError) Is(target error) bool {
	var errType *CreativeError
	return errors.As(target, &errType)
}

func NewCrossSellError(reason string) error {
	return &CrossSellError{
		Reason: reason,
	}
}

func (e *CrossSellError) Error() string {
	return "CrossSellError error: " + e.Reason
}

func (e *CrossSellError) Is(target error) bool {
	var errType *CrossSellError
	return errors.As(target, &errType)
}

func NewCurrencyCodeError(reason string) error {
	return &CurrencyCodeError{
		Reason: reason,
	}
}

func (e *CurrencyCodeError) Error() string {
	return "CurrencyCodeError error: " + e.Reason
}

func (e *CurrencyCodeError) Is(target error) bool {
	var errType *CurrencyCodeError
	return errors.As(target, &errType)
}

func NewCustomFieldValueError(reason string) error {
	return &CustomFieldValueError{
		Reason: reason,
	}
}

func (e *CustomFieldValueError) Error() string {
	return "CustomFieldValueError error: " + e.Reason
}

func (e *CustomFieldValueError) Is(target error) bool {
	var errType *CustomFieldValueError
	return errors.As(target, &errType)
}

func NewCustomTargetingError(reason string) error {
	return &CustomTargetingError{
		Reason: reason,
	}
}

func (e *CustomTargetingError) Error() string {
	return "CustomTargetingError error: " + e.Reason
}

func (e *CustomTargetingError) Is(target error) bool {
	var errType *CustomTargetingError
	return errors.As(target, &errType)
}

func NewDateTimeRangeTargetingError(reason string) error {
	return &DateTimeRangeTargetingError{
		Reason: reason,
	}
}

func (e *DateTimeRangeTargetingError) Error() string {
	return "DateTimeRangeTargetingError error: " + e.Reason
}

func (e *DateTimeRangeTargetingError) Is(target error) bool {
	var errType *DateTimeRangeTargetingError
	return errors.As(target, &errType)
}

func NewDayPartTargetingError(reason string) error {
	return &DayPartTargetingError{
		Reason: reason,
	}
}

func (e *DayPartTargetingError) Error() string {
	return "DayPartTargetingError error: " + e.Reason
}

func (e *DayPartTargetingError) Is(target error) bool {
	var errType *DayPartTargetingError
	return errors.As(target, &errType)
}

func NewEntityChildrenLimitReachedError(reason string) error {
	return &EntityChildrenLimitReachedError{
		Reason: reason,
	}
}

func (e *EntityChildrenLimitReachedError) Error() string {
	return "EntityChildrenLimitReachedError error: " + e.Reason
}

func (e *EntityChildrenLimitReachedError) Is(target error) bool {
	var errType *EntityChildrenLimitReachedError
	return errors.As(target, &errType)
}

func NewEntityLimitReachedError(reason string) error {
	return &EntityLimitReachedError{
		Reason: reason,
	}
}

func (e *EntityLimitReachedError) Error() string {
	return "EntityLimitReachedError error: " + e.Reason
}

func (e *EntityLimitReachedError) Is(target error) bool {
	var errType *EntityLimitReachedError
	return errors.As(target, &errType)
}

func NewFeatureError(reason string) error {
	return &FeatureError{
		Reason: reason,
	}
}

func (e *FeatureError) Error() string {
	return "FeatureError error: " + e.Reason
}

func (e *FeatureError) Is(target error) bool {
	var errType *FeatureError
	return errors.As(target, &errType)
}

func NewForecastError(reason string) error {
	return &ForecastError{
		Reason: reason,
	}
}

func (e *ForecastError) Error() string {
	return "ForecastError error: " + e.Reason
}

func (e *ForecastError) Is(target error) bool {
	var errType *ForecastError
	return errors.As(target, &errType)
}

func NewFrequencyCapError(reason string) error {
	return &FrequencyCapError{
		Reason: reason,
	}
}

func (e *FrequencyCapError) Error() string {
	return "FrequencyCapError error: " + e.Reason
}

func (e *FrequencyCapError) Is(target error) bool {
	var errType *FrequencyCapError
	return errors.As(target, &errType)
}

func NewGenericTargetingError(reason string) error {
	return &GenericTargetingError{
		Reason: reason,
	}
}

func (e *GenericTargetingError) Error() string {
	return "GenericTargetingError error: " + e.Reason
}

func (e *GenericTargetingError) Is(target error) bool {
	var errType *GenericTargetingError
	return errors.As(target, &errType)
}

func NewGeoTargetingError(reason string) error {
	return &GeoTargetingError{
		Reason: reason,
	}
}

func (e *GeoTargetingError) Error() string {
	return "GeoTargetingError error: " + e.Reason
}

func (e *GeoTargetingError) Is(target error) bool {
	var errType *GeoTargetingError
	return errors.As(target, &errType)
}

func NewGrpSettingsError(reason string) error {
	return &GrpSettingsError{
		Reason: reason,
	}
}

func (e *GrpSettingsError) Error() string {
	return "GrpSettingsError error: " + e.Reason
}

func (e *GrpSettingsError) Is(target error) bool {
	var errType *GrpSettingsError
	return errors.As(target, &errType)
}

func NewImageError(reason string) error {
	return &ImageError{
		Reason: reason,
	}
}

func (e *ImageError) Error() string {
	return "ImageError error: " + e.Reason
}

func (e *ImageError) Is(target error) bool {
	var errType *ImageError
	return errors.As(target, &errType)
}

func NewInternalApiError(reason string) error {
	return &InternalApiError{
		Reason: reason,
	}
}

func (e *InternalApiError) Error() string {
	return "InternalApiError error: " + e.Reason
}

func (e *InternalApiError) Is(target error) bool {
	var errType *InternalApiError
	return errors.As(target, &errType)
}

func NewInvalidEmailError(reason string) error {
	return &InvalidEmailError{
		Reason: reason,
	}
}

func (e *InvalidEmailError) Error() string {
	return "InvalidEmailError error: " + e.Reason
}

func (e *InvalidEmailError) Is(target error) bool {
	var errType *InvalidEmailError
	return errors.As(target, &errType)
}

func NewInvalidUrlError(reason string) error {
	return &InvalidUrlError{
		Reason: reason,
	}
}

func (e *InvalidUrlError) Error() string {
	return "InvalidUrlError error: " + e.Reason
}

func (e *InvalidUrlError) Is(target error) bool {
	var errType *InvalidUrlError
	return errors.As(target, &errType)
}

func NewInventoryTargetingError(reason string) error {
	return &InventoryTargetingError{
		Reason: reason,
	}
}

func (e *InventoryTargetingError) Error() string {
	return "InventoryTargetingError error: " + e.Reason
}

func (e *InventoryTargetingError) Is(target error) bool {
	var errType *InventoryTargetingError
	return errors.As(target, &errType)
}

func NewLabelEntityAssociationError(reason string) error {
	return &LabelEntityAssociationError{
		Reason: reason,
	}
}

func (e *LabelEntityAssociationError) Error() string {
	return "LabelEntityAssociationError error: " + e.Reason
}

func (e *LabelEntityAssociationError) Is(target error) bool {
	var errType *LabelEntityAssociationError
	return errors.As(target, &errType)
}

func NewLineItemActivityAssociationError(reason string) error {
	return &LineItemActivityAssociationError{
		Reason: reason,
	}
}

func (e *LineItemActivityAssociationError) Error() string {
	return "LineItemActivityAssociationError error: " + e.Reason
}

func (e *LineItemActivityAssociationError) Is(target error) bool {
	var errType *LineItemActivityAssociationError
	return errors.As(target, &errType)
}

func NewLineItemCreativeAssociationError(reason string) error {
	return &LineItemCreativeAssociationError{
		Reason: reason,
	}
}

func (e *LineItemCreativeAssociationError) Error() string {
	return "LineItemCreativeAssociationError error: " + e.Reason
}

func (e *LineItemCreativeAssociationError) Is(target error) bool {
	var errType *LineItemCreativeAssociationError
	return errors.As(target, &errType)
}

func NewLineItemError(reason string) error {
	return &LineItemError{
		Reason: reason,
	}
}

func (e *LineItemError) Error() string {
	return "LineItemError error: " + e.Reason
}

func (e *LineItemError) Is(target error) bool {
	var errType *LineItemError
	return errors.As(target, &errType)
}

func NewLineItemFlightDateError(reason string) error {
	return &LineItemFlightDateError{
		Reason: reason,
	}
}

func (e *LineItemFlightDateError) Error() string {
	return "LineItemFlightDateError error: " + e.Reason
}

func (e *LineItemFlightDateError) Is(target error) bool {
	var errType *LineItemFlightDateError
	return errors.As(target, &errType)
}

func NewLineItemOperationError(reason string) error {
	return &LineItemOperationError{
		Reason: reason,
	}
}

func (e *LineItemOperationError) Error() string {
	return "LineItemOperationError error: " + e.Reason
}

func (e *LineItemOperationError) Is(target error) bool {
	var errType *LineItemOperationError
	return errors.As(target, &errType)
}

func NewMobileApplicationTargetingError(reason string) error {
	return &MobileApplicationTargetingError{
		Reason: reason,
	}
}

func (e *MobileApplicationTargetingError) Error() string {
	return "MobileApplicationTargetingError error: " + e.Reason
}

func (e *MobileApplicationTargetingError) Is(target error) bool {
	var errType *MobileApplicationTargetingError
	return errors.As(target, &errType)
}

func NewNotNullError(reason string) error {
	return &NotNullError{
		Reason: reason,
	}
}

func (e *NotNullError) Error() string {
	return "NotNullError error: " + e.Reason
}

func (e *NotNullError) Is(target error) bool {
	var errType *NotNullError
	return errors.As(target, &errType)
}

func NewNullError(reason string) error {
	return &NullError{
		Reason: reason,
	}
}

func (e *NullError) Error() string {
	return "NullError error: " + e.Reason
}

func (e *NullError) Is(target error) bool {
	var errType *NullError
	return errors.As(target, &errType)
}

func NewOrderActionError(reason string) error {
	return &OrderActionError{
		Reason: reason,
	}
}

func (e *OrderActionError) Error() string {
	return "OrderActionError error: " + e.Reason
}

func (e *OrderActionError) Is(target error) bool {
	var errType *OrderActionError
	return errors.As(target, &errType)
}

func NewOrderError(reason string) error {
	return &OrderError{
		Reason: reason,
	}
}

func (e *OrderError) Error() string {
	return "OrderError error: " + e.Reason
}

func (e *OrderError) Is(target error) bool {
	var errType *OrderError
	return errors.As(target, &errType)
}

func NewParseError(reason string) error {
	return &ParseError{
		Reason: reason,
	}
}

func (e *ParseError) Error() string {
	return "ParseError error: " + e.Reason
}

func (e *ParseError) Is(target error) bool {
	var errType *ParseError
	return errors.As(target, &errType)
}

func NewPermissionError(reason string) error {
	return &PermissionError{
		Reason: reason,
	}
}

func (e *PermissionError) Error() string {
	return "PermissionError error: " + e.Reason
}

func (e *PermissionError) Is(target error) bool {
	var errType *PermissionError
	return errors.As(target, &errType)
}

func NewPrecisionError(reason string) error {
	return &PrecisionError{
		Reason: reason,
	}
}

func (e *PrecisionError) Error() string {
	return "PrecisionError error: " + e.Reason
}

func (e *PrecisionError) Is(target error) bool {
	var errType *PrecisionError
	return errors.As(target, &errType)
}

func NewProgrammaticError(reason string) error {
	return &ProgrammaticError{
		Reason: reason,
	}
}

func (e *ProgrammaticError) Error() string {
	return "ProgrammaticError error: " + e.Reason
}

func (e *ProgrammaticError) Is(target error) bool {
	var errType *ProgrammaticError
	return errors.As(target, &errType)
}

func NewPublisherQueryLanguageContextError(reason string) error {
	return &PublisherQueryLanguageContextError{
		Reason: reason,
	}
}

func (e *PublisherQueryLanguageContextError) Error() string {
	return "PublisherQueryLanguageContextError error: " + e.Reason
}

func (e *PublisherQueryLanguageContextError) Is(target error) bool {
	var errType *PublisherQueryLanguageContextError
	return errors.As(target, &errType)
}

func NewPublisherQueryLanguageSyntaxError(reason string) error {
	return &PublisherQueryLanguageSyntaxError{
		Reason: reason,
	}
}

func (e *PublisherQueryLanguageSyntaxError) Error() string {
	return "PublisherQueryLanguageSyntaxError error: " + e.Reason
}

func (e *PublisherQueryLanguageSyntaxError) Is(target error) bool {
	var errType *PublisherQueryLanguageSyntaxError
	return errors.As(target, &errType)
}

func NewQuotaError(reason string) error {
	return &QuotaError{
		Reason: reason,
	}
}

func (e *QuotaError) Error() string {
	return "QuotaError error: " + e.Reason
}

func (e *QuotaError) Is(target error) bool {
	var errType *QuotaError
	return errors.As(target, &errType)
}

func NewRangeError(reason string) error {
	return &RangeError{
		Reason: reason,
	}
}

func (e *RangeError) Error() string {
	return "RangeError error: " + e.Reason
}

func (e *RangeError) Is(target error) bool {
	var errType *RangeError
	return errors.As(target, &errType)
}

func NewRegExError(reason string) error {
	return &RegExError{
		Reason: reason,
	}
}

func (e *RegExError) Error() string {
	return "RegExError error: " + e.Reason
}

func (e *RegExError) Is(target error) bool {
	var errType *RegExError
	return errors.As(target, &errType)
}

func NewRequestPlatformTargetingError(reason string) error {
	return &RequestPlatformTargetingError{
		Reason: reason,
	}
}

func (e *RequestPlatformTargetingError) Error() string {
	return "RequestPlatformTargetingError error: " + e.Reason
}

func (e *RequestPlatformTargetingError) Is(target error) bool {
	var errType *RequestPlatformTargetingError
	return errors.As(target, &errType)
}

func NewRequiredCollectionError(reason string) error {
	return &RequiredCollectionError{
		Reason: reason,
	}
}

func (e *RequiredCollectionError) Error() string {
	return "RequiredCollectionError error: " + e.Reason
}

func (e *RequiredCollectionError) Is(target error) bool {
	var errType *RequiredCollectionError
	return errors.As(target, &errType)
}

func NewRequiredError(reason string) error {
	return &RequiredError{
		Reason: reason,
	}
}

func (e *RequiredError) Error() string {
	return "RequiredError error: " + e.Reason
}

func (e *RequiredError) Is(target error) bool {
	var errType *RequiredError
	return errors.As(target, &errType)
}

func NewRequiredNumberError(reason string) error {
	return &RequiredNumberError{
		Reason: reason,
	}
}

func (e *RequiredNumberError) Error() string {
	return "RequiredNumberError error: " + e.Reason
}

func (e *RequiredNumberError) Is(target error) bool {
	var errType *RequiredNumberError
	return errors.As(target, &errType)
}

func NewRequiredSizeError(reason string) error {
	return &RequiredSizeError{
		Reason: reason,
	}
}

func (e *RequiredSizeError) Error() string {
	return "RequiredSizeError error: " + e.Reason
}

func (e *RequiredSizeError) Is(target error) bool {
	var errType *RequiredSizeError
	return errors.As(target, &errType)
}

func NewReservationDetailsError(reason string) error {
	return &ReservationDetailsError{
		Reason: reason,
	}
}

func (e *ReservationDetailsError) Error() string {
	return "ReservationDetailsError error: " + e.Reason
}

func (e *ReservationDetailsError) Is(target error) bool {
	var errType *ReservationDetailsError
	return errors.As(target, &errType)
}

func NewAudienceSegmentError(reason string) error {
	return &AudienceSegmentError{
		Reason: reason,
	}
}

func (e *AudienceSegmentError) Error() string {
	return "AudienceSegmentError error: " + e.Reason
}

func (e *AudienceSegmentError) Is(target error) bool {
	var errType *AudienceSegmentError
	return errors.As(target, &errType)
}

func NewServerError(reason string) error {
	return &ServerError{
		Reason: reason,
	}
}

func (e *ServerError) Error() string {
	return "ServerError error: " + e.Reason
}

func (e *ServerError) Is(target error) bool {
	var errType *ServerError
	return errors.As(target, &errType)
}

func NewSetTopBoxLineItemError(reason string) error {
	return &SetTopBoxLineItemError{
		Reason: reason,
	}
}

func (e *SetTopBoxLineItemError) Error() string {
	return "SetTopBoxLineItemError error: " + e.Reason
}

func (e *SetTopBoxLineItemError) Is(target error) bool {
	var errType *SetTopBoxLineItemError
	return errors.As(target, &errType)
}

func NewStatementError(reason string) error {
	return &StatementError{
		Reason: reason,
	}
}

func (e *StatementError) Error() string {
	return "StatementError error: " + e.Reason
}

func (e *StatementError) Is(target error) bool {
	var errType *StatementError
	return errors.As(target, &errType)
}

func NewStringFormatError(reason string) error {
	return &StringFormatError{
		Reason: reason,
	}
}

func (e *StringFormatError) Error() string {
	return "StringFormatError error: " + e.Reason
}

func (e *StringFormatError) Is(target error) bool {
	var errType *StringFormatError
	return errors.As(target, &errType)
}

func NewStringLengthError(reason string) error {
	return &StringLengthError{
		Reason: reason,
	}
}

func (e *StringLengthError) Error() string {
	return "StringLengthError error: " + e.Reason
}

func (e *StringLengthError) Is(target error) bool {
	var errType *StringLengthError
	return errors.As(target, &errType)
}

func NewTeamError(reason string) error {
	return &TeamError{
		Reason: reason,
	}
}

func (e *TeamError) Error() string {
	return "TeamError error: " + e.Reason
}

func (e *TeamError) Is(target error) bool {
	var errType *TeamError
	return errors.As(target, &errType)
}

func NewTechnologyTargetingError(reason string) error {
	return &TechnologyTargetingError{
		Reason: reason,
	}
}

func (e *TechnologyTargetingError) Error() string {
	return "TechnologyTargetingError error: " + e.Reason
}

func (e *TechnologyTargetingError) Is(target error) bool {
	var errType *TechnologyTargetingError
	return errors.As(target, &errType)
}

func NewTemplateInstantiatedCreativeError(reason string) error {
	return &TemplateInstantiatedCreativeError{
		Reason: reason,
	}
}

func (e *TemplateInstantiatedCreativeError) Error() string {
	return "TemplateInstantiatedCreativeError error: " + e.Reason
}

func (e *TemplateInstantiatedCreativeError) Is(target error) bool {
	var errType *TemplateInstantiatedCreativeError
	return errors.As(target, &errType)
}

func NewTimeZoneError(reason string) error {
	return &TimeZoneError{
		Reason: reason,
	}
}

func (e *TimeZoneError) Error() string {
	return "TimeZoneError error: " + e.Reason
}

func (e *TimeZoneError) Is(target error) bool {
	var errType *TimeZoneError
	return errors.As(target, &errType)
}

func NewTranscodingError(reason string) error {
	return &TranscodingError{
		Reason: reason,
	}
}

func (e *TranscodingError) Error() string {
	return "TranscodingError error: " + e.Reason
}

func (e *TranscodingError) Is(target error) bool {
	var errType *TranscodingError
	return errors.As(target, &errType)
}

func NewTypeError(reason string) error {
	return &TypeError{
		Reason: reason,
	}
}

func (e *TypeError) Error() string {
	return "TypeError error: " + e.Reason
}

func (e *TypeError) Is(target error) bool {
	var errType *TypeError
	return errors.As(target, &errType)
}

func NewUniqueError(reason string) error {
	return &UniqueError{
		Reason: reason,
	}
}

func (e *UniqueError) Error() string {
	return "UniqueError error: " + e.Reason
}

func (e *UniqueError) Is(target error) bool {
	var errType *UniqueError
	return errors.As(target, &errType)
}

func NewUserDomainTargetingError(reason string) error {
	return &UserDomainTargetingError{
		Reason: reason,
	}
}

func (e *UserDomainTargetingError) Error() string {
	return "UserDomainTargetingError error: " + e.Reason
}

func (e *UserDomainTargetingError) Is(target error) bool {
	var errType *UserDomainTargetingError
	return errors.As(target, &errType)
}

func NewVideoPositionTargetingError(reason string) error {
	return &VideoPositionTargetingError{
		Reason: reason,
	}
}

func (e *VideoPositionTargetingError) Error() string {
	return "VideoPositionTargetingError error: " + e.Reason
}

func (e *VideoPositionTargetingError) Is(target error) bool {
	var errType *VideoPositionTargetingError
	return errors.As(target, &errType)
}

// RaiseError raises an error based on the error type and reason.
func RaiseError(errorType, reason string) error {
	switch errorType {

	case "ApiError":
		return NewApiError(reason)

	case "ApiException":
		return NewApiException(reason)

	case "ApiVersionError":
		return NewApiVersionError(reason)

	case "ApplicationException":
		return NewApplicationException(reason)

	case "AssetError":
		return NewAssetError(reason)

	case "AudienceExtensionError":
		return NewAudienceExtensionError(reason)

	case "AuthenticationError":
		return NewAuthenticationError(reason)

	case "ClickTrackingLineItemError":
		return NewClickTrackingLineItemError(reason)

	case "CollectionSizeError":
		return NewCollectionSizeError(reason)

	case "CommonError":
		return NewCommonError(reason)

	case "CompanyCreditStatusError":
		return NewCompanyCreditStatusError(reason)

	case "CreativeError":
		return NewCreativeError(reason)

	case "CrossSellError":
		return NewCrossSellError(reason)

	case "CurrencyCodeError":
		return NewCurrencyCodeError(reason)

	case "CustomFieldValueError":
		return NewCustomFieldValueError(reason)

	case "CustomTargetingError":
		return NewCustomTargetingError(reason)

	case "DateTimeRangeTargetingError":
		return NewDateTimeRangeTargetingError(reason)

	case "DayPartTargetingError":
		return NewDayPartTargetingError(reason)

	case "EntityChildrenLimitReachedError":
		return NewEntityChildrenLimitReachedError(reason)

	case "EntityLimitReachedError":
		return NewEntityLimitReachedError(reason)

	case "FeatureError":
		return NewFeatureError(reason)

	case "ForecastError":
		return NewForecastError(reason)

	case "FrequencyCapError":
		return NewFrequencyCapError(reason)

	case "GenericTargetingError":
		return NewGenericTargetingError(reason)

	case "GeoTargetingError":
		return NewGeoTargetingError(reason)

	case "GrpSettingsError":
		return NewGrpSettingsError(reason)

	case "ImageError":
		return NewImageError(reason)

	case "InternalApiError":
		return NewInternalApiError(reason)

	case "InvalidEmailError":
		return NewInvalidEmailError(reason)

	case "InvalidUrlError":
		return NewInvalidUrlError(reason)

	case "InventoryTargetingError":
		return NewInventoryTargetingError(reason)

	case "LabelEntityAssociationError":
		return NewLabelEntityAssociationError(reason)

	case "LineItemActivityAssociationError":
		return NewLineItemActivityAssociationError(reason)

	case "LineItemCreativeAssociationError":
		return NewLineItemCreativeAssociationError(reason)

	case "LineItemError":
		return NewLineItemError(reason)

	case "LineItemFlightDateError":
		return NewLineItemFlightDateError(reason)

	case "LineItemOperationError":
		return NewLineItemOperationError(reason)

	case "MobileApplicationTargetingError":
		return NewMobileApplicationTargetingError(reason)

	case "NotNullError":
		return NewNotNullError(reason)

	case "NullError":
		return NewNullError(reason)

	case "OrderActionError":
		return NewOrderActionError(reason)

	case "OrderError":
		return NewOrderError(reason)

	case "ParseError":
		return NewParseError(reason)

	case "PermissionError":
		return NewPermissionError(reason)

	case "PrecisionError":
		return NewPrecisionError(reason)

	case "ProgrammaticError":
		return NewProgrammaticError(reason)

	case "PublisherQueryLanguageContextError":
		return NewPublisherQueryLanguageContextError(reason)

	case "PublisherQueryLanguageSyntaxError":
		return NewPublisherQueryLanguageSyntaxError(reason)

	case "QuotaError":
		return NewQuotaError(reason)

	case "RangeError":
		return NewRangeError(reason)

	case "RegExError":
		return NewRegExError(reason)

	case "RequestPlatformTargetingError":
		return NewRequestPlatformTargetingError(reason)

	case "RequiredCollectionError":
		return NewRequiredCollectionError(reason)

	case "RequiredError":
		return NewRequiredError(reason)

	case "RequiredNumberError":
		return NewRequiredNumberError(reason)

	case "RequiredSizeError":
		return NewRequiredSizeError(reason)

	case "ReservationDetailsError":
		return NewReservationDetailsError(reason)

	case "AudienceSegmentError":
		return NewAudienceSegmentError(reason)

	case "ServerError":
		return NewServerError(reason)

	case "SetTopBoxLineItemError":
		return NewSetTopBoxLineItemError(reason)

	case "StatementError":
		return NewStatementError(reason)

	case "StringFormatError":
		return NewStringFormatError(reason)

	case "StringLengthError":
		return NewStringLengthError(reason)

	case "TeamError":
		return NewTeamError(reason)

	case "TechnologyTargetingError":
		return NewTechnologyTargetingError(reason)

	case "TemplateInstantiatedCreativeError":
		return NewTemplateInstantiatedCreativeError(reason)

	case "TimeZoneError":
		return NewTimeZoneError(reason)

	case "TranscodingError":
		return NewTranscodingError(reason)

	case "TypeError":
		return NewTypeError(reason)

	case "UniqueError":
		return NewUniqueError(reason)

	case "UserDomainTargetingError":
		return NewUserDomainTargetingError(reason)

	case "VideoPositionTargetingError":
		return NewVideoPositionTargetingError(reason)

	default:
		return nil
	}
}