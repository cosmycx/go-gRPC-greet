// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator.proto

package calculator_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Sum struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sum) Reset()         { *m = Sum{} }
func (m *Sum) String() string { return proto.CompactTextString(m) }
func (*Sum) ProtoMessage()    {}
func (*Sum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{0}
}

func (m *Sum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sum.Unmarshal(m, b)
}
func (m *Sum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sum.Marshal(b, m, deterministic)
}
func (m *Sum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sum.Merge(m, src)
}
func (m *Sum) XXX_Size() int {
	return xxx_messageInfo_Sum.Size(m)
}
func (m *Sum) XXX_DiscardUnknown() {
	xxx_messageInfo_Sum.DiscardUnknown(m)
}

var xxx_messageInfo_Sum proto.InternalMessageInfo

func (m *Sum) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Sum) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type SumRequest struct {
	Sum                  *Sum     `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{1}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetSum() *Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

type SumResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{2}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type FibonacciRequest struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonacciRequest) Reset()         { *m = FibonacciRequest{} }
func (m *FibonacciRequest) String() string { return proto.CompactTextString(m) }
func (*FibonacciRequest) ProtoMessage()    {}
func (*FibonacciRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{3}
}

func (m *FibonacciRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonacciRequest.Unmarshal(m, b)
}
func (m *FibonacciRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonacciRequest.Marshal(b, m, deterministic)
}
func (m *FibonacciRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonacciRequest.Merge(m, src)
}
func (m *FibonacciRequest) XXX_Size() int {
	return xxx_messageInfo_FibonacciRequest.Size(m)
}
func (m *FibonacciRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonacciRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FibonacciRequest proto.InternalMessageInfo

func (m *FibonacciRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type FibonacciResponse struct {
	ResultNum            int32    `protobuf:"varint,1,opt,name=resultNum,proto3" json:"resultNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonacciResponse) Reset()         { *m = FibonacciResponse{} }
func (m *FibonacciResponse) String() string { return proto.CompactTextString(m) }
func (*FibonacciResponse) ProtoMessage()    {}
func (*FibonacciResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{4}
}

func (m *FibonacciResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonacciResponse.Unmarshal(m, b)
}
func (m *FibonacciResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonacciResponse.Marshal(b, m, deterministic)
}
func (m *FibonacciResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonacciResponse.Merge(m, src)
}
func (m *FibonacciResponse) XXX_Size() int {
	return xxx_messageInfo_FibonacciResponse.Size(m)
}
func (m *FibonacciResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonacciResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FibonacciResponse proto.InternalMessageInfo

func (m *FibonacciResponse) GetResultNum() int32 {
	if m != nil {
		return m.ResultNum
	}
	return 0
}

func init() {
	proto.RegisterType((*Sum)(nil), "calculator.Sum")
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*FibonacciRequest)(nil), "calculator.FibonacciRequest")
	proto.RegisterType((*FibonacciResponse)(nil), "calculator.FibonacciResponse")
}

func init() { proto.RegisterFile("calculator.proto", fileDescriptor_c686ea360062a8cf) }

var fileDescriptor_c686ea360062a8cf = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x29, 0x72, 0x31, 0x07, 0x97, 0xe6, 0x0a, 0xf1, 0x70, 0x31, 0x26, 0x4a, 0x30, 0x2a, 0x30,
	0x6a, 0xb0, 0x06, 0x31, 0x26, 0x82, 0x78, 0x49, 0x12, 0x4c, 0x10, 0x5e, 0x92, 0x92, 0x3e, 0x17,
	0x57, 0x70, 0x69, 0x6e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x22, 0x17, 0x73, 0x71,
	0x69, 0x2e, 0x58, 0x2d, 0xb7, 0x11, 0xbf, 0x1e, 0x92, 0xe1, 0x20, 0x45, 0x20, 0x39, 0x25, 0x55,
	0x2e, 0x6e, 0xb0, 0x86, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0xa2, 0xd4,
	0xe2, 0xd2, 0x9c, 0x12, 0xa8, 0x05, 0x50, 0x9e, 0x92, 0x0a, 0x97, 0x80, 0x5b, 0x66, 0x52, 0x7e,
	0x5e, 0x62, 0x72, 0x72, 0x26, 0xcc, 0x74, 0x01, 0x2e, 0xe6, 0x3c, 0xa8, 0xe9, 0xac, 0x41, 0x20,
	0xa6, 0x92, 0x21, 0x97, 0x20, 0x92, 0x2a, 0xa8, 0x91, 0x32, 0x5c, 0x9c, 0x10, 0x43, 0xfc, 0xe0,
	0x8a, 0x11, 0x02, 0x46, 0x2b, 0x19, 0xb9, 0x04, 0x9d, 0xe1, 0xee, 0x0a, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0x15, 0xb2, 0x07, 0x7b, 0x03, 0xc6, 0x13, 0x43, 0x77, 0x39, 0xc4, 0x01, 0x52, 0xe2,
	0x18, 0xe2, 0x10, 0x2b, 0x95, 0x18, 0x84, 0x82, 0x91, 0xdc, 0x0b, 0x33, 0x46, 0x06, 0x59, 0x39,
	0xba, 0x6f, 0xa4, 0x64, 0x71, 0xc8, 0xc2, 0x8c, 0x34, 0x60, 0x74, 0xe2, 0x8f, 0xe2, 0x45, 0xa8,
	0x89, 0x2f, 0x48, 0x4a, 0x62, 0x03, 0xc7, 0x91, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xff, 0x1a,
	0xa8, 0xfa, 0xb7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	SumService(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	FibonacciService(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (CalculatorService_FibonacciServiceClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) SumService(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/SumService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) FibonacciService(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (CalculatorService_FibonacciServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/FibonacciService", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFibonacciServiceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_FibonacciServiceClient interface {
	Recv() (*FibonacciResponse, error)
	grpc.ClientStream
}

type calculatorServiceFibonacciServiceClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFibonacciServiceClient) Recv() (*FibonacciResponse, error) {
	m := new(FibonacciResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	SumService(context.Context, *SumRequest) (*SumResponse, error)
	FibonacciService(*FibonacciRequest, CalculatorService_FibonacciServiceServer) error
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_SumService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).SumService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/SumService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).SumService(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_FibonacciService_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FibonacciRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).FibonacciService(m, &calculatorServiceFibonacciServiceServer{stream})
}

type CalculatorService_FibonacciServiceServer interface {
	Send(*FibonacciResponse) error
	grpc.ServerStream
}

type calculatorServiceFibonacciServiceServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFibonacciServiceServer) Send(m *FibonacciResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SumService",
			Handler:    _CalculatorService_SumService_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FibonacciService",
			Handler:       _CalculatorService_FibonacciService_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
