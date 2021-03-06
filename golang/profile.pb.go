// Code generated by protoc-gen-go.
// source: profile.proto
// DO NOT EDIT!

package dialog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/gogo/protobuf/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Changing account's name
type RequestEditName struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *RequestEditName) Reset()                    { *m = RequestEditName{} }
func (m *RequestEditName) String() string            { return proto.CompactTextString(m) }
func (*RequestEditName) ProtoMessage()               {}
func (*RequestEditName) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{0} }

func (m *RequestEditName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Changing account's nickname
type RequestEditNickName struct {
	Nickname *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=nickname" json:"nickname,omitempty"`
}

func (m *RequestEditNickName) Reset()                    { *m = RequestEditNickName{} }
func (m *RequestEditNickName) String() string            { return proto.CompactTextString(m) }
func (*RequestEditNickName) ProtoMessage()               {}
func (*RequestEditNickName) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{1} }

func (m *RequestEditNickName) GetNickname() *google_protobuf.StringValue {
	if m != nil {
		return m.Nickname
	}
	return nil
}

// Checking availability of nickname
type RequestCheckNickName struct {
	Nickname string `protobuf:"bytes,1,opt,name=nickname" json:"nickname,omitempty"`
}

func (m *RequestCheckNickName) Reset()                    { *m = RequestCheckNickName{} }
func (m *RequestCheckNickName) String() string            { return proto.CompactTextString(m) }
func (*RequestCheckNickName) ProtoMessage()               {}
func (*RequestCheckNickName) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{2} }

func (m *RequestCheckNickName) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

// Changing about information
type RequestEditAbout struct {
	About *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=about" json:"about,omitempty"`
}

func (m *RequestEditAbout) Reset()                    { *m = RequestEditAbout{} }
func (m *RequestEditAbout) String() string            { return proto.CompactTextString(m) }
func (*RequestEditAbout) ProtoMessage()               {}
func (*RequestEditAbout) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{3} }

func (m *RequestEditAbout) GetAbout() *google_protobuf.StringValue {
	if m != nil {
		return m.About
	}
	return nil
}

// Changing account's avatar
type RequestEditAvatar struct {
	FileLocation *FileLocation `protobuf:"bytes,1,opt,name=file_location,json=fileLocation" json:"file_location,omitempty"`
}

func (m *RequestEditAvatar) Reset()                    { *m = RequestEditAvatar{} }
func (m *RequestEditAvatar) String() string            { return proto.CompactTextString(m) }
func (*RequestEditAvatar) ProtoMessage()               {}
func (*RequestEditAvatar) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{4} }

func (m *RequestEditAvatar) GetFileLocation() *FileLocation {
	if m != nil {
		return m.FileLocation
	}
	return nil
}

