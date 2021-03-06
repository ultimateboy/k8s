// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/pkg/apis/apps/v1beta1/generated.proto
// DO NOT EDIT!

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/pkg/apis/apps/v1beta1/generated.proto

	It has these top-level messages:
		StatefulSet
		StatefulSetList
		StatefulSetSpec
		StatefulSetStatus
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/ericchiang/k8s/api/resource"
import k8s_io_kubernetes_pkg_api_unversioned "github.com/ericchiang/k8s/api/unversioned"
import k8s_io_kubernetes_pkg_api_v1 "github.com/ericchiang/k8s/api/v1"
import _ "github.com/ericchiang/k8s/runtime"
import _ "github.com/ericchiang/k8s/util/intstr"

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

// StatefulSet represents a set of pods with consistent identities.
// Identities are defined as:
//  - Network: A single stable DNS and hostname.
//  - Storage: As many VolumeClaims as requested.
// The StatefulSet guarantees that a given network identity will always
// map to the same storage identity.
type StatefulSet struct {
	// +optional
	Metadata *k8s_io_kubernetes_pkg_api_v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Spec defines the desired identities of pods in this set.
	// +optional
	Spec *StatefulSetSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// Status is the current status of Pods in this StatefulSet. This data
	// may be out of date by some window of time.
	// +optional
	Status           *StatefulSetStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *StatefulSet) Reset()                    { *m = StatefulSet{} }
func (m *StatefulSet) String() string            { return proto.CompactTextString(m) }
func (*StatefulSet) ProtoMessage()               {}
func (*StatefulSet) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *StatefulSet) GetMetadata() *k8s_io_kubernetes_pkg_api_v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *StatefulSet) GetSpec() *StatefulSetSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *StatefulSet) GetStatus() *StatefulSetStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// StatefulSetList is a collection of StatefulSets.
type StatefulSetList struct {
	// +optional
	Metadata         *k8s_io_kubernetes_pkg_api_unversioned.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Items            []*StatefulSet                                  `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte                                          `json:"-"`
}

func (m *StatefulSetList) Reset()                    { *m = StatefulSetList{} }
func (m *StatefulSetList) String() string            { return proto.CompactTextString(m) }
func (*StatefulSetList) ProtoMessage()               {}
func (*StatefulSetList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *StatefulSetList) GetMetadata() *k8s_io_kubernetes_pkg_api_unversioned.ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *StatefulSetList) GetItems() []*StatefulSet {
	if m != nil {
		return m.Items
	}
	return nil
}

// A StatefulSetSpec is the specification of a StatefulSet.
type StatefulSetSpec struct {
	// Replicas is the desired number of replicas of the given Template.
	// These are replicas in the sense that they are instantiations of the
	// same Template, but individual replicas also have a consistent identity.
	// If unspecified, defaults to 1.
	// TODO: Consider a rename of this field.
	// +optional
	Replicas *int32 `protobuf:"varint,1,opt,name=replicas" json:"replicas,omitempty"`
	// Selector is a label query over pods that should match the replica count.
	// If empty, defaulted to labels on the pod template.
	// More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
	// +optional
	Selector *k8s_io_kubernetes_pkg_api_unversioned.LabelSelector `protobuf:"bytes,2,opt,name=selector" json:"selector,omitempty"`
	// Template is the object that describes the pod that will be created if
	// insufficient replicas are detected. Each pod stamped out by the StatefulSet
	// will fulfill this Template, but have a unique identity from the rest
	// of the StatefulSet.
	Template *k8s_io_kubernetes_pkg_api_v1.PodTemplateSpec `protobuf:"bytes,3,opt,name=template" json:"template,omitempty"`
	// VolumeClaimTemplates is a list of claims that pods are allowed to reference.
	// The StatefulSet controller is responsible for mapping network identities to
	// claims in a way that maintains the identity of a pod. Every claim in
	// this list must have at least one matching (by name) volumeMount in one
	// container in the template. A claim in this list takes precedence over
	// any volumes in the template, with the same name.
	// TODO: Define the behavior if a claim already exists with the same name.
	// +optional
	VolumeClaimTemplates []*k8s_io_kubernetes_pkg_api_v1.PersistentVolumeClaim `protobuf:"bytes,4,rep,name=volumeClaimTemplates" json:"volumeClaimTemplates,omitempty"`
	// ServiceName is the name of the service that governs this StatefulSet.
	// This service must exist before the StatefulSet, and is responsible for
	// the network identity of the set. Pods get DNS/hostnames that follow the
	// pattern: pod-specific-string.serviceName.default.svc.cluster.local
	// where "pod-specific-string" is managed by the StatefulSet controller.
	ServiceName      *string `protobuf:"bytes,5,opt,name=serviceName" json:"serviceName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StatefulSetSpec) Reset()                    { *m = StatefulSetSpec{} }
func (m *StatefulSetSpec) String() string            { return proto.CompactTextString(m) }
func (*StatefulSetSpec) ProtoMessage()               {}
func (*StatefulSetSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *StatefulSetSpec) GetReplicas() int32 {
	if m != nil && m.Replicas != nil {
		return *m.Replicas
	}
	return 0
}

func (m *StatefulSetSpec) GetSelector() *k8s_io_kubernetes_pkg_api_unversioned.LabelSelector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *StatefulSetSpec) GetTemplate() *k8s_io_kubernetes_pkg_api_v1.PodTemplateSpec {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *StatefulSetSpec) GetVolumeClaimTemplates() []*k8s_io_kubernetes_pkg_api_v1.PersistentVolumeClaim {
	if m != nil {
		return m.VolumeClaimTemplates
	}
	return nil
}

