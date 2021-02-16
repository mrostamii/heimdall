// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: heimdall/types/headers.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Checkpoint struct {
	Proposer   string `protobuf:"bytes,1,opt,name=proposer,proto3" json:"proposer,omitempty"`
	StartBlock uint64 `protobuf:"varint,2,opt,name=start_block,json=startBlock,proto3" json:"start_block,omitempty" yaml:"start_block"`
	EndBlock   uint64 `protobuf:"varint,3,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty" yaml:"end_block"`
	RootHash   string `protobuf:"bytes,4,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty" yaml:"root_hash"`
	BorChainID string `protobuf:"bytes,5,opt,name=BorChainID,proto3" json:"BorChainID,omitempty" yaml:"bor_chain_ID"`
	TimeStamp  uint64 `protobuf:"varint,6,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty" yaml:"time_stamp"`
}

func (m *Checkpoint) Reset()      { *m = Checkpoint{} }
func (*Checkpoint) ProtoMessage() {}
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7e74cbd87a76f33, []int{0}
}
func (m *Checkpoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Checkpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Checkpoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Checkpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkpoint.Merge(m, src)
}
func (m *Checkpoint) XXX_Size() int {
	return m.Size()
}
func (m *Checkpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Checkpoint proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Checkpoint)(nil), "heimdall.types.Checkpoint")
}

func init() { proto.RegisterFile("heimdall/types/headers.proto", fileDescriptor_c7e74cbd87a76f33) }

var fileDescriptor_c7e74cbd87a76f33 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0x93, 0xfe, 0xfb, 0x2f, 0xe9, 0x09, 0xa2, 0xb1, 0x6a, 0x28, 0x92, 0x94, 0x0c, 0xd2,
	0xa9, 0x41, 0x14, 0x0a, 0x9d, 0x24, 0xed, 0x60, 0xd7, 0xb8, 0xb9, 0x84, 0x4b, 0x72, 0xf4, 0x42,
	0x93, 0xbc, 0xe1, 0x72, 0x22, 0xfd, 0x06, 0x8e, 0x8e, 0x8e, 0xfd, 0x28, 0x8e, 0x8e, 0x1d, 0x9d,
	0x82, 0xb4, 0xdf, 0x20, 0x9f, 0x40, 0xee, 0xa2, 0xb1, 0x6e, 0xcf, 0xc3, 0xef, 0xfd, 0xc1, 0x0b,
	0x0f, 0xba, 0xa0, 0x24, 0x4e, 0x23, 0x9c, 0x24, 0x0e, 0x5f, 0xe5, 0xa4, 0x70, 0x28, 0xc1, 0x11,
	0x61, 0xc5, 0x28, 0x67, 0xc0, 0x41, 0x3f, 0xfc, 0xa1, 0x23, 0x49, 0xfb, 0xbd, 0x05, 0x2c, 0x40,
	0x22, 0x47, 0xa4, 0xfa, 0xca, 0x7e, 0x6b, 0x21, 0x34, 0xa5, 0x24, 0x5c, 0xe6, 0x10, 0x67, 0x5c,
	0xef, 0x23, 0x2d, 0x67, 0x90, 0x43, 0x41, 0x98, 0xa1, 0x0e, 0xd4, 0x61, 0xd7, 0x6b, 0xba, 0x3e,
	0x46, 0x07, 0x05, 0xc7, 0x8c, 0xfb, 0x41, 0x02, 0xe1, 0xd2, 0x68, 0x0d, 0xd4, 0x61, 0xdb, 0x3d,
	0xab, 0x4a, 0x4b, 0x5f, 0xe1, 0x34, 0x99, 0xd8, 0x7b, 0xd0, 0xf6, 0x90, 0x6c, 0xae, 0x28, 0xfa,
	0x15, 0xea, 0x92, 0x2c, 0xfa, 0xd6, 0xfe, 0x49, 0xad, 0x57, 0x95, 0xd6, 0x51, 0xad, 0x35, 0xc8,
	0xf6, 0x34, 0x92, 0x45, 0x8d, 0xc2, 0x00, 0xb8, 0x4f, 0x71, 0x41, 0x8d, 0xb6, 0x78, 0x64, 0x5f,
	0x69, 0x90, 0xed, 0x69, 0x22, 0xdf, 0xe1, 0x82, 0xea, 0x63, 0x84, 0x5c, 0x60, 0x53, 0x8a, 0xe3,
	0x6c, 0x3e, 0x33, 0xfe, 0x4b, 0xe7, 0xbc, 0x2a, 0xad, 0x93, 0xda, 0x09, 0x80, 0xf9, 0xa1, 0x80,
	0xfe, 0x7c, 0x66, 0x7b, 0x7b, 0xa7, 0xfa, 0x0d, 0x42, 0x3c, 0x4e, 0x89, 0x5f, 0x70, 0x9c, 0xe6,
	0x46, 0x47, 0xfe, 0x77, 0x5a, 0x95, 0xd6, 0x71, 0x2d, 0xfe, 0x32, 0xdb, 0xeb, 0x8a, 0x72, 0x2f,
	0xf2, 0x44, 0x7b, 0x5e, 0x5b, 0xca, 0xeb, 0xda, 0x52, 0xdc, 0xdb, 0xf7, 0xad, 0xa9, 0x6e, 0xb6,
	0xa6, 0xfa, 0xb9, 0x35, 0xd5, 0x97, 0x9d, 0xa9, 0x6c, 0x76, 0xa6, 0xf2, 0xb1, 0x33, 0x95, 0x87,
	0xcb, 0x45, 0xcc, 0xe9, 0x63, 0x30, 0x0a, 0x21, 0x75, 0x52, 0xcc, 0xe3, 0x30, 0x23, 0xfc, 0x09,
	0xd8, 0xd2, 0xf9, 0x3b, 0x5c, 0xd0, 0x91, 0x5b, 0x5c, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xb0,
	0x6b, 0x6b, 0x23, 0xd1, 0x01, 0x00, 0x00,
}

