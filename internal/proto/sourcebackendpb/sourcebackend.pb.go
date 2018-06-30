// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sourcebackend.proto

package sourcebackendpb // import "github.com/Debian/dcs/internal/proto/sourcebackendpb"

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

type SearchReply_Type int32

const (
	SearchReply_MATCH           SearchReply_Type = 0
	SearchReply_PROGRESS_UPDATE SearchReply_Type = 1
)

var SearchReply_Type_name = map[int32]string{
	0: "MATCH",
	1: "PROGRESS_UPDATE",
}
var SearchReply_Type_value = map[string]int32{
	"MATCH":           0,
	"PROGRESS_UPDATE": 1,
}

func (x SearchReply_Type) String() string {
	return proto.EnumName(SearchReply_Type_name, int32(x))
}
func (SearchReply_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_sourcebackend_1a3dc62c025055f3, []int{5, 0}
}

type FileRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileRequest) Reset()         { *m = FileRequest{} }
func (m *FileRequest) String() string { return proto.CompactTextString(m) }
func (*FileRequest) ProtoMessage()    {}
func (*FileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sourcebackend_1a3dc62c025055f3, []int{0}
}
func (m *FileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileRequest.Unmarshal(m, b)
}
func (m *FileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileRequest.Marshal(b, m, deterministic)
}
func (dst *FileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileRequest.Merge(dst, src)
}
func (m *FileRequest) XXX_Size() int {
	return xxx_messageInfo_FileRequest.Size(m)
}
func (m *FileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileRequest proto.InternalMessageInfo

func (m *FileRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type FileReply struct {
	Contents             []byte   `protobuf:"bytes,1,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileReply) Reset()         { *m = FileReply{} }
func (m *FileReply) String() string { return proto.CompactTextString(m) }
func (*FileReply) ProtoMessage()    {}
func (*FileReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_sourcebackend_1a3dc62c025055f3, []int{1}
}
func (m *FileReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileReply.Unmarshal(m, b)
}
func (m *FileReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileReply.Marshal(b, m, deterministic)
}
func (dst *FileReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileReply.Merge(dst, src)
}
func (m *FileReply) XXX_Size() int {
	return xxx_messageInfo_FileReply.Size(m)
}
func (m *FileReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FileReply.DiscardUnknown(m)
}

var xxx_messageInfo_FileReply proto.InternalMessageInfo

func (m *FileReply) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

type SearchRequest struct {
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// Rewritten URL (after RewriteQuery()) with all the parameters that
	// are relevant for ranking.
	RewrittenUrl         string   `protobuf:"bytes,2,opt,name=rewritten_url,json=rewrittenUrl,proto3" json:"rewritten_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sourcebackend_1a3dc62c025055f3, []int{2}
}
func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (dst *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(dst, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetRewrittenUrl() string {
	if m != nil {
		return m.RewrittenUrl
	}
	return ""
}

type Match struct {
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Line uint32 `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	// Contents of line-2.
	Ctxp2 string `protobuf:"bytes,3,opt,name=ctxp2,proto3" json:"ctxp2,omitempty"`
	// Contents of line-1.
	Ctxp1 string `protobuf:"bytes,4,opt,name=ctxp1,proto3" json:"ctxp1,omitempty"`
	// Contents of the line containing the match.
	Context string `protobuf:"bytes,5,opt,name=context,proto3" json:"context,omitempty"`
	// Contents of line+1.
	Ctxn1 string `protobuf:"bytes,6,opt,name=ctxn1,proto3" json:"ctxn1,omitempty"`
	// Contents of line+2.
	Ctxn2                string   `protobuf:"bytes,7,opt,name=ctxn2,proto3" json:"ctxn2,omitempty"`
	Pathrank             float32  `protobuf:"fixed32,8,opt,name=pathrank,proto3" json:"pathrank,omitempty"`
	Ranking              float32  `protobuf:"fixed32,9,opt,name=ranking,proto3" json:"ranking,omitempty"`
	Package              string   `protobuf:"bytes,10,opt,name=package,proto3" json:"package,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Match) Reset()         { *m = Match{} }
func (m *Match) String() string { return proto.CompactTextString(m) }
func (*Match) ProtoMessage()    {}
func (*Match) Descriptor() ([]byte, []int) {
	return fileDescriptor_sourcebackend_1a3dc62c025055f3, []int{3}
}
func (m *Match) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Match.Unmarshal(m, b)
}
func (m *Match) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Match.Marshal(b, m, deterministic)
}
func (dst *Match) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Match.Merge(dst, src)
}
func (m *Match) XXX_Size() int {
	return xxx_messageInfo_Match.Size(m)
}
func (m *Match) XXX_DiscardUnknown() {
	xxx_messageInfo_Match.DiscardUnknown(m)
}

var xxx_messageInfo_Match proto.InternalMessageInfo

func (m *Match) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Match) GetLine() uint32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *Match) GetCtxp2() string {
	if m != nil {
		return m.Ctxp2
	}
	return ""
}

