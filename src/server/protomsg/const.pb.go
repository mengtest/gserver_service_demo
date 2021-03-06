// Code generated by protoc-gen-go.
// source: const.proto
// DO NOT EDIT!

/*
Package protomsg is a generated protocol buffer package.

It is generated from these files:
	const.proto
	fight.proto
	game.proto
	gen.proto
	login.proto
	match.proto

It has these top-level messages:
	Vector2
	AOI
	AOIPushItem
	SpawnState
	CreateFightReq
	CreateFightAck
	JoinFightReq
	JoinFightAck
	JoinFightPush
	MoveReq
	MovePush
	AttackPush
	SpwanPush
	DeadSettlement
	DeadPush
	EndSettlement
	OverPush
	TestReq
	TestAck
	TestPush
	HeartbeatReq
	HeartbeatAck
	ErrorAck
	TouristReq
	TouristAck
	TouristLoginReq
	TouristLoginAck
	MatchReq
	MatchPush
*/
package protomsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TYPE int32

const (
	// 1v1战斗
	TYPE_F1V1 TYPE = 1
	// 3v3战斗
	TYPE_F3V3 TYPE = 3
	// 5v5战斗
	TYPE_F5V5 TYPE = 5
)

var TYPE_name = map[int32]string{
	1: "F1V1",
	3: "F3V3",
	5: "F5V5",
}
var TYPE_value = map[string]int32{
	"F1V1": 1,
	"F3V3": 3,
	"F5V5": 5,
}

func (x TYPE) Enum() *TYPE {
	p := new(TYPE)
	*p = x
	return p
}
func (x TYPE) String() string {
	return proto.EnumName(TYPE_name, int32(x))
}
func (x *TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TYPE_value, data, "TYPE")
	if err != nil {
		return err
	}
	*x = TYPE(value)
	return nil
}
func (TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ENTITY int32

const (
	// 中立方
	ENTITY_NEUTRAL_FLAG ENTITY = 0
	// 蓝色方战船
	ENTITY_BLUE_SHIP_FLAG ENTITY = 1
	// 红色方战船
	ENTITY_RED_SHIP_FLAG ENTITY = 2
	// 蓝色方小兵
	ENTITY_BLUE_SOLDIER_FLAG ENTITY = 3
	// 红色方小兵
	ENTITY_RED_SOLDIER_FLAG ENTITY = 4
	// 蓝色方防御塔
	ENTITY_BLUE_TOWER_FLAG ENTITY = 5
	// 红色方防御塔
	ENTITY_RED_TOWER_FLAG ENTITY = 6
	// 蓝色方水晶
	ENTITY_BLUE_CRYSTAL ENTITY = 7
	// 红色方水晶
	ENTITY_RED_CRYSTAL ENTITY = 8
)

var ENTITY_name = map[int32]string{
	0: "NEUTRAL_FLAG",
	1: "BLUE_SHIP_FLAG",
	2: "RED_SHIP_FLAG",
	3: "BLUE_SOLDIER_FLAG",
	4: "RED_SOLDIER_FLAG",
	5: "BLUE_TOWER_FLAG",
	6: "RED_TOWER_FLAG",
	7: "BLUE_CRYSTAL",
	8: "RED_CRYSTAL",
}
var ENTITY_value = map[string]int32{
	"NEUTRAL_FLAG":      0,
	"BLUE_SHIP_FLAG":    1,
	"RED_SHIP_FLAG":     2,
	"BLUE_SOLDIER_FLAG": 3,
	"RED_SOLDIER_FLAG":  4,
	"BLUE_TOWER_FLAG":   5,
	"RED_TOWER_FLAG":    6,
	"BLUE_CRYSTAL":      7,
	"RED_CRYSTAL":       8,
}

func (x ENTITY) Enum() *ENTITY {
	p := new(ENTITY)
	*p = x
	return p
}
func (x ENTITY) String() string {
	return proto.EnumName(ENTITY_name, int32(x))
}
func (x *ENTITY) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ENTITY_value, data, "ENTITY")
	if err != nil {
		return err
	}
	*x = ENTITY(value)
	return nil
}
func (ENTITY) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GameConst int32

const (
	// 3v3战斗创建失败
	GameConst_CREATE3V3FAILED GameConst = 1000
	// 3v3战斗id错误
	GameConst_FIGHTID3V3ERROR GameConst = 1001
	// 3v3战斗进程不存在
	GameConst_FIGHT3V3ERROR GameConst = 1002
	// 加入3v3战斗失败
	GameConst_JOIN3V3ERROR GameConst = 1003
	// 获取3v3战斗双方失败
	GameConst_GET3V3ERROR GameConst = 1004
	// 实体id错误
	GameConst_ENTITYIDERROR GameConst = 1005
)

var GameConst_name = map[int32]string{
	1000: "CREATE3V3FAILED",
	1001: "FIGHTID3V3ERROR",
	1002: "FIGHT3V3ERROR",
	1003: "JOIN3V3ERROR",
	1004: "GET3V3ERROR",
	1005: "ENTITYIDERROR",
}
var GameConst_value = map[string]int32{
	"CREATE3V3FAILED": 1000,
	"FIGHTID3V3ERROR": 1001,
	"FIGHT3V3ERROR":   1002,
	"JOIN3V3ERROR":    1003,
	"GET3V3ERROR":     1004,
	"ENTITYIDERROR":   1005,
}

func (x GameConst) Enum() *GameConst {
	p := new(GameConst)
	*p = x
	return p
}
func (x GameConst) String() string {
	return proto.EnumName(GameConst_name, int32(x))
}
func (x *GameConst) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GameConst_value, data, "GameConst")
	if err != nil {
		return err
	}
	*x = GameConst(value)
	return nil
}
func (GameConst) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SoldierStatus int32

