// Copyright 2022-2023 Tigris Data, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: server/v1/auth.proto

package api

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
	Auth_GetAccessToken_FullMethodName    = "/tigrisdata.auth.v1.Auth/GetAccessToken"
	Auth_CreateInvitations_FullMethodName = "/tigrisdata.auth.v1.Auth/CreateInvitations"
	Auth_DeleteInvitations_FullMethodName = "/tigrisdata.auth.v1.Auth/DeleteInvitations"
	Auth_ListInvitations_FullMethodName   = "/tigrisdata.auth.v1.Auth/ListInvitations"
	Auth_VerifyInvitation_FullMethodName  = "/tigrisdata.auth.v1.Auth/VerifyInvitation"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	// Endpoint for receiving access token from Tigris Server. The endpoint
	// requires Grant Type(`grant_type`) which has two possible values
	// <i>"REFRESH_TOKEN"</i> or <i>"CLIENT_CREDENTIALS"</i> based on which either
	// Refresh token(`refresh_token`) needs to be set or client
	// credentials(`client_id`, `client_secret`).
	GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*GetAccessTokenResponse, error)
	// Creates invitations to the namespace
	CreateInvitations(ctx context.Context, in *CreateInvitationsRequest, opts ...grpc.CallOption) (*CreateInvitationsResponse, error)
	// Deletes invitations to the namespace
	DeleteInvitations(ctx context.Context, in *DeleteInvitationsRequest, opts ...grpc.CallOption) (*DeleteInvitationsResponse, error)
	// Lists all the invitations to the namespace
	ListInvitations(ctx context.Context, in *ListInvitationsRequest, opts ...grpc.CallOption) (*ListInvitationsResponse, error)
	// Verify invitation
	VerifyInvitation(ctx context.Context, in *VerifyInvitationRequest, opts ...grpc.CallOption) (*VerifyInvitationResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*GetAccessTokenResponse, error) {
	out := new(GetAccessTokenResponse)
	err := c.cc.Invoke(ctx, Auth_GetAccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateInvitations(ctx context.Context, in *CreateInvitationsRequest, opts ...grpc.CallOption) (*CreateInvitationsResponse, error) {
	out := new(CreateInvitationsResponse)
	err := c.cc.Invoke(ctx, Auth_CreateInvitations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteInvitations(ctx context.Context, in *DeleteInvitationsRequest, opts ...grpc.CallOption) (*DeleteInvitationsResponse, error) {
	out := new(DeleteInvitationsResponse)
	err := c.cc.Invoke(ctx, Auth_DeleteInvitations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ListInvitations(ctx context.Context, in *ListInvitationsRequest, opts ...grpc.CallOption) (*ListInvitationsResponse, error) {
	out := new(ListInvitationsResponse)
	err := c.cc.Invoke(ctx, Auth_ListInvitations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) VerifyInvitation(ctx context.Context, in *VerifyInvitationRequest, opts ...grpc.CallOption) (*VerifyInvitationResponse, error) {
	out := new(VerifyInvitationResponse)
	err := c.cc.Invoke(ctx, Auth_VerifyInvitation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations should embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	// Endpoint for receiving access token from Tigris Server. The endpoint
	// requires Grant Type(`grant_type`) which has two possible values
	// <i>"REFRESH_TOKEN"</i> or <i>"CLIENT_CREDENTIALS"</i> based on which either
	// Refresh token(`refresh_token`) needs to be set or client
	// credentials(`client_id`, `client_secret`).
	GetAccessToken(context.Context, *GetAccessTokenRequest) (*GetAccessTokenResponse, error)
	// Creates invitations to the namespace
	CreateInvitations(context.Context, *CreateInvitationsRequest) (*CreateInvitationsResponse, error)
	// Deletes invitations to the namespace
	DeleteInvitations(context.Context, *DeleteInvitationsRequest) (*DeleteInvitationsResponse, error)
	// Lists all the invitations to the namespace
	ListInvitations(context.Context, *ListInvitationsRequest) (*ListInvitationsResponse, error)
	// Verify invitation
	VerifyInvitation(context.Context, *VerifyInvitationRequest) (*VerifyInvitationResponse, error)
}

// UnimplementedAuthServer should be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) GetAccessToken(context.Context, *GetAccessTokenRequest) (*GetAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
func (UnimplementedAuthServer) CreateInvitations(context.Context, *CreateInvitationsRequest) (*CreateInvitationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvitations not implemented")
}
func (UnimplementedAuthServer) DeleteInvitations(context.Context, *DeleteInvitationsRequest) (*DeleteInvitationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInvitations not implemented")
}
func (UnimplementedAuthServer) ListInvitations(context.Context, *ListInvitationsRequest) (*ListInvitationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInvitations not implemented")
}
func (UnimplementedAuthServer) VerifyInvitation(context.Context, *VerifyInvitationRequest) (*VerifyInvitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyInvitation not implemented")
}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_GetAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetAccessToken(ctx, req.(*GetAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CreateInvitations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvitationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateInvitations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_CreateInvitations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateInvitations(ctx, req.(*CreateInvitationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteInvitations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInvitationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteInvitations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_DeleteInvitations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteInvitations(ctx, req.(*DeleteInvitationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ListInvitations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInvitationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ListInvitations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ListInvitations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ListInvitations(ctx, req.(*ListInvitationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_VerifyInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).VerifyInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_VerifyInvitation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).VerifyInvitation(ctx, req.(*VerifyInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tigrisdata.auth.v1.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccessToken",
			Handler:    _Auth_GetAccessToken_Handler,
		},
		{
			MethodName: "CreateInvitations",
			Handler:    _Auth_CreateInvitations_Handler,
		},
		{
			MethodName: "DeleteInvitations",
			Handler:    _Auth_DeleteInvitations_Handler,
		},
		{
			MethodName: "ListInvitations",
			Handler:    _Auth_ListInvitations_Handler,
		},
		{
			MethodName: "VerifyInvitation",
			Handler:    _Auth_VerifyInvitation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/v1/auth.proto",
}
