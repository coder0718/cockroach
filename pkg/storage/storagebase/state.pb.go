// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/storage/storagebase/state.proto
// DO NOT EDIT!

package storagebase

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_storage_engine_enginepb "github.com/cockroachdb/cockroach/pkg/storage/engine/enginepb"
import cockroach_roachpb4 "github.com/cockroachdb/cockroach/pkg/roachpb"
import cockroach_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
import cockroach_roachpb1 "github.com/cockroachdb/cockroach/pkg/roachpb"
import cockroach_util_hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ReplicaState is the part of the Range Raft state machine which is cached in
// memory and which is manipulated exclusively through consensus.
//
// The struct is also used to transfer state to Replicas in the context of
// proposer-evaluated Raft, in which case it does not represent a complete
// state but instead an update to be applied to an existing state, with each
// field specified in the update overwriting its counterpart on the receiving
// ReplicaState.
//
// For the ReplicaState persisted on the Replica, all optional fields are
// populated (i.e. no nil pointers or enums with the default value).
type ReplicaState struct {
	// The highest (and last) index applied to the state machine.
	RaftAppliedIndex uint64 `protobuf:"varint,1,opt,name=raft_applied_index,json=raftAppliedIndex,proto3" json:"raft_applied_index,omitempty"`
	// The highest (and last) lease index applied to the state machine.
	LeaseAppliedIndex uint64 `protobuf:"varint,2,opt,name=lease_applied_index,json=leaseAppliedIndex,proto3" json:"lease_applied_index,omitempty"`
	// The Range descriptor.
	// The pointer may change, but the referenced RangeDescriptor struct itself
	// must be treated as immutable; it is leaked out of the lock.
	//
	// Changes of the descriptor should always go through one of the
	// (*Replica).setDesc* methods.
	Desc *cockroach_roachpb.RangeDescriptor `protobuf:"bytes,3,opt,name=desc" json:"desc,omitempty"`
	// The latest range lease.
	//
	// Note that this message is both sent over the network and used to model
	// replica state in memory. In memory (storage.Replica.mu.state), the lease
	// is never nil (and never zero-valued), but it may be nil when sent over
	// the network as part of ReplicatedEvalResult.
	//
	// TODO(tamird): consider changing the type of storage.Replica.mu.state to
	// avoid the nullability there.
	Lease *cockroach_roachpb1.Lease `protobuf:"bytes,4,opt,name=lease" json:"lease,omitempty"`
	// The truncation state of the Raft log.
	TruncatedState *cockroach_roachpb4.RaftTruncatedState `protobuf:"bytes,5,opt,name=truncated_state,json=truncatedState" json:"truncated_state,omitempty"`
	// gcThreshold is the GC threshold of the Range, typically updated when keys
	// are garbage collected. Reads and writes at timestamps <= this time will
	// not be served.
	//
	// TODO(tschottdorf): should be nullable to keep ReplicaState small as we are
	// sending it over the wire. Since we only ever increase gc_threshold, that's
	// the only upshot - fields which can return to the zero value must
	// special-case that value simply because otherwise there's no way of
	// distinguishing "no update" to and updating to the zero value.
	GCThreshold cockroach_util_hlc.Timestamp                `protobuf:"bytes,6,opt,name=gc_threshold,json=gcThreshold" json:"gc_threshold"`
	Stats       cockroach_storage_engine_enginepb.MVCCStats `protobuf:"bytes,7,opt,name=stats" json:"stats"`
	// txn_span_gc_threshold is the (maximum) timestamp below which transaction
	// records may have been garbage collected (as measured by txn.LastActive()).
	// Transaction at lower timestamps must not be allowed to write their initial
	// transaction entry.
	//
	// TODO(tschottdorf): should be nullable; see gc_threshold.
	TxnSpanGCThreshold cockroach_util_hlc.Timestamp `protobuf:"bytes,9,opt,name=txn_span_gc_threshold,json=txnSpanGcThreshold" json:"txn_span_gc_threshold"`
}

func (m *ReplicaState) Reset()                    { *m = ReplicaState{} }
func (m *ReplicaState) String() string            { return proto.CompactTextString(m) }
func (*ReplicaState) ProtoMessage()               {}
func (*ReplicaState) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{0} }

