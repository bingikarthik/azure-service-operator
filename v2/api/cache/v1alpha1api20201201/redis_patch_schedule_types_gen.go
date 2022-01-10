// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20201201storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_patchSchedules
type RedisPatchSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisPatchSchedules_Spec  `json:"spec,omitempty"`
	Status            RedisPatchSchedule_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisPatchSchedule{}

// GetConditions returns the conditions of the resource
func (redisPatchSchedule *RedisPatchSchedule) GetConditions() conditions.Conditions {
	return redisPatchSchedule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (redisPatchSchedule *RedisPatchSchedule) SetConditions(conditions conditions.Conditions) {
	redisPatchSchedule.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisPatchSchedule{}

// ConvertFrom populates our RedisPatchSchedule from the provided hub RedisPatchSchedule
func (redisPatchSchedule *RedisPatchSchedule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1alpha1api20201201storage.RedisPatchSchedule)
	if !ok {
		return fmt.Errorf("expected storage:cache/v1alpha1api20201201storage/RedisPatchSchedule but received %T instead", hub)
	}

	return redisPatchSchedule.AssignPropertiesFromRedisPatchSchedule(source)
}

// ConvertTo populates the provided hub RedisPatchSchedule from our RedisPatchSchedule
func (redisPatchSchedule *RedisPatchSchedule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20201201storage.RedisPatchSchedule)
	if !ok {
		return fmt.Errorf("expected storage:cache/v1alpha1api20201201storage/RedisPatchSchedule but received %T instead", hub)
	}

	return redisPatchSchedule.AssignPropertiesToRedisPatchSchedule(destination)
}

// +kubebuilder:webhook:path=/mutate-cache-azure-com-v1alpha1api20201201-redispatchschedule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redispatchschedules,verbs=create;update,versions=v1alpha1api20201201,name=default.v1alpha1api20201201.redispatchschedules.cache.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &RedisPatchSchedule{}

