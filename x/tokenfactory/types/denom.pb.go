// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sixswingchain/tokenfactory/denom.proto

package types

import (
	fmt "fmt"
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

type Denom struct {
	Denom              string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Description        string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Ticker             string `protobuf:"bytes,3,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Precision          int32  `protobuf:"varint,4,opt,name=precision,proto3" json:"precision,omitempty"`
	Url                string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	MaxSupply          int32  `protobuf:"varint,6,opt,name=maxSupply,proto3" json:"maxSupply,omitempty"`
	Supply             int32  `protobuf:"varint,7,opt,name=supply,proto3" json:"supply,omitempty"`
	CanChangeMaxSupply bool   `protobuf:"varint,8,opt,name=canChangeMaxSupply,proto3" json:"canChangeMaxSupply,omitempty"`
	Owner              string `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *Denom) Reset()         { *m = Denom{} }
func (m *Denom) String() string { return proto.CompactTextString(m) }
func (*Denom) ProtoMessage()    {}
func (*Denom) Descriptor() ([]byte, []int) {
	return fileDescriptor_b31d4da25876980c, []int{0}
}
func (m *Denom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Denom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Denom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Denom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Denom.Merge(m, src)
}
func (m *Denom) XXX_Size() int {
	return m.Size()
}
func (m *Denom) XXX_DiscardUnknown() {
	xxx_messageInfo_Denom.DiscardUnknown(m)
}

var xxx_messageInfo_Denom proto.InternalMessageInfo

func (m *Denom) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Denom) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Denom) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *Denom) GetPrecision() int32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *Denom) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Denom) GetMaxSupply() int32 {
	if m != nil {
		return m.MaxSupply
	}
	return 0
}

func (m *Denom) GetSupply() int32 {
	if m != nil {
		return m.Supply
	}
	return 0
}

func (m *Denom) GetCanChangeMaxSupply() bool {
	if m != nil {
		return m.CanChangeMaxSupply
	}
	return false
}

func (m *Denom) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterType((*Denom)(nil), "sixswingchain.tokenfactory.Denom")
}

func init() {
	proto.RegisterFile("sixswingchain/tokenfactory/denom.proto", fileDescriptor_b31d4da25876980c)
}

var fileDescriptor_b31d4da25876980c = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0xe3, 0x96, 0x84, 0xc6, 0x2c, 0xc8, 0x42, 0xc8, 0x42, 0xc8, 0x8a, 0x3a, 0xa0, 0x4c,
	0xe9, 0xc0, 0xca, 0x04, 0xac, 0x2c, 0x61, 0x63, 0x33, 0xae, 0x69, 0xad, 0xb6, 0xb6, 0x65, 0xbb,
	0x6a, 0x32, 0xf3, 0x02, 0x3c, 0x16, 0x63, 0x47, 0x46, 0x94, 0xbc, 0x08, 0xb2, 0x53, 0x95, 0x44,
	0xea, 0xe6, 0xbb, 0xfb, 0x4e, 0xd6, 0x7f, 0xf0, 0xce, 0x8a, 0xca, 0xee, 0x84, 0x5c, 0xb0, 0x25,
	0x15, 0x72, 0xe6, 0xd4, 0x8a, 0xcb, 0x0f, 0xca, 0x9c, 0x32, 0xf5, 0x6c, 0xce, 0xa5, 0xda, 0x14,
	0xda, 0x28, 0xa7, 0xd0, 0xcd, 0x80, 0x2b, 0xfa, 0xdc, 0xf4, 0x73, 0x04, 0xe3, 0x67, 0xcf, 0xa2,
	0x2b, 0x18, 0x87, 0x12, 0x06, 0x19, 0xc8, 0xd3, 0xb2, 0x13, 0x28, 0x83, 0x17, 0x73, 0x6e, 0x99,
	0x11, 0xda, 0x09, 0x25, 0xf1, 0x28, 0x64, 0x7d, 0x0b, 0x5d, 0xc3, 0xc4, 0x09, 0xb6, 0xe2, 0x06,
	0x8f, 0x43, 0x78, 0x50, 0xe8, 0x16, 0xa6, 0xda, 0x70, 0x26, 0xac, 0xef, 0x9d, 0x65, 0x20, 0x8f,
	0xcb, 0x7f, 0x03, 0x5d, 0xc2, 0xf1, 0xd6, 0xac, 0x71, 0x1c, 0x2a, 0xfe, 0xe9, 0xf9, 0x0d, 0xad,
	0x5e, 0xb7, 0x5a, 0xaf, 0x6b, 0x9c, 0x74, 0xfc, 0xd1, 0xf0, 0xbf, 0xd8, 0x2e, 0x3a, 0x0f, 0xd1,
	0x41, 0xa1, 0x02, 0x22, 0x46, 0xe5, 0xd3, 0x92, 0xca, 0x05, 0x7f, 0x39, 0xd6, 0x27, 0x19, 0xc8,
	0x27, 0xe5, 0x89, 0xc4, 0x5f, 0xa9, 0x76, 0x92, 0x1b, 0x9c, 0x76, 0x57, 0x06, 0xf1, 0xf8, 0xf0,
	0xdd, 0x10, 0xb0, 0x6f, 0x08, 0xf8, 0x6d, 0x08, 0xf8, 0x6a, 0x49, 0xb4, 0x6f, 0x49, 0xf4, 0xd3,
	0x92, 0xe8, 0x6d, 0x3a, 0xdc, 0xb8, 0x1a, 0xae, 0xec, 0x6a, 0xcd, 0xed, 0x7b, 0x12, 0x66, 0xbe,
	0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x4c, 0x3f, 0xcc, 0x90, 0x01, 0x00, 0x00,
}

