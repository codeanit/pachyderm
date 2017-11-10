// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/worker/worker_service.proto

/*
	Package worker is a generated protocol buffer package.

	It is generated from these files:
		server/worker/worker_service.proto

	It has these top-level messages:
		Input
		CancelRequest
		CancelResponse
		ChunkState
		Chunks
*/
package worker

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pfs "github.com/pachyderm/pachyderm/src/client/pfs"
import pps "github.com/pachyderm/pachyderm/src/client/pps"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf "github.com/gogo/protobuf/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ChunkState_State int32

const (
	ChunkState_RUNNING  ChunkState_State = 0
	ChunkState_COMPLETE ChunkState_State = 1
	ChunkState_FAILED   ChunkState_State = 3
)

var ChunkState_State_name = map[int32]string{
	0: "RUNNING",
	1: "COMPLETE",
	3: "FAILED",
}
var ChunkState_State_value = map[string]int32{
	"RUNNING":  0,
	"COMPLETE": 1,
	"FAILED":   3,
}

func (x ChunkState_State) String() string {
	return proto.EnumName(ChunkState_State_name, int32(x))
}
func (ChunkState_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorWorkerService, []int{3, 0}
}

type Input struct {
	FileInfo     *pfs.FileInfo `protobuf:"bytes,1,opt,name=file_info,json=fileInfo" json:"file_info,omitempty"`
	ParentCommit *pfs.Commit   `protobuf:"bytes,5,opt,name=parent_commit,json=parentCommit" json:"parent_commit,omitempty"`
	Name         string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lazy         bool          `protobuf:"varint,3,opt,name=lazy,proto3" json:"lazy,omitempty"`
	Branch       string        `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
	GitURL       string        `protobuf:"bytes,6,opt,name=git_url,json=gitUrl,proto3" json:"git_url,omitempty"`
}

func (m *Input) Reset()                    { *m = Input{} }
func (m *Input) String() string            { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()               {}
func (*Input) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{0} }

func (m *Input) GetFileInfo() *pfs.FileInfo {
	if m != nil {
		return m.FileInfo
	}
	return nil
}

func (m *Input) GetParentCommit() *pfs.Commit {
	if m != nil {
		return m.ParentCommit
	}
	return nil
}

func (m *Input) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Input) GetLazy() bool {
	if m != nil {
		return m.Lazy
	}
	return false
}

func (m *Input) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *Input) GetGitURL() string {
	if m != nil {
		return m.GitURL
	}
	return ""
}

type CancelRequest struct {
	JobID       string   `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	DataFilters []string `protobuf:"bytes,1,rep,name=data_filters,json=dataFilters" json:"data_filters,omitempty"`
}

func (m *CancelRequest) Reset()                    { *m = CancelRequest{} }
func (m *CancelRequest) String() string            { return proto.CompactTextString(m) }
func (*CancelRequest) ProtoMessage()               {}
func (*CancelRequest) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{1} }

func (m *CancelRequest) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *CancelRequest) GetDataFilters() []string {
	if m != nil {
		return m.DataFilters
	}
	return nil
}

type CancelResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *CancelResponse) Reset()                    { *m = CancelResponse{} }
func (m *CancelResponse) String() string            { return proto.CompactTextString(m) }
func (*CancelResponse) ProtoMessage()               {}
func (*CancelResponse) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{2} }

func (m *CancelResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ChunkState struct {
	State   ChunkState_State `protobuf:"varint,1,opt,name=state,proto3,enum=worker.ChunkState_State" json:"state,omitempty"`
	DatumID string           `protobuf:"bytes,2,opt,name=datum_id,json=datumId,proto3" json:"datum_id,omitempty"`
}

func (m *ChunkState) Reset()                    { *m = ChunkState{} }
func (m *ChunkState) String() string            { return proto.CompactTextString(m) }
func (*ChunkState) ProtoMessage()               {}
func (*ChunkState) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{3} }

func (m *ChunkState) GetState() ChunkState_State {
	if m != nil {
		return m.State
	}
	return ChunkState_RUNNING
}

func (m *ChunkState) GetDatumID() string {
	if m != nil {
		return m.DatumID
	}
	return ""
}

type Chunks struct {
	Chunks []int64 `protobuf:"varint,1,rep,packed,name=chunks" json:"chunks,omitempty"`
}

