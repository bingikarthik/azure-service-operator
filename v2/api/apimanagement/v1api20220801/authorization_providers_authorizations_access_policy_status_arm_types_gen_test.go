// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

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

func Test_AuthorizationAccessPolicyContractProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationAccessPolicyContractProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationAccessPolicyContractProperties_STATUS_ARM, AuthorizationAccessPolicyContractProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationAccessPolicyContractProperties_STATUS_ARM runs a test to see if a specific instance of AuthorizationAccessPolicyContractProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationAccessPolicyContractProperties_STATUS_ARM(subject AuthorizationAccessPolicyContractProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationAccessPolicyContractProperties_STATUS_ARM
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

// Generator of AuthorizationAccessPolicyContractProperties_STATUS_ARM instances for property testing - lazily
// instantiated by AuthorizationAccessPolicyContractProperties_STATUS_ARMGenerator()
var authorizationAccessPolicyContractProperties_STATUS_ARMGenerator gopter.Gen

// AuthorizationAccessPolicyContractProperties_STATUS_ARMGenerator returns a generator of AuthorizationAccessPolicyContractProperties_STATUS_ARM instances for property testing.
func AuthorizationAccessPolicyContractProperties_STATUS_ARMGenerator() gopter.Gen {
	if authorizationAccessPolicyContractProperties_STATUS_ARMGenerator != nil {
		return authorizationAccessPolicyContractProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationAccessPolicyContractProperties_STATUS_ARM(generators)
	authorizationAccessPolicyContractProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationAccessPolicyContractProperties_STATUS_ARM{}), generators)

	return authorizationAccessPolicyContractProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationAccessPolicyContractProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationAccessPolicyContractProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ObjectId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

func Test_AuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM, AuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM runs a test to see if a specific instance of AuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM(subject AuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM
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

// Generator of AuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM instances for property testing - lazily
// instantiated by AuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARMGenerator()
var authorizationProvidersAuthorizationsAccessPolicy_STATUS_ARMGenerator gopter.Gen

// AuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARMGenerator returns a generator of AuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM instances for property testing.
// We first initialize authorizationProvidersAuthorizationsAccessPolicy_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARMGenerator() gopter.Gen {
	if authorizationProvidersAuthorizationsAccessPolicy_STATUS_ARMGenerator != nil {
		return authorizationProvidersAuthorizationsAccessPolicy_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM(generators)
	authorizationProvidersAuthorizationsAccessPolicy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM(generators)
	authorizationProvidersAuthorizationsAccessPolicy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM{}), generators)

	return authorizationProvidersAuthorizationsAccessPolicy_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorizationsAccessPolicy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AuthorizationAccessPolicyContractProperties_STATUS_ARMGenerator())
}