func (m *Match) GetCtxp1() string {
	if m != nil {
		return m.Ctxp1
	}
	return ""
}

func (m *Match) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *Match) GetCtxn1() string {
	if m != nil {
		return m.Ctxn1
	}
	return ""
}

func (m *Match) GetCtxn2() string {
	if m != nil {
		return m.Ctxn2
	}
	return ""
}

func (m *Match) GetPathrank() float32 {
	if m != nil {
		return m.Pathrank
	}
	return 0
}

func (m *Match) GetRanking() float32 {
	if m != nil {
		return m.Ranking
	}
	return 0
}

func (m *Match) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

type ProgressUpdate struct {
	FilesProcessed       uint64   `protobuf:"varint,1,opt,name=files_processed,json=filesProcessed,proto3" json:"files_processed,omitempty"`
	FilesTotal           uint64   `protobuf:"varint,2,opt,name=files_total,json=filesTotal,proto3" json:"files_total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProgressUpdate) Reset()         { *m = ProgressUpdate{} }
func (m *ProgressUpdate) String() string { return proto.CompactTextString(m) }
func (*ProgressUpdate) ProtoMessage()    {}
func (*ProgressUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_sourcebackend_1a3dc62c025055f3, []int{4}
}
func (m *ProgressUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProgressUpdate.Unmarshal(m, b)
}
func (m *ProgressUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProgressUpdate.Marshal(b, m, deterministic)
}
func (dst *ProgressUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProgressUpdate.Merge(dst, src)
}
func (m *ProgressUpdate) XXX_Size() int {
	return xxx_messageInfo_ProgressUpdate.Size(m)
}
func (m *ProgressUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ProgressUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ProgressUpdate proto.InternalMessageInfo

func (m *ProgressUpdate) GetFilesProcessed() uint64 {
	if m != nil {
		return m.FilesProcessed
	}
	return 0
}

func (m *ProgressUpdate) GetFilesTotal() uint64 {
	if m != nil {
		return m.FilesTotal
	}
	return 0
}

type SearchReply struct {
	Type                 SearchReply_Type `protobuf:"varint,1,opt,name=type,proto3,enum=sourcebackendpb.SearchReply_Type" json:"type,omitempty"`
	Match                *Match           `protobuf:"bytes,2,opt,name=match,proto3" json:"match,omitempty"`
	ProgressUpdate       *ProgressUpdate  `protobuf:"bytes,3,opt,name=progress_update,json=progressUpdate,proto3" json:"progress_update,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SearchReply) Reset()         { *m = SearchReply{} }
func (m *SearchReply) String() string { return proto.CompactTextString(m) }
func (*SearchReply) ProtoMessage()    {}
func (*SearchReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_sourcebackend_1a3dc62c025055f3, []int{5}
}
func (m *SearchReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchReply.Unmarshal(m, b)
}
func (m *SearchReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchReply.Marshal(b, m, deterministic)
}
func (dst *SearchReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchReply.Merge(dst, src)
}
func (m *SearchReply) XXX_Size() int {
	return xxx_messageInfo_SearchReply.Size(m)
}
func (m *SearchReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchReply.DiscardUnknown(m)
}

var xxx_messageInfo_SearchReply proto.InternalMessageInfo

func (m *SearchReply) GetType() SearchReply_Type {
	if m != nil {
		return m.Type
	}
	return SearchReply_MATCH
}

func (m *SearchReply) GetMatch() *Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *SearchReply) GetProgressUpdate() *ProgressUpdate {
	if m != nil {
		return m.ProgressUpdate
	}
	return nil
}

