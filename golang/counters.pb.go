// Code generated by protoc-gen-go.
// source: counters.proto
// DO NOT EDIT!

package dialog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Unread dialogs
type UnreadDialog struct {
	Peer    *Peer `protobuf:"bytes,1,opt,name=peer" json:"peer,omitempty"`
	Counter int32 `protobuf:"varint,2,opt,name=counter" json:"counter,omitempty"`
}

func (m *UnreadDialog) Reset()                    { *m = UnreadDialog{} }
func (m *UnreadDialog) String() string            { return proto.CompactTextString(m) }
func (*UnreadDialog) ProtoMessage()               {}
func (*UnreadDialog) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *UnreadDialog) GetPeer() *Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

func (m *UnreadDialog) GetCounter() int32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

type AppCounters struct {
	// / Global unread counter
	GlobalCounter *google_protobuf.Int32Value `protobuf:"bytes,1,opt,name=global_counter,json=globalCounter" json:"global_counter,omitempty"`
	// / Global count of dialogs with positive counter values
	GlobalDialogsCounter *google_protobuf.Int32Value `protobuf:"bytes,2,opt,name=global_dialogs_counter,json=globalDialogsCounter" json:"global_dialogs_counter,omitempty"`
	// / map of chat peer id to unread to counters
	UnreadDialogs []*UnreadDialog `protobuf:"bytes,3,rep,name=unread_dialogs,json=unreadDialogs" json:"unread_dialogs,omitempty"`
}

func (m *AppCounters) Reset()                    { *m = AppCounters{} }
func (m *AppCounters) String() string            { return proto.CompactTextString(m) }
func (*AppCounters) ProtoMessage()               {}
func (*AppCounters) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *AppCounters) GetGlobalCounter() *google_protobuf.Int32Value {
	if m != nil {
		return m.GlobalCounter
	}
	return nil
}

func (m *AppCounters) GetGlobalDialogsCounter() *google_protobuf.Int32Value {
	if m != nil {
		return m.GlobalDialogsCounter
	}
	return nil
}

func (m *AppCounters) GetUnreadDialogs() []*UnreadDialog {
	if m != nil {
		return m.UnreadDialogs
	}
	return nil
}

// / deprecated
type UpdateCountersChanged struct {
	Counters *AppCounters                `protobuf:"bytes,1,opt,name=counters" json:"counters,omitempty"`
	Ts       *google_protobuf.Int64Value `protobuf:"bytes,2,opt,name=ts" json:"ts,omitempty"`
}

func (m *UpdateCountersChanged) Reset()                    { *m = UpdateCountersChanged{} }
func (m *UpdateCountersChanged) String() string            { return proto.CompactTextString(m) }
func (*UpdateCountersChanged) ProtoMessage()               {}
func (*UpdateCountersChanged) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *UpdateCountersChanged) GetCounters() *AppCounters {
	if m != nil {
		return m.Counters
	}
	return nil
}

func (m *UpdateCountersChanged) GetTs() *google_protobuf.Int64Value {
	if m != nil {
		return m.Ts
	}
	return nil
}

func init() {
	proto.RegisterType((*UnreadDialog)(nil), "dialog.UnreadDialog")
	proto.RegisterType((*AppCounters)(nil), "dialog.AppCounters")
	proto.RegisterType((*UpdateCountersChanged)(nil), "dialog.UpdateCountersChanged")
}

