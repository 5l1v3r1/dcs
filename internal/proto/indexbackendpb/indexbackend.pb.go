// Code generated by protoc-gen-go. DO NOT EDIT.
// source: indexbackend.proto

/*
Package indexbackendpb is a generated protocol buffer package.

It is generated from these files:
	indexbackend.proto

It has these top-level messages:
	FilesRequest
	FilesReply
	ReplaceIndexRequest
	ReplaceIndexReply
*/
package indexbackendpb

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

type FilesRequest struct {
	// Text query (e.g. “i3Font”) which will be translated into a trigram query
	// (e.g. "3Fo" "Fon" "i3F" "ont").
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *FilesRequest) Reset()                    { *m = FilesRequest{} }
func (m *FilesRequest) String() string            { return proto.CompactTextString(m) }
func (*FilesRequest) ProtoMessage()               {}
func (*FilesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FilesRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type FilesReply struct {
	// A path which match the requested trigram query (likely to match
	// the regular expression from which the trigram query was derived, but can
	// contain false positives).
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *FilesReply) Reset()                    { *m = FilesReply{} }
func (m *FilesReply) String() string            { return proto.CompactTextString(m) }
func (*FilesReply) ProtoMessage()               {}
func (*FilesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FilesReply) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ReplaceIndexRequest struct {
	ReplacementPath string `protobuf:"bytes,1,opt,name=replacement_path,json=replacementPath" json:"replacement_path,omitempty"`
}

func (m *ReplaceIndexRequest) Reset()                    { *m = ReplaceIndexRequest{} }
func (m *ReplaceIndexRequest) String() string            { return proto.CompactTextString(m) }
func (*ReplaceIndexRequest) ProtoMessage()               {}
func (*ReplaceIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReplaceIndexRequest) GetReplacementPath() string {
	if m != nil {
		return m.ReplacementPath
	}
	return ""
}

type ReplaceIndexReply struct {
}

func (m *ReplaceIndexReply) Reset()                    { *m = ReplaceIndexReply{} }
func (m *ReplaceIndexReply) String() string            { return proto.CompactTextString(m) }
func (*ReplaceIndexReply) ProtoMessage()               {}
func (*ReplaceIndexReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*FilesRequest)(nil), "indexbackendpb.FilesRequest")
	proto.RegisterType((*FilesReply)(nil), "indexbackendpb.FilesReply")
	proto.RegisterType((*ReplaceIndexRequest)(nil), "indexbackendpb.ReplaceIndexRequest")
	proto.RegisterType((*ReplaceIndexReply)(nil), "indexbackendpb.ReplaceIndexReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for IndexBackend service

type IndexBackendClient interface {
	// Files returns a list of files which match the specified query in the
	// trigram index.
	Files(ctx context.Context, in *FilesRequest, opts ...grpc.CallOption) (IndexBackend_FilesClient, error)
	// Replaces the loaded index with the specified replacement index. On a file
	// system level, the specified file is mv'ed to the file specified by
	// -index_path.
	ReplaceIndex(ctx context.Context, in *ReplaceIndexRequest, opts ...grpc.CallOption) (*ReplaceIndexReply, error)
}

type indexBackendClient struct {
	cc *grpc.ClientConn
}

func NewIndexBackendClient(cc *grpc.ClientConn) IndexBackendClient {
	return &indexBackendClient{cc}
}

func (c *indexBackendClient) Files(ctx context.Context, in *FilesRequest, opts ...grpc.CallOption) (IndexBackend_FilesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_IndexBackend_serviceDesc.Streams[0], c.cc, "/indexbackendpb.IndexBackend/Files", opts...)
	if err != nil {
		return nil, err
	}
	x := &indexBackendFilesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IndexBackend_FilesClient interface {
	Recv() (*FilesReply, error)
	grpc.ClientStream
}

type indexBackendFilesClient struct {
	grpc.ClientStream
}

func (x *indexBackendFilesClient) Recv() (*FilesReply, error) {
	m := new(FilesReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *indexBackendClient) ReplaceIndex(ctx context.Context, in *ReplaceIndexRequest, opts ...grpc.CallOption) (*ReplaceIndexReply, error) {
	out := new(ReplaceIndexReply)
	err := grpc.Invoke(ctx, "/indexbackendpb.IndexBackend/ReplaceIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IndexBackend service

type IndexBackendServer interface {
	// Files returns a list of files which match the specified query in the
	// trigram index.
	Files(*FilesRequest, IndexBackend_FilesServer) error
	// Replaces the loaded index with the specified replacement index. On a file
	// system level, the specified file is mv'ed to the file specified by
	// -index_path.
	ReplaceIndex(context.Context, *ReplaceIndexRequest) (*ReplaceIndexReply, error)
}

func RegisterIndexBackendServer(s *grpc.Server, srv IndexBackendServer) {
	s.RegisterService(&_IndexBackend_serviceDesc, srv)
}

func _IndexBackend_Files_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FilesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IndexBackendServer).Files(m, &indexBackendFilesServer{stream})
}

type IndexBackend_FilesServer interface {
	Send(*FilesReply) error
	grpc.ServerStream
}

type indexBackendFilesServer struct {
	grpc.ServerStream
}

func (x *indexBackendFilesServer) Send(m *FilesReply) error {
	return x.ServerStream.SendMsg(m)
}

func _IndexBackend_ReplaceIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexBackendServer).ReplaceIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexbackendpb.IndexBackend/ReplaceIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexBackendServer).ReplaceIndex(ctx, req.(*ReplaceIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IndexBackend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "indexbackendpb.IndexBackend",
	HandlerType: (*IndexBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReplaceIndex",
			Handler:    _IndexBackend_ReplaceIndex_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Files",
			Handler:       _IndexBackend_Files_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "indexbackend.proto",
}

func init() { proto.RegisterFile("indexbackend.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xcc, 0x4b, 0x49,
	0xad, 0x48, 0x4a, 0x4c, 0xce, 0x4e, 0xcd, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x43, 0x16, 0x2b, 0x48, 0x52, 0x52, 0xe1, 0xe2, 0x71, 0xcb, 0xcc, 0x49, 0x2d, 0x0e, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0x2c, 0x4d, 0x2d, 0xaa, 0x94, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0x14, 0xb8, 0xb8, 0xa0, 0xaa, 0x0a, 0x72, 0x2a, 0x85,
	0x84, 0xb8, 0x58, 0x0a, 0x12, 0x4b, 0x32, 0xa0, 0x4a, 0xc0, 0x6c, 0x25, 0x07, 0x2e, 0x61, 0x90,
	0x64, 0x62, 0x72, 0xaa, 0x27, 0xc8, 0x02, 0x98, 0x71, 0x9a, 0x5c, 0x02, 0x45, 0x10, 0xe1, 0xdc,
	0xd4, 0xbc, 0x92, 0x78, 0x24, 0x6d, 0xfc, 0x48, 0xe2, 0x01, 0x20, 0x13, 0x84, 0xb9, 0x04, 0x51,
	0x4d, 0x28, 0xc8, 0xa9, 0x34, 0x5a, 0xcf, 0xc8, 0xc5, 0x03, 0xe6, 0x3a, 0x41, 0x5c, 0x2c, 0xe4,
	0xca, 0xc5, 0x0a, 0x76, 0x89, 0x90, 0x8c, 0x1e, 0xaa, 0x4f, 0xf4, 0x90, 0xbd, 0x21, 0x25, 0x85,
	0x43, 0xb6, 0x20, 0xa7, 0x52, 0x89, 0xc1, 0x80, 0x51, 0x28, 0x82, 0x8b, 0x07, 0xd9, 0x32, 0x21,
	0x65, 0x74, 0xf5, 0x58, 0x3c, 0x23, 0xa5, 0x88, 0x5f, 0x11, 0xd8, 0xec, 0x24, 0x36, 0x70, 0x38,
	0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x97, 0x10, 0x08, 0x98, 0x7d, 0x01, 0x00, 0x00,
}