// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/concentratedliquidity/v1beta1/gov.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// CreateConcentratedLiquidityPoolsProposal is a gov Content type for creating
// concentrated liquidity pools. If a CreateConcentratedLiquidityPoolsProposal
// passes, the pools are created via pool manager module account.
type CreateConcentratedLiquidityPoolsProposal struct {
	Title       string       `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	PoolRecords []PoolRecord `protobuf:"bytes,3,rep,name=pool_records,json=poolRecords,proto3" json:"pool_records" yaml:"pool_records"`
}

func (m *CreateConcentratedLiquidityPoolsProposal) Reset() {
	*m = CreateConcentratedLiquidityPoolsProposal{}
}
func (*CreateConcentratedLiquidityPoolsProposal) ProtoMessage() {}
func (*CreateConcentratedLiquidityPoolsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a96adc35f4989ef7, []int{0}
}
func (m *CreateConcentratedLiquidityPoolsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateConcentratedLiquidityPoolsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateConcentratedLiquidityPoolsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateConcentratedLiquidityPoolsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConcentratedLiquidityPoolsProposal.Merge(m, src)
}
func (m *CreateConcentratedLiquidityPoolsProposal) XXX_Size() int {
	return m.Size()
}
func (m *CreateConcentratedLiquidityPoolsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConcentratedLiquidityPoolsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConcentratedLiquidityPoolsProposal proto.InternalMessageInfo

// TickSpacingDecreaseProposal is a gov Content type for proposing a tick
// spacing decrease for a pool. The proposal will fail if one of the pools do
// not exist, or if the new tick spacing is not less than the current tick
// spacing.
type TickSpacingDecreaseProposal struct {
	Title                      string                      `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description                string                      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	PoolIdToTickSpacingRecords []PoolIdToTickSpacingRecord `protobuf:"bytes,3,rep,name=pool_id_to_tick_spacing_records,json=poolIdToTickSpacingRecords,proto3" json:"pool_id_to_tick_spacing_records"`
}

func (m *TickSpacingDecreaseProposal) Reset()      { *m = TickSpacingDecreaseProposal{} }
func (*TickSpacingDecreaseProposal) ProtoMessage() {}
func (*TickSpacingDecreaseProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a96adc35f4989ef7, []int{1}
}
func (m *TickSpacingDecreaseProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TickSpacingDecreaseProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TickSpacingDecreaseProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TickSpacingDecreaseProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickSpacingDecreaseProposal.Merge(m, src)
}
func (m *TickSpacingDecreaseProposal) XXX_Size() int {
	return m.Size()
}
func (m *TickSpacingDecreaseProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_TickSpacingDecreaseProposal.DiscardUnknown(m)
}

var xxx_messageInfo_TickSpacingDecreaseProposal proto.InternalMessageInfo

// PoolIdToTickSpacingRecord is a struct that contains a pool id to new tick
// spacing pair.
type PoolIdToTickSpacingRecord struct {
	PoolId         uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	NewTickSpacing uint64 `protobuf:"varint,2,opt,name=new_tick_spacing,json=newTickSpacing,proto3" json:"new_tick_spacing,omitempty"`
}

func (m *PoolIdToTickSpacingRecord) Reset()         { *m = PoolIdToTickSpacingRecord{} }
func (m *PoolIdToTickSpacingRecord) String() string { return proto.CompactTextString(m) }
func (*PoolIdToTickSpacingRecord) ProtoMessage()    {}
func (*PoolIdToTickSpacingRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_a96adc35f4989ef7, []int{2}
}
func (m *PoolIdToTickSpacingRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolIdToTickSpacingRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolIdToTickSpacingRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolIdToTickSpacingRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolIdToTickSpacingRecord.Merge(m, src)
}
func (m *PoolIdToTickSpacingRecord) XXX_Size() int {
	return m.Size()
}
func (m *PoolIdToTickSpacingRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolIdToTickSpacingRecord.DiscardUnknown(m)
}

var xxx_messageInfo_PoolIdToTickSpacingRecord proto.InternalMessageInfo

func (m *PoolIdToTickSpacingRecord) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *PoolIdToTickSpacingRecord) GetNewTickSpacing() uint64 {
	if m != nil {
		return m.NewTickSpacing
	}
	return 0
}

type PoolRecord struct {
	Denom0       string                      `protobuf:"bytes,1,opt,name=denom0,proto3" json:"denom0,omitempty" yaml:"denom0"`
	Denom1       string                      `protobuf:"bytes,2,opt,name=denom1,proto3" json:"denom1,omitempty" yaml:"denom1"`
	TickSpacing  uint64                      `protobuf:"varint,3,opt,name=tick_spacing,json=tickSpacing,proto3" json:"tick_spacing,omitempty" yaml:"tick_spacing"`
	SpreadFactor cosmossdk_io_math.LegacyDec `protobuf:"bytes,5,opt,name=spread_factor,json=spreadFactor,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"spread_factor" yaml:"spread_factor"`
}

