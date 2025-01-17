/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// OrdererNodeEnrollmentApplyConfiguration represents an declarative configuration of the OrdererNodeEnrollment type for use
// with apply.
type OrdererNodeEnrollmentApplyConfiguration struct {
	TLS *OrdererNodeEnrollmentTLSApplyConfiguration `json:"tls,omitempty"`
}

// OrdererNodeEnrollmentApplyConfiguration constructs an declarative configuration of the OrdererNodeEnrollment type for use with
// apply.
func OrdererNodeEnrollment() *OrdererNodeEnrollmentApplyConfiguration {
	return &OrdererNodeEnrollmentApplyConfiguration{}
}

// WithTLS sets the TLS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLS field is set to the value of the last call.
func (b *OrdererNodeEnrollmentApplyConfiguration) WithTLS(value *OrdererNodeEnrollmentTLSApplyConfiguration) *OrdererNodeEnrollmentApplyConfiguration {
	b.TLS = value
	return b
}
