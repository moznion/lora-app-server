// Code generated by protoc-gen-go.
// source: as.proto
// DO NOT EDIT!

/*
Package as is a generated protocol buffer package.

It is generated from these files:
	as.proto

It has these top-level messages:
	DataRate
	RXInfo
	TXInfo
	JoinRequestRequest
	JoinRequestResponse
	HandleDataUpRequest
	GetDataDownRequest
	GetDataDownResponse
	HandleDataUpResponse
	HandleDataDownACKRequest
	HandleDataDownACKResponse
	HandleErrorRequest
	HandleErrorResponse
*/
package as

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RXWindow int32

const (
	RXWindow_RX1 RXWindow = 0
	RXWindow_RX2 RXWindow = 1
)

var RXWindow_name = map[int32]string{
	0: "RX1",
	1: "RX2",
}
var RXWindow_value = map[string]int32{
	"RX1": 0,
	"RX2": 1,
}

func (x RXWindow) String() string {
	return proto.EnumName(RXWindow_name, int32(x))
}
func (RXWindow) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ErrorType int32

const (
	ErrorType_Generic      ErrorType = 0
	ErrorType_OTAA         ErrorType = 1
	ErrorType_DATA_UP_FCNT ErrorType = 2
	ErrorType_DATA_UP_MIC  ErrorType = 3
)

var ErrorType_name = map[int32]string{
	0: "Generic",
	1: "OTAA",
	2: "DATA_UP_FCNT",
	3: "DATA_UP_MIC",
}
var ErrorType_value = map[string]int32{
	"Generic":      0,
	"OTAA":         1,
	"DATA_UP_FCNT": 2,
	"DATA_UP_MIC":  3,
}