func init() { proto.RegisterFile("counters.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4e, 0x2a, 0x31,
	0x18, 0xc5, 0x33, 0xc3, 0xbd, 0x70, 0x6f, 0x07, 0x48, 0x1c, 0x81, 0x20, 0x1a, 0x43, 0xd8, 0xc8,
	0xaa, 0x63, 0xc0, 0xb8, 0x70, 0xa1, 0x11, 0xdc, 0xe8, 0x4a, 0x4d, 0x70, 0x8b, 0x9d, 0x99, 0x32,
	0x34, 0xa9, 0x6d, 0xd3, 0x76, 0xf0, 0x1d, 0x5c, 0xb8, 0xf1, 0x0d, 0x5d, 0xfa, 0x14, 0xc6, 0x69,
	0x8b, 0x13, 0x34, 0x26, 0xae, 0x9a, 0x7c, 0x7f, 0x7e, 0xdf, 0x39, 0x27, 0x05, 0xcd, 0x84, 0xe7,
	0x4c, 0x63, 0xa9, 0xa0, 0x90, 0x5c, 0xf3, 0xb0, 0x9a, 0x12, 0x44, 0x79, 0xd6, 0xdb, 0xcf, 0x38,
	0xcf, 0x28, 0x8e, 0x8a, 0x6a, 0x9c, 0x2f, 0xa2, 0x47, 0x89, 0x84, 0x58, 0xcf, 0xf5, 0xf6, 0x6c,
	0x1f, 0x09, 0x12, 0x21, 0xc6, 0xb8, 0x46, 0x9a, 0x70, 0xe6, 0xba, 0x5b, 0x29, 0x5e, 0x10, 0x46,
	0xca, 0xa5, 0x40, 0xe0, 0xcf, 0xed, 0xb6, 0x4a, 0x10, 0x45, 0x22, 0x8e, 0xec, 0x6b, 0xca, 0x83,
	0x25, 0xa8, 0xcf, 0x98, 0xc4, 0x28, 0xbd, 0x28, 0x44, 0x84, 0x11, 0xf8, 0xf3, 0xb1, 0xd5, 0xf5,
	0xfa, 0xde, 0x30, 0x18, 0xd5, 0xa1, 0xd1, 0x06, 0xaf, 0x31, 0x96, 0x93, 0xc6, 0xd3, 0xdb, 0xe1,
	0x7f, 0x50, 0x5b, 0x11, 0x45, 0x62, 0x8a, 0x6f, 0x8b, 0xc1, 0xf0, 0x00, 0xd4, 0xac, 0x9f, 0xae,
	0xdf, 0xf7, 0x86, 0x7f, 0x37, 0xa7, 0x5c, 0x77, 0xf0, 0xec, 0x83, 0xe0, 0x5c, 0x88, 0xa9, 0x35,
	0x1f, 0xde, 0x80, 0x66, 0x46, 0x79, 0x8c, 0xe8, 0xdc, 0xed, 0x9b, 0x9b, 0xbb, 0xd0, 0xf8, 0x84,
	0x2e, 0x07, 0x78, 0xc9, 0xf4, 0x78, 0x74, 0x87, 0x68, 0x8e, 0x37, 0xe1, 0x0d, 0x43, 0xb0, 0xcc,
	0xf0, 0x1e, 0x74, 0x2c, 0xd2, 0xc8, 0x56, 0xf3, 0xb2, 0xb4, 0xdf, 0xa1, 0x5b, 0x86, 0x64, 0x62,
	0x51, 0xee, 0xc2, 0x15, 0x68, 0xe6, 0x45, 0x5c, 0xee, 0x42, 0xb7, 0xd2, 0xaf, 0x0c, 0x83, 0x51,
	0xcb, 0x05, 0x55, 0x0e, 0xf3, 0x8b, 0xda, 0xbc, 0xd4, 0x54, 0x83, 0x17, 0x0f, 0xb4, 0x67, 0x22,
	0x45, 0x1a, 0xbb, 0x4c, 0xa6, 0x4b, 0xc4, 0x32, 0x9c, 0x86, 0xa7, 0xe0, 0x9f, 0xfb, 0x23, 0x36,
	0x94, 0x6d, 0xc7, 0x2f, 0x25, 0xb8, 0x89, 0x5f, 0xef, 0x84, 0x27, 0xc0, 0xd7, 0xea, 0x27, 0xcf,
	0xc7, 0x47, 0xdf, 0x7a, 0xf6, 0xb5, 0x9a, 0xec, 0xbc, 0x9e, 0x75, 0x40, 0x8b, 0x3c, 0xc0, 0x94,
	0x66, 0x30, 0x93, 0x22, 0x81, 0x0a, 0xcb, 0x15, 0x49, 0xb0, 0x8a, 0xab, 0x05, 0x62, 0xfc, 0x1e,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0xae, 0xd4, 0xa5, 0xc1, 0x02, 0x00, 0x00,
}
