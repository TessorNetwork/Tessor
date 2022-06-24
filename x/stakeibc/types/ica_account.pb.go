// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stakeibc/ica_account.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type ICAAccountType int32

const (
	ICAAccountType_DELEGATION ICAAccountType = 0
	ICAAccountType_FEE        ICAAccountType = 1
	ICAAccountType_WITHDRAWAL ICAAccountType = 2
)

var ICAAccountType_name = map[int32]string{
	0: "DELEGATION",
	1: "FEE",
	2: "WITHDRAWAL",
}

var ICAAccountType_value = map[string]int32{
	"DELEGATION": 0,
	"FEE":        1,
	"WITHDRAWAL": 2,
}

func (x ICAAccountType) String() string {
	return proto.EnumName(ICAAccountType_name, int32(x))
}

func (ICAAccountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f7243c23ee376c2f, []int{0}
}

// TODO(TEST-XX): Update these fields to be more useful (e.g. balances should be coins, maybe store port name directly)
type ICAAccount struct {
	Address           string         `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Delegations       []*Delegation  `protobuf:"bytes,4,rep,name=delegations,proto3" json:"delegations,omitempty"`
	Target            ICAAccountType `protobuf:"varint,5,opt,name=target,proto3,enum=Stridelabs.stride.stakeibc.ICAAccountType" json:"target,omitempty"`
	WithdrawalBalance int64          `protobuf:"varint,8,opt,name=withdrawalBalance,proto3" json:"withdrawalBalance,omitempty"`
}

func (m *ICAAccount) Reset()         { *m = ICAAccount{} }
func (m *ICAAccount) String() string { return proto.CompactTextString(m) }
func (*ICAAccount) ProtoMessage()    {}
func (*ICAAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7243c23ee376c2f, []int{0}
}
func (m *ICAAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ICAAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ICAAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ICAAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ICAAccount.Merge(m, src)
}
func (m *ICAAccount) XXX_Size() int {
	return m.Size()
}
func (m *ICAAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ICAAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ICAAccount proto.InternalMessageInfo

func (m *ICAAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ICAAccount) GetDelegations() []*Delegation {
	if m != nil {
		return m.Delegations
	}
	return nil
}

func (m *ICAAccount) GetTarget() ICAAccountType {
	if m != nil {
		return m.Target
	}
	return ICAAccountType_DELEGATION
}

func (m *ICAAccount) GetWithdrawalBalance() int64 {
	if m != nil {
		return m.WithdrawalBalance
	}
	return 0
}

func init() {
	proto.RegisterEnum("Stridelabs.stride.stakeibc.ICAAccountType", ICAAccountType_name, ICAAccountType_value)
	proto.RegisterType((*ICAAccount)(nil), "Stridelabs.stride.stakeibc.ICAAccount")
}

func init() { proto.RegisterFile("stakeibc/ica_account.proto", fileDescriptor_f7243c23ee376c2f) }

var fileDescriptor_f7243c23ee376c2f = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x2e, 0x49, 0xcc,
	0x4e, 0xcd, 0x4c, 0x4a, 0xd6, 0xcf, 0x4c, 0x4e, 0x8c, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x0a, 0x2e, 0x29, 0xca, 0x4c, 0x49, 0xcd, 0x49,
	0x4c, 0x2a, 0xd6, 0x2b, 0x06, 0x33, 0xf5, 0x60, 0xaa, 0xa5, 0x24, 0xe1, 0xfa, 0x52, 0x52, 0x73,
	0x52, 0xd3, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0x20, 0xda, 0xa4, 0x24, 0x93, 0xf3, 0x8b, 0x73, 0xf3,
	0x8b, 0xe3, 0xc1, 0x3c, 0x7d, 0x08, 0x07, 0x22, 0xa5, 0xf4, 0x87, 0x91, 0x8b, 0xcb, 0xd3, 0xd9,
	0xd1, 0x11, 0x62, 0x8d, 0x90, 0x11, 0x17, 0x7b, 0x62, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0xa7, 0x93, 0xc4, 0xa5, 0x2d, 0xba, 0x22, 0x50, 0x1d, 0x8e, 0x10, 0x19,
	0x90, 0x1b, 0xf2, 0xd2, 0x83, 0x60, 0x0a, 0x85, 0x3c, 0xb8, 0xb8, 0x11, 0x36, 0x16, 0x4b, 0xb0,
	0x28, 0x30, 0x6b, 0x70, 0x1b, 0xa9, 0xe9, 0xe1, 0x76, 0xaa, 0x9e, 0x0b, 0x5c, 0x79, 0x10, 0xb2,
	0x56, 0x21, 0x27, 0x2e, 0xb6, 0x92, 0xc4, 0xa2, 0xf4, 0xd4, 0x12, 0x09, 0x56, 0x05, 0x46, 0x0d,
	0x3e, 0x23, 0x2d, 0x7c, 0x86, 0x20, 0x5c, 0x1d, 0x52, 0x59, 0x90, 0x1a, 0x04, 0xd5, 0x29, 0xa4,
	0xc3, 0x25, 0x58, 0x9e, 0x59, 0x92, 0x91, 0x52, 0x94, 0x58, 0x9e, 0x98, 0xe3, 0x94, 0x98, 0x93,
	0x98, 0x97, 0x9c, 0x2a, 0xc1, 0xa1, 0xc0, 0xa8, 0xc1, 0x1c, 0x84, 0x29, 0xa1, 0x65, 0xc9, 0xc5,
	0x87, 0x6a, 0x8e, 0x10, 0x1f, 0x17, 0x97, 0x8b, 0xab, 0x8f, 0xab, 0xbb, 0x63, 0x88, 0xa7, 0xbf,
	0x9f, 0x00, 0x83, 0x10, 0x3b, 0x17, 0xb3, 0x9b, 0xab, 0xab, 0x00, 0x23, 0x48, 0x22, 0xdc, 0x33,
	0xc4, 0xc3, 0x25, 0xc8, 0x31, 0xdc, 0xd1, 0x47, 0x80, 0xc9, 0xc9, 0xe3, 0xc4, 0x23, 0x39, 0xc6,
	0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39,
	0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xf4, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0x21, 0x1e, 0xd0, 0xf5, 0x49, 0x4c, 0x2a, 0xd6, 0x87, 0xf8, 0x40, 0xbf, 0x42, 0x1f,
	0x1e, 0x53, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0xa8, 0x30, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x3c, 0x46, 0x9b, 0xde, 0xfa, 0x01, 0x00, 0x00,
}

func (m *ICAAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ICAAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ICAAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WithdrawalBalance != 0 {
		i = encodeVarintIcaAccount(dAtA, i, uint64(m.WithdrawalBalance))
		i--
		dAtA[i] = 0x40
	}
	if m.Target != 0 {
		i = encodeVarintIcaAccount(dAtA, i, uint64(m.Target))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Delegations) > 0 {
		for iNdEx := len(m.Delegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Delegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIcaAccount(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintIcaAccount(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIcaAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovIcaAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ICAAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovIcaAccount(uint64(l))
	}
	if len(m.Delegations) > 0 {
		for _, e := range m.Delegations {
			l = e.Size()
			n += 1 + l + sovIcaAccount(uint64(l))
		}
	}
	if m.Target != 0 {
		n += 1 + sovIcaAccount(uint64(m.Target))
	}
	if m.WithdrawalBalance != 0 {
		n += 1 + sovIcaAccount(uint64(m.WithdrawalBalance))
	}
	return n
}

func sovIcaAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIcaAccount(x uint64) (n int) {
	return sovIcaAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ICAAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIcaAccount
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
			return fmt.Errorf("proto: ICAAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ICAAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaAccount
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
				return ErrInvalidLengthIcaAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaAccount
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
				return ErrInvalidLengthIcaAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIcaAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegations = append(m.Delegations, &Delegation{})
			if err := m.Delegations[len(m.Delegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			m.Target = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Target |= ICAAccountType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalBalance", wireType)
			}
			m.WithdrawalBalance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WithdrawalBalance |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIcaAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIcaAccount
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
func skipIcaAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIcaAccount
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
					return 0, ErrIntOverflowIcaAccount
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
					return 0, ErrIntOverflowIcaAccount
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
				return 0, ErrInvalidLengthIcaAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIcaAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIcaAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIcaAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIcaAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIcaAccount = fmt.Errorf("proto: unexpected end of group")
)
