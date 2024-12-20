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

func Test_ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded, ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded runs a test to see if a specific instance of ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded(subject ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded
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

// Generator of ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded instances for property
// testing - lazily instantiated by ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator()
var applicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator gopter.Gen

// ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator returns a generator of ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded instances for property testing.
func ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator() gopter.Gen {
	if applicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator != nil {
		return applicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded(generators)
	applicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded{}), generators)

	return applicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_Delegation_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Delegation via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDelegation, DelegationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDelegation runs a test to see if a specific instance of Delegation round trips to JSON and back losslessly
func RunJSONSerializationTestForDelegation(subject Delegation) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Delegation
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

// Generator of Delegation instances for property testing - lazily instantiated by DelegationGenerator()
var delegationGenerator gopter.Gen

// DelegationGenerator returns a generator of Delegation instances for property testing.
// We first initialize delegationGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DelegationGenerator() gopter.Gen {
	if delegationGenerator != nil {
		return delegationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDelegation(generators)
	delegationGenerator = gen.Struct(reflect.TypeOf(Delegation{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDelegation(generators)
	AddRelatedPropertyGeneratorsForDelegation(generators)
	delegationGenerator = gen.Struct(reflect.TypeOf(Delegation{}), generators)

	return delegationGenerator
}

// AddIndependentPropertyGeneratorsForDelegation is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDelegation(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDelegation is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDelegation(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServiceDelegationPropertiesFormatGenerator())
}

func Test_NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded, NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded runs a test to see if a specific instance of NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded(subject NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded
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

// Generator of NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded instances for property testing -
// lazily instantiated by NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator()
var networkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator gopter.Gen

// NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator returns a generator of NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded instances for property testing.
func NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator() gopter.Gen {
	if networkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator != nil {
		return networkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded(generators)
	networkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded{}), generators)

	return networkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded, RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded runs a test to see if a specific instance of RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded(subject RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded
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

// Generator of RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded instances for property testing - lazily
// instantiated by RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator()
var routeTableSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator gopter.Gen

// RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator returns a generator of RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded instances for property testing.
func RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator() gopter.Gen {
	if routeTableSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator != nil {
		return routeTableSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded(generators)
	routeTableSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded{}), generators)

	return routeTableSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForRouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServiceDelegationPropertiesFormat_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServiceDelegationPropertiesFormat via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServiceDelegationPropertiesFormat, ServiceDelegationPropertiesFormatGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServiceDelegationPropertiesFormat runs a test to see if a specific instance of ServiceDelegationPropertiesFormat round trips to JSON and back losslessly
func RunJSONSerializationTestForServiceDelegationPropertiesFormat(subject ServiceDelegationPropertiesFormat) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServiceDelegationPropertiesFormat
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

// Generator of ServiceDelegationPropertiesFormat instances for property testing - lazily instantiated by
// ServiceDelegationPropertiesFormatGenerator()
var serviceDelegationPropertiesFormatGenerator gopter.Gen

// ServiceDelegationPropertiesFormatGenerator returns a generator of ServiceDelegationPropertiesFormat instances for property testing.
func ServiceDelegationPropertiesFormatGenerator() gopter.Gen {
	if serviceDelegationPropertiesFormatGenerator != nil {
		return serviceDelegationPropertiesFormatGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServiceDelegationPropertiesFormat(generators)
	serviceDelegationPropertiesFormatGenerator = gen.Struct(reflect.TypeOf(ServiceDelegationPropertiesFormat{}), generators)

	return serviceDelegationPropertiesFormatGenerator
}

// AddIndependentPropertyGeneratorsForServiceDelegationPropertiesFormat is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServiceDelegationPropertiesFormat(gens map[string]gopter.Gen) {
	gens["ServiceName"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded, ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded runs a test to see if a specific instance of ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded(subject ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded
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

// Generator of ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded instances for property testing -
// lazily instantiated by ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator()
var serviceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator gopter.Gen

// ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator returns a generator of ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded instances for property testing.
func ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator() gopter.Gen {
	if serviceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator != nil {
		return serviceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded(generators)
	serviceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded{}), generators)

	return serviceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServiceEndpointPropertiesFormat_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServiceEndpointPropertiesFormat via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServiceEndpointPropertiesFormat, ServiceEndpointPropertiesFormatGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServiceEndpointPropertiesFormat runs a test to see if a specific instance of ServiceEndpointPropertiesFormat round trips to JSON and back losslessly
func RunJSONSerializationTestForServiceEndpointPropertiesFormat(subject ServiceEndpointPropertiesFormat) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServiceEndpointPropertiesFormat
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

// Generator of ServiceEndpointPropertiesFormat instances for property testing - lazily instantiated by
// ServiceEndpointPropertiesFormatGenerator()
var serviceEndpointPropertiesFormatGenerator gopter.Gen

// ServiceEndpointPropertiesFormatGenerator returns a generator of ServiceEndpointPropertiesFormat instances for property testing.
// We first initialize serviceEndpointPropertiesFormatGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServiceEndpointPropertiesFormatGenerator() gopter.Gen {
	if serviceEndpointPropertiesFormatGenerator != nil {
		return serviceEndpointPropertiesFormatGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServiceEndpointPropertiesFormat(generators)
	serviceEndpointPropertiesFormatGenerator = gen.Struct(reflect.TypeOf(ServiceEndpointPropertiesFormat{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServiceEndpointPropertiesFormat(generators)
	AddRelatedPropertyGeneratorsForServiceEndpointPropertiesFormat(generators)
	serviceEndpointPropertiesFormatGenerator = gen.Struct(reflect.TypeOf(ServiceEndpointPropertiesFormat{}), generators)

	return serviceEndpointPropertiesFormatGenerator
}

// AddIndependentPropertyGeneratorsForServiceEndpointPropertiesFormat is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServiceEndpointPropertiesFormat(gens map[string]gopter.Gen) {
	gens["Locations"] = gen.SliceOf(gen.AlphaString())
	gens["Service"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServiceEndpointPropertiesFormat is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServiceEndpointPropertiesFormat(gens map[string]gopter.Gen) {
	gens["NetworkIdentifier"] = gen.PtrOf(SubResourceGenerator())
}

func Test_SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded, SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded runs a test to see if a specific instance of SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForSubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded(subject SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded
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

// Generator of SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded instances for property testing -
// lazily instantiated by SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator()
var subnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator gopter.Gen

// SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator returns a generator of SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded instances for property testing.
// We first initialize subnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator() gopter.Gen {
	if subnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator != nil {
		return subnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded(generators)
	subnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded(generators)
	AddRelatedPropertyGeneratorsForSubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded(generators)
	subnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded{}), generators)

	return subnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForSubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["AddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["AddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DefaultOutboundAccess"] = gen.PtrOf(gen.Bool())
	gens["PrivateEndpointNetworkPolicies"] = gen.PtrOf(gen.OneConstOf(
		SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_Disabled,
		SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_Enabled,
		SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_NetworkSecurityGroupEnabled,
		SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_RouteTableEnabled))
	gens["PrivateLinkServiceNetworkPolicies"] = gen.PtrOf(gen.OneConstOf(SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_Disabled, SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_Enabled))
	gens["SharingScope"] = gen.PtrOf(gen.OneConstOf(SubnetPropertiesFormat_SharingScope_DelegatedServices, SubnetPropertiesFormat_SharingScope_Tenant))
}

// AddRelatedPropertyGeneratorsForSubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["ApplicationGatewayIPConfigurations"] = gen.SliceOf(ApplicationGatewayIPConfiguration_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator())
	gens["Delegations"] = gen.SliceOf(DelegationGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResourceGenerator())
	gens["NatGateway"] = gen.PtrOf(SubResourceGenerator())
	gens["NetworkSecurityGroup"] = gen.PtrOf(NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator())
	gens["RouteTable"] = gen.PtrOf(RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator())
	gens["ServiceEndpointPolicies"] = gen.SliceOf(ServiceEndpointPolicySpec_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator())
	gens["ServiceEndpoints"] = gen.SliceOf(ServiceEndpointPropertiesFormatGenerator())
}

func Test_VirtualNetworksSubnet_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworksSubnet_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksSubnet_Spec, VirtualNetworksSubnet_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksSubnet_Spec runs a test to see if a specific instance of VirtualNetworksSubnet_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksSubnet_Spec(subject VirtualNetworksSubnet_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworksSubnet_Spec
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

// Generator of VirtualNetworksSubnet_Spec instances for property testing - lazily instantiated by
// VirtualNetworksSubnet_SpecGenerator()
var virtualNetworksSubnet_SpecGenerator gopter.Gen

// VirtualNetworksSubnet_SpecGenerator returns a generator of VirtualNetworksSubnet_Spec instances for property testing.
// We first initialize virtualNetworksSubnet_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksSubnet_SpecGenerator() gopter.Gen {
	if virtualNetworksSubnet_SpecGenerator != nil {
		return virtualNetworksSubnet_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSubnet_Spec(generators)
	virtualNetworksSubnet_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksSubnet_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSubnet_Spec(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksSubnet_Spec(generators)
	virtualNetworksSubnet_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksSubnet_Spec{}), generators)

	return virtualNetworksSubnet_SpecGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksSubnet_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksSubnet_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForVirtualNetworksSubnet_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksSubnet_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SubnetPropertiesFormat_VirtualNetworks_Subnet_SubResourceEmbeddedGenerator())
}
