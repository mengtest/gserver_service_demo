// Code generated by protoc-gen-go.
// source: login.proto
// DO NOT EDIT!

package protomsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 创建游客账号请求500
type TouristReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *TouristReq) Reset()                    { *m = TouristReq{} }
func (m *TouristReq) String() string            { return proto.CompactTextString(m) }
func (*TouristReq) ProtoMessage()               {}
func (*TouristReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

// 创建游客账号请求返回501
type TouristAck struct {
	Id               *int32  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Account          *string `protobuf:"bytes,2,req,name=account" json:"account,omitempty"`
	Password         *string `protobuf:"bytes,3,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TouristAck) Reset()                    { *m = TouristAck{} }
func (m *TouristAck) String() string            { return proto.CompactTextString(m) }
func (*TouristAck) ProtoMessage()               {}
func (*TouristAck) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *TouristAck) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TouristAck) GetAccount() string {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return ""
}

func (m *TouristAck) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

// 游客账号登录502
type TouristLoginReq struct {
	Id               *int32  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Account          *string `protobuf:"bytes,2,req,name=account" json:"account,omitempty"`
	Password         *string `protobuf:"bytes,3,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TouristLoginReq) Reset()                    { *m = TouristLoginReq{} }
func (m *TouristLoginReq) String() string            { return proto.CompactTextString(m) }
func (*TouristLoginReq) ProtoMessage()               {}
func (*TouristLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *TouristLoginReq) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TouristLoginReq) GetAccount() string {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return ""
}

func (m *TouristLoginReq) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

// 游客账号登录返回503
type TouristLoginAck struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *TouristLoginAck) Reset()                    { *m = TouristLoginAck{} }
func (m *TouristLoginAck) String() string            { return proto.CompactTextString(m) }
func (*TouristLoginAck) ProtoMessage()               {}
func (*TouristLoginAck) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func init() {
	proto.RegisterType((*TouristReq)(nil), "protomsg.tourist_req")
	proto.RegisterType((*TouristAck)(nil), "protomsg.tourist_ack")
	proto.RegisterType((*TouristLoginReq)(nil), "protomsg.tourist_login_req")
	proto.RegisterType((*TouristLoginAck)(nil), "protomsg.tourist_login_ack")
}

var fileDescriptor4 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xc9, 0x4f, 0xcf,
	0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xb9, 0xc5, 0xe9, 0x4a, 0xbc,
	0x5c, 0xdc, 0x25, 0xf9, 0xa5, 0x45, 0x99, 0xc5, 0x25, 0xf1, 0x45, 0xa9, 0x85, 0x4a, 0xc1, 0x08,
	0x6e, 0x62, 0x72, 0xb6, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0x93, 0x06, 0x6b,
	0x10, 0x90, 0x25, 0x24, 0xc1, 0xc5, 0x9e, 0x98, 0x9c, 0x9c, 0x5f, 0x9a, 0x57, 0x22, 0xc1, 0x04,
	0x14, 0xe4, 0x0c, 0x82, 0x71, 0x85, 0xa4, 0xb8, 0x38, 0x0a, 0x12, 0x8b, 0x8b, 0xcb, 0xf3, 0x8b,
	0x52, 0x24, 0x98, 0xc1, 0x52, 0x70, 0xbe, 0x52, 0x24, 0x97, 0x20, 0xcc, 0x50, 0xb0, 0x23, 0x40,
	0x36, 0x51, 0xc9, 0x68, 0x61, 0x74, 0xa3, 0x81, 0xae, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbd,
	0x1f, 0x07, 0x43, 0xeb, 0x00, 0x00, 0x00,
}