/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// FabricPeerResourcesApplyConfiguration represents an declarative configuration of the FabricPeerResources type for use
// with apply.
type FabricPeerResourcesApplyConfiguration struct {
	Peer            *v1.ResourceRequirements `json:"peer,omitempty"`
	CouchDB         *v1.ResourceRequirements `json:"couchdb,omitempty"`
	Chaincode       *v1.ResourceRequirements `json:"chaincode,omitempty"`
	CouchDBExporter *v1.ResourceRequirements `json:"couchdbExporter,omitempty"`
	Proxy           *v1.ResourceRequirements `json:"proxy,omitempty"`
}

// FabricPeerResourcesApplyConfiguration constructs an declarative configuration of the FabricPeerResources type for use with
// apply.
func FabricPeerResources() *FabricPeerResourcesApplyConfiguration {
	return &FabricPeerResourcesApplyConfiguration{}
}

// WithPeer sets the Peer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Peer field is set to the value of the last call.
func (b *FabricPeerResourcesApplyConfiguration) WithPeer(value v1.ResourceRequirements) *FabricPeerResourcesApplyConfiguration {
	b.Peer = &value
	return b
}

// WithCouchDB sets the CouchDB field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CouchDB field is set to the value of the last call.
func (b *FabricPeerResourcesApplyConfiguration) WithCouchDB(value v1.ResourceRequirements) *FabricPeerResourcesApplyConfiguration {
	b.CouchDB = &value
	return b
}

// WithChaincode sets the Chaincode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Chaincode field is set to the value of the last call.
func (b *FabricPeerResourcesApplyConfiguration) WithChaincode(value v1.ResourceRequirements) *FabricPeerResourcesApplyConfiguration {
	b.Chaincode = &value
	return b
}

// WithCouchDBExporter sets the CouchDBExporter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CouchDBExporter field is set to the value of the last call.
func (b *FabricPeerResourcesApplyConfiguration) WithCouchDBExporter(value v1.ResourceRequirements) *FabricPeerResourcesApplyConfiguration {
	b.CouchDBExporter = &value
	return b
}

// WithProxy sets the Proxy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Proxy field is set to the value of the last call.
func (b *FabricPeerResourcesApplyConfiguration) WithProxy(value v1.ResourceRequirements) *FabricPeerResourcesApplyConfiguration {
	b.Proxy = &value
	return b
}