type ResponseEditAvatar struct {
	Avatar *Avatar `protobuf:"bytes,1,opt,name=avatar" json:"avatar,omitempty"`
	Seq    int32   `protobuf:"varint,2,opt,name=seq" json:"seq,omitempty"`
	State  []byte  `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
}

func (m *ResponseEditAvatar) Reset()                    { *m = ResponseEditAvatar{} }
func (m *ResponseEditAvatar) String() string            { return proto.CompactTextString(m) }
func (*ResponseEditAvatar) ProtoMessage()               {}
func (*ResponseEditAvatar) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{5} }

func (m *ResponseEditAvatar) GetAvatar() *Avatar {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *ResponseEditAvatar) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *ResponseEditAvatar) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

// Removing account's avatar
type RequestRemoveAvatar struct {
}

func (m *RequestRemoveAvatar) Reset()                    { *m = RequestRemoveAvatar{} }
func (m *RequestRemoveAvatar) String() string            { return proto.CompactTextString(m) }
func (*RequestRemoveAvatar) ProtoMessage()               {}
func (*RequestRemoveAvatar) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{6} }

// Updating user's time zone
type RequestEditMyTimeZone struct {
	Tz string `protobuf:"bytes,1,opt,name=tz" json:"tz,omitempty"`
}

func (m *RequestEditMyTimeZone) Reset()                    { *m = RequestEditMyTimeZone{} }
func (m *RequestEditMyTimeZone) String() string            { return proto.CompactTextString(m) }
func (*RequestEditMyTimeZone) ProtoMessage()               {}
func (*RequestEditMyTimeZone) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{7} }

func (m *RequestEditMyTimeZone) GetTz() string {
	if m != nil {
		return m.Tz
	}
	return ""
}

// Changing preffered languages
type RequestEditMyPreferredLanguages struct {
	PreferredLanguages []string `protobuf:"bytes,1,rep,name=preferred_languages,json=preferredLanguages" json:"preferred_languages,omitempty"`
}

func (m *RequestEditMyPreferredLanguages) Reset()         { *m = RequestEditMyPreferredLanguages{} }
func (m *RequestEditMyPreferredLanguages) String() string { return proto.CompactTextString(m) }
func (*RequestEditMyPreferredLanguages) ProtoMessage()    {}
func (*RequestEditMyPreferredLanguages) Descriptor() ([]byte, []int) {
	return fileDescriptor17, []int{8}
}

func (m *RequestEditMyPreferredLanguages) GetPreferredLanguages() []string {
	if m != nil {
		return m.PreferredLanguages
	}
	return nil
}

// Changing user's sex
type RequestEditSex struct {
	Sex Sex `protobuf:"varint,1,opt,name=sex,enum=dialog.Sex" json:"sex,omitempty"`
}

func (m *RequestEditSex) Reset()                    { *m = RequestEditSex{} }
func (m *RequestEditSex) String() string            { return proto.CompactTextString(m) }
func (*RequestEditSex) ProtoMessage()               {}
func (*RequestEditSex) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{9} }

func (m *RequestEditSex) GetSex() Sex {
	if m != nil {
		return m.Sex
	}
	return Sex_SEX_UNKNOWN
}

// Chaning user custom profile based on scheme
type RequestEditCustomProfile struct {
	CustomProfile string `protobuf:"bytes,1,opt,name=custom_profile,json=customProfile" json:"custom_profile,omitempty"`
}

func (m *RequestEditCustomProfile) Reset()                    { *m = RequestEditCustomProfile{} }
func (m *RequestEditCustomProfile) String() string            { return proto.CompactTextString(m) }
func (*RequestEditCustomProfile) ProtoMessage()               {}
func (*RequestEditCustomProfile) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{10} }

func (m *RequestEditCustomProfile) GetCustomProfile() string {
	if m != nil {
		return m.CustomProfile
	}
	return ""
}

// Changing user's status
type RequestChangeUserStatus struct {
	Status *UserStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *RequestChangeUserStatus) Reset()                    { *m = RequestChangeUserStatus{} }
func (m *RequestChangeUserStatus) String() string            { return proto.CompactTextString(m) }
func (*RequestChangeUserStatus) ProtoMessage()               {}
func (*RequestChangeUserStatus) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{11} }

func (m *RequestChangeUserStatus) GetStatus() *UserStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestEditName)(nil), "dialog.RequestEditName")
	proto.RegisterType((*RequestEditNickName)(nil), "dialog.RequestEditNickName")
	proto.RegisterType((*RequestCheckNickName)(nil), "dialog.RequestCheckNickName")
	proto.RegisterType((*RequestEditAbout)(nil), "dialog.RequestEditAbout")
	proto.RegisterType((*RequestEditAvatar)(nil), "dialog.RequestEditAvatar")
	proto.RegisterType((*ResponseEditAvatar)(nil), "dialog.ResponseEditAvatar")
	proto.RegisterType((*RequestRemoveAvatar)(nil), "dialog.RequestRemoveAvatar")
	proto.RegisterType((*RequestEditMyTimeZone)(nil), "dialog.RequestEditMyTimeZone")
	proto.RegisterType((*RequestEditMyPreferredLanguages)(nil), "dialog.RequestEditMyPreferredLanguages")
	proto.RegisterType((*RequestEditSex)(nil), "dialog.RequestEditSex")
	proto.RegisterType((*RequestEditCustomProfile)(nil), "dialog.RequestEditCustomProfile")
	proto.RegisterType((*RequestChangeUserStatus)(nil), "dialog.RequestChangeUserStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Profile service

type ProfileClient interface {
	EditName(ctx context.Context, in *RequestEditName, opts ...grpc.CallOption) (*ResponseSeq, error)
	EditNickName(ctx context.Context, in *RequestEditNickName, opts ...grpc.CallOption) (*ResponseSeq, error)
	CheckNickName(ctx context.Context, in *RequestCheckNickName, opts ...grpc.CallOption) (*ResponseBool, error)
	EditAbout(ctx context.Context, in *RequestEditAbout, opts ...grpc.CallOption) (*ResponseSeq, error)
	EditAvatar(ctx context.Context, in *RequestEditAvatar, opts ...grpc.CallOption) (*ResponseEditAvatar, error)
	RemoveAvatar(ctx context.Context, in *RequestRemoveAvatar, opts ...grpc.CallOption) (*ResponseSeq, error)
	EditMyTimeZone(ctx context.Context, in *RequestEditMyTimeZone, opts ...grpc.CallOption) (*ResponseSeq, error)
	EditMyPreferredLanguages(ctx context.Context, in *RequestEditMyPreferredLanguages, opts ...grpc.CallOption) (*ResponseSeq, error)
	EditSex(ctx context.Context, in *RequestEditSex, opts ...grpc.CallOption) (*ResponseSeq, error)
	EditCustomProfile(ctx context.Context, in *RequestEditCustomProfile, opts ...grpc.CallOption) (*ResponseSeq, error)
	ChangeUserStatus(ctx context.Context, in *RequestChangeUserStatus, opts ...grpc.CallOption) (*ResponseSeq, error)
}

type profileClient struct {
	cc *grpc.ClientConn
}

func NewProfileClient(cc *grpc.ClientConn) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) EditName(ctx context.Context, in *RequestEditName, opts ...grpc.CallOption) (*ResponseSeq, error) {
	out := new(ResponseSeq)
	err := grpc.Invoke(ctx, "/dialog.Profile/EditName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) EditNickName(ctx context.Context, in *RequestEditNickName, opts ...grpc.CallOption) (*ResponseSeq, error) {
	out := new(ResponseSeq)
	err := grpc.Invoke(ctx, "/dialog.Profile/EditNickName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) CheckNickName(ctx context.Context, in *RequestCheckNickName, opts ...grpc.CallOption) (*ResponseBool, error) {
	out := new(ResponseBool)
	err := grpc.Invoke(ctx, "/dialog.Profile/CheckNickName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) EditAbout(ctx context.Context, in *RequestEditAbout, opts ...grpc.CallOption) (*ResponseSeq, error) {
	out := new(ResponseSeq)
	err := grpc.Invoke(ctx, "/dialog.Profile/EditAbout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) EditAvatar(ctx context.Context, in *RequestEditAvatar, opts ...grpc.CallOption) (*ResponseEditAvatar, error) {
	out := new(ResponseEditAvatar)
	err := grpc.Invoke(ctx, "/dialog.Profile/EditAvatar", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) RemoveAvatar(ctx context.Context, in *RequestRemoveAvatar, opts ...grpc.CallOption) (*ResponseSeq, error) {
	out := new(ResponseSeq)
	err := grpc.Invoke(ctx, "/dialog.Profile/RemoveAvatar", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) EditMyTimeZone(ctx context.Context, in *RequestEditMyTimeZone, opts ...grpc.CallOption) (*ResponseSeq, error) {
	out := new(ResponseSeq)
	err := grpc.Invoke(ctx, "/dialog.Profile/EditMyTimeZone", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) EditMyPreferredLanguages(ctx context.Context, in *RequestEditMyPreferredLanguages, opts ...grpc.CallOption) (*ResponseSeq, error) {
	out := new(ResponseSeq)
	err := grpc.Invoke(ctx, "/dialog.Profile/EditMyPreferredLanguages", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) EditSex(ctx context.Context, in *RequestEditSex, opts ...grpc.CallOption) (*ResponseSeq, error) {
	out := new(ResponseSeq)
	err := grpc.Invoke(ctx, "/dialog.Profile/EditSex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) EditCustomProfile(ctx context.Context, in *RequestEditCustomProfile, opts ...grpc.CallOption) (*ResponseSeq, error) {
	out := new(ResponseSeq)
	err := grpc.Invoke(ctx, "/dialog.Profile/EditCustomProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) ChangeUserStatus(ctx context.Context, in *RequestChangeUserStatus, opts ...grpc.CallOption) (*ResponseSeq, error) {
	out := new(ResponseSeq)
	err := grpc.Invoke(ctx, "/dialog.Profile/ChangeUserStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Profile service

type ProfileServer interface {
	EditName(context.Context, *RequestEditName) (*ResponseSeq, error)
	EditNickName(context.Context, *RequestEditNickName) (*ResponseSeq, error)
	CheckNickName(context.Context, *RequestCheckNickName) (*ResponseBool, error)
	EditAbout(context.Context, *RequestEditAbout) (*ResponseSeq, error)
	EditAvatar(context.Context, *RequestEditAvatar) (*ResponseEditAvatar, error)
	RemoveAvatar(context.Context, *RequestRemoveAvatar) (*ResponseSeq, error)
	EditMyTimeZone(context.Context, *RequestEditMyTimeZone) (*ResponseSeq, error)
	EditMyPreferredLanguages(context.Context, *RequestEditMyPreferredLanguages) (*ResponseSeq, error)
	EditSex(context.Context, *RequestEditSex) (*ResponseSeq, error)
	EditCustomProfile(context.Context, *RequestEditCustomProfile) (*ResponseSeq, error)
	ChangeUserStatus(context.Context, *RequestChangeUserStatus) (*ResponseSeq, error)
}

func RegisterProfileServer(s *grpc.Server, srv ProfileServer) {
	s.RegisterService(&_Profile_serviceDesc, srv)
}

func _Profile_EditName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEditName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).EditName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Profile/EditName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).EditName(ctx, req.(*RequestEditName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_EditNickName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEditNickName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).EditNickName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Profile/EditNickName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).EditNickName(ctx, req.(*RequestEditNickName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_CheckNickName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCheckNickName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).CheckNickName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Profile/CheckNickName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).CheckNickName(ctx, req.(*RequestCheckNickName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_EditAbout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEditAbout)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).EditAbout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Profile/EditAbout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).EditAbout(ctx, req.(*RequestEditAbout))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_EditAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEditAvatar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).EditAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Profile/EditAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).EditAvatar(ctx, req.(*RequestEditAvatar))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_RemoveAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRemoveAvatar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).RemoveAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Profile/RemoveAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).RemoveAvatar(ctx, req.(*RequestRemoveAvatar))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_EditMyTimeZone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEditMyTimeZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).EditMyTimeZone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Profile/EditMyTimeZone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).EditMyTimeZone(ctx, req.(*RequestEditMyTimeZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_EditMyPreferredLanguages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEditMyPreferredLanguages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).EditMyPreferredLanguages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Profile/EditMyPreferredLanguages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).EditMyPreferredLanguages(ctx, req.(*RequestEditMyPreferredLanguages))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_EditSex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEditSex)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).EditSex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Profile/EditSex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).EditSex(ctx, req.(*RequestEditSex))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_EditCustomProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEditCustomProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).EditCustomProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Profile/EditCustomProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).EditCustomProfile(ctx, req.(*RequestEditCustomProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_ChangeUserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestChangeUserStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).ChangeUserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Profile/ChangeUserStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).ChangeUserStatus(ctx, req.(*RequestChangeUserStatus))
	}
	return interceptor(ctx, in, info, handler)
}

var _Profile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dialog.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EditName",
			Handler:    _Profile_EditName_Handler,
		},
		{
			MethodName: "EditNickName",
			Handler:    _Profile_EditNickName_Handler,
		},
		{
			MethodName: "CheckNickName",
			Handler:    _Profile_CheckNickName_Handler,
		},
		{
			MethodName: "EditAbout",
			Handler:    _Profile_EditAbout_Handler,
		},
		{
			MethodName: "EditAvatar",
			Handler:    _Profile_EditAvatar_Handler,
		},
		{
			MethodName: "RemoveAvatar",
			Handler:    _Profile_RemoveAvatar_Handler,
		},
		{
			MethodName: "EditMyTimeZone",
			Handler:    _Profile_EditMyTimeZone_Handler,
		},
		{
			MethodName: "EditMyPreferredLanguages",
			Handler:    _Profile_EditMyPreferredLanguages_Handler,
		},
		{
			MethodName: "EditSex",
			Handler:    _Profile_EditSex_Handler,
		},
		{
			MethodName: "EditCustomProfile",
			Handler:    _Profile_EditCustomProfile_Handler,
		},
		{
			MethodName: "ChangeUserStatus",
			Handler:    _Profile_ChangeUserStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}

func init() { proto.RegisterFile("profile.proto", fileDescriptor17) }

var fileDescriptor17 = []byte{
	// 923 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x96, 0xcf, 0x73, 0xdb, 0x44,
	0x14, 0xc7, 0x47, 0x09, 0x71, 0xe2, 0x97, 0xd8, 0x24, 0x8a, 0xd3, 0x28, 0x6e, 0x7e, 0x98, 0x6d,
	0x49, 0x93, 0x00, 0x72, 0x49, 0x7b, 0x21, 0x07, 0x02, 0xe9, 0x40, 0x39, 0x14, 0xa6, 0xc8, 0x85,
	0x03, 0x9d, 0xc1, 0xb3, 0x96, 0x9e, 0xd5, 0x25, 0x92, 0x56, 0xd1, 0x0f, 0xe3, 0x16, 0x66, 0x98,
	0x81, 0x5b, 0x0f, 0x5c, 0xf8, 0xd3, 0xe0, 0x2f, 0xc8, 0x70, 0xea, 0x5f, 0xc1, 0x68, 0x25, 0xd9,
	0x92, 0xb5, 0xc5, 0xd0, 0x53, 0x36, 0xef, 0xbb, 0xef, 0x7d, 0xde, 0xdb, 0x5d, 0x7d, 0xc7, 0xd0,
	0xf0, 0x03, 0x3e, 0x64, 0x0e, 0xea, 0x7e, 0xc0, 0x23, 0xae, 0xd6, 0x2c, 0x46, 0x1d, 0x6e, 0xb7,
	0xf7, 0x6d, 0xce, 0x6d, 0x07, 0xbb, 0x22, 0x3a, 0x88, 0x87, 0xdd, 0x1f, 0x03, 0xea, 0xfb, 0x18,
	0x84, 0xe9, 0xbe, 0xf6, 0x6e, 0xa6, 0x53, 0x9f, 0x75, 0xa9, 0xe7, 0xf1, 0x88, 0x46, 0x8c, 0x7b,
	0xb9, 0xba, 0x61, 0xe1, 0x90, 0x79, 0xac, 0x18, 0xda, 0x74, 0x59, 0x68, 0xa2, 0xe3, 0x50, 0x0f,
	0x79, 0x9c, 0x07, 0xb7, 0x5c, 0xb4, 0x18, 0xed, 0x53, 0xcf, 0xea, 0x27, 0x3d, 0xe4, 0xe1, 0xd5,
	0x38, 0x9c, 0x92, 0xb6, 0x42, 0x93, 0x3a, 0xd4, 0x1f, 0x74, 0xb3, 0xbf, 0x69, 0x98, 0x7c, 0x0d,
	0x6f, 0x1b, 0x78, 0x15, 0x63, 0x18, 0x7d, 0x66, 0xb1, 0xe8, 0x2b, 0xea, 0xa2, 0xda, 0x81, 0xb7,
	0x3c, 0xea, 0xa2, 0xa6, 0x74, 0x94, 0xa3, 0xfa, 0xc5, 0xda, 0xcb, 0x57, 0x77, 0x57, 0xa0, 0xf6,
	0x8c, 0x59, 0x16, 0x7a, 0x86, 0x50, 0xce, 0x76, 0xaf, 0xcf, 0x77, 0x60, 0x9b, 0xb9, 0xba, 0xe5,
	0xd8, 0xba, 0x1d, 0xf8, 0xa6, 0xfe, 0x30, 0xf0, 0xcd, 0xac, 0x0e, 0xf9, 0x19, 0x36, 0x8b, 0x25,
	0x99, 0x79, 0x29, 0xca, 0x3e, 0x84, 0x15, 0x8f, 0x99, 0x97, 0x93, 0xd2, 0xab, 0xa7, 0xbb, 0x7a,
	0x3a, 0xbd, 0x9e, 0x9f, 0x8e, 0xde, 0x8b, 0x02, 0xe6, 0xd9, 0xdf, 0x52, 0x27, 0xc6, 0x8b, 0xc6,
	0xcb, 0x57, 0x77, 0xeb, 0xb0, 0x3c, 0x62, 0x21, 0x1b, 0x38, 0x68, 0x4c, 0x92, 0xe7, 0xd0, 0xfb,
	0xd0, 0xca, 0x96, 0x0f, 0x9e, 0xa1, 0x79, 0x39, 0xc1, 0x1f, 0xcf, 0xe0, 0xeb, 0x6f, 0x0a, 0x08,
	0x60, 0xbd, 0x30, 0xde, 0xa7, 0x03, 0x1e, 0x47, 0xea, 0x27, 0xb0, 0x44, 0x93, 0xc5, 0x7f, 0x1a,
	0xac, 0x7c, 0xa2, 0x69, 0xe2, 0x1c, 0xe6, 0x4f, 0xb0, 0x51, 0x64, 0x8e, 0x68, 0x44, 0x03, 0xf5,
	0x0b, 0x68, 0x24, 0xb7, 0xdd, 0x77, 0xb8, 0x29, 0x5e, 0x4d, 0x06, 0x6f, 0xe9, 0xe9, 0xdb, 0xd3,
	0x3f, 0x67, 0x0e, 0x3e, 0xca, 0xb4, 0x7c, 0x58, 0x93, 0xbb, 0x3e, 0x35, 0x23, 0x63, 0x6d, 0x58,
	0x10, 0xe7, 0xc0, 0x7f, 0x53, 0x40, 0x35, 0x30, 0xf4, 0xb9, 0x17, 0x62, 0x01, 0x7f, 0x08, 0x35,
	0x2a, 0x56, 0x19, 0xb7, 0x99, 0x73, 0x53, 0xdd, 0xc8, 0x54, 0x75, 0x1d, 0x16, 0x43, 0xbc, 0xd2,
	0x16, 0x3a, 0xca, 0xd1, 0x92, 0x91, 0x2c, 0xd5, 0x16, 0x2c, 0x85, 0x11, 0x8d, 0x50, 0x5b, 0xec,
	0x28, 0x47, 0x6b, 0x46, 0xfa, 0xcf, 0xd9, 0xde, 0xf5, 0x79, 0x1b, 0xb4, 0x6a, 0x13, 0x29, 0x98,
	0xdc, 0x9b, 0xbc, 0x2a, 0x03, 0x5d, 0x3e, 0xc2, 0x94, 0x32, 0xa7, 0xf5, 0x27, 0xb0, 0x55, 0x38,
	0xb7, 0x2f, 0x9f, 0x3f, 0x61, 0x2e, 0x7e, 0xc7, 0x3d, 0x54, 0xf7, 0x60, 0x21, 0x7a, 0x21, 0x7f,
	0x07, 0x0b, 0xd1, 0x8b, 0x39, 0x55, 0x7f, 0x81, 0x83, 0x52, 0xd5, 0xc7, 0x01, 0x0e, 0x31, 0x08,
	0xd0, 0x7a, 0x44, 0x3d, 0x3b, 0xa6, 0x36, 0x86, 0xea, 0xc7, 0xb0, 0xe9, 0xe7, 0xd1, 0xbe, 0x93,
	0x87, 0x35, 0xa5, 0xb3, 0x58, 0x05, 0xaa, 0x7e, 0x25, 0x7f, 0x4e, 0x03, 0x4f, 0xa1, 0x59, 0x68,
	0xa0, 0x87, 0x63, 0xf5, 0xbd, 0xe4, 0x90, 0xc7, 0x62, 0xa0, 0xe6, 0xe9, 0x6a, 0x7e, 0x13, 0x3d,
	0x1c, 0xcf, 0xbc, 0xb6, 0x64, 0xd7, 0x9c, 0xe2, 0x1e, 0x68, 0x85, 0xe2, 0x0f, 0xe2, 0x30, 0xe2,
	0xee, 0xe3, 0xd4, 0xdc, 0xd4, 0xfb, 0xd0, 0x34, 0x45, 0xa0, 0x9f, 0xd9, 0x9d, 0xfc, 0x08, 0x1b,
	0x66, 0x31, 0x6b, 0xee, 0xf7, 0xb4, 0x3d, 0xf9, 0x60, 0xa9, 0x67, 0xe3, 0x37, 0x21, 0x06, 0xbd,
	0x88, 0x46, 0x71, 0xa8, 0x7e, 0x04, 0xb5, 0x50, 0xac, 0xb2, 0x27, 0xa6, 0xe6, 0x83, 0x4d, 0xf7,
	0xcc, 0xa2, 0xb3, 0x84, 0x7f, 0x67, 0x9e, 0xfe, 0x55, 0x87, 0xe5, 0x7c, 0xa6, 0xef, 0x61, 0x65,
	0x62, 0x7d, 0xdb, 0x39, 0x60, 0xc6, 0x13, 0xdb, 0x9b, 0x53, 0x21, 0x7d, 0x8f, 0x3d, 0xbc, 0x22,
	0xb7, 0x7f, 0xfd, 0xf3, 0xef, 0x3f, 0x16, 0xf6, 0xc9, 0x4e, 0x77, 0xf4, 0x61, 0x37, 0x61, 0x74,
	0xb3, 0x9a, 0xdd, 0x3c, 0xef, 0x4c, 0x39, 0x51, 0x7f, 0x80, 0xb5, 0x92, 0x0f, 0xde, 0x94, 0x31,
	0x32, 0x51, 0xce, 0x39, 0x12, 0x1c, 0x42, 0xf6, 0xe4, 0x9c, 0x2c, 0x37, 0x61, 0x79, 0xd0, 0x28,
	0xbb, 0xde, 0xee, 0x0c, 0xac, 0xa4, 0xb6, 0x5b, 0xb3, 0xb4, 0x0b, 0xce, 0x1d, 0x72, 0x2c, 0x70,
	0xb7, 0xc8, 0x7e, 0x05, 0x57, 0xca, 0x4e, 0x78, 0x14, 0xea, 0x53, 0x13, 0xd4, 0x24, 0x83, 0x09,
	0x45, 0x3e, 0xd5, 0xbb, 0x02, 0x73, 0x40, 0xda, 0xd2, 0xa9, 0x44, 0x62, 0x82, 0xb8, 0x04, 0x28,
	0x98, 0xce, 0x8e, 0x8c, 0x21, 0xa4, 0x76, 0x7b, 0x16, 0x32, 0xd5, 0xc8, 0xa1, 0x60, 0x75, 0xc8,
	0x4d, 0x39, 0x2b, 0xb5, 0x92, 0xf4, 0xae, 0x8a, 0xee, 0x52, 0xb9, 0xab, 0xa2, 0xf8, 0x7f, 0xef,
	0xaa, 0x64, 0x5b, 0xca, 0x89, 0xea, 0x43, 0x73, 0xd6, 0x94, 0x24, 0xc3, 0x4d, 0x65, 0x39, 0xef,
	0x44, 0xf0, 0x6e, 0x93, 0x03, 0xe9, 0x64, 0xd3, 0xec, 0x84, 0xf8, 0xbb, 0x02, 0xda, 0x6b, 0x1d,
	0xeb, 0x8e, 0x14, 0x5e, 0xdd, 0x28, 0x6f, 0xe3, 0xbe, 0x68, 0x43, 0x27, 0xc7, 0xaf, 0x69, 0xa3,
	0x5a, 0x27, 0x69, 0xe8, 0x29, 0x2c, 0xe7, 0x06, 0x76, 0x43, 0x82, 0xef, 0xe1, 0x58, 0x4e, 0xbb,
	0x25, 0x68, 0x7b, 0x44, 0x93, 0xd2, 0x7a, 0x38, 0x4e, 0x8a, 0x3f, 0x87, 0x8d, 0xaa, 0x81, 0x75,
	0x24, 0x98, 0xd2, 0x0e, 0x39, 0xf0, 0x03, 0x01, 0xbc, 0x43, 0x88, 0x14, 0x58, 0x2a, 0x90, 0xa0,
	0x47, 0xb0, 0x5e, 0xf1, 0xb2, 0x83, 0xca, 0x97, 0x58, 0xde, 0x20, 0x07, 0xbf, 0x2f, 0xc0, 0x87,
	0xe4, 0x1d, 0xc9, 0xb7, 0x58, 0xce, 0x3f, 0x53, 0x4e, 0x2e, 0x76, 0xae, 0xcf, 0x6f, 0x40, 0xab,
	0x68, 0x7a, 0x21, 0x06, 0x23, 0x66, 0x62, 0x38, 0xa8, 0x89, 0x9f, 0x22, 0xf7, 0xfe, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xaf, 0xad, 0x17, 0xd2, 0xa8, 0x0a, 0x00, 0x00,
}
