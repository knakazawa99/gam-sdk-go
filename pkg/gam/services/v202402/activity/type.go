// Code generated by gamwsdl/main.go. DO NOT EDIT.
package activity

// ObjectValue
//
// Contains an object value.
// <p>
// <b>This object is experimental!
// <code>ObjectValue</code> is an experimental, innovative, and rapidly
// changing new feature for Ad Manager. Unfortunately, being on the bleeding edge means that
// we may make backwards-incompatible changes to
// <code>ObjectValue</code>. We will inform the community when this feature
// is no longer experimental.</b>
//

type ObjectValue struct {
}

// Activity
//
// An activity is a specific user action that an advertiser wants to track, such as the completion
// of a purchase or a visit to a webpage. You create and manage activities in Ad Manager. When a
// user performs the action after seeing an advertiser's ad, that's a conversion.
//
// <p>For example, you set up an activity in Ad Manager to track how many users visit an
// advertiser's promotional website after viewing or clicking on an ad. When a user views an ad,
// then visits the page, that's one conversion.
//

type Activity struct {

	// Id
	//
	// The unique ID of the {@code Activity}. This value is readonly and is assigned by Google.
	//

	Id *int64 `xml:"id"`

	// ActivityGroupId
	//
	// The ID of the {@link ActivityGroup} that this {@link Activity} belongs to.
	//

	ActivityGroupId *int64 `xml:"activityGroupId"`

	// Name
	//
	// The name of the {@code Activity}. This attribute is required and has a maximum length of 255
	// characters.
	//

	Name *string `xml:"name"`

	// ExpectedURL
	//
	// The URL of the webpage where the tags from this activity will be placed. This attribute is
	// optional.
	//

	ExpectedURL *string `xml:"expectedURL"`

	// Status
	//
	// The status of this activity. This attribute is readonly.
	//

	Status *ActivityStatus `xml:"status"`

	// Type
	//
	// The activity type. This attribute is optional and defaults to {@link Activity.Type#PAGE_VIEWS}
	//

	Type *ActivityType `xml:"type"`
}

// ActivityError
//
// Errors relating to Activity and Activity Group services.
//

type ActivityError struct {
	Reason string `xml:"reason"`
}

// ActivityPage
//
// Captures a page of {@link Activity} objects.
//

type ActivityPage struct {

	// TotalResultSetSize
	//
	// The size of the total result set to which this page belongs.
	//

	TotalResultSetSize *int `xml:"totalResultSetSize"`

	// StartIndex
	//
	// The absolute index in the total result set on which this page begins.
	//

	StartIndex *int `xml:"startIndex"`

	// Results
	//
	// The collection of activities contained within this page.
	//

	Results []*Activity `xml:"results"`
}

// ApiError
//
// The API error base class that provides details about an error that occurred
// while processing a service request.
//
// <p>The OGNL field path is provided for parsers to identify the request data
// element that may have caused the error.</p>
//

type ApiError struct {
	Reason string `xml:"reason"`
}

// ApiException
//
// Exception class for holding a list of service errors.
//

type ApiException struct {
	Reason string `xml:"reason"`
}

// ApiVersionError
//
// Errors related to the usage of API versions.
//

type ApiVersionError struct {
	Reason string `xml:"reason"`
}

// ApplicationException
//
// Base class for exceptions.
//

type ApplicationException struct {
	Reason string `xml:"reason"`
}

// AuthenticationError
//
// An error for an exception that occurred when authenticating.
//

type AuthenticationError struct {
	Reason string `xml:"reason"`
}

// BooleanValue
//
// Contains a boolean value.
//

type BooleanValue struct {

	// Value
	//
	// The boolean value.
	//

	Value *bool `xml:"value"`
}

// CollectionSizeError
//
// Error for the size of the collection being too large
//

type CollectionSizeError struct {
	Reason string `xml:"reason"`
}

// CommonError
//
// A place for common errors that can be used across services.
//

type CommonError struct {
	Reason string `xml:"reason"`
}

// Date
//
// Represents a date.
//

type Date struct {

	// Year
	//
	// Year (e.g., 2009)
	//

	Year *int `xml:"year"`

	// Month
	//
	// Month (1..12)
	//

	Month *int `xml:"month"`

	// Day
	//
	// Day (1..31)
	//

	Day *int `xml:"day"`
}

