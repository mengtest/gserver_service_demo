// Code generated by protoc-gen-go.
// source: match.proto
// DO NOT EDIT!

package protomsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 开始匹配请求1500
type MatchReq struct {
	Fighttype        *int32 `protobuf:"varint,1,req,name=fighttype" json:"fighttype,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MatchReq) Reset()                    { *m = MatchReq{} }
func (m *MatchReq) String() string            { return proto.CompactTextString(m) }
func (*MatchReq) ProtoMessage()               {}
func (*MatchReq) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *MatchReq) GetFighttype() int32 {
	if m != nil && m.Fighttype != nil {
		return *m.Fighttype
	}
	return 0
}

// 匹配成功推送1501
type MatchPush struct {
	Fighttype        *int32  `protobuf:"varint,1,req,name=fighttype" json:"fighttype,omitempty"`
	Fightid          *string `protobuf:"bytes,2,req,name=fightid" json:"fightid,omitempty"`
	Playertype       *int32  `protobuf:"varint,3,req,name=playertype" json:"playertype,omitempty"`
	Reds             []int32 `protobuf:"varint,4,rep,name=reds" json:"reds,omitempty"`
	Blues            []int32 `protobuf:"varint,5,rep,name=blues" json:"blues,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MatchPush) Reset()                    { *m = MatchPush{} }
func (m *MatchPush) String() string            { return proto.CompactTextString(m) }
func (*MatchPush) ProtoMessage()               {}
func (*MatchPush) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *MatchPush) GetFighttype() int32 {
	if m != nil && m.Fighttype != nil {
		return *m.Fighttype
	}
	return 0
}

func (m *MatchPush) GetFightid() string {
	if m != nil && m.Fightid != nil {
		return *m.Fightid
	}
	return ""
}

func (m *MatchPush) GetPlayertype() int32 {
	if m != nil && m.Playertype != nil {
		return *m.Playertype
	}
	return 0
}

func (m *MatchPush) GetReds() []int32 {
	if m != nil {
		return m.Reds
	}
	return nil
}

func (m *MatchPush) GetBlues() []int32 {
	if m != nil {
		return m.Blues
	}
	return nil
}

func init() {
	proto.RegisterType((*MatchReq)(nil), "protomsg.match_req")
	proto.RegisterType((*MatchPush)(nil), "protomsg.match_push")
}

var fileDescriptor5 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x4d, 0x2c, 0x49,
	0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xb9, 0xc5, 0xe9, 0x4a, 0x9a,
	0x5c, 0x9c, 0x60, 0x89, 0xf8, 0xa2, 0xd4, 0x42, 0x21, 0x19, 0x2e, 0xce, 0xb4, 0xcc, 0xf4, 0x8c,
	0x92, 0x92, 0xca, 0x82, 0x54, 0x09, 0x46, 0x05, 0x26, 0x0d, 0xd6, 0x20, 0x84, 0x80, 0x52, 0x1f,
	0x23, 0x17, 0x17, 0x44, 0x6d, 0x41, 0x69, 0x71, 0x06, 0x7e, 0xc5, 0x42, 0x12, 0x5c, 0xec, 0x60,
	0x4e, 0x66, 0x8a, 0x04, 0x13, 0x50, 0x8e, 0x33, 0x08, 0xc6, 0x15, 0x92, 0xe3, 0xe2, 0x2a, 0xc8,
	0x49, 0xac, 0x4c, 0x2d, 0x02, 0x6b, 0x64, 0x06, 0x6b, 0x44, 0x12, 0x11, 0x12, 0xe2, 0x62, 0x29,
	0x4a, 0x4d, 0x29, 0x96, 0x60, 0x51, 0x60, 0x06, 0xca, 0x80, 0xd9, 0x42, 0x22, 0x5c, 0xac, 0x49,
	0x39, 0xa5, 0xa9, 0xc5, 0x12, 0xac, 0x60, 0x41, 0x08, 0x07, 0x10, 0x00, 0x00, 0xff, 0xff, 0x24,
	0xbb, 0xb7, 0xc8, 0xd3, 0x00, 0x00, 0x00,
}
