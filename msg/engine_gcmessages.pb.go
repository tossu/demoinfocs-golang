// Code generated by protoc-gen-go.
// source: engine_gcmessages.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CEngineGotvSyncPacket struct {
	MatchId          *uint64  `protobuf:"varint,1,opt,name=match_id" json:"match_id,omitempty"`
	InstanceId       *uint32  `protobuf:"varint,2,opt,name=instance_id" json:"instance_id,omitempty"`
	Signupfragment   *uint32  `protobuf:"varint,3,opt,name=signupfragment" json:"signupfragment,omitempty"`
	Currentfragment  *uint32  `protobuf:"varint,4,opt,name=currentfragment" json:"currentfragment,omitempty"`
	Tickrate         *float32 `protobuf:"fixed32,5,opt,name=tickrate" json:"tickrate,omitempty"`
	Tick             *uint32  `protobuf:"varint,6,opt,name=tick" json:"tick,omitempty"`
	Rtdelay          *float32 `protobuf:"fixed32,8,opt,name=rtdelay" json:"rtdelay,omitempty"`
	Rcvage           *float32 `protobuf:"fixed32,9,opt,name=rcvage" json:"rcvage,omitempty"`
	KeyframeInterval *float32 `protobuf:"fixed32,10,opt,name=keyframe_interval" json:"keyframe_interval,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CEngineGotvSyncPacket) Reset()                    { *m = CEngineGotvSyncPacket{} }
func (m *CEngineGotvSyncPacket) String() string            { return proto.CompactTextString(m) }
func (*CEngineGotvSyncPacket) ProtoMessage()               {}
func (*CEngineGotvSyncPacket) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *CEngineGotvSyncPacket) GetMatchId() uint64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetInstanceId() uint32 {
	if m != nil && m.InstanceId != nil {
		return *m.InstanceId
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetSignupfragment() uint32 {
	if m != nil && m.Signupfragment != nil {
		return *m.Signupfragment
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetCurrentfragment() uint32 {
	if m != nil && m.Currentfragment != nil {
		return *m.Currentfragment
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetTickrate() float32 {
	if m != nil && m.Tickrate != nil {
		return *m.Tickrate
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetTick() uint32 {
	if m != nil && m.Tick != nil {
		return *m.Tick
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetRtdelay() float32 {
	if m != nil && m.Rtdelay != nil {
		return *m.Rtdelay
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetRcvage() float32 {
	if m != nil && m.Rcvage != nil {
		return *m.Rcvage
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetKeyframeInterval() float32 {
	if m != nil && m.KeyframeInterval != nil {
		return *m.KeyframeInterval
	}
	return 0
}

func init() {
	proto.RegisterType((*CEngineGotvSyncPacket)(nil), "CEngineGotvSyncPacket")
}

func init() { proto.RegisterFile("engine_gcmessages.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0xc9, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x80, 0x61, 0xdc, 0x86, 0x12, 0x0e, 0x68, 0x4b, 0x10, 0xd4, 0x6c, 0x11, 0x53, 0x26, 0x1e,
	0x02, 0x84, 0x58, 0x91, 0xd8, 0x58, 0xa2, 0x93, 0x7b, 0x35, 0x56, 0xea, 0x4b, 0x75, 0xbe, 0x44,
	0xca, 0xc6, 0xfb, 0xf1, 0x52, 0x28, 0x1e, 0x18, 0xff, 0xff, 0x83, 0x1d, 0xb1, 0x0f, 0x4c, 0xad,
	0x77, 0x91, 0x52, 0x42, 0x4f, 0xe9, 0xf9, 0x24, 0xbd, 0xf6, 0x4f, 0xbf, 0x06, 0xee, 0x5f, 0xdf,
	0x32, 0xbe, 0xf7, 0x3a, 0x7e, 0x4e, 0xec, 0x3e, 0xd0, 0x75, 0xa4, 0xd5, 0x16, 0xca, 0x88, 0xea,
	0xbe, 0xdb, 0xb0, 0xb7, 0xa6, 0x36, 0x4d, 0x51, 0xdd, 0xc1, 0x55, 0xe0, 0xa4, 0xc8, 0x8e, 0xe6,
	0xb9, 0xa8, 0x4d, 0x73, 0x53, 0x3d, 0xc0, 0x3a, 0x05, 0xcf, 0xc3, 0xe9, 0x20, 0xe8, 0x23, 0xb1,
	0xda, 0x65, 0xfe, 0x3b, 0xd8, 0xb8, 0x41, 0x84, 0x58, 0xff, 0xa1, 0xc8, 0xb0, 0x85, 0x52, 0x83,
	0xeb, 0x04, 0x95, 0xec, 0x79, 0x6d, 0x9a, 0x45, 0x75, 0x0d, 0xc5, 0x7c, 0xec, 0x2a, 0xfb, 0x06,
	0x2e, 0x44, 0xf7, 0x74, 0xc4, 0xc9, 0x96, 0x99, 0xd7, 0xb0, 0x12, 0x37, 0xa2, 0x27, 0x7b, 0x99,
	0xfb, 0x11, 0x6e, 0x3b, 0x9a, 0x0e, 0x82, 0x91, 0xda, 0xc0, 0x4a, 0x32, 0xe2, 0xd1, 0xc2, 0x4c,
	0x2f, 0xe5, 0xd7, 0x32, 0x26, 0xff, 0x63, 0xce, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x65,
	0xc2, 0x03, 0xf1, 0x00, 0x00, 0x00,
}