// DateTime
//
// Represents a date combined with the time of day.
//

type DateTime struct {

	// Date

	Date *Date `xml:"date"`

	// Hour

	Hour *int `xml:"hour"`

	// Minute

	Minute *int `xml:"minute"`

	// Second

	Second *int `xml:"second"`

	// TimeZoneId

	TimeZoneId *string `xml:"timeZoneId"`
}

// DateTimeValue
//
// Contains a date-time value.
//

type DateTimeValue struct {

	// Value
	//
	// The {@code DateTime} value.
	//

	Value *DateTime `xml:"value"`
}

// DateValue
//
// Contains a date value.
//

type DateValue struct {

	// Value
	//
	// The {@code Date} value.
	//

	Value *Date `xml:"value"`
}

// FeatureError
//
// Errors related to feature management.  If you attempt using a feature that is not available to
// the current network you'll receive a FeatureError with the missing feature as the trigger.
//

type FeatureError struct {
	Reason string `xml:"reason"`
}

// FieldPathElement
//
// A segment of a field path. Each dot in a field path defines a new segment.
//

type FieldPathElement struct {

	// Field
	//
	// The name of a field in lower camelcase. (e.g. "biddingStrategy")
	//

	Field *string `xml:"field"`

	// Index
	//
	// For list fields, this is a 0-indexed position in the list. Null for non-list fields.
	//

	Index *int `xml:"index"`
}

// InternalApiError
//
// Indicates that a server-side error has occured. {@code InternalApiError}s
// are generally not the result of an invalid request or message sent by the
// client.
//

type InternalApiError struct {
	Reason string `xml:"reason"`
}

// NotNullError
//
// Caused by supplying a null value for an attribute that cannot be null.
//

type NotNullError struct {
	Reason string `xml:"reason"`
}

// NumberValue
//
// Contains a numeric value.
//

type NumberValue struct {

	// Value
	//
	// The numeric value represented as a string.
	//

	Value *string `xml:"value"`
}

// ParseError
//
// Lists errors related to parsing.
//

type ParseError struct {
	Reason string `xml:"reason"`
}

// PermissionError
//
// Errors related to incorrect permission.
//

type PermissionError struct {
	Reason string `xml:"reason"`
}

// PublisherQueryLanguageContextError
//
// An error that occurs while executing a PQL query contained in
// a {@link Statement} object.
//

type PublisherQueryLanguageContextError struct {
	Reason string `xml:"reason"`
}

// PublisherQueryLanguageSyntaxError
//
// An error that occurs while parsing a PQL query contained in a
// {@link Statement} object.
//

type PublisherQueryLanguageSyntaxError struct {
	Reason string `xml:"reason"`
}

// QuotaError
//
// Describes a client-side error on which a user is attempting
// to perform an action to which they have no quota remaining.
//

type QuotaError struct {
	Reason string `xml:"reason"`
}

// RangeError
//
// A list of all errors associated with the Range constraint.
//

type RangeError struct {
	Reason string `xml:"reason"`
}

// RequiredCollectionError
//
// A list of all errors to be used for validating sizes of collections.
//

type RequiredCollectionError struct {
	Reason string `xml:"reason"`
}

// RequiredError
//
// Errors due to missing required field.
//

type RequiredError struct {
	Reason string `xml:"reason"`
}

// ServerError
//
// Errors related to the server.
//

type ServerError struct {
	Reason string `xml:"reason"`
}

// SetValue
//
// Contains a set of {@link Value Values}. May not contain duplicates.
//

type SetValue struct {

	// Value
	*Value

	// Values
	//
	// The values. They must all be the same type of {@code Value} and not contain duplicates.
	//

	Values []*Value `xml:"values"`
}

// SoapRequestHeader
//
// Represents the SOAP request header used by API requests.
//

type SoapRequestHeader struct {

	// NetworkCode
	//
	// The network code to use in the context of a request.
	//

	NetworkCode *string `xml:"networkCode"`

	// ApplicationName
	//
	// The name of client library application.
	//

	ApplicationName *string `xml:"applicationName"`
}

// SoapResponseHeader
//
// Represents the SOAP request header used by API responses.
//

