/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// ComponentApplyConfiguration represents an declarative configuration of the Component type for use
// with apply.
type ComponentApplyConfiguration struct {
	Cahost       *string                                `json:"cahost,omitempty"`
	Caname       *string                                `json:"caname,omitempty"`
	Caport       *int                                   `json:"caport,omitempty"`
	Catls        *CatlsApplyConfiguration               `json:"catls,omitempty"`
	Enrollid     *string                                `json:"enrollid,omitempty"`
	Enrollsecret *string                                `json:"enrollsecret,omitempty"`
	External     *ExternalCertificateApplyConfiguration `json:"external,omitempty"`
}

// ComponentApplyConfiguration constructs an declarative configuration of the Component type for use with
// apply.
func Component() *ComponentApplyConfiguration {
	return &ComponentApplyConfiguration{}
}

// WithCahost sets the Cahost field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cahost field is set to the value of the last call.
func (b *ComponentApplyConfiguration) WithCahost(value string) *ComponentApplyConfiguration {
	b.Cahost = &value
	return b
}

// WithCaname sets the Caname field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Caname field is set to the value of the last call.
func (b *ComponentApplyConfiguration) WithCaname(value string) *ComponentApplyConfiguration {
	b.Caname = &value
	return b
}

// WithCaport sets the Caport field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Caport field is set to the value of the last call.
func (b *ComponentApplyConfiguration) WithCaport(value int) *ComponentApplyConfiguration {
	b.Caport = &value
	return b
}

// WithCatls sets the Catls field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Catls field is set to the value of the last call.
func (b *ComponentApplyConfiguration) WithCatls(value *CatlsApplyConfiguration) *ComponentApplyConfiguration {
	b.Catls = value
	return b
}

// WithEnrollid sets the Enrollid field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enrollid field is set to the value of the last call.
func (b *ComponentApplyConfiguration) WithEnrollid(value string) *ComponentApplyConfiguration {
	b.Enrollid = &value
	return b
}

// WithEnrollsecret sets the Enrollsecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enrollsecret field is set to the value of the last call.
func (b *ComponentApplyConfiguration) WithEnrollsecret(value string) *ComponentApplyConfiguration {
	b.Enrollsecret = &value
	return b
}

// WithExternal sets the External field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the External field is set to the value of the last call.
func (b *ComponentApplyConfiguration) WithExternal(value *ExternalCertificateApplyConfiguration) *ComponentApplyConfiguration {
	b.External = value
	return b
}
