// Copyright 2020 Prysmatic Labs.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: protos/attestation.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	github_com_prysmaticlabs_go_bitfield "github.com/prysmaticlabs/go-bitfield"
	github_com_theQRL_zond_consensus_types_primitives "github.com/theQRL/zond/consensus-types/primitives"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Attestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A bitfield representation of validator indices that have voted exactly
	// the same vote and have been aggregated into this attestation.
	AggregationBits github_com_prysmaticlabs_go_bitfield.Bitlist `protobuf:"bytes,1,opt,name=aggregation_bits,json=aggregationBits,proto3" json:"aggregation_bits,omitempty" cast-type:"github.com/prysmaticlabs/go-bitfield.Bitlist" ssz-max:"2048"`
	Data            *AttestationData                             `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// 96 byte BLS aggregate signature.
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty" ssz-size:"96"`
}

func (x *Attestation) Reset() {
	*x = Attestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_attestation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation) ProtoMessage() {}

func (x *Attestation) ProtoReflect() protoreflect.Message {
	mi := &file_protos_attestation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation.ProtoReflect.Descriptor instead.
func (*Attestation) Descriptor() ([]byte, []int) {
	return file_protos_attestation_proto_rawDescGZIP(), []int{0}
}

func (x *Attestation) GetAggregationBits() github_com_prysmaticlabs_go_bitfield.Bitlist {
	if x != nil {
		return x.AggregationBits
	}
	return github_com_prysmaticlabs_go_bitfield.Bitlist(nil)
}

func (x *Attestation) GetData() *AttestationData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Attestation) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type AggregateAttestationAndProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The aggregator index that submitted this aggregated attestation and proof.
	AggregatorIndex github_com_theQRL_zond_consensus_types_primitives.ValidatorIndex `protobuf:"varint,1,opt,name=aggregator_index,json=aggregatorIndex,proto3" json:"aggregator_index,omitempty" cast-type:"github.com/theQRL/zond/consensus-types/primitives.ValidatorIndex"`
	// The aggregated attestation that was submitted.
	Aggregate *Attestation `protobuf:"bytes,3,opt,name=aggregate,proto3" json:"aggregate,omitempty"`
	// 96 byte selection proof signed by the aggregator, which is the signature of the slot to aggregate.
	SelectionProof []byte `protobuf:"bytes,2,opt,name=selection_proof,json=selectionProof,proto3" json:"selection_proof,omitempty" ssz-size:"96"`
}

func (x *AggregateAttestationAndProof) Reset() {
	*x = AggregateAttestationAndProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_attestation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregateAttestationAndProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateAttestationAndProof) ProtoMessage() {}

func (x *AggregateAttestationAndProof) ProtoReflect() protoreflect.Message {
	mi := &file_protos_attestation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateAttestationAndProof.ProtoReflect.Descriptor instead.
func (*AggregateAttestationAndProof) Descriptor() ([]byte, []int) {
	return file_protos_attestation_proto_rawDescGZIP(), []int{1}
}

func (x *AggregateAttestationAndProof) GetAggregatorIndex() github_com_theQRL_zond_consensus_types_primitives.ValidatorIndex {
	if x != nil {
		return x.AggregatorIndex
	}
	return github_com_theQRL_zond_consensus_types_primitives.ValidatorIndex(0)
}

func (x *AggregateAttestationAndProof) GetAggregate() *Attestation {
	if x != nil {
		return x.Aggregate
	}
	return nil
}

func (x *AggregateAttestationAndProof) GetSelectionProof() []byte {
	if x != nil {
		return x.SelectionProof
	}
	return nil
}

type SignedAggregateAttestationAndProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The aggregated attestation and selection proof itself.
	Message *AggregateAttestationAndProof `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// 96 byte BLS aggregate signature signed by the aggregator over the message.
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty" ssz-size:"96"`
}

func (x *SignedAggregateAttestationAndProof) Reset() {
	*x = SignedAggregateAttestationAndProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_attestation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedAggregateAttestationAndProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedAggregateAttestationAndProof) ProtoMessage() {}