type SoapResponseHeader struct {

	// RequestId

	RequestId *string `xml:"requestId"`

	// ResponseTime

	ResponseTime *int64 `xml:"responseTime"`
}

// Statement
//
// Captures the {@code WHERE}, {@code ORDER BY} and {@code LIMIT} clauses of a
// PQL query. Statements are typically used to retrieve objects of a predefined
// domain type, which makes SELECT clause unnecessary.
// <p>
// An example query text might be {@code "WHERE status = 'ACTIVE' ORDER BY id
// LIMIT 30"}.
// </p>
// <p>
// Statements support bind variables. These are substitutes for literals
// and can be thought of as input parameters to a PQL query.
// </p>
// <p>
// An example of such a query might be {@code "WHERE id = :idValue"}.
// </p>
// <p>
// Statements also support use of the LIKE keyword. This provides wildcard string matching.
// </p>
// <p>
// An example of such a query might be {@code "WHERE name LIKE '%searchString%'"}.
// </p>
// The value for the variable idValue must then be set with an object of type
// {@link Value}, e.g., {@link NumberValue}, {@link TextValue} or
// {@link BooleanValue}.
//

type Statement struct {

	// Query
	//
	// Holds the query in PQL syntax. The syntax is:<br>
	// <code>[WHERE <condition> {[AND | OR] <condition> ...}]</code><br>
	// <code>[ORDER BY <property> [ASC | DESC]]</code><br>
	// <code>[LIMIT {[<offset>,] <count>} | {<count> OFFSET <offset>}]</code><br>
	// <p>
	// <code><condition></code><br>
	// &nbsp;&nbsp;&nbsp;&nbsp;
	// <code>:= <property> {< | <= | > | >= | = | != } <value></code><br>
	// <code><condition></code><br>
	// &nbsp;&nbsp;&nbsp;&nbsp;
	// <code>:= <property> {< | <= | > | >= | = | != } <bind variable></code><br>
	// <code><condition> := <property> IN <list></code><br>
	// <code><condition> := <property> IS NULL</code><br>
	// <code><condition> := <property> LIKE <wildcard%match></code><br>
	// <code><bind variable> := :<name></code><br>
	// </p>
	//

	Query *string `xml:"query"`

	// Values
	//
	// Holds keys and values for bind variables and their values. The key is the
	// name of the bind variable. The value is the literal value of the variable.
	// <p>
	// In the example {@code "WHERE status = :bindStatus ORDER BY id LIMIT 30"},
	// the bind variable, represented by {@code :bindStatus} is named {@code
	// bindStatus}, which would also be the parameter map key. The bind variable's
	// value would be represented by a parameter map value of type
	// {@link TextValue}. The final result, for example, would be an entry of
	// {@code "bindStatus" => StringParam("ACTIVE")}.
	// </p>
	//

	Values []*String_ValueMapEntry `xml:"values"`
}

// StatementError
//
// An error that occurs while parsing {@link Statement} objects.
//

type StatementError struct {
	Reason string `xml:"reason"`
}

// StringFormatError
//
// A list of error code for reporting invalid content of input strings.
//

type StringFormatError struct {
	Reason string `xml:"reason"`
}

// StringLengthError
//
// Errors for Strings which do not meet given length constraints.
//

type StringLengthError struct {
	Reason string `xml:"reason"`
}

// String_ValueMapEntry
//
// This represents an entry in a map with a key of type String
// and value of type Value.
//

type String_ValueMapEntry struct {

	// Key

	Key *string `xml:"key"`

	// Value

	Value *Value `xml:"value"`
}

// TextValue
//
// Contains a string value.
//

type TextValue struct {

	// Value
	//
	// The string value.
	//

	Value *string `xml:"value"`
}

// UniqueError
//
// An error for a field which must satisfy a uniqueness constraint
//

type UniqueError struct {
	Reason string `xml:"reason"`
}

// Value
//
// {@code Value} represents a value.
//

type Value struct {
}

// ActivityStatusType
//
// The activity status.
type ActivityStatus string

const (

	// ActivityStatusTypeActive
	ActivityStatusTypeActive ActivityStatus = "ACTIVE"

	// ActivityStatusTypeInactive
	ActivityStatusTypeInactive ActivityStatus = "INACTIVE"
)

