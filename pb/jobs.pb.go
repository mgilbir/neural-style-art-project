// Code generated by protoc-gen-go.
// source: jobs.proto
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

type JobRequest struct {
}

func (m *JobRequest) Reset()                    { *m = JobRequest{} }
func (m *JobRequest) String() string            { return proto.CompactTextString(m) }
func (*JobRequest) ProtoMessage()               {}
func (*JobRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type JobAck struct {
}

func (m *JobAck) Reset()                    { *m = JobAck{} }
func (m *JobAck) String() string            { return proto.CompactTextString(m) }
func (*JobAck) ProtoMessage()               {}
func (*JobAck) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type Job struct {
	Id      string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name    string      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Style   *InputImage `protobuf:"bytes,4,opt,name=style" json:"style,omitempty"`
	Content *InputImage `protobuf:"bytes,5,opt,name=content" json:"content,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *Job) GetStyle() *InputImage {
	if m != nil {
		return m.Style
	}
	return nil
}

func (m *Job) GetContent() *InputImage {
	if m != nil {
		return m.Content
	}
	return nil
}

type JobResult struct {
	Id            string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name          string      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ProgressCount int32       `protobuf:"varint,3,opt,name=progress_count" json:"progress_count,omitempty"`
	Format        ImageFormat `protobuf:"varint,4,opt,name=format,enum=ImageFormat" json:"format,omitempty"`
	Image         []byte      `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
}

func (m *JobResult) Reset()                    { *m = JobResult{} }
func (m *JobResult) String() string            { return proto.CompactTextString(m) }
func (*JobResult) ProtoMessage()               {}
func (*JobResult) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type JobResultResponse struct {
}

func (m *JobResultResponse) Reset()                    { *m = JobResultResponse{} }
func (m *JobResultResponse) String() string            { return proto.CompactTextString(m) }
func (*JobResultResponse) ProtoMessage()               {}
func (*JobResultResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

type JobFail struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *JobFail) Reset()                    { *m = JobFail{} }
func (m *JobFail) String() string            { return proto.CompactTextString(m) }
func (*JobFail) ProtoMessage()               {}
func (*JobFail) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

type JobProgressResponse struct {
}

func (m *JobProgressResponse) Reset()                    { *m = JobProgressResponse{} }
func (m *JobProgressResponse) String() string            { return proto.CompactTextString(m) }
func (*JobProgressResponse) ProtoMessage()               {}
func (*JobProgressResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func init() {
	proto.RegisterType((*JobRequest)(nil), "JobRequest")
	proto.RegisterType((*JobAck)(nil), "JobAck")
	proto.RegisterType((*Job)(nil), "Job")
	proto.RegisterType((*JobResult)(nil), "JobResult")
	proto.RegisterType((*JobResultResponse)(nil), "JobResultResponse")
	proto.RegisterType((*JobFail)(nil), "JobFail")
	proto.RegisterType((*JobProgressResponse)(nil), "JobProgressResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for NeuralStyleWorker service

type NeuralStyleWorkerClient interface {
	RequestJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*Job, error)
	AcknowledgeJob(ctx context.Context, in *JobAck, opts ...grpc.CallOption) (*JobAck, error)
	ProgressReport(ctx context.Context, in *JobResult, opts ...grpc.CallOption) (*JobProgressResponse, error)
	CompleteJob(ctx context.Context, in *JobResult, opts ...grpc.CallOption) (*JobResultResponse, error)
	FailJob(ctx context.Context, in *JobFail, opts ...grpc.CallOption) (*JobFail, error)
}

type neuralStyleWorkerClient struct {
	cc *grpc.ClientConn
}

func NewNeuralStyleWorkerClient(cc *grpc.ClientConn) NeuralStyleWorkerClient {
	return &neuralStyleWorkerClient{cc}
}

func (c *neuralStyleWorkerClient) RequestJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/NeuralStyleWorker/RequestJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neuralStyleWorkerClient) AcknowledgeJob(ctx context.Context, in *JobAck, opts ...grpc.CallOption) (*JobAck, error) {
	out := new(JobAck)
	err := grpc.Invoke(ctx, "/NeuralStyleWorker/AcknowledgeJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neuralStyleWorkerClient) ProgressReport(ctx context.Context, in *JobResult, opts ...grpc.CallOption) (*JobProgressResponse, error) {
	out := new(JobProgressResponse)
	err := grpc.Invoke(ctx, "/NeuralStyleWorker/ProgressReport", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neuralStyleWorkerClient) CompleteJob(ctx context.Context, in *JobResult, opts ...grpc.CallOption) (*JobResultResponse, error) {
	out := new(JobResultResponse)
	err := grpc.Invoke(ctx, "/NeuralStyleWorker/CompleteJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neuralStyleWorkerClient) FailJob(ctx context.Context, in *JobFail, opts ...grpc.CallOption) (*JobFail, error) {
	out := new(JobFail)
	err := grpc.Invoke(ctx, "/NeuralStyleWorker/FailJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NeuralStyleWorker service

type NeuralStyleWorkerServer interface {
	RequestJob(context.Context, *JobRequest) (*Job, error)
	AcknowledgeJob(context.Context, *JobAck) (*JobAck, error)
	ProgressReport(context.Context, *JobResult) (*JobProgressResponse, error)
	CompleteJob(context.Context, *JobResult) (*JobResultResponse, error)
	FailJob(context.Context, *JobFail) (*JobFail, error)
}

func RegisterNeuralStyleWorkerServer(s *grpc.Server, srv NeuralStyleWorkerServer) {
	s.RegisterService(&_NeuralStyleWorker_serviceDesc, srv)
}

func _NeuralStyleWorker_RequestJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NeuralStyleWorkerServer).RequestJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _NeuralStyleWorker_AcknowledgeJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobAck)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NeuralStyleWorkerServer).AcknowledgeJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _NeuralStyleWorker_ProgressReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NeuralStyleWorkerServer).ProgressReport(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _NeuralStyleWorker_CompleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NeuralStyleWorkerServer).CompleteJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _NeuralStyleWorker_FailJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobFail)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NeuralStyleWorkerServer).FailJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _NeuralStyleWorker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NeuralStyleWorker",
	HandlerType: (*NeuralStyleWorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestJob",
			Handler:    _NeuralStyleWorker_RequestJob_Handler,
		},
		{
			MethodName: "AcknowledgeJob",
			Handler:    _NeuralStyleWorker_AcknowledgeJob_Handler,
		},
		{
			MethodName: "ProgressReport",
			Handler:    _NeuralStyleWorker_ProgressReport_Handler,
		},
		{
			MethodName: "CompleteJob",
			Handler:    _NeuralStyleWorker_CompleteJob_Handler,
		},
		{
			MethodName: "FailJob",
			Handler:    _NeuralStyleWorker_FailJob_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor2 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x8e, 0xda, 0x30,
	0x10, 0x55, 0x80, 0x10, 0x98, 0xd0, 0x48, 0x18, 0x2a, 0x45, 0x91, 0xaa, 0x16, 0xb7, 0x48, 0xbd,
	0xe0, 0x43, 0xfa, 0x05, 0x2d, 0x12, 0x12, 0x3d, 0x54, 0x55, 0x7a, 0xe8, 0x71, 0x15, 0x82, 0x61,
	0x81, 0xc4, 0xce, 0xda, 0x8e, 0x56, 0xfb, 0x1b, 0xfb, 0x79, 0xfb, 0x35, 0x6b, 0x3b, 0x09, 0xac,
	0xc4, 0x1e, 0xb8, 0x44, 0x33, 0x6f, 0x26, 0x33, 0x6f, 0xde, 0x33, 0xc0, 0x91, 0x6f, 0x24, 0x29,
	0x05, 0x57, 0x3c, 0xf2, 0x0f, 0x45, 0xba, 0xa7, 0x75, 0x82, 0x47, 0x00, 0xbf, 0xf9, 0x26, 0xa1,
	0x0f, 0x15, 0x95, 0x0a, 0x0f, 0xa0, 0xaf, 0xb3, 0x9f, 0xd9, 0x09, 0x4b, 0xe8, 0xea, 0x08, 0x05,
	0xd0, 0x39, 0x6c, 0x43, 0xe7, 0x8b, 0xf3, 0x7d, 0x98, 0xe8, 0x08, 0x21, 0xe8, 0xb1, 0xb4, 0xa0,
	0x61, 0xc7, 0x22, 0x36, 0x46, 0x33, 0x70, 0xa5, 0x7a, 0xca, 0x69, 0xd8, 0xd3, 0xa0, 0x1f, 0xfb,
	0x64, 0xcd, 0xca, 0x4a, 0xad, 0xcd, 0x92, 0xa4, 0xae, 0xa0, 0x39, 0x78, 0x19, 0x67, 0x8a, 0x32,
	0x15, 0xba, 0xd7, 0x4d, 0x6d, 0x0d, 0x3f, 0x3b, 0x30, 0xb4, 0x6c, 0x64, 0x95, 0xab, 0x9b, 0x76,
	0xcf, 0x21, 0xd0, 0x77, 0xec, 0x05, 0x95, 0xf2, 0x2e, 0xe3, 0x95, 0x9e, 0xdf, 0xd5, 0x55, 0x37,
	0xf9, 0xd0, 0xa2, 0x4b, 0x03, 0xa2, 0x6f, 0xd0, 0xdf, 0x71, 0x51, 0xa4, 0xca, 0x72, 0x0c, 0xe2,
	0x11, 0xb1, 0x9b, 0x57, 0x16, 0x4b, 0x9a, 0x1a, 0x9a, 0x82, 0x6b, 0xa5, 0xb1, 0x1c, 0x47, 0x49,
	0x9d, 0xe0, 0x09, 0x8c, 0xcf, 0x9c, 0xf4, 0xb7, 0xe4, 0x4c, 0x52, 0xbc, 0x00, 0x4f, 0x83, 0xab,
	0xf4, 0x90, 0xdf, 0x42, 0x13, 0x7f, 0x84, 0x89, 0x6e, 0xff, 0xdb, 0x70, 0x6a, 0xa7, 0xc4, 0x2f,
	0x0e, 0x8c, 0xff, 0xd0, 0x4a, 0xa4, 0xf9, 0x3f, 0x23, 0xd3, 0x7f, 0x2e, 0x4e, 0x54, 0xa0, 0xcf,
	0x00, 0x8d, 0x1f, 0xc6, 0x01, 0x9f, 0x5c, 0xfc, 0x89, 0x7a, 0x26, 0x41, 0x18, 0x02, 0x6d, 0x11,
	0xe3, 0x8f, 0x39, 0xdd, 0xee, 0xa9, 0x41, 0x3c, 0x52, 0xdb, 0x16, 0xb5, 0x01, 0x8a, 0x21, 0xb8,
	0xac, 0x2b, 0xb9, 0x50, 0x08, 0xc8, 0xf9, 0x8c, 0x68, 0x4a, 0xde, 0xa1, 0x83, 0x16, 0xe0, 0x2f,
	0x79, 0x51, 0xe6, 0x54, 0xd9, 0xa1, 0x6f, 0x7f, 0x40, 0xe4, 0x4a, 0x03, 0xf4, 0x09, 0x3c, 0x23,
	0x80, 0x69, 0x1d, 0x90, 0x46, 0x8d, 0xe8, 0x1c, 0xfd, 0xfa, 0x0a, 0x33, 0x46, 0x15, 0xd9, 0x89,
	0x94, 0x65, 0xf7, 0x15, 0x61, 0xf6, 0x4e, 0xfb, 0x1c, 0x52, 0xa1, 0xb4, 0x3d, 0x47, 0x9a, 0xa9,
	0x4d, 0xdf, 0xbe, 0xc2, 0x1f, 0xaf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x23, 0xa3, 0x3c, 0x9c, 0xa0,
	0x02, 0x00, 0x00,
}