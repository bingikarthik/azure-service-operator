// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RuleSet_Spec_ARM struct {
	Name string `json:"name,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RuleSet_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (ruleSet RuleSet_Spec_ARM) GetAPIVersion() string {
	return "2023-05-01"
}

// GetName returns the Name of the resource
func (ruleSet *RuleSet_Spec_ARM) GetName() string {
	return ruleSet.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles/ruleSets"
func (ruleSet *RuleSet_Spec_ARM) GetType() string {
	return "Microsoft.Cdn/profiles/ruleSets"
}