func (m *Chunks) Reset()                    { *m = Chunks{} }
func (m *Chunks) String() string            { return proto.CompactTextString(m) }
func (*Chunks) ProtoMessage()               {}
func (*Chunks) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{4} }

func (m *Chunks) GetChunks() []int64 {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func init() {
	proto.RegisterType((*Input)(nil), "worker.Input")
	proto.RegisterType((*CancelRequest)(nil), "worker.CancelRequest")
	proto.RegisterType((*CancelResponse)(nil), "worker.CancelResponse")
	proto.RegisterType((*ChunkState)(nil), "worker.ChunkState")
	proto.RegisterType((*Chunks)(nil), "worker.Chunks")
	proto.RegisterEnum("worker.ChunkState_State", ChunkState_State_name, ChunkState_State_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Worker service

type WorkerClient interface {
	Status(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*pps.WorkerStatus, error)
	Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) Status(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*pps.WorkerStatus, error) {
	out := new(pps.WorkerStatus)
	err := grpc.Invoke(ctx, "/worker.Worker/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error) {
	out := new(CancelResponse)
	err := grpc.Invoke(ctx, "/worker.Worker/Cancel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Worker service

type WorkerServer interface {
	Status(context.Context, *google_protobuf.Empty) (*pps.WorkerStatus, error)
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.Worker/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Status(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.Worker/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Cancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "worker.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Worker_Status_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Worker_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/worker/worker_service.proto",
}

func (m *Input) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Input) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FileInfo != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(m.FileInfo.Size()))
		n1, err := m.FileInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Lazy {
		dAtA[i] = 0x18
		i++
		if m.Lazy {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Branch) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(len(m.Branch)))
		i += copy(dAtA[i:], m.Branch)
	}
	if m.ParentCommit != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(m.ParentCommit.Size()))
		n2, err := m.ParentCommit.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.GitURL) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(len(m.GitURL)))
		i += copy(dAtA[i:], m.GitURL)
	}
	return i, nil
}

func (m *CancelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DataFilters) > 0 {
		for _, s := range m.DataFilters {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.JobID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(len(m.JobID)))
		i += copy(dAtA[i:], m.JobID)
	}
	return i, nil
}

func (m *CancelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Success {
		dAtA[i] = 0x8
		i++
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *ChunkState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChunkState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.State != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(m.State))
	}
	if len(m.DatumID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(len(m.DatumID)))
		i += copy(dAtA[i:], m.DatumID)
	}
	return i, nil
}

func (m *Chunks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Chunks) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Chunks) > 0 {
		dAtA4 := make([]byte, len(m.Chunks)*10)
		var j3 int
		for _, num1 := range m.Chunks {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(j3))
		i += copy(dAtA[i:], dAtA4[:j3])
	}
	return i, nil
}

func encodeVarintWorkerService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Input) Size() (n int) {
	var l int
	_ = l
	if m.FileInfo != nil {
		l = m.FileInfo.Size()
		n += 1 + l + sovWorkerService(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWorkerService(uint64(l))
	}
	if m.Lazy {
		n += 2
	}
	l = len(m.Branch)
	if l > 0 {
		n += 1 + l + sovWorkerService(uint64(l))
	}
	if m.ParentCommit != nil {
		l = m.ParentCommit.Size()
		n += 1 + l + sovWorkerService(uint64(l))
	}
	l = len(m.GitURL)
	if l > 0 {
		n += 1 + l + sovWorkerService(uint64(l))
	}
	return n
}

func (m *CancelRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.DataFilters) > 0 {
		for _, s := range m.DataFilters {
			l = len(s)
			n += 1 + l + sovWorkerService(uint64(l))
		}
	}
	l = len(m.JobID)
	if l > 0 {
		n += 1 + l + sovWorkerService(uint64(l))
	}
	return n
}

