// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

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

func Test_Topic_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Topic_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopic_Spec_ARM, Topic_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopic_Spec_ARM runs a test to see if a specific instance of Topic_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTopic_Spec_ARM(subject Topic_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Topic_Spec_ARM
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

// Generator of Topic_Spec_ARM instances for property testing - lazily instantiated by Topic_Spec_ARMGenerator()
var topic_Spec_ARMGenerator gopter.Gen

// Topic_Spec_ARMGenerator returns a generator of Topic_Spec_ARM instances for property testing.
func Topic_Spec_ARMGenerator() gopter.Gen {
	if topic_Spec_ARMGenerator != nil {
		return topic_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopic_Spec_ARM(generators)
	topic_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Topic_Spec_ARM{}), generators)

	return topic_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForTopic_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTopic_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}