// ActivityTypeType
//
// The activity type.
type ActivityType string

const (

	// ActivityTypeTypePageViews
	//
	// Tracks conversions for each visit to a webpage. This is a counter type.
	//
	ActivityTypeTypePageViews ActivityType = "PAGE_VIEWS"

	// ActivityTypeTypeDailyVisits
	//
	// Tracks conversions for visits to a webpage, but only counts one conversion per user per day,
	// even if a user visits the page multiple times. This is a counter type.
	//
	ActivityTypeTypeDailyVisits ActivityType = "DAILY_VISITS"

	// ActivityTypeTypeCustom
	//
	// Tracks conversions for visits to a webpage, but only counts one conversion per user per user
	// session. Session length is set by the advertiser. This is a counter type.
	//
	ActivityTypeTypeCustom ActivityType = "CUSTOM"

	// ActivityTypeTypeItemsPurchased
	//
	// Tracks conversions where the user has made a purchase, the monetary value of each purchase,
	// plus the number of items that were purchased and the order ID. This is a sales type.
	//
	ActivityTypeTypeItemsPurchased ActivityType = "ITEMS_PURCHASED"

	// ActivityTypeTypeTransactions
	//
	// Tracks conversions where the user has made a purchase, the monetary value of each purchase,
	// plus the order ID (but not the number of items purchased). This is a sales type.
	//
	ActivityTypeTypeTransactions ActivityType = "TRANSACTIONS"

	// ActivityTypeTypeIosApplicationDownloads
	//
	// Tracks conversions where the user has installed an iOS application. This is a counter type.
	//
	ActivityTypeTypeIosApplicationDownloads ActivityType = "IOS_APPLICATION_DOWNLOADS"

	// ActivityTypeTypeAndroidApplicationDownloads
	//
	// Tracks conversions where the user has installed an Android application. This is a counter
	// type.
	//
	ActivityTypeTypeAndroidApplicationDownloads ActivityType = "ANDROID_APPLICATION_DOWNLOADS"

	// ActivityTypeTypeUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	ActivityTypeTypeUnknown ActivityType = "UNKNOWN"
)

// ActivityErrorReasonReason
//
// The reasons for the target error.
const (

	// ActivityErrorReasonReasonActivitiesFeatureRequired
	//
	// The 'activities' feature is required but not enabled.
	//
	ActivityErrorReasonReasonActivitiesFeatureRequired = "ACTIVITIES_FEATURE_REQUIRED"

	// ActivityErrorReasonReasonUnsupportedCompanyType
	//
	// Activity group cannot be associated with the company services.
	//
	ActivityErrorReasonReasonUnsupportedCompanyType = "UNSUPPORTED_COMPANY_TYPE"

	// ActivityErrorReasonReasonDeprecated
	//
	// Activities and ActivityGroups are discontinued as part of Spotlight deprecation.
	//
	ActivityErrorReasonReasonDeprecated = "DEPRECATED"

	// ActivityErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	ActivityErrorReasonReasonUnknown = "UNKNOWN"
)

// ApiVersionErrorReasonReason
const (

	// ApiVersionErrorReasonReasonUpdateToNewerVersion
	//
	// Indicates that the operation is not allowed in the version the request
	// was made in.
	//
	ApiVersionErrorReasonReasonUpdateToNewerVersion = "UPDATE_TO_NEWER_VERSION"

	// ApiVersionErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	ApiVersionErrorReasonReasonUnknown = "UNKNOWN"
)