func (x *SignedAggregateAttestationAndProof) ProtoReflect() protoreflect.Message {
	mi := &file_protos_attestation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedAggregateAttestationAndProof.ProtoReflect.Descriptor instead.
func (*SignedAggregateAttestationAndProof) Descriptor() ([]byte, []int) {
	return file_protos_attestation_proto_rawDescGZIP(), []int{2}
}

func (x *SignedAggregateAttestationAndProof) GetMessage() *AggregateAttestationAndProof {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *SignedAggregateAttestationAndProof) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type AttestationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Slot of the attestation attesting for.
	Slot github_com_theQRL_zond_consensus_types_primitives.Slot `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty" cast-type:"github.com/theQRL/zond/consensus-types/primitives.Slot"`
	// The committee index that submitted this attestation.
	CommitteeIndex github_com_theQRL_zond_consensus_types_primitives.CommitteeIndex `protobuf:"varint,2,opt,name=committee_index,json=committeeIndex,proto3" json:"committee_index,omitempty" cast-type:"github.com/theQRL/zond/consensus-types/primitives.CommitteeIndex"`
	// 32 byte root of the LMD GHOST block vote.
	BeaconBlockRoot []byte `protobuf:"bytes,3,opt,name=beacon_block_root,json=beaconBlockRoot,proto3" json:"beacon_block_root,omitempty" ssz-size:"32"`
	// The most recent justified checkpoint in the beacon state
	Source *Checkpoint `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	// The checkpoint attempting to be justified for the current epoch and its epoch boundary block
	Target *Checkpoint `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *AttestationData) Reset() {
	*x = AttestationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_attestation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestationData) ProtoMessage() {}

func (x *AttestationData) ProtoReflect() protoreflect.Message {
	mi := &file_protos_attestation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestationData.ProtoReflect.Descriptor instead.
func (*AttestationData) Descriptor() ([]byte, []int) {
	return file_protos_attestation_proto_rawDescGZIP(), []int{3}
}

func (x *AttestationData) GetSlot() github_com_theQRL_zond_consensus_types_primitives.Slot {
	if x != nil {
		return x.Slot
	}
	return github_com_theQRL_zond_consensus_types_primitives.Slot(0)
}

func (x *AttestationData) GetCommitteeIndex() github_com_theQRL_zond_consensus_types_primitives.CommitteeIndex {
	if x != nil {
		return x.CommitteeIndex
	}
	return github_com_theQRL_zond_consensus_types_primitives.CommitteeIndex(0)
}

func (x *AttestationData) GetBeaconBlockRoot() []byte {
	if x != nil {
		return x.BeaconBlockRoot
	}
	return nil
}

func (x *AttestationData) GetSource() *Checkpoint {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *AttestationData) GetTarget() *Checkpoint {
	if x != nil {
		return x.Target
	}
	return nil
}

type Checkpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Epoch the checkpoint references.
	Epoch github_com_theQRL_zond_consensus_types_primitives.Epoch `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty" cast-type:"github.com/theQRL/zond/consensus-types/primitives.Epoch"`
	// Block root of the checkpoint references.
	Root []byte `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty" ssz-size:"32"`
}

func (x *Checkpoint) Reset() {
	*x = Checkpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_attestation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Checkpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Checkpoint) ProtoMessage() {}

func (x *Checkpoint) ProtoReflect() protoreflect.Message {
	mi := &file_protos_attestation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Checkpoint.ProtoReflect.Descriptor instead.
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return file_protos_attestation_proto_rawDescGZIP(), []int{4}
}

func (x *Checkpoint) GetEpoch() github_com_theQRL_zond_consensus_types_primitives.Epoch {
	if x != nil {
		return x.Epoch
	}
	return github_com_theQRL_zond_consensus_types_primitives.Epoch(0)
}

func (x *Checkpoint) GetRoot() []byte {
	if x != nil {
		return x.Root
	}
	return nil
}

