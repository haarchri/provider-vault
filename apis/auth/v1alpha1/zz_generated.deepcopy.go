//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backend) DeepCopyInto(out *Backend) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backend.
func (in *Backend) DeepCopy() *Backend {
	if in == nil {
		return nil
	}
	out := new(Backend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Backend) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendInitParameters) DeepCopyInto(out *BackendInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableRemount != nil {
		in, out := &in.DisableRemount, &out.DisableRemount
		*out = new(bool)
		**out = **in
	}
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(bool)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Tune != nil {
		in, out := &in.Tune, &out.Tune
		*out = make([]TuneInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendInitParameters.
func (in *BackendInitParameters) DeepCopy() *BackendInitParameters {
	if in == nil {
		return nil
	}
	out := new(BackendInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendList) DeepCopyInto(out *BackendList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Backend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendList.
func (in *BackendList) DeepCopy() *BackendList {
	if in == nil {
		return nil
	}
	out := new(BackendList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackendList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendObservation) DeepCopyInto(out *BackendObservation) {
	*out = *in
	if in.Accessor != nil {
		in, out := &in.Accessor, &out.Accessor
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableRemount != nil {
		in, out := &in.DisableRemount, &out.DisableRemount
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(bool)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Tune != nil {
		in, out := &in.Tune, &out.Tune
		*out = make([]TuneObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendObservation.
func (in *BackendObservation) DeepCopy() *BackendObservation {
	if in == nil {
		return nil
	}
	out := new(BackendObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendParameters) DeepCopyInto(out *BackendParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableRemount != nil {
		in, out := &in.DisableRemount, &out.DisableRemount
		*out = new(bool)
		**out = **in
	}
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(bool)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Tune != nil {
		in, out := &in.Tune, &out.Tune
		*out = make([]TuneParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendParameters.
func (in *BackendParameters) DeepCopy() *BackendParameters {
	if in == nil {
		return nil
	}
	out := new(BackendParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendSpec) DeepCopyInto(out *BackendSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendSpec.
func (in *BackendSpec) DeepCopy() *BackendSpec {
	if in == nil {
		return nil
	}
	out := new(BackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendStatus) DeepCopyInto(out *BackendStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendStatus.
func (in *BackendStatus) DeepCopy() *BackendStatus {
	if in == nil {
		return nil
	}
	out := new(BackendStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TuneInitParameters) DeepCopyInto(out *TuneInitParameters) {
	*out = *in
	if in.AllowedResponseHeaders != nil {
		in, out := &in.AllowedResponseHeaders, &out.AllowedResponseHeaders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AuditNonHMACRequestKeys != nil {
		in, out := &in.AuditNonHMACRequestKeys, &out.AuditNonHMACRequestKeys
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AuditNonHMACResponseKeys != nil {
		in, out := &in.AuditNonHMACResponseKeys, &out.AuditNonHMACResponseKeys
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DefaultLeaseTTL != nil {
		in, out := &in.DefaultLeaseTTL, &out.DefaultLeaseTTL
		*out = new(string)
		**out = **in
	}
	if in.ListingVisibility != nil {
		in, out := &in.ListingVisibility, &out.ListingVisibility
		*out = new(string)
		**out = **in
	}
	if in.MaxLeaseTTL != nil {
		in, out := &in.MaxLeaseTTL, &out.MaxLeaseTTL
		*out = new(string)
		**out = **in
	}
	if in.PassthroughRequestHeaders != nil {
		in, out := &in.PassthroughRequestHeaders, &out.PassthroughRequestHeaders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenType != nil {
		in, out := &in.TokenType, &out.TokenType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TuneInitParameters.
func (in *TuneInitParameters) DeepCopy() *TuneInitParameters {
	if in == nil {
		return nil
	}
	out := new(TuneInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TuneObservation) DeepCopyInto(out *TuneObservation) {
	*out = *in
	if in.AllowedResponseHeaders != nil {
		in, out := &in.AllowedResponseHeaders, &out.AllowedResponseHeaders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AuditNonHMACRequestKeys != nil {
		in, out := &in.AuditNonHMACRequestKeys, &out.AuditNonHMACRequestKeys
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AuditNonHMACResponseKeys != nil {
		in, out := &in.AuditNonHMACResponseKeys, &out.AuditNonHMACResponseKeys
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DefaultLeaseTTL != nil {
		in, out := &in.DefaultLeaseTTL, &out.DefaultLeaseTTL
		*out = new(string)
		**out = **in
	}
	if in.ListingVisibility != nil {
		in, out := &in.ListingVisibility, &out.ListingVisibility
		*out = new(string)
		**out = **in
	}
	if in.MaxLeaseTTL != nil {
		in, out := &in.MaxLeaseTTL, &out.MaxLeaseTTL
		*out = new(string)
		**out = **in
	}
	if in.PassthroughRequestHeaders != nil {
		in, out := &in.PassthroughRequestHeaders, &out.PassthroughRequestHeaders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenType != nil {
		in, out := &in.TokenType, &out.TokenType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TuneObservation.
func (in *TuneObservation) DeepCopy() *TuneObservation {
	if in == nil {
		return nil
	}
	out := new(TuneObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TuneParameters) DeepCopyInto(out *TuneParameters) {
	*out = *in
	if in.AllowedResponseHeaders != nil {
		in, out := &in.AllowedResponseHeaders, &out.AllowedResponseHeaders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AuditNonHMACRequestKeys != nil {
		in, out := &in.AuditNonHMACRequestKeys, &out.AuditNonHMACRequestKeys
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AuditNonHMACResponseKeys != nil {
		in, out := &in.AuditNonHMACResponseKeys, &out.AuditNonHMACResponseKeys
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DefaultLeaseTTL != nil {
		in, out := &in.DefaultLeaseTTL, &out.DefaultLeaseTTL
		*out = new(string)
		**out = **in
	}
	if in.ListingVisibility != nil {
		in, out := &in.ListingVisibility, &out.ListingVisibility
		*out = new(string)
		**out = **in
	}
	if in.MaxLeaseTTL != nil {
		in, out := &in.MaxLeaseTTL, &out.MaxLeaseTTL
		*out = new(string)
		**out = **in
	}
	if in.PassthroughRequestHeaders != nil {
		in, out := &in.PassthroughRequestHeaders, &out.PassthroughRequestHeaders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenType != nil {
		in, out := &in.TokenType, &out.TokenType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TuneParameters.
func (in *TuneParameters) DeepCopy() *TuneParameters {
	if in == nil {
		return nil
	}
	out := new(TuneParameters)
	in.DeepCopyInto(out)
	return out
}