// AuthenticationErrorReasonReason
const (

	// AuthenticationErrorReasonReasonAmbiguousSoapRequestHeader
	//
	// The SOAP message contains a request header with an ambiguous definition of the authentication
	// header fields. This means either the {@code authToken} and {@code oAuthToken} fields were
	// both null or both were specified. Exactly one value should be specified with each request.
	//
	AuthenticationErrorReasonReasonAmbiguousSoapRequestHeader = "AMBIGUOUS_SOAP_REQUEST_HEADER"

	// AuthenticationErrorReasonReasonInvalidEmail
	//
	// The login provided is invalid.
	//
	AuthenticationErrorReasonReasonInvalidEmail = "INVALID_EMAIL"

	// AuthenticationErrorReasonReasonAuthenticationFailed
	//
	// Tried to authenticate with provided information, but failed.
	//
	AuthenticationErrorReasonReasonAuthenticationFailed = "AUTHENTICATION_FAILED"

	// AuthenticationErrorReasonReasonInvalidOauthSignature
	//
	// The OAuth provided is invalid.
	//
	AuthenticationErrorReasonReasonInvalidOauthSignature = "INVALID_OAUTH_SIGNATURE"

	// AuthenticationErrorReasonReasonInvalidService
	//
	// The specified service to use was not recognized.
	//
	AuthenticationErrorReasonReasonInvalidService = "INVALID_SERVICE"

	// AuthenticationErrorReasonReasonMissingSoapRequestHeader
	//
	// The SOAP message is missing a request header with an {@code authToken} and optional {@code
	// networkCode}.
	//
	AuthenticationErrorReasonReasonMissingSoapRequestHeader = "MISSING_SOAP_REQUEST_HEADER"

	// AuthenticationErrorReasonReasonMissingAuthenticationHttpHeader
	//
	// The HTTP request is missing a request header with an {@code authToken}
	//
	AuthenticationErrorReasonReasonMissingAuthenticationHttpHeader = "MISSING_AUTHENTICATION_HTTP_HEADER"

	// AuthenticationErrorReasonReasonMissingAuthentication
	//
	// The request is missing an {@code authToken}
	//
	AuthenticationErrorReasonReasonMissingAuthentication = "MISSING_AUTHENTICATION"

	// AuthenticationErrorReasonReasonNetworkApiAccessDisabled
	//
	// The network does not have API access enabled.
	//
	AuthenticationErrorReasonReasonNetworkApiAccessDisabled = "NETWORK_API_ACCESS_DISABLED"

	// AuthenticationErrorReasonReasonNoNetworksToAccess
	//
	// The user is not associated with any network.
	//
	AuthenticationErrorReasonReasonNoNetworksToAccess = "NO_NETWORKS_TO_ACCESS"

	// AuthenticationErrorReasonReasonNetworkNotFound
	//
	// No network for the given {@code networkCode} was found.
	//
	AuthenticationErrorReasonReasonNetworkNotFound = "NETWORK_NOT_FOUND"

	// AuthenticationErrorReasonReasonNetworkCodeRequired
	//
	// The user has access to more than one network, but did not provide a {@code networkCode}.
	//
	AuthenticationErrorReasonReasonNetworkCodeRequired = "NETWORK_CODE_REQUIRED"

	// AuthenticationErrorReasonReasonConnectionError
	//
	// An error happened on the server side during connection to authentication service.
	//
	AuthenticationErrorReasonReasonConnectionError = "CONNECTION_ERROR"

	// AuthenticationErrorReasonReasonGoogleAccountAlreadyAssociatedWithNetwork
	//
	// The user tried to create a test network using an account that already is associated with a
	// network.
	//
	AuthenticationErrorReasonReasonGoogleAccountAlreadyAssociatedWithNetwork = "GOOGLE_ACCOUNT_ALREADY_ASSOCIATED_WITH_NETWORK"

	// AuthenticationErrorReasonReasonUnderInvestigation
	//
	// The account is blocked and under investigation by the collections team. Please contact Google
	// for more information.
	//
	AuthenticationErrorReasonReasonUnderInvestigation = "UNDER_INVESTIGATION"

	// AuthenticationErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	AuthenticationErrorReasonReasonUnknown = "UNKNOWN"
)

// CollectionSizeErrorReasonReason
const (

	// CollectionSizeErrorReasonReasonTooLarge
	CollectionSizeErrorReasonReasonTooLarge = "TOO_LARGE"

	// CollectionSizeErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	CollectionSizeErrorReasonReasonUnknown = "UNKNOWN"
)