func (m *CancelResponse) Size() (n int) {
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func (m *ChunkState) Size() (n int) {
	var l int
	_ = l
	if m.State != 0 {
		n += 1 + sovWorkerService(uint64(m.State))
	}
	l = len(m.DatumID)
	if l > 0 {
		n += 1 + l + sovWorkerService(uint64(l))
	}
	return n
}

func (m *Chunks) Size() (n int) {
	var l int
	_ = l
	if len(m.Chunks) > 0 {
		l = 0
		for _, e := range m.Chunks {
			l += sovWorkerService(uint64(e))
		}
		n += 1 + sovWorkerService(uint64(l)) + l
	}
	return n
}

func sovWorkerService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWorkerService(x uint64) (n int) {
	return sovWorkerService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Input) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Input: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Input: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FileInfo == nil {
				m.FileInfo = &pfs.FileInfo{}
			}
			if err := m.FileInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lazy", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Lazy = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Branch", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Branch = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentCommit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ParentCommit == nil {
				m.ParentCommit = &pfs.Commit{}
			}
			if err := m.ParentCommit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GitURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GitURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CancelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CancelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataFilters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataFilters = append(m.DataFilters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CancelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CancelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipWorkerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ChunkState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChunkState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChunkState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (ChunkState_State(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DatumID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DatumID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Chunks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Chunks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Chunks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWorkerService
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Chunks = append(m.Chunks, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWorkerService
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthWorkerService
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWorkerService
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Chunks = append(m.Chunks, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWorkerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWorkerService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorkerService
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
					return 0, ErrIntOverflowWorkerService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWorkerService
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthWorkerService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWorkerService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipWorkerService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthWorkerService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorkerService   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("server/worker/worker_service.proto", fileDescriptorWorkerService) }

var fileDescriptorWorkerService = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x49, 0xbd, 0x49, 0x26, 0x6d, 0x15, 0x56, 0x50, 0x59, 0x45, 0x4a, 0x8d, 0x91, 0x50,
	0xd4, 0x83, 0x83, 0x8a, 0x38, 0x70, 0xa4, 0xf9, 0xa9, 0x8c, 0x42, 0x40, 0x4b, 0x23, 0x8e, 0x96,
	0xed, 0xac, 0x5d, 0xb7, 0x8e, 0xd7, 0x78, 0xd7, 0xa0, 0xf6, 0x39, 0x38, 0xf0, 0x48, 0x88, 0x13,
	0x4f, 0x10, 0x21, 0xf3, 0x22, 0x68, 0x77, 0x93, 0x06, 0x71, 0x58, 0xfb, 0x9b, 0x6f, 0x3e, 0xcd,
	0xb7, 0x3b, 0x33, 0xe0, 0x70, 0x5a, 0x7e, 0xa1, 0xe5, 0xf0, 0x2b, 0x2b, 0x6f, 0xee, 0x7f, 0xbe,
	0x24, 0xd3, 0x88, 0xba, 0x45, 0xc9, 0x04, 0xc3, 0x48, 0xb3, 0xc7, 0x8f, 0xa2, 0x2c, 0xa5, 0xb9,
	0x18, 0x16, 0x31, 0x97, 0x47, 0x67, 0x77, 0x6c, 0xc1, 0xe5, 0xd9, 0xb2, 0x09, 0x4b, 0x98, 0x82,
	0x43, 0x89, 0x36, 0xec, 0x93, 0x84, 0xb1, 0x24, 0xa3, 0x43, 0x15, 0x85, 0x55, 0x3c, 0xa4, 0xab,
	0x42, 0xdc, 0xea, 0xa4, 0xf3, 0xd3, 0x00, 0xd3, 0xcb, 0x8b, 0x4a, 0xe0, 0x53, 0xe8, 0xc4, 0x69,
	0x46, 0xfd, 0x34, 0x8f, 0x99, 0x65, 0xd8, 0xc6, 0xa0, 0x7b, 0x76, 0xe0, 0x4a, 0xc7, 0x69, 0x9a,
	0x51, 0x2f, 0x8f, 0x19, 0x69, 0xc7, 0x1b, 0x84, 0x31, 0xec, 0xe5, 0xc1, 0x8a, 0x5a, 0x0f, 0x6c,
	0x63, 0xd0, 0x21, 0x0a, 0x4b, 0x2e, 0x0b, 0xee, 0x6e, 0xad, 0xa6, 0x6d, 0x0c, 0xda, 0x44, 0x61,
	0x7c, 0x04, 0x28, 0x2c, 0x83, 0x3c, 0xba, 0xb2, 0xf6, 0x94, 0x72, 0x13, 0xe1, 0x17, 0x70, 0x50,
	0x04, 0x25, 0xcd, 0x85, 0x1f, 0xb1, 0xd5, 0x2a, 0x15, 0x96, 0xa9, 0xfc, 0xba, 0xca, 0x6f, 0xa4,
	0x28, 0xb2, 0xaf, 0x15, 0x3a, 0xc2, 0xcf, 0xa0, 0x95, 0xa4, 0xc2, 0xaf, 0xca, 0xcc, 0x42, 0xb2,
	0xd4, 0x39, 0xd4, 0xeb, 0x13, 0x74, 0x91, 0x8a, 0x05, 0x99, 0x11, 0x94, 0xa4, 0x62, 0x51, 0x66,
	0xce, 0x25, 0x1c, 0x8c, 0x82, 0x3c, 0xa2, 0x19, 0xa1, 0x9f, 0x2b, 0xca, 0x05, 0x7e, 0x0a, 0xfb,
	0xcb, 0x40, 0x04, 0x7e, 0x9c, 0x66, 0x82, 0x96, 0xdc, 0x32, 0xec, 0xe6, 0xa0, 0x43, 0xba, 0x92,
	0x9b, 0x6a, 0x0a, 0xdb, 0x80, 0xae, 0x59, 0xe8, 0xa7, 0x4b, 0xfd, 0x98, 0xf3, 0x4e, 0xbd, 0x3e,
	0x31, 0xdf, 0xb2, 0xd0, 0x1b, 0x13, 0xf3, 0x9a, 0x85, 0xde, 0xd2, 0x39, 0x85, 0xc3, 0x6d, 0x55,
	0x5e, 0xb0, 0x9c, 0x53, 0x6c, 0x41, 0x8b, 0x57, 0x51, 0x44, 0x39, 0x57, 0x8d, 0x6a, 0x93, 0x6d,
	0xe8, 0x7c, 0x33, 0x00, 0x46, 0x57, 0x55, 0x7e, 0xf3, 0x51, 0x04, 0x82, 0x62, 0x17, 0x4c, 0x2e,
	0x81, 0x92, 0x1d, 0x9e, 0x59, 0xae, 0x1e, 0xaa, 0xbb, 0x93, 0xb8, 0xea, 0x4b, 0xb4, 0x0c, 0x3f,
	0x87, 0xf6, 0x32, 0x10, 0xd5, 0x6a, 0x77, 0x9d, 0x6e, 0xbd, 0x3e, 0x69, 0x8d, 0x25, 0xe7, 0x8d,
	0x49, 0x4b, 0x25, 0xbd, 0xa5, 0xe3, 0x82, 0xa9, 0x0d, 0xba, 0xd0, 0x22, 0x8b, 0xf9, 0xdc, 0x9b,
	0x5f, 0xf4, 0x1a, 0x78, 0x1f, 0xda, 0xa3, 0xf7, 0xef, 0x3e, 0xcc, 0x26, 0x97, 0x93, 0x9e, 0x81,
	0x01, 0xd0, 0xf4, 0x8d, 0x37, 0x9b, 0x8c, 0x7b, 0x4d, 0xc7, 0x06, 0xa4, 0x2c, 0xb9, 0x9c, 0x48,
	0xa4, 0x90, 0xea, 0x45, 0x93, 0x6c, 0xa2, 0xb3, 0x3b, 0x40, 0x9f, 0xd4, 0xdd, 0xf0, 0x2b, 0x40,
	0xb2, 0x76, 0xc5, 0xf1, 0x91, 0xab, 0x37, 0xc7, 0xdd, 0x6e, 0x8e, 0x3b, 0x91, 0x9b, 0x73, 0xfc,
	0xd0, 0x95, 0x2b, 0xa7, 0xe5, 0x5a, 0xea, 0x34, 0xf0, 0x6b, 0x40, 0xba, 0x4b, 0xf8, 0xf1, 0xfd,
	0x2b, 0xff, 0x9d, 0xc5, 0xf1, 0xd1, 0xff, 0xb4, 0x6e, 0xa6, 0xd3, 0x38, 0xef, 0xfd, 0xa8, 0xfb,
	0xc6, 0xaf, 0xba, 0x6f, 0xfc, 0xae, 0xfb, 0xc6, 0xf7, 0x3f, 0xfd, 0x46, 0x88, 0x94, 0xe3, 0xcb,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x22, 0x54, 0x70, 0x1d, 0x29, 0x03, 0x00, 0x00,
}
