// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220301

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Site_Spec. Use v1api20220301.Site_Spec instead
type Site_Spec_ARM struct {
	ExtendedLocation *ExtendedLocation_ARM       `json:"extendedLocation,omitempty"`
	Identity         *ManagedServiceIdentity_ARM `json:"identity,omitempty"`
	Kind             *string                     `json:"kind,omitempty"`
	Location         *string                     `json:"location,omitempty"`
	Name             string                      `json:"name,omitempty"`
	Properties       *Site_Properties_Spec_ARM   `json:"properties,omitempty"`
	Tags             map[string]string           `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Site_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-03-01"
func (site Site_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (site *Site_Spec_ARM) GetName() string {
	return site.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Web/sites"
func (site *Site_Spec_ARM) GetType() string {
	return "Microsoft.Web/sites"
}

// Deprecated version of ManagedServiceIdentity. Use v1api20220301.ManagedServiceIdentity instead
type ManagedServiceIdentity_ARM struct {
	Type                   *ManagedServiceIdentity_Type               `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of Site_Properties_Spec. Use v1api20220301.Site_Properties_Spec instead
type Site_Properties_Spec_ARM struct {
	ClientAffinityEnabled      *bool                                `json:"clientAffinityEnabled,omitempty"`
	ClientCertEnabled          *bool                                `json:"clientCertEnabled,omitempty"`
	ClientCertExclusionPaths   *string                              `json:"clientCertExclusionPaths,omitempty"`
	ClientCertMode             *Site_Properties_ClientCertMode_Spec `json:"clientCertMode,omitempty"`
	CloningInfo                *CloningInfo_ARM                     `json:"cloningInfo,omitempty"`
	ContainerSize              *int                                 `json:"containerSize,omitempty"`
	CustomDomainVerificationId *string                              `json:"customDomainVerificationId,omitempty"`
	DailyMemoryTimeQuota       *int                                 `json:"dailyMemoryTimeQuota,omitempty"`
	Enabled                    *bool                                `json:"enabled,omitempty"`
	HostNameSslStates          []HostNameSslState_ARM               `json:"hostNameSslStates,omitempty"`
	HostNamesDisabled          *bool                                `json:"hostNamesDisabled,omitempty"`
	HostingEnvironmentProfile  *HostingEnvironmentProfile_ARM       `json:"hostingEnvironmentProfile,omitempty"`
	HttpsOnly                  *bool                                `json:"httpsOnly,omitempty"`
	HyperV                     *bool                                `json:"hyperV,omitempty"`
	IsXenon                    *bool                                `json:"isXenon,omitempty"`
	KeyVaultReferenceIdentity  *string                              `json:"keyVaultReferenceIdentity,omitempty"`
	PublicNetworkAccess        *string                              `json:"publicNetworkAccess,omitempty"`
	RedundancyMode             *Site_Properties_RedundancyMode_Spec `json:"redundancyMode,omitempty"`
	Reserved                   *bool                                `json:"reserved,omitempty"`
	ScmSiteAlsoStopped         *bool                                `json:"scmSiteAlsoStopped,omitempty"`
	ServerFarmId               *string                              `json:"serverFarmId,omitempty"`
	SiteConfig                 *SiteConfig_ARM                      `json:"siteConfig,omitempty"`
	StorageAccountRequired     *bool                                `json:"storageAccountRequired,omitempty"`
	VirtualNetworkSubnetId     *string                              `json:"virtualNetworkSubnetId,omitempty"`
	VnetContentShareEnabled    *bool                                `json:"vnetContentShareEnabled,omitempty"`
	VnetImagePullEnabled       *bool                                `json:"vnetImagePullEnabled,omitempty"`
	VnetRouteAllEnabled        *bool                                `json:"vnetRouteAllEnabled,omitempty"`
}

// Deprecated version of CloningInfo. Use v1api20220301.CloningInfo instead
type CloningInfo_ARM struct {
	AppSettingsOverrides      map[string]string `json:"appSettingsOverrides,omitempty"`
	CloneCustomHostNames      *bool             `json:"cloneCustomHostNames,omitempty"`
	CloneSourceControl        *bool             `json:"cloneSourceControl,omitempty"`
	ConfigureLoadBalancing    *bool             `json:"configureLoadBalancing,omitempty"`
	CorrelationId             *string           `json:"correlationId,omitempty"`
	HostingEnvironment        *string           `json:"hostingEnvironment,omitempty"`
	Overwrite                 *bool             `json:"overwrite,omitempty"`
	SourceWebAppId            *string           `json:"sourceWebAppId,omitempty"`
	SourceWebAppLocation      *string           `json:"sourceWebAppLocation,omitempty"`
	TrafficManagerProfileId   *string           `json:"trafficManagerProfileId,omitempty"`
	TrafficManagerProfileName *string           `json:"trafficManagerProfileName,omitempty"`
}

// Deprecated version of HostNameSslState. Use v1api20220301.HostNameSslState instead
type HostNameSslState_ARM struct {
	HostType   *HostNameSslState_HostType `json:"hostType,omitempty"`
	Name       *string                    `json:"name,omitempty"`
	SslState   *HostNameSslState_SslState `json:"sslState,omitempty"`
	Thumbprint *string                    `json:"thumbprint,omitempty"`
	ToUpdate   *bool                      `json:"toUpdate,omitempty"`
	VirtualIP  *string                    `json:"virtualIP,omitempty"`
}

// Deprecated version of ManagedServiceIdentity_Type. Use v1api20220301.ManagedServiceIdentity_Type instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type ManagedServiceIdentity_Type string

const (
	ManagedServiceIdentity_Type_None                       = ManagedServiceIdentity_Type("None")
	ManagedServiceIdentity_Type_SystemAssigned             = ManagedServiceIdentity_Type("SystemAssigned")
	ManagedServiceIdentity_Type_SystemAssignedUserAssigned = ManagedServiceIdentity_Type("SystemAssigned, UserAssigned")
	ManagedServiceIdentity_Type_UserAssigned               = ManagedServiceIdentity_Type("UserAssigned")
)

// Deprecated version of SiteConfig. Use v1api20220301.SiteConfig instead
type SiteConfig_ARM struct {
	AcrUseManagedIdentityCreds             *bool                                `json:"acrUseManagedIdentityCreds,omitempty"`
	AcrUserManagedIdentityID               *string                              `json:"acrUserManagedIdentityID,omitempty"`
	AlwaysOn                               *bool                                `json:"alwaysOn,omitempty"`
	ApiDefinition                          *ApiDefinitionInfo_ARM               `json:"apiDefinition,omitempty"`
	ApiManagementConfig                    *ApiManagementConfig_ARM             `json:"apiManagementConfig,omitempty"`
	AppCommandLine                         *string                              `json:"appCommandLine,omitempty"`
	AppSettings                            []NameValuePair_ARM                  `json:"appSettings,omitempty"`
	AutoHealEnabled                        *bool                                `json:"autoHealEnabled,omitempty"`
	AutoHealRules                          *AutoHealRules_ARM                   `json:"autoHealRules,omitempty"`
	AutoSwapSlotName                       *string                              `json:"autoSwapSlotName,omitempty"`
	AzureStorageAccounts                   map[string]AzureStorageInfoValue_ARM `json:"azureStorageAccounts,omitempty"`
	ConnectionStrings                      []ConnStringInfo_ARM                 `json:"connectionStrings,omitempty"`
	Cors                                   *CorsSettings_ARM                    `json:"cors,omitempty"`
	DefaultDocuments                       []string                             `json:"defaultDocuments,omitempty"`
	DetailedErrorLoggingEnabled            *bool                                `json:"detailedErrorLoggingEnabled,omitempty"`
	DocumentRoot                           *string                              `json:"documentRoot,omitempty"`
	Experiments                            *Experiments_ARM                     `json:"experiments,omitempty"`
	FtpsState                              *SiteConfig_FtpsState                `json:"ftpsState,omitempty"`
	FunctionAppScaleLimit                  *int                                 `json:"functionAppScaleLimit,omitempty"`
	FunctionsRuntimeScaleMonitoringEnabled *bool                                `json:"functionsRuntimeScaleMonitoringEnabled,omitempty"`
	HandlerMappings                        []HandlerMapping_ARM                 `json:"handlerMappings,omitempty"`
	HealthCheckPath                        *string                              `json:"healthCheckPath,omitempty"`
	Http20Enabled                          *bool                                `json:"http20Enabled,omitempty"`
	HttpLoggingEnabled                     *bool                                `json:"httpLoggingEnabled,omitempty"`
	IpSecurityRestrictions                 []IpSecurityRestriction_ARM          `json:"ipSecurityRestrictions,omitempty"`
	JavaContainer                          *string                              `json:"javaContainer,omitempty"`
	JavaContainerVersion                   *string                              `json:"javaContainerVersion,omitempty"`
	JavaVersion                            *string                              `json:"javaVersion,omitempty"`
	KeyVaultReferenceIdentity              *string                              `json:"keyVaultReferenceIdentity,omitempty"`
	Limits                                 *SiteLimits_ARM                      `json:"limits,omitempty"`
	LinuxFxVersion                         *string                              `json:"linuxFxVersion,omitempty"`
	LoadBalancing                          *SiteConfig_LoadBalancing            `json:"loadBalancing,omitempty"`
	LocalMySqlEnabled                      *bool                                `json:"localMySqlEnabled,omitempty"`
	LogsDirectorySizeLimit                 *int                                 `json:"logsDirectorySizeLimit,omitempty"`
	ManagedPipelineMode                    *SiteConfig_ManagedPipelineMode      `json:"managedPipelineMode,omitempty"`
	ManagedServiceIdentityId               *int                                 `json:"managedServiceIdentityId,omitempty"`
	MinTlsVersion                          *SiteConfig_MinTlsVersion            `json:"minTlsVersion,omitempty"`
	MinimumElasticInstanceCount            *int                                 `json:"minimumElasticInstanceCount,omitempty"`
	NetFrameworkVersion                    *string                              `json:"netFrameworkVersion,omitempty"`
	NodeVersion                            *string                              `json:"nodeVersion,omitempty"`
	NumberOfWorkers                        *int                                 `json:"numberOfWorkers,omitempty"`
	PhpVersion                             *string                              `json:"phpVersion,omitempty"`
	PowerShellVersion                      *string                              `json:"powerShellVersion,omitempty"`
	PreWarmedInstanceCount                 *int                                 `json:"preWarmedInstanceCount,omitempty"`
	PublicNetworkAccess                    *string                              `json:"publicNetworkAccess,omitempty"`
	PublishingUsername                     *string                              `json:"publishingUsername,omitempty"`
	Push                                   *PushSettings_ARM                    `json:"push,omitempty"`
	PythonVersion                          *string                              `json:"pythonVersion,omitempty"`
	RemoteDebuggingEnabled                 *bool                                `json:"remoteDebuggingEnabled,omitempty"`
	RemoteDebuggingVersion                 *string                              `json:"remoteDebuggingVersion,omitempty"`
	RequestTracingEnabled                  *bool                                `json:"requestTracingEnabled,omitempty"`
	RequestTracingExpirationTime           *string                              `json:"requestTracingExpirationTime,omitempty"`
	ScmIpSecurityRestrictions              []IpSecurityRestriction_ARM          `json:"scmIpSecurityRestrictions,omitempty"`
	ScmIpSecurityRestrictionsUseMain       *bool                                `json:"scmIpSecurityRestrictionsUseMain,omitempty"`
	ScmMinTlsVersion                       *SiteConfig_ScmMinTlsVersion         `json:"scmMinTlsVersion,omitempty"`
	ScmType                                *SiteConfig_ScmType                  `json:"scmType,omitempty"`
	TracingOptions                         *string                              `json:"tracingOptions,omitempty"`
	Use32BitWorkerProcess                  *bool                                `json:"use32BitWorkerProcess,omitempty"`
	VirtualApplications                    []VirtualApplication_ARM             `json:"virtualApplications,omitempty"`
	VnetName                               *string                              `json:"vnetName,omitempty"`
	VnetPrivatePortsCount                  *int                                 `json:"vnetPrivatePortsCount,omitempty"`
	VnetRouteAllEnabled                    *bool                                `json:"vnetRouteAllEnabled,omitempty"`
	WebSocketsEnabled                      *bool                                `json:"webSocketsEnabled,omitempty"`
	WebsiteTimeZone                        *string                              `json:"websiteTimeZone,omitempty"`
	WindowsFxVersion                       *string                              `json:"windowsFxVersion,omitempty"`
	XManagedServiceIdentityId              *int                                 `json:"xManagedServiceIdentityId,omitempty"`
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// Deprecated version of ApiDefinitionInfo. Use v1api20220301.ApiDefinitionInfo instead
type ApiDefinitionInfo_ARM struct {
	Url *string `json:"url,omitempty"`
}

// Deprecated version of ApiManagementConfig. Use v1api20220301.ApiManagementConfig instead
type ApiManagementConfig_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of AutoHealRules. Use v1api20220301.AutoHealRules instead
type AutoHealRules_ARM struct {
	Actions  *AutoHealActions_ARM  `json:"actions,omitempty"`
	Triggers *AutoHealTriggers_ARM `json:"triggers,omitempty"`
}

// Deprecated version of AzureStorageInfoValue. Use v1api20220301.AzureStorageInfoValue instead
type AzureStorageInfoValue_ARM struct {
	AccessKey   *string                     `json:"accessKey,omitempty"`
	AccountName *string                     `json:"accountName,omitempty"`
	MountPath   *string                     `json:"mountPath,omitempty"`
	ShareName   *string                     `json:"shareName,omitempty"`
	Type        *AzureStorageInfoValue_Type `json:"type,omitempty"`
}

// Deprecated version of ConnStringInfo. Use v1api20220301.ConnStringInfo instead
type ConnStringInfo_ARM struct {
	ConnectionString *string              `json:"connectionString,omitempty"`
	Name             *string              `json:"name,omitempty"`
	Type             *ConnStringInfo_Type `json:"type,omitempty"`
}

// Deprecated version of CorsSettings. Use v1api20220301.CorsSettings instead
type CorsSettings_ARM struct {
	AllowedOrigins     []string `json:"allowedOrigins,omitempty"`
	SupportCredentials *bool    `json:"supportCredentials,omitempty"`
}

// Deprecated version of Experiments. Use v1api20220301.Experiments instead
type Experiments_ARM struct {
	RampUpRules []RampUpRule_ARM `json:"rampUpRules,omitempty"`
}

// Deprecated version of HandlerMapping. Use v1api20220301.HandlerMapping instead
type HandlerMapping_ARM struct {
	Arguments       *string `json:"arguments,omitempty"`
	Extension       *string `json:"extension,omitempty"`
	ScriptProcessor *string `json:"scriptProcessor,omitempty"`
}

// Deprecated version of IpSecurityRestriction. Use v1api20220301.IpSecurityRestriction instead
type IpSecurityRestriction_ARM struct {
	Action               *string                    `json:"action,omitempty"`
	Description          *string                    `json:"description,omitempty"`
	Headers              map[string][]string        `json:"headers,omitempty"`
	IpAddress            *string                    `json:"ipAddress,omitempty"`
	Name                 *string                    `json:"name,omitempty"`
	Priority             *int                       `json:"priority,omitempty"`
	SubnetMask           *string                    `json:"subnetMask,omitempty"`
	SubnetTrafficTag     *int                       `json:"subnetTrafficTag,omitempty"`
	Tag                  *IpSecurityRestriction_Tag `json:"tag,omitempty"`
	VnetSubnetResourceId *string                    `json:"vnetSubnetResourceId,omitempty"`
	VnetTrafficTag       *int                       `json:"vnetTrafficTag,omitempty"`
}

// Deprecated version of NameValuePair. Use v1api20220301.NameValuePair instead
type NameValuePair_ARM struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Deprecated version of PushSettings. Use v1api20220301.PushSettings instead
type PushSettings_ARM struct {
	Kind       *string                      `json:"kind,omitempty"`
	Properties *PushSettings_Properties_ARM `json:"properties,omitempty"`
}

// Deprecated version of SiteLimits. Use v1api20220301.SiteLimits instead
type SiteLimits_ARM struct {
	MaxDiskSizeInMb  *int     `json:"maxDiskSizeInMb,omitempty"`
	MaxMemoryInMb    *int     `json:"maxMemoryInMb,omitempty"`
	MaxPercentageCpu *float64 `json:"maxPercentageCpu,omitempty"`
}

// Deprecated version of VirtualApplication. Use v1api20220301.VirtualApplication instead
type VirtualApplication_ARM struct {
	PhysicalPath       *string                `json:"physicalPath,omitempty"`
	PreloadEnabled     *bool                  `json:"preloadEnabled,omitempty"`
	VirtualDirectories []VirtualDirectory_ARM `json:"virtualDirectories,omitempty"`
	VirtualPath        *string                `json:"virtualPath,omitempty"`
}

// Deprecated version of AutoHealActions. Use v1api20220301.AutoHealActions instead
type AutoHealActions_ARM struct {
	ActionType              *AutoHealActions_ActionType `json:"actionType,omitempty"`
	CustomAction            *AutoHealCustomAction_ARM   `json:"customAction,omitempty"`
	MinProcessExecutionTime *string                     `json:"minProcessExecutionTime,omitempty"`
}

// Deprecated version of AutoHealTriggers. Use v1api20220301.AutoHealTriggers instead
type AutoHealTriggers_ARM struct {
	PrivateBytesInKB     *int                               `json:"privateBytesInKB,omitempty"`
	Requests             *RequestsBasedTrigger_ARM          `json:"requests,omitempty"`
	SlowRequests         *SlowRequestsBasedTrigger_ARM      `json:"slowRequests,omitempty"`
	SlowRequestsWithPath []SlowRequestsBasedTrigger_ARM     `json:"slowRequestsWithPath,omitempty"`
	StatusCodes          []StatusCodesBasedTrigger_ARM      `json:"statusCodes,omitempty"`
	StatusCodesRange     []StatusCodesRangeBasedTrigger_ARM `json:"statusCodesRange,omitempty"`
}

// Deprecated version of PushSettings_Properties. Use v1api20220301.PushSettings_Properties instead
type PushSettings_Properties_ARM struct {
	DynamicTagsJson   *string `json:"dynamicTagsJson,omitempty"`
	IsPushEnabled     *bool   `json:"isPushEnabled,omitempty"`
	TagWhitelistJson  *string `json:"tagWhitelistJson,omitempty"`
	TagsRequiringAuth *string `json:"tagsRequiringAuth,omitempty"`
}

// Deprecated version of RampUpRule. Use v1api20220301.RampUpRule instead
type RampUpRule_ARM struct {
	ActionHostName            *string  `json:"actionHostName,omitempty"`
	ChangeDecisionCallbackUrl *string  `json:"changeDecisionCallbackUrl,omitempty"`
	ChangeIntervalInMinutes   *int     `json:"changeIntervalInMinutes,omitempty"`
	ChangeStep                *float64 `json:"changeStep,omitempty"`
	MaxReroutePercentage      *float64 `json:"maxReroutePercentage,omitempty"`
	MinReroutePercentage      *float64 `json:"minReroutePercentage,omitempty"`
	Name                      *string  `json:"name,omitempty"`
	ReroutePercentage         *float64 `json:"reroutePercentage,omitempty"`
}

// Deprecated version of VirtualDirectory. Use v1api20220301.VirtualDirectory instead
type VirtualDirectory_ARM struct {
	PhysicalPath *string `json:"physicalPath,omitempty"`
	VirtualPath  *string `json:"virtualPath,omitempty"`
}

// Deprecated version of AutoHealCustomAction. Use v1api20220301.AutoHealCustomAction instead
type AutoHealCustomAction_ARM struct {
	Exe        *string `json:"exe,omitempty"`
	Parameters *string `json:"parameters,omitempty"`
}

// Deprecated version of RequestsBasedTrigger. Use v1api20220301.RequestsBasedTrigger instead
type RequestsBasedTrigger_ARM struct {
	Count        *int    `json:"count,omitempty"`
	TimeInterval *string `json:"timeInterval,omitempty"`
}

// Deprecated version of SlowRequestsBasedTrigger. Use v1api20220301.SlowRequestsBasedTrigger instead
type SlowRequestsBasedTrigger_ARM struct {
	Count        *int    `json:"count,omitempty"`
	Path         *string `json:"path,omitempty"`
	TimeInterval *string `json:"timeInterval,omitempty"`
	TimeTaken    *string `json:"timeTaken,omitempty"`
}

// Deprecated version of StatusCodesBasedTrigger. Use v1api20220301.StatusCodesBasedTrigger instead
type StatusCodesBasedTrigger_ARM struct {
	Count        *int    `json:"count,omitempty"`
	Path         *string `json:"path,omitempty"`
	Status       *int    `json:"status,omitempty"`
	SubStatus    *int    `json:"subStatus,omitempty"`
	TimeInterval *string `json:"timeInterval,omitempty"`
	Win32Status  *int    `json:"win32Status,omitempty"`
}

// Deprecated version of StatusCodesRangeBasedTrigger. Use v1api20220301.StatusCodesRangeBasedTrigger instead
type StatusCodesRangeBasedTrigger_ARM struct {
	Count        *int    `json:"count,omitempty"`
	Path         *string `json:"path,omitempty"`
	StatusCodes  *string `json:"statusCodes,omitempty"`
	TimeInterval *string `json:"timeInterval,omitempty"`
}