// CommonErrorReasonReason
//
// Describes reasons for common errors
const (

	// CommonErrorReasonReasonNotFound
	//
	// Indicates that an attempt was made to retrieve an entity that does not
	// exist.
	//
	CommonErrorReasonReasonNotFound = "NOT_FOUND"

	// CommonErrorReasonReasonAlreadyExists
	//
	// Indicates that an attempt was made to create an entity that already
	// exists.
	//
	CommonErrorReasonReasonAlreadyExists = "ALREADY_EXISTS"

	// CommonErrorReasonReasonNotApplicable
	//
	// Indicates that a value is not applicable for given use case.
	//
	CommonErrorReasonReasonNotApplicable = "NOT_APPLICABLE"

	// CommonErrorReasonReasonDuplicateObject
	//
	// Indicates that two elements in the collection were identical.
	//
	CommonErrorReasonReasonDuplicateObject = "DUPLICATE_OBJECT"

	// CommonErrorReasonReasonCannotUpdate
	//
	// Indicates that an attempt was made to change an immutable field.
	//
	CommonErrorReasonReasonCannotUpdate = "CANNOT_UPDATE"

	// CommonErrorReasonReasonUnsupportedOperation
	//
	// Indicates that the requested operation is not supported.
	//
	CommonErrorReasonReasonUnsupportedOperation = "UNSUPPORTED_OPERATION"

	// CommonErrorReasonReasonConcurrentModification
	//
	// Indicates that another request attempted to update the same data in the same network
	// at about the same time. Please wait and try the request again.
	//
	CommonErrorReasonReasonConcurrentModification = "CONCURRENT_MODIFICATION"

	// CommonErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	CommonErrorReasonReasonUnknown = "UNKNOWN"
)

// FeatureErrorReasonReason
const (

	// FeatureErrorReasonReasonMissingFeature
	//
	// A feature is being used that is not enabled on the current network.
	//
	FeatureErrorReasonReasonMissingFeature = "MISSING_FEATURE"

	// FeatureErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	FeatureErrorReasonReasonUnknown = "UNKNOWN"
)

// InternalApiErrorReasonReason
//
// The single reason for the internal API error.
const (

	// InternalApiErrorReasonReasonUnexpectedInternalApiError
	//
	// API encountered an unexpected internal error.
	//
	InternalApiErrorReasonReasonUnexpectedInternalApiError = "UNEXPECTED_INTERNAL_API_ERROR"

	// InternalApiErrorReasonReasonTransientError
	//
	// A temporary error occurred during the request. Please retry.
	//
	InternalApiErrorReasonReasonTransientError = "TRANSIENT_ERROR"

	// InternalApiErrorReasonReasonUnknown
	//
	// The cause of the error is not known or only defined in newer versions.
	//
	InternalApiErrorReasonReasonUnknown = "UNKNOWN"

	// InternalApiErrorReasonReasonDowntime
	//
	// The API is currently unavailable for a planned downtime.
	//
	InternalApiErrorReasonReasonDowntime = "DOWNTIME"

	// InternalApiErrorReasonReasonErrorGeneratingResponse
	//
	// Mutate succeeded but server was unable to build response. Client should not retry mutate.
	//
	InternalApiErrorReasonReasonErrorGeneratingResponse = "ERROR_GENERATING_RESPONSE"
)

// NotNullErrorReasonReason
//
// The reasons for the target error.
const (

	// NotNullErrorReasonReasonArg1Null
	//
	// Assuming that a method will not have more than 3 arguments, if it does,
	// return NULL
	//
	NotNullErrorReasonReasonArg1Null = "ARG1_NULL"

	// NotNullErrorReasonReasonArg2Null
	NotNullErrorReasonReasonArg2Null = "ARG2_NULL"

	// NotNullErrorReasonReasonArg3Null
	NotNullErrorReasonReasonArg3Null = "ARG3_NULL"

	// NotNullErrorReasonReasonNull
	NotNullErrorReasonReasonNull = "NULL"

	// NotNullErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	NotNullErrorReasonReasonUnknown = "UNKNOWN"
)

// ParseErrorReasonReason
//
// The reasons for the target error.
const (

	// ParseErrorReasonReasonUnparsable
	//
	// Indicates an error in parsing an attribute.
	//
	ParseErrorReasonReasonUnparsable = "UNPARSABLE"

	// ParseErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	ParseErrorReasonReasonUnknown = "UNKNOWN"
)

