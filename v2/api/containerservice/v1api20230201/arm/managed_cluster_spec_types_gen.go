// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ManagedCluster_Spec struct {
	ExtendedLocation *ExtendedLocation         `json:"extendedLocation,omitempty"`
	Identity         *ManagedClusterIdentity   `json:"identity,omitempty"`
	Location         *string                   `json:"location,omitempty"`
	Name             string                    `json:"name,omitempty"`
	Properties       *ManagedClusterProperties `json:"properties,omitempty"`
	Sku              *ManagedClusterSKU        `json:"sku,omitempty"`
	Tags             map[string]string         `json:"tags" serializationType:"explicitEmptyCollection"`
}

var _ genruntime.ARMResourceSpec = &ManagedCluster_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-02-01"
func (cluster ManagedCluster_Spec) GetAPIVersion() string {
	return "2023-02-01"
}

// GetName returns the Name of the resource
func (cluster *ManagedCluster_Spec) GetName() string {
	return cluster.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/managedClusters"
func (cluster *ManagedCluster_Spec) GetType() string {
	return "Microsoft.ContainerService/managedClusters"
}

type ExtendedLocation struct {
	Name *string               `json:"name,omitempty"`
	Type *ExtendedLocationType `json:"type,omitempty"`
}

type ManagedClusterIdentity struct {
	Type                   *ManagedClusterIdentity_Type           `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

type ManagedClusterProperties struct {
	AadProfile                *ManagedClusterAADProfile                     `json:"aadProfile,omitempty"`
	AddonProfiles             map[string]ManagedClusterAddonProfile         `json:"addonProfiles"`
	AgentPoolProfiles         []ManagedClusterAgentPoolProfile              `json:"agentPoolProfiles"`
	ApiServerAccessProfile    *ManagedClusterAPIServerAccessProfile         `json:"apiServerAccessProfile,omitempty"`
	AutoScalerProfile         *ManagedClusterProperties_AutoScalerProfile   `json:"autoScalerProfile,omitempty"`
	AutoUpgradeProfile        *ManagedClusterAutoUpgradeProfile             `json:"autoUpgradeProfile,omitempty"`
	AzureMonitorProfile       *ManagedClusterAzureMonitorProfile            `json:"azureMonitorProfile,omitempty"`
	DisableLocalAccounts      *bool                                         `json:"disableLocalAccounts,omitempty"`
	DiskEncryptionSetID       *string                                       `json:"diskEncryptionSetID,omitempty"`
	DnsPrefix                 *string                                       `json:"dnsPrefix,omitempty"`
	EnablePodSecurityPolicy   *bool                                         `json:"enablePodSecurityPolicy,omitempty"`
	EnableRBAC                *bool                                         `json:"enableRBAC,omitempty"`
	FqdnSubdomain             *string                                       `json:"fqdnSubdomain,omitempty"`
	HttpProxyConfig           *ManagedClusterHTTPProxyConfig                `json:"httpProxyConfig,omitempty"`
	IdentityProfile           map[string]UserAssignedIdentity               `json:"identityProfile"`
	KubernetesVersion         *string                                       `json:"kubernetesVersion,omitempty"`
	LinuxProfile              *ContainerServiceLinuxProfile                 `json:"linuxProfile,omitempty"`
	NetworkProfile            *ContainerServiceNetworkProfile               `json:"networkProfile,omitempty"`
	NodeResourceGroup         *string                                       `json:"nodeResourceGroup,omitempty"`
	OidcIssuerProfile         *ManagedClusterOIDCIssuerProfile              `json:"oidcIssuerProfile,omitempty"`
	PodIdentityProfile        *ManagedClusterPodIdentityProfile             `json:"podIdentityProfile,omitempty"`
	PrivateLinkResources      []PrivateLinkResource                         `json:"privateLinkResources"`
	PublicNetworkAccess       *ManagedClusterProperties_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
	SecurityProfile           *ManagedClusterSecurityProfile                `json:"securityProfile,omitempty"`
	ServicePrincipalProfile   *ManagedClusterServicePrincipalProfile        `json:"servicePrincipalProfile,omitempty"`
	StorageProfile            *ManagedClusterStorageProfile                 `json:"storageProfile,omitempty"`
	WindowsProfile            *ManagedClusterWindowsProfile                 `json:"windowsProfile,omitempty"`
	WorkloadAutoScalerProfile *ManagedClusterWorkloadAutoScalerProfile      `json:"workloadAutoScalerProfile,omitempty"`
}

type ManagedClusterSKU struct {
	Name *ManagedClusterSKU_Name `json:"name,omitempty"`
	Tier *ManagedClusterSKU_Tier `json:"tier,omitempty"`
}

type ContainerServiceLinuxProfile struct {
	AdminUsername *string                           `json:"adminUsername,omitempty"`
	Ssh           *ContainerServiceSshConfiguration `json:"ssh,omitempty"`
}

type ContainerServiceNetworkProfile struct {
	DnsServiceIP        *string                                           `json:"dnsServiceIP,omitempty"`
	DockerBridgeCidr    *string                                           `json:"dockerBridgeCidr,omitempty"`
	IpFamilies          []ContainerServiceNetworkProfile_IpFamilies       `json:"ipFamilies"`
	LoadBalancerProfile *ManagedClusterLoadBalancerProfile                `json:"loadBalancerProfile,omitempty"`
	LoadBalancerSku     *ContainerServiceNetworkProfile_LoadBalancerSku   `json:"loadBalancerSku,omitempty"`
	NatGatewayProfile   *ManagedClusterNATGatewayProfile                  `json:"natGatewayProfile,omitempty"`
	NetworkDataplane    *ContainerServiceNetworkProfile_NetworkDataplane  `json:"networkDataplane,omitempty"`
	NetworkMode         *ContainerServiceNetworkProfile_NetworkMode       `json:"networkMode,omitempty"`
	NetworkPlugin       *ContainerServiceNetworkProfile_NetworkPlugin     `json:"networkPlugin,omitempty"`
	NetworkPluginMode   *ContainerServiceNetworkProfile_NetworkPluginMode `json:"networkPluginMode,omitempty"`
	NetworkPolicy       *ContainerServiceNetworkProfile_NetworkPolicy     `json:"networkPolicy,omitempty"`
	OutboundType        *ContainerServiceNetworkProfile_OutboundType      `json:"outboundType,omitempty"`
	PodCidr             *string                                           `json:"podCidr,omitempty"`
	PodCidrs            []string                                          `json:"podCidrs"`
	ServiceCidr         *string                                           `json:"serviceCidr,omitempty"`
	ServiceCidrs        []string                                          `json:"serviceCidrs"`
}

// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocationType string

const ExtendedLocationType_EdgeZone = ExtendedLocationType("EdgeZone")

// Mapping from string to ExtendedLocationType
var extendedLocationType_Values = map[string]ExtendedLocationType{
	"edgezone": ExtendedLocationType_EdgeZone,
}

type ManagedClusterAADProfile struct {
	AdminGroupObjectIDs []string `json:"adminGroupObjectIDs"`
	ClientAppID         *string  `json:"clientAppID,omitempty"`
	EnableAzureRBAC     *bool    `json:"enableAzureRBAC,omitempty"`
	Managed             *bool    `json:"managed,omitempty"`
	ServerAppID         *string  `json:"serverAppID,omitempty"`
	ServerAppSecret     *string  `json:"serverAppSecret,omitempty"`
	TenantID            *string  `json:"tenantID,omitempty"`
}

type ManagedClusterAddonProfile struct {
	Config  map[string]string `json:"config"`
	Enabled *bool             `json:"enabled,omitempty"`
}

type ManagedClusterAgentPoolProfile struct {
	AvailabilityZones         []string                  `json:"availabilityZones"`
	Count                     *int                      `json:"count,omitempty"`
	CreationData              *CreationData             `json:"creationData,omitempty"`
	EnableAutoScaling         *bool                     `json:"enableAutoScaling,omitempty"`
	EnableEncryptionAtHost    *bool                     `json:"enableEncryptionAtHost,omitempty"`
	EnableFIPS                *bool                     `json:"enableFIPS,omitempty"`
	EnableNodePublicIP        *bool                     `json:"enableNodePublicIP,omitempty"`
	EnableUltraSSD            *bool                     `json:"enableUltraSSD,omitempty"`
	GpuInstanceProfile        *GPUInstanceProfile       `json:"gpuInstanceProfile,omitempty"`
	HostGroupID               *string                   `json:"hostGroupID,omitempty"`
	KubeletConfig             *KubeletConfig            `json:"kubeletConfig,omitempty"`
	KubeletDiskType           *KubeletDiskType          `json:"kubeletDiskType,omitempty"`
	LinuxOSConfig             *LinuxOSConfig            `json:"linuxOSConfig,omitempty"`
	MaxCount                  *int                      `json:"maxCount,omitempty"`
	MaxPods                   *int                      `json:"maxPods,omitempty"`
	MinCount                  *int                      `json:"minCount,omitempty"`
	Mode                      *AgentPoolMode            `json:"mode,omitempty"`
	Name                      *string                   `json:"name,omitempty"`
	NodeLabels                map[string]string         `json:"nodeLabels" serializationType:"explicitEmptyCollection"`
	NodePublicIPPrefixID      *string                   `json:"nodePublicIPPrefixID,omitempty"`
	NodeTaints                []string                  `json:"nodeTaints" serializationType:"explicitEmptyCollection"`
	OrchestratorVersion       *string                   `json:"orchestratorVersion,omitempty"`
	OsDiskSizeGB              *int                      `json:"osDiskSizeGB,omitempty"`
	OsDiskType                *OSDiskType               `json:"osDiskType,omitempty"`
	OsSKU                     *OSSKU                    `json:"osSKU,omitempty"`
	OsType                    *OSType                   `json:"osType,omitempty"`
	PodSubnetID               *string                   `json:"podSubnetID,omitempty"`
	PowerState                *PowerState               `json:"powerState,omitempty"`
	ProximityPlacementGroupID *string                   `json:"proximityPlacementGroupID,omitempty"`
	ScaleDownMode             *ScaleDownMode            `json:"scaleDownMode,omitempty"`
	ScaleSetEvictionPolicy    *ScaleSetEvictionPolicy   `json:"scaleSetEvictionPolicy,omitempty"`
	ScaleSetPriority          *ScaleSetPriority         `json:"scaleSetPriority,omitempty"`
	SpotMaxPrice              *float64                  `json:"spotMaxPrice,omitempty"`
	Tags                      map[string]string         `json:"tags" serializationType:"explicitEmptyCollection"`
	Type                      *AgentPoolType            `json:"type,omitempty"`
	UpgradeSettings           *AgentPoolUpgradeSettings `json:"upgradeSettings,omitempty"`
	VmSize                    *string                   `json:"vmSize,omitempty"`
	VnetSubnetID              *string                   `json:"vnetSubnetID,omitempty"`
	WorkloadRuntime           *WorkloadRuntime          `json:"workloadRuntime,omitempty"`
}

type ManagedClusterAPIServerAccessProfile struct {
	AuthorizedIPRanges             []string `json:"authorizedIPRanges"`
	DisableRunCommand              *bool    `json:"disableRunCommand,omitempty"`
	EnablePrivateCluster           *bool    `json:"enablePrivateCluster,omitempty"`
	EnablePrivateClusterPublicFQDN *bool    `json:"enablePrivateClusterPublicFQDN,omitempty"`
	PrivateDNSZone                 *string  `json:"privateDNSZone,omitempty"`
}

type ManagedClusterAutoUpgradeProfile struct {
	UpgradeChannel *ManagedClusterAutoUpgradeProfile_UpgradeChannel `json:"upgradeChannel,omitempty"`
}

type ManagedClusterAzureMonitorProfile struct {
	Metrics *ManagedClusterAzureMonitorProfileMetrics `json:"metrics,omitempty"`
}

type ManagedClusterHTTPProxyConfig struct {
	HttpProxy  *string  `json:"httpProxy,omitempty"`
	HttpsProxy *string  `json:"httpsProxy,omitempty"`
	NoProxy    []string `json:"noProxy"`
	TrustedCa  *string  `json:"trustedCa,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","UserAssigned"}
type ManagedClusterIdentity_Type string

const (
	ManagedClusterIdentity_Type_None           = ManagedClusterIdentity_Type("None")
	ManagedClusterIdentity_Type_SystemAssigned = ManagedClusterIdentity_Type("SystemAssigned")
	ManagedClusterIdentity_Type_UserAssigned   = ManagedClusterIdentity_Type("UserAssigned")
)

// Mapping from string to ManagedClusterIdentity_Type
var managedClusterIdentity_Type_Values = map[string]ManagedClusterIdentity_Type{
	"none":           ManagedClusterIdentity_Type_None,
	"systemassigned": ManagedClusterIdentity_Type_SystemAssigned,
	"userassigned":   ManagedClusterIdentity_Type_UserAssigned,
}

type ManagedClusterOIDCIssuerProfile struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type ManagedClusterPodIdentityProfile struct {
	AllowNetworkPluginKubenet      *bool                                `json:"allowNetworkPluginKubenet,omitempty"`
	Enabled                        *bool                                `json:"enabled,omitempty"`
	UserAssignedIdentities         []ManagedClusterPodIdentity          `json:"userAssignedIdentities"`
	UserAssignedIdentityExceptions []ManagedClusterPodIdentityException `json:"userAssignedIdentityExceptions"`
}

type ManagedClusterProperties_AutoScalerProfile struct {
	BalanceSimilarNodeGroups      *string                                              `json:"balance-similar-node-groups,omitempty"`
	Expander                      *ManagedClusterProperties_AutoScalerProfile_Expander `json:"expander,omitempty"`
	MaxEmptyBulkDelete            *string                                              `json:"max-empty-bulk-delete,omitempty"`
	MaxGracefulTerminationSec     *string                                              `json:"max-graceful-termination-sec,omitempty"`
	MaxNodeProvisionTime          *string                                              `json:"max-node-provision-time,omitempty"`
	MaxTotalUnreadyPercentage     *string                                              `json:"max-total-unready-percentage,omitempty"`
	NewPodScaleUpDelay            *string                                              `json:"new-pod-scale-up-delay,omitempty"`
	OkTotalUnreadyCount           *string                                              `json:"ok-total-unready-count,omitempty"`
	ScaleDownDelayAfterAdd        *string                                              `json:"scale-down-delay-after-add,omitempty"`
	ScaleDownDelayAfterDelete     *string                                              `json:"scale-down-delay-after-delete,omitempty"`
	ScaleDownDelayAfterFailure    *string                                              `json:"scale-down-delay-after-failure,omitempty"`
	ScaleDownUnneededTime         *string                                              `json:"scale-down-unneeded-time,omitempty"`
	ScaleDownUnreadyTime          *string                                              `json:"scale-down-unready-time,omitempty"`
	ScaleDownUtilizationThreshold *string                                              `json:"scale-down-utilization-threshold,omitempty"`
	ScanInterval                  *string                                              `json:"scan-interval,omitempty"`
	SkipNodesWithLocalStorage     *string                                              `json:"skip-nodes-with-local-storage,omitempty"`
	SkipNodesWithSystemPods       *string                                              `json:"skip-nodes-with-system-pods,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ManagedClusterProperties_PublicNetworkAccess string

const (
	ManagedClusterProperties_PublicNetworkAccess_Disabled = ManagedClusterProperties_PublicNetworkAccess("Disabled")
	ManagedClusterProperties_PublicNetworkAccess_Enabled  = ManagedClusterProperties_PublicNetworkAccess("Enabled")
)

// Mapping from string to ManagedClusterProperties_PublicNetworkAccess
var managedClusterProperties_PublicNetworkAccess_Values = map[string]ManagedClusterProperties_PublicNetworkAccess{
	"disabled": ManagedClusterProperties_PublicNetworkAccess_Disabled,
	"enabled":  ManagedClusterProperties_PublicNetworkAccess_Enabled,
}

type ManagedClusterSecurityProfile struct {
	AzureKeyVaultKms *AzureKeyVaultKms                              `json:"azureKeyVaultKms,omitempty"`
	Defender         *ManagedClusterSecurityProfileDefender         `json:"defender,omitempty"`
	ImageCleaner     *ManagedClusterSecurityProfileImageCleaner     `json:"imageCleaner,omitempty"`
	WorkloadIdentity *ManagedClusterSecurityProfileWorkloadIdentity `json:"workloadIdentity,omitempty"`
}

type ManagedClusterServicePrincipalProfile struct {
	ClientId *string `json:"clientId,omitempty"`
	Secret   *string `json:"secret,omitempty"`
}

// +kubebuilder:validation:Enum={"Base"}
type ManagedClusterSKU_Name string

const ManagedClusterSKU_Name_Base = ManagedClusterSKU_Name("Base")

// Mapping from string to ManagedClusterSKU_Name
var managedClusterSKU_Name_Values = map[string]ManagedClusterSKU_Name{
	"base": ManagedClusterSKU_Name_Base,
}

// +kubebuilder:validation:Enum={"Free","Standard"}
type ManagedClusterSKU_Tier string

const (
	ManagedClusterSKU_Tier_Free     = ManagedClusterSKU_Tier("Free")
	ManagedClusterSKU_Tier_Standard = ManagedClusterSKU_Tier("Standard")
)

// Mapping from string to ManagedClusterSKU_Tier
var managedClusterSKU_Tier_Values = map[string]ManagedClusterSKU_Tier{
	"free":     ManagedClusterSKU_Tier_Free,
	"standard": ManagedClusterSKU_Tier_Standard,
}

type ManagedClusterStorageProfile struct {
	BlobCSIDriver      *ManagedClusterStorageProfileBlobCSIDriver      `json:"blobCSIDriver,omitempty"`
	DiskCSIDriver      *ManagedClusterStorageProfileDiskCSIDriver      `json:"diskCSIDriver,omitempty"`
	FileCSIDriver      *ManagedClusterStorageProfileFileCSIDriver      `json:"fileCSIDriver,omitempty"`
	SnapshotController *ManagedClusterStorageProfileSnapshotController `json:"snapshotController,omitempty"`
}

type ManagedClusterWindowsProfile struct {
	AdminPassword  *string                                   `json:"adminPassword,omitempty"`
	AdminUsername  *string                                   `json:"adminUsername,omitempty"`
	EnableCSIProxy *bool                                     `json:"enableCSIProxy,omitempty"`
	GmsaProfile    *WindowsGmsaProfile                       `json:"gmsaProfile,omitempty"`
	LicenseType    *ManagedClusterWindowsProfile_LicenseType `json:"licenseType,omitempty"`
}

type ManagedClusterWorkloadAutoScalerProfile struct {
	Keda *ManagedClusterWorkloadAutoScalerProfileKeda `json:"keda,omitempty"`
}

type PrivateLinkResource struct {
	GroupId         *string  `json:"groupId,omitempty"`
	Id              *string  `json:"id,omitempty"`
	Name            *string  `json:"name,omitempty"`
	RequiredMembers []string `json:"requiredMembers"`
	Type            *string  `json:"type,omitempty"`
}

type UserAssignedIdentity struct {
	ClientId   *string `json:"clientId,omitempty"`
	ObjectId   *string `json:"objectId,omitempty"`
	ResourceId *string `json:"resourceId,omitempty"`
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
}

type AzureKeyVaultKms struct {
	Enabled               *bool                                   `json:"enabled,omitempty"`
	KeyId                 *string                                 `json:"keyId,omitempty"`
	KeyVaultNetworkAccess *AzureKeyVaultKms_KeyVaultNetworkAccess `json:"keyVaultNetworkAccess,omitempty"`
	KeyVaultResourceId    *string                                 `json:"keyVaultResourceId,omitempty"`
}

// +kubebuilder:validation:Enum={"IPv4","IPv6"}
type ContainerServiceNetworkProfile_IpFamilies string

const (
	ContainerServiceNetworkProfile_IpFamilies_IPv4 = ContainerServiceNetworkProfile_IpFamilies("IPv4")
	ContainerServiceNetworkProfile_IpFamilies_IPv6 = ContainerServiceNetworkProfile_IpFamilies("IPv6")
)

// Mapping from string to ContainerServiceNetworkProfile_IpFamilies
var containerServiceNetworkProfile_IpFamilies_Values = map[string]ContainerServiceNetworkProfile_IpFamilies{
	"ipv4": ContainerServiceNetworkProfile_IpFamilies_IPv4,
	"ipv6": ContainerServiceNetworkProfile_IpFamilies_IPv6,
}

// +kubebuilder:validation:Enum={"basic","standard"}
type ContainerServiceNetworkProfile_LoadBalancerSku string

const (
	ContainerServiceNetworkProfile_LoadBalancerSku_Basic    = ContainerServiceNetworkProfile_LoadBalancerSku("basic")
	ContainerServiceNetworkProfile_LoadBalancerSku_Standard = ContainerServiceNetworkProfile_LoadBalancerSku("standard")
)

// Mapping from string to ContainerServiceNetworkProfile_LoadBalancerSku
var containerServiceNetworkProfile_LoadBalancerSku_Values = map[string]ContainerServiceNetworkProfile_LoadBalancerSku{
	"basic":    ContainerServiceNetworkProfile_LoadBalancerSku_Basic,
	"standard": ContainerServiceNetworkProfile_LoadBalancerSku_Standard,
}

// +kubebuilder:validation:Enum={"azure","cilium"}
type ContainerServiceNetworkProfile_NetworkDataplane string

const (
	ContainerServiceNetworkProfile_NetworkDataplane_Azure  = ContainerServiceNetworkProfile_NetworkDataplane("azure")
	ContainerServiceNetworkProfile_NetworkDataplane_Cilium = ContainerServiceNetworkProfile_NetworkDataplane("cilium")
)

// Mapping from string to ContainerServiceNetworkProfile_NetworkDataplane
var containerServiceNetworkProfile_NetworkDataplane_Values = map[string]ContainerServiceNetworkProfile_NetworkDataplane{
	"azure":  ContainerServiceNetworkProfile_NetworkDataplane_Azure,
	"cilium": ContainerServiceNetworkProfile_NetworkDataplane_Cilium,
}

// +kubebuilder:validation:Enum={"bridge","transparent"}
type ContainerServiceNetworkProfile_NetworkMode string

const (
	ContainerServiceNetworkProfile_NetworkMode_Bridge      = ContainerServiceNetworkProfile_NetworkMode("bridge")
	ContainerServiceNetworkProfile_NetworkMode_Transparent = ContainerServiceNetworkProfile_NetworkMode("transparent")
)

// Mapping from string to ContainerServiceNetworkProfile_NetworkMode
var containerServiceNetworkProfile_NetworkMode_Values = map[string]ContainerServiceNetworkProfile_NetworkMode{
	"bridge":      ContainerServiceNetworkProfile_NetworkMode_Bridge,
	"transparent": ContainerServiceNetworkProfile_NetworkMode_Transparent,
}

// +kubebuilder:validation:Enum={"azure","kubenet","none"}
type ContainerServiceNetworkProfile_NetworkPlugin string

const (
	ContainerServiceNetworkProfile_NetworkPlugin_Azure   = ContainerServiceNetworkProfile_NetworkPlugin("azure")
	ContainerServiceNetworkProfile_NetworkPlugin_Kubenet = ContainerServiceNetworkProfile_NetworkPlugin("kubenet")
	ContainerServiceNetworkProfile_NetworkPlugin_None    = ContainerServiceNetworkProfile_NetworkPlugin("none")
)

// Mapping from string to ContainerServiceNetworkProfile_NetworkPlugin
var containerServiceNetworkProfile_NetworkPlugin_Values = map[string]ContainerServiceNetworkProfile_NetworkPlugin{
	"azure":   ContainerServiceNetworkProfile_NetworkPlugin_Azure,
	"kubenet": ContainerServiceNetworkProfile_NetworkPlugin_Kubenet,
	"none":    ContainerServiceNetworkProfile_NetworkPlugin_None,
}

// +kubebuilder:validation:Enum={"overlay"}
type ContainerServiceNetworkProfile_NetworkPluginMode string

const ContainerServiceNetworkProfile_NetworkPluginMode_Overlay = ContainerServiceNetworkProfile_NetworkPluginMode("overlay")

// Mapping from string to ContainerServiceNetworkProfile_NetworkPluginMode
var containerServiceNetworkProfile_NetworkPluginMode_Values = map[string]ContainerServiceNetworkProfile_NetworkPluginMode{
	"overlay": ContainerServiceNetworkProfile_NetworkPluginMode_Overlay,
}

// +kubebuilder:validation:Enum={"azure","calico","cilium"}
type ContainerServiceNetworkProfile_NetworkPolicy string

const (
	ContainerServiceNetworkProfile_NetworkPolicy_Azure  = ContainerServiceNetworkProfile_NetworkPolicy("azure")
	ContainerServiceNetworkProfile_NetworkPolicy_Calico = ContainerServiceNetworkProfile_NetworkPolicy("calico")
	ContainerServiceNetworkProfile_NetworkPolicy_Cilium = ContainerServiceNetworkProfile_NetworkPolicy("cilium")
)

// Mapping from string to ContainerServiceNetworkProfile_NetworkPolicy
var containerServiceNetworkProfile_NetworkPolicy_Values = map[string]ContainerServiceNetworkProfile_NetworkPolicy{
	"azure":  ContainerServiceNetworkProfile_NetworkPolicy_Azure,
	"calico": ContainerServiceNetworkProfile_NetworkPolicy_Calico,
	"cilium": ContainerServiceNetworkProfile_NetworkPolicy_Cilium,
}

// +kubebuilder:validation:Enum={"loadBalancer","managedNATGateway","userAssignedNATGateway","userDefinedRouting"}
type ContainerServiceNetworkProfile_OutboundType string

const (
	ContainerServiceNetworkProfile_OutboundType_LoadBalancer           = ContainerServiceNetworkProfile_OutboundType("loadBalancer")
	ContainerServiceNetworkProfile_OutboundType_ManagedNATGateway      = ContainerServiceNetworkProfile_OutboundType("managedNATGateway")
	ContainerServiceNetworkProfile_OutboundType_UserAssignedNATGateway = ContainerServiceNetworkProfile_OutboundType("userAssignedNATGateway")
	ContainerServiceNetworkProfile_OutboundType_UserDefinedRouting     = ContainerServiceNetworkProfile_OutboundType("userDefinedRouting")
)

// Mapping from string to ContainerServiceNetworkProfile_OutboundType
var containerServiceNetworkProfile_OutboundType_Values = map[string]ContainerServiceNetworkProfile_OutboundType{
	"loadbalancer":           ContainerServiceNetworkProfile_OutboundType_LoadBalancer,
	"managednatgateway":      ContainerServiceNetworkProfile_OutboundType_ManagedNATGateway,
	"userassignednatgateway": ContainerServiceNetworkProfile_OutboundType_UserAssignedNATGateway,
	"userdefinedrouting":     ContainerServiceNetworkProfile_OutboundType_UserDefinedRouting,
}

type ContainerServiceSshConfiguration struct {
	PublicKeys []ContainerServiceSshPublicKey `json:"publicKeys"`
}

// +kubebuilder:validation:Enum={"node-image","none","patch","rapid","stable"}
type ManagedClusterAutoUpgradeProfile_UpgradeChannel string

const (
	ManagedClusterAutoUpgradeProfile_UpgradeChannel_NodeImage = ManagedClusterAutoUpgradeProfile_UpgradeChannel("node-image")
	ManagedClusterAutoUpgradeProfile_UpgradeChannel_None      = ManagedClusterAutoUpgradeProfile_UpgradeChannel("none")
	ManagedClusterAutoUpgradeProfile_UpgradeChannel_Patch     = ManagedClusterAutoUpgradeProfile_UpgradeChannel("patch")
	ManagedClusterAutoUpgradeProfile_UpgradeChannel_Rapid     = ManagedClusterAutoUpgradeProfile_UpgradeChannel("rapid")
	ManagedClusterAutoUpgradeProfile_UpgradeChannel_Stable    = ManagedClusterAutoUpgradeProfile_UpgradeChannel("stable")
)

// Mapping from string to ManagedClusterAutoUpgradeProfile_UpgradeChannel
var managedClusterAutoUpgradeProfile_UpgradeChannel_Values = map[string]ManagedClusterAutoUpgradeProfile_UpgradeChannel{
	"node-image": ManagedClusterAutoUpgradeProfile_UpgradeChannel_NodeImage,
	"none":       ManagedClusterAutoUpgradeProfile_UpgradeChannel_None,
	"patch":      ManagedClusterAutoUpgradeProfile_UpgradeChannel_Patch,
	"rapid":      ManagedClusterAutoUpgradeProfile_UpgradeChannel_Rapid,
	"stable":     ManagedClusterAutoUpgradeProfile_UpgradeChannel_Stable,
}

type ManagedClusterAzureMonitorProfileMetrics struct {
	Enabled          *bool                                              `json:"enabled,omitempty"`
	KubeStateMetrics *ManagedClusterAzureMonitorProfileKubeStateMetrics `json:"kubeStateMetrics,omitempty"`
}

type ManagedClusterLoadBalancerProfile struct {
	AllocatedOutboundPorts              *int                                                  `json:"allocatedOutboundPorts,omitempty"`
	EffectiveOutboundIPs                []ResourceReference                                   `json:"effectiveOutboundIPs"`
	EnableMultipleStandardLoadBalancers *bool                                                 `json:"enableMultipleStandardLoadBalancers,omitempty"`
	IdleTimeoutInMinutes                *int                                                  `json:"idleTimeoutInMinutes,omitempty"`
	ManagedOutboundIPs                  *ManagedClusterLoadBalancerProfile_ManagedOutboundIPs `json:"managedOutboundIPs,omitempty"`
	OutboundIPPrefixes                  *ManagedClusterLoadBalancerProfile_OutboundIPPrefixes `json:"outboundIPPrefixes,omitempty"`
	OutboundIPs                         *ManagedClusterLoadBalancerProfile_OutboundIPs        `json:"outboundIPs,omitempty"`
}

type ManagedClusterNATGatewayProfile struct {
	EffectiveOutboundIPs     []ResourceReference                     `json:"effectiveOutboundIPs"`
	IdleTimeoutInMinutes     *int                                    `json:"idleTimeoutInMinutes,omitempty"`
	ManagedOutboundIPProfile *ManagedClusterManagedOutboundIPProfile `json:"managedOutboundIPProfile,omitempty"`
}

type ManagedClusterPodIdentity struct {
	BindingSelector *string               `json:"bindingSelector,omitempty"`
	Identity        *UserAssignedIdentity `json:"identity,omitempty"`
	Name            *string               `json:"name,omitempty"`
	Namespace       *string               `json:"namespace,omitempty"`
}

type ManagedClusterPodIdentityException struct {
	Name      *string           `json:"name,omitempty"`
	Namespace *string           `json:"namespace,omitempty"`
	PodLabels map[string]string `json:"podLabels"`
}

// +kubebuilder:validation:Enum={"least-waste","most-pods","priority","random"}
type ManagedClusterProperties_AutoScalerProfile_Expander string

const (
	ManagedClusterProperties_AutoScalerProfile_Expander_LeastWaste = ManagedClusterProperties_AutoScalerProfile_Expander("least-waste")
	ManagedClusterProperties_AutoScalerProfile_Expander_MostPods   = ManagedClusterProperties_AutoScalerProfile_Expander("most-pods")
	ManagedClusterProperties_AutoScalerProfile_Expander_Priority   = ManagedClusterProperties_AutoScalerProfile_Expander("priority")
	ManagedClusterProperties_AutoScalerProfile_Expander_Random     = ManagedClusterProperties_AutoScalerProfile_Expander("random")
)

// Mapping from string to ManagedClusterProperties_AutoScalerProfile_Expander
var managedClusterProperties_AutoScalerProfile_Expander_Values = map[string]ManagedClusterProperties_AutoScalerProfile_Expander{
	"least-waste": ManagedClusterProperties_AutoScalerProfile_Expander_LeastWaste,
	"most-pods":   ManagedClusterProperties_AutoScalerProfile_Expander_MostPods,
	"priority":    ManagedClusterProperties_AutoScalerProfile_Expander_Priority,
	"random":      ManagedClusterProperties_AutoScalerProfile_Expander_Random,
}

type ManagedClusterSecurityProfileDefender struct {
	LogAnalyticsWorkspaceResourceId *string                                                  `json:"logAnalyticsWorkspaceResourceId,omitempty"`
	SecurityMonitoring              *ManagedClusterSecurityProfileDefenderSecurityMonitoring `json:"securityMonitoring,omitempty"`
}

type ManagedClusterSecurityProfileImageCleaner struct {
	Enabled       *bool `json:"enabled,omitempty"`
	IntervalHours *int  `json:"intervalHours,omitempty"`
}

type ManagedClusterSecurityProfileWorkloadIdentity struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type ManagedClusterStorageProfileBlobCSIDriver struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type ManagedClusterStorageProfileDiskCSIDriver struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type ManagedClusterStorageProfileFileCSIDriver struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type ManagedClusterStorageProfileSnapshotController struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// +kubebuilder:validation:Enum={"None","Windows_Server"}
type ManagedClusterWindowsProfile_LicenseType string

const (
	ManagedClusterWindowsProfile_LicenseType_None           = ManagedClusterWindowsProfile_LicenseType("None")
	ManagedClusterWindowsProfile_LicenseType_Windows_Server = ManagedClusterWindowsProfile_LicenseType("Windows_Server")
)

// Mapping from string to ManagedClusterWindowsProfile_LicenseType
var managedClusterWindowsProfile_LicenseType_Values = map[string]ManagedClusterWindowsProfile_LicenseType{
	"none":           ManagedClusterWindowsProfile_LicenseType_None,
	"windows_server": ManagedClusterWindowsProfile_LicenseType_Windows_Server,
}

type ManagedClusterWorkloadAutoScalerProfileKeda struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type WindowsGmsaProfile struct {
	DnsServer      *string `json:"dnsServer,omitempty"`
	Enabled        *bool   `json:"enabled,omitempty"`
	RootDomainName *string `json:"rootDomainName,omitempty"`
}

// +kubebuilder:validation:Enum={"Private","Public"}
type AzureKeyVaultKms_KeyVaultNetworkAccess string

const (
	AzureKeyVaultKms_KeyVaultNetworkAccess_Private = AzureKeyVaultKms_KeyVaultNetworkAccess("Private")
	AzureKeyVaultKms_KeyVaultNetworkAccess_Public  = AzureKeyVaultKms_KeyVaultNetworkAccess("Public")
)

// Mapping from string to AzureKeyVaultKms_KeyVaultNetworkAccess
var azureKeyVaultKms_KeyVaultNetworkAccess_Values = map[string]AzureKeyVaultKms_KeyVaultNetworkAccess{
	"private": AzureKeyVaultKms_KeyVaultNetworkAccess_Private,
	"public":  AzureKeyVaultKms_KeyVaultNetworkAccess_Public,
}

type ContainerServiceSshPublicKey struct {
	KeyData *string `json:"keyData,omitempty"`
}

type ManagedClusterAzureMonitorProfileKubeStateMetrics struct {
	MetricAnnotationsAllowList *string `json:"metricAnnotationsAllowList,omitempty"`
	MetricLabelsAllowlist      *string `json:"metricLabelsAllowlist,omitempty"`
}

type ManagedClusterLoadBalancerProfile_ManagedOutboundIPs struct {
	Count     *int `json:"count,omitempty"`
	CountIPv6 *int `json:"countIPv6,omitempty"`
}

type ManagedClusterLoadBalancerProfile_OutboundIPPrefixes struct {
	PublicIPPrefixes []ResourceReference `json:"publicIPPrefixes"`
}

type ManagedClusterLoadBalancerProfile_OutboundIPs struct {
	PublicIPs []ResourceReference `json:"publicIPs"`
}

type ManagedClusterManagedOutboundIPProfile struct {
	Count *int `json:"count,omitempty"`
}

type ManagedClusterSecurityProfileDefenderSecurityMonitoring struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type ResourceReference struct {
	Id *string `json:"id,omitempty"`
}
