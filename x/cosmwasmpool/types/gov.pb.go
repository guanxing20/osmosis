// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/cosmwasmpool/v1beta1/gov.proto

package types

import (
	bytes "bytes"
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

// UploadCosmWasmPoolCodeAndWhiteListProposal is a gov Content type for
// uploading coswasm pool code and adding it to internal whitelist. Only the
// code ids created by this message are eligible for being x/cosmwasmpool pools.
type UploadCosmWasmPoolCodeAndWhiteListProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// WASMByteCode can be raw or gzip compressed
	WASMByteCode []byte `protobuf:"bytes,3,opt,name=wasm_byte_code,json=wasmByteCode,proto3" json:"wasm_byte_code,omitempty"`
}

func (m *UploadCosmWasmPoolCodeAndWhiteListProposal) Reset() {
	*m = UploadCosmWasmPoolCodeAndWhiteListProposal{}
}
func (*UploadCosmWasmPoolCodeAndWhiteListProposal) ProtoMessage() {}
func (*UploadCosmWasmPoolCodeAndWhiteListProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_c184a48c55bbcf5c, []int{0}
}
func (m *UploadCosmWasmPoolCodeAndWhiteListProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UploadCosmWasmPoolCodeAndWhiteListProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UploadCosmWasmPoolCodeAndWhiteListProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UploadCosmWasmPoolCodeAndWhiteListProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadCosmWasmPoolCodeAndWhiteListProposal.Merge(m, src)
}
func (m *UploadCosmWasmPoolCodeAndWhiteListProposal) XXX_Size() int {
	return m.Size()
}
func (m *UploadCosmWasmPoolCodeAndWhiteListProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadCosmWasmPoolCodeAndWhiteListProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UploadCosmWasmPoolCodeAndWhiteListProposal proto.InternalMessageInfo

// MigratePoolContractsProposal is a gov Content type for
// migrating  given pools to the new contract code and adding to internal
// whitelist if needed. It has two options to perform the migration:
//
// 1. If the codeID is non-zero, it will migrate the pool contracts to a given
// codeID assuming that it has already been uploaded. uploadByteCode must be
// empty in such a case. Fails if codeID does not exist. Fails if uploadByteCode
// is not empty.
//
// 2. If the codeID is zero, it will upload the given uploadByteCode and use the
// new resulting code id to migrate the pool to. Errors if uploadByteCode is
// empty or invalid.
//
// In both cases, if one of the pools specified by the given poolID does not
// exist, the proposal fails.
//
// The reason for having poolIDs be a slice of ids is to account for the
// potential need for emergency migration of all old code ids associated with
// particular pools to new code ids, or simply having the flexibility of
// migrating multiple older pool contracts to a new one at once when there is a
// release.
//
// poolD count to be submitted at once is gated by a governance paramets (20 at
// launch). The proposal fails if more. Note that 20 was chosen arbitrarily to
// have a constant bound on the number of pools migrated at once. This size will
// be configured by a module parameter so it can be changed by a constant.
type MigratePoolContractsProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// pool_ids are the pool ids of the contracts to be migrated
	// either to the new_code_id that is already uploaded to chain or to
	// the given wasm_byte_code.
	PoolIds []uint64 `protobuf:"varint,3,rep,packed,name=pool_ids,json=poolIds,proto3" json:"pool_ids,omitempty"`
	// new_code_id is the code id of the contract code to migrate to.
	// Assumes that the code is already uploaded to chain. Only one of
	// new_code_id and wasm_byte_code should be set.
	NewCodeId uint64 `protobuf:"varint,4,opt,name=new_code_id,json=newCodeId,proto3" json:"new_code_id,omitempty"`
	// WASMByteCode can be raw or gzip compressed. Assumes that the code id
	// has not been uploaded yet so uploads the given code and migrates to it.
	// Only one of new_code_id and wasm_byte_code should be set.
	WASMByteCode []byte `protobuf:"bytes,5,opt,name=wasm_byte_code,json=wasmByteCode,proto3" json:"wasm_byte_code,omitempty"`
	// MigrateMsg migrate message to be used for migrating the pool contracts.
	MigrateMsg []byte `protobuf:"bytes,6,opt,name=migrate_msg,json=migrateMsg,proto3" json:"migrate_msg,omitempty"`
}

func (m *MigratePoolContractsProposal) Reset()      { *m = MigratePoolContractsProposal{} }
func (*MigratePoolContractsProposal) ProtoMessage() {}
func (*MigratePoolContractsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_c184a48c55bbcf5c, []int{1}
}
func (m *MigratePoolContractsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MigratePoolContractsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MigratePoolContractsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MigratePoolContractsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MigratePoolContractsProposal.Merge(m, src)
}
func (m *MigratePoolContractsProposal) XXX_Size() int {
	return m.Size()
}
func (m *MigratePoolContractsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_MigratePoolContractsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_MigratePoolContractsProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UploadCosmWasmPoolCodeAndWhiteListProposal)(nil), "osmosis.cosmwasmpool.v1beta1.UploadCosmWasmPoolCodeAndWhiteListProposal")
	proto.RegisterType((*MigratePoolContractsProposal)(nil), "osmosis.cosmwasmpool.v1beta1.MigratePoolContractsProposal")
}

