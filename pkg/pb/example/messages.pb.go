// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package example

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Probe struct {
	Mac                  string   `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	Bssid                string   `protobuf:"bytes,2,opt,name=bssid,proto3" json:"bssid,omitempty"`
	Timestamp            string   `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Ssid                 string   `protobuf:"bytes,4,opt,name=ssid,proto3" json:"ssid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Probe) Reset()         { *m = Probe{} }
func (m *Probe) String() string { return proto.CompactTextString(m) }
func (*Probe) ProtoMessage()    {}
func (*Probe) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *Probe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Probe.Unmarshal(m, b)
}
func (m *Probe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Probe.Marshal(b, m, deterministic)
}
func (m *Probe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Probe.Merge(m, src)
}
func (m *Probe) XXX_Size() int {
	return xxx_messageInfo_Probe.Size(m)
}
func (m *Probe) XXX_DiscardUnknown() {
	xxx_messageInfo_Probe.DiscardUnknown(m)
}

var xxx_messageInfo_Probe proto.InternalMessageInfo

func (m *Probe) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *Probe) GetBssid() string {
	if m != nil {
		return m.Bssid
	}
	return ""
}

func (m *Probe) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Probe) GetSsid() string {
	if m != nil {
		return m.Ssid
	}
	return ""
}

type ExampleRequest struct {
	ApId                 string   `protobuf:"bytes,1,opt,name=ap_id,json=apId,proto3" json:"ap_id,omitempty"`
	ProbeRequests        []*Probe `protobuf:"bytes,2,rep,name=probe_requests,json=probeRequests,proto3" json:"probe_requests,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExampleRequest) Reset()         { *m = ExampleRequest{} }
func (m *ExampleRequest) String() string { return proto.CompactTextString(m) }
func (*ExampleRequest) ProtoMessage()    {}
func (*ExampleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *ExampleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExampleRequest.Unmarshal(m, b)
}
func (m *ExampleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExampleRequest.Marshal(b, m, deterministic)
}
func (m *ExampleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExampleRequest.Merge(m, src)
}
func (m *ExampleRequest) XXX_Size() int {
	return xxx_messageInfo_ExampleRequest.Size(m)
}
func (m *ExampleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExampleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExampleRequest proto.InternalMessageInfo

func (m *ExampleRequest) GetApId() string {
	if m != nil {
		return m.ApId
	}
	return ""
}

func (m *ExampleRequest) GetProbeRequests() []*Probe {
	if m != nil {
		return m.ProbeRequests
	}
	return nil
}

type ExampleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExampleResponse) Reset()         { *m = ExampleResponse{} }
func (m *ExampleResponse) String() string { return proto.CompactTextString(m) }
func (*ExampleResponse) ProtoMessage()    {}
func (*ExampleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *ExampleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExampleResponse.Unmarshal(m, b)
}
func (m *ExampleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExampleResponse.Marshal(b, m, deterministic)
}
func (m *ExampleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExampleResponse.Merge(m, src)
}
func (m *ExampleResponse) XXX_Size() int {
	return xxx_messageInfo_ExampleResponse.Size(m)
}
func (m *ExampleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExampleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExampleResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Probe)(nil), "example.Probe")
	proto.RegisterType((*ExampleRequest)(nil), "example.ExampleRequest")
	proto.RegisterType((*ExampleResponse)(nil), "example.ExampleResponse")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x69, 0xfe, 0x28, 0xdd, 0x62, 0xac, 0xa3, 0x68, 0x28, 0x3d, 0x94, 0x9c, 0x8a, 0x87,
	0x04, 0x2b, 0x5e, 0xbc, 0xf7, 0xe0, 0x4d, 0xe2, 0xb5, 0x50, 0x26, 0xcd, 0x10, 0x16, 0x9a, 0xdd,
	0x35, 0xb3, 0x16, 0xcf, 0xbe, 0x82, 0x8f, 0xe6, 0x2b, 0xf8, 0x20, 0xd2, 0xcd, 0x12, 0x11, 0x6f,
	0xb3, 0xf3, 0xfd, 0xf6, 0xfb, 0xbe, 0x5d, 0x91, 0xb4, 0xc4, 0x8c, 0x0d, 0x71, 0x6e, 0x3a, 0x6d,
	0x35, 0x9c, 0xd2, 0x3b, 0xb6, 0x66, 0x4f, 0xb3, 0x79, 0xa3, 0x75, 0xb3, 0xa7, 0x02, 0x8d, 0x2c,
	0x50, 0x29, 0x6d, 0xd1, 0x4a, 0xad, 0x3c, 0x96, 0xa1, 0x88, 0x9f, 0x3b, 0x5d, 0x11, 0x4c, 0x45,
	0xd8, 0xe2, 0x2e, 0x1d, 0x2d, 0x46, 0xcb, 0x71, 0x79, 0x1c, 0xe1, 0x4a, 0xc4, 0x15, 0xb3, 0xac,
	0xd3, 0xc0, 0xed, 0xfa, 0x03, 0xcc, 0xc5, 0xd8, 0xca, 0x96, 0xd8, 0x62, 0x6b, 0xd2, 0xd0, 0x29,
	0xbf, 0x0b, 0x00, 0x11, 0xb9, 0x2b, 0x91, 0x13, 0xdc, 0x9c, 0x6d, 0x44, 0xb2, 0xee, 0xbb, 0x94,
	0xf4, 0xfa, 0x46, 0x6c, 0xe1, 0x52, 0xc4, 0x68, 0xb6, 0xb2, 0xf6, 0x69, 0x11, 0x9a, 0xa7, 0x1a,
	0x1e, 0x44, 0x62, 0x8e, 0x4d, 0xb6, 0x5d, 0x4f, 0x71, 0x1a, 0x2c, 0xc2, 0xe5, 0x64, 0x95, 0xe4,
	0xfe, 0x25, 0xb9, 0x2b, 0x5a, 0x9e, 0x39, 0xca, 0x5b, 0x71, 0x76, 0x21, 0xce, 0x07, 0x77, 0x36,
	0x5a, 0x31, 0xad, 0xd4, 0x10, 0xf8, 0x42, 0xdd, 0x41, 0xee, 0x08, 0x36, 0x03, 0xb4, 0x56, 0xb5,
	0xd1, 0x52, 0x59, 0xb8, 0x19, 0x6c, 0xff, 0x96, 0x9b, 0xa5, 0xff, 0x85, 0xde, 0x37, 0xbb, 0xfe,
	0xf8, 0xfa, 0xfe, 0x0c, 0xa6, 0xd9, 0xa4, 0x38, 0xdc, 0x15, 0x1e, 0x7a, 0x1c, 0xdd, 0x56, 0x27,
	0xee, 0x2b, 0xef, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x46, 0x50, 0x95, 0x83, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleServiceClient interface {
	ExampleEndpoint(ctx context.Context, in *ExampleRequest, opts ...grpc.CallOption) (*ExampleResponse, error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) ExampleEndpoint(ctx context.Context, in *ExampleRequest, opts ...grpc.CallOption) (*ExampleResponse, error) {
	out := new(ExampleResponse)
	err := c.cc.Invoke(ctx, "/example.ExampleService/ExampleEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
type ExampleServiceServer interface {
	ExampleEndpoint(context.Context, *ExampleRequest) (*ExampleResponse, error)
}

// UnimplementedExampleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (*UnimplementedExampleServiceServer) ExampleEndpoint(ctx context.Context, req *ExampleRequest) (*ExampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExampleEndpoint not implemented")
}

func RegisterExampleServiceServer(s *grpc.Server, srv ExampleServiceServer) {
	s.RegisterService(&_ExampleService_serviceDesc, srv)
}

func _ExampleService_ExampleEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).ExampleEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.ExampleService/ExampleEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).ExampleEndpoint(ctx, req.(*ExampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExampleEndpoint",
			Handler:    _ExampleService_ExampleEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messages.proto",
}