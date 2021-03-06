// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apiv1

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

// DomainAPIClient is the client API for DomainAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DomainAPIClient interface {
	// RegisterDomain creates a new domain which can be used as a container for all resources.  Domain is a top level
	// entity within Cadence, used as a container for all resources like workflow executions, task lists, etc.  Domain
	// acts as a sandbox and provides isolation for all resources within the domain.  All resources belongs to exactly one
	// domain.
	RegisterDomain(ctx context.Context, in *RegisterDomainRequest, opts ...grpc.CallOption) (*RegisterDomainResponse, error)
	// DescribeDomain returns the information and configuration for a registered domain.
	DescribeDomain(ctx context.Context, in *DescribeDomainRequest, opts ...grpc.CallOption) (*DescribeDomainResponse, error)
	// ListDomains returns the information and configuration for all domains.
	ListDomains(ctx context.Context, in *ListDomainsRequest, opts ...grpc.CallOption) (*ListDomainsResponse, error)
	// UpdateDomain is used to update the information and configuration for a registered domain.
	UpdateDomain(ctx context.Context, in *UpdateDomainRequest, opts ...grpc.CallOption) (*UpdateDomainResponse, error)
	// DeprecateDomain us used to update status of a registered domain to DEPRECATED.  Once the domain is deprecated
	// it cannot be used to start new workflow executions.  Existing workflow executions will continue to run on
	// deprecated domains.
	DeprecateDomain(ctx context.Context, in *DeprecateDomainRequest, opts ...grpc.CallOption) (*DeprecateDomainResponse, error)
}

type domainAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewDomainAPIClient(cc grpc.ClientConnInterface) DomainAPIClient {
	return &domainAPIClient{cc}
}

