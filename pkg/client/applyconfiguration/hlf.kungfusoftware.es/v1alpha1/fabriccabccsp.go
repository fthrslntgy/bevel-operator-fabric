/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricCABCCSPApplyConfiguration represents an declarative configuration of the FabricCABCCSP type for use
// with apply.
type FabricCABCCSPApplyConfiguration struct {
	Default *string                            `json:"default,omitempty"`
	SW      *FabricCABCCSPSWApplyConfiguration `json:"sw,omitempty"`
}

// FabricCABCCSPApplyConfiguration constructs an declarative configuration of the FabricCABCCSP type for use with
// apply.
func FabricCABCCSP() *FabricCABCCSPApplyConfiguration {
	return &FabricCABCCSPApplyConfiguration{}
}

// WithDefault sets the Default field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Default field is set to the value of the last call.
func (b *FabricCABCCSPApplyConfiguration) WithDefault(value string) *FabricCABCCSPApplyConfiguration {
	b.Default = &value
	return b
}

// WithSW sets the SW field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SW field is set to the value of the last call.
func (b *FabricCABCCSPApplyConfiguration) WithSW(value *FabricCABCCSPSWApplyConfiguration) *FabricCABCCSPApplyConfiguration {
	b.SW = value
	return b
}
