// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230601preview

type FlexibleServersConfiguration_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. E.g.
	// "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The properties of a configuration.
	Properties *ConfigurationProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The properties of a configuration.
type ConfigurationProperties_STATUS_ARM struct {
	// AllowedValues: Allowed values of the configuration.
	AllowedValues *string `json:"allowedValues,omitempty"`

	// DataType: Data type of the configuration.
	DataType *ConfigurationProperties_DataType_STATUS_ARM `json:"dataType,omitempty"`

	// DefaultValue: Default value of the configuration.
	DefaultValue *string `json:"defaultValue,omitempty"`

	// Description: Description of the configuration.
	Description *string `json:"description,omitempty"`

	// DocumentationLink: Configuration documentation link.
	DocumentationLink *string `json:"documentationLink,omitempty"`

	// IsConfigPendingRestart: Configuration is pending restart or not.
	IsConfigPendingRestart *bool `json:"isConfigPendingRestart,omitempty"`

	// IsDynamicConfig: Configuration dynamic or static.
	IsDynamicConfig *bool `json:"isDynamicConfig,omitempty"`

	// IsReadOnly: Configuration read-only or not.
	IsReadOnly *bool `json:"isReadOnly,omitempty"`

	// Source: Source of the configuration.
	Source *string `json:"source,omitempty"`

	// Unit: Configuration unit.
	Unit *string `json:"unit,omitempty"`

	// Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}

type ConfigurationProperties_DataType_STATUS_ARM string

const (
	ConfigurationProperties_DataType_STATUS_ARM_Boolean     = ConfigurationProperties_DataType_STATUS_ARM("Boolean")
	ConfigurationProperties_DataType_STATUS_ARM_Enumeration = ConfigurationProperties_DataType_STATUS_ARM("Enumeration")
	ConfigurationProperties_DataType_STATUS_ARM_Integer     = ConfigurationProperties_DataType_STATUS_ARM("Integer")
	ConfigurationProperties_DataType_STATUS_ARM_Numeric     = ConfigurationProperties_DataType_STATUS_ARM("Numeric")
)

// Mapping from string to ConfigurationProperties_DataType_STATUS_ARM
var configurationProperties_DataType_STATUS_ARM_Values = map[string]ConfigurationProperties_DataType_STATUS_ARM{
	"boolean":     ConfigurationProperties_DataType_STATUS_ARM_Boolean,
	"enumeration": ConfigurationProperties_DataType_STATUS_ARM_Enumeration,
	"integer":     ConfigurationProperties_DataType_STATUS_ARM_Integer,
	"numeric":     ConfigurationProperties_DataType_STATUS_ARM_Numeric,
}
