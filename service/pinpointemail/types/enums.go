// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type BehaviorOnMxFailure string

// Enum values for BehaviorOnMxFailure
const (
	BehaviorOnMxFailureUse_default_value BehaviorOnMxFailure = "USE_DEFAULT_VALUE"
	BehaviorOnMxFailureReject_message    BehaviorOnMxFailure = "REJECT_MESSAGE"
)

type DeliverabilityDashboardAccountStatus string

// Enum values for DeliverabilityDashboardAccountStatus
const (
	DeliverabilityDashboardAccountStatusActive             DeliverabilityDashboardAccountStatus = "ACTIVE"
	DeliverabilityDashboardAccountStatusPending_expiration DeliverabilityDashboardAccountStatus = "PENDING_EXPIRATION"
	DeliverabilityDashboardAccountStatusDisabled           DeliverabilityDashboardAccountStatus = "DISABLED"
)

type DeliverabilityTestStatus string

// Enum values for DeliverabilityTestStatus
const (
	DeliverabilityTestStatusIn_progress DeliverabilityTestStatus = "IN_PROGRESS"
	DeliverabilityTestStatusCompleted   DeliverabilityTestStatus = "COMPLETED"
)

type DimensionValueSource string

// Enum values for DimensionValueSource
const (
	DimensionValueSourceMessage_tag  DimensionValueSource = "MESSAGE_TAG"
	DimensionValueSourceEmail_header DimensionValueSource = "EMAIL_HEADER"
	DimensionValueSourceLink_tag     DimensionValueSource = "LINK_TAG"
)

type DkimStatus string

// Enum values for DkimStatus
const (
	DkimStatusPending           DkimStatus = "PENDING"
	DkimStatusSuccess           DkimStatus = "SUCCESS"
	DkimStatusFailed            DkimStatus = "FAILED"
	DkimStatusTemporary_failure DkimStatus = "TEMPORARY_FAILURE"
	DkimStatusNot_started       DkimStatus = "NOT_STARTED"
)

type EventType string

// Enum values for EventType
const (
	EventTypeSend              EventType = "SEND"
	EventTypeReject            EventType = "REJECT"
	EventTypeBounce            EventType = "BOUNCE"
	EventTypeComplaint         EventType = "COMPLAINT"
	EventTypeDelivery          EventType = "DELIVERY"
	EventTypeOpen              EventType = "OPEN"
	EventTypeClick             EventType = "CLICK"
	EventTypeRendering_failure EventType = "RENDERING_FAILURE"
)

type IdentityType string

// Enum values for IdentityType
const (
	IdentityTypeEmail_address  IdentityType = "EMAIL_ADDRESS"
	IdentityTypeDomain         IdentityType = "DOMAIN"
	IdentityTypeManaged_domain IdentityType = "MANAGED_DOMAIN"
)

type MailFromDomainStatus string

// Enum values for MailFromDomainStatus
const (
	MailFromDomainStatusPending           MailFromDomainStatus = "PENDING"
	MailFromDomainStatusSuccess           MailFromDomainStatus = "SUCCESS"
	MailFromDomainStatusFailed            MailFromDomainStatus = "FAILED"
	MailFromDomainStatusTemporary_failure MailFromDomainStatus = "TEMPORARY_FAILURE"
)

type TlsPolicy string

// Enum values for TlsPolicy
const (
	TlsPolicyRequire  TlsPolicy = "REQUIRE"
	TlsPolicyOptional TlsPolicy = "OPTIONAL"
)

type WarmupStatus string

// Enum values for WarmupStatus
const (
	WarmupStatusIn_progress WarmupStatus = "IN_PROGRESS"
	WarmupStatusDone        WarmupStatus = "DONE"
)