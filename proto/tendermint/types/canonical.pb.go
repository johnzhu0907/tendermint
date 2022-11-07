// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/types/canonical.proto

package types

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CanonicalBlockID struct {
	Hash          []byte                 `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	PartSetHeader CanonicalPartSetHeader `protobuf:"bytes,2,opt,name=part_set_header,json=partSetHeader,proto3" json:"part_set_header"`
}

func (m *CanonicalBlockID) Reset()         { *m = CanonicalBlockID{} }
func (m *CanonicalBlockID) String() string { return proto.CompactTextString(m) }
func (*CanonicalBlockID) ProtoMessage()    {}
func (*CanonicalBlockID) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d1a1a84ff7267ed, []int{0}
}
func (m *CanonicalBlockID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CanonicalBlockID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CanonicalBlockID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CanonicalBlockID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CanonicalBlockID.Merge(m, src)
}
func (m *CanonicalBlockID) XXX_Size() int {
	return m.Size()
}
func (m *CanonicalBlockID) XXX_DiscardUnknown() {
	xxx_messageInfo_CanonicalBlockID.DiscardUnknown(m)
}

var xxx_messageInfo_CanonicalBlockID proto.InternalMessageInfo

func (m *CanonicalBlockID) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *CanonicalBlockID) GetPartSetHeader() CanonicalPartSetHeader {
	if m != nil {
		return m.PartSetHeader
	}
	return CanonicalPartSetHeader{}
}

type CanonicalPartSetHeader struct {
	Total uint32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Hash  []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *CanonicalPartSetHeader) Reset()         { *m = CanonicalPartSetHeader{} }
func (m *CanonicalPartSetHeader) String() string { return proto.CompactTextString(m) }
func (*CanonicalPartSetHeader) ProtoMessage()    {}
func (*CanonicalPartSetHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d1a1a84ff7267ed, []int{1}
}
func (m *CanonicalPartSetHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CanonicalPartSetHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CanonicalPartSetHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CanonicalPartSetHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CanonicalPartSetHeader.Merge(m, src)
}
func (m *CanonicalPartSetHeader) XXX_Size() int {
	return m.Size()
}
func (m *CanonicalPartSetHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_CanonicalPartSetHeader.DiscardUnknown(m)
}

var xxx_messageInfo_CanonicalPartSetHeader proto.InternalMessageInfo

func (m *CanonicalPartSetHeader) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *CanonicalPartSetHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type CanonicalProposal struct {
	Type      SignedMsgType     `protobuf:"varint,1,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"type,omitempty"`
	Height    int64             `protobuf:"fixed64,2,opt,name=height,proto3" json:"height,omitempty"`
	Round     int64             `protobuf:"fixed64,3,opt,name=round,proto3" json:"round,omitempty"`
	POLRound  int64             `protobuf:"varint,4,opt,name=pol_round,json=polRound,proto3" json:"pol_round,omitempty"`
	BlockID   *CanonicalBlockID `protobuf:"bytes,5,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Timestamp time.Time         `protobuf:"bytes,6,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	ChainID   string            `protobuf:"bytes,7,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *CanonicalProposal) Reset()         { *m = CanonicalProposal{} }
func (m *CanonicalProposal) String() string { return proto.CompactTextString(m) }
func (*CanonicalProposal) ProtoMessage()    {}
func (*CanonicalProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d1a1a84ff7267ed, []int{2}
}
func (m *CanonicalProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CanonicalProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CanonicalProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CanonicalProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CanonicalProposal.Merge(m, src)
}
func (m *CanonicalProposal) XXX_Size() int {
	return m.Size()
}
func (m *CanonicalProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CanonicalProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CanonicalProposal proto.InternalMessageInfo

func (m *CanonicalProposal) GetType() SignedMsgType {
	if m != nil {
		return m.Type
	}
	return UnknownType
}

func (m *CanonicalProposal) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *CanonicalProposal) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *CanonicalProposal) GetPOLRound() int64 {
	if m != nil {
		return m.POLRound
	}
	return 0
}

func (m *CanonicalProposal) GetBlockID() *CanonicalBlockID {
	if m != nil {
		return m.BlockID
	}
	return nil
}

func (m *CanonicalProposal) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *CanonicalProposal) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

type CanonicalVote struct {
	Type          SignedMsgType        `protobuf:"varint,1,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"type,omitempty"`
	Height        int64                `protobuf:"fixed64,2,opt,name=height,proto3" json:"height,omitempty"`
	Round         int64                `protobuf:"fixed64,3,opt,name=round,proto3" json:"round,omitempty"`
	BlockID       *CanonicalBlockID    `protobuf:"bytes,4,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Timestamp     time.Time            `protobuf:"bytes,5,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	ChainID       string               `protobuf:"bytes,6,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	VoteExtension *VoteExtensionToSign `protobuf:"bytes,7,opt,name=vote_extension,json=voteExtension,proto3" json:"vote_extension,omitempty"`
}