var File_protos_attestation_proto protoreflect.FileDescriptor

var file_protos_attestation_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x0b, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x63, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x38, 0x92, 0xb5, 0x18, 0x04, 0x32, 0x30, 0x34, 0x38, 0x82, 0xb5, 0x18, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d,
	0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x69, 0x74, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x2e, 0x42, 0x69, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x0f, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x74, 0x73, 0x12, 0x2b, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a,
	0xb5, 0x18, 0x02, 0x39, 0x36, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x22, 0xf3, 0x01, 0x0a, 0x1c, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x12, 0x6f, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x44, 0x82, 0xb5, 0x18,
	0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x51,
	0x52, 0x4c, 0x2f, 0x7a, 0x6f, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76,
	0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x31, 0x0a, 0x09, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06,
	0x8a, 0xb5, 0x18, 0x02, 0x39, 0x36, 0x52, 0x0e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x8a, 0x01, 0x0a, 0x22, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x3e, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x39, 0x36, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x22, 0xdc, 0x02, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4e, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x3a, 0x82, 0xb5, 0x18, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x51, 0x52, 0x4c, 0x2f, 0x7a, 0x6f, 0x6e,
	0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x6d, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x44, 0x82, 0xb5, 0x18, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x68, 0x65, 0x51, 0x52, 0x4c, 0x2f, 0x7a, 0x6f, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x69,
	0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x32, 0x0a, 0x11, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x0f, 0x62, 0x65, 0x61, 0x63, 0x6f,
	0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x22, 0x7b, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x51, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x3b, 0x82, 0xb5, 0x18, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x68, 0x65, 0x51, 0x52, 0x4c, 0x2f, 0x7a, 0x6f, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x05, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x12, 0x1a, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x42,
	0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68,
	0x65, 0x51, 0x52, 0x4c, 0x2f, 0x7a, 0x6f, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_attestation_proto_rawDescOnce sync.Once
	file_protos_attestation_proto_rawDescData = file_protos_attestation_proto_rawDesc
)

func file_protos_attestation_proto_rawDescGZIP() []byte {
	file_protos_attestation_proto_rawDescOnce.Do(func() {
		file_protos_attestation_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_attestation_proto_rawDescData)
	})
	return file_protos_attestation_proto_rawDescData
}

var file_protos_attestation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_attestation_proto_goTypes = []interface{}{
	(*Attestation)(nil),                        // 0: protos.Attestation
	(*AggregateAttestationAndProof)(nil),       // 1: protos.AggregateAttestationAndProof
	(*SignedAggregateAttestationAndProof)(nil), // 2: protos.SignedAggregateAttestationAndProof
	(*AttestationData)(nil),                    // 3: protos.AttestationData
	(*Checkpoint)(nil),                         // 4: protos.Checkpoint
}
var file_protos_attestation_proto_depIdxs = []int32{
	3, // 0: protos.Attestation.data:type_name -> protos.AttestationData
	0, // 1: protos.AggregateAttestationAndProof.aggregate:type_name -> protos.Attestation
	1, // 2: protos.SignedAggregateAttestationAndProof.message:type_name -> protos.AggregateAttestationAndProof
	4, // 3: protos.AttestationData.source:type_name -> protos.Checkpoint
	4, // 4: protos.AttestationData.target:type_name -> protos.Checkpoint
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_attestation_proto_init() }
func file_protos_attestation_proto_init() {
	if File_protos_attestation_proto != nil {
		return
	}
	file_protos_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_attestation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attestation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_attestation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregateAttestationAndProof); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_attestation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedAggregateAttestationAndProof); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_attestation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_attestation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Checkpoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_attestation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_attestation_proto_goTypes,
		DependencyIndexes: file_protos_attestation_proto_depIdxs,
		MessageInfos:      file_protos_attestation_proto_msgTypes,
	}.Build()
	File_protos_attestation_proto = out.File
	file_protos_attestation_proto_rawDesc = nil
	file_protos_attestation_proto_goTypes = nil
	file_protos_attestation_proto_depIdxs = nil
}