type ReplaceIndexRequest struct {
	ReplacementPath      string   `protobuf:"bytes,1,opt,name=replacement_path,json=replacementPath,proto3" json:"replacement_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplaceIndexRequest) Reset()         { *m = ReplaceIndexRequest{} }
func (m *ReplaceIndexRequest) String() string { return proto.CompactTextString(m) }
func (*ReplaceIndexRequest) ProtoMessage()    {}
func (*ReplaceIndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sourcebackend_1a3dc62c025055f3, []int{6}
}
func (m *ReplaceIndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplaceIndexRequest.Unmarshal(m, b)
}
func (m *ReplaceIndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplaceIndexRequest.Marshal(b, m, deterministic)
}
func (dst *ReplaceIndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceIndexRequest.Merge(dst, src)
}
func (m *ReplaceIndexRequest) XXX_Size() int {
	return xxx_messageInfo_ReplaceIndexRequest.Size(m)
}
func (m *ReplaceIndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceIndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceIndexRequest proto.InternalMessageInfo

func (m *ReplaceIndexRequest) GetReplacementPath() string {
	if m != nil {
		return m.ReplacementPath
	}
	return ""
}

type ReplaceIndexReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplaceIndexReply) Reset()         { *m = ReplaceIndexReply{} }
func (m *ReplaceIndexReply) String() string { return proto.CompactTextString(m) }
func (*ReplaceIndexReply) ProtoMessage()    {}
func (*ReplaceIndexReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_sourcebackend_1a3dc62c025055f3, []int{7}
}
func (m *ReplaceIndexReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplaceIndexReply.Unmarshal(m, b)
}
func (m *ReplaceIndexReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplaceIndexReply.Marshal(b, m, deterministic)
}
func (dst *ReplaceIndexReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceIndexReply.Merge(dst, src)
}
func (m *ReplaceIndexReply) XXX_Size() int {
	return xxx_messageInfo_ReplaceIndexReply.Size(m)
}
func (m *ReplaceIndexReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceIndexReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceIndexReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FileRequest)(nil), "sourcebackendpb.FileRequest")
	proto.RegisterType((*FileReply)(nil), "sourcebackendpb.FileReply")
	proto.RegisterType((*SearchRequest)(nil), "sourcebackendpb.SearchRequest")
	proto.RegisterType((*Match)(nil), "sourcebackendpb.Match")
	proto.RegisterType((*ProgressUpdate)(nil), "sourcebackendpb.ProgressUpdate")
	proto.RegisterType((*SearchReply)(nil), "sourcebackendpb.SearchReply")
	proto.RegisterType((*ReplaceIndexRequest)(nil), "sourcebackendpb.ReplaceIndexRequest")
	proto.RegisterType((*ReplaceIndexReply)(nil), "sourcebackendpb.ReplaceIndexReply")
	proto.RegisterEnum("sourcebackendpb.SearchReply_Type", SearchReply_Type_name, SearchReply_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SourceBackendClient is the client API for SourceBackend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SourceBackendClient interface {
	// File reads the file and returns its contents.
	File(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileReply, error)
	// Search performs the given query and streams matches/progress updates.
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (SourceBackend_SearchClient, error)
	// Replaces the loaded index with the specified replacement index. On a file
	// system level, the specified file is mv'ed to the file specified by
	// -index_path.
	ReplaceIndex(ctx context.Context, in *ReplaceIndexRequest, opts ...grpc.CallOption) (*ReplaceIndexReply, error)
}

type sourceBackendClient struct {
	cc *grpc.ClientConn
}

func NewSourceBackendClient(cc *grpc.ClientConn) SourceBackendClient {
	return &sourceBackendClient{cc}
}

func (c *sourceBackendClient) File(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileReply, error) {
	out := new(FileReply)
	err := c.cc.Invoke(ctx, "/sourcebackendpb.SourceBackend/File", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceBackendClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (SourceBackend_SearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SourceBackend_serviceDesc.Streams[0], "/sourcebackendpb.SourceBackend/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &sourceBackendSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SourceBackend_SearchClient interface {
	Recv() (*SearchReply, error)
	grpc.ClientStream
}

type sourceBackendSearchClient struct {
	grpc.ClientStream
}

func (x *sourceBackendSearchClient) Recv() (*SearchReply, error) {
	m := new(SearchReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sourceBackendClient) ReplaceIndex(ctx context.Context, in *ReplaceIndexRequest, opts ...grpc.CallOption) (*ReplaceIndexReply, error) {
	out := new(ReplaceIndexReply)
	err := c.cc.Invoke(ctx, "/sourcebackendpb.SourceBackend/ReplaceIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SourceBackendServer is the server API for SourceBackend service.
type SourceBackendServer interface {
	// File reads the file and returns its contents.
	File(context.Context, *FileRequest) (*FileReply, error)
	// Search performs the given query and streams matches/progress updates.
	Search(*SearchRequest, SourceBackend_SearchServer) error
	// Replaces the loaded index with the specified replacement index. On a file
	// system level, the specified file is mv'ed to the file specified by
	// -index_path.
	ReplaceIndex(context.Context, *ReplaceIndexRequest) (*ReplaceIndexReply, error)
}

func RegisterSourceBackendServer(s *grpc.Server, srv SourceBackendServer) {
	s.RegisterService(&_SourceBackend_serviceDesc, srv)
}

func _SourceBackend_File_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceBackendServer).File(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sourcebackendpb.SourceBackend/File",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceBackendServer).File(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourceBackend_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SourceBackendServer).Search(m, &sourceBackendSearchServer{stream})
}

type SourceBackend_SearchServer interface {
	Send(*SearchReply) error
	grpc.ServerStream
}

type sourceBackendSearchServer struct {
	grpc.ServerStream
}

func (x *sourceBackendSearchServer) Send(m *SearchReply) error {
	return x.ServerStream.SendMsg(m)
}

func _SourceBackend_ReplaceIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceBackendServer).ReplaceIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sourcebackendpb.SourceBackend/ReplaceIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceBackendServer).ReplaceIndex(ctx, req.(*ReplaceIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SourceBackend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sourcebackendpb.SourceBackend",
	HandlerType: (*SourceBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "File",
			Handler:    _SourceBackend_File_Handler,
		},
		{
			MethodName: "ReplaceIndex",
			Handler:    _SourceBackend_ReplaceIndex_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _SourceBackend_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sourcebackend.proto",
}

func init() { proto.RegisterFile("sourcebackend.proto", fileDescriptor_sourcebackend_1a3dc62c025055f3) }

var fileDescriptor_sourcebackend_1a3dc62c025055f3 = []byte{
	// 582 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5f, 0x6f, 0xda, 0x3e,
	0x14, 0x6d, 0xfa, 0x0b, 0x6d, 0xb9, 0x14, 0xe8, 0xcf, 0x4c, 0x53, 0x84, 0xaa, 0xb5, 0xcd, 0xa6,
	0xb5, 0x93, 0x26, 0x18, 0xd9, 0x9f, 0xe7, 0xb5, 0x6b, 0xb7, 0xae, 0x52, 0x35, 0x64, 0xe8, 0x0b,
	0x2f, 0xc8, 0x84, 0x3b, 0x88, 0x08, 0x8e, 0xeb, 0x38, 0x1a, 0xf9, 0xbe, 0xfb, 0x02, 0x7b, 0xde,
	0xcb, 0x64, 0x87, 0x50, 0x28, 0xed, 0xf6, 0x44, 0xce, 0xf1, 0xf1, 0xbd, 0xe7, 0xfe, 0x31, 0x50,
	0x8b, 0xa3, 0x44, 0xfa, 0x38, 0x60, 0xfe, 0x04, 0xf9, 0xb0, 0x21, 0x64, 0xa4, 0x22, 0x52, 0x5d,
	0x21, 0xc5, 0xc0, 0x3d, 0x82, 0xd2, 0xe7, 0x20, 0x44, 0x8a, 0xb7, 0x09, 0xc6, 0x8a, 0x10, 0xb0,
	0x05, 0x53, 0x63, 0xc7, 0x3a, 0xb4, 0x4e, 0x8a, 0xd4, 0x7c, 0xbb, 0xc7, 0x50, 0xcc, 0x24, 0x22,
	0x4c, 0x49, 0x1d, 0x76, 0xfc, 0x88, 0x2b, 0xe4, 0x2a, 0x36, 0xa2, 0x5d, 0xba, 0xc0, 0xee, 0x15,
	0x94, 0x3b, 0xc8, 0xa4, 0x3f, 0xce, 0xa3, 0x3d, 0x81, 0xc2, 0x6d, 0x82, 0x32, 0x9d, 0x87, 0xcb,
	0x00, 0x79, 0x0e, 0x65, 0x89, 0x3f, 0x64, 0xa0, 0x14, 0xf2, 0x7e, 0x22, 0x43, 0x67, 0xd3, 0x9c,
	0xee, 0x2e, 0xc8, 0x1b, 0x19, 0xba, 0xbf, 0x2c, 0x28, 0x5c, 0x33, 0xe5, 0x8f, 0x1f, 0xb2, 0xa4,
	0xb9, 0x30, 0xe0, 0x68, 0x6e, 0x96, 0xa9, 0xf9, 0xd6, 0xc9, 0x7c, 0x35, 0x13, 0x9e, 0xf3, 0x5f,
	0x96, 0xcc, 0x80, 0x9c, 0x6d, 0x39, 0xf6, 0x1d, 0xdb, 0x22, 0x0e, 0x6c, 0x1b, 0xd7, 0x33, 0xe5,
	0x14, 0x0c, 0x9f, 0xc3, 0xb9, 0x9e, 0xb7, 0x9c, 0xad, 0x85, 0x9e, 0xb7, 0x72, 0xd6, 0x73, 0xb6,
	0xef, 0x58, 0x4f, 0xf7, 0x42, 0xbb, 0x91, 0x8c, 0x4f, 0x9c, 0x9d, 0x43, 0xeb, 0x64, 0x93, 0x2e,
	0xb0, 0xce, 0xa0, 0x7f, 0x03, 0x3e, 0x72, 0x8a, 0xe6, 0x28, 0x87, 0xfa, 0x44, 0x30, 0x7f, 0xc2,
	0x46, 0xe8, 0x40, 0x96, 0x7b, 0x0e, 0xdd, 0x1e, 0x54, 0xda, 0x32, 0x1a, 0x49, 0x8c, 0xe3, 0x1b,
	0x31, 0x64, 0x0a, 0xc9, 0x31, 0x54, 0xbf, 0x07, 0x21, 0xc6, 0x7d, 0x21, 0x23, 0x1f, 0xe3, 0x18,
	0x87, 0xa6, 0x0d, 0x36, 0xad, 0x18, 0xba, 0x9d, 0xb3, 0xe4, 0x00, 0x4a, 0x99, 0x50, 0x45, 0x8a,
	0x65, 0x1d, 0xb5, 0x29, 0x18, 0xaa, 0xab, 0x19, 0xf7, 0xa7, 0x05, 0xa5, 0x7c, 0x38, 0x7a, 0x8e,
	0xef, 0xc1, 0x56, 0xa9, 0x40, 0x13, 0xae, 0xe2, 0x1d, 0x35, 0xee, 0xed, 0x45, 0x63, 0x49, 0xdb,
	0xe8, 0xa6, 0x02, 0xa9, 0x91, 0x93, 0xd7, 0x50, 0x98, 0xea, 0xa9, 0x98, 0x0c, 0x25, 0xef, 0xe9,
	0xda, 0x3d, 0x33, 0x33, 0x9a, 0x89, 0xc8, 0x25, 0x54, 0xc5, 0xbc, 0xa0, 0x7e, 0x62, 0x2a, 0x32,
	0xc3, 0x29, 0x79, 0x07, 0x6b, 0xf7, 0x56, 0x0b, 0xa7, 0x15, 0xb1, 0x82, 0xdd, 0x97, 0x60, 0x6b,
	0x17, 0xa4, 0x08, 0x85, 0xeb, 0xd3, 0xee, 0xa7, 0xcb, 0xbd, 0x0d, 0x52, 0x83, 0x6a, 0x9b, 0x7e,
	0xfb, 0x42, 0x2f, 0x3a, 0x9d, 0xfe, 0x4d, 0xfb, 0xfc, 0xb4, 0x7b, 0xb1, 0x67, 0xb9, 0x1f, 0xa1,
	0xa6, 0x3d, 0x33, 0x1f, 0xbf, 0xf2, 0x21, 0xce, 0xf2, 0x45, 0x7c, 0x05, 0x7b, 0x32, 0xa3, 0xa7,
	0xc8, 0x55, 0x7f, 0x69, 0x9f, 0xaa, 0x4b, 0x7c, 0x5b, 0x6f, 0x7b, 0x0d, 0xfe, 0x5f, 0x8d, 0x20,
	0xc2, 0xd4, 0xfb, 0x6d, 0x41, 0xb9, 0x63, 0x1c, 0x9f, 0x65, 0x8e, 0xc9, 0x19, 0xd8, 0xfa, 0x51,
	0x90, 0xfd, 0xb5, 0x4a, 0x96, 0x9e, 0x53, 0xbd, 0xfe, 0xc8, 0xa9, 0x08, 0x53, 0x77, 0x83, 0x5c,
	0xc1, 0x56, 0xd6, 0x66, 0xf2, 0xec, 0xd1, 0xfe, 0x67, 0x71, 0xf6, 0xff, 0x36, 0x1f, 0x77, 0xe3,
	0x8d, 0x45, 0x7a, 0xb0, 0xbb, 0x6c, 0x9b, 0xbc, 0x58, 0xbb, 0xf1, 0x40, 0x5f, 0xea, 0xee, 0x3f,
	0x54, 0x26, 0xfa, 0xd9, 0x87, 0xde, 0xbb, 0x51, 0xa0, 0xc6, 0xc9, 0xa0, 0xe1, 0x47, 0xd3, 0xe6,
	0x39, 0x0e, 0x02, 0xc6, 0x9b, 0x43, 0x3f, 0x6e, 0x06, 0x5c, 0xa1, 0xe4, 0x2c, 0x6c, 0x9a, 0x3f,
	0x97, 0xe6, 0xbd, 0x58, 0x83, 0x2d, 0x43, 0xbf, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x0a,
	0x96, 0x6b, 0x8a, 0x04, 0x00, 0x00,
}
