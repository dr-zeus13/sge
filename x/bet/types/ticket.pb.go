// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/bet/ticket.proto

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

type BetPlacementTicketPayload struct {
	// OddsUID is the selected odds id
	OddsUID string `protobuf:"bytes,1,opt,name=OddsUID,json=odds_uid,proto3" json:"odds_uid"`
	// Odds is a list of odds at the bet placement time
	Odds []*BetOdds `protobuf:"bytes,2,rep,name=odds,proto3" json:"odds,omitempty"`
	// kyc_data contains the details of user KYC
	KycData *KycDataPayload `protobuf:"bytes,3,opt,name=kyc_data,json=kycData,proto3" json:"kyc_data,omitempty"`
}

func (m *BetPlacementTicketPayload) Reset()         { *m = BetPlacementTicketPayload{} }
func (m *BetPlacementTicketPayload) String() string { return proto.CompactTextString(m) }
func (*BetPlacementTicketPayload) ProtoMessage()    {}
func (*BetPlacementTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6959e7db451613, []int{0}
}
func (m *BetPlacementTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BetPlacementTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BetPlacementTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BetPlacementTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BetPlacementTicketPayload.Merge(m, src)
}
func (m *BetPlacementTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *BetPlacementTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_BetPlacementTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_BetPlacementTicketPayload proto.InternalMessageInfo

func (m *BetPlacementTicketPayload) GetOddsUID() string {
	if m != nil {
		return m.OddsUID
	}
	return ""
}

func (m *BetPlacementTicketPayload) GetOdds() []*BetOdds {
	if m != nil {
		return m.Odds
	}
	return nil
}

func (m *BetPlacementTicketPayload) GetKycData() *KycDataPayload {
	if m != nil {
		return m.KycData
	}
	return nil
}

type KycDataPayload struct {
	KycRequired bool   `protobuf:"varint,1,opt,name=kyc_required,json=kycRequired,proto3" json:"kyc_required,omitempty"`
	KycApproved bool   `protobuf:"varint,2,opt,name=kyc_approved,json=kycApproved,proto3" json:"kyc_approved,omitempty"`
	KycId       string `protobuf:"bytes,3,opt,name=kyc_id,json=kycId,proto3" json:"kyc_id,omitempty"`
}

func (m *KycDataPayload) Reset()         { *m = KycDataPayload{} }
func (m *KycDataPayload) String() string { return proto.CompactTextString(m) }
func (*KycDataPayload) ProtoMessage()    {}
func (*KycDataPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6959e7db451613, []int{1}
}
func (m *KycDataPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KycDataPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KycDataPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KycDataPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KycDataPayload.Merge(m, src)
}
func (m *KycDataPayload) XXX_Size() int {
	return m.Size()
}
func (m *KycDataPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_KycDataPayload.DiscardUnknown(m)
}

var xxx_messageInfo_KycDataPayload proto.InternalMessageInfo

func (m *KycDataPayload) GetKycRequired() bool {
	if m != nil {
		return m.KycRequired
	}
	return false
}

func (m *KycDataPayload) GetKycApproved() bool {
	if m != nil {
		return m.KycApproved
	}
	return false
}

func (m *KycDataPayload) GetKycId() string {
	if m != nil {
		return m.KycId
	}
	return ""
}

func init() {
	proto.RegisterType((*BetPlacementTicketPayload)(nil), "sgenetwork.sge.bet.BetPlacementTicketPayload")
	proto.RegisterType((*KycDataPayload)(nil), "sgenetwork.sge.bet.KycDataPayload")
}

func init() { proto.RegisterFile("sge/bet/ticket.proto", fileDescriptor_cf6959e7db451613) }

var fileDescriptor_cf6959e7db451613 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x9b, 0x16, 0xfa, 0xe3, 0x22, 0x06, 0xab, 0x40, 0x28, 0x52, 0x5a, 0x3a, 0xa0, 0x2e,
	0x24, 0x52, 0x99, 0x91, 0x20, 0xea, 0x52, 0x31, 0x50, 0x45, 0xb0, 0xb0, 0x54, 0x8e, 0x7d, 0x65,
	0xa2, 0xb4, 0x75, 0x88, 0x5d, 0x20, 0x6f, 0xc1, 0x1b, 0xb1, 0x32, 0x76, 0x64, 0xaa, 0x50, 0xba,
	0xf1, 0x14, 0xc8, 0x49, 0x5a, 0x09, 0xc1, 0x76, 0x7c, 0xee, 0x77, 0x8f, 0x8e, 0x6d, 0xd4, 0x92,
	0x1c, 0x1c, 0x1f, 0x94, 0xa3, 0x02, 0x1a, 0x82, 0xb2, 0xa3, 0x58, 0x28, 0x81, 0xb1, 0xe4, 0x30,
	0x07, 0xf5, 0x22, 0xe2, 0xd0, 0x96, 0x1c, 0x6c, 0x1f, 0x54, 0xbb, 0xc5, 0x05, 0x17, 0xd9, 0xd8,
	0xd1, 0x2a, 0x27, 0xdb, 0x87, 0x9b, 0x7d, 0x1f, 0xd4, 0x44, 0x30, 0x26, 0x73, 0xbf, 0xf7, 0x6e,
	0xa0, 0x63, 0x17, 0xd4, 0x78, 0x4a, 0x28, 0xcc, 0x60, 0xae, 0xee, 0xb2, 0xf8, 0x31, 0x49, 0xa6,
	0x82, 0x30, 0x3c, 0x40, 0xb5, 0x5b, 0xc6, 0xe4, 0xfd, 0x68, 0x68, 0x1a, 0x5d, 0xa3, 0xdf, 0x70,
	0x8f, 0xd2, 0x55, 0x67, 0x63, 0x7d, 0xaf, 0x3a, 0x75, 0x9d, 0x34, 0x59, 0x04, 0xcc, 0xdb, 0x2a,
	0xec, 0xa0, 0x1d, 0xad, 0xcd, 0x72, 0xb7, 0xd2, 0x6f, 0x0e, 0x4e, 0xec, 0xbf, 0x15, 0x6d, 0x17,
	0x94, 0xce, 0xf0, 0x32, 0x10, 0x5f, 0xa2, 0x7a, 0x98, 0xd0, 0x09, 0x23, 0x8a, 0x98, 0x95, 0xae,
	0xd1, 0x6f, 0x0e, 0x7a, 0xff, 0x2d, 0xdd, 0x24, 0x74, 0x48, 0x14, 0x29, 0xaa, 0x79, 0xb5, 0x30,
	0x3f, 0xf7, 0x66, 0x68, 0xff, 0xf7, 0x08, 0x9f, 0xa2, 0x3d, 0x1d, 0x18, 0xc3, 0xd3, 0x22, 0x88,
	0x81, 0x65, 0xd5, 0xeb, 0x5e, 0x33, 0x4c, 0xa8, 0x57, 0x58, 0x1b, 0x84, 0x44, 0x51, 0x2c, 0x9e,
	0x81, 0x99, 0xe5, 0x2d, 0x72, 0x5d, 0x58, 0xf8, 0x00, 0x55, 0x35, 0x12, 0xb0, 0xac, 0x54, 0xc3,
	0xdb, 0x0d, 0x13, 0x3a, 0x62, 0xee, 0xd5, 0x47, 0x6a, 0x19, 0xcb, 0xd4, 0x32, 0xbe, 0x52, 0xcb,
	0x78, 0x5b, 0x5b, 0xa5, 0xe5, 0xda, 0x2a, 0x7d, 0xae, 0xad, 0xd2, 0xc3, 0x19, 0x0f, 0xd4, 0xe3,
	0xc2, 0xb7, 0xa9, 0x98, 0x39, 0x92, 0xc3, 0x79, 0x71, 0x01, 0xad, 0x9d, 0xd7, 0xfc, 0xef, 0x92,
	0x08, 0xa4, 0x5f, 0xcd, 0x5e, 0xfe, 0xe2, 0x27, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xfa, 0xe5, 0x70,
	0xd3, 0x01, 0x00, 0x00,
}