func (m *PoolRecord) Reset()         { *m = PoolRecord{} }
func (m *PoolRecord) String() string { return proto.CompactTextString(m) }
func (*PoolRecord) ProtoMessage()    {}
func (*PoolRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_a96adc35f4989ef7, []int{3}
}
func (m *PoolRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolRecord.Merge(m, src)
}
func (m *PoolRecord) XXX_Size() int {
	return m.Size()
}
func (m *PoolRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolRecord.DiscardUnknown(m)
}

var xxx_messageInfo_PoolRecord proto.InternalMessageInfo

func (m *PoolRecord) GetDenom0() string {
	if m != nil {
		return m.Denom0
	}
	return ""
}

func (m *PoolRecord) GetDenom1() string {
	if m != nil {
		return m.Denom1
	}
	return ""
}

func (m *PoolRecord) GetTickSpacing() uint64 {
	if m != nil {
		return m.TickSpacing
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateConcentratedLiquidityPoolsProposal)(nil), "osmosis.concentratedliquidity.v1beta1.CreateConcentratedLiquidityPoolsProposal")
	proto.RegisterType((*TickSpacingDecreaseProposal)(nil), "osmosis.concentratedliquidity.v1beta1.TickSpacingDecreaseProposal")
	proto.RegisterType((*PoolIdToTickSpacingRecord)(nil), "osmosis.concentratedliquidity.v1beta1.PoolIdToTickSpacingRecord")
	proto.RegisterType((*PoolRecord)(nil), "osmosis.concentratedliquidity.v1beta1.PoolRecord")
}

func init() {
	proto.RegisterFile("osmosis/concentratedliquidity/v1beta1/gov.proto", fileDescriptor_a96adc35f4989ef7)
}