func (c *domainAPIClient) RegisterDomain(ctx context.Context, in *RegisterDomainRequest, opts ...grpc.CallOption) (*RegisterDomainResponse, error) {
	out := new(RegisterDomainResponse)
	err := c.cc.Invoke(ctx, "/uber.cadence.api.v1.DomainAPI/RegisterDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainAPIClient) DescribeDomain(ctx context.Context, in *DescribeDomainRequest, opts ...grpc.CallOption) (*DescribeDomainResponse, error) {
	out := new(DescribeDomainResponse)
	err := c.cc.Invoke(ctx, "/uber.cadence.api.v1.DomainAPI/DescribeDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainAPIClient) ListDomains(ctx context.Context, in *ListDomainsRequest, opts ...grpc.CallOption) (*ListDomainsResponse, error) {
	out := new(ListDomainsResponse)
	err := c.cc.Invoke(ctx, "/uber.cadence.api.v1.DomainAPI/ListDomains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainAPIClient) UpdateDomain(ctx context.Context, in *UpdateDomainRequest, opts ...grpc.CallOption) (*UpdateDomainResponse, error) {
	out := new(UpdateDomainResponse)
	err := c.cc.Invoke(ctx, "/uber.cadence.api.v1.DomainAPI/UpdateDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainAPIClient) DeprecateDomain(ctx context.Context, in *DeprecateDomainRequest, opts ...grpc.CallOption) (*DeprecateDomainResponse, error) {
	out := new(DeprecateDomainResponse)
	err := c.cc.Invoke(ctx, "/uber.cadence.api.v1.DomainAPI/DeprecateDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomainAPIServer is the server API for DomainAPI service.
// All implementations must embed UnimplementedDomainAPIServer
// for forward compatibility
type DomainAPIServer interface {
	// RegisterDomain creates a new domain which can be used as a container for all resources.  Domain is a top level
	// entity within Cadence, used as a container for all resources like workflow executions, task lists, etc.  Domain
	// acts as a sandbox and provides isolation for all resources within the domain.  All resources belongs to exactly one
	// domain.
	RegisterDomain(context.Context, *RegisterDomainRequest) (*RegisterDomainResponse, error)
	// DescribeDomain returns the information and configuration for a registered domain.
	DescribeDomain(context.Context, *DescribeDomainRequest) (*DescribeDomainResponse, error)
	// ListDomains returns the information and configuration for all domains.
	ListDomains(context.Context, *ListDomainsRequest) (*ListDomainsResponse, error)
	// UpdateDomain is used to update the information and configuration for a registered domain.
	UpdateDomain(context.Context, *UpdateDomainRequest) (*UpdateDomainResponse, error)
	// DeprecateDomain us used to update status of a registered domain to DEPRECATED.  Once the domain is deprecated
	// it cannot be used to start new workflow executions.  Existing workflow executions will continue to run on
	// deprecated domains.
	DeprecateDomain(context.Context, *DeprecateDomainRequest) (*DeprecateDomainResponse, error)
	mustEmbedUnimplementedDomainAPIServer()
}

// UnimplementedDomainAPIServer must be embedded to have forward compatible implementations.
type UnimplementedDomainAPIServer struct {
}

func (UnimplementedDomainAPIServer) RegisterDomain(context.Context, *RegisterDomainRequest) (*RegisterDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDomain not implemented")
}
func (UnimplementedDomainAPIServer) DescribeDomain(context.Context, *DescribeDomainRequest) (*DescribeDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeDomain not implemented")
}
func (UnimplementedDomainAPIServer) ListDomains(context.Context, *ListDomainsRequest) (*ListDomainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDomains not implemented")
}
func (UnimplementedDomainAPIServer) UpdateDomain(context.Context, *UpdateDomainRequest) (*UpdateDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDomain not implemented")
}
func (UnimplementedDomainAPIServer) DeprecateDomain(context.Context, *DeprecateDomainRequest) (*DeprecateDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeprecateDomain not implemented")
}
func (UnimplementedDomainAPIServer) mustEmbedUnimplementedDomainAPIServer() {}

// UnsafeDomainAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DomainAPIServer will
// result in compilation errors.
type UnsafeDomainAPIServer interface {
	mustEmbedUnimplementedDomainAPIServer()
}

func RegisterDomainAPIServer(s grpc.ServiceRegistrar, srv DomainAPIServer) {
	s.RegisterService(&DomainAPI_ServiceDesc, srv)
}

func _DomainAPI_RegisterDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainAPIServer).RegisterDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.cadence.api.v1.DomainAPI/RegisterDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainAPIServer).RegisterDomain(ctx, req.(*RegisterDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainAPI_DescribeDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainAPIServer).DescribeDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.cadence.api.v1.DomainAPI/DescribeDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainAPIServer).DescribeDomain(ctx, req.(*DescribeDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainAPI_ListDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainAPIServer).ListDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.cadence.api.v1.DomainAPI/ListDomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainAPIServer).ListDomains(ctx, req.(*ListDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainAPI_UpdateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainAPIServer).UpdateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.cadence.api.v1.DomainAPI/UpdateDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainAPIServer).UpdateDomain(ctx, req.(*UpdateDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainAPI_DeprecateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeprecateDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainAPIServer).DeprecateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.cadence.api.v1.DomainAPI/DeprecateDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainAPIServer).DeprecateDomain(ctx, req.(*DeprecateDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DomainAPI_ServiceDesc is the grpc.ServiceDesc for DomainAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DomainAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "uber.cadence.api.v1.DomainAPI",
	HandlerType: (*DomainAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDomain",
			Handler:    _DomainAPI_RegisterDomain_Handler,
		},
		{
			MethodName: "DescribeDomain",
			Handler:    _DomainAPI_DescribeDomain_Handler,
		},
		{
			MethodName: "ListDomains",
			Handler:    _DomainAPI_ListDomains_Handler,
		},
		{
			MethodName: "UpdateDomain",
			Handler:    _DomainAPI_UpdateDomain_Handler,
		},
		{
			MethodName: "DeprecateDomain",
			Handler:    _DomainAPI_DeprecateDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uber/cadence/api/v1/service_domain.proto",
}
