// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230630

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServersAdministrator_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties of an administrator.
	Properties *AdministratorProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersAdministrator_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-06-30"
func (administrator FlexibleServersAdministrator_Spec_ARM) GetAPIVersion() string {
	return "2023-06-30"
}

// GetName returns the Name of the resource
func (administrator *FlexibleServersAdministrator_Spec_ARM) GetName() string {
	return administrator.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/administrators"
func (administrator *FlexibleServersAdministrator_Spec_ARM) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/administrators"
}

// The properties of an administrator.
type AdministratorProperties_ARM struct {
	// AdministratorType: Type of the sever administrator.
	AdministratorType  *AdministratorProperties_AdministratorType_ARM `json:"administratorType,omitempty"`
	IdentityResourceId *string                                        `json:"identityResourceId,omitempty"`

	// Login: Login name of the server administrator.
	Login *string `json:"login,omitempty"`

	// Sid: SID (object ID) of the server administrator.
	Sid *string `json:"sid,omitempty" optionalConfigMapPair:"Sid"`

	// TenantId: Tenant ID of the administrator.
	TenantId *string `json:"tenantId,omitempty" optionalConfigMapPair:"TenantId"`
}

// +kubebuilder:validation:Enum={"ActiveDirectory"}
type AdministratorProperties_AdministratorType_ARM string

const AdministratorProperties_AdministratorType_ARM_ActiveDirectory = AdministratorProperties_AdministratorType_ARM("ActiveDirectory")

// Mapping from string to AdministratorProperties_AdministratorType_ARM
var administratorProperties_AdministratorType_ARM_Values = map[string]AdministratorProperties_AdministratorType_ARM{
	"activedirectory": AdministratorProperties_AdministratorType_ARM_ActiveDirectory,
}
