// Code generated by protoc-gen-go.
// source: imager.proto
// DO NOT EDIT!

package pb

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

type CreateJobRequest struct {
	Name    string      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Content *InputImage `protobuf:"bytes,5,opt,name=content" json:"content,omitempty"`
}

func (m *CreateJobRequest) Reset()                    { *m = CreateJobRequest{} }
func (m *CreateJobRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateJobRequest) ProtoMessage()               {}
func (*CreateJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateJobRequest) GetContent() *InputImage {
	if m != nil {
		return m.Content
	}
	return nil
}

type CreateJobResponse struct {
}

func (m *CreateJobResponse) Reset()                    { *m = CreateJobResponse{} }
func (m *CreateJobResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateJobResponse) ProtoMessage()               {}
func (*CreateJobResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type CreateFullJobRequest struct {
	Name    string      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Style   *InputImage `protobuf:"bytes,3,opt,name=style" json:"style,omitempty"`
	Content *InputImage `protobuf:"bytes,5,opt,name=content" json:"content,omitempty"`
}

func (m *CreateFullJobRequest) Reset()                    { *m = CreateFullJobRequest{} }
func (m *CreateFullJobRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateFullJobRequest) ProtoMessage()               {}
func (*CreateFullJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *CreateFullJobRequest) GetStyle() *InputImage {
	if m != nil {
		return m.Style
	}
	return nil
}

func (m *CreateFullJobRequest) GetContent() *InputImage {
	if m != nil {
		return m.Content
	}
	return nil
}

type CreateFullJobResponse struct {
}

func (m *CreateFullJobResponse) Reset()                    { *m = CreateFullJobResponse{} }
func (m *CreateFullJobResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateFullJobResponse) ProtoMessage()               {}
func (*CreateFullJobResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func init() {
	proto.RegisterType((*CreateJobRequest)(nil), "CreateJobRequest")
	proto.RegisterType((*CreateJobResponse)(nil), "CreateJobResponse")
	proto.RegisterType((*CreateFullJobRequest)(nil), "CreateFullJobRequest")
	proto.RegisterType((*CreateFullJobResponse)(nil), "CreateFullJobResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for NeuralStyleImager service

type NeuralStyleImagerClient interface {
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobResponse, error)
	CreateFullJob(ctx context.Context, in *CreateFullJobRequest, opts ...grpc.CallOption) (*CreateFullJobResponse, error)
}

type neuralStyleImagerClient struct {
	cc *grpc.ClientConn
}

func NewNeuralStyleImagerClient(cc *grpc.ClientConn) NeuralStyleImagerClient {
	return &neuralStyleImagerClient{cc}
}

func (c *neuralStyleImagerClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobResponse, error) {
	out := new(CreateJobResponse)
	err := grpc.Invoke(ctx, "/NeuralStyleImager/CreateJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neuralStyleImagerClient) CreateFullJob(ctx context.Context, in *CreateFullJobRequest, opts ...grpc.CallOption) (*CreateFullJobResponse, error) {
	out := new(CreateFullJobResponse)
	err := grpc.Invoke(ctx, "/NeuralStyleImager/CreateFullJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NeuralStyleImager service

type NeuralStyleImagerServer interface {
	CreateJob(context.Context, *CreateJobRequest) (*CreateJobResponse, error)
	CreateFullJob(context.Context, *CreateFullJobRequest) (*CreateFullJobResponse, error)
}

func RegisterNeuralStyleImagerServer(s *grpc.Server, srv NeuralStyleImagerServer) {
	s.RegisterService(&_NeuralStyleImager_serviceDesc, srv)
}

func _NeuralStyleImager_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NeuralStyleImagerServer).CreateJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _NeuralStyleImager_CreateFullJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateFullJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NeuralStyleImagerServer).CreateFullJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _NeuralStyleImager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NeuralStyleImager",
	HandlerType: (*NeuralStyleImagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _NeuralStyleImager_CreateJob_Handler,
		},
		{
			MethodName: "CreateFullJob",
			Handler:    _NeuralStyleImager_CreateFullJob_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor1 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xcc, 0x4d, 0x4c,
	0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0x06, 0xf3, 0x20, 0x1c, 0x25, 0x5f,
	0x2e, 0x01, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0x54, 0xaf, 0xfc, 0xa4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4,
	0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x30, 0x5b, 0x48, 0x95, 0x8b, 0x3d, 0x39, 0x3f, 0xaf, 0x24, 0x35, 0xaf, 0x44, 0x82, 0x15,
	0x28, 0xcc, 0x6d, 0xc4, 0xad, 0xe7, 0x99, 0x57, 0x50, 0x5a, 0xe2, 0x09, 0x32, 0x2b, 0x08, 0x26,
	0xa7, 0x24, 0xcc, 0x25, 0x88, 0x64, 0x5c, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x52, 0x09, 0x97,
	0x08, 0x44, 0xd0, 0xad, 0x34, 0x27, 0x87, 0x80, 0x3d, 0x8a, 0x5c, 0xac, 0xc5, 0x25, 0x95, 0x39,
	0xa9, 0x12, 0xcc, 0x98, 0xb6, 0x40, 0x64, 0x88, 0x75, 0x8a, 0x38, 0x97, 0x28, 0x9a, 0xad, 0x10,
	0xe7, 0x18, 0xb5, 0x33, 0x72, 0x09, 0xfa, 0xa5, 0x96, 0x16, 0x25, 0xe6, 0x04, 0x83, 0xcc, 0x03,
	0x6b, 0x2b, 0x12, 0x32, 0xe2, 0xe2, 0x84, 0xbb, 0x5c, 0x48, 0x50, 0x0f, 0x3d, 0x50, 0xa4, 0x84,
	0xf4, 0x30, 0x3c, 0x26, 0x64, 0xc7, 0xc5, 0x8b, 0x62, 0x85, 0x90, 0xa8, 0x1e, 0x36, 0x8f, 0x4a,
	0x89, 0xe9, 0x61, 0x75, 0x89, 0x93, 0x32, 0x97, 0x62, 0x5e, 0x6a, 0x89, 0x5e, 0x5a, 0x51, 0x62,
	0x5e, 0x72, 0x46, 0xa9, 0x5e, 0x1e, 0xd8, 0x51, 0x60, 0x4f, 0x26, 0x16, 0x95, 0x00, 0xe3, 0x27,
	0x2b, 0x35, 0xb9, 0x24, 0x89, 0x0d, 0x1c, 0x51, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c,
	0x7b, 0x70, 0xda, 0xc5, 0x01, 0x00, 0x00,
}