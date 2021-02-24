// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EmrServiceClient is the client API for EmrService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmrServiceClient interface {
	PatientInfo(ctx context.Context, in *RequestCid, opts ...grpc.CallOption) (*InfoResponse, error)
	GetServices(ctx context.Context, in *RequestCid, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type emrServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmrServiceClient(cc grpc.ClientConnInterface) EmrServiceClient {
	return &emrServiceClient{cc}
}

func (c *emrServiceClient) PatientInfo(ctx context.Context, in *RequestCid, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/proto.EmrService/PatientInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emrServiceClient) GetServices(ctx context.Context, in *RequestCid, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.EmrService/GetServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmrServiceServer is the server API for EmrService service.
// All implementations must embed UnimplementedEmrServiceServer
// for forward compatibility
type EmrServiceServer interface {
	PatientInfo(context.Context, *RequestCid) (*InfoResponse, error)
	GetServices(context.Context, *RequestCid) (*ServiceResponse, error)
	mustEmbedUnimplementedEmrServiceServer()
}

// UnimplementedEmrServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmrServiceServer struct {
}

func (UnimplementedEmrServiceServer) PatientInfo(context.Context, *RequestCid) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatientInfo not implemented")
}
func (UnimplementedEmrServiceServer) GetServices(context.Context, *RequestCid) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (UnimplementedEmrServiceServer) mustEmbedUnimplementedEmrServiceServer() {}

// UnsafeEmrServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmrServiceServer will
// result in compilation errors.
type UnsafeEmrServiceServer interface {
	mustEmbedUnimplementedEmrServiceServer()
}

func RegisterEmrServiceServer(s grpc.ServiceRegistrar, srv EmrServiceServer) {
	s.RegisterService(&EmrService_ServiceDesc, srv)
}

func _EmrService_PatientInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmrServiceServer).PatientInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmrService/PatientInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmrServiceServer).PatientInfo(ctx, req.(*RequestCid))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmrService_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmrServiceServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmrService/GetServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmrServiceServer).GetServices(ctx, req.(*RequestCid))
	}
	return interceptor(ctx, in, info, handler)
}

// EmrService_ServiceDesc is the grpc.ServiceDesc for EmrService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmrService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EmrService",
	HandlerType: (*EmrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PatientInfo",
			Handler:    _EmrService_PatientInfo_Handler,
		},
		{
			MethodName: "GetServices",
			Handler:    _EmrService_GetServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hosxp.proto",
}

// DoctorServiceClient is the client API for DoctorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoctorServiceClient interface {
	DoctorList(ctx context.Context, in *RequestHospcode, opts ...grpc.CallOption) (*DoctorResponse, error)
}

type doctorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDoctorServiceClient(cc grpc.ClientConnInterface) DoctorServiceClient {
	return &doctorServiceClient{cc}
}

func (c *doctorServiceClient) DoctorList(ctx context.Context, in *RequestHospcode, opts ...grpc.CallOption) (*DoctorResponse, error) {
	out := new(DoctorResponse)
	err := c.cc.Invoke(ctx, "/proto.DoctorService/DoctorList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoctorServiceServer is the server API for DoctorService service.
// All implementations must embed UnimplementedDoctorServiceServer
// for forward compatibility
type DoctorServiceServer interface {
	DoctorList(context.Context, *RequestHospcode) (*DoctorResponse, error)
	mustEmbedUnimplementedDoctorServiceServer()
}

// UnimplementedDoctorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDoctorServiceServer struct {
}

func (UnimplementedDoctorServiceServer) DoctorList(context.Context, *RequestHospcode) (*DoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoctorList not implemented")
}
func (UnimplementedDoctorServiceServer) mustEmbedUnimplementedDoctorServiceServer() {}

// UnsafeDoctorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoctorServiceServer will
// result in compilation errors.
type UnsafeDoctorServiceServer interface {
	mustEmbedUnimplementedDoctorServiceServer()
}

func RegisterDoctorServiceServer(s grpc.ServiceRegistrar, srv DoctorServiceServer) {
	s.RegisterService(&DoctorService_ServiceDesc, srv)
}

func _DoctorService_DoctorList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestHospcode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).DoctorList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DoctorService/DoctorList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).DoctorList(ctx, req.(*RequestHospcode))
	}
	return interceptor(ctx, in, info, handler)
}

// DoctorService_ServiceDesc is the grpc.ServiceDesc for DoctorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoctorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DoctorService",
	HandlerType: (*DoctorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoctorList",
			Handler:    _DoctorService_DoctorList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hosxp.proto",
}
