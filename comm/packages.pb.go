// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comm/packages.proto

/*
Package comm is a generated protocol buffer package.

It is generated from these files:
	comm/packages.proto

It has these top-level messages:
	Envelope
	TimeSyncRequest
	TimeSyncResponse
	QueueChunkRequest
	PingMessage
	PongMessage
	SetVolumeRequest
	SubscribeChannelRequest
	NewSongInfo
	ChunkInfo
	PauseInfo
*/
package comm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Channel int32

const (
	Channel_AUDIO  Channel = 0
	Channel_META   Channel = 1
	Channel_LYRICS Channel = 2
)

var Channel_name = map[int32]string{
	0: "AUDIO",
	1: "META",
	2: "LYRICS",
}
var Channel_value = map[string]int32{
	"AUDIO":  0,
	"META":   1,
	"LYRICS": 2,
}

func (x Channel) String() string {
	return proto.EnumName(Channel_name, int32(x))
}
func (Channel) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Envelope struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Envelope) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Envelope) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type TimeSyncRequest struct {
	ClientSend int64 `protobuf:"varint,1,opt,name=clientSend" json:"clientSend,omitempty"`
}

func (m *TimeSyncRequest) Reset()                    { *m = TimeSyncRequest{} }
func (m *TimeSyncRequest) String() string            { return proto.CompactTextString(m) }
func (*TimeSyncRequest) ProtoMessage()               {}
func (*TimeSyncRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TimeSyncRequest) GetClientSend() int64 {
	if m != nil {
		return m.ClientSend
	}
	return 0
}

type TimeSyncResponse struct {
	ClientSendTime int64 `protobuf:"varint,1,opt,name=clientSendTime" json:"clientSendTime,omitempty"`
	ServerRecvTime int64 `protobuf:"varint,2,opt,name=serverRecvTime" json:"serverRecvTime,omitempty"`
	ServerSendTime int64 `protobuf:"varint,3,opt,name=serverSendTime" json:"serverSendTime,omitempty"`
}

func (m *TimeSyncResponse) Reset()                    { *m = TimeSyncResponse{} }
func (m *TimeSyncResponse) String() string            { return proto.CompactTextString(m) }
func (*TimeSyncResponse) ProtoMessage()               {}
func (*TimeSyncResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TimeSyncResponse) GetClientSendTime() int64 {
	if m != nil {
		return m.ClientSendTime
	}
	return 0
}

func (m *TimeSyncResponse) GetServerRecvTime() int64 {
	if m != nil {
		return m.ServerRecvTime
	}
	return 0
}

func (m *TimeSyncResponse) GetServerSendTime() int64 {
	if m != nil {
		return m.ServerSendTime
	}
	return 0
}

type QueueChunkRequest struct {
	StartTime        int64     `protobuf:"varint,1,opt,name=startTime" json:"startTime,omitempty"`
	ChunkId          int64     `protobuf:"varint,2,opt,name=chunkId" json:"chunkId,omitempty"`
	SampleLow        []float64 `protobuf:"fixed64,3,rep,packed,name=sampleLow" json:"sampleLow,omitempty"`
	SampleHigh       []float64 `protobuf:"fixed64,4,rep,packed,name=sampleHigh" json:"sampleHigh,omitempty"`
	FirstSampleIndex uint64    `protobuf:"varint,5,opt,name=firstSampleIndex" json:"firstSampleIndex,omitempty"`
}

func (m *QueueChunkRequest) Reset()                    { *m = QueueChunkRequest{} }
func (m *QueueChunkRequest) String() string            { return proto.CompactTextString(m) }
func (*QueueChunkRequest) ProtoMessage()               {}
func (*QueueChunkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QueueChunkRequest) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *QueueChunkRequest) GetChunkId() int64 {
	if m != nil {
		return m.ChunkId
	}
	return 0
}

func (m *QueueChunkRequest) GetSampleLow() []float64 {
	if m != nil {
		return m.SampleLow
	}
	return nil
}

func (m *QueueChunkRequest) GetSampleHigh() []float64 {
	if m != nil {
		return m.SampleHigh
	}
	return nil
}

