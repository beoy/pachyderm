// Code generated by protoc-gen-go.
// source: server/pps/persist/persist.proto
// DO NOT EDIT!

/*
Package persist is a generated protocol buffer package.

It is generated from these files:
	server/pps/persist/persist.proto

It has these top-level messages:
	JobInfo
	JobInfos
	JobOutput
	JobState
	PipelineInfo
	PipelineInfos
*/
package persist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"
import google_protobuf1 "go.pedge.io/pb/go/google/protobuf"
import pfs "github.com/pachyderm/pachyderm/src/client/pfs"
import pachyderm_pps "github.com/pachyderm/pachyderm/src/client/pps"

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
const _ = proto.ProtoPackageIsVersion1

type JobInfo struct {
	JobID        string                      `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	Transform    *pachyderm_pps.Transform    `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	PipelineName string                      `protobuf:"bytes,3,opt,name=pipeline_name,json=pipelineName" json:"pipeline_name,omitempty"`
	Shards       uint64                      `protobuf:"varint,4,opt,name=shards" json:"shards,omitempty"`
	Inputs       []*pachyderm_pps.JobInput   `protobuf:"bytes,5,rep,name=inputs" json:"inputs,omitempty"`
	ParentJob    *pachyderm_pps.Job          `protobuf:"bytes,6,opt,name=parent_job,json=parentJob" json:"parent_job,omitempty"`
	CreatedAt    *google_protobuf1.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	OutputCommit *pfs.Commit                 `protobuf:"bytes,8,opt,name=output_commit,json=outputCommit" json:"output_commit,omitempty"`
	State        pachyderm_pps.JobState      `protobuf:"varint,9,opt,name=state,enum=pachyderm.pps.JobState" json:"state,omitempty"`
	CommitIndex  string                      `protobuf:"bytes,10,opt,name=commit_index,json=commitIndex" json:"commit_index,omitempty"`
}

func (m *JobInfo) Reset()                    { *m = JobInfo{} }
func (m *JobInfo) String() string            { return proto.CompactTextString(m) }
func (*JobInfo) ProtoMessage()               {}
func (*JobInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *JobInfo) GetTransform() *pachyderm_pps.Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *JobInfo) GetInputs() []*pachyderm_pps.JobInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *JobInfo) GetParentJob() *pachyderm_pps.Job {
	if m != nil {
		return m.ParentJob
	}
	return nil
}

func (m *JobInfo) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *JobInfo) GetOutputCommit() *pfs.Commit {
	if m != nil {
		return m.OutputCommit
	}
	return nil
}

type JobInfos struct {
	JobInfo []*JobInfo `protobuf:"bytes,1,rep,name=job_info,json=jobInfo" json:"job_info,omitempty"`
}

func (m *JobInfos) Reset()                    { *m = JobInfos{} }
func (m *JobInfos) String() string            { return proto.CompactTextString(m) }
func (*JobInfos) ProtoMessage()               {}
func (*JobInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JobInfos) GetJobInfo() []*JobInfo {
	if m != nil {
		return m.JobInfo
	}
	return nil
}

