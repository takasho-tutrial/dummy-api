// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: dummy-api.proto

package dummy_api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DataType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensorName string `protobuf:"bytes,1,opt,name=sensor_name,json=sensorName,proto3" json:"sensor_name,omitempty"`
}

func (x *DataType) Reset() {
	*x = DataType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dummy_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType) ProtoMessage() {}

func (x *DataType) ProtoReflect() protoreflect.Message {
	mi := &file_dummy_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType.ProtoReflect.Descriptor instead.
func (*DataType) Descriptor() ([]byte, []int) {
	return file_dummy_api_proto_rawDescGZIP(), []int{0}
}

func (x *DataType) GetSensorName() string {
	if x != nil {
		return x.SensorName
	}
	return ""
}

type SensingData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensorName   string  `protobuf:"bytes,1,opt,name=sensor_name,json=sensorName,proto3" json:"sensor_name,omitempty"`
	SensorValues []int32 `protobuf:"varint,2,rep,packed,name=sensor_values,json=sensorValues,proto3" json:"sensor_values,omitempty"`
}

func (x *SensingData) Reset() {
	*x = SensingData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dummy_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SensingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensingData) ProtoMessage() {}

func (x *SensingData) ProtoReflect() protoreflect.Message {
	mi := &file_dummy_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensingData.ProtoReflect.Descriptor instead.
func (*SensingData) Descriptor() ([]byte, []int) {
	return file_dummy_api_proto_rawDescGZIP(), []int{1}
}

func (x *SensingData) GetSensorName() string {
	if x != nil {
		return x.SensorName
	}
	return ""
}

func (x *SensingData) GetSensorValues() []int32 {
	if x != nil {
		return x.SensorValues
	}
	return nil
}

type StatsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensorName string  `protobuf:"bytes,1,opt,name=sensor_name,json=sensorName,proto3" json:"sensor_name,omitempty"`
	StatsInfo  float32 `protobuf:"fixed32,2,opt,name=stats_info,json=statsInfo,proto3" json:"stats_info,omitempty"`
}

func (x *StatsData) Reset() {
	*x = StatsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dummy_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsData) ProtoMessage() {}

func (x *StatsData) ProtoReflect() protoreflect.Message {
	mi := &file_dummy_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsData.ProtoReflect.Descriptor instead.
func (*StatsData) Descriptor() ([]byte, []int) {
	return file_dummy_api_proto_rawDescGZIP(), []int{2}
}

func (x *StatsData) GetSensorName() string {
	if x != nil {
		return x.SensorName
	}
	return ""
}

func (x *StatsData) GetStatsInfo() float32 {
	if x != nil {
		return x.StatsInfo
	}
	return 0
}

var File_dummy_api_proto protoreflect.FileDescriptor

var file_dummy_api_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x41, 0x70, 0x69, 0x22, 0x2b, 0x0a, 0x08, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x53, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x73,
	0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0c, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x4b, 0x0a,
	0x09, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xfd, 0x01, 0x0a, 0x0f, 0x44,
	0x75, 0x6d, 0x6d, 0x79, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x2e, 0x64, 0x75, 0x6d, 0x6d,
	0x79, 0x41, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x15, 0x2e,
	0x64, 0x75, 0x6d, 0x6d, 0x79, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x6e, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x41, 0x70, 0x69, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x13, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79,
	0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x28, 0x01, 0x12,
	0x39, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x2e,
	0x64, 0x75, 0x6d, 0x6d, 0x79, 0x41, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x1a, 0x15, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e,
	0x73, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x30, 0x01, 0x12, 0x3e, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x71, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x2e,
	0x64, 0x75, 0x6d, 0x6d, 0x79, 0x41, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x1a, 0x13, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x28, 0x01, 0x30, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x6b, 0x61, 0x73, 0x68, 0x6f,
	0x2d, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2d,
	0x61, 0x70, 0x69, 0x3b, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dummy_api_proto_rawDescOnce sync.Once
	file_dummy_api_proto_rawDescData = file_dummy_api_proto_rawDesc
)

func file_dummy_api_proto_rawDescGZIP() []byte {
	file_dummy_api_proto_rawDescOnce.Do(func() {
		file_dummy_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_dummy_api_proto_rawDescData)
	})
	return file_dummy_api_proto_rawDescData
}

