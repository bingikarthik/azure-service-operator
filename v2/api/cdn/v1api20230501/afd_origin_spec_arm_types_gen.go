// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type AfdOrigin_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: The JSON object that contains the properties of the origin.
	Properties *AFDOriginProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &AfdOrigin_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (origin AfdOrigin_Spec_ARM) GetAPIVersion() string {
	return "2023-05-01"
}

// GetName returns the Name of the resource
func (origin *AfdOrigin_Spec_ARM) GetName() string {
	return origin.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles/originGroups/origins"
func (origin *AfdOrigin_Spec_ARM) GetType() string {
	return "Microsoft.Cdn/profiles/originGroups/origins"
}

// The JSON object that contains the properties of the origin.
type AFDOriginProperties_ARM struct {
	// AzureOrigin: Resource reference to the Azure origin resource.
	AzureOrigin *ResourceReference_ARM `json:"azureOrigin,omitempty"`

	// EnabledState: Whether to enable health probes to be made against backends defined under backendPools. Health probes can
	// only be disabled if there is a single enabled backend in single enabled backend pool.
	EnabledState *AFDOriginProperties_EnabledState_ARM `json:"enabledState,omitempty"`

	// EnforceCertificateNameCheck: Whether to enable certificate name check at origin level
	EnforceCertificateNameCheck *bool `json:"enforceCertificateNameCheck,omitempty"`

	// HostName: The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be
	// unique across all origins in an endpoint.
	HostName *string `json:"hostName,omitempty"`

	// HttpPort: The value of the HTTP port. Must be between 1 and 65535.
	HttpPort *int `json:"httpPort,omitempty"`

	// HttpsPort: The value of the HTTPS port. Must be between 1 and 65535.
	HttpsPort *int `json:"httpsPort,omitempty"`

	// OriginHostHeader: The host header value sent to the origin with each request. If you leave this blank, the request
	// hostname determines this value. Azure Front Door origins, such as Web Apps, Blob Storage, and Cloud Services require
	// this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
	OriginHostHeader *string `json:"originHostHeader,omitempty"`

	// Priority: Priority of origin in given origin group for load balancing. Higher priorities will not be used for load
	// balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority *int `json:"priority,omitempty"`

	// SharedPrivateLinkResource: The properties of the private link resource for private origin.
	SharedPrivateLinkResource *SharedPrivateLinkResourceProperties_ARM `json:"sharedPrivateLinkResource,omitempty"`

	// Weight: Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
	Weight *int `json:"weight,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type AFDOriginProperties_EnabledState_ARM string

const (
	AFDOriginProperties_EnabledState_ARM_Disabled = AFDOriginProperties_EnabledState_ARM("Disabled")
	AFDOriginProperties_EnabledState_ARM_Enabled  = AFDOriginProperties_EnabledState_ARM("Enabled")
)

// Mapping from string to AFDOriginProperties_EnabledState_ARM
var aFDOriginProperties_EnabledState_ARM_Values = map[string]AFDOriginProperties_EnabledState_ARM{
	"disabled": AFDOriginProperties_EnabledState_ARM_Disabled,
	"enabled":  AFDOriginProperties_EnabledState_ARM_Enabled,
}

// Describes the properties of an existing Shared Private Link Resource to use when connecting to a private origin.
type SharedPrivateLinkResourceProperties_ARM struct {
	// GroupId: The group id from the provider of resource the shared private link resource is for.
	GroupId *string `json:"groupId,omitempty"`

	// PrivateLink: The resource id of the resource the shared private link resource is for.
	PrivateLink *ResourceReference_ARM `json:"privateLink,omitempty"`

	// PrivateLinkLocation: The location of the shared private link resource
	PrivateLinkLocation *string `json:"privateLinkLocation,omitempty"`

	// RequestMessage: The request message for requesting approval of the shared private link resource.
	RequestMessage *string `json:"requestMessage,omitempty"`

	// Status: Status of the shared private link resource. Can be Pending, Approved, Rejected, Disconnected, or Timeout.
	Status *SharedPrivateLinkResourceProperties_Status_ARM `json:"status,omitempty"`
}

// +kubebuilder:validation:Enum={"Approved","Disconnected","Pending","Rejected","Timeout"}
type SharedPrivateLinkResourceProperties_Status_ARM string

const (
	SharedPrivateLinkResourceProperties_Status_ARM_Approved     = SharedPrivateLinkResourceProperties_Status_ARM("Approved")
	SharedPrivateLinkResourceProperties_Status_ARM_Disconnected = SharedPrivateLinkResourceProperties_Status_ARM("Disconnected")
	SharedPrivateLinkResourceProperties_Status_ARM_Pending      = SharedPrivateLinkResourceProperties_Status_ARM("Pending")
	SharedPrivateLinkResourceProperties_Status_ARM_Rejected     = SharedPrivateLinkResourceProperties_Status_ARM("Rejected")
	SharedPrivateLinkResourceProperties_Status_ARM_Timeout      = SharedPrivateLinkResourceProperties_Status_ARM("Timeout")
)

// Mapping from string to SharedPrivateLinkResourceProperties_Status_ARM
var sharedPrivateLinkResourceProperties_Status_ARM_Values = map[string]SharedPrivateLinkResourceProperties_Status_ARM{
	"approved":     SharedPrivateLinkResourceProperties_Status_ARM_Approved,
	"disconnected": SharedPrivateLinkResourceProperties_Status_ARM_Disconnected,
	"pending":      SharedPrivateLinkResourceProperties_Status_ARM_Pending,
	"rejected":     SharedPrivateLinkResourceProperties_Status_ARM_Rejected,
	"timeout":      SharedPrivateLinkResourceProperties_Status_ARM_Timeout,
}
