// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1/trustbundle.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using TrustBundle within kubernetes types, where deepcopy-gen is used.
func (in *TrustBundle) DeepCopyInto(out *TrustBundle) {
	p := proto.Clone(in).(*TrustBundle)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrustBundle. Required by controller-gen.
func (in *TrustBundle) DeepCopy() *TrustBundle {
	if in == nil {
		return nil
	}
	out := new(TrustBundle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TrustBundle. Required by controller-gen.
func (in *TrustBundle) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TrustBundleRequest within kubernetes types, where deepcopy-gen is used.
func (in *TrustBundleRequest) DeepCopyInto(out *TrustBundleRequest) {
	p := proto.Clone(in).(*TrustBundleRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrustBundleRequest. Required by controller-gen.
func (in *TrustBundleRequest) DeepCopy() *TrustBundleRequest {
	if in == nil {
		return nil
	}
	out := new(TrustBundleRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TrustBundleRequest. Required by controller-gen.
func (in *TrustBundleRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TrustBundleResponse within kubernetes types, where deepcopy-gen is used.
func (in *TrustBundleResponse) DeepCopyInto(out *TrustBundleResponse) {
	p := proto.Clone(in).(*TrustBundleResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrustBundleResponse. Required by controller-gen.
func (in *TrustBundleResponse) DeepCopy() *TrustBundleResponse {
	if in == nil {
		return nil
	}
	out := new(TrustBundleResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TrustBundleResponse. Required by controller-gen.
func (in *TrustBundleResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
