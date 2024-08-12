// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commitserver/commit/commit.proto

package apiclient

import (
	context "context"
	fmt "fmt"
	v1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// CommitHydratedManifestsRequest is the request to commit hydrated manifests to a repository.
type CommitHydratedManifestsRequest struct {
	// Repo contains repository information including, at minimum, the URL of the repository. Generally it will contain
	// repo credentials.
	Repo *v1alpha1.Repository `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	// SyncBranch is the branch Argo CD syncs from, i.e. the hydrated branch.
	SyncBranch string `protobuf:"bytes,2,opt,name=syncBranch,proto3" json:"syncBranch,omitempty"`
	// TargetBranch is the branch Argo CD is committing to, i.e. the branch that will be updated.
	TargetBranch string `protobuf:"bytes,3,opt,name=targetBranch,proto3" json:"targetBranch,omitempty"`
	// DrySha is the commit SHA from the dry branch, i.e. pre-rendered manifest branch.
	DrySha string `protobuf:"bytes,4,opt,name=drySha,proto3" json:"drySha,omitempty"`
	// CommitMessage is the commit message to use when committing changes.
	CommitMessage string `protobuf:"bytes,5,opt,name=commitMessage,proto3" json:"commitMessage,omitempty"`
	// Paths contains the paths to write hydrated manifests to, along with the manifests and commands to execute.
	Paths                []*PathDetails `protobuf:"bytes,6,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CommitHydratedManifestsRequest) Reset()         { *m = CommitHydratedManifestsRequest{} }
func (m *CommitHydratedManifestsRequest) String() string { return proto.CompactTextString(m) }
func (*CommitHydratedManifestsRequest) ProtoMessage()    {}
func (*CommitHydratedManifestsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf3a3abbc35e3069, []int{0}
}
func (m *CommitHydratedManifestsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommitHydratedManifestsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommitHydratedManifestsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommitHydratedManifestsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitHydratedManifestsRequest.Merge(m, src)
}
func (m *CommitHydratedManifestsRequest) XXX_Size() int {
	return m.Size()
}
func (m *CommitHydratedManifestsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitHydratedManifestsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommitHydratedManifestsRequest proto.InternalMessageInfo

func (m *CommitHydratedManifestsRequest) GetRepo() *v1alpha1.Repository {
	if m != nil {
		return m.Repo
	}
	return nil
}

func (m *CommitHydratedManifestsRequest) GetSyncBranch() string {
	if m != nil {
		return m.SyncBranch
	}
	return ""
}

func (m *CommitHydratedManifestsRequest) GetTargetBranch() string {
	if m != nil {
		return m.TargetBranch
	}
	return ""
}

func (m *CommitHydratedManifestsRequest) GetDrySha() string {
	if m != nil {
		return m.DrySha
	}
	return ""
}

func (m *CommitHydratedManifestsRequest) GetCommitMessage() string {
	if m != nil {
		return m.CommitMessage
	}
	return ""
}

func (m *CommitHydratedManifestsRequest) GetPaths() []*PathDetails {
	if m != nil {
		return m.Paths
	}
	return nil
}

// PathDetails holds information about hydrated manifests to be written to a particular path in the hydrated manifests
// commit.
type PathDetails struct {
	// Path is the path to write the hydrated manifests to.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Manifests contains the manifests to write to the path.
	Manifests []*HydratedManifestDetails `protobuf:"bytes,2,rep,name=manifests,proto3" json:"manifests,omitempty"`
	// Commands contains the commands executed when hydrating the manifests.
	Commands             []string `protobuf:"bytes,3,rep,name=commands,proto3" json:"commands,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PathDetails) Reset()         { *m = PathDetails{} }
func (m *PathDetails) String() string { return proto.CompactTextString(m) }
func (*PathDetails) ProtoMessage()    {}
func (*PathDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf3a3abbc35e3069, []int{1}
}
func (m *PathDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PathDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PathDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PathDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathDetails.Merge(m, src)
}
func (m *PathDetails) XXX_Size() int {
	return m.Size()
}
func (m *PathDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_PathDetails.DiscardUnknown(m)
}

var xxx_messageInfo_PathDetails proto.InternalMessageInfo

func (m *PathDetails) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PathDetails) GetManifests() []*HydratedManifestDetails {
	if m != nil {
		return m.Manifests
	}
	return nil
}

func (m *PathDetails) GetCommands() []string {
	if m != nil {
		return m.Commands
	}
	return nil
}

// ManifestDetails contains the hydrated manifests.
type HydratedManifestDetails struct {
	// ManifestJSON is the hydrated manifest as JSON.
	ManifestJSON         string   `protobuf:"bytes,1,opt,name=manifestJSON,proto3" json:"manifestJSON,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HydratedManifestDetails) Reset()         { *m = HydratedManifestDetails{} }
func (m *HydratedManifestDetails) String() string { return proto.CompactTextString(m) }
func (*HydratedManifestDetails) ProtoMessage()    {}
func (*HydratedManifestDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf3a3abbc35e3069, []int{2}
}
func (m *HydratedManifestDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HydratedManifestDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HydratedManifestDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HydratedManifestDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HydratedManifestDetails.Merge(m, src)
}
func (m *HydratedManifestDetails) XXX_Size() int {
	return m.Size()
}
func (m *HydratedManifestDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_HydratedManifestDetails.DiscardUnknown(m)
}

var xxx_messageInfo_HydratedManifestDetails proto.InternalMessageInfo

func (m *HydratedManifestDetails) GetManifestJSON() string {
	if m != nil {
		return m.ManifestJSON
	}
	return ""
}

// ManifestsResponse is the response to the ManifestsRequest.
type CommitHydratedManifestsResponse struct {
	// HydratedSha is the commit SHA of the hydrated manifests commit.
	HydratedSha          string   `protobuf:"bytes,1,opt,name=hydratedSha,proto3" json:"hydratedSha,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitHydratedManifestsResponse) Reset()         { *m = CommitHydratedManifestsResponse{} }
func (m *CommitHydratedManifestsResponse) String() string { return proto.CompactTextString(m) }
func (*CommitHydratedManifestsResponse) ProtoMessage()    {}
func (*CommitHydratedManifestsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf3a3abbc35e3069, []int{3}
}
func (m *CommitHydratedManifestsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommitHydratedManifestsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommitHydratedManifestsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommitHydratedManifestsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitHydratedManifestsResponse.Merge(m, src)
}
func (m *CommitHydratedManifestsResponse) XXX_Size() int {
	return m.Size()
}
func (m *CommitHydratedManifestsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitHydratedManifestsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommitHydratedManifestsResponse proto.InternalMessageInfo

func (m *CommitHydratedManifestsResponse) GetHydratedSha() string {
	if m != nil {
		return m.HydratedSha
	}
	return ""
}

func init() {
	proto.RegisterType((*CommitHydratedManifestsRequest)(nil), "CommitHydratedManifestsRequest")
	proto.RegisterType((*PathDetails)(nil), "PathDetails")
	proto.RegisterType((*HydratedManifestDetails)(nil), "HydratedManifestDetails")
	proto.RegisterType((*CommitHydratedManifestsResponse)(nil), "CommitHydratedManifestsResponse")
}

func init() { proto.RegisterFile("commitserver/commit/commit.proto", fileDescriptor_cf3a3abbc35e3069) }

var fileDescriptor_cf3a3abbc35e3069 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0x24, 0x8d, 0xc8, 0xa4, 0xbd, 0xec, 0x81, 0x5a, 0x39, 0xb8, 0x96, 0xc5, 0x21,
	0x17, 0xd6, 0xaa, 0x11, 0xdc, 0xb8, 0x34, 0x1c, 0x2a, 0x44, 0x01, 0x39, 0x37, 0x54, 0x09, 0x6d,
	0xd7, 0x83, 0xbd, 0x34, 0xf6, 0x2e, 0xbb, 0x1b, 0x4b, 0x79, 0x1f, 0x1e, 0x86, 0x23, 0x8f, 0x80,
	0xf2, 0x24, 0xc8, 0x6b, 0x9b, 0xc6, 0x48, 0x69, 0x4f, 0x9e, 0xf9, 0x67, 0xf4, 0xcd, 0xe8, 0xf7,
	0x2c, 0x84, 0x5c, 0x96, 0xa5, 0xb0, 0x06, 0x75, 0x8d, 0x3a, 0x6e, 0x93, 0xee, 0x43, 0x95, 0x96,
	0x56, 0x2e, 0x3e, 0xe4, 0xc2, 0x16, 0xdb, 0x3b, 0xca, 0x65, 0x19, 0x33, 0x9d, 0x4b, 0xa5, 0xe5,
	0x77, 0x17, 0xbc, 0xe4, 0x59, 0x5c, 0x27, 0xb1, 0xba, 0xcf, 0x63, 0xa6, 0x84, 0x89, 0x99, 0x52,
	0x1b, 0xc1, 0x99, 0x15, 0xb2, 0x8a, 0xeb, 0x4b, 0xb6, 0x51, 0x05, 0xbb, 0x8c, 0x73, 0xac, 0x50,
	0x33, 0x8b, 0x59, 0x4b, 0x8b, 0x7e, 0x8e, 0x20, 0x58, 0x39, 0xfc, 0xf5, 0x2e, 0x73, 0x85, 0x1b,
	0x56, 0x89, 0x6f, 0x68, 0xac, 0x49, 0xf1, 0xc7, 0x16, 0x8d, 0x25, 0xb7, 0x30, 0xd1, 0xa8, 0xa4,
	0xef, 0x85, 0xde, 0x72, 0x9e, 0x5c, 0xd3, 0x87, 0xf9, 0xb4, 0x9f, 0xef, 0x82, 0xaf, 0x3c, 0xa3,
	0x75, 0x42, 0xd5, 0x7d, 0x4e, 0x9b, 0xf9, 0xf4, 0x60, 0x3e, 0xed, 0xe7, 0xd3, 0x14, 0x95, 0x34,
	0xc2, 0x4a, 0xbd, 0x4b, 0x1d, 0x95, 0x04, 0x00, 0x66, 0x57, 0xf1, 0x2b, 0xcd, 0x2a, 0x5e, 0xf8,
	0xa3, 0xd0, 0x5b, 0xce, 0xd2, 0x03, 0x85, 0x44, 0x70, 0x6a, 0x99, 0xce, 0xd1, 0x76, 0x1d, 0x63,
	0xd7, 0x31, 0xd0, 0xc8, 0x73, 0x98, 0x66, 0x7a, 0xb7, 0x2e, 0x98, 0x3f, 0x71, 0xd5, 0x2e, 0x23,
	0x2f, 0xe0, 0xac, 0xb5, 0xee, 0x06, 0x8d, 0x61, 0x39, 0xfa, 0x27, 0xae, 0x3c, 0x14, 0x49, 0x04,
	0x27, 0x8a, 0xd9, 0xc2, 0xf8, 0xd3, 0x70, 0xbc, 0x9c, 0x27, 0xa7, 0xf4, 0x33, 0xb3, 0xc5, 0x3b,
	0xb4, 0x4c, 0x6c, 0x4c, 0xda, 0x96, 0xa2, 0x2d, 0xcc, 0x0f, 0x54, 0x42, 0x60, 0xd2, 0xe8, 0xce,
	0x92, 0x59, 0xea, 0x62, 0xf2, 0x06, 0x66, 0x65, 0x6f, 0x9d, 0x3f, 0x72, 0x28, 0x9f, 0xfe, 0x6f,
	0x6a, 0x8f, 0x7d, 0x68, 0x25, 0x0b, 0x78, 0xd6, 0xec, 0xc3, 0xaa, 0xcc, 0xf8, 0xe3, 0x70, 0xbc,
	0x9c, 0xa5, 0xff, 0xf2, 0xe8, 0x2d, 0x9c, 0x1f, 0x21, 0x34, 0xbe, 0xf4, 0x8c, 0xf7, 0xeb, 0x4f,
	0x1f, 0xbb, 0x55, 0x06, 0x5a, 0xb4, 0x82, 0x8b, 0xa3, 0xff, 0xd6, 0x28, 0x59, 0x19, 0x24, 0x21,
	0xcc, 0x8b, 0xae, 0xd8, 0xf8, 0xd7, 0x52, 0x0e, 0xa5, 0xa4, 0x84, 0xb3, 0x16, 0xb2, 0x46, 0x5d,
	0x0b, 0x8e, 0xe4, 0x16, 0xce, 0x8f, 0x50, 0xc9, 0x05, 0x7d, 0xfc, 0x96, 0x16, 0x21, 0x7d, 0x62,
	0xa1, 0xab, 0xd5, 0xaf, 0x7d, 0xe0, 0xfd, 0xde, 0x07, 0xde, 0x9f, 0x7d, 0xe0, 0x7d, 0x79, 0xfd,
	0xc4, 0xb1, 0x0f, 0x5e, 0x0b, 0x53, 0x82, 0x6f, 0x04, 0x56, 0xf6, 0x6e, 0xea, 0x8e, 0xfb, 0xd5,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0xb8, 0xdf, 0x48, 0x4e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommitServiceClient is the client API for CommitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommitServiceClient interface {
	// Commit commits hydrated manifests to a repository.
	CommitHydratedManifests(ctx context.Context, in *CommitHydratedManifestsRequest, opts ...grpc.CallOption) (*CommitHydratedManifestsResponse, error)
}

type commitServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommitServiceClient(cc *grpc.ClientConn) CommitServiceClient {
	return &commitServiceClient{cc}
}

func (c *commitServiceClient) CommitHydratedManifests(ctx context.Context, in *CommitHydratedManifestsRequest, opts ...grpc.CallOption) (*CommitHydratedManifestsResponse, error) {
	out := new(CommitHydratedManifestsResponse)
	err := c.cc.Invoke(ctx, "/CommitService/CommitHydratedManifests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommitServiceServer is the server API for CommitService service.
type CommitServiceServer interface {
	// Commit commits hydrated manifests to a repository.
	CommitHydratedManifests(context.Context, *CommitHydratedManifestsRequest) (*CommitHydratedManifestsResponse, error)
}

// UnimplementedCommitServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommitServiceServer struct {
}

func (*UnimplementedCommitServiceServer) CommitHydratedManifests(ctx context.Context, req *CommitHydratedManifestsRequest) (*CommitHydratedManifestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitHydratedManifests not implemented")
}

func RegisterCommitServiceServer(s *grpc.Server, srv CommitServiceServer) {
	s.RegisterService(&_CommitService_serviceDesc, srv)
}

func _CommitService_CommitHydratedManifests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitHydratedManifestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommitServiceServer).CommitHydratedManifests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommitService/CommitHydratedManifests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommitServiceServer).CommitHydratedManifests(ctx, req.(*CommitHydratedManifestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CommitService",
	HandlerType: (*CommitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommitHydratedManifests",
			Handler:    _CommitService_CommitHydratedManifests_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commitserver/commit/commit.proto",
}

func (m *CommitHydratedManifestsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommitHydratedManifestsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommitHydratedManifestsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Paths) > 0 {
		for iNdEx := len(m.Paths) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Paths[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommit(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.CommitMessage) > 0 {
		i -= len(m.CommitMessage)
		copy(dAtA[i:], m.CommitMessage)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.CommitMessage)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DrySha) > 0 {
		i -= len(m.DrySha)
		copy(dAtA[i:], m.DrySha)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.DrySha)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TargetBranch) > 0 {
		i -= len(m.TargetBranch)
		copy(dAtA[i:], m.TargetBranch)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.TargetBranch)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SyncBranch) > 0 {
		i -= len(m.SyncBranch)
		copy(dAtA[i:], m.SyncBranch)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.SyncBranch)))
		i--
		dAtA[i] = 0x12
	}
	if m.Repo != nil {
		{
			size, err := m.Repo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PathDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PathDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PathDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Commands) > 0 {
		for iNdEx := len(m.Commands) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Commands[iNdEx])
			copy(dAtA[i:], m.Commands[iNdEx])
			i = encodeVarintCommit(dAtA, i, uint64(len(m.Commands[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Manifests) > 0 {
		for iNdEx := len(m.Manifests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Manifests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommit(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HydratedManifestDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HydratedManifestDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HydratedManifestDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ManifestJSON) > 0 {
		i -= len(m.ManifestJSON)
		copy(dAtA[i:], m.ManifestJSON)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.ManifestJSON)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommitHydratedManifestsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommitHydratedManifestsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommitHydratedManifestsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.HydratedSha) > 0 {
		i -= len(m.HydratedSha)
		copy(dAtA[i:], m.HydratedSha)
		i = encodeVarintCommit(dAtA, i, uint64(len(m.HydratedSha)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommit(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CommitHydratedManifestsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Repo != nil {
		l = m.Repo.Size()
		n += 1 + l + sovCommit(uint64(l))
	}
	l = len(m.SyncBranch)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	l = len(m.TargetBranch)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	l = len(m.DrySha)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	l = len(m.CommitMessage)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	if len(m.Paths) > 0 {
		for _, e := range m.Paths {
			l = e.Size()
			n += 1 + l + sovCommit(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PathDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	if len(m.Manifests) > 0 {
		for _, e := range m.Manifests {
			l = e.Size()
			n += 1 + l + sovCommit(uint64(l))
		}
	}
	if len(m.Commands) > 0 {
		for _, s := range m.Commands {
			l = len(s)
			n += 1 + l + sovCommit(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HydratedManifestDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ManifestJSON)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CommitHydratedManifestsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HydratedSha)
	if l > 0 {
		n += 1 + l + sovCommit(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCommit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommit(x uint64) (n int) {
	return sovCommit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommitHydratedManifestsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommitHydratedManifestsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommitHydratedManifestsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Repo == nil {
				m.Repo = &v1alpha1.Repository{}
			}
			if err := m.Repo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SyncBranch", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SyncBranch = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetBranch", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetBranch = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DrySha", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DrySha = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommitMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paths", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Paths = append(m.Paths, &PathDetails{})
			if err := m.Paths[len(m.Paths)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PathDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PathDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PathDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Manifests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Manifests = append(m.Manifests, &HydratedManifestDetails{})
			if err := m.Manifests[len(m.Manifests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commands", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commands = append(m.Commands, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HydratedManifestDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HydratedManifestDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HydratedManifestDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ManifestJSON", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ManifestJSON = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CommitHydratedManifestsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommitHydratedManifestsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommitHydratedManifestsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HydratedSha", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HydratedSha = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCommit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommit
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCommit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommit = fmt.Errorf("proto: unexpected end of group")
)