var fileDescriptor_a96adc35f4989ef7 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xb5, 0x5b, 0x27, 0xdf, 0xc7, 0x24, 0x45, 0xc5, 0x44, 0x6a, 0x68, 0x24, 0x3b, 0xb2, 0x84,
	0x14, 0x16, 0xb5, 0x6b, 0xba, 0x0b, 0x1b, 0x94, 0x56, 0x48, 0xa0, 0x2e, 0x2a, 0xd3, 0x15, 0x42,
	0x0a, 0x93, 0xf1, 0xe0, 0x8e, 0xe2, 0xf8, 0xba, 0x9e, 0x69, 0x4a, 0xde, 0x00, 0x09, 0x16, 0x2c,
	0x59, 0xe6, 0x71, 0xba, 0xec, 0x12, 0xb1, 0x88, 0x50, 0xb2, 0x61, 0x4b, 0x9e, 0x00, 0x65, 0x26,
	0x3f, 0x4e, 0xa1, 0x12, 0x88, 0x9d, 0x67, 0xe6, 0x9c, 0x73, 0xcf, 0xb9, 0xd7, 0x17, 0x79, 0xc0,
	0x7b, 0xc0, 0x19, 0xf7, 0x08, 0x24, 0x84, 0x26, 0x22, 0xc3, 0x82, 0x86, 0x31, 0x3b, 0xbf, 0x60,
	0x21, 0x13, 0x03, 0xaf, 0xef, 0x77, 0xa8, 0xc0, 0xbe, 0x17, 0x41, 0xdf, 0x4d, 0x33, 0x10, 0x60,
	0x3e, 0x9c, 0x13, 0xdc, 0xdf, 0x12, 0xdc, 0x39, 0x61, 0xb7, 0x12, 0x41, 0x04, 0x92, 0xe1, 0xcd,
	0xbe, 0x14, 0xd9, 0x99, 0xe8, 0xa8, 0x71, 0x98, 0x51, 0x2c, 0xe8, 0x61, 0x8e, 0x7d, 0xbc, 0x60,
	0x9f, 0x00, 0xc4, 0xfc, 0x24, 0x83, 0x14, 0x38, 0x8e, 0xcd, 0x0a, 0x2a, 0x08, 0x26, 0x62, 0x5a,
	0xd5, 0xeb, 0x7a, 0xe3, 0x4e, 0xa0, 0x0e, 0x66, 0x1d, 0x95, 0x42, 0xca, 0x49, 0xc6, 0x52, 0xc1,
	0x20, 0xa9, 0x6e, 0xc8, 0xb7, 0xfc, 0x95, 0x79, 0x8e, 0xca, 0x29, 0x40, 0xdc, 0xce, 0x28, 0x81,
	0x2c, 0xe4, 0xd5, 0xcd, 0xfa, 0x66, 0xa3, 0xf4, 0xd8, 0x77, 0xff, 0xc8, 0xb8, 0x3b, 0xf3, 0x10,
	0x48, 0x66, 0xab, 0x76, 0x35, 0xb2, 0xb5, 0xe9, 0xc8, 0xbe, 0x3f, 0xc0, 0xbd, 0xb8, 0xe9, 0xe4,
	0x45, 0x9d, 0xa0, 0x94, 0x2e, 0x81, 0xbc, 0x59, 0x7e, 0x3f, 0xb4, 0xb5, 0xcf, 0x43, 0x5b, 0xfb,
	0x3e, 0xb4, 0x75, 0xe7, 0x87, 0x8e, 0x6a, 0xa7, 0x8c, 0x74, 0x5f, 0xa6, 0x98, 0xb0, 0x24, 0x3a,
	0xa2, 0x24, 0xa3, 0x98, 0xd3, 0x7f, 0x0e, 0xf6, 0x41, 0x47, 0xb6, 0x34, 0xc1, 0xc2, 0xb6, 0x80,
	0xb6, 0x60, 0xa4, 0xdb, 0xe6, 0xaa, 0xc6, 0x8d, 0xb0, 0x4f, 0xff, 0x22, 0xec, 0xf3, 0xf0, 0x14,
	0x72, 0x6e, 0xe7, 0xd9, 0x8d, 0x59, 0xf6, 0x60, 0x37, 0xbd, 0x0d, 0x70, 0x33, 0x73, 0x88, 0x1e,
	0xdc, 0x2a, 0x66, 0xee, 0xa0, 0xff, 0xe6, 0xbe, 0x65, 0x64, 0x23, 0x28, 0x2a, 0x5d, 0xb3, 0x81,
	0xb6, 0x13, 0x7a, 0xb9, 0x96, 0x44, 0x06, 0x37, 0x82, 0xbb, 0x09, 0xbd, 0xcc, 0x09, 0x35, 0x0d,
	0x59, 0xe5, 0xe3, 0x06, 0x42, 0xab, 0x01, 0x99, 0x8f, 0x50, 0x31, 0xa4, 0x09, 0xf4, 0xf6, 0x55,
	0x27, 0x5b, 0xf7, 0xa6, 0x23, 0x7b, 0x4b, 0x0d, 0x4b, 0xdd, 0x3b, 0xc1, 0x1c, 0xb0, 0x84, 0xfa,
	0xaa, 0xb1, 0xbf, 0x40, 0xfd, 0x05, 0xd4, 0x37, 0x9b, 0xa8, 0xbc, 0x66, 0x68, 0x73, 0x66, 0xa8,
	0xb5, 0xb3, 0xfa, 0x11, 0xf2, 0xaf, 0x4e, 0x50, 0x12, 0x2b, 0x9b, 0xe6, 0x1b, 0xb4, 0xc5, 0xd3,
	0x8c, 0xe2, 0xb0, 0xfd, 0x16, 0x13, 0x01, 0x59, 0xb5, 0x20, 0xab, 0x3d, 0x99, 0x75, 0xf3, 0xeb,
	0xc8, 0xae, 0x11, 0x39, 0x17, 0x1e, 0x76, 0x5d, 0x06, 0x5e, 0x0f, 0x8b, 0x33, 0xf7, 0x98, 0x46,
	0x98, 0x0c, 0x8e, 0x28, 0x99, 0x8e, 0xec, 0x8a, 0xd2, 0x5f, 0x53, 0x70, 0x82, 0xb2, 0x3a, 0x3f,
	0x93, 0x47, 0xd5, 0x88, 0x17, 0xc6, 0xff, 0xc6, 0x76, 0xa1, 0xf5, 0xfa, 0x6a, 0x6c, 0xe9, 0xd7,
	0x63, 0x4b, 0xff, 0x36, 0xb6, 0xf4, 0x4f, 0x13, 0x4b, 0xbb, 0x9e, 0x58, 0xda, 0x97, 0x89, 0xa5,
	0xbd, 0x6a, 0x45, 0x4c, 0x9c, 0x5d, 0x74, 0x5c, 0x02, 0xbd, 0xc5, 0x86, 0xef, 0xc5, 0xb8, 0xc3,
	0x97, 0xeb, 0xde, 0x3f, 0xd8, 0xf7, 0xde, 0xad, 0x2d, 0xfd, 0xde, 0x6a, 0xeb, 0xc5, 0x20, 0xa5,
	0xbc, 0x53, 0x94, 0x3b, 0x7b, 0xf0, 0x33, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x5c, 0x07, 0x2f, 0x23,
	0x04, 0x00, 0x00,
}