func (m *CanonicalVote) Reset()         { *m = CanonicalVote{} }
func (m *CanonicalVote) String() string { return proto.CompactTextString(m) }
func (*CanonicalVote) ProtoMessage()    {}
func (*CanonicalVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d1a1a84ff7267ed, []int{3}
}
func (m *CanonicalVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CanonicalVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CanonicalVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CanonicalVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CanonicalVote.Merge(m, src)
}
func (m *CanonicalVote) XXX_Size() int {
	return m.Size()
}
func (m *CanonicalVote) XXX_DiscardUnknown() {
	xxx_messageInfo_CanonicalVote.DiscardUnknown(m)
}

var xxx_messageInfo_CanonicalVote proto.InternalMessageInfo

func (m *CanonicalVote) GetType() SignedMsgType {
	if m != nil {
		return m.Type
	}
	return UnknownType
}

func (m *CanonicalVote) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *CanonicalVote) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *CanonicalVote) GetBlockID() *CanonicalBlockID {
	if m != nil {
		return m.BlockID
	}
	return nil
}

func (m *CanonicalVote) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *CanonicalVote) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *CanonicalVote) GetVoteExtension() *VoteExtensionToSign {
	if m != nil {
		return m.VoteExtension
	}
	return nil
}

func init() {
	proto.RegisterType((*CanonicalBlockID)(nil), "tendermint.types.CanonicalBlockID")
	proto.RegisterType((*CanonicalPartSetHeader)(nil), "tendermint.types.CanonicalPartSetHeader")
	proto.RegisterType((*CanonicalProposal)(nil), "tendermint.types.CanonicalProposal")
	proto.RegisterType((*CanonicalVote)(nil), "tendermint.types.CanonicalVote")
}

func init() { proto.RegisterFile("tendermint/types/canonical.proto", fileDescriptor_8d1a1a84ff7267ed) }