type JobOutput struct {
	JobID        string      `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	OutputCommit *pfs.Commit `protobuf:"bytes,2,opt,name=output_commit,json=outputCommit" json:"output_commit,omitempty"`
}

func (m *JobOutput) Reset()                    { *m = JobOutput{} }
func (m *JobOutput) String() string            { return proto.CompactTextString(m) }
func (*JobOutput) ProtoMessage()               {}
func (*JobOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *JobOutput) GetOutputCommit() *pfs.Commit {
	if m != nil {
		return m.OutputCommit
	}
	return nil
}

type JobState struct {
	JobID string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	State pachyderm_pps.JobState `protobuf:"varint,2,opt,name=state,enum=pachyderm.pps.JobState" json:"state,omitempty"`
}

func (m *JobState) Reset()                    { *m = JobState{} }
func (m *JobState) String() string            { return proto.CompactTextString(m) }
func (*JobState) ProtoMessage()               {}
func (*JobState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type PipelineInfo struct {
	PipelineName string                         `protobuf:"bytes,1,opt,name=pipeline_name,json=pipelineName" json:"pipeline_name,omitempty"`
	Transform    *pachyderm_pps.Transform       `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Shards       uint64                         `protobuf:"varint,3,opt,name=shards" json:"shards,omitempty"`
	Inputs       []*pachyderm_pps.PipelineInput `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty"`
	OutputRepo   *pfs.Repo                      `protobuf:"bytes,5,opt,name=output_repo,json=outputRepo" json:"output_repo,omitempty"`
	CreatedAt    *google_protobuf1.Timestamp    `protobuf:"bytes,6,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *PipelineInfo) Reset()                    { *m = PipelineInfo{} }
func (m *PipelineInfo) String() string            { return proto.CompactTextString(m) }
func (*PipelineInfo) ProtoMessage()               {}
func (*PipelineInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PipelineInfo) GetTransform() *pachyderm_pps.Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *PipelineInfo) GetInputs() []*pachyderm_pps.PipelineInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PipelineInfo) GetOutputRepo() *pfs.Repo {
	if m != nil {
		return m.OutputRepo
	}
	return nil
}

func (m *PipelineInfo) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type PipelineInfos struct {
	PipelineInfo []*PipelineInfo `protobuf:"bytes,1,rep,name=pipeline_info,json=pipelineInfo" json:"pipeline_info,omitempty"`
}

func (m *PipelineInfos) Reset()                    { *m = PipelineInfos{} }
func (m *PipelineInfos) String() string            { return proto.CompactTextString(m) }
func (*PipelineInfos) ProtoMessage()               {}
func (*PipelineInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PipelineInfos) GetPipelineInfo() []*PipelineInfo {
	if m != nil {
		return m.PipelineInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*JobInfo)(nil), "pachyderm.pps.persist.JobInfo")
	proto.RegisterType((*JobInfos)(nil), "pachyderm.pps.persist.JobInfos")
	proto.RegisterType((*JobOutput)(nil), "pachyderm.pps.persist.JobOutput")
	proto.RegisterType((*JobState)(nil), "pachyderm.pps.persist.JobState")
	proto.RegisterType((*PipelineInfo)(nil), "pachyderm.pps.persist.PipelineInfo")
	proto.RegisterType((*PipelineInfos)(nil), "pachyderm.pps.persist.PipelineInfos")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for API service

type APIClient interface {
	// Job rpcs
	// job_id cannot be set
	// timestamp cannot be set
	CreateJobInfo(ctx context.Context, in *JobInfo, opts ...grpc.CallOption) (*JobInfo, error)
	InspectJob(ctx context.Context, in *pachyderm_pps.InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error)
	// ordered by time, latest to earliest
	ListJobInfos(ctx context.Context, in *pachyderm_pps.ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error)
	// should only be called when rolling back if a Job does not start!
	DeleteJobInfo(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// JobOutput rpcs
	CreateJobOutput(ctx context.Context, in *JobOutput, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// JobState rpcs
	CreateJobState(ctx context.Context, in *JobState, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// Pipeline rpcs
	CreatePipelineInfo(ctx context.Context, in *PipelineInfo, opts ...grpc.CallOption) (*PipelineInfo, error)
	GetPipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*PipelineInfo, error)
	// ordered by time, latest to earliest
	ListPipelineInfos(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PipelineInfos, error)
	DeletePipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) CreateJobInfo(ctx context.Context, in *JobInfo, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreateJobInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) InspectJob(ctx context.Context, in *pachyderm_pps.InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/InspectJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListJobInfos(ctx context.Context, in *pachyderm_pps.ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error) {
	out := new(JobInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/ListJobInfos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteJobInfo(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/DeleteJobInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateJobOutput(ctx context.Context, in *JobOutput, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreateJobOutput", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateJobState(ctx context.Context, in *JobState, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreateJobState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreatePipelineInfo(ctx context.Context, in *PipelineInfo, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreatePipelineInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetPipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/GetPipelineInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListPipelineInfos(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PipelineInfos, error) {
	out := new(PipelineInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/ListPipelineInfos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeletePipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/DeletePipelineInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	// Job rpcs
	// job_id cannot be set
	// timestamp cannot be set
	CreateJobInfo(context.Context, *JobInfo) (*JobInfo, error)
	InspectJob(context.Context, *pachyderm_pps.InspectJobRequest) (*JobInfo, error)
	// ordered by time, latest to earliest
	ListJobInfos(context.Context, *pachyderm_pps.ListJobRequest) (*JobInfos, error)
	// should only be called when rolling back if a Job does not start!
	DeleteJobInfo(context.Context, *pachyderm_pps.Job) (*google_protobuf.Empty, error)
	// JobOutput rpcs
	CreateJobOutput(context.Context, *JobOutput) (*google_protobuf.Empty, error)
	// JobState rpcs
	CreateJobState(context.Context, *JobState) (*google_protobuf.Empty, error)
	// Pipeline rpcs
	CreatePipelineInfo(context.Context, *PipelineInfo) (*PipelineInfo, error)
	GetPipelineInfo(context.Context, *pachyderm_pps.Pipeline) (*PipelineInfo, error)
	// ordered by time, latest to earliest
	ListPipelineInfos(context.Context, *google_protobuf.Empty) (*PipelineInfos, error)
	DeletePipelineInfo(context.Context, *pachyderm_pps.Pipeline) (*google_protobuf.Empty, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_CreateJobInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateJobInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_InspectJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.InspectJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).InspectJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_ListJobInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.ListJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).ListJobInfos(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_DeleteJobInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).DeleteJobInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_CreateJobOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobOutput)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateJobOutput(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_CreateJobState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobState)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateJobState(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_CreatePipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PipelineInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreatePipelineInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetPipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetPipelineInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_ListPipelineInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).ListPipelineInfos(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_DeletePipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).DeletePipelineInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.persist.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJobInfo",
			Handler:    _API_CreateJobInfo_Handler,
		},
		{
			MethodName: "InspectJob",
			Handler:    _API_InspectJob_Handler,
		},
		{
			MethodName: "ListJobInfos",
			Handler:    _API_ListJobInfos_Handler,
		},
		{
			MethodName: "DeleteJobInfo",
			Handler:    _API_DeleteJobInfo_Handler,
		},
		{
			MethodName: "CreateJobOutput",
			Handler:    _API_CreateJobOutput_Handler,
		},
		{
			MethodName: "CreateJobState",
			Handler:    _API_CreateJobState_Handler,
		},
		{
			MethodName: "CreatePipelineInfo",
			Handler:    _API_CreatePipelineInfo_Handler,
		},
		{
			MethodName: "GetPipelineInfo",
			Handler:    _API_GetPipelineInfo_Handler,
		},
		{
			MethodName: "ListPipelineInfos",
			Handler:    _API_ListPipelineInfos_Handler,
		},
		{
			MethodName: "DeletePipelineInfo",
			Handler:    _API_DeletePipelineInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 700 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x7e, 0x9b, 0xb4, 0x4e, 0x3d, 0x49, 0x5a, 0xbd, 0x2b, 0x68, 0x2d, 0xf3, 0xd1, 0xe0, 0x72,
	0x40, 0x48, 0x38, 0x50, 0x10, 0x12, 0x07, 0x0e, 0xa5, 0x54, 0x25, 0x7c, 0xa6, 0x6e, 0x2f, 0x70,
	0x09, 0x4e, 0xb2, 0x69, 0x5d, 0xc5, 0xf6, 0xe2, 0xdd, 0x20, 0xfa, 0x6f, 0x38, 0xf3, 0x0b, 0xf8,
	0x79, 0x8c, 0x77, 0xed, 0x7c, 0x38, 0x71, 0x62, 0x38, 0x44, 0xf1, 0xce, 0xc7, 0x33, 0x33, 0xcf,
	0x3c, 0xbb, 0xd0, 0xe0, 0x34, 0xfa, 0x4e, 0xa3, 0x26, 0x63, 0xbc, 0xc9, 0x68, 0xc4, 0x3d, 0x2e,
	0xd2, 0x7f, 0x9b, 0x45, 0xa1, 0x08, 0xc9, 0x4d, 0xe6, 0xf6, 0x2e, 0xaf, 0xfb, 0x34, 0xf2, 0x6d,
	0x0c, 0xb2, 0x13, 0xa7, 0x79, 0xeb, 0x22, 0x0c, 0x2f, 0x86, 0xb4, 0x29, 0x83, 0xba, 0xa3, 0x41,
	0x93, 0xfa, 0x4c, 0x5c, 0xab, 0x1c, 0x73, 0x2f, 0xeb, 0x14, 0x9e, 0x4f, 0xb9, 0x70, 0x7d, 0x96,
	0x04, 0xdc, 0xe8, 0x0d, 0x3d, 0x1a, 0x60, 0xa9, 0x01, 0x8f, 0x7f, 0x59, 0x6b, 0xdc, 0x0c, 0x4b,
	0xac, 0xd6, 0xef, 0x32, 0x54, 0xde, 0x86, 0xdd, 0x56, 0x30, 0xc0, 0x66, 0x40, 0xbb, 0x0a, 0xbb,
	0x1d, 0xaf, 0x6f, 0xac, 0x35, 0xd6, 0x1e, 0xe8, 0xce, 0x06, 0x9e, 0x5a, 0x7d, 0xf2, 0x1c, 0x74,
	0x11, 0xb9, 0x01, 0x1f, 0x84, 0x91, 0x6f, 0x94, 0xd0, 0x53, 0x3d, 0x30, 0xec, 0xd9, 0xbe, 0xcf,
	0x53, 0xbf, 0x33, 0x09, 0x25, 0xfb, 0x50, 0x67, 0x1e, 0xa3, 0x43, 0x2f, 0xa0, 0x9d, 0xc0, 0xf5,
	0xa9, 0x51, 0x96, 0xa8, 0xb5, 0xd4, 0xf8, 0x11, 0x6d, 0x64, 0x07, 0x34, 0x7e, 0xe9, 0x46, 0x7d,
	0x6e, 0xac, 0xa3, 0x77, 0xdd, 0x49, 0x4e, 0xa4, 0x09, 0x9a, 0x17, 0xb0, 0x91, 0xe0, 0xc6, 0x46,
	0xa3, 0x8c, 0x15, 0x77, 0x33, 0x15, 0x65, 0xcf, 0xe8, 0x77, 0x92, 0x30, 0xf2, 0x04, 0x80, 0xb9,
	0x11, 0x0e, 0xd8, 0xc1, 0xae, 0x0d, 0x4d, 0xb6, 0x49, 0xe6, 0x93, 0x1c, 0x5d, 0x45, 0xe1, 0x27,
	0x79, 0x01, 0xd0, 0x8b, 0xa8, 0x2b, 0x68, 0xbf, 0xe3, 0x0a, 0xa3, 0x22, 0x53, 0x4c, 0x5b, 0xb1,
	0x6b, 0xa7, 0xec, 0xda, 0xe7, 0x29, 0xbb, 0x8e, 0x9e, 0x44, 0x1f, 0x0a, 0xf2, 0x18, 0xea, 0xe1,
	0x48, 0x60, 0xe1, 0x4e, 0x2f, 0xf4, 0x7d, 0x4f, 0x18, 0x9b, 0x32, 0xbb, 0x6a, 0xc7, 0x7c, 0x1f,
	0x49, 0x93, 0x53, 0x53, 0x11, 0xea, 0x44, 0x1e, 0xc1, 0x06, 0xa2, 0x08, 0x6a, 0xe8, 0x18, 0xb9,
	0xb5, 0x68, 0x9e, 0xb3, 0xd8, 0xed, 0xa8, 0x28, 0x72, 0x0f, 0x6a, 0x0a, 0xb9, 0xe3, 0x05, 0x7d,
	0xfa, 0xc3, 0x00, 0xc9, 0x5d, 0x55, 0xd9, 0x5a, 0xb1, 0xc9, 0x3a, 0x86, 0xcd, 0x64, 0x73, 0x1c,
	0x47, 0xd9, 0x94, 0xab, 0xc3, 0x03, 0x2e, 0x2f, 0x26, 0xec, 0xae, 0xbd, 0x50, 0x5a, 0x76, 0x92,
	0xe2, 0x54, 0xae, 0xd4, 0x87, 0x75, 0x0e, 0x3a, 0xda, 0x3e, 0xc9, 0x5e, 0xf3, 0x24, 0x30, 0x37,
	0x6e, 0x69, 0xc5, 0xb8, 0x56, 0x5b, 0x36, 0x27, 0x47, 0xca, 0x03, 0x1d, 0x33, 0x52, 0x2a, 0xc2,
	0x88, 0xf5, 0xb3, 0x04, 0xb5, 0x76, 0x22, 0x1d, 0x29, 0xd7, 0x39, 0x7d, 0xad, 0x2d, 0xd0, 0xd7,
	0xbf, 0x8a, 0x77, 0xa2, 0xcb, 0xf2, 0x8c, 0x2e, 0x9f, 0x8d, 0x75, 0xb9, 0x2e, 0x69, 0xbe, 0x9d,
	0x01, 0x9b, 0x74, 0x38, 0x2d, 0xce, 0x87, 0x50, 0x4d, 0xf8, 0x8b, 0x28, 0x0b, 0x51, 0xd2, 0x71,
	0x1f, 0xba, 0x64, 0xcf, 0x41, 0x83, 0x03, 0xca, 0x1b, 0x7f, 0x67, 0x54, 0xa9, 0xfd, 0x85, 0x2a,
	0xad, 0xcf, 0x50, 0x9f, 0x66, 0x88, 0x93, 0x37, 0x53, 0x14, 0x4d, 0x69, 0x63, 0x3f, 0x47, 0x1b,
	0xd3, 0xc9, 0x13, 0x1e, 0xe3, 0xd3, 0xc1, 0x2f, 0x0d, 0xca, 0x87, 0xed, 0x16, 0x39, 0x85, 0xfa,
	0x91, 0xac, 0x97, 0x3e, 0x1a, 0x2b, 0x74, 0x66, 0xae, 0xf0, 0x5b, 0xff, 0x91, 0x36, 0x40, 0x2b,
	0xe0, 0x8c, 0xf6, 0xe4, 0xa5, 0x6c, 0x64, 0xe2, 0x27, 0x2e, 0x87, 0x7e, 0x1b, 0xe1, 0xcc, 0x85,
	0x10, 0x6b, 0xef, 0xd1, 0x32, 0xbe, 0x1d, 0x77, 0x32, 0x19, 0x89, 0x33, 0x05, 0xdc, 0x5b, 0x0e,
	0xc8, 0x11, 0xf1, 0x25, 0xd4, 0x5f, 0xd3, 0x21, 0x9d, 0x8c, 0xbd, 0xe0, 0x69, 0x31, 0x77, 0xe6,
	0xb6, 0x74, 0x1c, 0x3f, 0xdb, 0x98, 0xfe, 0x01, 0xb6, 0xc7, 0xac, 0x25, 0x37, 0xad, 0x91, 0x5f,
	0x54, 0x45, 0x2c, 0x81, 0x7b, 0x07, 0x5b, 0x63, 0x38, 0x75, 0xc5, 0x96, 0x8c, 0x20, 0x03, 0x96,
	0x80, 0x7d, 0x05, 0xa2, 0xc0, 0x66, 0x2f, 0x57, 0x01, 0x89, 0x98, 0x45, 0x82, 0xb0, 0xc2, 0x29,
	0x6c, 0x9f, 0x50, 0x31, 0x03, 0xbf, 0x9b, 0x73, 0x6d, 0x8a, 0x42, 0x9e, 0xc1, 0xff, 0xf1, 0x12,
	0x67, 0xd5, 0x9e, 0x33, 0xa3, 0x79, 0xbf, 0x00, 0x66, 0xbc, 0xe4, 0x13, 0x20, 0x6a, 0xc9, 0xc5,
	0x5a, 0xcd, 0xa5, 0xf4, 0x95, 0xfe, 0xa5, 0x92, 0xd4, 0xe8, 0x6a, 0xd2, 0xf9, 0xf4, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x7e, 0x36, 0xef, 0x26, 0x0b, 0x08, 0x00, 0x00,
}
