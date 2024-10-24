// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230630

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

func Test_DatabaseProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseProperties_ARM, DatabaseProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseProperties_ARM runs a test to see if a specific instance of DatabaseProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseProperties_ARM(subject DatabaseProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseProperties_ARM
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

// Generator of DatabaseProperties_ARM instances for property testing - lazily instantiated by
// DatabaseProperties_ARMGenerator()
var databaseProperties_ARMGenerator gopter.Gen

// DatabaseProperties_ARMGenerator returns a generator of DatabaseProperties_ARM instances for property testing.
func DatabaseProperties_ARMGenerator() gopter.Gen {
	if databaseProperties_ARMGenerator != nil {
		return databaseProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties_ARM(generators)
	databaseProperties_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_ARM{}), generators)

	return databaseProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseProperties_ARM(gens map[string]gopter.Gen) {
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
}

func Test_FlexibleServersDatabase_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersDatabase_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersDatabase_Spec_ARM, FlexibleServersDatabase_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersDatabase_Spec_ARM runs a test to see if a specific instance of FlexibleServersDatabase_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersDatabase_Spec_ARM(subject FlexibleServersDatabase_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersDatabase_Spec_ARM
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

// Generator of FlexibleServersDatabase_Spec_ARM instances for property testing - lazily instantiated by
// FlexibleServersDatabase_Spec_ARMGenerator()
var flexibleServersDatabase_Spec_ARMGenerator gopter.Gen

// FlexibleServersDatabase_Spec_ARMGenerator returns a generator of FlexibleServersDatabase_Spec_ARM instances for property testing.
// We first initialize flexibleServersDatabase_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersDatabase_Spec_ARMGenerator() gopter.Gen {
	if flexibleServersDatabase_Spec_ARMGenerator != nil {
		return flexibleServersDatabase_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersDatabase_Spec_ARM(generators)
	flexibleServersDatabase_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersDatabase_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersDatabase_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersDatabase_Spec_ARM(generators)
	flexibleServersDatabase_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersDatabase_Spec_ARM{}), generators)

	return flexibleServersDatabase_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersDatabase_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersDatabase_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForFlexibleServersDatabase_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersDatabase_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabaseProperties_ARMGenerator())
}