// Default applies defaults to the RedisPatchSchedule resource
func (redisPatchSchedule *RedisPatchSchedule) Default() {
	redisPatchSchedule.defaultImpl()
	var temp interface{} = redisPatchSchedule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the RedisPatchSchedule resource
func (redisPatchSchedule *RedisPatchSchedule) defaultImpl() {}

var _ genruntime.KubernetesResource = &RedisPatchSchedule{}

// AzureName returns the Azure name of the resource (always "default")
func (redisPatchSchedule *RedisPatchSchedule) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (redisPatchSchedule RedisPatchSchedule) GetAPIVersion() string {
	return "2020-12-01"
}

// GetResourceKind returns the kind of the resource
func (redisPatchSchedule *RedisPatchSchedule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (redisPatchSchedule *RedisPatchSchedule) GetSpec() genruntime.ConvertibleSpec {
	return &redisPatchSchedule.Spec
}

// GetStatus returns the status of this resource
func (redisPatchSchedule *RedisPatchSchedule) GetStatus() genruntime.ConvertibleStatus {
	return &redisPatchSchedule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/patchSchedules"
func (redisPatchSchedule *RedisPatchSchedule) GetType() string {
	return "Microsoft.Cache/redis/patchSchedules"
}

// NewEmptyStatus returns a new empty (blank) status
func (redisPatchSchedule *RedisPatchSchedule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisPatchSchedule_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (redisPatchSchedule *RedisPatchSchedule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(redisPatchSchedule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  redisPatchSchedule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (redisPatchSchedule *RedisPatchSchedule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisPatchSchedule_Status); ok {
		redisPatchSchedule.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisPatchSchedule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	redisPatchSchedule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-cache-azure-com-v1alpha1api20201201-redispatchschedule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redispatchschedules,verbs=create;update,versions=v1alpha1api20201201,name=validate.v1alpha1api20201201.redispatchschedules.cache.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &RedisPatchSchedule{}

// ValidateCreate validates the creation of the resource
func (redisPatchSchedule *RedisPatchSchedule) ValidateCreate() error {
	validations := redisPatchSchedule.createValidations()
	var temp interface{} = redisPatchSchedule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (redisPatchSchedule *RedisPatchSchedule) ValidateDelete() error {
	validations := redisPatchSchedule.deleteValidations()
	var temp interface{} = redisPatchSchedule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (redisPatchSchedule *RedisPatchSchedule) ValidateUpdate(old runtime.Object) error {
	validations := redisPatchSchedule.updateValidations()
	var temp interface{} = redisPatchSchedule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (redisPatchSchedule *RedisPatchSchedule) createValidations() []func() error {
	return []func() error{redisPatchSchedule.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (redisPatchSchedule *RedisPatchSchedule) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (redisPatchSchedule *RedisPatchSchedule) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return redisPatchSchedule.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (redisPatchSchedule *RedisPatchSchedule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&redisPatchSchedule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromRedisPatchSchedule populates our RedisPatchSchedule from the provided source RedisPatchSchedule
func (redisPatchSchedule *RedisPatchSchedule) AssignPropertiesFromRedisPatchSchedule(source *v1alpha1api20201201storage.RedisPatchSchedule) error {

	// ObjectMeta
	redisPatchSchedule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisPatchSchedules_Spec
	err := spec.AssignPropertiesFromRedisPatchSchedulesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisPatchSchedulesSpec() to populate field Spec")
	}
	redisPatchSchedule.Spec = spec

	// Status
	var status RedisPatchSchedule_Status
	err = status.AssignPropertiesFromRedisPatchScheduleStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisPatchScheduleStatus() to populate field Status")
	}
	redisPatchSchedule.Status = status

	// No error
	return nil
}

// AssignPropertiesToRedisPatchSchedule populates the provided destination RedisPatchSchedule from our RedisPatchSchedule
func (redisPatchSchedule *RedisPatchSchedule) AssignPropertiesToRedisPatchSchedule(destination *v1alpha1api20201201storage.RedisPatchSchedule) error {

	// ObjectMeta
	destination.ObjectMeta = *redisPatchSchedule.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20201201storage.RedisPatchSchedules_Spec
	err := redisPatchSchedule.Spec.AssignPropertiesToRedisPatchSchedulesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisPatchSchedulesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20201201storage.RedisPatchSchedule_Status
	err = redisPatchSchedule.Status.AssignPropertiesToRedisPatchScheduleStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisPatchScheduleStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (redisPatchSchedule *RedisPatchSchedule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: redisPatchSchedule.Spec.OriginalVersion(),
		Kind:    "RedisPatchSchedule",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_patchSchedules
type RedisPatchScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisPatchSchedule `json:"items"`
}

type RedisPatchSchedule_Status struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//ScheduleEntries: List of patch schedules for a Redis cache.
	ScheduleEntries []ScheduleEntry_Status `json:"scheduleEntries,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
	//"Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisPatchSchedule_Status{}

// ConvertStatusFrom populates our RedisPatchSchedule_Status from the provided source
func (redisPatchScheduleStatus *RedisPatchSchedule_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20201201storage.RedisPatchSchedule_Status)
	if ok {
		// Populate our instance from source
		return redisPatchScheduleStatus.AssignPropertiesFromRedisPatchScheduleStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20201201storage.RedisPatchSchedule_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = redisPatchScheduleStatus.AssignPropertiesFromRedisPatchScheduleStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RedisPatchSchedule_Status
func (redisPatchScheduleStatus *RedisPatchSchedule_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20201201storage.RedisPatchSchedule_Status)
	if ok {
		// Populate destination from our instance
		return redisPatchScheduleStatus.AssignPropertiesToRedisPatchScheduleStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20201201storage.RedisPatchSchedule_Status{}
	err := redisPatchScheduleStatus.AssignPropertiesToRedisPatchScheduleStatus(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &RedisPatchSchedule_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (redisPatchScheduleStatus *RedisPatchSchedule_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisPatchSchedule_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (redisPatchScheduleStatus *RedisPatchSchedule_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisPatchSchedule_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisPatchSchedule_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		redisPatchScheduleStatus.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		redisPatchScheduleStatus.Name = &name
	}

	// Set property ‘ScheduleEntries’:
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.ScheduleEntries {
			var item1 ScheduleEntry_Status
			err := item1.PopulateFromARM(owner, item)
			if err != nil {
				return err
			}
			redisPatchScheduleStatus.ScheduleEntries = append(redisPatchScheduleStatus.ScheduleEntries, item1)
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		redisPatchScheduleStatus.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromRedisPatchScheduleStatus populates our RedisPatchSchedule_Status from the provided source RedisPatchSchedule_Status
func (redisPatchScheduleStatus *RedisPatchSchedule_Status) AssignPropertiesFromRedisPatchScheduleStatus(source *v1alpha1api20201201storage.RedisPatchSchedule_Status) error {

	// Conditions
	redisPatchScheduleStatus.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	redisPatchScheduleStatus.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	redisPatchScheduleStatus.Name = genruntime.ClonePointerToString(source.Name)

	// ScheduleEntries
	if source.ScheduleEntries != nil {
		scheduleEntryList := make([]ScheduleEntry_Status, len(source.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range source.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry ScheduleEntry_Status
			err := scheduleEntry.AssignPropertiesFromScheduleEntryStatus(&scheduleEntryItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromScheduleEntryStatus() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		redisPatchScheduleStatus.ScheduleEntries = scheduleEntryList
	} else {
		redisPatchScheduleStatus.ScheduleEntries = nil
	}

	// Type
	redisPatchScheduleStatus.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToRedisPatchScheduleStatus populates the provided destination RedisPatchSchedule_Status from our RedisPatchSchedule_Status
func (redisPatchScheduleStatus *RedisPatchSchedule_Status) AssignPropertiesToRedisPatchScheduleStatus(destination *v1alpha1api20201201storage.RedisPatchSchedule_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(redisPatchScheduleStatus.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(redisPatchScheduleStatus.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(redisPatchScheduleStatus.Name)

	// ScheduleEntries
	if redisPatchScheduleStatus.ScheduleEntries != nil {
		scheduleEntryList := make([]v1alpha1api20201201storage.ScheduleEntry_Status, len(redisPatchScheduleStatus.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range redisPatchScheduleStatus.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry v1alpha1api20201201storage.ScheduleEntry_Status
			err := scheduleEntryItem.AssignPropertiesToScheduleEntryStatus(&scheduleEntry)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToScheduleEntryStatus() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		destination.ScheduleEntries = scheduleEntryList
	} else {
		destination.ScheduleEntries = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(redisPatchScheduleStatus.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"2020-12-01"}
type RedisPatchSchedulesSpecAPIVersion string

const RedisPatchSchedulesSpecAPIVersion20201201 = RedisPatchSchedulesSpecAPIVersion("2020-12-01")

type RedisPatchSchedules_Spec struct {
	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner" kind:"Redis"`

	// +kubebuilder:validation:Required
	//ScheduleEntries: List of patch schedules for a Redis cache.
	ScheduleEntries []ScheduleEntry `json:"scheduleEntries"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &RedisPatchSchedules_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (redisPatchSchedulesSpec *RedisPatchSchedules_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if redisPatchSchedulesSpec == nil {
		return nil, nil
	}
	var result RedisPatchSchedules_SpecARM

	// Set property ‘Location’:
	if redisPatchSchedulesSpec.Location != nil {
		location := *redisPatchSchedulesSpec.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	for _, item := range redisPatchSchedulesSpec.ScheduleEntries {
		itemARM, err := item.ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		result.Properties.ScheduleEntries = append(result.Properties.ScheduleEntries, itemARM.(ScheduleEntryARM))
	}

	// Set property ‘Tags’:
	if redisPatchSchedulesSpec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range redisPatchSchedulesSpec.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (redisPatchSchedulesSpec *RedisPatchSchedules_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisPatchSchedules_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (redisPatchSchedulesSpec *RedisPatchSchedules_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisPatchSchedules_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisPatchSchedules_SpecARM, got %T", armInput)
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		redisPatchSchedulesSpec.Location = &location
	}

	// Set property ‘Owner’:
	redisPatchSchedulesSpec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘ScheduleEntries’:
	// copying flattened property:
	for _, item := range typedInput.Properties.ScheduleEntries {
		var item1 ScheduleEntry
		err := item1.PopulateFromARM(owner, item)
		if err != nil {
			return err
		}
		redisPatchSchedulesSpec.ScheduleEntries = append(redisPatchSchedulesSpec.ScheduleEntries, item1)
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		redisPatchSchedulesSpec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			redisPatchSchedulesSpec.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RedisPatchSchedules_Spec{}

// ConvertSpecFrom populates our RedisPatchSchedules_Spec from the provided source
func (redisPatchSchedulesSpec *RedisPatchSchedules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20201201storage.RedisPatchSchedules_Spec)
	if ok {
		// Populate our instance from source
		return redisPatchSchedulesSpec.AssignPropertiesFromRedisPatchSchedulesSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20201201storage.RedisPatchSchedules_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = redisPatchSchedulesSpec.AssignPropertiesFromRedisPatchSchedulesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisPatchSchedules_Spec
func (redisPatchSchedulesSpec *RedisPatchSchedules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20201201storage.RedisPatchSchedules_Spec)
	if ok {
		// Populate destination from our instance
		return redisPatchSchedulesSpec.AssignPropertiesToRedisPatchSchedulesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20201201storage.RedisPatchSchedules_Spec{}
	err := redisPatchSchedulesSpec.AssignPropertiesToRedisPatchSchedulesSpec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromRedisPatchSchedulesSpec populates our RedisPatchSchedules_Spec from the provided source RedisPatchSchedules_Spec
func (redisPatchSchedulesSpec *RedisPatchSchedules_Spec) AssignPropertiesFromRedisPatchSchedulesSpec(source *v1alpha1api20201201storage.RedisPatchSchedules_Spec) error {

	// Location
	redisPatchSchedulesSpec.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	redisPatchSchedulesSpec.Owner = source.Owner.Copy()

	// ScheduleEntries
	if source.ScheduleEntries != nil {
		scheduleEntryList := make([]ScheduleEntry, len(source.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range source.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry ScheduleEntry
			err := scheduleEntry.AssignPropertiesFromScheduleEntry(&scheduleEntryItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromScheduleEntry() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		redisPatchSchedulesSpec.ScheduleEntries = scheduleEntryList
	} else {
		redisPatchSchedulesSpec.ScheduleEntries = nil
	}

	// Tags
	redisPatchSchedulesSpec.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToRedisPatchSchedulesSpec populates the provided destination RedisPatchSchedules_Spec from our RedisPatchSchedules_Spec
func (redisPatchSchedulesSpec *RedisPatchSchedules_Spec) AssignPropertiesToRedisPatchSchedulesSpec(destination *v1alpha1api20201201storage.RedisPatchSchedules_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Location
	destination.Location = genruntime.ClonePointerToString(redisPatchSchedulesSpec.Location)

	// OriginalVersion
	destination.OriginalVersion = redisPatchSchedulesSpec.OriginalVersion()

	// Owner
	destination.Owner = redisPatchSchedulesSpec.Owner.Copy()

	// ScheduleEntries
	if redisPatchSchedulesSpec.ScheduleEntries != nil {
		scheduleEntryList := make([]v1alpha1api20201201storage.ScheduleEntry, len(redisPatchSchedulesSpec.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range redisPatchSchedulesSpec.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry v1alpha1api20201201storage.ScheduleEntry
			err := scheduleEntryItem.AssignPropertiesToScheduleEntry(&scheduleEntry)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToScheduleEntry() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		destination.ScheduleEntries = scheduleEntryList
	} else {
		destination.ScheduleEntries = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(redisPatchSchedulesSpec.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (redisPatchSchedulesSpec *RedisPatchSchedules_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

//Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/ScheduleEntry
type ScheduleEntry struct {
	// +kubebuilder:validation:Required
	//DayOfWeek: Day of the week when a cache can be patched.
	DayOfWeek ScheduleEntryDayOfWeek `json:"dayOfWeek"`

	//MaintenanceWindow: ISO8601 timespan specifying how much time cache patching can
	//take.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty"`

	// +kubebuilder:validation:Required
	//StartHourUtc: Start hour after which cache patching can start.
	StartHourUtc int `json:"startHourUtc"`
}

var _ genruntime.ARMTransformer = &ScheduleEntry{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (scheduleEntry *ScheduleEntry) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if scheduleEntry == nil {
		return nil, nil
	}
	var result ScheduleEntryARM

	// Set property ‘DayOfWeek’:
	result.DayOfWeek = scheduleEntry.DayOfWeek

	// Set property ‘MaintenanceWindow’:
	if scheduleEntry.MaintenanceWindow != nil {
		maintenanceWindow := *scheduleEntry.MaintenanceWindow
		result.MaintenanceWindow = &maintenanceWindow
	}

	// Set property ‘StartHourUtc’:
	result.StartHourUtc = scheduleEntry.StartHourUtc
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (scheduleEntry *ScheduleEntry) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ScheduleEntryARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (scheduleEntry *ScheduleEntry) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ScheduleEntryARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ScheduleEntryARM, got %T", armInput)
	}

	// Set property ‘DayOfWeek’:
	scheduleEntry.DayOfWeek = typedInput.DayOfWeek

	// Set property ‘MaintenanceWindow’:
	if typedInput.MaintenanceWindow != nil {
		maintenanceWindow := *typedInput.MaintenanceWindow
		scheduleEntry.MaintenanceWindow = &maintenanceWindow
	}

	// Set property ‘StartHourUtc’:
	scheduleEntry.StartHourUtc = typedInput.StartHourUtc

	// No error
	return nil
}

// AssignPropertiesFromScheduleEntry populates our ScheduleEntry from the provided source ScheduleEntry
func (scheduleEntry *ScheduleEntry) AssignPropertiesFromScheduleEntry(source *v1alpha1api20201201storage.ScheduleEntry) error {

	// DayOfWeek
	if source.DayOfWeek != nil {
		scheduleEntry.DayOfWeek = ScheduleEntryDayOfWeek(*source.DayOfWeek)
	} else {
		scheduleEntry.DayOfWeek = ""
	}

	// MaintenanceWindow
	if source.MaintenanceWindow != nil {
		maintenanceWindow := *source.MaintenanceWindow
		scheduleEntry.MaintenanceWindow = &maintenanceWindow
	} else {
		scheduleEntry.MaintenanceWindow = nil
	}

	// StartHourUtc
	scheduleEntry.StartHourUtc = genruntime.GetOptionalIntValue(source.StartHourUtc)

	// No error
	return nil
}

// AssignPropertiesToScheduleEntry populates the provided destination ScheduleEntry from our ScheduleEntry
func (scheduleEntry *ScheduleEntry) AssignPropertiesToScheduleEntry(destination *v1alpha1api20201201storage.ScheduleEntry) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// DayOfWeek
	dayOfWeek := string(scheduleEntry.DayOfWeek)
	destination.DayOfWeek = &dayOfWeek

	// MaintenanceWindow
	if scheduleEntry.MaintenanceWindow != nil {
		maintenanceWindow := *scheduleEntry.MaintenanceWindow
		destination.MaintenanceWindow = &maintenanceWindow
	} else {
		destination.MaintenanceWindow = nil
	}

	// StartHourUtc
	startHourUtc := scheduleEntry.StartHourUtc
	destination.StartHourUtc = &startHourUtc

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type ScheduleEntry_Status struct {
	// +kubebuilder:validation:Required
	//DayOfWeek: Day of the week when a cache can be patched.
	DayOfWeek ScheduleEntryStatusDayOfWeek `json:"dayOfWeek"`

	//MaintenanceWindow: ISO8601 timespan specifying how much time cache patching can
	//take.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty"`

	// +kubebuilder:validation:Required
	//StartHourUtc: Start hour after which cache patching can start.
	StartHourUtc int `json:"startHourUtc"`
}

var _ genruntime.FromARMConverter = &ScheduleEntry_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (scheduleEntryStatus *ScheduleEntry_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ScheduleEntry_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (scheduleEntryStatus *ScheduleEntry_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ScheduleEntry_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ScheduleEntry_StatusARM, got %T", armInput)
	}

	// Set property ‘DayOfWeek’:
	scheduleEntryStatus.DayOfWeek = typedInput.DayOfWeek

	// Set property ‘MaintenanceWindow’:
	if typedInput.MaintenanceWindow != nil {
		maintenanceWindow := *typedInput.MaintenanceWindow
		scheduleEntryStatus.MaintenanceWindow = &maintenanceWindow
	}

	// Set property ‘StartHourUtc’:
	scheduleEntryStatus.StartHourUtc = typedInput.StartHourUtc

	// No error
	return nil
}

// AssignPropertiesFromScheduleEntryStatus populates our ScheduleEntry_Status from the provided source ScheduleEntry_Status
func (scheduleEntryStatus *ScheduleEntry_Status) AssignPropertiesFromScheduleEntryStatus(source *v1alpha1api20201201storage.ScheduleEntry_Status) error {

	// DayOfWeek
	if source.DayOfWeek != nil {
		scheduleEntryStatus.DayOfWeek = ScheduleEntryStatusDayOfWeek(*source.DayOfWeek)
	} else {
		scheduleEntryStatus.DayOfWeek = ""
	}

	// MaintenanceWindow
	scheduleEntryStatus.MaintenanceWindow = genruntime.ClonePointerToString(source.MaintenanceWindow)

	// StartHourUtc
	scheduleEntryStatus.StartHourUtc = genruntime.GetOptionalIntValue(source.StartHourUtc)

	// No error
	return nil
}

// AssignPropertiesToScheduleEntryStatus populates the provided destination ScheduleEntry_Status from our ScheduleEntry_Status
func (scheduleEntryStatus *ScheduleEntry_Status) AssignPropertiesToScheduleEntryStatus(destination *v1alpha1api20201201storage.ScheduleEntry_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// DayOfWeek
	dayOfWeek := string(scheduleEntryStatus.DayOfWeek)
	destination.DayOfWeek = &dayOfWeek

	// MaintenanceWindow
	destination.MaintenanceWindow = genruntime.ClonePointerToString(scheduleEntryStatus.MaintenanceWindow)

	// StartHourUtc
	startHourUtc := scheduleEntryStatus.StartHourUtc
	destination.StartHourUtc = &startHourUtc

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"Everyday","Friday","Monday","Saturday","Sunday","Thursday","Tuesday","Wednesday","Weekend"}
type ScheduleEntryDayOfWeek string

const (
	ScheduleEntryDayOfWeekEveryday  = ScheduleEntryDayOfWeek("Everyday")
	ScheduleEntryDayOfWeekFriday    = ScheduleEntryDayOfWeek("Friday")
	ScheduleEntryDayOfWeekMonday    = ScheduleEntryDayOfWeek("Monday")
	ScheduleEntryDayOfWeekSaturday  = ScheduleEntryDayOfWeek("Saturday")
	ScheduleEntryDayOfWeekSunday    = ScheduleEntryDayOfWeek("Sunday")
	ScheduleEntryDayOfWeekThursday  = ScheduleEntryDayOfWeek("Thursday")
	ScheduleEntryDayOfWeekTuesday   = ScheduleEntryDayOfWeek("Tuesday")
	ScheduleEntryDayOfWeekWednesday = ScheduleEntryDayOfWeek("Wednesday")
	ScheduleEntryDayOfWeekWeekend   = ScheduleEntryDayOfWeek("Weekend")
)

type ScheduleEntryStatusDayOfWeek string

const (
	ScheduleEntryStatusDayOfWeekEveryday  = ScheduleEntryStatusDayOfWeek("Everyday")
	ScheduleEntryStatusDayOfWeekFriday    = ScheduleEntryStatusDayOfWeek("Friday")
	ScheduleEntryStatusDayOfWeekMonday    = ScheduleEntryStatusDayOfWeek("Monday")
	ScheduleEntryStatusDayOfWeekSaturday  = ScheduleEntryStatusDayOfWeek("Saturday")
	ScheduleEntryStatusDayOfWeekSunday    = ScheduleEntryStatusDayOfWeek("Sunday")
	ScheduleEntryStatusDayOfWeekThursday  = ScheduleEntryStatusDayOfWeek("Thursday")
	ScheduleEntryStatusDayOfWeekTuesday   = ScheduleEntryStatusDayOfWeek("Tuesday")
	ScheduleEntryStatusDayOfWeekWednesday = ScheduleEntryStatusDayOfWeek("Wednesday")
	ScheduleEntryStatusDayOfWeekWeekend   = ScheduleEntryStatusDayOfWeek("Weekend")
)

func init() {
	SchemeBuilder.Register(&RedisPatchSchedule{}, &RedisPatchScheduleList{})
}