var fileDescriptor_8d1a1a84ff7267ed = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0xe3, 0xd4, 0x49, 0x9c, 0x4b, 0x53, 0xc2, 0xa9, 0xaa, 0xac, 0x08, 0xd9, 0x96, 0x25,
	0x90, 0x59, 0x6c, 0x29, 0x1d, 0xd8, 0x5d, 0x90, 0x08, 0x2a, 0xa2, 0x5c, 0xa3, 0x0e, 0x2c, 0xd6,
	0xc5, 0x3e, 0x6c, 0x0b, 0xc7, 0x67, 0xd9, 0x97, 0x8a, 0x2e, 0x7c, 0x86, 0x7e, 0xac, 0x8e, 0x1d,
	0x61, 0x09, 0xc8, 0xf9, 0x12, 0x8c, 0xe8, 0xce, 0x49, 0x1c, 0x25, 0xc0, 0x02, 0xea, 0x12, 0xbd,
	0x7f, 0x1e, 0xbf, 0xef, 0xa3, 0xdf, 0xab, 0x1c, 0x30, 0x18, 0x49, 0x03, 0x92, 0xcf, 0xe2, 0x94,
	0x39, 0xec, 0x26, 0x23, 0x85, 0xe3, 0xe3, 0x94, 0xa6, 0xb1, 0x8f, 0x13, 0x3b, 0xcb, 0x29, 0xa3,
	0x70, 0x50, 0x2b, 0x6c, 0xa1, 0x18, 0x1e, 0x87, 0x34, 0xa4, 0xa2, 0xe9, 0xf0, 0xa8, 0xd2, 0x0d,
	0x9f, 0xec, 0x4d, 0x12, 0xbf, 0xab, 0xae, 0x1e, 0x52, 0x1a, 0x26, 0xc4, 0x11, 0xd9, 0x74, 0xfe,
	0xd1, 0x61, 0xf1, 0x8c, 0x14, 0x0c, 0xcf, 0xb2, 0x4a, 0x60, 0x7e, 0x01, 0x83, 0xb3, 0xf5, 0x66,
	0x37, 0xa1, 0xfe, 0xa7, 0xf1, 0x4b, 0x08, 0x81, 0x1c, 0xe1, 0x22, 0x52, 0x25, 0x43, 0xb2, 0x0e,
	0x91, 0x88, 0xe1, 0x15, 0x78, 0x94, 0xe1, 0x9c, 0x79, 0x05, 0x61, 0x5e, 0x44, 0x70, 0x40, 0x72,
	0xb5, 0x69, 0x48, 0x56, 0x6f, 0x64, 0xd9, 0xbb, 0x46, 0xed, 0xcd, 0xc0, 0x0b, 0x9c, 0xb3, 0x4b,
	0xc2, 0x5e, 0x0b, 0xbd, 0x2b, 0xdf, 0x2d, 0xf4, 0x06, 0xea, 0x67, 0xdb, 0x45, 0xd3, 0x05, 0x27,
	0xbf, 0x97, 0xc3, 0x63, 0xd0, 0x62, 0x94, 0xe1, 0x44, 0xd8, 0xe8, 0xa3, 0x2a, 0xd9, 0x78, 0x6b,
	0xd6, 0xde, 0xcc, 0x6f, 0x4d, 0xf0, 0xb8, 0x1e, 0x92, 0xd3, 0x8c, 0x16, 0x38, 0x81, 0xa7, 0x40,
	0xe6, 0x76, 0xc4, 0xe7, 0x47, 0x23, 0x7d, 0xdf, 0xe6, 0x65, 0x1c, 0xa6, 0x24, 0x78, 0x5b, 0x84,
	0x93, 0x9b, 0x8c, 0x20, 0x21, 0x86, 0x27, 0xa0, 0x1d, 0x91, 0x38, 0x8c, 0x98, 0x58, 0x30, 0x40,
	0xab, 0x8c, 0x9b, 0xc9, 0xe9, 0x3c, 0x0d, 0xd4, 0x03, 0x51, 0xae, 0x12, 0xf8, 0x1c, 0x74, 0x33,
	0x9a, 0x78, 0x55, 0x47, 0x36, 0x24, 0xeb, 0xc0, 0x3d, 0x2c, 0x17, 0xba, 0x72, 0xf1, 0xee, 0x1c,
	0xf1, 0x1a, 0x52, 0x32, 0x9a, 0x88, 0x08, 0xbe, 0x01, 0xca, 0x94, 0xe3, 0xf5, 0xe2, 0x40, 0x6d,
	0x09, 0x70, 0xe6, 0x5f, 0xc0, 0xad, 0x2e, 0xe1, 0xf6, 0xca, 0x85, 0xde, 0x59, 0x25, 0xa8, 0x23,
	0x06, 0x8c, 0x03, 0xe8, 0x82, 0xee, 0xe6, 0x8c, 0x6a, 0x5b, 0x0c, 0x1b, 0xda, 0xd5, 0xa1, 0xed,
	0xf5, 0xa1, 0xed, 0xc9, 0x5a, 0xe1, 0x2a, 0x9c, 0xfb, 0xed, 0x77, 0x5d, 0x42, 0xf5, 0x67, 0xf0,
	0x19, 0x50, 0xfc, 0x08, 0xc7, 0x29, 0xf7, 0xd3, 0x31, 0x24, 0xab, 0x5b, 0xed, 0x3a, 0xe3, 0x35,
	0xbe, 0x4b, 0x34, 0xc7, 0x81, 0xf9, 0xb3, 0x09, 0xfa, 0x1b, 0x5b, 0x57, 0x94, 0x91, 0x87, 0xe0,
	0xba, 0x0d, 0x4b, 0xfe, 0x9f, 0xb0, 0x5a, 0xff, 0x0e, 0xab, 0xfd, 0x67, 0x58, 0xf0, 0x1c, 0x1c,
	0x5d, 0x53, 0x46, 0x3c, 0xf2, 0x99, 0x91, 0xb4, 0x88, 0x69, 0x2a, 0xd0, 0xf6, 0x46, 0x4f, 0xf7,
	0xdd, 0x73, 0x94, 0xaf, 0xd6, 0xb2, 0x09, 0xe5, 0xcc, 0x50, 0xff, 0x7a, 0xbb, 0xe8, 0xbe, 0xbf,
	0x2b, 0x35, 0xe9, 0xbe, 0xd4, 0xa4, 0x1f, 0xa5, 0x26, 0xdd, 0x2e, 0xb5, 0xc6, 0xfd, 0x52, 0x6b,
	0x7c, 0x5d, 0x6a, 0x8d, 0x0f, 0x2f, 0xc2, 0x98, 0x45, 0xf3, 0xa9, 0xed, 0xd3, 0x99, 0xb3, 0xfd,
	0xf7, 0xaf, 0xc3, 0xea, 0x99, 0xd8, 0x7d, 0x1a, 0xa6, 0x6d, 0x51, 0x3f, 0xfd, 0x15, 0x00, 0x00,
	0xff, 0xff, 0x4e, 0x61, 0xcb, 0xc4, 0x7f, 0x04, 0x00, 0x00,
}

