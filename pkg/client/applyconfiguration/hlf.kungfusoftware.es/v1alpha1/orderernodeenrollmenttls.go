/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// OrdererNodeEnrollmentTLSApplyConfiguration represents an declarative configuration of the OrdererNodeEnrollmentTLS type for use
// with apply.
type OrdererNodeEnrollmentTLSApplyConfiguration struct {
	Csr *CsrApplyConfiguration `json:"csr,omitempty"`
}

// OrdererNodeEnrollmentTLSApplyConfiguration constructs an declarative configuration of the OrdererNodeEnrollmentTLS type for use with
// apply.
func OrdererNodeEnrollmentTLS() *OrdererNodeEnrollmentTLSApplyConfiguration {
	return &OrdererNodeEnrollmentTLSApplyConfiguration{}
}

// WithCsr sets the Csr field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Csr field is set to the value of the last call.
func (b *OrdererNodeEnrollmentTLSApplyConfiguration) WithCsr(value *CsrApplyConfiguration) *OrdererNodeEnrollmentTLSApplyConfiguration {
	b.Csr = value
	return b
}