func (this *CreateConcentratedLiquidityPoolsProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateConcentratedLiquidityPoolsProposal)
	if !ok {
		that2, ok := that.(CreateConcentratedLiquidityPoolsProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.PoolRecords) != len(that1.PoolRecords) {
		return false
	}
	for i := range this.PoolRecords {
		if !this.PoolRecords[i].Equal(&that1.PoolRecords[i]) {
			return false
		}
	}
	return true
}
func (this *TickSpacingDecreaseProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TickSpacingDecreaseProposal)
	if !ok {
		that2, ok := that.(TickSpacingDecreaseProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.PoolIdToTickSpacingRecords) != len(that1.PoolIdToTickSpacingRecords) {
		return false
	}
	for i := range this.PoolIdToTickSpacingRecords {
		if !this.PoolIdToTickSpacingRecords[i].Equal(&that1.PoolIdToTickSpacingRecords[i]) {
			return false
		}
	}
	return true
}
func (this *PoolIdToTickSpacingRecord) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PoolIdToTickSpacingRecord)
	if !ok {
		that2, ok := that.(PoolIdToTickSpacingRecord)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PoolId != that1.PoolId {
		return false
	}
	if this.NewTickSpacing != that1.NewTickSpacing {
		return false
	}
	return true
}
func (this *PoolRecord) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PoolRecord)
	if !ok {
		that2, ok := that.(PoolRecord)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Denom0 != that1.Denom0 {
		return false
	}
	if this.Denom1 != that1.Denom1 {
		return false
	}
	if this.TickSpacing != that1.TickSpacing {
		return false
	}
	if !this.SpreadFactor.Equal(that1.SpreadFactor) {
		return false
	}
	return true
}
func (m *CreateConcentratedLiquidityPoolsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateConcentratedLiquidityPoolsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateConcentratedLiquidityPoolsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolRecords) > 0 {
		for iNdEx := len(m.PoolRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TickSpacingDecreaseProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TickSpacingDecreaseProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TickSpacingDecreaseProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolIdToTickSpacingRecords) > 0 {
		for iNdEx := len(m.PoolIdToTickSpacingRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolIdToTickSpacingRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PoolIdToTickSpacingRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolIdToTickSpacingRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolIdToTickSpacingRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NewTickSpacing != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.NewTickSpacing))
		i--
		dAtA[i] = 0x10
	}
	if m.PoolId != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PoolRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SpreadFactor.Size()
		i -= size
		if _, err := m.SpreadFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.TickSpacing != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.TickSpacing))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Denom1) > 0 {
		i -= len(m.Denom1)
		copy(dAtA[i:], m.Denom1)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Denom1)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom0) > 0 {
		i -= len(m.Denom0)
		copy(dAtA[i:], m.Denom0)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Denom0)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreateConcentratedLiquidityPoolsProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.PoolRecords) > 0 {
		for _, e := range m.PoolRecords {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func (m *TickSpacingDecreaseProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.PoolIdToTickSpacingRecords) > 0 {
		for _, e := range m.PoolIdToTickSpacingRecords {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func (m *PoolIdToTickSpacingRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovGov(uint64(m.PoolId))
	}
	if m.NewTickSpacing != 0 {
		n += 1 + sovGov(uint64(m.NewTickSpacing))
	}
	return n
}

func (m *PoolRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom0)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Denom1)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if m.TickSpacing != 0 {
		n += 1 + sovGov(uint64(m.TickSpacing))
	}
	l = m.SpreadFactor.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateConcentratedLiquidityPoolsProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateConcentratedLiquidityPoolsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateConcentratedLiquidityPoolsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolRecords = append(m.PoolRecords, PoolRecord{})
			if err := m.PoolRecords[len(m.PoolRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TickSpacingDecreaseProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TickSpacingDecreaseProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickSpacingDecreaseProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolIdToTickSpacingRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolIdToTickSpacingRecords = append(m.PoolIdToTickSpacingRecords, PoolIdToTickSpacingRecord{})
			if err := m.PoolIdToTickSpacingRecords[len(m.PoolIdToTickSpacingRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PoolIdToTickSpacingRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PoolIdToTickSpacingRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolIdToTickSpacingRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewTickSpacing", wireType)
			}
			m.NewTickSpacing = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewTickSpacing |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PoolRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PoolRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom0", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom0 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickSpacing", wireType)
			}
			m.TickSpacing = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickSpacing |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpreadFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SpreadFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGov
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGov
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
