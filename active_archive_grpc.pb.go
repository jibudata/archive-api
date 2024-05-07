// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: active_archive.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ActiveArchiveClient is the client API for ActiveArchive service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActiveArchiveClient interface {
	ListMediumInfo(ctx context.Context, in *DefaultResourceRequest, opts ...grpc.CallOption) (*MediumInfo, error)
	GetMediaInfo(ctx context.Context, in *LibraryManagerResourceKey, opts ...grpc.CallOption) (*MediaInfoReply, error)
	GetDrivesInfo(ctx context.Context, in *DefaultResourceRequest, opts ...grpc.CallOption) (*DrivesInfo, error)
	GetPoolsInfo(ctx context.Context, in *DefaultResourceRequest, opts ...grpc.CallOption) (*PoolsInfo, error)
	CreatePool(ctx context.Context, in *CreatePoolParams, opts ...grpc.CallOption) (*ReplyMessage, error)
	DeletePool(ctx context.Context, in *LibraryManagerResourceKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddMediaToPool(ctx context.Context, in *PoolAddRequest, opts ...grpc.CallOption) (*ReplyMessage, error)
	RemoveMediaFromPool(ctx context.Context, in *PoolRemoveRequest, opts ...grpc.CallOption) (*ReplyMessage, error)
	Migrate(ctx context.Context, in *MigrateRequest, opts ...grpc.CallOption) (*MigrationStatus, error)
	Recall(ctx context.Context, in *RecallRequest, opts ...grpc.CallOption) (*MigrationStatus, error)
	Retrieve(ctx context.Context, in *DefaultResourceRequest, opts ...grpc.CallOption) (*ReplyMessage, error)
	MigrateAsync(ctx context.Context, in *MigrateRequest, opts ...grpc.CallOption) (*MigrationStatus, error)
	RecallAsync(ctx context.Context, in *RecallRequest, opts ...grpc.CallOption) (*MigrationStatus, error)
	GetAsyncStatus(ctx context.Context, in *AsyncStatusRequest, opts ...grpc.CallOption) (*MigrationStatus, error)
	GetFileInfo(ctx context.Context, in *FileInfoRequest, opts ...grpc.CallOption) (*FileInfo, error)
	SearchFile(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	PrepareFileList(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*PrepareFileListResponse, error)
}

type activeArchiveClient struct {
	cc grpc.ClientConnInterface
}

func NewActiveArchiveClient(cc grpc.ClientConnInterface) ActiveArchiveClient {
	return &activeArchiveClient{cc}
}

func (c *activeArchiveClient) ListMediumInfo(ctx context.Context, in *DefaultResourceRequest, opts ...grpc.CallOption) (*MediumInfo, error) {
	out := new(MediumInfo)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/ListMediumInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) GetMediaInfo(ctx context.Context, in *LibraryManagerResourceKey, opts ...grpc.CallOption) (*MediaInfoReply, error) {
	out := new(MediaInfoReply)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/GetMediaInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) GetDrivesInfo(ctx context.Context, in *DefaultResourceRequest, opts ...grpc.CallOption) (*DrivesInfo, error) {
	out := new(DrivesInfo)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/GetDrivesInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) GetPoolsInfo(ctx context.Context, in *DefaultResourceRequest, opts ...grpc.CallOption) (*PoolsInfo, error) {
	out := new(PoolsInfo)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/GetPoolsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) CreatePool(ctx context.Context, in *CreatePoolParams, opts ...grpc.CallOption) (*ReplyMessage, error) {
	out := new(ReplyMessage)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/CreatePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) DeletePool(ctx context.Context, in *LibraryManagerResourceKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/DeletePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) AddMediaToPool(ctx context.Context, in *PoolAddRequest, opts ...grpc.CallOption) (*ReplyMessage, error) {
	out := new(ReplyMessage)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/AddMediaToPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) RemoveMediaFromPool(ctx context.Context, in *PoolRemoveRequest, opts ...grpc.CallOption) (*ReplyMessage, error) {
	out := new(ReplyMessage)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/RemoveMediaFromPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) Migrate(ctx context.Context, in *MigrateRequest, opts ...grpc.CallOption) (*MigrationStatus, error) {
	out := new(MigrationStatus)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/Migrate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) Recall(ctx context.Context, in *RecallRequest, opts ...grpc.CallOption) (*MigrationStatus, error) {
	out := new(MigrationStatus)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/Recall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) Retrieve(ctx context.Context, in *DefaultResourceRequest, opts ...grpc.CallOption) (*ReplyMessage, error) {
	out := new(ReplyMessage)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) MigrateAsync(ctx context.Context, in *MigrateRequest, opts ...grpc.CallOption) (*MigrationStatus, error) {
	out := new(MigrationStatus)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/MigrateAsync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) RecallAsync(ctx context.Context, in *RecallRequest, opts ...grpc.CallOption) (*MigrationStatus, error) {
	out := new(MigrationStatus)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/RecallAsync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) GetAsyncStatus(ctx context.Context, in *AsyncStatusRequest, opts ...grpc.CallOption) (*MigrationStatus, error) {
	out := new(MigrationStatus)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/GetAsyncStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) GetFileInfo(ctx context.Context, in *FileInfoRequest, opts ...grpc.CallOption) (*FileInfo, error) {
	out := new(FileInfo)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/GetFileInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) SearchFile(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/SearchFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeArchiveClient) PrepareFileList(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*PrepareFileListResponse, error) {
	out := new(PrepareFileListResponse)
	err := c.cc.Invoke(ctx, "/v1.ActiveArchive/PrepareFileList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActiveArchiveServer is the server API for ActiveArchive service.
// All implementations must embed UnimplementedActiveArchiveServer
// for forward compatibility
type ActiveArchiveServer interface {
	ListMediumInfo(context.Context, *DefaultResourceRequest) (*MediumInfo, error)
	GetMediaInfo(context.Context, *LibraryManagerResourceKey) (*MediaInfoReply, error)
	GetDrivesInfo(context.Context, *DefaultResourceRequest) (*DrivesInfo, error)
	GetPoolsInfo(context.Context, *DefaultResourceRequest) (*PoolsInfo, error)
	CreatePool(context.Context, *CreatePoolParams) (*ReplyMessage, error)
	DeletePool(context.Context, *LibraryManagerResourceKey) (*emptypb.Empty, error)
	AddMediaToPool(context.Context, *PoolAddRequest) (*ReplyMessage, error)
	RemoveMediaFromPool(context.Context, *PoolRemoveRequest) (*ReplyMessage, error)
	Migrate(context.Context, *MigrateRequest) (*MigrationStatus, error)
	Recall(context.Context, *RecallRequest) (*MigrationStatus, error)
	Retrieve(context.Context, *DefaultResourceRequest) (*ReplyMessage, error)
	MigrateAsync(context.Context, *MigrateRequest) (*MigrationStatus, error)
	RecallAsync(context.Context, *RecallRequest) (*MigrationStatus, error)
	GetAsyncStatus(context.Context, *AsyncStatusRequest) (*MigrationStatus, error)
	GetFileInfo(context.Context, *FileInfoRequest) (*FileInfo, error)
	SearchFile(context.Context, *SearchRequest) (*SearchResponse, error)
	PrepareFileList(context.Context, *SearchRequest) (*PrepareFileListResponse, error)
	mustEmbedUnimplementedActiveArchiveServer()
}

// UnimplementedActiveArchiveServer must be embedded to have forward compatible implementations.
type UnimplementedActiveArchiveServer struct {
}

func (UnimplementedActiveArchiveServer) ListMediumInfo(context.Context, *DefaultResourceRequest) (*MediumInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMediumInfo not implemented")
}
func (UnimplementedActiveArchiveServer) GetMediaInfo(context.Context, *LibraryManagerResourceKey) (*MediaInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMediaInfo not implemented")
}
func (UnimplementedActiveArchiveServer) GetDrivesInfo(context.Context, *DefaultResourceRequest) (*DrivesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDrivesInfo not implemented")
}
func (UnimplementedActiveArchiveServer) GetPoolsInfo(context.Context, *DefaultResourceRequest) (*PoolsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoolsInfo not implemented")
}
func (UnimplementedActiveArchiveServer) CreatePool(context.Context, *CreatePoolParams) (*ReplyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}
func (UnimplementedActiveArchiveServer) DeletePool(context.Context, *LibraryManagerResourceKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePool not implemented")
}
func (UnimplementedActiveArchiveServer) AddMediaToPool(context.Context, *PoolAddRequest) (*ReplyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMediaToPool not implemented")
}
func (UnimplementedActiveArchiveServer) RemoveMediaFromPool(context.Context, *PoolRemoveRequest) (*ReplyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMediaFromPool not implemented")
}
func (UnimplementedActiveArchiveServer) Migrate(context.Context, *MigrateRequest) (*MigrationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Migrate not implemented")
}
func (UnimplementedActiveArchiveServer) Recall(context.Context, *RecallRequest) (*MigrationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recall not implemented")
}
func (UnimplementedActiveArchiveServer) Retrieve(context.Context, *DefaultResourceRequest) (*ReplyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedActiveArchiveServer) MigrateAsync(context.Context, *MigrateRequest) (*MigrationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MigrateAsync not implemented")
}
func (UnimplementedActiveArchiveServer) RecallAsync(context.Context, *RecallRequest) (*MigrationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecallAsync not implemented")
}
func (UnimplementedActiveArchiveServer) GetAsyncStatus(context.Context, *AsyncStatusRequest) (*MigrationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAsyncStatus not implemented")
}
func (UnimplementedActiveArchiveServer) GetFileInfo(context.Context, *FileInfoRequest) (*FileInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileInfo not implemented")
}
func (UnimplementedActiveArchiveServer) SearchFile(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchFile not implemented")
}
func (UnimplementedActiveArchiveServer) PrepareFileList(context.Context, *SearchRequest) (*PrepareFileListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareFileList not implemented")
}
func (UnimplementedActiveArchiveServer) mustEmbedUnimplementedActiveArchiveServer() {}

// UnsafeActiveArchiveServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActiveArchiveServer will
// result in compilation errors.
type UnsafeActiveArchiveServer interface {
	mustEmbedUnimplementedActiveArchiveServer()
}

func RegisterActiveArchiveServer(s grpc.ServiceRegistrar, srv ActiveArchiveServer) {
	s.RegisterService(&ActiveArchive_ServiceDesc, srv)
}

func _ActiveArchive_ListMediumInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).ListMediumInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/ListMediumInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).ListMediumInfo(ctx, req.(*DefaultResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_GetMediaInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LibraryManagerResourceKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).GetMediaInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/GetMediaInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).GetMediaInfo(ctx, req.(*LibraryManagerResourceKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_GetDrivesInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).GetDrivesInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/GetDrivesInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).GetDrivesInfo(ctx, req.(*DefaultResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_GetPoolsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).GetPoolsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/GetPoolsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).GetPoolsInfo(ctx, req.(*DefaultResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePoolParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/CreatePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).CreatePool(ctx, req.(*CreatePoolParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_DeletePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LibraryManagerResourceKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).DeletePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/DeletePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).DeletePool(ctx, req.(*LibraryManagerResourceKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_AddMediaToPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PoolAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).AddMediaToPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/AddMediaToPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).AddMediaToPool(ctx, req.(*PoolAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_RemoveMediaFromPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PoolRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).RemoveMediaFromPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/RemoveMediaFromPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).RemoveMediaFromPool(ctx, req.(*PoolRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_Migrate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MigrateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).Migrate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/Migrate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).Migrate(ctx, req.(*MigrateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_Recall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).Recall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/Recall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).Recall(ctx, req.(*RecallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).Retrieve(ctx, req.(*DefaultResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_MigrateAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MigrateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).MigrateAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/MigrateAsync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).MigrateAsync(ctx, req.(*MigrateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_RecallAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).RecallAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/RecallAsync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).RecallAsync(ctx, req.(*RecallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_GetAsyncStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AsyncStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).GetAsyncStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/GetAsyncStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).GetAsyncStatus(ctx, req.(*AsyncStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_GetFileInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).GetFileInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/GetFileInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).GetFileInfo(ctx, req.(*FileInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_SearchFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).SearchFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/SearchFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).SearchFile(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveArchive_PrepareFileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveArchiveServer).PrepareFileList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ActiveArchive/PrepareFileList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveArchiveServer).PrepareFileList(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActiveArchive_ServiceDesc is the grpc.ServiceDesc for ActiveArchive service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActiveArchive_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ActiveArchive",
	HandlerType: (*ActiveArchiveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMediumInfo",
			Handler:    _ActiveArchive_ListMediumInfo_Handler,
		},
		{
			MethodName: "GetMediaInfo",
			Handler:    _ActiveArchive_GetMediaInfo_Handler,
		},
		{
			MethodName: "GetDrivesInfo",
			Handler:    _ActiveArchive_GetDrivesInfo_Handler,
		},
		{
			MethodName: "GetPoolsInfo",
			Handler:    _ActiveArchive_GetPoolsInfo_Handler,
		},
		{
			MethodName: "CreatePool",
			Handler:    _ActiveArchive_CreatePool_Handler,
		},
		{
			MethodName: "DeletePool",
			Handler:    _ActiveArchive_DeletePool_Handler,
		},
		{
			MethodName: "AddMediaToPool",
			Handler:    _ActiveArchive_AddMediaToPool_Handler,
		},
		{
			MethodName: "RemoveMediaFromPool",
			Handler:    _ActiveArchive_RemoveMediaFromPool_Handler,
		},
		{
			MethodName: "Migrate",
			Handler:    _ActiveArchive_Migrate_Handler,
		},
		{
			MethodName: "Recall",
			Handler:    _ActiveArchive_Recall_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _ActiveArchive_Retrieve_Handler,
		},
		{
			MethodName: "MigrateAsync",
			Handler:    _ActiveArchive_MigrateAsync_Handler,
		},
		{
			MethodName: "RecallAsync",
			Handler:    _ActiveArchive_RecallAsync_Handler,
		},
		{
			MethodName: "GetAsyncStatus",
			Handler:    _ActiveArchive_GetAsyncStatus_Handler,
		},
		{
			MethodName: "GetFileInfo",
			Handler:    _ActiveArchive_GetFileInfo_Handler,
		},
		{
			MethodName: "SearchFile",
			Handler:    _ActiveArchive_SearchFile_Handler,
		},
		{
			MethodName: "PrepareFileList",
			Handler:    _ActiveArchive_PrepareFileList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "active_archive.proto",
}