func (x ErrorType) String() string {
	return proto.EnumName(ErrorType_name, int32(x))
}
func (ErrorType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DataRate struct {
	Modulation   string `protobuf:"bytes,1,opt,name=modulation" json:"modulation,omitempty"`
	BandWidth    uint32 `protobuf:"varint,2,opt,name=bandWidth" json:"bandWidth,omitempty"`
	SpreadFactor uint32 `protobuf:"varint,3,opt,name=spreadFactor" json:"spreadFactor,omitempty"`
	Bitrate      uint32 `protobuf:"varint,4,opt,name=bitrate" json:"bitrate,omitempty"`
}

func (m *DataRate) Reset()                    { *m = DataRate{} }
func (m *DataRate) String() string            { return proto.CompactTextString(m) }
func (*DataRate) ProtoMessage()               {}
func (*DataRate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RXInfo struct {
	Mac     []byte  `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	Time    string  `protobuf:"bytes,2,opt,name=time" json:"time,omitempty"`
	Rssi    int32   `protobuf:"varint,3,opt,name=rssi" json:"rssi,omitempty"`
	LoRaSNR float64 `protobuf:"fixed64,4,opt,name=loRaSNR" json:"loRaSNR,omitempty"`
}

func (m *RXInfo) Reset()                    { *m = RXInfo{} }
func (m *RXInfo) String() string            { return proto.CompactTextString(m) }
func (*RXInfo) ProtoMessage()               {}
func (*RXInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type TXInfo struct {
	Frequency int64     `protobuf:"varint,1,opt,name=frequency" json:"frequency,omitempty"`
	DataRate  *DataRate `protobuf:"bytes,2,opt,name=dataRate" json:"dataRate,omitempty"`
	Adr       bool      `protobuf:"varint,3,opt,name=adr" json:"adr,omitempty"`
	CodeRate  string    `protobuf:"bytes,4,opt,name=codeRate" json:"codeRate,omitempty"`
}

func (m *TXInfo) Reset()                    { *m = TXInfo{} }
func (m *TXInfo) String() string            { return proto.CompactTextString(m) }
func (*TXInfo) ProtoMessage()               {}
func (*TXInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TXInfo) GetDataRate() *DataRate {
	if m != nil {
		return m.DataRate
	}
	return nil
}

type JoinRequestRequest struct {
	PhyPayload []byte `protobuf:"bytes,1,opt,name=phyPayload,proto3" json:"phyPayload,omitempty"`
	DevAddr    []byte `protobuf:"bytes,2,opt,name=devAddr,proto3" json:"devAddr,omitempty"`
	NetID      []byte `protobuf:"bytes,3,opt,name=netID,proto3" json:"netID,omitempty"`
}

func (m *JoinRequestRequest) Reset()                    { *m = JoinRequestRequest{} }
func (m *JoinRequestRequest) String() string            { return proto.CompactTextString(m) }
func (*JoinRequestRequest) ProtoMessage()               {}
func (*JoinRequestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type JoinRequestResponse struct {
	PhyPayload  []byte   `protobuf:"bytes,1,opt,name=phyPayload,proto3" json:"phyPayload,omitempty"`
	NwkSKey     []byte   `protobuf:"bytes,2,opt,name=nwkSKey,proto3" json:"nwkSKey,omitempty"`
	RxDelay     uint32   `protobuf:"varint,3,opt,name=rxDelay" json:"rxDelay,omitempty"`
	Rx1DROffset uint32   `protobuf:"varint,4,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	CFList      []uint32 `protobuf:"varint,5,rep,packed,name=cFList" json:"cFList,omitempty"`
	RxWindow    RXWindow `protobuf:"varint,6,opt,name=rxWindow,enum=as.RXWindow" json:"rxWindow,omitempty"`
	Rx2DR       uint32   `protobuf:"varint,7,opt,name=rx2DR" json:"rx2DR,omitempty"`
	RelaxFCnt   bool     `protobuf:"varint,8,opt,name=relaxFCnt" json:"relaxFCnt,omitempty"`
}

func (m *JoinRequestResponse) Reset()                    { *m = JoinRequestResponse{} }
func (m *JoinRequestResponse) String() string            { return proto.CompactTextString(m) }
func (*JoinRequestResponse) ProtoMessage()               {}
func (*JoinRequestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type HandleDataUpRequest struct {
	DevEUI []byte    `protobuf:"bytes,1,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	AppEUI []byte    `protobuf:"bytes,2,opt,name=appEUI,proto3" json:"appEUI,omitempty"`
	FCnt   uint32    `protobuf:"varint,3,opt,name=fCnt" json:"fCnt,omitempty"`
	FPort  uint32    `protobuf:"varint,4,opt,name=fPort" json:"fPort,omitempty"`
	Data   []byte    `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	TxInfo *TXInfo   `protobuf:"bytes,6,opt,name=txInfo" json:"txInfo,omitempty"`
	RxInfo []*RXInfo `protobuf:"bytes,7,rep,name=rxInfo" json:"rxInfo,omitempty"`
}

func (m *HandleDataUpRequest) Reset()                    { *m = HandleDataUpRequest{} }
func (m *HandleDataUpRequest) String() string            { return proto.CompactTextString(m) }
func (*HandleDataUpRequest) ProtoMessage()               {}
func (*HandleDataUpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *HandleDataUpRequest) GetTxInfo() *TXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *HandleDataUpRequest) GetRxInfo() []*RXInfo {
	if m != nil {
		return m.RxInfo
	}
	return nil
}

type GetDataDownRequest struct {
	DevEUI         []byte `protobuf:"bytes,1,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	AppEUI         []byte `protobuf:"bytes,2,opt,name=appEUI,proto3" json:"appEUI,omitempty"`
	MaxPayloadSize uint32 `protobuf:"varint,3,opt,name=maxPayloadSize" json:"maxPayloadSize,omitempty"`
	FCnt           uint32 `protobuf:"varint,4,opt,name=fCnt" json:"fCnt,omitempty"`
}

func (m *GetDataDownRequest) Reset()                    { *m = GetDataDownRequest{} }
func (m *GetDataDownRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDataDownRequest) ProtoMessage()               {}
func (*GetDataDownRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type GetDataDownResponse struct {
	Data      []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Confirmed bool   `protobuf:"varint,2,opt,name=confirmed" json:"confirmed,omitempty"`
	FPort     uint32 `protobuf:"varint,3,opt,name=fPort" json:"fPort,omitempty"`
	MoreData  bool   `protobuf:"varint,4,opt,name=moreData" json:"moreData,omitempty"`
}

func (m *GetDataDownResponse) Reset()                    { *m = GetDataDownResponse{} }
func (m *GetDataDownResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDataDownResponse) ProtoMessage()               {}
func (*GetDataDownResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type HandleDataUpResponse struct {
}

func (m *HandleDataUpResponse) Reset()                    { *m = HandleDataUpResponse{} }
func (m *HandleDataUpResponse) String() string            { return proto.CompactTextString(m) }
func (*HandleDataUpResponse) ProtoMessage()               {}
func (*HandleDataUpResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type HandleDataDownACKRequest struct {
	DevEUI []byte `protobuf:"bytes,1,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	AppEUI []byte `protobuf:"bytes,2,opt,name=appEUI,proto3" json:"appEUI,omitempty"`
	FCnt   uint32 `protobuf:"varint,3,opt,name=fCnt" json:"fCnt,omitempty"`
}

func (m *HandleDataDownACKRequest) Reset()                    { *m = HandleDataDownACKRequest{} }
func (m *HandleDataDownACKRequest) String() string            { return proto.CompactTextString(m) }
func (*HandleDataDownACKRequest) ProtoMessage()               {}
func (*HandleDataDownACKRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type HandleDataDownACKResponse struct {
}

func (m *HandleDataDownACKResponse) Reset()                    { *m = HandleDataDownACKResponse{} }
func (m *HandleDataDownACKResponse) String() string            { return proto.CompactTextString(m) }
func (*HandleDataDownACKResponse) ProtoMessage()               {}
func (*HandleDataDownACKResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type HandleErrorRequest struct {
	DevEUI []byte    `protobuf:"bytes,1,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	AppEUI []byte    `protobuf:"bytes,2,opt,name=appEUI,proto3" json:"appEUI,omitempty"`
	Type   ErrorType `protobuf:"varint,3,opt,name=type,enum=as.ErrorType" json:"type,omitempty"`
	Error  string    `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *HandleErrorRequest) Reset()                    { *m = HandleErrorRequest{} }
func (m *HandleErrorRequest) String() string            { return proto.CompactTextString(m) }
func (*HandleErrorRequest) ProtoMessage()               {}
func (*HandleErrorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type HandleErrorResponse struct {
}

func (m *HandleErrorResponse) Reset()                    { *m = HandleErrorResponse{} }
func (m *HandleErrorResponse) String() string            { return proto.CompactTextString(m) }
func (*HandleErrorResponse) ProtoMessage()               {}
func (*HandleErrorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func init() {
	proto.RegisterType((*DataRate)(nil), "as.DataRate")
	proto.RegisterType((*RXInfo)(nil), "as.RXInfo")
	proto.RegisterType((*TXInfo)(nil), "as.TXInfo")
	proto.RegisterType((*JoinRequestRequest)(nil), "as.JoinRequestRequest")
	proto.RegisterType((*JoinRequestResponse)(nil), "as.JoinRequestResponse")
	proto.RegisterType((*HandleDataUpRequest)(nil), "as.HandleDataUpRequest")
	proto.RegisterType((*GetDataDownRequest)(nil), "as.GetDataDownRequest")
	proto.RegisterType((*GetDataDownResponse)(nil), "as.GetDataDownResponse")
	proto.RegisterType((*HandleDataUpResponse)(nil), "as.HandleDataUpResponse")
	proto.RegisterType((*HandleDataDownACKRequest)(nil), "as.HandleDataDownACKRequest")
	proto.RegisterType((*HandleDataDownACKResponse)(nil), "as.HandleDataDownACKResponse")
	proto.RegisterType((*HandleErrorRequest)(nil), "as.HandleErrorRequest")
	proto.RegisterType((*HandleErrorResponse)(nil), "as.HandleErrorResponse")
	proto.RegisterEnum("as.RXWindow", RXWindow_name, RXWindow_value)
	proto.RegisterEnum("as.ErrorType", ErrorType_name, ErrorType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ApplicationServer service

type ApplicationServerClient interface {
	// JoinRequest requests the application server to validate the join-request and return an encrypted join-accept.
	JoinRequest(ctx context.Context, in *JoinRequestRequest, opts ...grpc.CallOption) (*JoinRequestResponse, error)
	// HandleDataUp publishes data received from an end-device.
	HandleDataUp(ctx context.Context, in *HandleDataUpRequest, opts ...grpc.CallOption) (*HandleDataUpResponse, error)
	// GetDataDown gets data from the downlink queue.
	GetDataDown(ctx context.Context, in *GetDataDownRequest, opts ...grpc.CallOption) (*GetDataDownResponse, error)
	// HandleDataDownACK publishes a data-down ack response.
	HandleDataDownACK(ctx context.Context, in *HandleDataDownACKRequest, opts ...grpc.CallOption) (*HandleDataDownACKResponse, error)
	// HandleError publishes an error message.
	HandleError(ctx context.Context, in *HandleErrorRequest, opts ...grpc.CallOption) (*HandleErrorResponse, error)
}

type applicationServerClient struct {
	cc *grpc.ClientConn
}

func NewApplicationServerClient(cc *grpc.ClientConn) ApplicationServerClient {
	return &applicationServerClient{cc}
}

func (c *applicationServerClient) JoinRequest(ctx context.Context, in *JoinRequestRequest, opts ...grpc.CallOption) (*JoinRequestResponse, error) {
	out := new(JoinRequestResponse)
	err := grpc.Invoke(ctx, "/as.ApplicationServer/JoinRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerClient) HandleDataUp(ctx context.Context, in *HandleDataUpRequest, opts ...grpc.CallOption) (*HandleDataUpResponse, error) {
	out := new(HandleDataUpResponse)
	err := grpc.Invoke(ctx, "/as.ApplicationServer/HandleDataUp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerClient) GetDataDown(ctx context.Context, in *GetDataDownRequest, opts ...grpc.CallOption) (*GetDataDownResponse, error) {
	out := new(GetDataDownResponse)
	err := grpc.Invoke(ctx, "/as.ApplicationServer/GetDataDown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerClient) HandleDataDownACK(ctx context.Context, in *HandleDataDownACKRequest, opts ...grpc.CallOption) (*HandleDataDownACKResponse, error) {
	out := new(HandleDataDownACKResponse)
	err := grpc.Invoke(ctx, "/as.ApplicationServer/HandleDataDownACK", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerClient) HandleError(ctx context.Context, in *HandleErrorRequest, opts ...grpc.CallOption) (*HandleErrorResponse, error) {
	out := new(HandleErrorResponse)
	err := grpc.Invoke(ctx, "/as.ApplicationServer/HandleError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApplicationServer service

type ApplicationServerServer interface {
	// JoinRequest requests the application server to validate the join-request and return an encrypted join-accept.
	JoinRequest(context.Context, *JoinRequestRequest) (*JoinRequestResponse, error)
	// HandleDataUp publishes data received from an end-device.
	HandleDataUp(context.Context, *HandleDataUpRequest) (*HandleDataUpResponse, error)
	// GetDataDown gets data from the downlink queue.
	GetDataDown(context.Context, *GetDataDownRequest) (*GetDataDownResponse, error)
	// HandleDataDownACK publishes a data-down ack response.
	HandleDataDownACK(context.Context, *HandleDataDownACKRequest) (*HandleDataDownACKResponse, error)
	// HandleError publishes an error message.
	HandleError(context.Context, *HandleErrorRequest) (*HandleErrorResponse, error)
}

func RegisterApplicationServerServer(s *grpc.Server, srv ApplicationServerServer) {
	s.RegisterService(&_ApplicationServer_serviceDesc, srv)
}

func _ApplicationServer_JoinRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServer).JoinRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServer/JoinRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServer).JoinRequest(ctx, req.(*JoinRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServer_HandleDataUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleDataUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServer).HandleDataUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServer/HandleDataUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServer).HandleDataUp(ctx, req.(*HandleDataUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServer_GetDataDown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataDownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServer).GetDataDown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServer/GetDataDown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServer).GetDataDown(ctx, req.(*GetDataDownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServer_HandleDataDownACK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleDataDownACKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServer).HandleDataDownACK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServer/HandleDataDownACK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServer).HandleDataDownACK(ctx, req.(*HandleDataDownACKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServer_HandleError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServer).HandleError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServer/HandleError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServer).HandleError(ctx, req.(*HandleErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "as.ApplicationServer",
	HandlerType: (*ApplicationServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinRequest",
			Handler:    _ApplicationServer_JoinRequest_Handler,
		},
		{
			MethodName: "HandleDataUp",
			Handler:    _ApplicationServer_HandleDataUp_Handler,
		},
		{
			MethodName: "GetDataDown",
			Handler:    _ApplicationServer_GetDataDown_Handler,
		},
		{
			MethodName: "HandleDataDownACK",
			Handler:    _ApplicationServer_HandleDataDownACK_Handler,
		},
		{
			MethodName: "HandleError",
			Handler:    _ApplicationServer_HandleError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("as.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 843 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x8e, 0xe3, 0xfc, 0x38, 0x27, 0x69, 0xf1, 0x4e, 0x97, 0xac, 0x09, 0x01, 0x05, 0x5f, 0xa0,
	0x68, 0x2f, 0x2a, 0x6d, 0x78, 0x81, 0x8d, 0x92, 0x66, 0x09, 0x0b, 0xbb, 0xd5, 0x24, 0xd5, 0xf6,
	0x02, 0xb1, 0x9a, 0x7a, 0x26, 0x5a, 0x8b, 0xc4, 0x63, 0xc6, 0xd3, 0x36, 0x46, 0x02, 0xc1, 0x0d,
	0x2f, 0xc7, 0xdb, 0xf0, 0x04, 0x68, 0x7e, 0x9c, 0x38, 0xa4, 0x48, 0xa8, 0xe2, 0x2a, 0xf3, 0x7d,
	0x67, 0x72, 0xbe, 0xf3, 0x9d, 0x73, 0x6c, 0x83, 0x47, 0xb2, 0xf3, 0x54, 0x70, 0xc9, 0x51, 0x95,
	0x64, 0xe1, 0x1f, 0x0e, 0x78, 0x53, 0x22, 0x09, 0x26, 0x92, 0xa1, 0xcf, 0x01, 0x36, 0x9c, 0xde,
	0xae, 0x89, 0x8c, 0x79, 0x12, 0x38, 0x03, 0x67, 0xd8, 0xc2, 0x25, 0x06, 0xf5, 0xa1, 0x75, 0x43,
	0x12, 0xfa, 0x2e, 0xa6, 0xf2, 0x43, 0x50, 0x1d, 0x38, 0xc3, 0x13, 0xbc, 0x27, 0x50, 0x08, 0x9d,
	0x2c, 0x15, 0x8c, 0xd0, 0x19, 0x89, 0x24, 0x17, 0x81, 0xab, 0x2f, 0x1c, 0x70, 0x28, 0x80, 0xe6,
	0x4d, 0x2c, 0x05, 0x91, 0x2c, 0xa8, 0xe9, 0x70, 0x01, 0xc3, 0xef, 0xa1, 0x81, 0xaf, 0xe7, 0xc9,
	0x8a, 0x23, 0x1f, 0xdc, 0x0d, 0x89, 0xb4, 0x7c, 0x07, 0xab, 0x23, 0x42, 0x50, 0x93, 0xf1, 0x86,
	0x69, 0xc9, 0x16, 0xd6, 0x67, 0xc5, 0x89, 0x2c, 0x8b, 0xb5, 0x4a, 0x1d, 0xeb, 0xb3, 0xca, 0xbe,
	0xe6, 0x98, 0x2c, 0xde, 0x60, 0x9d, 0xdd, 0xc1, 0x05, 0x0c, 0x7f, 0x85, 0xc6, 0xd2, 0x64, 0xef,
	0x43, 0x6b, 0x25, 0xd8, 0x4f, 0xb7, 0x2c, 0x89, 0x72, 0xad, 0xe1, 0xe2, 0x3d, 0x81, 0x86, 0xe0,
	0x51, 0xdb, 0x0d, 0xad, 0xd6, 0x1e, 0x75, 0xce, 0x49, 0x76, 0x5e, 0x74, 0x08, 0xef, 0xa2, 0xaa,
	0x4a, 0x42, 0x8d, 0x49, 0x0f, 0xab, 0x23, 0xea, 0x81, 0x17, 0x71, 0xca, 0x70, 0x61, 0xae, 0x85,
	0x77, 0x38, 0xa4, 0x80, 0xbe, 0xe1, 0x71, 0x82, 0x95, 0x4e, 0x26, 0xed, 0x8f, 0xea, 0x77, 0xfa,
	0x21, 0xbf, 0x24, 0xf9, 0x9a, 0x13, 0x6a, 0x0d, 0x97, 0x18, 0xe5, 0x87, 0xb2, 0xbb, 0x31, 0xa5,
	0x42, 0x17, 0xd3, 0xc1, 0x05, 0x44, 0x4f, 0xa1, 0x9e, 0x30, 0x39, 0x9f, 0x6a, 0xfd, 0x0e, 0x36,
	0x20, 0xfc, 0xbd, 0x0a, 0x67, 0x07, 0x32, 0x59, 0xca, 0x93, 0x8c, 0xfd, 0x17, 0x9d, 0xe4, 0xfe,
	0xc7, 0xc5, 0x6b, 0x96, 0x17, 0x3a, 0x16, 0xaa, 0x88, 0xd8, 0x4e, 0xd9, 0x9a, 0xe4, 0x76, 0x9c,
	0x05, 0x44, 0x03, 0x68, 0x8b, 0xed, 0x8b, 0x29, 0x7e, 0xbb, 0x5a, 0x65, 0x4c, 0xda, 0x69, 0x96,
	0x29, 0xd4, 0x85, 0x46, 0x34, 0xfb, 0x36, 0xce, 0x64, 0x50, 0x1f, 0xb8, 0xc3, 0x13, 0x6c, 0x91,
	0xea, 0xb1, 0xd8, 0xbe, 0x8b, 0x13, 0xca, 0xef, 0x83, 0xc6, 0xc0, 0x19, 0x9e, 0x9a, 0x1e, 0xe3,
	0x6b, 0xc3, 0xe1, 0x5d, 0x54, 0xb9, 0x14, 0xdb, 0xd1, 0x14, 0x07, 0x4d, 0x9d, 0xdd, 0x00, 0x35,
	0x41, 0xc1, 0xd6, 0x64, 0x3b, 0x9b, 0x24, 0x32, 0xf0, 0x74, 0xff, 0xf7, 0x44, 0xf8, 0xa7, 0x03,
	0x67, 0x5f, 0x93, 0x84, 0xae, 0x99, 0x1a, 0xda, 0x55, 0x5a, 0xf4, 0xba, 0x0b, 0x0d, 0xca, 0xee,
	0x2e, 0xae, 0xe6, 0xd6, 0xbf, 0x45, 0x8a, 0x27, 0x69, 0xaa, 0x78, 0x63, 0xdd, 0x22, 0xb5, 0x5f,
	0x2b, 0x25, 0x60, 0x6c, 0xeb, 0xb3, 0xaa, 0x67, 0x75, 0xc9, 0x45, 0xe1, 0xd6, 0x00, 0x75, 0x53,
	0x6d, 0x45, 0x50, 0xd7, 0xff, 0xd7, 0x67, 0x14, 0x42, 0x43, 0x6e, 0xd5, 0xbe, 0x69, 0x87, 0xed,
	0x11, 0x28, 0x87, 0x66, 0x03, 0xb1, 0x8d, 0xa8, 0x3b, 0xc2, 0xdc, 0x69, 0x0e, 0xdc, 0xe2, 0x0e,
	0xb6, 0x77, 0x4c, 0x24, 0xfc, 0xcd, 0x01, 0xf4, 0x8a, 0x49, 0x65, 0x65, 0xca, 0xef, 0x93, 0xc7,
	0x9a, 0xf9, 0x12, 0x4e, 0x37, 0x64, 0x6b, 0xc7, 0xbd, 0x88, 0x7f, 0x66, 0xd6, 0xd6, 0x3f, 0xd8,
	0x9d, 0xe9, 0xda, 0xde, 0x74, 0x98, 0xc3, 0xd9, 0x41, 0x05, 0x76, 0xa7, 0x0a, 0xd7, 0x4e, 0xc9,
	0x75, 0x1f, 0x5a, 0x11, 0x4f, 0x56, 0xb1, 0xd8, 0x30, 0xaa, 0x2b, 0xf0, 0xf0, 0x9e, 0xd8, 0x77,
	0xcf, 0x2d, 0x77, 0xaf, 0x07, 0xde, 0x86, 0x0b, 0x3d, 0x2c, 0x2d, 0xeb, 0xe1, 0x1d, 0x0e, 0xbb,
	0xf0, 0xf4, 0x70, 0x94, 0x46, 0x3b, 0xfc, 0x01, 0x82, 0x3d, 0xaf, 0xaa, 0x1a, 0x4f, 0x5e, 0xff,
	0x8f, 0x73, 0x0e, 0x3f, 0x85, 0x4f, 0x1e, 0xc8, 0x6f, 0xc5, 0x7f, 0x01, 0x64, 0x82, 0x17, 0x42,
	0x70, 0xf1, 0x58, 0xd9, 0x2f, 0xa0, 0x26, 0xf3, 0xd4, 0xcc, 0xe1, 0x74, 0x74, 0xa2, 0x46, 0xaf,
	0xf3, 0x2d, 0xf3, 0x94, 0x61, 0x1d, 0x52, 0xfd, 0x62, 0x8a, 0xb2, 0x2f, 0x13, 0x03, 0xc2, 0x8f,
	0x8b, 0xf5, 0xb6, 0xf2, 0xa6, 0xaa, 0xe7, 0x7d, 0xf0, 0x8a, 0x07, 0x08, 0x35, 0xc1, 0xc5, 0xd7,
	0x2f, 0xfc, 0x8a, 0x39, 0x8c, 0x7c, 0xe7, 0xf9, 0x05, 0xb4, 0x76, 0xd9, 0x51, 0x1b, 0x9a, 0xaf,
	0x58, 0xc2, 0x44, 0x1c, 0xf9, 0x15, 0xe4, 0x41, 0xed, 0xed, 0x72, 0x3c, 0xf6, 0x1d, 0xe4, 0x43,
	0x67, 0x3a, 0x5e, 0x8e, 0xdf, 0x5f, 0x5d, 0xbe, 0x9f, 0x4d, 0xde, 0x2c, 0xfd, 0x2a, 0xfa, 0x08,
	0xda, 0x05, 0xf3, 0xdd, 0x7c, 0xe2, 0xbb, 0xa3, 0xbf, 0xaa, 0xf0, 0x64, 0x9c, 0xa6, 0xeb, 0x38,
	0xd2, 0xdf, 0x83, 0x05, 0x13, 0x77, 0x4c, 0xa0, 0x97, 0xd0, 0x2e, 0xbd, 0x74, 0x50, 0x57, 0x79,
	0x39, 0x7e, 0xd9, 0xf5, 0x9e, 0x1d, 0xf1, 0xb6, 0xa1, 0x15, 0x34, 0x81, 0x4e, 0x79, 0xce, 0x48,
	0x5f, 0x7d, 0xe0, 0x21, 0xee, 0x05, 0xc7, 0x81, 0x5d, 0x92, 0x97, 0xd0, 0x2e, 0xed, 0xa9, 0x29,
	0xe3, 0xf8, 0xd1, 0x31, 0x65, 0x3c, 0xb0, 0xd0, 0x61, 0x05, 0x61, 0x78, 0x72, 0x34, 0x76, 0xd4,
	0x3f, 0x94, 0x3c, 0xdc, 0xb6, 0xde, 0x67, 0xff, 0x12, 0x2d, 0x57, 0x55, 0x1a, 0x97, 0xa9, 0xea,
	0x78, 0x7d, 0x7a, 0xcf, 0x8e, 0xf8, 0x22, 0xc3, 0x4d, 0x43, 0x7f, 0xac, 0xbf, 0xfa, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x81, 0x7b, 0x7e, 0xb9, 0xb8, 0x07, 0x00, 0x00,
}