// PermissionErrorReasonReason
//
// Describes reasons for permission errors.
const (

	// PermissionErrorReasonReasonPermissionDenied
	//
	// User does not have the required permission for the request.
	//
	PermissionErrorReasonReasonPermissionDenied = "PERMISSION_DENIED"

	// PermissionErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	PermissionErrorReasonReasonUnknown = "UNKNOWN"
)

// PublisherQueryLanguageContextErrorReasonReason
//
// The reasons for the target error.
const (

	// PublisherQueryLanguageContextErrorReasonReasonUnexecutable
	//
	// Indicates that there was an error executing the PQL.
	//
	PublisherQueryLanguageContextErrorReasonReasonUnexecutable = "UNEXECUTABLE"

	// PublisherQueryLanguageContextErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	PublisherQueryLanguageContextErrorReasonReasonUnknown = "UNKNOWN"
)

// PublisherQueryLanguageSyntaxErrorReasonReason
//
// The reasons for the target error.
const (

	// PublisherQueryLanguageSyntaxErrorReasonReasonUnparsable
	//
	// Indicates that there was a PQL syntax error.
	//
	PublisherQueryLanguageSyntaxErrorReasonReasonUnparsable = "UNPARSABLE"

	// PublisherQueryLanguageSyntaxErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	PublisherQueryLanguageSyntaxErrorReasonReasonUnknown = "UNKNOWN"
)

// QuotaErrorReasonReason
const (

	// QuotaErrorReasonReasonExceededQuota
	//
	// The number of requests made per second is too high and has exceeded the
	// allowable limit. The recommended approach to handle this error is to wait
	// about 5 seconds and then retry the request. Note that this does not
	// guarantee the request will succeed. If it fails again, try increasing the
	// wait time.
	// <p>Another way to mitigate this error is to limit requests to 8 per second for Ad Manager
	// 360 accounts, or 2 per second for Ad Manager accounts. Once again
	// this does not guarantee that every request will succeed, but may help
	// reduce the number of times you receive this error.
	//
	QuotaErrorReasonReasonExceededQuota = "EXCEEDED_QUOTA"

	// QuotaErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	QuotaErrorReasonReasonUnknown = "UNKNOWN"

	// QuotaErrorReasonReasonReportJobLimit
	//
	// This user has exceeded the allowed number of new report requests per hour
	// (this includes both reports run via the UI and reports
	// run via {@link ReportService#runReportJob}).
	// The recommended approach to handle this error is to wait about 10 minutes
	// and then retry the request. Note that this does not guarantee the request
	// will succeed. If it fails again, try increasing the wait time.
	// <p>Another way to mitigate this error is to limit the number of new report
	// requests to 250 per hour per user. Once again, this does not guarantee that
	// every request will succeed, but may help reduce the number of times you
	// receive this error.
	//
	QuotaErrorReasonReasonReportJobLimit = "REPORT_JOB_LIMIT"

	// QuotaErrorReasonReasonSegmentPopulationLimit
	//
	// This network has exceeded the allowed number of identifiers uploaded within a 24 hour period.
	// The recommended approach to handle this error is to wait 30 minutes and then retry the
	// request. Note that this does not guarantee the request will succeed. If it fails again, try
	// increasing the wait time.
	//
	QuotaErrorReasonReasonSegmentPopulationLimit = "SEGMENT_POPULATION_LIMIT"
)

// RangeErrorReasonReason
const (

	// RangeErrorReasonReasonTooHigh
	RangeErrorReasonReasonTooHigh = "TOO_HIGH"

	// RangeErrorReasonReasonTooLow
	RangeErrorReasonReasonTooLow = "TOO_LOW"

	// RangeErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	RangeErrorReasonReasonUnknown = "UNKNOWN"
)

// RequiredCollectionErrorReasonReason
const (

	// RequiredCollectionErrorReasonReasonRequired
	//
	// A required collection is missing.
	//
	RequiredCollectionErrorReasonReasonRequired = "REQUIRED"

	// RequiredCollectionErrorReasonReasonTooLarge
	//
	// Collection size is too large.
	//
	RequiredCollectionErrorReasonReasonTooLarge = "TOO_LARGE"

	// RequiredCollectionErrorReasonReasonTooSmall
	//
	// Collection size is too small.
	//
	RequiredCollectionErrorReasonReasonTooSmall = "TOO_SMALL"

	// RequiredCollectionErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	RequiredCollectionErrorReasonReasonUnknown = "UNKNOWN"
)