func (m *Checkpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Checkpoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Checkpoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeStamp != 0 {
		i = encodeVarintHeaders(dAtA, i, uint64(m.TimeStamp))
		i--
		dAtA[i] = 0x30
	}
	if len(m.BorChainID) > 0 {
		i -= len(m.BorChainID)
		copy(dAtA[i:], m.BorChainID)
		i = encodeVarintHeaders(dAtA, i, uint64(len(m.BorChainID)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.RootHash) > 0 {
		i -= len(m.RootHash)
		copy(dAtA[i:], m.RootHash)
		i = encodeVarintHeaders(dAtA, i, uint64(len(m.RootHash)))
		i--
		dAtA[i] = 0x22
	}
	if m.EndBlock != 0 {
		i = encodeVarintHeaders(dAtA, i, uint64(m.EndBlock))
		i--
		dAtA[i] = 0x18
	}
	if m.StartBlock != 0 {
		i = encodeVarintHeaders(dAtA, i, uint64(m.StartBlock))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintHeaders(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHeaders(dAtA []byte, offset int, v uint64) int {
	offset -= sovHeaders(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Checkpoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovHeaders(uint64(l))
	}
	if m.StartBlock != 0 {
		n += 1 + sovHeaders(uint64(m.StartBlock))
	}
	if m.EndBlock != 0 {
		n += 1 + sovHeaders(uint64(m.EndBlock))
	}
	l = len(m.RootHash)
	if l > 0 {
		n += 1 + l + sovHeaders(uint64(l))
	}
	l = len(m.BorChainID)
	if l > 0 {
		n += 1 + l + sovHeaders(uint64(l))
	}
	if m.TimeStamp != 0 {
		n += 1 + sovHeaders(uint64(m.TimeStamp))
	}
	return n
}

func sovHeaders(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHeaders(x uint64) (n int) {
	return sovHeaders(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Checkpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeaders
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
			return fmt.Errorf("proto: Checkpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Checkpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
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
				return ErrInvalidLengthHeaders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHeaders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBlock", wireType)
			}
			m.StartBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlock", wireType)
			}
			m.EndBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
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
				return ErrInvalidLengthHeaders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHeaders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
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
				return ErrInvalidLengthHeaders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHeaders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BorChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeStamp", wireType)
			}
			m.TimeStamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeStamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHeaders(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeaders
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHeaders
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
func skipHeaders(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeaders
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
					return 0, ErrIntOverflowHeaders
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
					return 0, ErrIntOverflowHeaders
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
				return 0, ErrInvalidLengthHeaders
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHeaders
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHeaders
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHeaders        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeaders          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHeaders = fmt.Errorf("proto: unexpected end of group")
)