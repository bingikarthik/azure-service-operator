// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type DatabaseAccount_Spec_ARM struct {
	// Identity: Identity for the resource.
	Identity *ManagedServiceIdentity_ARM `json:"identity,omitempty"`

	// Kind: Indicates the type of database account. This can only be set at database account creation.
	Kind *DatabaseAccount_Kind_Spec `json:"kind,omitempty"`

	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties to create and update Azure Cosmos DB database accounts.
	Properties *DatabaseAccountCreateUpdateProperties_ARM `json:"properties,omitempty"`
	Tags       map[string]string                          `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccount_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (account DatabaseAccount_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (account *DatabaseAccount_Spec_ARM) GetName() string {
	return account.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts"
func (account *DatabaseAccount_Spec_ARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts"
}

// +kubebuilder:validation:Enum={"GlobalDocumentDB","MongoDB","Parse"}
type DatabaseAccount_Kind_Spec string

const (
	DatabaseAccount_Kind_Spec_GlobalDocumentDB = DatabaseAccount_Kind_Spec("GlobalDocumentDB")
	DatabaseAccount_Kind_Spec_MongoDB          = DatabaseAccount_Kind_Spec("MongoDB")
	DatabaseAccount_Kind_Spec_Parse            = DatabaseAccount_Kind_Spec("Parse")
)

// Properties to create and update Azure Cosmos DB database accounts.
type DatabaseAccountCreateUpdateProperties_ARM struct {
	// AnalyticalStorageConfiguration: Analytical storage specific properties.
	AnalyticalStorageConfiguration *AnalyticalStorageConfiguration_ARM `json:"analyticalStorageConfiguration,omitempty"`

	// ApiProperties: API specific properties. Currently, supported only for MongoDB API.
	ApiProperties *ApiProperties_ARM `json:"apiProperties,omitempty"`

	// BackupPolicy: The object representing the policy for taking backups on an account.
	BackupPolicy *BackupPolicy_ARM `json:"backupPolicy,omitempty"`

	// Capabilities: List of Cosmos DB capabilities for the account
	Capabilities []Capability_ARM `json:"capabilities,omitempty"`

	// ConnectorOffer: The cassandra connector offer type for the Cosmos DB database C* account.
	ConnectorOffer *ConnectorOffer `json:"connectorOffer,omitempty"`

	// ConsistencyPolicy: The consistency policy for the Cosmos DB account.
	ConsistencyPolicy *ConsistencyPolicy_ARM `json:"consistencyPolicy,omitempty"`

	// Cors: The CORS policy for the Cosmos DB database account.
	Cors []CorsPolicy_ARM `json:"cors,omitempty"`

	// DatabaseAccountOfferType: The offer type for the database
	DatabaseAccountOfferType *DatabaseAccountOfferType `json:"databaseAccountOfferType,omitempty"`

	// DefaultIdentity: The default identity for accessing key vault used in features like customer managed keys. The default
	// identity needs to be explicitly set by the users. It can be "FirstPartyIdentity", "SystemAssignedIdentity" and more.
	DefaultIdentity *string `json:"defaultIdentity,omitempty"`

	// DisableKeyBasedMetadataWriteAccess: Disable write operations on metadata resources (databases, containers, throughput)
	// via account keys
	DisableKeyBasedMetadataWriteAccess *bool `json:"disableKeyBasedMetadataWriteAccess,omitempty"`

	// EnableAnalyticalStorage: Flag to indicate whether to enable storage analytics.
	EnableAnalyticalStorage *bool `json:"enableAnalyticalStorage,omitempty"`

	// EnableAutomaticFailover: Enables automatic failover of the write region in the rare event that the region is unavailable
	// due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the
	// failover priorities configured for the account.
	EnableAutomaticFailover *bool `json:"enableAutomaticFailover,omitempty"`

	// EnableCassandraConnector: Enables the cassandra connector on the Cosmos DB C* account
	EnableCassandraConnector *bool `json:"enableCassandraConnector,omitempty"`

	// EnableFreeTier: Flag to indicate whether Free Tier is enabled.
	EnableFreeTier *bool `json:"enableFreeTier,omitempty"`

	// EnableMultipleWriteLocations: Enables the account to write in multiple locations
	EnableMultipleWriteLocations *bool `json:"enableMultipleWriteLocations,omitempty"`

	// IpRules: List of IpRules.
	IpRules []IpAddressOrRange_ARM `json:"ipRules,omitempty"`

	// IsVirtualNetworkFilterEnabled: Flag to indicate whether to enable/disable Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled *bool `json:"isVirtualNetworkFilterEnabled,omitempty"`

	// KeyVaultKeyUri: The URI of the key vault
	KeyVaultKeyUri *string `json:"keyVaultKeyUri,omitempty"`

	// Locations: An array that contains the georeplication locations enabled for the Cosmos DB account.
	Locations []Location_ARM `json:"locations,omitempty"`

	// NetworkAclBypass: Indicates what services are allowed to bypass firewall checks.
	NetworkAclBypass *NetworkAclBypass `json:"networkAclBypass,omitempty"`

	// NetworkAclBypassResourceIds: An array that contains the Resource Ids for Network Acl Bypass for the Cosmos DB account.
	NetworkAclBypassResourceIds []string `json:"networkAclBypassResourceIds,omitempty"`

	// PublicNetworkAccess: Whether requests from Public Network are allowed
	PublicNetworkAccess *PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// VirtualNetworkRules: List of Virtual Network ACL rules configured for the Cosmos DB account.
	VirtualNetworkRules []VirtualNetworkRule_ARM `json:"virtualNetworkRules,omitempty"`
}

// Identity for the resource.
type ManagedServiceIdentity_ARM struct {
	// Type: The type of identity used for the resource. The type 'SystemAssigned,UserAssigned' includes both an implicitly
	// created identity and a set of user assigned identities. The type 'None' will remove any identities from the service.
	Type                   *ManagedServiceIdentity_Type               `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// Analytical storage specific properties.