var file_dummy_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dummy_api_proto_goTypes = []interface{}{
	(*DataType)(nil),    // 0: dummyApi.DataType
	(*SensingData)(nil), // 1: dummyApi.SensingData
	(*StatsData)(nil),   // 2: dummyApi.StatsData
}
var file_dummy_api_proto_depIdxs = []int32{
	0, // 0: dummyApi.DummyApiService.GetData:input_type -> dummyApi.DataType
	0, // 1: dummyApi.DummyApiService.GetStatsData:input_type -> dummyApi.DataType
	0, // 2: dummyApi.DummyApiService.GetSeqData:input_type -> dummyApi.DataType
	0, // 3: dummyApi.DummyApiService.GetSeqStatsData:input_type -> dummyApi.DataType
	1, // 4: dummyApi.DummyApiService.GetData:output_type -> dummyApi.SensingData
	2, // 5: dummyApi.DummyApiService.GetStatsData:output_type -> dummyApi.StatsData
	1, // 6: dummyApi.DummyApiService.GetSeqData:output_type -> dummyApi.SensingData
	2, // 7: dummyApi.DummyApiService.GetSeqStatsData:output_type -> dummyApi.StatsData
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dummy_api_proto_init() }
func file_dummy_api_proto_init() {
	if File_dummy_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dummy_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dummy_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SensingData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dummy_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dummy_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dummy_api_proto_goTypes,
		DependencyIndexes: file_dummy_api_proto_depIdxs,
		MessageInfos:      file_dummy_api_proto_msgTypes,
	}.Build()
	File_dummy_api_proto = out.File
	file_dummy_api_proto_rawDesc = nil
	file_dummy_api_proto_goTypes = nil
	file_dummy_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DummyApiServiceClient is the client API for DummyApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DummyApiServiceClient interface {
	GetData(ctx context.Context, in *DataType, opts ...grpc.CallOption) (*SensingData, error)
	GetStatsData(ctx context.Context, opts ...grpc.CallOption) (DummyApiService_GetStatsDataClient, error)
	GetSeqData(ctx context.Context, in *DataType, opts ...grpc.CallOption) (DummyApiService_GetSeqDataClient, error)
	GetSeqStatsData(ctx context.Context, opts ...grpc.CallOption) (DummyApiService_GetSeqStatsDataClient, error)
}

type dummyApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDummyApiServiceClient(cc grpc.ClientConnInterface) DummyApiServiceClient {
	return &dummyApiServiceClient{cc}
}

func (c *dummyApiServiceClient) GetData(ctx context.Context, in *DataType, opts ...grpc.CallOption) (*SensingData, error) {
	out := new(SensingData)
	err := c.cc.Invoke(ctx, "/dummyApi.DummyApiService/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dummyApiServiceClient) GetStatsData(ctx context.Context, opts ...grpc.CallOption) (DummyApiService_GetStatsDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DummyApiService_serviceDesc.Streams[0], "/dummyApi.DummyApiService/GetStatsData", opts...)
	if err != nil {
		return nil, err
	}
	x := &dummyApiServiceGetStatsDataClient{stream}
	return x, nil
}

type DummyApiService_GetStatsDataClient interface {
	Send(*DataType) error
	CloseAndRecv() (*StatsData, error)
	grpc.ClientStream
}

type dummyApiServiceGetStatsDataClient struct {
	grpc.ClientStream
}

func (x *dummyApiServiceGetStatsDataClient) Send(m *DataType) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dummyApiServiceGetStatsDataClient) CloseAndRecv() (*StatsData, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StatsData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dummyApiServiceClient) GetSeqData(ctx context.Context, in *DataType, opts ...grpc.CallOption) (DummyApiService_GetSeqDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DummyApiService_serviceDesc.Streams[1], "/dummyApi.DummyApiService/GetSeqData", opts...)
	if err != nil {
		return nil, err
	}
	x := &dummyApiServiceGetSeqDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DummyApiService_GetSeqDataClient interface {
	Recv() (*SensingData, error)
	grpc.ClientStream
}

type dummyApiServiceGetSeqDataClient struct {
	grpc.ClientStream
}