func init() {
	proto.RegisterFile("osmosis/cosmwasmpool/v1beta1/gov.proto", fileDescriptor_c184a48c55bbcf5c)
}

var fileDescriptor_c184a48c55bbcf5c = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xb1, 0x8b, 0xd4, 0x40,
	0x14, 0xc6, 0x33, 0xee, 0xde, 0xe9, 0xcd, 0x06, 0x91, 0x70, 0x45, 0x94, 0x23, 0x09, 0x57, 0x48,
	0x10, 0x4c, 0x3c, 0x0e, 0x44, 0xec, 0x6e, 0xaf, 0x3a, 0x70, 0xe1, 0x88, 0xc8, 0x82, 0x4d, 0x98,
	0x64, 0x86, 0xdc, 0x40, 0x26, 0x2f, 0xe4, 0x8d, 0xbb, 0xee, 0x7f, 0x60, 0x69, 0x69, 0xb9, 0xb5,
	0x7f, 0x89, 0xe5, 0x96, 0x56, 0x22, 0xd9, 0xc6, 0x3f, 0xc1, 0x52, 0x26, 0xc9, 0xc2, 0x2e, 0x58,
	0x08, 0x76, 0xf3, 0xbd, 0xf9, 0x1e, 0xf3, 0xfb, 0x86, 0x8f, 0x3e, 0x05, 0x54, 0x80, 0x12, 0xe3,
	0x1c, 0x50, 0x2d, 0x19, 0xaa, 0x1a, 0xa0, 0x8c, 0x17, 0x17, 0x99, 0xd0, 0xec, 0x22, 0x2e, 0x60,
	0x11, 0xd5, 0x0d, 0x68, 0x70, 0xce, 0x06, 0x5f, 0xb4, 0xef, 0x8b, 0x06, 0xdf, 0x93, 0xd3, 0x02,
	0x0a, 0xe8, 0x8c, 0xb1, 0x39, 0xf5, 0x3b, 0xe7, 0x5f, 0x09, 0x7d, 0xf6, 0xae, 0x2e, 0x81, 0xf1,
	0x6b, 0x40, 0x35, 0x67, 0xa8, 0x6e, 0x01, 0xca, 0x6b, 0xe0, 0xe2, 0xaa, 0xe2, 0xf3, 0x3b, 0xa9,
	0xc5, 0x1b, 0x89, 0xfa, 0xb6, 0x81, 0x1a, 0x90, 0x95, 0xce, 0x29, 0x3d, 0xd2, 0x52, 0x97, 0xc2,
	0x25, 0x01, 0x09, 0x4f, 0x92, 0x5e, 0x38, 0x01, 0x9d, 0x70, 0x81, 0x79, 0x23, 0x6b, 0x2d, 0xa1,
	0x72, 0xef, 0x75, 0x77, 0xfb, 0x23, 0xe7, 0x25, 0x7d, 0x68, 0x80, 0xd2, 0x6c, 0xa5, 0x45, 0x9a,
	0x03, 0x17, 0xee, 0x28, 0x20, 0xa1, 0x3d, 0x7d, 0xd4, 0xfe, 0xf0, 0xed, 0xf9, 0xd5, 0xdb, 0xd9,
	0x74, 0xa5, 0x85, 0x79, 0x35, 0xb1, 0x8d, 0x6f, 0xa7, 0x5e, 0xdb, 0x9f, 0xd6, 0xbe, 0xf5, 0x65,
	0xed, 0x5b, 0xbf, 0xd6, 0x3e, 0x39, 0xff, 0x4d, 0xe8, 0xd9, 0x4c, 0x16, 0x0d, 0xd3, 0xa2, 0xa7,
	0xac, 0x74, 0xc3, 0x72, 0x8d, 0xff, 0x8d, 0xf7, 0x98, 0x3e, 0x30, 0x7f, 0x95, 0x4a, 0x8e, 0xee,
	0x28, 0x18, 0x85, 0xe3, 0xe4, 0xbe, 0xd1, 0x37, 0x1c, 0x1d, 0x8f, 0x4e, 0x2a, 0xb1, 0xec, 0x98,
	0x53, 0xc9, 0xdd, 0x71, 0x40, 0xc2, 0x71, 0x72, 0x52, 0x89, 0xa5, 0xe1, 0xbb, 0xe1, 0x7f, 0x49,
	0x76, 0xf4, 0x2f, 0xc9, 0x1c, 0x9f, 0x4e, 0x54, 0x1f, 0x25, 0x55, 0x58, 0xb8, 0xc7, 0x66, 0x29,
	0xa1, 0xc3, 0x68, 0x86, 0xc5, 0x61, 0xf4, 0x69, 0xf2, 0xad, 0xf5, 0xc8, 0xa6, 0xf5, 0xc8, 0xcf,
	0xd6, 0x23, 0x9f, 0xb7, 0x9e, 0xb5, 0xd9, 0x7a, 0xd6, 0xf7, 0xad, 0x67, 0xbd, 0x7f, 0x55, 0x48,
	0x7d, 0xf7, 0x21, 0x8b, 0x72, 0x50, 0xf1, 0x50, 0x80, 0xe7, 0x25, 0xcb, 0x70, 0x27, 0xe2, 0xc5,
	0xe5, 0x8b, 0xf8, 0xe3, 0x61, 0x77, 0xf4, 0xaa, 0x16, 0x98, 0x1d, 0x77, 0x15, 0xb8, 0xfc, 0x13,
	0x00, 0x00, 0xff, 0xff, 0xc3, 0x9c, 0x16, 0xfe, 0x60, 0x02, 0x00, 0x00,
}