// RequiredErrorReasonReason
//
// The reasons for the target error.
const (

	// RequiredErrorReasonReasonRequired
	//
	// Missing required field.
	//
	RequiredErrorReasonReasonRequired = "REQUIRED"
)

// ServerErrorReasonReason
//
// Describes reasons for server errors
const (

	// ServerErrorReasonReasonServerError
	//
	// Indicates that an unexpected error occured.
	//
	ServerErrorReasonReasonServerError = "SERVER_ERROR"

	// ServerErrorReasonReasonServerBusy
	//
	// Indicates that the server is currently experiencing a high load. Please
	// wait and try your request again.
	//
	ServerErrorReasonReasonServerBusy = "SERVER_BUSY"

	// ServerErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	ServerErrorReasonReasonUnknown = "UNKNOWN"
)

// StatementErrorReasonReason
const (

	// StatementErrorReasonReasonVariableNotBoundToValue
	//
	// A bind variable has not been bound to a value.
	//
	StatementErrorReasonReasonVariableNotBoundToValue = "VARIABLE_NOT_BOUND_TO_VALUE"

	// StatementErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	StatementErrorReasonReasonUnknown = "UNKNOWN"
)

// StringFormatErrorReasonReason
//
// The reasons for the target error.
const (

	// StringFormatErrorReasonReasonUnknown
	StringFormatErrorReasonReasonUnknown = "UNKNOWN"

	// StringFormatErrorReasonReasonIllegalChars
	//
	// The input string value contains disallowed characters.
	//
	StringFormatErrorReasonReasonIllegalChars = "ILLEGAL_CHARS"

	// StringFormatErrorReasonReasonInvalidFormat
	//
	// The input string value is invalid for the associated field.
	//
	StringFormatErrorReasonReasonInvalidFormat = "INVALID_FORMAT"
)

// StringLengthErrorReasonReason
const (

	// StringLengthErrorReasonReasonTooLong
	StringLengthErrorReasonReasonTooLong = "TOO_LONG"

	// StringLengthErrorReasonReasonTooShort
	StringLengthErrorReasonReasonTooShort = "TOO_SHORT"

	// StringLengthErrorReasonReasonUnknown
	//
	// The value returned if the actual value is not exposed by the requested API version.
	//
	StringLengthErrorReasonReasonUnknown = "UNKNOWN"
)

// createActivities
//
// Creates a new {@link Activity} objects.
type CreateActivities struct {
	Activities []*Activity `xml:"activities"`
}

// createActivitiesResponse
type CreateActivitiesResponse struct {
	Rval []*Activity `xml:"rval"`
}

// getActivitiesByStatement
//
// Gets an {@link ActivityPage} of {@link Activity} objects that satisfy the given {@link
// Statement#query}. The following fields are supported for filtering:
//
// <table>
// <tr>
// <th scope="col">PQL Property</th> <th scope="col">Object Property</th>
// </tr>
// <tr>
// <td>{@code id}</td>
// <td>{@link Activity#id}</td>
// </tr>
// <tr>
// <td>{@code name}</td>
// <td>{@link Activity#name}</td>
// </tr>
// <tr>
// <td>{@code expectedURL}</td>
// <td>{@link Activity#expectedURL}</td>
// </tr>
// <tr>
// <td>{@code status}</td>
// <td>{@link Activity#status}</td>
// </tr>
// <tr>
// <td>{@code activityGroupId}</td>
// <td>{@link Activity#activityGroupId}</td>
// </tr>
// </table>
type GetActivitiesByStatement struct {
	FilterStatement *Statement `xml:"filterStatement"`
}

// getActivitiesByStatementResponse
type GetActivitiesByStatementResponse struct {
	Rval *ActivityPage `xml:"rval"`
}

// updateActivities
//
// Updates the specified {@link Activity} objects.
type UpdateActivities struct {
	Activities []*Activity `xml:"activities"`
}

// updateActivitiesResponse
type UpdateActivitiesResponse struct {
	Rval []*Activity `xml:"rval"`
}

// RequestHeader
type RequestHeader struct {
}

// ResponseHeader
type ResponseHeader struct {
}