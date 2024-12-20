// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_AccessPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AccessPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAccessPolicy, AccessPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAccessPolicy runs a test to see if a specific instance of AccessPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForAccessPolicy(subject AccessPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AccessPolicy
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AccessPolicy instances for property testing - lazily instantiated by AccessPolicyGenerator()
var accessPolicyGenerator gopter.Gen

// AccessPolicyGenerator returns a generator of AccessPolicy instances for property testing.
func AccessPolicyGenerator() gopter.Gen {
	if accessPolicyGenerator != nil {
		return accessPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAccessPolicy(generators)
	accessPolicyGenerator = gen.Struct(reflect.TypeOf(AccessPolicy{}), generators)

	return accessPolicyGenerator
}

// AddIndependentPropertyGeneratorsForAccessPolicy is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAccessPolicy(gens map[string]gopter.Gen) {
	gens["ExpiryTime"] = gen.PtrOf(gen.AlphaString())
	gens["Permission"] = gen.PtrOf(gen.AlphaString())
	gens["StartTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_FileShareProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FileShareProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFileShareProperties, FileSharePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFileShareProperties runs a test to see if a specific instance of FileShareProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForFileShareProperties(subject FileShareProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FileShareProperties
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of FileShareProperties instances for property testing - lazily instantiated by
// FileSharePropertiesGenerator()
var fileSharePropertiesGenerator gopter.Gen

// FileSharePropertiesGenerator returns a generator of FileShareProperties instances for property testing.
// We first initialize fileSharePropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FileSharePropertiesGenerator() gopter.Gen {
	if fileSharePropertiesGenerator != nil {
		return fileSharePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFileShareProperties(generators)
	fileSharePropertiesGenerator = gen.Struct(reflect.TypeOf(FileShareProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFileShareProperties(generators)
	AddRelatedPropertyGeneratorsForFileShareProperties(generators)
	fileSharePropertiesGenerator = gen.Struct(reflect.TypeOf(FileShareProperties{}), generators)

	return fileSharePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForFileShareProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFileShareProperties(gens map[string]gopter.Gen) {
	gens["AccessTier"] = gen.PtrOf(gen.OneConstOf(
		FileShareProperties_AccessTier_Cool,
		FileShareProperties_AccessTier_Hot,
		FileShareProperties_AccessTier_Premium,
		FileShareProperties_AccessTier_TransactionOptimized))
	gens["EnabledProtocols"] = gen.PtrOf(gen.OneConstOf(FileShareProperties_EnabledProtocols_NFS, FileShareProperties_EnabledProtocols_SMB))
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["RootSquash"] = gen.PtrOf(gen.OneConstOf(FileShareProperties_RootSquash_AllSquash, FileShareProperties_RootSquash_NoRootSquash, FileShareProperties_RootSquash_RootSquash))
	gens["ShareQuota"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForFileShareProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFileShareProperties(gens map[string]gopter.Gen) {
	gens["SignedIdentifiers"] = gen.SliceOf(SignedIdentifierGenerator())
}

func Test_SignedIdentifier_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SignedIdentifier via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSignedIdentifier, SignedIdentifierGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSignedIdentifier runs a test to see if a specific instance of SignedIdentifier round trips to JSON and back losslessly
func RunJSONSerializationTestForSignedIdentifier(subject SignedIdentifier) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SignedIdentifier
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SignedIdentifier instances for property testing - lazily instantiated by SignedIdentifierGenerator()
var signedIdentifierGenerator gopter.Gen

// SignedIdentifierGenerator returns a generator of SignedIdentifier instances for property testing.
// We first initialize signedIdentifierGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SignedIdentifierGenerator() gopter.Gen {
	if signedIdentifierGenerator != nil {
		return signedIdentifierGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSignedIdentifier(generators)
	signedIdentifierGenerator = gen.Struct(reflect.TypeOf(SignedIdentifier{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSignedIdentifier(generators)
	AddRelatedPropertyGeneratorsForSignedIdentifier(generators)
	signedIdentifierGenerator = gen.Struct(reflect.TypeOf(SignedIdentifier{}), generators)

	return signedIdentifierGenerator
}

// AddIndependentPropertyGeneratorsForSignedIdentifier is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSignedIdentifier(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSignedIdentifier is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSignedIdentifier(gens map[string]gopter.Gen) {
	gens["AccessPolicy"] = gen.PtrOf(AccessPolicyGenerator())
}

func Test_StorageAccountsFileServicesShare_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsFileServicesShare_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsFileServicesShare_Spec, StorageAccountsFileServicesShare_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsFileServicesShare_Spec runs a test to see if a specific instance of StorageAccountsFileServicesShare_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsFileServicesShare_Spec(subject StorageAccountsFileServicesShare_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsFileServicesShare_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of StorageAccountsFileServicesShare_Spec instances for property testing - lazily instantiated by
// StorageAccountsFileServicesShare_SpecGenerator()
var storageAccountsFileServicesShare_SpecGenerator gopter.Gen

// StorageAccountsFileServicesShare_SpecGenerator returns a generator of StorageAccountsFileServicesShare_Spec instances for property testing.
// We first initialize storageAccountsFileServicesShare_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsFileServicesShare_SpecGenerator() gopter.Gen {
	if storageAccountsFileServicesShare_SpecGenerator != nil {
		return storageAccountsFileServicesShare_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsFileServicesShare_Spec(generators)
	storageAccountsFileServicesShare_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsFileServicesShare_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsFileServicesShare_Spec(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsFileServicesShare_Spec(generators)
	storageAccountsFileServicesShare_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsFileServicesShare_Spec{}), generators)

	return storageAccountsFileServicesShare_SpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsFileServicesShare_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsFileServicesShare_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForStorageAccountsFileServicesShare_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsFileServicesShare_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FileSharePropertiesGenerator())
}
