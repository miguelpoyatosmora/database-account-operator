//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLAccount) DeepCopyInto(out *PostgreSQLAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLAccount.
func (in *PostgreSQLAccount) DeepCopy() *PostgreSQLAccount {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLAccountList) DeepCopyInto(out *PostgreSQLAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgreSQLAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLAccountList.
func (in *PostgreSQLAccountList) DeepCopy() *PostgreSQLAccountList {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLAccountSpec) DeepCopyInto(out *PostgreSQLAccountSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLAccountSpec.
func (in *PostgreSQLAccountSpec) DeepCopy() *PostgreSQLAccountSpec {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLAccountStatus) DeepCopyInto(out *PostgreSQLAccountStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLAccountStatus.
func (in *PostgreSQLAccountStatus) DeepCopy() *PostgreSQLAccountStatus {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLDatabase) DeepCopyInto(out *PostgreSQLDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLDatabase.
func (in *PostgreSQLDatabase) DeepCopy() *PostgreSQLDatabase {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLDatabaseList) DeepCopyInto(out *PostgreSQLDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgreSQLDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLDatabaseList.
func (in *PostgreSQLDatabaseList) DeepCopy() *PostgreSQLDatabaseList {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLDatabaseSpec) DeepCopyInto(out *PostgreSQLDatabaseSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLDatabaseSpec.
func (in *PostgreSQLDatabaseSpec) DeepCopy() *PostgreSQLDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLDatabaseStatus) DeepCopyInto(out *PostgreSQLDatabaseStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLDatabaseStatus.
func (in *PostgreSQLDatabaseStatus) DeepCopy() *PostgreSQLDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLGrant) DeepCopyInto(out *PostgreSQLGrant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLGrant.
func (in *PostgreSQLGrant) DeepCopy() *PostgreSQLGrant {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLGrant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLGrant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLGrantList) DeepCopyInto(out *PostgreSQLGrantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgreSQLGrant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLGrantList.
func (in *PostgreSQLGrantList) DeepCopy() *PostgreSQLGrantList {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLGrantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLGrantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLGrantSpec) DeepCopyInto(out *PostgreSQLGrantSpec) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLGrantSpec.
func (in *PostgreSQLGrantSpec) DeepCopy() *PostgreSQLGrantSpec {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLGrantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLGrantStatus) DeepCopyInto(out *PostgreSQLGrantStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLGrantStatus.
func (in *PostgreSQLGrantStatus) DeepCopy() *PostgreSQLGrantStatus {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLGrantStatus)
	in.DeepCopyInto(out)
	return out
}