func (m *CanonicalBlockID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CanonicalBlockID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CanonicalBlockID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PartSetHeader.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCanonical(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintCanonical(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CanonicalPartSetHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CanonicalPartSetHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CanonicalPartSetHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintCanonical(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Total != 0 {
		i = encodeVarintCanonical(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CanonicalProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CanonicalProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CanonicalProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintCanonical(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x3a
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintCanonical(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x32
	if m.BlockID != nil {
		{
			size, err := m.BlockID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCanonical(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.POLRound != 0 {
		i = encodeVarintCanonical(dAtA, i, uint64(m.POLRound))
		i--
		dAtA[i] = 0x20
	}
	if m.Round != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Round))
		i--
		dAtA[i] = 0x19
	}
	if m.Height != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Height))
		i--
		dAtA[i] = 0x11
	}
	if m.Type != 0 {
		i = encodeVarintCanonical(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CanonicalVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CanonicalVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CanonicalVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VoteExtension != nil {
		{
			size, err := m.VoteExtension.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCanonical(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintCanonical(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x32
	}
	n5, err5 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintCanonical(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0x2a
	if m.BlockID != nil {
		{
			size, err := m.BlockID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCanonical(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Round != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Round))
		i--
		dAtA[i] = 0x19
	}
	if m.Height != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Height))
		i--
		dAtA[i] = 0x11
	}
	if m.Type != 0 {
		i = encodeVarintCanonical(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCanonical(dAtA []byte, offset int, v uint64) int {
	offset -= sovCanonical(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CanonicalBlockID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovCanonical(uint64(l))
	}
	l = m.PartSetHeader.Size()
	n += 1 + l + sovCanonical(uint64(l))
	return n
}

func (m *CanonicalPartSetHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Total != 0 {
		n += 1 + sovCanonical(uint64(m.Total))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovCanonical(uint64(l))
	}
	return n
}

func (m *CanonicalProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCanonical(uint64(m.Type))
	}
	if m.Height != 0 {
		n += 9
	}
	if m.Round != 0 {
		n += 9
	}
	if m.POLRound != 0 {
		n += 1 + sovCanonical(uint64(m.POLRound))
	}
	if m.BlockID != nil {
		l = m.BlockID.Size()
		n += 1 + l + sovCanonical(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovCanonical(uint64(l))
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovCanonical(uint64(l))
	}
	return n
}

func (m *CanonicalVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCanonical(uint64(m.Type))
	}
	if m.Height != 0 {
		n += 9
	}
	if m.Round != 0 {
		n += 9
	}
	if m.BlockID != nil {
		l = m.BlockID.Size()
		n += 1 + l + sovCanonical(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovCanonical(uint64(l))
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovCanonical(uint64(l))
	}
	if m.VoteExtension != nil {
		l = m.VoteExtension.Size()
		n += 1 + l + sovCanonical(uint64(l))
	}
	return n
}

func sovCanonical(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCanonical(x uint64) (n int) {
	return sovCanonical(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CanonicalBlockID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCanonical
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
			return fmt.Errorf("proto: CanonicalBlockID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CanonicalBlockID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartSetHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PartSetHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCanonical(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCanonical
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
func (m *CanonicalPartSetHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCanonical
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
			return fmt.Errorf("proto: CanonicalPartSetHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CanonicalPartSetHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCanonical(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCanonical
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
func (m *CanonicalProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCanonical
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
			return fmt.Errorf("proto: CanonicalProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CanonicalProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= SignedMsgType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Height = int64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Round = int64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field POLRound", wireType)
			}
			m.POLRound = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.POLRound |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockID == nil {
				m.BlockID = &CanonicalBlockID{}
			}
			if err := m.BlockID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCanonical(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCanonical
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
func (m *CanonicalVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCanonical
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
			return fmt.Errorf("proto: CanonicalVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CanonicalVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= SignedMsgType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Height = int64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Round = int64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockID == nil {
				m.BlockID = &CanonicalBlockID{}
			}
			if err := m.BlockID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteExtension", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VoteExtension == nil {
				m.VoteExtension = &VoteExtensionToSign{}
			}
			if err := m.VoteExtension.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCanonical(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCanonical
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
func skipCanonical(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCanonical
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
					return 0, ErrIntOverflowCanonical
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
					return 0, ErrIntOverflowCanonical
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
				return 0, ErrInvalidLengthCanonical
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCanonical
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCanonical
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCanonical        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCanonical          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCanonical = fmt.Errorf("proto: unexpected end of group")
)