func (m *Denom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Denom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Denom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintDenom(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x4a
	}
	if m.CanChangeMaxSupply {
		i--
		if m.CanChangeMaxSupply {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.Supply != 0 {
		i = encodeVarintDenom(dAtA, i, uint64(m.Supply))
		i--
		dAtA[i] = 0x38
	}
	if m.MaxSupply != 0 {
		i = encodeVarintDenom(dAtA, i, uint64(m.MaxSupply))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintDenom(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Precision != 0 {
		i = encodeVarintDenom(dAtA, i, uint64(m.Precision))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Ticker) > 0 {
		i -= len(m.Ticker)
		copy(dAtA[i:], m.Ticker)
		i = encodeVarintDenom(dAtA, i, uint64(len(m.Ticker)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintDenom(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintDenom(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDenom(dAtA []byte, offset int, v uint64) int {
	offset -= sovDenom(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Denom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovDenom(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovDenom(uint64(l))
	}
	l = len(m.Ticker)
	if l > 0 {
		n += 1 + l + sovDenom(uint64(l))
	}
	if m.Precision != 0 {
		n += 1 + sovDenom(uint64(m.Precision))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovDenom(uint64(l))
	}
	if m.MaxSupply != 0 {
		n += 1 + sovDenom(uint64(m.MaxSupply))
	}
	if m.Supply != 0 {
		n += 1 + sovDenom(uint64(m.Supply))
	}
	if m.CanChangeMaxSupply {
		n += 2
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovDenom(uint64(l))
	}
	return n
}

func sovDenom(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDenom(x uint64) (n int) {
	return sovDenom(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Denom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDenom
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
			return fmt.Errorf("proto: Denom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Denom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ticker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Precision", wireType)
			}
			m.Precision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Precision |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			m.MaxSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSupply |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
			}
			m.Supply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Supply |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanChangeMaxSupply", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanChangeMaxSupply = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDenom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDenom
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
func skipDenom(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDenom
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
					return 0, ErrIntOverflowDenom
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
					return 0, ErrIntOverflowDenom
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
				return 0, ErrInvalidLengthDenom
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDenom
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDenom
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDenom        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDenom          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDenom = fmt.Errorf("proto: unexpected end of group")
)