type RangeInfo struct {
	ReplicaState `protobuf:"bytes,1,opt,name=state,embedded=state" json:"state"`
	// The highest (and last) index in the Raft log.
	LastIndex  uint64 `protobuf:"varint,2,opt,name=lastIndex,proto3" json:"lastIndex,omitempty"`
	NumPending uint64 `protobuf:"varint,3,opt,name=num_pending,json=numPending,proto3" json:"num_pending,omitempty"`
	NumDropped uint64 `protobuf:"varint,5,opt,name=num_dropped,json=numDropped,proto3" json:"num_dropped,omitempty"`
	// raft_log_size may be initially inaccurate after a server restart.
	// See storage.Replica.mu.raftLogSize.
	RaftLogSize int64 `protobuf:"varint,6,opt,name=raft_log_size,json=raftLogSize,proto3" json:"raft_log_size,omitempty"`
	// Approximately the amount of quota available.
	ApproximateProposalQuota int64 `protobuf:"varint,7,opt,name=approximate_proposal_quota,json=approximateProposalQuota,proto3" json:"approximate_proposal_quota,omitempty"`
}

func (m *RangeInfo) Reset()                    { *m = RangeInfo{} }
func (m *RangeInfo) String() string            { return proto.CompactTextString(m) }
func (*RangeInfo) ProtoMessage()               {}
func (*RangeInfo) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{1} }

func init() {
	proto.RegisterType((*ReplicaState)(nil), "cockroach.storage.storagebase.ReplicaState")
	proto.RegisterType((*RangeInfo)(nil), "cockroach.storage.storagebase.RangeInfo")
}
func (this *ReplicaState) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ReplicaState)
	if !ok {
		that2, ok := that.(ReplicaState)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.RaftAppliedIndex != that1.RaftAppliedIndex {
		return false
	}
	if this.LeaseAppliedIndex != that1.LeaseAppliedIndex {
		return false
	}
	if !this.Desc.Equal(that1.Desc) {
		return false
	}
	if !this.Lease.Equal(that1.Lease) {
		return false
	}
	if !this.TruncatedState.Equal(that1.TruncatedState) {
		return false
	}
	if !this.GCThreshold.Equal(&that1.GCThreshold) {
		return false
	}
	if !this.Stats.Equal(&that1.Stats) {
		return false
	}
	if !this.TxnSpanGCThreshold.Equal(&that1.TxnSpanGCThreshold) {
		return false
	}
	return true
}
func (this *RangeInfo) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RangeInfo)
	if !ok {
		that2, ok := that.(RangeInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.ReplicaState.Equal(&that1.ReplicaState) {
		return false
	}
	if this.LastIndex != that1.LastIndex {
		return false
	}
	if this.NumPending != that1.NumPending {
		return false
	}
	if this.NumDropped != that1.NumDropped {
		return false
	}
	if this.RaftLogSize != that1.RaftLogSize {
		return false
	}
	if this.ApproximateProposalQuota != that1.ApproximateProposalQuota {
		return false
	}
	return true
}
func (m *ReplicaState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplicaState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RaftAppliedIndex != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintState(dAtA, i, uint64(m.RaftAppliedIndex))
	}
	if m.LeaseAppliedIndex != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintState(dAtA, i, uint64(m.LeaseAppliedIndex))
	}
	if m.Desc != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintState(dAtA, i, uint64(m.Desc.Size()))
		n1, err := m.Desc.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Lease != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintState(dAtA, i, uint64(m.Lease.Size()))
		n2, err := m.Lease.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.TruncatedState != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintState(dAtA, i, uint64(m.TruncatedState.Size()))
		n3, err := m.TruncatedState.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	dAtA[i] = 0x32
	i++
	i = encodeVarintState(dAtA, i, uint64(m.GCThreshold.Size()))
	n4, err := m.GCThreshold.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x3a
	i++
	i = encodeVarintState(dAtA, i, uint64(m.Stats.Size()))
	n5, err := m.Stats.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x4a
	i++
	i = encodeVarintState(dAtA, i, uint64(m.TxnSpanGCThreshold.Size()))
	n6, err := m.TxnSpanGCThreshold.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	return i, nil
}

