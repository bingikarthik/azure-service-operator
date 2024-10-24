// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230131

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FederatedIdentityCredential_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties associated with the federated identity credential.
	Properties *FederatedIdentityCredentialProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FederatedIdentityCredential_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-01-31"
func (credential FederatedIdentityCredential_Spec_ARM) GetAPIVersion() string {
	return "2023-01-31"
}

// GetName returns the Name of the resource
func (credential *FederatedIdentityCredential_Spec_ARM) GetName() string {
	return credential.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"
func (credential *FederatedIdentityCredential_Spec_ARM) GetType() string {
	return "Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"
}

// The properties associated with a federated identity credential.
type FederatedIdentityCredentialProperties_ARM struct {
	// Audiences: The list of audiences that can appear in the issued token.
	Audiences []string `json:"audiences,omitempty"`

	// Issuer: The URL of the issuer to be trusted.
	Issuer *string `json:"issuer,omitempty" optionalConfigMapPair:"Issuer"`

	// Subject: The identifier of the external identity.
	Subject *string `json:"subject,omitempty" optionalConfigMapPair:"Subject"`
}