type AnalyticalStorageConfiguration_ARM struct {
	// SchemaType: Describes the types of schema for analytical storage.
	SchemaType *AnalyticalStorageSchemaType `json:"schemaType,omitempty"`
}

type ApiProperties_ARM struct {
	// ServerVersion: Describes the ServerVersion of an a MongoDB account.
	ServerVersion *ApiProperties_ServerVersion `json:"serverVersion,omitempty"`
}

type BackupPolicy_ARM struct {
	// Continuous: Mutually exclusive with all other properties
	Continuous *ContinuousModeBackupPolicy_ARM `json:"continuous,omitempty"`

	// Periodic: Mutually exclusive with all other properties
	Periodic *PeriodicModeBackupPolicy_ARM `json:"periodic,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BackupPolicy_ARM represents a discriminated union (JSON OneOf)
func (policy BackupPolicy_ARM) MarshalJSON() ([]byte, error) {
	if policy.Continuous != nil {
		return json.Marshal(policy.Continuous)
	}
	if policy.Periodic != nil {
		return json.Marshal(policy.Periodic)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BackupPolicy_ARM
func (policy *BackupPolicy_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["type"]
	if discriminator == "Continuous" {
		policy.Continuous = &ContinuousModeBackupPolicy_ARM{}
		return json.Unmarshal(data, policy.Continuous)
	}
	if discriminator == "Periodic" {
		policy.Periodic = &PeriodicModeBackupPolicy_ARM{}
		return json.Unmarshal(data, policy.Periodic)
	}

	// No error
	return nil
}

// Cosmos DB capability object
type Capability_ARM struct {
	// Name: Name of the Cosmos DB capability. For example, "name": "EnableCassandra". Current values also include
	// "EnableTable" and "EnableGremlin".
	Name *string `json:"name,omitempty"`
}

// The consistency policy for the Cosmos DB database account.
type ConsistencyPolicy_ARM struct {
	// DefaultConsistencyLevel: The default consistency level and configuration settings of the Cosmos DB account.
	DefaultConsistencyLevel *ConsistencyPolicy_DefaultConsistencyLevel `json:"defaultConsistencyLevel,omitempty"`

	// MaxIntervalInSeconds: When used with the Bounded Staleness consistency level, this value represents the time amount of
	// staleness (in seconds) tolerated. Accepted range for this value is 5 - 86400. Required when defaultConsistencyPolicy is
	// set to 'BoundedStaleness'.
	MaxIntervalInSeconds *int `json:"maxIntervalInSeconds,omitempty"`

	// MaxStalenessPrefix: When used with the Bounded Staleness consistency level, this value represents the number of stale
	// requests tolerated. Accepted range for this value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set
	// to 'BoundedStaleness'.
	MaxStalenessPrefix *int `json:"maxStalenessPrefix,omitempty"`
}

// The CORS policy for the Cosmos DB database account.
type CorsPolicy_ARM struct {
	// AllowedHeaders: The request headers that the origin domain may specify on the CORS request.
	AllowedHeaders *string `json:"allowedHeaders,omitempty"`

	// AllowedMethods: The methods (HTTP request verbs) that the origin domain may use for a CORS request.
	AllowedMethods *string `json:"allowedMethods,omitempty"`

	// AllowedOrigins: The origin domains that are permitted to make a request against the service via CORS.
	AllowedOrigins *string `json:"allowedOrigins,omitempty"`

	// ExposedHeaders: The response headers that may be sent in the response to the CORS request and exposed by the browser to
	// the request issuer.
	ExposedHeaders *string `json:"exposedHeaders,omitempty"`

	// MaxAgeInSeconds: The maximum amount time that a browser should cache the preflight OPTIONS request.
	MaxAgeInSeconds *int `json:"maxAgeInSeconds,omitempty"`
}

// IpAddressOrRange object
type IpAddressOrRange_ARM struct {
	// IpAddressOrRange: A single IPv4 address or a single IPv4 address range in CIDR format. Provided IPs must be
	// well-formatted and cannot be contained in one of the following ranges: 10.0.0.0/8, 100.64.0.0/10, 172.16.0.0/12,
	// 192.168.0.0/16, since these are not enforceable by the IP address filter. Example of valid inputs: “23.40.210.245”
	// or “23.40.210.0/8”.
	IpAddressOrRange *string `json:"ipAddressOrRange,omitempty"`
}

// A region in which the Azure Cosmos DB database account is deployed.
type Location_ARM struct {
	// FailoverPriority: The failover priority of the region. A failover priority of 0 indicates a write region. The maximum
	// value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the
	// regions in which the database account exists.
	FailoverPriority *int `json:"failoverPriority,omitempty"`

	// IsZoneRedundant: Flag to indicate whether or not this region is an AvailabilityZone region
	IsZoneRedundant *bool `json:"isZoneRedundant,omitempty"`

	// LocationName: The name of the region.
	LocationName *string `json:"locationName,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type ManagedServiceIdentity_Type string

const (
	ManagedServiceIdentity_Type_None                       = ManagedServiceIdentity_Type("None")
	ManagedServiceIdentity_Type_SystemAssigned             = ManagedServiceIdentity_Type("SystemAssigned")
	ManagedServiceIdentity_Type_SystemAssignedUserAssigned = ManagedServiceIdentity_Type("SystemAssigned,UserAssigned")
	ManagedServiceIdentity_Type_UserAssigned               = ManagedServiceIdentity_Type("UserAssigned")
)

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// Virtual Network ACL Rule object
type VirtualNetworkRule_ARM struct {
	Id *string `json:"id,omitempty"`

	// IgnoreMissingVNetServiceEndpoint: Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVNetServiceEndpoint *bool `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
}

type ContinuousModeBackupPolicy_ARM struct {
	Type ContinuousModeBackupPolicy_Type `json:"type,omitempty"`
}

type PeriodicModeBackupPolicy_ARM struct {
	// PeriodicModeProperties: Configuration values for periodic mode backup
	PeriodicModeProperties *PeriodicModeProperties_ARM   `json:"periodicModeProperties,omitempty"`
	Type                   PeriodicModeBackupPolicy_Type `json:"type,omitempty"`
}

// Configuration values for periodic mode backup
type PeriodicModeProperties_ARM struct {
	// BackupIntervalInMinutes: An integer representing the interval in minutes between two backups
	BackupIntervalInMinutes *int `json:"backupIntervalInMinutes,omitempty"`

	// BackupRetentionIntervalInHours: An integer representing the time (in hours) that each backup is retained
	BackupRetentionIntervalInHours *int `json:"backupRetentionIntervalInHours,omitempty"`
}