func (m *QueueChunkRequest) GetFirstSampleIndex() uint64 {
	if m != nil {
		return m.FirstSampleIndex
	}
	return 0
}

type PingMessage struct {
}

func (m *PingMessage) Reset()                    { *m = PingMessage{} }
func (m *PingMessage) String() string            { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()               {}
func (*PingMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type PongMessage struct {
}

func (m *PongMessage) Reset()                    { *m = PongMessage{} }
func (m *PongMessage) String() string            { return proto.CompactTextString(m) }
func (*PongMessage) ProtoMessage()               {}
func (*PongMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type SetVolumeRequest struct {
	Volume float64 `protobuf:"fixed64,1,opt,name=volume" json:"volume,omitempty"`
}

func (m *SetVolumeRequest) Reset()                    { *m = SetVolumeRequest{} }
func (m *SetVolumeRequest) String() string            { return proto.CompactTextString(m) }
func (*SetVolumeRequest) ProtoMessage()               {}
func (*SetVolumeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SetVolumeRequest) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type SubscribeChannelRequest struct {
	Channel Channel `protobuf:"varint,1,opt,name=channel,enum=comm.Channel" json:"channel,omitempty"`
}

func (m *SubscribeChannelRequest) Reset()                    { *m = SubscribeChannelRequest{} }
func (m *SubscribeChannelRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeChannelRequest) ProtoMessage()               {}
func (*SubscribeChannelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SubscribeChannelRequest) GetChannel() Channel {
	if m != nil {
		return m.Channel
	}
	return Channel_AUDIO
}

type NewSongInfo struct {
	FirstSampleOfSongIndex uint64 `protobuf:"varint,1,opt,name=firstSampleOfSongIndex" json:"firstSampleOfSongIndex,omitempty"`
	SongFileName           string `protobuf:"bytes,2,opt,name=songFileName" json:"songFileName,omitempty"`
	SongLength             int64  `protobuf:"varint,3,opt,name=songLength" json:"songLength,omitempty"`
}

func (m *NewSongInfo) Reset()                    { *m = NewSongInfo{} }
func (m *NewSongInfo) String() string            { return proto.CompactTextString(m) }
func (*NewSongInfo) ProtoMessage()               {}
func (*NewSongInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *NewSongInfo) GetFirstSampleOfSongIndex() uint64 {
	if m != nil {
		return m.FirstSampleOfSongIndex
	}
	return 0
}

func (m *NewSongInfo) GetSongFileName() string {
	if m != nil {
		return m.SongFileName
	}
	return ""
}

func (m *NewSongInfo) GetSongLength() int64 {
	if m != nil {
		return m.SongLength
	}
	return 0
}

type ChunkInfo struct {
	StartTime        int64  `protobuf:"varint,1,opt,name=startTime" json:"startTime,omitempty"`
	FirstSampleIndex uint64 `protobuf:"varint,2,opt,name=firstSampleIndex" json:"firstSampleIndex,omitempty"`
	ChunkSize        uint64 `protobuf:"varint,3,opt,name=chunkSize" json:"chunkSize,omitempty"`
}

func (m *ChunkInfo) Reset()                    { *m = ChunkInfo{} }
func (m *ChunkInfo) String() string            { return proto.CompactTextString(m) }
func (*ChunkInfo) ProtoMessage()               {}
func (*ChunkInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ChunkInfo) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *ChunkInfo) GetFirstSampleIndex() uint64 {
	if m != nil {
		return m.FirstSampleIndex
	}
	return 0
}

func (m *ChunkInfo) GetChunkSize() uint64 {
	if m != nil {
		return m.ChunkSize
	}
	return 0
}

type PauseInfo struct {
	Playing           bool   `protobuf:"varint,1,opt,name=playing" json:"playing,omitempty"`
	ToggleSampleIndex uint64 `protobuf:"varint,2,opt,name=toggleSampleIndex" json:"toggleSampleIndex,omitempty"`
}

func (m *PauseInfo) Reset()                    { *m = PauseInfo{} }
func (m *PauseInfo) String() string            { return proto.CompactTextString(m) }
func (*PauseInfo) ProtoMessage()               {}
func (*PauseInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PauseInfo) GetPlaying() bool {
	if m != nil {
		return m.Playing
	}
	return false
}

func (m *PauseInfo) GetToggleSampleIndex() uint64 {
	if m != nil {
		return m.ToggleSampleIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*Envelope)(nil), "comm.Envelope")
	proto.RegisterType((*TimeSyncRequest)(nil), "comm.TimeSyncRequest")
	proto.RegisterType((*TimeSyncResponse)(nil), "comm.TimeSyncResponse")
	proto.RegisterType((*QueueChunkRequest)(nil), "comm.QueueChunkRequest")
	proto.RegisterType((*PingMessage)(nil), "comm.PingMessage")
	proto.RegisterType((*PongMessage)(nil), "comm.PongMessage")
	proto.RegisterType((*SetVolumeRequest)(nil), "comm.SetVolumeRequest")
	proto.RegisterType((*SubscribeChannelRequest)(nil), "comm.SubscribeChannelRequest")
	proto.RegisterType((*NewSongInfo)(nil), "comm.NewSongInfo")
	proto.RegisterType((*ChunkInfo)(nil), "comm.ChunkInfo")
	proto.RegisterType((*PauseInfo)(nil), "comm.PauseInfo")
	proto.RegisterEnum("comm.Channel", Channel_name, Channel_value)
}

func init() { proto.RegisterFile("comm/packages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x71, 0x9b, 0x6d, 0xcd, 0xed, 0x07, 0x99, 0x91, 0x46, 0x1e, 0x10, 0xaa, 0xf2, 0x00,
	0x55, 0x85, 0x8a, 0x18, 0x12, 0xef, 0x5b, 0x19, 0xa2, 0x52, 0xf7, 0x03, 0x67, 0x20, 0xf1, 0xe8,
	0xa6, 0xd7, 0x34, 0x5a, 0x6a, 0x87, 0xd8, 0xe9, 0x28, 0xff, 0x01, 0xfc, 0x4f, 0xfc, 0x6f, 0xc8,
	0x6e, 0xb2, 0x64, 0x74, 0xf0, 0x96, 0xef, 0xe7, 0xbe, 0xe7, 0xf3, 0xf9, 0x2e, 0xf0, 0x24, 0x92,
	0x8b, 0xc5, 0xeb, 0x8c, 0x47, 0x37, 0x3c, 0x46, 0x35, 0xc8, 0x72, 0xa9, 0x25, 0x75, 0x0c, 0x0c,
	0x8e, 0xa1, 0x73, 0x26, 0x96, 0x98, 0xca, 0x0c, 0x29, 0x05, 0x47, 0xaf, 0x32, 0xf4, 0x49, 0x97,
	0xf4, 0x5c, 0x66, 0xbf, 0x0d, 0x9b, 0x72, 0xcd, 0xfd, 0x56, 0x97, 0xf4, 0xf6, 0x98, 0xfd, 0x0e,
	0xde, 0xc0, 0xe3, 0xeb, 0x64, 0x81, 0xe1, 0x4a, 0x44, 0x0c, 0xbf, 0x15, 0xa8, 0x34, 0x7d, 0x0e,
	0x10, 0xa5, 0x09, 0x0a, 0x1d, 0xa2, 0x98, 0xda, 0x03, 0xda, 0xac, 0x41, 0x82, 0x5f, 0x04, 0xbc,
	0x3a, 0x47, 0x65, 0x52, 0x28, 0xa4, 0x2f, 0xe0, 0xa0, 0xb6, 0x98, 0x68, 0x99, 0xf8, 0x17, 0x35,
	0x3e, 0x85, 0xf9, 0x12, 0x73, 0x86, 0xd1, 0xd2, 0xfa, 0x5a, 0x6b, 0xdf, 0x7d, 0x5a, 0xfb, 0xee,
	0xce, 0x6b, 0x37, 0x7d, 0x15, 0x0d, 0x7e, 0x13, 0x38, 0xfc, 0x54, 0x60, 0x81, 0xc3, 0x79, 0x21,
	0x6e, 0xaa, 0x16, 0x9e, 0x81, 0xab, 0x34, 0xcf, 0x75, 0xe3, 0x22, 0x35, 0xa0, 0x3e, 0xec, 0x44,
	0xc6, 0x3d, 0x9a, 0x96, 0xc5, 0x2b, 0x49, 0xbb, 0xe0, 0x2a, 0xbe, 0xc8, 0x52, 0x1c, 0xcb, 0x5b,
	0xbf, 0xdd, 0x6d, 0xf7, 0xc8, 0x69, 0xcb, 0x23, 0xac, 0x86, 0x34, 0x00, 0x58, 0x8b, 0x8f, 0x49,
	0x3c, 0xf7, 0x9d, 0x3b, 0x4b, 0x83, 0xd2, 0x3e, 0x78, 0xb3, 0x24, 0x57, 0x3a, 0xb4, 0x68, 0x24,
	0xa6, 0xf8, 0xdd, 0xdf, 0xea, 0x92, 0x9e, 0xc3, 0x36, 0x78, 0xb0, 0x0f, 0xbb, 0x57, 0x89, 0x88,
	0xcf, 0x51, 0x29, 0x1e, 0xa3, 0x95, 0xb2, 0x96, 0x7d, 0xf0, 0x42, 0xd4, 0x5f, 0x64, 0x5a, 0x2c,
	0xb0, 0xea, 0xed, 0x08, 0xb6, 0x97, 0x16, 0xd8, 0xc6, 0x08, 0x2b, 0x55, 0x70, 0x0a, 0x4f, 0xc3,
	0x62, 0xa2, 0xa2, 0x3c, 0x99, 0xe0, 0x70, 0xce, 0x85, 0xc0, 0xb4, 0x4a, 0x79, 0x69, 0x1a, 0xb6,
	0xc4, 0xe6, 0x1c, 0x1c, 0xef, 0x0f, 0xcc, 0xc2, 0x0c, 0x2a, 0x5b, 0x15, 0x0d, 0x7e, 0x12, 0xd8,
	0xbd, 0xc0, 0xdb, 0x50, 0x8a, 0x78, 0x24, 0x66, 0x92, 0xbe, 0x83, 0xa3, 0xc6, 0x8d, 0x2f, 0x67,
	0xeb, 0x80, 0xe9, 0x87, 0xd8, 0x7e, 0xfe, 0x11, 0xa5, 0x01, 0xec, 0x29, 0x29, 0xe2, 0x0f, 0x49,
	0x8a, 0x17, 0xbc, 0x9c, 0xb1, 0xcb, 0xee, 0x31, 0xb3, 0x66, 0x46, 0x8f, 0x51, 0xc4, 0x7a, 0x5e,
	0x4e, 0xb7, 0x41, 0x02, 0x05, 0xae, 0x9d, 0xa9, 0xbd, 0xc8, 0xff, 0x07, 0xfa, 0xd0, 0x83, 0xb7,
	0x1e, 0x7e, 0x70, 0x73, 0x92, 0x9d, 0x76, 0x98, 0xfc, 0x58, 0xef, 0x94, 0xc3, 0x6a, 0x10, 0x84,
	0xe0, 0x5e, 0xf1, 0x42, 0xa1, 0x2d, 0xea, 0xc3, 0x4e, 0x96, 0xf2, 0x55, 0x22, 0x62, 0x5b, 0xb2,
	0xc3, 0x2a, 0x49, 0x5f, 0xc1, 0xa1, 0x96, 0x71, 0x9c, 0xe2, 0x66, 0xc5, 0xcd, 0x40, 0xbf, 0x0f,
	0x3b, 0xe5, 0x4b, 0x53, 0x17, 0xb6, 0x4e, 0x3e, 0xbf, 0x1f, 0x5d, 0x7a, 0x8f, 0x68, 0x07, 0x9c,
	0xf3, 0xb3, 0xeb, 0x13, 0x8f, 0x50, 0x80, 0xed, 0xf1, 0x57, 0x36, 0x1a, 0x86, 0x5e, 0x6b, 0xb2,
	0x6d, 0x7f, 0xe8, 0xb7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x50, 0xbc, 0x1c, 0xed, 0xe7, 0x03,
	0x00, 0x00,
}
