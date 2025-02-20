// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: hyle/hyle/zktx/v1/tx.proto

package zktxv1

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

const (
	Msg_ExecuteStateChange_FullMethodName = "/hyle.hyle.zktx.v1.Msg/ExecuteStateChange"
	Msg_VerifyProof_FullMethodName        = "/hyle.hyle.zktx.v1.Msg/VerifyProof"
	Msg_RegisterContract_FullMethodName   = "/hyle.hyle.zktx.v1.Msg/RegisterContract"
	Msg_UpdateParams_FullMethodName       = "/hyle.hyle.zktx.v1.Msg/UpdateParams"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// execute a zk-proven state change
	ExecuteStateChange(ctx context.Context, in *MsgExecuteStateChange, opts ...grpc.CallOption) (*MsgExecuteStateChangeResponse, error)
	// Only verify a ZK proof
	VerifyProof(ctx context.Context, in *MsgVerifyProof, opts ...grpc.CallOption) (*MsgVerifyProofResponse, error)
	// RegisterContract registers a contract
	RegisterContract(ctx context.Context, in *MsgRegisterContract, opts ...grpc.CallOption) (*MsgRegisterContractResponse, error)
	// UpdateParams updates the module parameters.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ExecuteStateChange(ctx context.Context, in *MsgExecuteStateChange, opts ...grpc.CallOption) (*MsgExecuteStateChangeResponse, error) {
	out := new(MsgExecuteStateChangeResponse)
	err := c.cc.Invoke(ctx, Msg_ExecuteStateChange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) VerifyProof(ctx context.Context, in *MsgVerifyProof, opts ...grpc.CallOption) (*MsgVerifyProofResponse, error) {
	out := new(MsgVerifyProofResponse)
	err := c.cc.Invoke(ctx, Msg_VerifyProof_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RegisterContract(ctx context.Context, in *MsgRegisterContract, opts ...grpc.CallOption) (*MsgRegisterContractResponse, error) {
	out := new(MsgRegisterContractResponse)
	err := c.cc.Invoke(ctx, Msg_RegisterContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// execute a zk-proven state change
	ExecuteStateChange(context.Context, *MsgExecuteStateChange) (*MsgExecuteStateChangeResponse, error)
	// Only verify a ZK proof
	VerifyProof(context.Context, *MsgVerifyProof) (*MsgVerifyProofResponse, error)
	// RegisterContract registers a contract
	RegisterContract(context.Context, *MsgRegisterContract) (*MsgRegisterContractResponse, error)
	// UpdateParams updates the module parameters.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) ExecuteStateChange(context.Context, *MsgExecuteStateChange) (*MsgExecuteStateChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteStateChange not implemented")
}
func (UnimplementedMsgServer) VerifyProof(context.Context, *MsgVerifyProof) (*MsgVerifyProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyProof not implemented")
}
func (UnimplementedMsgServer) RegisterContract(context.Context, *MsgRegisterContract) (*MsgRegisterContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterContract not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_ExecuteStateChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExecuteStateChange)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExecuteStateChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ExecuteStateChange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExecuteStateChange(ctx, req.(*MsgExecuteStateChange))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_VerifyProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVerifyProof)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).VerifyProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_VerifyProof_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).VerifyProof(ctx, req.(*MsgVerifyProof))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RegisterContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RegisterContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterContract(ctx, req.(*MsgRegisterContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hyle.hyle.zktx.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteStateChange",
			Handler:    _Msg_ExecuteStateChange_Handler,
		},
		{
			MethodName: "VerifyProof",
			Handler:    _Msg_VerifyProof_Handler,
		},
		{
			MethodName: "RegisterContract",
			Handler:    _Msg_RegisterContract_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hyle/hyle/zktx/v1/tx.proto",
}