func (m *RangeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RangeInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintState(dAtA, i, uint64(m.ReplicaState.Size()))
	n7, err := m.ReplicaState.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	if m.LastIndex != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintState(dAtA, i, uint64(m.LastIndex))
	}
	if m.NumPending != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintState(dAtA, i, uint64(m.NumPending))
	}
	if m.NumDropped != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintState(dAtA, i, uint64(m.NumDropped))
	}
	if m.RaftLogSize != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintState(dAtA, i, uint64(m.RaftLogSize))
	}
	if m.ApproximateProposalQuota != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintState(dAtA, i, uint64(m.ApproximateProposalQuota))
	}
	return i, nil
}

func encodeFixed64State(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32State(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ReplicaState) Size() (n int) {
	var l int
	_ = l
	if m.RaftAppliedIndex != 0 {
		n += 1 + sovState(uint64(m.RaftAppliedIndex))
	}
	if m.LeaseAppliedIndex != 0 {
		n += 1 + sovState(uint64(m.LeaseAppliedIndex))
	}
	if m.Desc != nil {
		l = m.Desc.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.Lease != nil {
		l = m.Lease.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.TruncatedState != nil {
		l = m.TruncatedState.Size()
		n += 1 + l + sovState(uint64(l))
	}
	l = m.GCThreshold.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Stats.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.TxnSpanGCThreshold.Size()
	n += 1 + l + sovState(uint64(l))
	return n
}

func (m *RangeInfo) Size() (n int) {
	var l int
	_ = l
	l = m.ReplicaState.Size()
	n += 1 + l + sovState(uint64(l))
	if m.LastIndex != 0 {
		n += 1 + sovState(uint64(m.LastIndex))
	}
	if m.NumPending != 0 {
		n += 1 + sovState(uint64(m.NumPending))
	}
	if m.NumDropped != 0 {
		n += 1 + sovState(uint64(m.NumDropped))
	}
	if m.RaftLogSize != 0 {
		n += 1 + sovState(uint64(m.RaftLogSize))
	}
	if m.ApproximateProposalQuota != 0 {
		n += 1 + sovState(uint64(m.ApproximateProposalQuota))
	}
	return n
}

func sovState(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReplicaState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: ReplicaState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftAppliedIndex", wireType)
			}
			m.RaftAppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftAppliedIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseAppliedIndex", wireType)
			}
			m.LeaseAppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LeaseAppliedIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Desc == nil {
				m.Desc = &cockroach_roachpb.RangeDescriptor{}
			}
			if err := m.Desc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lease == nil {
				m.Lease = &cockroach_roachpb1.Lease{}
			}
			if err := m.Lease.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TruncatedState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TruncatedState == nil {
				m.TruncatedState = &cockroach_roachpb4.RaftTruncatedState{}
			}
			if err := m.TruncatedState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GCThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GCThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxnSpanGCThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxnSpanGCThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
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
func (m *RangeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: RangeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RangeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicaState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReplicaState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastIndex", wireType)
			}
			m.LastIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumPending", wireType)
			}
			m.NumPending = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumPending |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumDropped", wireType)
			}
			m.NumDropped = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumDropped |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftLogSize", wireType)
			}
			m.RaftLogSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftLogSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApproximateProposalQuota", wireType)
			}
			m.ApproximateProposalQuota = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApproximateProposalQuota |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowState
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
				next, err := skipState(dAtA[start:])
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
	ErrInvalidLengthState = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/pkg/storage/storagebase/state.proto", fileDescriptorState) }

