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

func Test_BackupShortTermRetentionPolicyProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackupShortTermRetentionPolicyProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackupShortTermRetentionPolicyProperties, BackupShortTermRetentionPolicyPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackupShortTermRetentionPolicyProperties runs a test to see if a specific instance of BackupShortTermRetentionPolicyProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForBackupShortTermRetentionPolicyProperties(subject BackupShortTermRetentionPolicyProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackupShortTermRetentionPolicyProperties
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

// Generator of BackupShortTermRetentionPolicyProperties instances for property testing - lazily instantiated by
// BackupShortTermRetentionPolicyPropertiesGenerator()
var backupShortTermRetentionPolicyPropertiesGenerator gopter.Gen

// BackupShortTermRetentionPolicyPropertiesGenerator returns a generator of BackupShortTermRetentionPolicyProperties instances for property testing.
func BackupShortTermRetentionPolicyPropertiesGenerator() gopter.Gen {
	if backupShortTermRetentionPolicyPropertiesGenerator != nil {
		return backupShortTermRetentionPolicyPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupShortTermRetentionPolicyProperties(generators)
	backupShortTermRetentionPolicyPropertiesGenerator = gen.Struct(reflect.TypeOf(BackupShortTermRetentionPolicyProperties{}), generators)

	return backupShortTermRetentionPolicyPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForBackupShortTermRetentionPolicyProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackupShortTermRetentionPolicyProperties(gens map[string]gopter.Gen) {
	gens["DiffBackupIntervalInHours"] = gen.PtrOf(gen.OneConstOf(BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_12, BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_24))
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
}

func Test_ServersDatabasesBackupShortTermRetentionPolicy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesBackupShortTermRetentionPolicy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesBackupShortTermRetentionPolicy_Spec, ServersDatabasesBackupShortTermRetentionPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesBackupShortTermRetentionPolicy_Spec runs a test to see if a specific instance of ServersDatabasesBackupShortTermRetentionPolicy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesBackupShortTermRetentionPolicy_Spec(subject ServersDatabasesBackupShortTermRetentionPolicy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesBackupShortTermRetentionPolicy_Spec
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

// Generator of ServersDatabasesBackupShortTermRetentionPolicy_Spec instances for property testing - lazily instantiated
// by ServersDatabasesBackupShortTermRetentionPolicy_SpecGenerator()
var serversDatabasesBackupShortTermRetentionPolicy_SpecGenerator gopter.Gen

// ServersDatabasesBackupShortTermRetentionPolicy_SpecGenerator returns a generator of ServersDatabasesBackupShortTermRetentionPolicy_Spec instances for property testing.
// We first initialize serversDatabasesBackupShortTermRetentionPolicy_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersDatabasesBackupShortTermRetentionPolicy_SpecGenerator() gopter.Gen {
	if serversDatabasesBackupShortTermRetentionPolicy_SpecGenerator != nil {
		return serversDatabasesBackupShortTermRetentionPolicy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabasesBackupShortTermRetentionPolicy_Spec(generators)
	serversDatabasesBackupShortTermRetentionPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesBackupShortTermRetentionPolicy_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabasesBackupShortTermRetentionPolicy_Spec(generators)
	AddRelatedPropertyGeneratorsForServersDatabasesBackupShortTermRetentionPolicy_Spec(generators)
	serversDatabasesBackupShortTermRetentionPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesBackupShortTermRetentionPolicy_Spec{}), generators)

	return serversDatabasesBackupShortTermRetentionPolicy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServersDatabasesBackupShortTermRetentionPolicy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersDatabasesBackupShortTermRetentionPolicy_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForServersDatabasesBackupShortTermRetentionPolicy_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersDatabasesBackupShortTermRetentionPolicy_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(BackupShortTermRetentionPolicyPropertiesGenerator())
}
