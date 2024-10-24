// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ApiVersionSet_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: API VersionSet contract properties.
	Properties *ApiVersionSetContractProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ApiVersionSet_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01-preview"
func (versionSet ApiVersionSet_Spec_ARM) GetAPIVersion() string {
	return "2023-05-01-preview"
}

// GetName returns the Name of the resource
func (versionSet *ApiVersionSet_Spec_ARM) GetName() string {
	return versionSet.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/apiVersionSets"
func (versionSet *ApiVersionSet_Spec_ARM) GetType() string {
	return "Microsoft.ApiManagement/service/apiVersionSets"
}

// Properties of an API Version Set.
type ApiVersionSetContractProperties_ARM struct {
	// Description: Description of API Version Set.
	Description *string `json:"description,omitempty"`

	// DisplayName: Name of API Version Set
	DisplayName *string `json:"displayName,omitempty"`

	// VersionHeaderName: Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
	VersionHeaderName *string `json:"versionHeaderName,omitempty"`

	// VersionQueryName: Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
	VersionQueryName *string `json:"versionQueryName,omitempty"`

	// VersioningScheme: An value that determines where the API Version identifier will be located in a HTTP request.
	VersioningScheme *ApiVersionSetContractProperties_VersioningScheme_ARM `json:"versioningScheme,omitempty"`
}

// +kubebuilder:validation:Enum={"Header","Query","Segment"}
type ApiVersionSetContractProperties_VersioningScheme_ARM string

const (
	ApiVersionSetContractProperties_VersioningScheme_ARM_Header  = ApiVersionSetContractProperties_VersioningScheme_ARM("Header")
	ApiVersionSetContractProperties_VersioningScheme_ARM_Query   = ApiVersionSetContractProperties_VersioningScheme_ARM("Query")
	ApiVersionSetContractProperties_VersioningScheme_ARM_Segment = ApiVersionSetContractProperties_VersioningScheme_ARM("Segment")
)

// Mapping from string to ApiVersionSetContractProperties_VersioningScheme_ARM
var apiVersionSetContractProperties_VersioningScheme_ARM_Values = map[string]ApiVersionSetContractProperties_VersioningScheme_ARM{
	"header":  ApiVersionSetContractProperties_VersioningScheme_ARM_Header,
	"query":   ApiVersionSetContractProperties_VersioningScheme_ARM_Query,
	"segment": ApiVersionSetContractProperties_VersioningScheme_ARM_Segment,
}