var fileDescriptorState = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x97, 0x2d, 0x1d, 0xad, 0x33, 0xa0, 0x78, 0x20, 0x45, 0x15, 0x6b, 0xa7, 0x8a, 0xa1,
	0x21, 0x86, 0x83, 0x86, 0xc4, 0x01, 0x71, 0xa1, 0x9b, 0x04, 0x1b, 0x03, 0x8d, 0xac, 0x70, 0xe0,
	0x12, 0xb9, 0x8e, 0x97, 0x46, 0x4b, 0x6c, 0x13, 0xbb, 0xa8, 0xda, 0xa7, 0xe0, 0x23, 0x70, 0xe2,
	0x9b, 0x20, 0xed, 0xb8, 0x23, 0xa7, 0x09, 0xca, 0x85, 0x8f, 0x81, 0xec, 0x24, 0x6b, 0x0a, 0x15,
	0xe2, 0x14, 0xe7, 0xbd, 0xdf, 0x7b, 0xf9, 0xfb, 0xbd, 0x7f, 0xc0, 0x03, 0xc2, 0xc9, 0x49, 0xc6,
	0x31, 0x19, 0x7a, 0xe2, 0x24, 0xf2, 0xa4, 0xe2, 0x19, 0x8e, 0x68, 0xf9, 0x1c, 0x60, 0xa9, 0xcf,
	0x58, 0x51, 0x24, 0x32, 0xae, 0x38, 0x5c, 0xbb, 0xc4, 0x51, 0x81, 0xa0, 0x0a, 0xda, 0x7a, 0x38,
	0xbf, 0x1b, 0x65, 0x51, 0xcc, 0xca, 0x87, 0x18, 0x78, 0xe9, 0x47, 0x42, 0xf2, 0x86, 0xad, 0x7b,
	0xb3, 0x15, 0xe6, 0x24, 0x06, 0x5e, 0xcc, 0x14, 0xcd, 0x18, 0x4e, 0x82, 0x0c, 0x1f, 0xab, 0x02,
	0xbd, 0x33, 0x1f, 0x4d, 0xa9, 0xc2, 0x21, 0x56, 0xb8, 0xa0, 0xd6, 0xe7, 0x53, 0x15, 0xe2, 0xee,
	0x2c, 0x31, 0x52, 0x71, 0xe2, 0x0d, 0x13, 0xe2, 0xa9, 0x38, 0xa5, 0x52, 0xe1, 0x54, 0x14, 0xdc,
	0xcd, 0x88, 0x47, 0xdc, 0x1c, 0x3d, 0x7d, 0xca, 0xa3, 0xdd, 0xaf, 0x36, 0x58, 0xf1, 0xa9, 0x48,
	0x62, 0x82, 0x8f, 0xf4, 0x60, 0xe0, 0x16, 0x80, 0x5a, 0x64, 0x80, 0x85, 0x48, 0x62, 0x1a, 0x06,
	0x31, 0x0b, 0xe9, 0xd8, 0xb5, 0xd6, 0xad, 0x4d, 0xdb, 0x6f, 0xea, 0xcc, 0xb3, 0x3c, 0xb1, 0xa7,
	0xe3, 0x10, 0x81, 0xd5, 0x84, 0x62, 0x49, 0xff, 0xc0, 0x17, 0x0d, 0x7e, 0xc3, 0xa4, 0x66, 0xf8,
	0xc7, 0xc0, 0x0e, 0xa9, 0x24, 0xee, 0xd2, 0xba, 0xb5, 0xe9, 0x6c, 0x77, 0xd1, 0x74, 0xfe, 0xc5,
	0xcd, 0x90, 0x8f, 0x59, 0x44, 0x77, 0xa9, 0x24, 0x59, 0x2c, 0x14, 0xcf, 0x7c, 0xc3, 0x43, 0x04,
	0x6a, 0xa6, 0x99, 0x6b, 0x9b, 0x42, 0x77, 0x4e, 0xe1, 0x81, 0xce, 0xfb, 0x39, 0x06, 0x5f, 0x83,
	0xeb, 0x2a, 0x1b, 0x31, 0x82, 0x15, 0x0d, 0x03, 0xb3, 0x71, 0xb7, 0x66, 0x2a, 0x37, 0xe6, 0x7e,
	0xf2, 0x58, 0xf5, 0x4b, 0xda, 0x4c, 0xc1, 0xbf, 0xa6, 0x66, 0xde, 0xe1, 0x5b, 0xb0, 0x12, 0x91,
	0x40, 0x0d, 0x33, 0x2a, 0x87, 0x3c, 0x09, 0xdd, 0x65, 0xd3, 0x6c, 0xad, 0xd2, 0x4c, 0xcf, 0x1d,
	0x0d, 0x13, 0x82, 0xfa, 0xe5, 0xdc, 0x7b, 0xab, 0x67, 0x17, 0x9d, 0x85, 0xc9, 0x45, 0xc7, 0x79,
	0xbe, 0xd3, 0x2f, 0x2b, 0x7d, 0x27, 0x22, 0x97, 0x2f, 0xf0, 0x05, 0xa8, 0x69, 0x71, 0xd2, 0xbd,
	0x62, 0xfa, 0x6d, 0xa1, 0xbf, 0xfd, 0x98, 0xbb, 0x0c, 0x95, 0x66, 0x43, 0xaf, 0xde, 0xed, 0xec,
	0x68, 0x4d, 0xb2, 0x67, 0xeb, 0xf6, 0x7e, 0xde, 0x00, 0x26, 0xe0, 0x96, 0x1a, 0xb3, 0x40, 0x0a,
	0xcc, 0x82, 0x19, 0xa5, 0x8d, 0xff, 0x51, 0xda, 0x2a, 0x94, 0xc2, 0xfe, 0x98, 0x1d, 0x09, 0xcc,
	0xaa, 0x82, 0xa1, 0x2a, 0x62, 0x53, 0xdd, 0x4f, 0xec, 0x5f, 0x9f, 0x3b, 0xd6, 0xbe, 0x5d, 0xaf,
	0x37, 0x1b, 0xfb, 0x76, 0x1d, 0x34, 0x9d, 0xee, 0x97, 0x45, 0xd0, 0x30, 0xab, 0xdb, 0x63, 0xc7,
	0x1c, 0xbe, 0xcc, 0xef, 0x45, 0x8d, 0x6f, 0x9c, 0xed, 0xfb, 0xe8, 0x9f, 0xff, 0x19, 0xaa, 0x1a,
	0xb0, 0x57, 0xd7, 0x5a, 0xce, 0x2f, 0x3a, 0x56, 0x7e, 0x35, 0x0a, 0x6f, 0x83, 0x46, 0x82, 0xa5,
	0xda, 0xab, 0x38, 0x6b, 0x1a, 0x80, 0x1d, 0xe0, 0xb0, 0x51, 0x1a, 0x08, 0xca, 0xc2, 0x98, 0x45,
	0xc6, 0x58, 0xb6, 0x0f, 0xd8, 0x28, 0x3d, 0xcc, 0x23, 0x25, 0x10, 0x66, 0x5c, 0x08, 0x1a, 0x1a,
	0x1b, 0xe4, 0xc0, 0x6e, 0x1e, 0x81, 0x5d, 0x70, 0xd5, 0x38, 0x3e, 0xe1, 0x51, 0x20, 0xe3, 0x53,
	0x6a, 0x96, 0xbb, 0xe4, 0x3b, 0x3a, 0x78, 0xc0, 0xa3, 0xa3, 0xf8, 0x94, 0xc2, 0xa7, 0xa0, 0x85,
	0x85, 0xc8, 0xf8, 0x38, 0x4e, 0xb1, 0xa2, 0x81, 0xc8, 0xb8, 0xe0, 0x12, 0x27, 0xc1, 0x87, 0x11,
	0x57, 0xd8, 0x6c, 0x6f, 0xc9, 0x77, 0x2b, 0xc4, 0x61, 0x01, 0xbc, 0xd1, 0xf9, 0xcb, 0x71, 0xd9,
	0xcd, 0x5a, 0x6f, 0xe3, 0xec, 0x47, 0x7b, 0xe1, 0x6c, 0xd2, 0xb6, 0xce, 0x27, 0x6d, 0xeb, 0xdb,
	0xa4, 0x6d, 0x7d, 0x9f, 0xb4, 0xad, 0x4f, 0x3f, 0xdb, 0x0b, 0xef, 0x9d, 0xca, 0x48, 0x06, 0xcb,
	0xe6, 0xf7, 0x7c, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x91, 0xeb, 0x20, 0xd1, 0x04, 0x00,
	0x00,
}
