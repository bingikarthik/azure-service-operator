// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ServersOutboundFirewallRule_Spec struct {
	Name string `json:"name,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ServersOutboundFirewallRule_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule ServersOutboundFirewallRule_Spec) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (rule *ServersOutboundFirewallRule_Spec) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/outboundFirewallRules"
func (rule *ServersOutboundFirewallRule_Spec) GetType() string {
	return "Microsoft.Sql/servers/outboundFirewallRules"
}
