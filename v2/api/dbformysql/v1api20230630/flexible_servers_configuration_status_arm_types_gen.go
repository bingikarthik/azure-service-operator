// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230630

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

	// CurrentValue: Current value of the configuration.
	CurrentValue *string `json:"currentValue,omitempty"`

	// DataType: Data type of the configuration.
	DataType *string `json:"dataType,omitempty"`

	// DefaultValue: Default value of the configuration.
	DefaultValue *string `json:"defaultValue,omitempty"`

	// Description: Description of the configuration.
	Description *string `json:"description,omitempty"`

	// DocumentationLink: The link used to get the document from community or Azure site.
	DocumentationLink *string `json:"documentationLink,omitempty"`

	// IsConfigPendingRestart: If is the configuration pending restart or not.
	IsConfigPendingRestart *ConfigurationProperties_IsConfigPendingRestart_STATUS_ARM `json:"isConfigPendingRestart,omitempty"`

	// IsDynamicConfig: If is the configuration dynamic.
	IsDynamicConfig *ConfigurationProperties_IsDynamicConfig_STATUS_ARM `json:"isDynamicConfig,omitempty"`

	// IsReadOnly: If is the configuration read only.
	IsReadOnly *ConfigurationProperties_IsReadOnly_STATUS_ARM `json:"isReadOnly,omitempty"`

	// Source: Source of the configuration.
	Source *ConfigurationProperties_Source_STATUS_ARM `json:"source,omitempty"`

	// Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}

type ConfigurationProperties_IsConfigPendingRestart_STATUS_ARM string

const (
	ConfigurationProperties_IsConfigPendingRestart_STATUS_ARM_False = ConfigurationProperties_IsConfigPendingRestart_STATUS_ARM("False")
	ConfigurationProperties_IsConfigPendingRestart_STATUS_ARM_True  = ConfigurationProperties_IsConfigPendingRestart_STATUS_ARM("True")
)

// Mapping from string to ConfigurationProperties_IsConfigPendingRestart_STATUS_ARM
var configurationProperties_IsConfigPendingRestart_STATUS_ARM_Values = map[string]ConfigurationProperties_IsConfigPendingRestart_STATUS_ARM{
	"false": ConfigurationProperties_IsConfigPendingRestart_STATUS_ARM_False,
	"true":  ConfigurationProperties_IsConfigPendingRestart_STATUS_ARM_True,
}

type ConfigurationProperties_IsDynamicConfig_STATUS_ARM string

const (
	ConfigurationProperties_IsDynamicConfig_STATUS_ARM_False = ConfigurationProperties_IsDynamicConfig_STATUS_ARM("False")
	ConfigurationProperties_IsDynamicConfig_STATUS_ARM_True  = ConfigurationProperties_IsDynamicConfig_STATUS_ARM("True")
)

// Mapping from string to ConfigurationProperties_IsDynamicConfig_STATUS_ARM
var configurationProperties_IsDynamicConfig_STATUS_ARM_Values = map[string]ConfigurationProperties_IsDynamicConfig_STATUS_ARM{
	"false": ConfigurationProperties_IsDynamicConfig_STATUS_ARM_False,
	"true":  ConfigurationProperties_IsDynamicConfig_STATUS_ARM_True,
}

type ConfigurationProperties_IsReadOnly_STATUS_ARM string

const (
	ConfigurationProperties_IsReadOnly_STATUS_ARM_False = ConfigurationProperties_IsReadOnly_STATUS_ARM("False")
	ConfigurationProperties_IsReadOnly_STATUS_ARM_True  = ConfigurationProperties_IsReadOnly_STATUS_ARM("True")
)

// Mapping from string to ConfigurationProperties_IsReadOnly_STATUS_ARM
var configurationProperties_IsReadOnly_STATUS_ARM_Values = map[string]ConfigurationProperties_IsReadOnly_STATUS_ARM{
	"false": ConfigurationProperties_IsReadOnly_STATUS_ARM_False,
	"true":  ConfigurationProperties_IsReadOnly_STATUS_ARM_True,
}

type ConfigurationProperties_Source_STATUS_ARM string

const (
	ConfigurationProperties_Source_STATUS_ARM_SystemDefault = ConfigurationProperties_Source_STATUS_ARM("system-default")
	ConfigurationProperties_Source_STATUS_ARM_UserOverride  = ConfigurationProperties_Source_STATUS_ARM("user-override")
)

// Mapping from string to ConfigurationProperties_Source_STATUS_ARM
var configurationProperties_Source_STATUS_ARM_Values = map[string]ConfigurationProperties_Source_STATUS_ARM{
	"system-default": ConfigurationProperties_Source_STATUS_ARM_SystemDefault,
	"user-override":  ConfigurationProperties_Source_STATUS_ARM_UserOverride,
}