func (m *BetPlacementTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BetPlacementTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BetPlacementTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.KycData != nil {
		{
			size, err := m.KycData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTicket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Odds) > 0 {
		for iNdEx := len(m.Odds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Odds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTicket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.OddsUID) > 0 {
		i -= len(m.OddsUID)
		copy(dAtA[i:], m.OddsUID)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.OddsUID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KycDataPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KycDataPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KycDataPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.KycId) > 0 {
		i -= len(m.KycId)
		copy(dAtA[i:], m.KycId)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.KycId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.KycApproved {
		i--
		if m.KycApproved {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.KycRequired {
		i--
		if m.KycRequired {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTicket(dAtA []byte, offset int, v uint64) int {
	offset -= sovTicket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BetPlacementTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OddsUID)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	if len(m.Odds) > 0 {
		for _, e := range m.Odds {
			l = e.Size()
			n += 1 + l + sovTicket(uint64(l))
		}
	}
	if m.KycData != nil {
		l = m.KycData.Size()
		n += 1 + l + sovTicket(uint64(l))
	}
	return n
}

func (m *KycDataPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.KycRequired {
		n += 2
	}
	if m.KycApproved {
		n += 2
	}
	l = len(m.KycId)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	return n
}

func sovTicket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTicket(x uint64) (n int) {
	return sovTicket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BetPlacementTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: BetPlacementTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BetPlacementTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OddsUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OddsUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Odds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Odds = append(m.Odds, &BetOdds{})
			if err := m.Odds[len(m.Odds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KycData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.KycData == nil {
				m.KycData = &KycDataPayload{}
			}
			if err := m.KycData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func (m *KycDataPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: KycDataPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KycDataPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KycRequired", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
			m.KycRequired = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KycApproved", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
			m.KycApproved = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KycId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KycId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func skipTicket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
				return 0, ErrInvalidLengthTicket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTicket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTicket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTicket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTicket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTicket = fmt.Errorf("proto: unexpected end of group")
)