func (this *UploadCosmWasmPoolCodeAndWhiteListProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UploadCosmWasmPoolCodeAndWhiteListProposal)
	if !ok {
		that2, ok := that.(UploadCosmWasmPoolCodeAndWhiteListProposal)
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
	if !bytes.Equal(this.WASMByteCode, that1.WASMByteCode) {
		return false
	}
	return true
}
func (this *MigratePoolContractsProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MigratePoolContractsProposal)
	if !ok {
		that2, ok := that.(MigratePoolContractsProposal)
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
	if len(this.PoolIds) != len(that1.PoolIds) {
		return false
	}
	for i := range this.PoolIds {
		if this.PoolIds[i] != that1.PoolIds[i] {
			return false
		}
	}
	if this.NewCodeId != that1.NewCodeId {
		return false
	}
	if !bytes.Equal(this.WASMByteCode, that1.WASMByteCode) {
		return false
	}
	if !bytes.Equal(this.MigrateMsg, that1.MigrateMsg) {
		return false
	}
	return true
}
func (m *UploadCosmWasmPoolCodeAndWhiteListProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadCosmWasmPoolCodeAndWhiteListProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UploadCosmWasmPoolCodeAndWhiteListProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WASMByteCode) > 0 {
		i -= len(m.WASMByteCode)
		copy(dAtA[i:], m.WASMByteCode)
		i = encodeVarintGov(dAtA, i, uint64(len(m.WASMByteCode)))
		i--
		dAtA[i] = 0x1a
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

func (m *MigratePoolContractsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MigratePoolContractsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MigratePoolContractsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MigrateMsg) > 0 {
		i -= len(m.MigrateMsg)
		copy(dAtA[i:], m.MigrateMsg)
		i = encodeVarintGov(dAtA, i, uint64(len(m.MigrateMsg)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.WASMByteCode) > 0 {
		i -= len(m.WASMByteCode)
		copy(dAtA[i:], m.WASMByteCode)
		i = encodeVarintGov(dAtA, i, uint64(len(m.WASMByteCode)))
		i--
		dAtA[i] = 0x2a
	}
	if m.NewCodeId != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.NewCodeId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.PoolIds) > 0 {
		dAtA2 := make([]byte, len(m.PoolIds)*10)
		var j1 int
		for _, num := range m.PoolIds {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintGov(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x1a
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
func (m *UploadCosmWasmPoolCodeAndWhiteListProposal) Size() (n int) {
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
	l = len(m.WASMByteCode)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func (m *MigratePoolContractsProposal) Size() (n int) {
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
	if len(m.PoolIds) > 0 {
		l = 0
		for _, e := range m.PoolIds {
			l += sovGov(uint64(e))
		}
		n += 1 + sovGov(uint64(l)) + l
	}
	if m.NewCodeId != 0 {
		n += 1 + sovGov(uint64(m.NewCodeId))
	}
	l = len(m.WASMByteCode)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.MigrateMsg)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UploadCosmWasmPoolCodeAndWhiteListProposal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UploadCosmWasmPoolCodeAndWhiteListProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadCosmWasmPoolCodeAndWhiteListProposal: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field WASMByteCode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WASMByteCode = append(m.WASMByteCode[:0], dAtA[iNdEx:postIndex]...)
			if m.WASMByteCode == nil {
				m.WASMByteCode = []byte{}
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
func (m *MigratePoolContractsProposal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MigratePoolContractsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MigratePoolContractsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.PoolIds = append(m.PoolIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGov
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGov
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGov
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.PoolIds) == 0 {
					m.PoolIds = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGov
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.PoolIds = append(m.PoolIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolIds", wireType)
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewCodeId", wireType)
			}
			m.NewCodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewCodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WASMByteCode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WASMByteCode = append(m.WASMByteCode[:0], dAtA[iNdEx:postIndex]...)
			if m.WASMByteCode == nil {
				m.WASMByteCode = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MigrateMsg", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MigrateMsg = append(m.MigrateMsg[:0], dAtA[iNdEx:postIndex]...)
			if m.MigrateMsg == nil {
				m.MigrateMsg = []byte{}
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