func (x *dummyApiServiceGetSeqDataClient) Recv() (*SensingData, error) {
	m := new(SensingData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dummyApiServiceClient) GetSeqStatsData(ctx context.Context, opts ...grpc.CallOption) (DummyApiService_GetSeqStatsDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DummyApiService_serviceDesc.Streams[2], "/dummyApi.DummyApiService/GetSeqStatsData", opts...)
	if err != nil {
		return nil, err
	}
	x := &dummyApiServiceGetSeqStatsDataClient{stream}
	return x, nil
}

type DummyApiService_GetSeqStatsDataClient interface {
	Send(*DataType) error
	Recv() (*StatsData, error)
	grpc.ClientStream
}

type dummyApiServiceGetSeqStatsDataClient struct {
	grpc.ClientStream
}

func (x *dummyApiServiceGetSeqStatsDataClient) Send(m *DataType) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dummyApiServiceGetSeqStatsDataClient) Recv() (*StatsData, error) {
	m := new(StatsData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DummyApiServiceServer is the server API for DummyApiService service.
type DummyApiServiceServer interface {
	GetData(context.Context, *DataType) (*SensingData, error)
	GetStatsData(DummyApiService_GetStatsDataServer) error
	GetSeqData(*DataType, DummyApiService_GetSeqDataServer) error
	GetSeqStatsData(DummyApiService_GetSeqStatsDataServer) error
}

// UnimplementedDummyApiServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDummyApiServiceServer struct {
}

func (*UnimplementedDummyApiServiceServer) GetData(context.Context, *DataType) (*SensingData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (*UnimplementedDummyApiServiceServer) GetStatsData(DummyApiService_GetStatsDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStatsData not implemented")
}
func (*UnimplementedDummyApiServiceServer) GetSeqData(*DataType, DummyApiService_GetSeqDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSeqData not implemented")
}
func (*UnimplementedDummyApiServiceServer) GetSeqStatsData(DummyApiService_GetSeqStatsDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSeqStatsData not implemented")
}

func RegisterDummyApiServiceServer(s *grpc.Server, srv DummyApiServiceServer) {
	s.RegisterService(&_DummyApiService_serviceDesc, srv)
}

func _DummyApiService_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DummyApiServiceServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dummyApi.DummyApiService/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DummyApiServiceServer).GetData(ctx, req.(*DataType))
	}
	return interceptor(ctx, in, info, handler)
}

func _DummyApiService_GetStatsData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DummyApiServiceServer).GetStatsData(&dummyApiServiceGetStatsDataServer{stream})
}

type DummyApiService_GetStatsDataServer interface {
	SendAndClose(*StatsData) error
	Recv() (*DataType, error)
	grpc.ServerStream
}

type dummyApiServiceGetStatsDataServer struct {
	grpc.ServerStream
}

func (x *dummyApiServiceGetStatsDataServer) SendAndClose(m *StatsData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dummyApiServiceGetStatsDataServer) Recv() (*DataType, error) {
	m := new(DataType)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DummyApiService_GetSeqData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DataType)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DummyApiServiceServer).GetSeqData(m, &dummyApiServiceGetSeqDataServer{stream})
}

type DummyApiService_GetSeqDataServer interface {
	Send(*SensingData) error
	grpc.ServerStream
}

type dummyApiServiceGetSeqDataServer struct {
	grpc.ServerStream
}

func (x *dummyApiServiceGetSeqDataServer) Send(m *SensingData) error {
	return x.ServerStream.SendMsg(m)
}

func _DummyApiService_GetSeqStatsData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DummyApiServiceServer).GetSeqStatsData(&dummyApiServiceGetSeqStatsDataServer{stream})
}

type DummyApiService_GetSeqStatsDataServer interface {
	Send(*StatsData) error
	Recv() (*DataType, error)
	grpc.ServerStream
}

type dummyApiServiceGetSeqStatsDataServer struct {
	grpc.ServerStream
}

func (x *dummyApiServiceGetSeqStatsDataServer) Send(m *StatsData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dummyApiServiceGetSeqStatsDataServer) Recv() (*DataType, error) {
	m := new(DataType)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DummyApiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dummyApi.DummyApiService",
	HandlerType: (*DummyApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _DummyApiService_GetData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStatsData",
			Handler:       _DummyApiService_GetStatsData_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetSeqData",
			Handler:       _DummyApiService_GetSeqData_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetSeqStatsData",
			Handler:       _DummyApiService_GetSeqStatsData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "dummy-api.proto",
}