const (
	// 移动
	SoldierStatus_MOVE SoldierStatus = 1
	// 停止
	SoldierStatus_STOP SoldierStatus = 2
	// 攻击
	SoldierStatus_ATTACK SoldierStatus = 3
)

var SoldierStatus_name = map[int32]string{
	1: "MOVE",
	2: "STOP",
	3: "ATTACK",
}
var SoldierStatus_value = map[string]int32{
	"MOVE":   1,
	"STOP":   2,
	"ATTACK": 3,
}

func (x SoldierStatus) Enum() *SoldierStatus {
	p := new(SoldierStatus)
	*p = x
	return p
}
func (x SoldierStatus) String() string {
	return proto.EnumName(SoldierStatus_name, int32(x))
}
func (x *SoldierStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SoldierStatus_value, data, "SoldierStatus")
	if err != nil {
		return err
	}
	*x = SoldierStatus(value)
	return nil
}
func (SoldierStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterEnum("protomsg.TYPE", TYPE_name, TYPE_value)
	proto.RegisterEnum("protomsg.ENTITY", ENTITY_name, ENTITY_value)
	proto.RegisterEnum("protomsg.GameConst", GameConst_name, GameConst_value)
	proto.RegisterEnum("protomsg.SoldierStatus", SoldierStatus_name, SoldierStatus_value)
}

var fileDescriptor0 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0x5d, 0x4e, 0xf2, 0x40,
	0x14, 0x86, 0xbf, 0x7e, 0xfc, 0x14, 0x0f, 0x20, 0xc3, 0xa8, 0x3b, 0xf0, 0xaa, 0x17, 0x1a, 0x42,
	0x58, 0xc0, 0xd8, 0x9e, 0x96, 0xd1, 0xca, 0x90, 0xe9, 0x50, 0xc3, 0x15, 0x21, 0x4a, 0x8c, 0x89,
	0x58, 0x03, 0x75, 0x07, 0xae, 0xca, 0x15, 0xf9, 0xbb, 0x07, 0xe7, 0x27, 0x60, 0xaf, 0xe6, 0xe4,
	0xc9, 0x33, 0x79, 0xcf, 0x79, 0xa1, 0x7d, 0x5b, 0x3c, 0x6d, 0xcb, 0xb3, 0xe7, 0x4d, 0x51, 0x16,
	0xb4, 0x65, 0x9f, 0xf5, 0xf6, 0x3e, 0x38, 0x85, 0xba, 0x9a, 0x4f, 0x91, 0xb6, 0xa0, 0x1e, 0x0f,
	0xf2, 0x01, 0xf1, 0xec, 0x34, 0xcc, 0x87, 0xa4, 0x66, 0xa7, 0x51, 0x3e, 0x22, 0x8d, 0xe0, 0xcd,
	0x83, 0x26, 0x4e, 0x14, 0x57, 0x73, 0x4a, 0xa0, 0x33, 0xc1, 0x99, 0x92, 0x2c, 0x5d, 0xc4, 0x29,
	0x4b, 0xc8, 0x3f, 0x4a, 0xe1, 0xf0, 0x22, 0x9d, 0xe1, 0x22, 0x1b, 0xf3, 0xa9, 0x63, 0x1e, 0xed,
	0x43, 0x57, 0x62, 0x54, 0x41, 0xff, 0xe9, 0x09, 0xf4, 0x9d, 0x26, 0xd2, 0x88, 0xa3, 0x74, 0xb8,
	0x46, 0x8f, 0x81, 0x58, 0xb3, 0x4a, 0xeb, 0xf4, 0x08, 0x7a, 0x56, 0x56, 0xe2, 0x66, 0x07, 0x1b,
	0x26, 0xc8, 0xa8, 0x15, 0xd6, 0x34, 0xeb, 0x58, 0x31, 0x94, 0xf3, 0x4c, 0xb1, 0x94, 0xf8, 0xb4,
	0x07, 0x6d, 0x63, 0xed, 0x40, 0x2b, 0x78, 0xf5, 0xe0, 0x20, 0x59, 0xae, 0x57, 0xa1, 0x29, 0x40,
	0xe7, 0xf5, 0x42, 0x89, 0x4c, 0xa1, 0xbe, 0x31, 0x66, 0x3c, 0xc5, 0x88, 0xbc, 0xfb, 0x86, 0xc6,
	0x3c, 0x19, 0x2b, 0x1e, 0x69, 0x8c, 0x52, 0x0a, 0x49, 0x3e, 0x7c, 0x1d, 0xd8, 0xb5, 0x74, 0xcf,
	0x3e, 0x7d, 0x7d, 0x59, 0xe7, 0x52, 0xf0, 0xc9, 0x1e, 0x7d, 0xf9, 0x7a, 0x87, 0x76, 0x82, 0x7f,
	0xd2, 0xb7, 0xfd, 0xe8, 0xea, 0xe2, 0x91, 0x63, 0x3f, 0x7e, 0x70, 0x0e, 0xdd, 0xac, 0x78, 0xbc,
	0x7b, 0x58, 0x6d, 0xb2, 0x72, 0x59, 0xbe, 0x6c, 0x4d, 0xbd, 0xd7, 0x22, 0x47, 0x57, 0x79, 0xa6,
	0xc4, 0x54, 0x97, 0x04, 0xd0, 0x64, 0x4a, 0xb1, 0xf0, 0x8a, 0xd4, 0x7e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x28, 0x8f, 0x91, 0xfb, 0xb2, 0x01, 0x00, 0x00,
}