func (m *StatefulSetSpec) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

// StatefulSetStatus represents the current state of a StatefulSet.
type StatefulSetStatus struct {
	// most recent generation observed by this autoscaler.
	// +optional
	ObservedGeneration *int64 `protobuf:"varint,1,opt,name=observedGeneration" json:"observedGeneration,omitempty"`
	// Replicas is the number of actual replicas.
	Replicas         *int32 `protobuf:"varint,2,opt,name=replicas" json:"replicas,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *StatefulSetStatus) Reset()                    { *m = StatefulSetStatus{} }
func (m *StatefulSetStatus) String() string            { return proto.CompactTextString(m) }
func (*StatefulSetStatus) ProtoMessage()               {}
func (*StatefulSetStatus) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func (m *StatefulSetStatus) GetObservedGeneration() int64 {
	if m != nil && m.ObservedGeneration != nil {
		return *m.ObservedGeneration
	}
	return 0
}

func (m *StatefulSetStatus) GetReplicas() int32 {
	if m != nil && m.Replicas != nil {
		return *m.Replicas
	}
	return 0
}

func init() {
	proto.RegisterType((*StatefulSet)(nil), "github.com/ericchiang.k8s.apis.apps.v1beta1.StatefulSet")
	proto.RegisterType((*StatefulSetList)(nil), "github.com/ericchiang.k8s.apis.apps.v1beta1.StatefulSetList")
	proto.RegisterType((*StatefulSetSpec)(nil), "github.com/ericchiang.k8s.apis.apps.v1beta1.StatefulSetSpec")
	proto.RegisterType((*StatefulSetStatus)(nil), "github.com/ericchiang.k8s.apis.apps.v1beta1.StatefulSetStatus")
}
func (m *StatefulSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatefulSet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Spec != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
		n2, err := m.Spec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Status != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Status.Size()))
		n3, err := m.Status.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StatefulSetList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatefulSetList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Metadata.Size()))
		n4, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StatefulSetSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatefulSetSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Replicas != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.Replicas))
	}
	if m.Selector != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Selector.Size()))
		n5, err := m.Selector.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Template != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Template.Size()))
		n6, err := m.Template.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if len(m.VolumeClaimTemplates) > 0 {
		for _, msg := range m.VolumeClaimTemplates {
			dAtA[i] = 0x22
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.ServiceName != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.ServiceName)))
		i += copy(dAtA[i:], *m.ServiceName)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StatefulSetStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatefulSetStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ObservedGeneration != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.ObservedGeneration))
	}
	if m.Replicas != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.Replicas))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StatefulSet) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StatefulSetList) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StatefulSetSpec) Size() (n int) {
	var l int
	_ = l
	if m.Replicas != nil {
		n += 1 + sovGenerated(uint64(*m.Replicas))
	}
	if m.Selector != nil {
		l = m.Selector.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Template != nil {
		l = m.Template.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.VolumeClaimTemplates) > 0 {
		for _, e := range m.VolumeClaimTemplates {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.ServiceName != nil {
		l = len(*m.ServiceName)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StatefulSetStatus) Size() (n int) {
	var l int
	_ = l
	if m.ObservedGeneration != nil {
		n += 1 + sovGenerated(uint64(*m.ObservedGeneration))
	}
	if m.Replicas != nil {
		n += 1 + sovGenerated(uint64(*m.Replicas))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StatefulSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: StatefulSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatefulSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &k8s_io_kubernetes_pkg_api_v1.ObjectMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &StatefulSetSpec{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &StatefulSetStatus{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *StatefulSetList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: StatefulSetList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatefulSetList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &k8s_io_kubernetes_pkg_api_unversioned.ListMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &StatefulSet{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *StatefulSetSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: StatefulSetSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatefulSetSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Replicas = &v
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Selector == nil {
				m.Selector = &k8s_io_kubernetes_pkg_api_unversioned.LabelSelector{}
			}
			if err := m.Selector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Template", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Template == nil {
				m.Template = &k8s_io_kubernetes_pkg_api_v1.PodTemplateSpec{}
			}
			if err := m.Template.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VolumeClaimTemplates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VolumeClaimTemplates = append(m.VolumeClaimTemplates, &k8s_io_kubernetes_pkg_api_v1.PersistentVolumeClaim{})
			if err := m.VolumeClaimTemplates[len(m.VolumeClaimTemplates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.ServiceName = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *StatefulSetStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: StatefulSetStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatefulSetStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObservedGeneration", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
			m.ObservedGeneration = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Replicas = &v
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/ericchiang/k8s/apis/apps/v1beta1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc7, 0x49, 0x7b, 0x85, 0xe2, 0x0e, 0x08, 0x8b, 0x21, 0x74, 0xa8, 0xaa, 0x2e, 0x74, 0xe0,
	0x6c, 0xb5, 0x1c, 0xe2, 0xc4, 0x08, 0x48, 0x08, 0x38, 0xe0, 0xe4, 0x22, 0x06, 0x16, 0xe4, 0x24,
	0x8f, 0xca, 0x34, 0x89, 0x2d, 0xfb, 0x25, 0x9f, 0x85, 0xcf, 0xc0, 0x27, 0x61, 0x64, 0x61, 0x47,
	0xe5, 0x33, 0xb0, 0xa3, 0xa4, 0xa1, 0xd7, 0x36, 0xcd, 0xdd, 0xa9, 0x63, 0xe2, 0xf7, 0xfb, 0xe5,
	0xbd, 0xf7, 0x77, 0xc8, 0x93, 0xc5, 0xa9, 0x63, 0x4a, 0xf3, 0x45, 0x16, 0x80, 0x4d, 0x01, 0xc1,
	0x71, 0xb3, 0x98, 0x73, 0x69, 0x94, 0xe3, 0xd2, 0x18, 0xc7, 0xf3, 0x49, 0x00, 0x28, 0x27, 0x7c,
	0x0e, 0x29, 0x58, 0x89, 0x10, 0x31, 0x63, 0x35, 0x6a, 0xfa, 0x60, 0x05, 0xb2, 0x0b, 0x90, 0x99,
	0xc5, 0x9c, 0x15, 0x20, 0x2b, 0x40, 0x56, 0x81, 0xfd, 0x69, 0xe3, 0x17, 0xb8, 0x05, 0xa7, 0x33,
	0x1b, 0xc2, 0xae, 0xbc, 0xff, 0xb8, 0x99, 0xc9, 0xd2, 0x1c, 0xac, 0x53, 0x3a, 0x85, 0xa8, 0x86,
	0x3d, 0x6c, 0xc6, 0xf2, 0xda, 0x04, 0xfd, 0xe3, 0xfd, 0xd5, 0x36, 0x4b, 0x51, 0x25, 0xf5, 0x9e,
	0x26, 0xfb, 0xcb, 0x33, 0x54, 0x31, 0x57, 0x29, 0x3a, 0xb4, 0xbb, 0xc8, 0xe8, 0xaf, 0x47, 0x7a,
	0x33, 0x94, 0x08, 0x5f, 0xb2, 0x78, 0x06, 0x48, 0x5f, 0x90, 0x6e, 0x02, 0x28, 0x23, 0x89, 0xd2,
	0xf7, 0x86, 0xde, 0xb8, 0x37, 0x1d, 0xb3, 0xc6, 0x35, 0xb2, 0x7c, 0xc2, 0xde, 0x07, 0x5f, 0x21,
	0xc4, 0xb7, 0x80, 0x52, 0xac, 0x49, 0x7a, 0x46, 0x8e, 0x9c, 0x81, 0xd0, 0x6f, 0x95, 0x86, 0x53,
	0x76, 0xcd, 0x20, 0xd8, 0x46, 0x27, 0x33, 0x03, 0xa1, 0x28, 0x2d, 0x54, 0x90, 0x9b, 0x0e, 0x25,
	0x66, 0xce, 0x6f, 0x97, 0xbe, 0xa7, 0x07, 0xf9, 0x4a, 0x83, 0xa8, 0x4c, 0xa3, 0xef, 0x1e, 0xb9,
	0xb3, 0x71, 0x7a, 0xa6, 0x1c, 0xd2, 0x37, 0xb5, 0xd9, 0xf9, 0x25, 0xb3, 0x6f, 0xa4, 0xcc, 0x0a,
	0x7c, 0x67, 0x05, 0xaf, 0x49, 0x47, 0x21, 0x24, 0xce, 0x6f, 0x0d, 0xdb, 0xe3, 0xde, 0xf4, 0xe4,
	0x90, 0x9e, 0xc5, 0x4a, 0x31, 0xfa, 0xd5, 0xda, 0x6a, 0xb6, 0x58, 0x0d, 0xed, 0x93, 0xae, 0x05,
	0x13, 0xab, 0x50, 0xba, 0xb2, 0xd9, 0x8e, 0x58, 0x3f, 0xd3, 0x73, 0xd2, 0x75, 0x10, 0x43, 0x88,
	0xda, 0x56, 0x11, 0x9c, 0x5c, 0x77, 0x10, 0x19, 0x40, 0x3c, 0xab, 0x58, 0xb1, 0xb6, 0xd0, 0x57,
	0xa4, 0x8b, 0x90, 0x98, 0x58, 0x22, 0x54, 0x21, 0x1c, 0x5f, 0x7e, 0x2d, 0xce, 0x75, 0xf4, 0xa1,
	0x02, 0xca, 0x24, 0xd7, 0x38, 0x9d, 0x93, 0x7b, 0xb9, 0x8e, 0xb3, 0x04, 0x9e, 0xc7, 0x52, 0x25,
	0xff, 0x8b, 0x9c, 0x7f, 0x54, 0xee, 0xe9, 0xd1, 0x15, 0xda, 0xa2, 0x53, 0x87, 0x90, 0xe2, 0xc7,
	0x0b, 0x87, 0xd8, 0x2b, 0xa4, 0x43, 0xd2, 0x73, 0x60, 0x73, 0x15, 0xc2, 0x3b, 0x99, 0x80, 0xdf,
	0x19, 0x7a, 0xe3, 0xdb, 0x62, 0xf3, 0xd5, 0xe8, 0x33, 0xb9, 0x5b, 0xbb, 0x21, 0x94, 0x11, 0xaa,
	0x83, 0xa2, 0x0a, 0xa2, 0x97, 0xab, 0x9f, 0x45, 0xe9, 0xb4, 0x5c, 0x71, 0x5b, 0xec, 0x39, 0xd9,
	0x0a, 0xa2, 0xb5, 0x1d, 0xc4, 0xb3, 0xfb, 0x3f, 0x96, 0x03, 0xef, 0xe7, 0x72, 0xe0, 0xfd, 0x5e,
	0x0e, 0xbc, 0x6f, 0x7f, 0x06, 0x37, 0x3e, 0xdd, 0xaa, 0x62, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff,
	0x76, 0xa6, 0x33, 0xba, 0xd6, 0x04, 0x00, 0x00,
}
