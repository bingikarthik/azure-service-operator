// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180501preview

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

func Test_Webtest_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Webtest_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebtest_Spec_ARM, Webtest_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebtest_Spec_ARM runs a test to see if a specific instance of Webtest_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebtest_Spec_ARM(subject Webtest_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Webtest_Spec_ARM
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

// Generator of Webtest_Spec_ARM instances for property testing - lazily instantiated by Webtest_Spec_ARMGenerator()
var webtest_Spec_ARMGenerator gopter.Gen

// Webtest_Spec_ARMGenerator returns a generator of Webtest_Spec_ARM instances for property testing.
// We first initialize webtest_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Webtest_Spec_ARMGenerator() gopter.Gen {
	if webtest_Spec_ARMGenerator != nil {
		return webtest_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebtest_Spec_ARM(generators)
	webtest_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Webtest_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebtest_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForWebtest_Spec_ARM(generators)
	webtest_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Webtest_Spec_ARM{}), generators)

	return webtest_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebtest_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebtest_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWebtest_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebtest_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(WebTestProperties_ARMGenerator())
}

func Test_WebTestProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestProperties_ARM, WebTestProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestProperties_ARM runs a test to see if a specific instance of WebTestProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestProperties_ARM(subject WebTestProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_ARM
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

// Generator of WebTestProperties_ARM instances for property testing - lazily instantiated by
// WebTestProperties_ARMGenerator()
var webTestProperties_ARMGenerator gopter.Gen

// WebTestProperties_ARMGenerator returns a generator of WebTestProperties_ARM instances for property testing.
// We first initialize webTestProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WebTestProperties_ARMGenerator() gopter.Gen {
	if webTestProperties_ARMGenerator != nil {
		return webTestProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_ARM(generators)
	webTestProperties_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForWebTestProperties_ARM(generators)
	webTestProperties_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_ARM{}), generators)

	return webTestProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestProperties_ARM(gens map[string]gopter.Gen) {
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Frequency"] = gen.PtrOf(gen.Int())
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(
		WebTestProperties_Kind_Basic,
		WebTestProperties_Kind_Multistep,
		WebTestProperties_Kind_Ping,
		WebTestProperties_Kind_Standard))
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RetryEnabled"] = gen.PtrOf(gen.Bool())
	gens["SyntheticMonitorId"] = gen.PtrOf(gen.AlphaString())
	gens["Timeout"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForWebTestProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebTestProperties_ARM(gens map[string]gopter.Gen) {
	gens["Configuration"] = gen.PtrOf(WebTestPropertiesConfiguration_ARMGenerator())
	gens["Locations"] = gen.SliceOf(WebTestGeolocation_ARMGenerator())
	gens["Request"] = gen.PtrOf(WebTestPropertiesRequest_ARMGenerator())
	gens["ValidationRules"] = gen.PtrOf(WebTestPropertiesValidationRules_ARMGenerator())
}

func Test_WebTestGeolocation_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestGeolocation_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestGeolocation_ARM, WebTestGeolocation_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestGeolocation_ARM runs a test to see if a specific instance of WebTestGeolocation_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestGeolocation_ARM(subject WebTestGeolocation_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestGeolocation_ARM
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

// Generator of WebTestGeolocation_ARM instances for property testing - lazily instantiated by
// WebTestGeolocation_ARMGenerator()
var webTestGeolocation_ARMGenerator gopter.Gen

// WebTestGeolocation_ARMGenerator returns a generator of WebTestGeolocation_ARM instances for property testing.
func WebTestGeolocation_ARMGenerator() gopter.Gen {
	if webTestGeolocation_ARMGenerator != nil {
		return webTestGeolocation_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestGeolocation_ARM(generators)
	webTestGeolocation_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestGeolocation_ARM{}), generators)

	return webTestGeolocation_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestGeolocation_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestGeolocation_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_WebTestPropertiesConfiguration_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestPropertiesConfiguration_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestPropertiesConfiguration_ARM, WebTestPropertiesConfiguration_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestPropertiesConfiguration_ARM runs a test to see if a specific instance of WebTestPropertiesConfiguration_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestPropertiesConfiguration_ARM(subject WebTestPropertiesConfiguration_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestPropertiesConfiguration_ARM
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

// Generator of WebTestPropertiesConfiguration_ARM instances for property testing - lazily instantiated by
// WebTestPropertiesConfiguration_ARMGenerator()
var webTestPropertiesConfiguration_ARMGenerator gopter.Gen

// WebTestPropertiesConfiguration_ARMGenerator returns a generator of WebTestPropertiesConfiguration_ARM instances for property testing.
func WebTestPropertiesConfiguration_ARMGenerator() gopter.Gen {
	if webTestPropertiesConfiguration_ARMGenerator != nil {
		return webTestPropertiesConfiguration_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestPropertiesConfiguration_ARM(generators)
	webTestPropertiesConfiguration_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestPropertiesConfiguration_ARM{}), generators)

	return webTestPropertiesConfiguration_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestPropertiesConfiguration_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestPropertiesConfiguration_ARM(gens map[string]gopter.Gen) {
	gens["WebTest"] = gen.PtrOf(gen.AlphaString())
}

func Test_WebTestPropertiesRequest_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestPropertiesRequest_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestPropertiesRequest_ARM, WebTestPropertiesRequest_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestPropertiesRequest_ARM runs a test to see if a specific instance of WebTestPropertiesRequest_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestPropertiesRequest_ARM(subject WebTestPropertiesRequest_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestPropertiesRequest_ARM
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

// Generator of WebTestPropertiesRequest_ARM instances for property testing - lazily instantiated by
// WebTestPropertiesRequest_ARMGenerator()
var webTestPropertiesRequest_ARMGenerator gopter.Gen

// WebTestPropertiesRequest_ARMGenerator returns a generator of WebTestPropertiesRequest_ARM instances for property testing.
// We first initialize webTestPropertiesRequest_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WebTestPropertiesRequest_ARMGenerator() gopter.Gen {
	if webTestPropertiesRequest_ARMGenerator != nil {
		return webTestPropertiesRequest_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestPropertiesRequest_ARM(generators)
	webTestPropertiesRequest_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestPropertiesRequest_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestPropertiesRequest_ARM(generators)
	AddRelatedPropertyGeneratorsForWebTestPropertiesRequest_ARM(generators)
	webTestPropertiesRequest_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestPropertiesRequest_ARM{}), generators)

	return webTestPropertiesRequest_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestPropertiesRequest_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestPropertiesRequest_ARM(gens map[string]gopter.Gen) {
	gens["FollowRedirects"] = gen.PtrOf(gen.Bool())
	gens["HttpVerb"] = gen.PtrOf(gen.AlphaString())
	gens["ParseDependentRequests"] = gen.PtrOf(gen.Bool())
	gens["RequestBody"] = gen.PtrOf(gen.AlphaString())
	gens["RequestUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWebTestPropertiesRequest_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebTestPropertiesRequest_ARM(gens map[string]gopter.Gen) {
	gens["Headers"] = gen.SliceOf(HeaderField_ARMGenerator())
}

func Test_WebTestPropertiesValidationRules_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestPropertiesValidationRules_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestPropertiesValidationRules_ARM, WebTestPropertiesValidationRules_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestPropertiesValidationRules_ARM runs a test to see if a specific instance of WebTestPropertiesValidationRules_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestPropertiesValidationRules_ARM(subject WebTestPropertiesValidationRules_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestPropertiesValidationRules_ARM
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

// Generator of WebTestPropertiesValidationRules_ARM instances for property testing - lazily instantiated by
// WebTestPropertiesValidationRules_ARMGenerator()
var webTestPropertiesValidationRules_ARMGenerator gopter.Gen

// WebTestPropertiesValidationRules_ARMGenerator returns a generator of WebTestPropertiesValidationRules_ARM instances for property testing.
// We first initialize webTestPropertiesValidationRules_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WebTestPropertiesValidationRules_ARMGenerator() gopter.Gen {
	if webTestPropertiesValidationRules_ARMGenerator != nil {
		return webTestPropertiesValidationRules_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestPropertiesValidationRules_ARM(generators)
	webTestPropertiesValidationRules_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestPropertiesValidationRules_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestPropertiesValidationRules_ARM(generators)
	AddRelatedPropertyGeneratorsForWebTestPropertiesValidationRules_ARM(generators)
	webTestPropertiesValidationRules_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestPropertiesValidationRules_ARM{}), generators)

	return webTestPropertiesValidationRules_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestPropertiesValidationRules_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestPropertiesValidationRules_ARM(gens map[string]gopter.Gen) {
	gens["ExpectedHttpStatusCode"] = gen.PtrOf(gen.Int())
	gens["IgnoreHttpsStatusCode"] = gen.PtrOf(gen.Bool())
	gens["SSLCertRemainingLifetimeCheck"] = gen.PtrOf(gen.Int())
	gens["SSLCheck"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForWebTestPropertiesValidationRules_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebTestPropertiesValidationRules_ARM(gens map[string]gopter.Gen) {
	gens["ContentValidation"] = gen.PtrOf(WebTestPropertiesValidationRulesContentValidation_ARMGenerator())
}

func Test_HeaderField_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HeaderField_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHeaderField_ARM, HeaderField_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHeaderField_ARM runs a test to see if a specific instance of HeaderField_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForHeaderField_ARM(subject HeaderField_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HeaderField_ARM
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

// Generator of HeaderField_ARM instances for property testing - lazily instantiated by HeaderField_ARMGenerator()
var headerField_ARMGenerator gopter.Gen

// HeaderField_ARMGenerator returns a generator of HeaderField_ARM instances for property testing.
func HeaderField_ARMGenerator() gopter.Gen {
	if headerField_ARMGenerator != nil {
		return headerField_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHeaderField_ARM(generators)
	headerField_ARMGenerator = gen.Struct(reflect.TypeOf(HeaderField_ARM{}), generators)

	return headerField_ARMGenerator
}

// AddIndependentPropertyGeneratorsForHeaderField_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHeaderField_ARM(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_WebTestPropertiesValidationRulesContentValidation_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestPropertiesValidationRulesContentValidation_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestPropertiesValidationRulesContentValidation_ARM, WebTestPropertiesValidationRulesContentValidation_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestPropertiesValidationRulesContentValidation_ARM runs a test to see if a specific instance of WebTestPropertiesValidationRulesContentValidation_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestPropertiesValidationRulesContentValidation_ARM(subject WebTestPropertiesValidationRulesContentValidation_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestPropertiesValidationRulesContentValidation_ARM
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

// Generator of WebTestPropertiesValidationRulesContentValidation_ARM instances for property testing - lazily
// instantiated by WebTestPropertiesValidationRulesContentValidation_ARMGenerator()
var webTestPropertiesValidationRulesContentValidation_ARMGenerator gopter.Gen

// WebTestPropertiesValidationRulesContentValidation_ARMGenerator returns a generator of WebTestPropertiesValidationRulesContentValidation_ARM instances for property testing.
func WebTestPropertiesValidationRulesContentValidation_ARMGenerator() gopter.Gen {
	if webTestPropertiesValidationRulesContentValidation_ARMGenerator != nil {
		return webTestPropertiesValidationRulesContentValidation_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestPropertiesValidationRulesContentValidation_ARM(generators)
	webTestPropertiesValidationRulesContentValidation_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestPropertiesValidationRulesContentValidation_ARM{}), generators)

	return webTestPropertiesValidationRulesContentValidation_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestPropertiesValidationRulesContentValidation_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestPropertiesValidationRulesContentValidation_ARM(gens map[string]gopter.Gen) {
	gens["ContentMatch"] = gen.PtrOf(gen.AlphaString())
	gens["IgnoreCase"] = gen.PtrOf(gen.Bool())
	gens["PassIfTextFound"] = gen.PtrOf(gen.Bool())
}